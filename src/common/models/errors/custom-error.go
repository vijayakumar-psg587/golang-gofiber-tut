package errors

type CustomError struct {
	Err error
	Message string
	Status int32
	Code HTTPStatusCode
	Timestamp int64
}

type HTTPStatusCode string


func (ce *CustomError) Error() string {
	return ce.Message
}
