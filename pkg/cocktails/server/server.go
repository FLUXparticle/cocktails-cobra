package server

import (
	"context"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"moul.io/chizap"
	"moul.io/zapgorm2"
	"net"
	"net/http"
	"time"
)

func NewChiHandler(handler *CocktailsHandler, log *zap.Logger) http.Handler {
	r := chi.NewRouter()
	r.Use(chizap.New(log, &chizap.Opts{
		WithReferer:   true,
		WithUserAgent: true,
	}))
	r.Get("/cocktails", handler.CocktailHandler)
	return r
}

func NewGinHandler(handler *CocktailsHandler, log *zap.Logger) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(ginzap.Ginzap(log, time.RFC3339, true))
	r.Use(gin.Recovery())

	r.GET("/cocktails", handler.CocktailList)
	r.GET("/cocktails/:id", handler.CocktailDetails)

	return r
}

func NewHTTPServer(lc fx.Lifecycle, handler http.Handler, log *zap.Logger) *http.Server {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			//fmt.Println("Starting HTTP server at", srv.Addr)
			log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func NewZapLogger() *zap.Logger {
	//zapLogger, _ := zap.NewProduction(zap.AddCaller() /*, zap.AddCallerSkip(1)*/)
	//return zapLogger
	return zap.NewExample(zap.AddCaller() /*, zap.AddCallerSkip(1)*/)
}

func NewDatabase(log *zap.Logger) *gorm.DB {
	cfg := &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
		Logger: zapgorm2.New(log).LogMode(logger.Info),
	}
	db, err := gorm.Open(sqlite.Open("cocktails.db"), cfg)
	if err != nil {
		panic(err)
	}
	return db
}

func RunServer(useChi bool) {
	constructors := []any{
		NewHTTPServer,
		NewCocktailsHandler,
		NewZapLogger,
		NewDatabase,
	}

	if useChi {
		constructors = append(constructors, NewChiHandler)
	} else {
		constructors = append(constructors, NewGinHandler)
	}

	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(constructors...),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
