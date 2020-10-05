package here_test

import (
	"github.com/dc0d/here"
)

func firstFn() error {
	return here.Mark(rootCause)
}

func anotherFn() error {
	return here.Mark(rootCause)
}

func secondFn() error {
	return here.Mark(firstFn())
}

func thirdFn() error {
	return here.Mark(secondFn())
}

func callAnonymousFunc() error {
	fn := func() error { return here.Mark(rootCause) }
	return here.Mark(fn())
}

func whereIsThisPlace() here.Loc {
	return here.Here()
}

func returnTheCaller() here.Loc {
	return here.Here(here.WithSkip(2))
}

func theCaller() here.Loc {
	return returnTheCaller()
}

func lessThanOneSkipIsIgnored() here.Loc {
	return here.Here(here.WithSkip(0))
}

func inShortWhereIsThisPlace() here.Loc {
	return here.Here(here.WithShortFile(), here.WithShortFunc())
}
