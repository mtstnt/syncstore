package helpers

import "fmt"

var DatabaseErr = NewDatabaseError(nil)
var IntegrityErr = NewIntegrityError(nil, "")

var _ error = (*DatabaseError)(nil)

type DatabaseError struct {
	err error
}

func NewDatabaseError(err error) DatabaseError {
	return DatabaseError{err}
}

func (d DatabaseError) Is(err error) bool {
	_, ok := err.(DatabaseError)
	return ok
}

func (d DatabaseError) Error() string {
	return fmt.Sprintf("database error: %s", d.err.Error())
}

// TODO: Iki ga terlalu ngerti buat apa? Krn semua errornya lho DB error ? Bee mesti service/usecase layer.
type IntegrityError struct {
	err     error
	message string
}

func NewIntegrityError(err error, message string) IntegrityError {
	return IntegrityError{err, message}
}

func (e IntegrityError) Error() string {
	return fmt.Sprintf("Error: %s", e.message)
}

func (e IntegrityError) Is(err error) bool {
	_, ok := err.(IntegrityError)
	return ok
}
