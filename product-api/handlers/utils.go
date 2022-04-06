package handlers

import (
	"regexp"
)

func isValid(id []byte) []byte {
	re := regexp.MustCompile("([0-9]+)")
	return re.Find(id)
}
