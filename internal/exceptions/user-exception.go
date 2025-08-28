package exceptions

import (
	"fmt"
)

type UserNotFoundErr struct {
	UserInfo any
}

func (e UserNotFoundErr) Error() string {
	return fmt.Sprintf("User not found for user information: %v", e.UserInfo)
}

type UserAlreadyExistsErr struct {
	UserInfo any
}

func (e UserAlreadyExistsErr) Error() string {
	return fmt.Sprintf("User already exists for user information %v", e.UserInfo)
}
