package coredns

import (
	"path/filepath"
	"strings"
)

func getCoreDNSPath(fqdn string, basePath string, suffixes []string) string {
	fqdnSlice := strings.Split(strings.TrimSuffix(fqdn, "."), ".")

	// reverse
	last := len(fqdnSlice) - 1
	for i := 0; i < len(fqdnSlice)/2; i++ {
		fqdnSlice[i], fqdnSlice[last-i] = fqdnSlice[last-i], fqdnSlice[i]
	}

	var path []string
	path = append(path, basePath)
	path = append(path, fqdnSlice...)
	path = append(path, suffixes...)
	return filepath.Join(path...)
}
