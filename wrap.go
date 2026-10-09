package errors

import stderrors "errors"

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
//
// Unwrap only calls a method of the form:
//
//	Unwrap() error
//
// In particular, Unwrap does not unwrap errors returned by Unwrap() []error.
func Unwrap(err error) error {
	return stderrors.Unwrap(err)
}

// Is reports whether any error in err's tree matches target.
//
// The tree consists of err itself, followed by the errors obtained by
// repeatedly calling Unwrap. When err wraps multiple errors, Is examines
// err followed by a depth-first traversal of its children.
//
// An error is considered to match target if it is equal to target, or if
// its Is(error) bool method returns true.
func Is(err, target error) bool {
	return stderrors.Is(err, target)
}

// As finds the first error in err's tree that matches target, and if one
// is found, sets target to that error value and returns true.
//
// The tree consists of err itself, followed by the errors obtained by
// repeatedly calling Unwrap. When err wraps multiple errors, As examines
// err followed by a depth-first traversal of its children.
//
// An error matches target if the error's concrete value is assignable to
// the value pointed to by target, or if the error has an As(any) bool
// method that returns true.
//
// As panics if target is nil, or if target is not a non-nil pointer to
// either a type that implements error or any interface type.
func As(err error, target any) bool {
	return stderrors.As(err, target)
}

// AsType finds the first error in err's tree that can be assigned to E,
// and returns it along with a boolean indicating whether the match was
// successful.
//
// If no matching error is found, AsType returns the zero value of E
// and false.
func AsType[E error](err error) (E, bool) {
	return stderrors.AsType[E](err)
}
