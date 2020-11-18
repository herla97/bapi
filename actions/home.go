package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
// @Summary Show home page buffalo
// @Description Get welcomen json buffalo
// @Tags home
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Welcome to Buffalo!"
// @Router / [get]
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Welcome to Buffalo!"}))
}
