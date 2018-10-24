package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"path/filepath"

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
	dirname := os.Getenv(AppLogDirectory)
	if dirname == "" {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("APP_LOG_DIR environment not defined.")
		return
	}
	if _, err := os.Stat(dirname); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("APP_LOG_DIR is not found.:%v", err)
		return
	}
	w.WriteJson("ok")
}

// log is application logging.
func saveGreeting(message *string) error {
	logMutex.Lock()
	defer logMutex.Unlock()

	dirname := os.Getenv(AppLogDirectory)
	filename := filepath.Join(dirname, "app.log")
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("saveGreeting error:%v", err)
	}
	defer file.Close()
	fmt.Fprintln(file, *message)
	return nil
}

// Greeting response structure.
type Greeting struct {
	Message string
}

// Gretting reply.
func hello(w rest.ResponseWriter, r *rest.Request) {
	id := r.URL.Query()["id"]
	log.Printf("id is %v", id)

	if id != nil && len(id) > 0 {
		saveGreeting(&id[0])
	}

	message := "Hello"
	greeting := Greeting{Message: message}
	w.WriteJson(greeting)
}
