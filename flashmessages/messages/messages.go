package messages

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

const sessionName = "fmessages"

func getCookieStore() *sessions.CookieStore {
	// In real-world application, use env variables to store the session key.
	sessionKey := "test-session-key"
	return sessions.NewCookieStore([]byte(sessionKey))
}

// Set sets new message into the cookie storage.
func Set(c echo.Context, name, value string) {
	session, _ := getCookieStore().Get(c.Request(), sessionName)
	session.AddFlash(value, name)

	session.Save(c.Request(), c.Response())
}

// Get gets flash messages from the cookie storage.
func Get(c echo.Context, name string) []string {
	session, _ := getCookieStore().Get(c.Request(), sessionName)
	fm := session.Flashes(name)
	// If we have some messages.
	if len(fm) > 0 {
		session.Save(c.Request(), c.Response())
		// Initiate a strings slice to store messages.
		var flashes []string
		for _, fl := range fm {
			// Add message to the slice.
			flashes = append(flashes, fl.(string))
		}

		return flashes
	}
	return nil
}
