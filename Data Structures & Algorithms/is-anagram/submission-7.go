func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashMapOfS := make(map[byte]int)
	hashMapOfT := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, ok := hashMapOfS[s[i]]; !ok {
			hashMapOfS[s[i]] = 1
		} else {
			hashMapOfS[s[i]] += 1
		}
        if _, ok := hashMapOfT[t[i]]; !ok {
			hashMapOfT[t[i]] = 1
		} else {
			hashMapOfT[t[i]] += 1
		}
	}

	for key, value := range hashMapOfS {
		value2, ok := hashMapOfT[key]
        if !ok {
            return false
        }
        if value != value2 {
            return false
        }
	}
	return true
}