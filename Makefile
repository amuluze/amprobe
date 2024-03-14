APP 			= amprobe
SERVER  		= ./cmd/${APP}
RELEASE_SERVER 	= release/${APP}/${APP}

GO 				= go
CONFIG_FILE := $(shell pwd)/configs/config.toml

.PHONY: dev
# run dev
dev:
	@$(GO) run $(SERVER)/main.go run -c $(CONFIG_FILE)

.PHONY: bin
# build bin
bin:
	$(GO) build -o $(APP) $(SERVER)

.PHONY: wire
# generate wire
wire:
	wire ./service/
