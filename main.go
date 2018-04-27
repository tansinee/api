package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ndidplatform/api/rp"
	"github.com/tylerb/graceful"
)

// Health ckecking Handler
func well(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func main() {
	// Echo instance
	e := echo.New()
	// e.File("/favicon.ico", "images/favicon.ico")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", well)

	e.POST(rp.CreatePath, rp.Create)

	// Start server
	e.Server.Addr = ":1323"

	fmt.Printf("⇨ http server started on %s\n", e.Server.Addr)
	// Serve it like a boss
	graceful.ListenAndServe(e.Server, 5*time.Second)
}
