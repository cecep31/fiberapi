package exception

type ValidationError struct {
	Message string
}

func (validationerror ValidationError) Error() string {
	return validationerror.Message
}
