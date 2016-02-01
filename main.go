package main

import (
	"vcloud/was"

	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()

	//Routes
	m.Post("/listcontainers", was.ListContainers)

	m.RunOnAddr("localhost:8081")
}
