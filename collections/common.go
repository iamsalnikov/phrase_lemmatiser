package collections

func IsStringSlicesEquals(a, b []string) bool {
	resultMap := make(map[string]int64)

	for _, av := range a {
		if _, ok := resultMap[av]; !ok {
			resultMap[av] = 1
		} else {
			resultMap[av] += 1
		}
	}

	for _, bv := range b {
		if _, ok := resultMap[bv]; !ok {
			return false
		} else {
			resultMap[bv] -= 1
		}
	}

	for _, r := range resultMap {
		if r != 0 {
			return false
		}
	}

	return true
}
