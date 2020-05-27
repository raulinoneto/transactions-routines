package apierror

import (
	"encoding/json"
	"fmt"
	"log"
)

type Severity int

const (
	Debug = iota
	Warning
	Critical
)

type Stack struct {
	Message string
}

type ApiError struct {
	Code           string
	Message        string
	Err            error
	HttpStatusCode int
	Severity
}

func NewApiError(code, message string, err error, httpStatusCode int, severity Severity) (errReturned *ApiError) {
	return &ApiError{code, message, err, httpStatusCode, severity}
}

func (a *ApiError) Error() string {
	return fmt.Sprintf("%s:%s", a.Code, a.Message)
}

func GetStack(e *ApiError) []Stack {
	res := []Stack{{e.Error()}}

	subErr, ok := e.Err.(*ApiError)
	if !ok {
		return res
	}
	return append(res, GetStack(subErr)...)
}

func NewDebug(code string, message string, err error) *ApiError {
	return NewApiError(code, message, err, 0, Debug)
}

func NewCritical(code string, message string, httpStatusCode int, err error) *ApiError {
	errReturned := NewApiError(code, message, err, httpStatusCode, Critical)
	errJson, err := json.Marshal(errReturned)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(errJson))
	return errReturned
}

func NewWarning(code string, message string, httpStatusCode int, err error) *ApiError {
	return NewApiError(code, message, err, httpStatusCode, Warning)
}
