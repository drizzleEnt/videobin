package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"videobin/internal/api"
	"videobin/internal/api/filectrl"
	"videobin/internal/middleware"
	"videobin/internal/repository"
	"videobin/internal/routes"
	"videobin/internal/service"
	"videobin/internal/service/filesrv"
)

type App struct {
	httpServer *http.Server

	database             repository.DatabaseStorage
	fileStorage          repository.FileStorage
	fileService          service.FileService
	fileController       api.FileController
	middlewareController middleware.Middleware
}

func New(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDebs(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run() error {
	go func() {
		if err := a.runHTTPServer(); err != nil {
			if err == http.ErrServerClosed {
				return
			}
			log.Printf("failed to run HTTP server: %v", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	log.Printf("Shutting down http server...\n")
	if err := a.httpServer.Shutdown(ctx); err != nil {
		log.Printf("failed Shutting down http server %s.\n", err.Error())
	}
	log.Printf("HTTP server stopped.\n")

	return nil
}

func (a *App) initDebs(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}

	}

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	engine := routes.InitRoutes(a.FileController(ctx), a.Middleware(ctx))

	srv := http.Server{
		Addr:           ":8080",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	a.httpServer = &srv

	return nil
}

func (a *App) runHTTPServer() error {
	log.Printf("server run on :8080")
	err := a.httpServer.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}

func (a *App) Database(ctx context.Context) repository.DatabaseStorage {
	if a.database == nil {
		//a.database =
	}

	return a.database
}

func (a *App) FileStorage(ctx context.Context) repository.FileStorage {
	if a.fileStorage == nil {
		//a.fileStorage =
	}

	return a.fileStorage
}

func (a *App) FileService(ctx context.Context) service.FileService {
	if a.fileService == nil {
		a.fileService = filesrv.New()
	}

	return a.fileService
}

func (a *App) FileController(ctx context.Context) api.FileController {
	if a.fileController == nil {
		a.fileController = filectrl.New(a.FileService(ctx))
	}

	return a.fileController
}

func (a *App) Middleware(ctx context.Context) middleware.Middleware {
	if a.middlewareController == nil {
		a.middlewareController = middleware.New()
	}

	return a.middlewareController
}
