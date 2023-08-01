package main

import "strings"

func main() {

}

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

func reverseLeftWords2(s string, n int) string {
	sb := strings.Builder{}
	for i := n; i < len(s); i++ {
		sb.WriteByte(s[i])
	}
	for i := 0; i < n; i++ {
		sb.WriteByte(s[i])
	}
	return sb.String()
}
