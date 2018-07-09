package sublist

type Relation string

func Sublist(a, b []int) Relation {
	if isSublist(a, b) {
		if isSublist(b, a) {
			return "equal"
		}
		return "sublist"
	}
	if isSublist(b, a) {
		return "superlist"
	}
	return "unequal"
}

func isSublist(a, b []int) bool {
	if len(a) == 0 {
		return true
	}
	for i := 0; i < len(b)-len(a)+1; i++ {
		if equalArrays(a, b[i:i+len(a)]) {
			return true
		}
	}
	return false
}

func equalArrays(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
