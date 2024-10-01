package user

import (
	"github.com/go-chi/render"
	"net/http"
	"vetback/internal/model"
)

// @Summary Sign out
// @Tags auth
// @Description Sign out
// @Accept json
// @Produce json
// @Router /sign-out [get]
func (a *userApi) NewSignOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.setToken(w, "")

		render.JSON(w, r, model.NewResponse("вы вышли из аккаунта"))
	}
}
