package stringcompression

import "strconv"

func compress(chars []byte) int {
	write := 0
	read := 0

	for read < len(chars) {

		ch := chars[read]
		count := 0

		for read < len(chars) && chars[read] == ch {
			read++
			count++
		}

		chars[write] = ch
		write++

		if count > 1 {
			digits := strconv.Itoa(count)
			for i := 0; i < len(digits); i++ {
				chars[write] = digits[i]
				write++
			}

		}
	}

	return write
}
