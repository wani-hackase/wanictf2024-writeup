package main

import (
	"embed"
	"html/template"
	"net/url"

	"github.com/go-ldap/ldap"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slices"
)

//go:embed templates
var fs embed.FS

//go:embed flag.txt
var flag string

var ldap_allow = [...]string{"localhost", "127.0.0.1", "::1", "ldap.example.com", "ldap.wani.example"}

func main() {
	t, err := template.ParseFS(fs, "templates/index.html")
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return t.Execute(c.Response().Writer, "sign in to get your flag!")
	})

	e.POST("/", func(c echo.Context) error {
		getMsg := func() string {
			userid := c.FormValue("userid")
			password := c.FormValue("password")
			ldapUrl := c.FormValue("ldap_url")

			u, err := url.Parse(ldapUrl)
			if err != nil {
				return "invalid ldap url"
			}
			if !slices.Contains(ldap_allow[:], u.Hostname()) {
				return "ldap url not allowed"
			}

			l, err := ldap.DialURL(ldapUrl)
			if err != nil {
				return "failed to connect to ldap server"
			}
			defer l.Close()

			err = l.Bind("userid="+userid+",ou=people,dc=wani,dc=example,dc=com", password)
			if err != nil {
				return "userid or password is incorrect"
			}

			return "congratulations! your flag is " + flag
		}

		return t.Execute(c.Response().Writer, getMsg())
	})

	e.Logger.Fatal(e.Start(":6867"))
}
