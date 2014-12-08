package main

import "github.com/go-martini/martini"

func main() {
	server := martini.Classic()

	server.Get("/", func() string {
		return "<h1>140Pl.us</h1><p>Hello 140Pl.us!</p>"
	})

	server.Get("/:name", func(args martini.Params) string {
		return "Hello " + args["name"] + "!"
	})

	server.Run()
}
