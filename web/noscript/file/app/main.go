package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type InMemoryDB struct {
	data map[string][2]string
	mu   sync.RWMutex
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: make(map[string][2]string),
	}
}

func (db *InMemoryDB) Set(key, value1, value2 string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = [2]string{value1, value2}
}

func (db *InMemoryDB) Get(key string) ([2]string, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	vals, exists := db.data[key]
	return vals, exists
}

func (db *InMemoryDB) Delete(key string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.data, key)
}

func main() {
	ctx := context.Background()

	db := NewInMemoryDB()

	redisAddr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Home page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Noscript!",
		})
	})

	// Sign in
	r.POST("/signin", func(c *gin.Context) {
		id := uuid.New().String()
		db.Set(id, "test user", "test profile")
		c.Redirect(http.StatusMovedPermanently, "/user/"+id)
	})

	// Get user profiles
	r.GET("/user/:id", func(c *gin.Context) {
		c.Header("Content-Security-Policy", "default-src 'self', script-src 'none'")
		id := c.Param("id")
		re := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
		if re.MatchString(id) {
			if val, ok := db.Get(id); ok {
				params := map[string]interface{}{
					"id":       id,
					"username": val[0],
					"profile":  template.HTML(val[1]),
				}
				c.HTML(http.StatusOK, "user.html", params)
			} else {
				_, _ = c.Writer.WriteString("<p>user not found <a href='/'>Home</a></p>")
			}
		} else {
			_, _ = c.Writer.WriteString("<p>invalid id <a href='/'>Home</a></p>")
		}
	})

	// Modify user profiles
	r.POST("/user/:id/", func(c *gin.Context) {
		id := c.Param("id")
		re := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
		if re.MatchString(id) {
			if _, ok := db.Get(id); ok {
				username := c.PostForm("username")
				profile := c.PostForm("profile")
				db.Delete(id)
				db.Set(id, username, profile)
				if _, ok := db.Get(id); ok {
					c.Redirect(http.StatusMovedPermanently, "/user/"+id)
				} else {
					_, _ = c.Writer.WriteString("<p>user not found <a href='/'>Home</a></p>")
				}
			} else {
				_, _ = c.Writer.WriteString("<p>user not found <a href='/'>Home</a></p>")
			}
		} else {
			_, _ = c.Writer.WriteString("<p>invalid id <a href='/'>Home</a></p>")
		}
	})

	// Get username API
	r.GET("/username/:id", func(c *gin.Context) {
		id := c.Param("id")
		re := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
		if re.MatchString(id) {
			if val, ok := db.Get(id); ok {
				_, _ = c.Writer.WriteString(val[0])
			} else {
				_, _ = c.Writer.WriteString("<p>user not found <a href='/'>Home</a></p>")
			}
		} else {
			_, _ = c.Writer.WriteString("<p>invalid id <a href='/'>Home</a></p>")
		}
	})

	// Report API
	r.POST("/report", func(c *gin.Context) {
		url := c.PostForm("url") // URL to report, example : "/user/ce93310c-b549-4fe2-9afa-a298dc4cb78d"
		re := regexp.MustCompile("^/user/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
		if re.MatchString(url) {
			if err := redisClient.RPush(ctx, "url", url).Err(); err != nil {
				_, _ = c.Writer.WriteString("<p>Failed to report <a href='/'>Home</a></p>")
				return
			}
			if err := redisClient.Incr(ctx, "queued_count").Err(); err != nil {
				_, _ = c.Writer.WriteString("<p>Failed to report <a href='/'>Home</a></p>")
				return
			}
			_, _ = c.Writer.WriteString("<p>Reported! <a href='/'>Home</a></p>")
		} else {
			_, _ = c.Writer.WriteString("<p>invalid url <a href='/'>Home</a></p>")
		}
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}
