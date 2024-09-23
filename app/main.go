package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	name := os.Getenv("NAME")
	if len(name) == 0 {
		name = "chris"
	}
	hello := fmt.Sprintf("Hello %s ", name)
	http.Handle("/hello/", http.StripPrefix("/hello/", http.FileServer(http.Dir("static"))))

	f, err := os.OpenFile("./static/index.html", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(hello); err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}
