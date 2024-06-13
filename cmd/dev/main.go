package main

import (
	"fmt"

	"github.com/leapkit/core/gloves"
	"map/internal"
)

func main() {
	err := gloves.Start(
		"cmd/app/main.go",

		internal.GlovesOptions...,
	)

	if err != nil {
		fmt.Println(err)
	}
}
