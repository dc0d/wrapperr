package here_test

import (
	"github.com/dc0d/here"
)

func fn1() here.Calls {
	return here.Mark()
}

func fn2() here.Calls {
	return here.Mark(here.WithShortFiles())
}

func fn3() here.Calls {
	return here.Mark(here.WithShortFuncs())
}

func fn4() here.Calls {
	return here.Mark(here.WithSkip(3), here.WithSkip(-1))
}

func fn5() here.Calls {
	return fn4()
}
