import "slices"

func isAnagram(s string, t string) bool {
	runesOfS := []rune(s)
	runesOfT := []rune(t)

	slices.Sort(runesOfS)
	slices.Sort(runesOfT)

	sortedS := string(runesOfS)
	sortedT := string(runesOfT)

	return sortedS == sortedT
}
