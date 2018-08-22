package indefinite

import (
	"fmt"
	"regexp"
)

// All the logic below is sourced from https://github.com/tandrewnichols/indefinite
// That library is released under the MIT license

const (
	acronymString          = "^[A-Z]+$"
	irregularAcronymString = "^[UFHLMNRSX]"
	leadingVowelString     = "^[aeiouAEIOU]"
)

var acronymPattern = regexp.MustCompile(acronymString)
var irregularAcronymPattern = regexp.MustCompile(irregularAcronymString)
var leadingVowelPattern = regexp.MustCompile(leadingVowelString)

// GetArticle returns the article for a word
func GetArticle(word string) string {
	startsWithVowel := leadingVowelPattern.MatchString(word)
	var irreg bool

	if acronymPattern.MatchString(word) {
		irreg = irregularAcronymPattern.MatchString(word)
	} else {
		_, irreg = irregularMap[word]
	}

	if startsWithVowel == irreg {
		return "a"
	}
	return "an"
}

// AddArticle prepends the article to the word and returns both
func AddArticle(word string) string {
	return fmt.Sprintf("%s %s", GetArticle(word), word)
}
