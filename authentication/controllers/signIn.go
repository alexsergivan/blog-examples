package controllers

import (
	"github.com/alexsergivan/blog-examples/authentication/auth"
	"github.com/alexsergivan/blog-examples/authentication/user"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"path"
)

// SignInForm responsible for signIn Form rendering.
func SignInForm() echo.HandlerFunc {
	return func(c echo.Context) error {
		fp := path.Join("templates", "signIn.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if err := tmpl.Execute(c.Response().Writer, nil); err != nil {
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
			// If the two passwords don't match, return a 401 status
			return echo.NewHTTPError(http.StatusUnauthorized, "Password is incorrect")
		}
		// If password is correct, generate tokens and set cookies.
		err := auth.GenerateTokensAndSetCookies(storedUser, c)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
		}

		return c.Redirect(http.StatusMovedPermanently, "/admin")
	}
}