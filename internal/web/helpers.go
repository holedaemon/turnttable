package web

import "strings"

func checkHeaderOrder(row []string) bool {
	csvHeaderRows := []string{"title", "artist", "label", "cn", "genre", "released", "purchased", "medium"}

	for i, r := range row {
		if !strings.EqualFold(r, csvHeaderRows[i]) {
			return false
		}
	}

	return true
}
