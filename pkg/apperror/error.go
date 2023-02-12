package apperror

type Code string

const (
	CodeError   Code = "error"
	CodeInvalid Code = "invalid"
	CodeSuccess Code = "success"
)
