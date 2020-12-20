package dao

type CommonError struct {
	Message string `json:"message"`
}

var (
	PayloadParsingError = CommonError{Message: "Encountered Error while parsing request payload"}
)