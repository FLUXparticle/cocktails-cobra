package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func resetRatings(db *gorm.DB) {
	db.Exec("DELETE FROM ratings")
	db.Exec("DELETE FROM sqlite_sequence WHERE name = 'ratings'")
}

// setupTestHandler stellt über Fx die Handler-Instanz und die Gorm-Datenbank bereit,
// sodass sie in den Tests verwendet werden können.
func setupTestHandler(t *testing.T) http.Handler {
	var handler http.Handler

	constructors := baseConstructors()

	app := fx.New(
		fx.Provide(constructors...),
		fx.Invoke(resetRatings),
		fx.Populate(&handler),
	)

	// Starte die Anwendung und blockiere bei Fehlern
	{
		startCtx, cancel := context.WithTimeout(context.Background(), fx.DefaultTimeout)
		defer cancel()
		if err := app.Start(startCtx); err != nil {
			t.Fatalf("Failed to start Fx app: %v", err)
		}
	}

	t.Cleanup(func() {
		stopCtx, cancel := context.WithTimeout(context.Background(), fx.DefaultTimeout)
		defer cancel()
		if err := app.Stop(stopCtx); err != nil {
			t.Fatalf("Failed to stop Fx app: %v", err)
		}
	})

	return handler
}

func TestCocktailListHandler(t *testing.T) {
	// Gin-Router über setupTestHandler bereitstellen
	handler := setupTestHandler(t)

	// Erstelle eine neue HTTP-Anfrage für /cocktails
	req := httptest.NewRequest("GET", "/cocktails", nil)
	w := httptest.NewRecorder()

	// Anfrage über den Handler ausführen
	handler.ServeHTTP(w, req)

	// Überprüfe die Antwort
	resp := w.Result()
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, string(body), "Yellow Boxer")
}

func TestCocktailDetailsHandler(t *testing.T) {
	handler := setupTestHandler(t)

	req := httptest.NewRequest("GET", "/cocktails/1", nil)
	w := httptest.NewRecorder()

	// Anfrage über den Handler ausführen
	handler.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, string(body), "Orange Velvet")
}

func TestCocktailListWithIngredientHandler(t *testing.T) {
	handler := setupTestHandler(t)

	// Erstelle eine neue HTTP-Anfrage für /cocktails?ingredient=Milch
	req := httptest.NewRequest("GET", "/cocktails?ingredient=Milch", nil)
	w := httptest.NewRecorder()

	// Anfrage über den Handler ausführen
	handler.ServeHTTP(w, req)

	// Überprüfe die Antwort
	resp := w.Result()
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, string(body), "Milkshake")
	assert.NotContains(t, string(body), "Mandarinetto")
}

func TestIngredientsListHandler(t *testing.T) {
	handler := setupTestHandler(t)

	// Erstelle eine neue HTTP-Anfrage für /ingredients
	req := httptest.NewRequest("GET", "/ingredients", nil)
	w := httptest.NewRecorder()

	// Anfrage über den Handler ausführen
	handler.ServeHTTP(w, req)

	// Überprüfe die Antwort
	resp := w.Result()
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// Erwarteter Statuscode
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Beispiel: Überprüfe, ob einige bekannte Zutaten enthalten sind
	assert.Contains(t, string(body), "Milch")
	assert.Contains(t, string(body), "Orangensaft")
	assert.Contains(t, string(body), "Sahne")
}

func TestAddRatingHandler(t *testing.T) {
	handler := setupTestHandler(t)

	ratingJSON := `{"score": 5}`
	req := httptest.NewRequest("POST", "/ratings/cocktails/1", strings.NewReader(ratingJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestAddMultipleRatingsAndAverage(t *testing.T) {
	handler := setupTestHandler(t)

	// Ratings, die hinzugefügt werden sollen
	ratings := []int{5, 3, 4}
	expectedAverage := 4.0

	// Füge mehrere Ratings für den Cocktail mit der ID 1 hinzu
	for _, score := range ratings {
		ratingJSON := fmt.Sprintf(`{"score": %d}`, score)
		req := httptest.NewRequest("POST", "/ratings/cocktails/1", strings.NewReader(ratingJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		resp := w.Result()
		resp.Body.Close()
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	}

	// Testen des Durchschnittsendpoints
	req := httptest.NewRequest("GET", "/ratings/cocktails/1/average", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	// Antwort auslesen und prüfen
	body, _ := io.ReadAll(resp.Body)

	// Parse JSON-Antwort
	var result map[string]float64
	err := json.Unmarshal(body, &result)
	assert.NoError(t, err, "Fehler beim Parsen der JSON-Antwort")

	// Durchschnittswert überprüfen
	assert.Equal(t, expectedAverage, result["average_rating"], "Durchschnittswert ist nicht korrekt")
}
