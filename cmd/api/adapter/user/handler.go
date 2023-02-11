package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/mahiro72/go-user-api/pkg/presenter"
	"github.com/mahiro72/go-user-api/pkg/registry"
	"github.com/mahiro72/go-user-api/pkg/usecase/user"
)

type Handler struct {
	usecase *user.Usecase
}

func NewHandler(repo *registry.Repository) *Handler {
	usecase := user.NewUsecase(
		repo.NewUser(),
	)
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	idParamString := chi.URLParam(r, "id")
	if idParamString == "" {
		presenter.BadRequestErrResponse(w, fmt.Errorf("error: is is empty"))
		return
	}
	idParam, err := strconv.Atoi(idParamString)
	if err != nil {
		presenter.BadRequestErrResponse(w, err)
		return
	}
	inp := &user.GetInput{
		Id: idParam,
	}
	out, err := h.usecase.Get(r.Context(), inp)
	if err != nil {
		presenter.BadRequestErrResponse(w, err)
		return
	}
	u := NewUserView(out.User)
	presenter.Response(w, u)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

}
