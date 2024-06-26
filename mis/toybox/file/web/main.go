package main

import (
	"embed"
	"html/template"
	"io"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed templates
var fs embed.FS

func main() {
	t, err := template.ParseFS(fs, "templates/*.html")
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("10K")) // tiny executables only

	e.GET("/", func(c echo.Context) error {
		return t.ExecuteTemplate(c.Response().Writer, "index", nil)
	})
	e.POST("/upload", func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		id := uuid.New().String()
		dst, err := os.Create("./data/" + id)
		if err != nil {
			return err
		}
		defer dst.Close()

		if err := dst.Chmod(0755); err != nil {
			return err
		}

		if _, err := io.Copy(dst, src); err != nil {
			return err
		}

		return t.ExecuteTemplate(c.Response().Writer, "upload", struct {
			Id         string
			SocketHost string
			SocketPort string
		}{
			Id:         id,
			SocketHost: strings.Split(c.Request().Host, ":")[0],
			SocketPort: "9893",
		})
	})

	e.Logger.Fatal(e.Start(":1850"))
}
