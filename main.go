package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/unrolled/render"
)

var rend *render.Render

func serverInit() (server *http.Server) {
	rend = render.New()
	port := ":8081"
	fmt.Println("livlon server started on port " + port + "!")

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/hello", helloHandler)

	server = &http.Server{
		Addr:    port,
		Handler: router, // instead of stl - http.HandlerFunc(handler)
	}
	return
}

func main() {
	server := serverInit()

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to start to server", err)
	}
}

func helloHandler(wri http.ResponseWriter, req *http.Request) {
	res := map[string]string{"hello": "json"}
	rend.JSON(wri, http.StatusOK, res)
}
