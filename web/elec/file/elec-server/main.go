package main

import (
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {
	db, err := createDb()
	if err != nil {
		panic(err)
	}

	adminCh := makeAdminCh()
	startAdminWorker(adminCh)

	e := echo.New()

	e.IPExtractor = echo.ExtractIPDirect()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("1M"))
	e.Use(CspMiddleware)

	e.Renderer = NewTemplateRenderer()

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	e.POST("/", func(c echo.Context) error {
		id := uuid.New().String()

		err := db.Create(&post{
			ID:      id,
			Title:   c.FormValue("title"),
			Content: c.FormValue("content"),
		}).Error
		if err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, "/article/"+id)
	})

	e.GET("/article/:id", func(c echo.Context) error {
		id := c.Param("id")

		var p post
		err := db.Find(&p, "id = ?", id).Error
		if err != nil {
			return err
		}

		if p.ID == "" {
			return echo.ErrNotFound
		}

		return c.Render(http.StatusOK, "article", &p)
	})

	e.POST("/report/:id", func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return err
		}

		adminCh <- id

		return c.Render(http.StatusOK, "report", nil)
	}, middleware.RateLimiter(middleware.NewRateLimiterMemoryStoreWithConfig(middleware.RateLimiterMemoryStoreConfig{
		Rate:  rate.Every(20 * time.Second),
		Burst: 3,
	})))

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func CspMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' https://cdn.jsdelivr.net https://esm.sh 'unsafe-inline'; style-src 'self' https://cdn.jsdelivr.net; object-src 'none'; base-uri 'none'; frame-ancestors 'none'")
		return next(c)
	}
}
