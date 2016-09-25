// Package house solves an Exercism challenge.
package house

import (
	"strings"
)

var phrases = []string{
	"the horse and the hound and the horn\nthat belonged to",
	"the farmer sowing his corn\nthat kept",
	"the rooster that crowed in the morn\nthat woke",
	"the priest all shaven and shorn\nthat married",
	"the man all tattered and torn\nthat kissed",
	"the maiden all forlorn\nthat milked",
	"the cow with the crumpled horn\nthat tossed",
	"the dog\nthat worried",
	"the cat\nthat killed",
	"the rat\nthat ate",
	"the malt\nthat lay in",
}

// Embed just puts two phrases together with a space.
func Embed(l, p string) string {
	return l + " " + p
}

// Verse makes a single verse of the song.
func Verse(subject string, relPhrases []string, nounPhrase string) string {
	if len(relPhrases) == 0 {
		return Embed(subject, nounPhrase)
	}
	return Verse(Embed(subject, relPhrases[0]), relPhrases[1:], nounPhrase)
}

// Song returns the entire song.
func Song() string {
	var verses = make([]string, 0)
	for i := len(phrases); i >= 0; i-- {
		verses = append(verses, Verse("This is", phrases[i:], "the house that Jack built."))
	}
	return strings.Join(verses, "\n\n")
}
