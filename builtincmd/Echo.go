package builtincmd

import "strings"

func Echo (parts []string) string {
	if len(parts) <= 1 {
		return ""
	}
	return strings.Join(parts[1:], " ")
}
