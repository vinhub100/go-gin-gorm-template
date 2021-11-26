package main

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/multitemplate"

	"path/filepath"

	"blog-app/views"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.HTMLRender = loadTemplates("./templates")

	homeRoute := r.Group("/")
	{
		homeRoute.GET("/", views.Home)
		homeRoute.GET("/hello", views.Hello)
	}

	r.Run(":8080")
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	components, err := filepath.Glob(templatesDir + "/components/*.html")
	if err != nil {
		panic(err.Error())
	}

	views, err := filepath.Glob(templatesDir + "/views/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our components/ and views/ directories
	for _, view := range views {
		layoutCopy := make([]string, len(components))
		copy(layoutCopy, components)
		files := append(layoutCopy, view)
		r.AddFromFiles(filepath.Base(view), files...)
	}
	return r
}
