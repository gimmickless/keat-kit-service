package custom

import "fmt"

// ElemNotFoundError is a custom error type.
// Wraps the errors when get/delete/update operations are performed on a non-existing element.
type ElemNotFoundError struct {
	ID  string
	Err error
}

func (e *ElemNotFoundError) Error() string {
	return fmt.Sprintf("not found %s: %v", e.ID, e.Err)
}
