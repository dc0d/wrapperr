# here

Some Go modules/packages that help with instrumenting code with logging, can be configured in a way to enhance the log entry with the information about _where the logging happens_ inside the code.

This package provides an `error` type which tells _where the error happened_ in the code. It's compatible with the `errors` package in Go 1.13 and later and the root error can be unwrapped using `errors.Unwrap(...)` function.

## TODO

- ...
