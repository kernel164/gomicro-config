package ssm

import (
	"strings"
)

func format1(ct string) string {
	parts := strings.Split(ct, "/")
	if len(parts) <= 1 {
		return ct
	}
	return parts[1]
}
