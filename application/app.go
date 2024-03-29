package application

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() (*App) { 
	return &App { 
		router: loadRoutes(),
	}
}

func (a *App) Start(ctx context.Context) error { 
	server := &http.Server { 
		Addr: ":8081",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil { 
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}