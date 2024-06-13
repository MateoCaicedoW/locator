package main

import (
	"fmt"
	"net/http"
	"os"

	"map/internal"
	"map/internal/ai_integration"

	"github.com/leapkit/core/envor"
	"github.com/leapkit/core/server"
)

func main() {
	s := server.New(
		server.WithHost(envor.Get("HOST", "0.0.0.0")),
		server.WithPort(envor.Get("PORT", "3000")),
	)

	s.Use(server.InCtxMiddleware("aiGenerator", ai_integration.Generator))
	if err := internal.AddRoutes(s); err != nil {
		os.Exit(1)
	}

	fmt.Println("Server started at", s.Addr())
	err := http.ListenAndServe(s.Addr(), s.Handler())
	if err != nil {
		fmt.Println(err)
	}
}
