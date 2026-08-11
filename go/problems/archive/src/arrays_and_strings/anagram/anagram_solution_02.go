package anagram

// See https://leetcode.com/problems/valid-anagram/

func isAnagramSolution2(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune](int))
	tMap := make(map[rune](int))

	for _, sChar := range s {
		sMap[sChar] = sMap[sChar] + 1
	}

	for _, sChar := range t {
		tMap[sChar] = tMap[sChar] + 1
	}

	for key, val := range sMap {
		if val != tMap[key] {
			return false
		}
	}
	return true
}
