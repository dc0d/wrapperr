package wrapperr_test

import (
	"github.com/dc0d/wrapperr"
)

func fix1() error {
	return steps{
		func() error { return nil },
		func() error { return nil },
		func() error { return wrapperr.WithStack(errRootCause, "annotated root cause") },
		func() error { return nil },
	}.run()
}

func fix2() error {
	return wrapperr.WithStack(errRootCause, "annotated root cause")
}

func fix3() error {
	if err := fix2(); err != nil {
		return wrapperr.WithStackf(err, "sample %v", "annotation")
	}

	return nil
}

func fix4() error {
	if err := fix3(); err != nil {
		return wrapperr.WithStack(err, "second sample", "annotation")
	}

	return nil
}

type steps []step

func (steps steps) run() error {
	for _, fn := range steps {
		if err := fn(); err != nil {
			return wrapperr.WithStackf(err, "sample %v", "annotation")
		}
	}

	return nil
}

type step func() error
