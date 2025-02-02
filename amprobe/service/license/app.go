// Package license
// Date: 2023/6/7 12:56
// Author: Amu
// Description:
package license

import (
	"amprobe/service/license/api"
	"amprobe/service/license/repository"
	"amprobe/service/license/service"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	api.Set,
	service.Set,
	repository.Set,
)
