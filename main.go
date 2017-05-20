package main

import (
	"fmt"
	"github.com/takama/router"
	"net/http"
)

// Run server: go run main.go
// Try few requests:
// - curl http://localhost:8000
// - curl http://localhost:8000/test-some-path
func main() {

	r := router.New()
	r.GET("/home", home)
	r.GET("/mw", mw(5))
	r.GET("/mw2", mw2(5, abc))

	r.Listen(":8000")
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
	fmt.Fprintf(c.Writer, "a = %v\n", 6)
}


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
