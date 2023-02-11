package response

import (
	"net/http"

	"github.com/mahiro72/go-user-api/pkg/apperror"
)

var CodeStatuses = map[apperror.Code]int{
	apperror.CodeError:   http.StatusInternalServerError,
	apperror.CodeSuccess: http.StatusOK,
	apperror.CodeInvalid: http.StatusBadRequest,
}
