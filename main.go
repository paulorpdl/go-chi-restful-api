package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kelseyhightower/envconfig"

	"github.com/paulorpdl/go-chi-restful-api/routes"
)

func main() {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	config := &Config{}

	envconfig.Process("SERVER", config)

	log.Printf("Starting up on http://%s:%s", config.Addr, config.Port)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get(config.Path, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Mount(fmt.Sprintf("%s/%s", config.Path, "posts"), routes.PostsResource{}.Routes())

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.Addr, config.Port), r))
}
