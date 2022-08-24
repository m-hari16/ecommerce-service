package helper

import (
	"net/http"
	"strconv"
)

func Unprocessable() Response {
	return Response{Code: strconv.Itoa(http.StatusUnprocessableEntity), Success: false, Message: "Unprocessable entity"}
}

func ErrValidate(message interface{}) Response {
	return Response{Code: strconv.Itoa(http.StatusUnprocessableEntity), Success: false, Message: message}

}

func HasStored() Response {
	return Response{Code: strconv.Itoa(http.StatusOK), Success: true, Message: "Data has been stored"}
}

func HasOk(data interface{}) Response {
	return Response{Code: strconv.Itoa(http.StatusOK), Success: true, Message: "Ok", Data: data}
}
