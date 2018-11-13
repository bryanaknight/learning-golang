// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	if YellingNonQuestion(remark) {
		return "Whoa, chill out!"
	} else if YellingQuestion(remark){
		return "Calm down, I know what I'm doing!"
	} else if Question(remark) {
		return "Sure."
	} else {
		return "Whatever."
	}
}

func Question(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func AllCaps(remark string) bool {
	return strings.ToUpper(remark) == remark
}

func YellingNonQuestion(remark string) bool {
	Yelling := AllCaps(remark)
	Questioning := Question(remark)

	return Yelling && !Questioning
}

func YellingQuestion(remark string) bool {
	Yelling := AllCaps(remark)
	Questioning := Question(remark)

	return Yelling && Questioning
}
