package contracts

import (
	"errors"
	"fmt"
	"net/http"
)

// Code is a stable, documented API error code. Clients branch on it; messages
// are for humans and may be localized.
type Code string

const (
	CodeValidation          Code = "VALIDATION_ERROR"
	CodeUnauthenticated     Code = "UNAUTHENTICATED"
	CodeForbidden           Code = "FORBIDDEN"
	CodeNotFound            Code = "NOT_FOUND"
	CodeConflict            Code = "CONFLICT"
	CodeRateLimited         Code = "RATE_LIMITED"
	CodeProfileIncomplete   Code = "PROFILE_INCOMPLETE"
	CodeInvalidServiceArea  Code = "INVALID_SERVICE_AREA"
	CodeEmailNotConfirmed   Code = "EMAIL_NOT_CONFIRMED"
	CodeTokenExpired        Code = "TOKEN_EXPIRED"
	CodeTokenInvalid        Code = "TOKEN_INVALID"
	CodeInvalidAvailability Code = "INVALID_AVAILABILITY"
	CodeInvalidImage        Code = "INVALID_IMAGE"
	CodeInvalidPassword     Code = "INVALID_PASSWORD"
	CodeCSRFInvalid         Code = "CSRF_INVALID"
	CodeMFARequired         Code = "MFA_REQUIRED"
	CodeInvalidTransition   Code = "INVALID_STATE_TRANSITION"
	CodePIINotAllowed       Code = "PII_NOT_ALLOWED"
	CodeInternal            Code = "INTERNAL_ERROR"
)

var statusByCode = map[Code]int{
	CodeValidation:          http.StatusBadRequest,
	CodeUnauthenticated:     http.StatusUnauthorized,
	CodeForbidden:           http.StatusForbidden,
	CodeNotFound:            http.StatusNotFound,
	CodeConflict:            http.StatusConflict,
	CodeRateLimited:         http.StatusTooManyRequests,
	CodeProfileIncomplete:   http.StatusUnprocessableEntity,
	CodeInvalidServiceArea:  http.StatusUnprocessableEntity,
	CodeEmailNotConfirmed:   http.StatusForbidden,
	CodeTokenExpired:        http.StatusUnprocessableEntity,
	CodeTokenInvalid:        http.StatusUnprocessableEntity,
	CodeInvalidAvailability: http.StatusUnprocessableEntity,
	CodeInvalidImage:        http.StatusUnprocessableEntity,
	CodeInvalidPassword:     http.StatusUnprocessableEntity,
	CodeCSRFInvalid:         http.StatusForbidden,
	CodeMFARequired:         http.StatusForbidden,
	CodeInvalidTransition:   http.StatusConflict,
	CodePIINotAllowed:       http.StatusUnprocessableEntity,
	CodeInternal:            http.StatusInternalServerError,
}

var defaultMessage = map[Code]string{
	CodeValidation:          "Confira os campos destacados.",
	CodeUnauthenticated:     "Entre para continuar.",
	CodeForbidden:           "Você não tem acesso a esta ação.",
	CodeNotFound:            "Não encontramos este recurso.",
	CodeConflict:            "A ação conflita com o estado atual.",
	CodeRateLimited:         "Muitas tentativas. Tente novamente em alguns minutos.",
	CodeProfileIncomplete:   "Complete nome, categoria e área para publicar.",
	CodeInvalidServiceArea:  "Informe uma área de atendimento válida.",
	CodeEmailNotConfirmed:   "Confirme seu e-mail para continuar.",
	CodeTokenExpired:        "Este link expirou. Solicite outro.",
	CodeTokenInvalid:        "Este link não é válido. Solicite outro.",
	CodeInvalidAvailability: "Confira os dias e horários.",
	CodeInvalidImage:        "Escolha uma imagem JPG, PNG ou WebP de até 2 MB.",
	CodeInvalidPassword:     "A senha atual não confere.",
	CodeCSRFInvalid:         "Atualize a página e tente novamente.",
	CodeMFARequired:         "Confirme o código de verificação.",
	CodeInvalidTransition:   "Esta ação não é permitida no estado atual.",
	CodePIINotAllowed:       "Remova dados pessoais deste campo público.",
	CodeInternal:            "Algo deu errado. Tente novamente.",
}

// Error is the shared application error. Services return it; HTTP handlers add
// the correlation id and never expose internal details.
type Error struct {
	Code       Code
	Message    string
	Details    any
	HTTPStatus int
	cause      error
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	if e.cause != nil {
		return fmt.Sprintf("%s: %v", e.Code, e.cause)
	}
	return string(e.Code)
}

func (e *Error) Unwrap() error { return e.cause }

func New(code Code, message string) *Error {
	if message == "" {
		message = defaultMessage[code]
	}
	status := statusByCode[code]
	if status == 0 {
		status = http.StatusInternalServerError
	}
	return &Error{Code: code, Message: message, HTTPStatus: status}
}

func Newf(code Code, format string, args ...any) *Error {
	return New(code, fmt.Sprintf(format, args...))
}

func Wrap(code Code, cause error) *Error {
	e := New(code, "")
	e.cause = cause
	return e
}

func Validation(message string) *Error { return New(CodeValidation, message) }

// As extracts a *Error from an error chain, falling back to INTERNAL_ERROR.
func As(err error) *Error {
	var e *Error
	if errors.As(err, &e) {
		return e
	}
	internal := New(CodeInternal, "")
	internal.cause = err
	return internal
}
