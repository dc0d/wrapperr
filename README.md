[![PkgGoDev](https://pkg.go.dev/badge/dc0d/wrapperr)](https://pkg.go.dev/dc0d/wrapperr)

# wrapperr

> a pick at _where the logging happens_ vs _where the error happened_ in Go

Many Go modules for logging provide the option to log where the methods of the logger are called. For example the `logger.Error(someError)` could be called inside a file named _service.go_, from a function named `gitrepo.com/user/module/service-pkg/Connect(...)`.

That's useful information.

But **where** the actual error - `someError` - is coming from?

This library provides a utility for enhancing the error with information about the call stack. Which is very helpful especially while working on legacy code-bases.

Also the root cause error can be accessed using the standard `errors.Unwrap(error)` function at any step.