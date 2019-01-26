package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/ant0ine/go-json-rest/rest"
)

const (
	// AppLogDirectory is container-shared test file directory.
	AppLogDirectory = "APP_LOG_DIR"
)

var (
	logMutex = new(sync.Mutex)
)

// main logic.
func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("application starting...")

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/hello", hello),
		rest.Get("/health", health),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))

}

// Health Check API
func health(w rest.ResponseWriter, req *rest.Request) {
	w.WriteJson("ok")
}

// Greeting response structure.
type Greeting struct {
	Message string
}

// Gretting reply.
func hello(w rest.ResponseWriter, r *rest.Request) {
	message := "Hello world!This is next version."
	greeting := Greeting{Message: message}
	w.WriteJson(greeting)
}
