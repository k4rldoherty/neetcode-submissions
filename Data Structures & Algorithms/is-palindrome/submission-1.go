func isPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	sb := strings.Builder{}
	for i := range s {
		if s[i] < '0' || (s[i] > '9' && s[i] < 'A') || (s[i] > 'Z' && s[i] < 'a') || s[i] > 'z' {
			continue
		} 
		sb.WriteByte(s[i])
	}

	e := sb.Len() - 1
	s = sb.String()
	for i := range s {
		if s[i] != s[e] { return false }
		e--
	}
	return true
}
