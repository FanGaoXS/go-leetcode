package main

import "strings"

func main() {

}

func replaceSpace(s string) string {
	sb := strings.Builder{}
	for _, c := range s {
		if c == ' ' {
			sb.WriteString("%20")
			continue
		}
		sb.WriteString(string(c))
	}
	return sb.String()
}
