package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"

)

var (
	ValidationError = "VALIDATION_ERROR"
	RedisSetupError = "REDIS_SETUP_ERROR"
	NoRecordError   = "NO_RECORD_FOUND_ERROR"
	NoResourceError = "NO_RESOURCE_FOUND_ERROR"
	CreateError     = "CREATE_ERROR"
	UpdateError     = "UPDATE_ERROR"
	LogError     = "LOG_ERROR"
	MongoDBError = "MONGO_DB_ERROR"
	InvalidResource = "INVALID_RESOURCE_ERROR"
)

var CustomError = map[string]int{
	ValidationError: 400,
	RedisSetupError: 500,
	NoRecordError: 204,
	InvalidResource: 422,
    CreateError: 500,
	UpdateError: 500,
    LogError: 500,
	MongoDBError: 500,
	NoResourceError: 404,
}


func (err ErrorResponse) Error() string {
	var errorBody ErrorBody
	return fmt.Sprintf("%v", errorBody)

}
func ErrorArrayToError(errorBody []validator.FieldError) error {
	var errorResponse ErrorResponse
	errorResponse.TimeStamp = time.Now().Format(time.RFC3339)
	errorResponse.ErrorReference = uuid.New()
    errorResponse.ErrorType = ValidationError
	errorResponse.Code = CustomError[ValidationError]
	for _, value := range errorBody {
		body := ErrorBody{Message: value.Error()}
		errorResponse.Errors = append(errorResponse.Errors, body)
	}
	return errorResponse
}
func ErrorMessage(errorType string, message string) error {
	var errorResponse ErrorResponse
	errorResponse.TimeStamp = time.Now().Format(time.RFC3339)
	errorResponse.ErrorReference = uuid.New()
	errorResponse.ErrorType = errorType
	errorResponse.Code = CustomError[errorType]
	errorResponse.Errors = append(errorResponse.Errors, ErrorBody{Message: message})
	return errorResponse
}

type ErrorBody struct {
	//Code    string      `json:"code"`
	Message interface{} `json:"message"`
	//Source  string      `json:"source"`
}
type ErrorResponse struct {
	ErrorReference uuid.UUID   `json:"error_reference"`
	ErrorType      string 		`json:"error_type"`
	TimeStamp      string      `json:"timestamp"`
	Code      	   int 			`json:"code"`
	Errors         []ErrorBody `json:"errors"`
}