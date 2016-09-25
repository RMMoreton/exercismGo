// Pacakge foodchain is a solution to an Exercism challenge.
package foodchain

import (
	"fmt"
	"strings"
)

const testVersion = 2

type animal struct {
	name, line, lineEnd string
}

var animals = []animal{
	animal{"fly", "", ""},
	animal{"spider", "It wriggled and jiggled and tickled inside her.\n",
		" that wriggled and jiggled and tickled inside her"},
	animal{"bird", "How absurd to swallow a bird!\n", ""},
	animal{"cat", "Imagine that, to swallow a cat!\n", ""},
	animal{"dog", "What a hog, to swallow a dog!\n", ""},
	animal{"goat", "Just opened her throat and swallowed a goat!\n", ""},
	animal{"cow", "I don't know how she swallowed a cow!\n", ""},
	animal{"horse", "She's dead, of course!", ""},
}

var firstLine = "I know an old lady who swallowed a %s.\n"
var whySwallowed = "She swallowed the %s to catch the %s%s.\n"
var lastLine = "I don't know why she swallowed the fly. Perhaps she'll die."

// Verse returns a single verse of the song.
func Verse(v int) string {
	if v < 1 || v > len(animals) { // v must be a valid verse number
		return ""
	}
	var verse []string
	a := animals[v-1]
	verse = append(verse, fmt.Sprintf(firstLine, a.name))
	verse = append(verse, a.line)
	if v == len(animals) { // Don't do the whole 'she swalled the...' if she swallowed the horse
		return strings.Join(verse, "")
	}
	for i := v - 2; i >= 0; i-- {
		a = animals[i]
		verse = append(verse, fmt.Sprintf(whySwallowed, animals[i+1].name, a.name, a.lineEnd))
	}
	verse = append(verse, lastLine)
	return strings.Join(verse, "")
}

// Verses passes two ints, and returns a string of all the verses from start to end.
func Verses(start, end int) string {
	var verses []string
	for ; start <= end; start++ {
		verses = append(verses, Verse(start))
	}
	return strings.Join(verses, "\n\n")
}

// Song returns the entire song.
func Song() string {
	return Verses(1, 8)
}
