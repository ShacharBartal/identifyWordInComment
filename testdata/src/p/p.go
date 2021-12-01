package p

import (
	"fmt"
)

//func myLog(format string, args ...interface{}) {
//	const prefix = "[my] "
//	log.Printf(prefix + format, args...)
//}

// nolint - above // want "word" "is contained"
func checkSpecialWordInFunc() {
	// nolint - in // want "word `nolint` is contained"
}
func checkWordInString() {
	fmt.Println("dominos") // todo kuku // want "word `todo` is contained"
}
