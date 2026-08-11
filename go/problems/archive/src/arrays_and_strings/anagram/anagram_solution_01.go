package anagram

// See https://leetcode.com/problems/valid-anagram/

func isAnagramSolution02(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	// convert to rune due to usage of unicode
	ra := []rune(s)
	rt := []rune(t)

	for _, ch := range ra {
		for j, cj := range rt {
			if ch == cj {
				// save state
				rt[j] = rt[len(rt)-1]
				rt = rt[:len(rt)-1]
				break
			}
		}
	}

	return len(rt) <= 0
}
