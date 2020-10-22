package container

import "strings"

func getContainerName(out string) string {
	return out[:strings.Index(out, ` `)]
}
