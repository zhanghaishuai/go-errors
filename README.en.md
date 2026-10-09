[简体中文](README.md) | [English](README.en.md)

# errors

A Go error handling package inspired by [`github.com/pkg/errors`](https://github.com/pkg/errors), with additional APIs aligned with the standard library's error handling model.

This package provides error wrapping, stack trace support, error inspection, and error composition, while avoiding unnecessary stack capture when creating sentinel errors.

## Requirements

* **Go 1.26 or later**

Go 1.26 is required for the standard library's `errors.AsType` API used by this package.

## Installation

```bash
go get github.com/zhanghaishuai/go-errors
```

## Features

* **Contextual errors:** Add context to errors with `Wrap` and `Wrapf`.
* **Stack traces:** Capture and format stack traces where they are useful for debugging.
* **Sentinel errors without stack traces:** `New` creates a new error without capturing a stack trace. This avoids capturing an initialization-time stack for sentinel errors.
* **Error inspection:** Inspect wrapped errors with `Is` and `As`.
* **Error unwrapping:** Retrieve wrapped errors with `Unwrap` and, where supported, retrieve the original cause with `Cause`.
* **Error composition:** Combine multiple errors with `Join`.
* **Generic type-based inspection:** Use `AsType` to retrieve an error of a specified type.

## Creating errors

Use `New` to create a simple error, including sentinel errors:

```go
var ErrNotFound = errors.New("not found")
```

Unlike implementations that capture a stack trace when creating an error, this package's `New` does not capture a stack trace. This is useful for sentinel errors initialized at package load time, where a captured stack would point to initialization rather than the location where the error is later returned or handled.

## Adding context to an error

Use `Wrap` or `Wrapf` to add context while preserving the underlying error:

```go
data, err := io.ReadAll(r)
if err != nil {
    return errors.Wrap(err, "read failed")
}
```

When formatted with `%+v`, wrapped errors can include stack trace information, depending on how the error was created and wrapped.

## Inspecting errors

Use `Is` to check whether an error matches a target, including errors in its wrapping chain:

```go
if errors.Is(err, ErrNotFound) {
    // Handle the sentinel error.
}
```

Use `As` to find an error assignable to a specified target type:

```go
var pathErr *fs.PathError
if errors.As(err, &pathErr) {
    // Handle the path error.
}
```

For generic type-based inspection, use `AsType`:

```go
pathErr, ok := errors.AsType[*fs.PathError](err)
if ok {
    // Handle the path error.
}
```

## Unwrapping errors

Use `Unwrap` to retrieve the next error in a wrapping chain:

```go
cause := errors.Unwrap(err)
```

If your code uses the original `pkg/errors` cause convention, `Cause` can be used to retrieve the underlying cause from errors that implement `Cause() error`.

```go
type causer interface {
    Cause() error
}
```

`Cause` and `Unwrap` serve related but distinct purposes. `Unwrap` follows the standard `Unwrap() error` convention, while `Cause` follows the older `Cause() error` convention.

## Combining errors

Use `Join` to combine multiple errors into a single error value:

```go
err := errors.Join(errRead, errClose)
if errors.Is(err, errRead) {
    // Handle the read error.
}
```

The resulting error can be inspected with `Is` and `As`, allowing callers to check individual errors in the combined result.

## Compatibility

This package builds on the error handling APIs popularized by `github.com/pkg/errors` and incorporates APIs available in newer Go standard library versions, including `Join` and `AsType`.

## License

BSD-2-Clause. See [LICENSE](LICENSE) for details.
