// Package errors package will have custom errors for the expenses package
package errors

import "fmt"

var (
	ErrNotFound = fmt.Errorf("not found") // ErrNotFound is used for not found data
)
