package users

import "fmt"

type UserRepositoryError interface {
	error
	UserRepositoryError()
}

type InvalidPropertyError[T any] struct {
	Property string
	Value    T
	Reason   string
}

// Reports Property, Value and Message
func (ipe *InvalidPropertyError[T]) Error() string {
	return fmt.Sprintf(
		"Invalid property %s with value %v, %s",
		ipe.Property,
		ipe.Value,
		ipe.Reason,
	)
}

func (ipe *InvalidPropertyError[T]) UserRepositoryError() {
	// This function is used to check that this is a User Repo Error
}

var _ UserRepositoryError = &InvalidPropertyError[string]{} // Ensure that the interface is implemented

type NotFoundError[T any] struct {
	Resource string
	Property string
	Value    T
}

// Reports of not finding an entity with property = value
func (nfe *NotFoundError[T]) Error() string {
	return fmt.Sprintf(
		"Could not find entity with %s = %v",
		nfe.Property,
		nfe.Value,
	)
}

func (nfe *NotFoundError[T]) UserRepositoryError() {
	// This function is used to check that this is a User Repo Error
}

var _ UserRepositoryError = &NotFoundError[string]{} // Ensure that the interface is implemented

type AlreadyExistsError[T any] struct {
	Property string
	Value    T
}

// Reports of an entity with property = value already existing
func (aee *AlreadyExistsError[T]) Error() string {
	return fmt.Sprintf(
		"Entity with %s = %v already exists",
		aee.Property,
		aee.Value,
	)
}

func (aee *AlreadyExistsError[T]) UserRepositoryError() {
	// This function is used to check that this is a User Repo Error
}

var _ UserRepositoryError = &AlreadyExistsError[string]{} // Ensure that the interface is implemented

type InternalError struct {
	Message string
}

// Reports of an internal error
func (ie *InternalError) Error() string {
	return fmt.Sprintf(
		"Internal error: %s",
		ie.Message,
	)
}

func (ie *InternalError) UserRepositoryError() {
	// This function is used to check that this is a User Repo Error
}

var _ UserRepositoryError = &InternalError{} // Ensure that the interface is implemented
