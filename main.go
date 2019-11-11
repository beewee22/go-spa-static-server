package main

import (
	"flag"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dir := flag.String("dir", "dist", "Directory to serve")
	port := flag.String("port", "80", "Port")
	if !flag.Parsed() {
		flag.Parse()
	}
	// Echo instance
	e := echo.New()

	//#region Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// HSTS
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		HSTSMaxAge: 3600,
	}))
	e.Use(middleware.Static(*dir))
	// #endregion

	// Routes
	e.File("*", *dir+"/index.html")
	// e.Static("/", "dist")

	// Start server
	e.Logger.Fatal(e.Start(":" + *port))
}
