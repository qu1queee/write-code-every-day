package strstr

// From Leetcode https://leetcode.com/problems/implement-strstr/

func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-n; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}
