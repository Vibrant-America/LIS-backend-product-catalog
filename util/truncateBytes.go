package util

// TruncateString truncate string to specific length
func TruncateString(s string, n int) string {
	if n > len(s) {
		return s
	}
	return s[:n]
}

func TruncateBytes(s []byte, n int) []byte {
	if n > len(s) {
		return s
	}
	return s[:n]
}
