package utility

import (
	"strings"
)

//CleanMagnet extracts URI from a magnet
func CleanMagnet(text string) (string, string) {

	split := strings.Split(text, "&")

	split2 := strings.Split(split[0], ":")

	return split[0], split2[len(split2)-1]
}
