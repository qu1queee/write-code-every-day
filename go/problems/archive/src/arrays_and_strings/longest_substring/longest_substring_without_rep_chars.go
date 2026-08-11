package lsworc

// From Leetcode https://leetcode.com/problems/longest-substring-without-repeating-characters/

/*
Some personal notes:
- Tagged as Medium
*/

func lengthOfLongestSubstring(s string) int {
	var longest, start int

	for index, ch := range s {
		for j, ch2 := range s[start:index] {
			if ch == ch2 {
				if index-start > longest {
					longest = index - start
				}
				start += j + 1
				break
			}
		}
	}

	if len(s)-start > longest {
		longest = len(s) - start
	}
	return longest
}
