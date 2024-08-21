// Package compose
// Date       : 2024/8/21 18:22
// Author     : Amu
// Description:
package compose

import "testing"

func TestGenerateDockerCompose(t *testing.T) {
	filePath := "/Users/amu/Desktop/docker-compose.yml"
	err := GenerateDockerCompose(filePath)
	if err != nil {
		t.Fatal(err)
	}
}
