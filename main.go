package main

import (
	"fmt"
	"github.com/takama/router"
	"net/http"
	"os"
	"log"
)

// Run server: go run main.go
// Try few requests:
// - curl http://localhost:8000
// - curl http://localhost:8000/test-some-path
func main() {

	logger := log.New(os.Stdout, "[step-by-step]", log.LstdFlags)

	logger.Print("asd")

	port := os.Getenv("SERVICE_PORT")

	if len(port) == 0 {
		logger.Fatal("Порт не задан")
	}
	r := router.New()
	r.GET("/home", home)
	r.GET("/mw", mw(5))
	r.GET("/mw2", mw2(5, abc))
	r.GET("/abc/:num", abc)

	r.Listen(":"+port)
}

func home(c *router.Control) {
	fmt.Fprintf(c.Writer, "URL.Path = %q\n", c.Request.URL.Path)

}

func mw(a int) router.Handle {
	return func(c *router.Control) {
		fmt.Fprintf(c.Writer, "a = %v\n", a)
	}
}

func mw2(a int, h router.Handle) router.Handle {

	return h
}
func abc(c *router.Control) {
	fmt.Fprintf(c.Writer, "a = %v\n", c.Get(":num"))
}


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
