// Package nginx
// Date: 2024/8/6 17:15
// Author: Amu
// Description:
package nginx

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func UUID() string {
	return uuid.Must(uuid.NewRandom()).String()
}

func validateName(name string) bool {
	return !strings.ContainsAny(name, " _")
}

func GenerateUpstreamName(site *Site) (string, error) {
	name := site.Name
	if !validateName(name) {
		return "", fmt.Errorf("invalid site name [%s]", name)
	}
	return strings.Join([]string{UpstreamPrefix, UUID(), name}, GeneralSep), nil
}
