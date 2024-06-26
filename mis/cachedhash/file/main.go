package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	mode := flag.String("mode", "serve", "hash or serve")
	file := flag.String("file", "", "file to process")
	flag.Parse()

	if *mode == "hash" {
		hash(*file)
	} else if *mode == "serve" {
		serve(*file)
	} else {
		panic("invalid mode")
	}
}

func hash(file string) {
	dat, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x\n", sha1.Sum(dat))
}

func serve(file string) {
	dat, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("sha1(flag) = %s", dat))
	})
	e.Logger.Fatal(e.Start(":5089"))
}
