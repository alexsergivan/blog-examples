package controllers

import (
	"github.com/alexsergivan/blog-examples/flashmessages/messages"
	"github.com/labstack/echo/v4"
	"html/template"
	"net/http"
	"path"
)

func Admin() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmpl, err := template.ParseFiles(path.Join("templates", "admin.html"), path.Join("templates", "messages.html"))
		data  := make(map[string]interface{})
		data["messages"] = messages.Get(c, "message")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if err := tmpl.Execute(c.Response().Writer, data); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return nil
	}
}