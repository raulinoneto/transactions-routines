package httpadapters

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/raulinoneto/transactions-routines/internal/apierror"
)

type Error struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Error   string           `json:"error,omitempty"`
	Errors  []string         `json:"errors,omitempty"`
	Stack   []apierror.Stack `json:"stack"`
}

type Payload struct {
	*Error `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

func writeResponse(status int, body []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin-Type", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers")
	w.WriteHeader(status)
	if _, err := w.Write(body); err != nil {
		log.Fatalln(err.Error())
	}
}

func BuildCreatedResponse(result interface{}, apierr error, w http.ResponseWriter) {
	BuildResponse(result, apierr, http.StatusCreated, w)
}

func BuildOkResponse(result interface{}, apierr error, w http.ResponseWriter) {
	BuildResponse(result, apierr, http.StatusOK, w)
}

func BuildUnauthorizedResponse(w http.ResponseWriter) {
	payload := new(Payload)
	payload.Error = &Error{
		Code:    "not_authorized",
		Message: "Not Authorized",
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		asError(
			errors.New("Something are wrong in your body, please verify "),
			http.StatusBadRequest,
			w,
		)
		return
	}
	writeResponse(
		http.StatusUnauthorized,
		payloadJson,
		w,
	)
}
func BuildBadRequestResponse(err error, w http.ResponseWriter) {
	payload := new(Payload)
	payload.Error = &Error{
		Code:    "bad_request",
		Message: err.Error(),
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		asError(
			errors.New("Something are wrong in your body, please verify "),
			http.StatusBadRequest,
			w,
		)
		return
	}
	writeResponse(
		http.StatusBadRequest,
		payloadJson,
		w,
	)
}

func BuildBadRequestResponseWithErrorArray(apierrs []*apierror.ApiError, code, message string, w http.ResponseWriter) {
	payload := new(Payload)
	payload.Error = &Error{
		Code:    code,
		Message: message,
		Errors:  getErrMessages(apierrs),
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		asError(
			errors.New("Something are wrong in your body, please verify "),
			http.StatusBadRequest,
			w,
		)
		return
	}
	writeResponse(
		http.StatusBadRequest,
		payloadJson,
		w,
	)
}

func BuildResponse(result interface{}, apierr error, status int, w http.ResponseWriter) {
	if apierr != nil {
		asError(apierr, http.StatusInternalServerError, w)
		return
	}
	response := &Response{
		StatusCode: status,
		Data:       result,
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		asError(apierr, status, w)
		return
	}

	writeResponse(status, responseJson, w)
}

func getErrMessages(errs []*apierror.ApiError) []string {
	errMessages := make([]string, 0)
	for _, err := range errs {
		if err != nil {
			errMessages = append(errMessages, err.Error())
		}
	}
	return errMessages
}

func asError(err error, status int, w http.ResponseWriter) {
	payload := new(Payload)
	switch err.(type) {
	case *apierror.ApiError:
		apierr := err.(*apierror.ApiError)
		payload.Error = &Error{
			Code:    apierr.Code,
			Message: apierr.Message,
			Error:   apierr.Error(),
			Stack:   apierror.GetStack(apierr),
		}
		status = apierr.HttpStatusCode
	default:
		payload.Error = &Error{
			Code:    "generic.internal_error",
			Message: "an internal error occurred",
			Error:   err.Error(),
			Stack:   nil,
		}
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err.Error())
	}
	writeResponse(status, payloadJson, w)
}
