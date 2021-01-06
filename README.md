[![PkgGoDev](https://pkg.go.dev/badge/dc0d/wrapperr)](https://pkg.go.dev/github.com/dc0d/wrapperr) [![Go Report Card](https://goreportcard.com/badge/github.com/dc0d/wrapperr)](https://goreportcard.com/report/github.com/dc0d/wrapperr) [![Maintainability](https://api.codeclimate.com/v1/badges/c0fdd128cafcb6ce0c52/maintainability)](https://codeclimate.com/github/dc0d/wrapperr/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/c0fdd128cafcb6ce0c52/test_coverage)](https://codeclimate.com/github/dc0d/wrapperr/test_coverage)

# wrapperr

_Where_ did that error happen down the call chain?

> Right at the bottom!

<div align="center">
<img src="./images/github_com_dc0d_wrapperr.png" width="80%" alt="https://github.com/dc0d/wrapperr">
</div>

<br />

All you need to do is, instead of:

```go
return nil, err
```

Do:

```go
return nil, wrapperr.WithStack(err)
```

Also, it is possible to annotate the stack in the middle:

<div align="center">
<img src="./images/github_com_dc0d_wrapperr_annotate.png" width="80%" alt="https://github.com/dc0d/wrapperr">
</div>

<br />

And to get the original error, just used the standard `errors.Unwrap(error)` function from built-in `errors` package.
