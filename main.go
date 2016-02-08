package main

import (
	"github.com/bertsinnema/vcloud/was"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		IndentJSON: true,
		IndentXML:  true,
	}))

	//Routes
	m.Get("/", was.ListContainers)

	m.RunOnAddr("localhost:8081")
}
