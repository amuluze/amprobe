// Package host
// Date: 2024/3/6 12:43
// Author: Amu
// Description:
package host

import (
	"github.com/google/wire"
	
	"github.com/amuluze/amprobe/service/host/api"
	"github.com/amuluze/amprobe/service/host/repository"
	"github.com/amuluze/amprobe/service/host/service"
)

var Set = wire.NewSet(
	api.Set,
	service.Set,
	repository.Set,
)
