// Package repository
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package repository

import "github.com/google/wire"

var Set = wire.NewSet(
	AuditRepoSet,
)
