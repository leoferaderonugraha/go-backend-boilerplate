package errors

type UserError struct {
    message string
}

func NewUserError(message string) *UserError {
    return &UserError{message: message}
}

func (e *UserError) Error() string {
    return e.message
}

var (
    USER_ALREADY_EXISTS = NewUserError("User already exists")
    USER_NOT_FOUND = NewUserError("User not found")
)
