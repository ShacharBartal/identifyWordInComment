//args: -EidentifyWordInComment

package p

//nolint - before function declaration // want "word `nolint` is contained"

func checkSpecialWordInCommentTst() {
	//nolint - in function // want "word `nolint` is contained"
}

func checkSpecialWordInCommentTst2() int {
	return 1 // nolint:check for nolint with a linter report // want "word `nolint` is contained"
}

func checkSpecialWordInCommentTst3() bool {
	return true // nolint:check // with some explain // want "word `nolint` is contained"
}

func checkSpecialWordInCommentTst4() string {
	return "dominos" // dominos should not be in our code // want "word `dominos` is contained"
}

// nolint free in the air // want "word `nolint` is contained"

func checkWordInCommentTst5() { // here please todo things // want "word `todo` is contained"
}
