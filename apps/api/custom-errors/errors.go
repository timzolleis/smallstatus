package custom_errors

import "fmt"

type ResourceNotFoundError struct {
	Message string
}

type DatabaseError struct {
	Message string
}

func (e *ResourceNotFoundError) Error() string {
	return fmt.Sprintf("Resource not found: %s", e.Message)
}

func NewResourceNotFoundError(message string) ResourceNotFoundError {
	return ResourceNotFoundError{Message: message}
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("Database error: %s", e.Message)
}

func NewDatabaseError(message string) *DatabaseError {
	return &DatabaseError{Message: message}
}
