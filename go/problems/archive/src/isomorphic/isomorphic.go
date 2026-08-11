package isisomorphic

func isIsomorphic(s string, t string) bool {
	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sc, tc := s[i], t[i]
		if val, ok := sToT[sc]; ok && val != tc {
			return false
		}
		sToT[sc] = tc
		if val, ok := tToS[tc]; ok && val != sc {
			return false
		}

		tToS[tc] = sc

	}
	return true
}
