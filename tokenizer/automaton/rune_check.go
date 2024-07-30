package automaton

func isBracket(s rune) bool {
	return s == '{' || s == '}' || s == '[' || s == ']' || s == '(' || s == ')'
}
