package controllers

import (
	"github.com/alexsergivan/blog-examples/flashmessages/auth"
	"github.com/alexsergivan/blog-examples/flashmessages/messages"
	"github.com/alexsergivan/blog-examples/flashmessages/user"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"path"
)

// SignInForm responsible for signIn Form rendering.
func SignInForm() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmpl, err := template.ParseFiles(path.Join("templates", "signIn.html"), path.Join("templates", "messages.html"))
		data := make(map[string]interface{})
		data["errors"] = messages.Get(c, "error")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if err := tmpl.Execute(c.Response().Writer, data); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return nil
	}
}

// SignIn will be executed after SignInForm submission.
func SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Load our "test" user.
		storedUser := user.LoadTestUser()
		// Initiate a new User struct.
		u := new(user.User)
		// Parse the submitted data and fill the User struct with the data from the SignIn form.
		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		// Compare the stored hashed password, with the hashed version of the password that was received
		if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(u.Password)); err != nil {
			// If the two passwords don't match, set a message and reload the page.
            messages.Set(c, "error", "Password is incorrect!")
			return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("userSignInForm"))
		}
		// If password is correct, generate tokens and set cookies.
		err := auth.GenerateTokensAndSetCookies(storedUser, c)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
		}
		messages.Set(c, "message", "Password is correct, you have been authenticated!")
		return c.Redirect(http.StatusMovedPermanently, "/admin")
	}
}