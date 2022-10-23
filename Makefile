# Note: to make a plugin compatible with a binary built in debug mode, add `-gcflags='all=-N -l'`

PLUGIN_OS ?= linux
PLUGIN_ARCH ?= amd64

plugin_sqlite: bin/$(PLUGIN_OS)$(PLUGIN_ARCH)/sqlite.so

bin/$(PLUGIN_OS)$(PLUGIN_ARCH)/sqlite.so: pkg/plugins/glauth-sqlite/sqlite.go
	GOOS=$(PLUGIN_OS) GOARCH=$(PLUGIN_ARCH) go build ${TRIM_FLAGS} -ldflags "${BUILD_VARS}" -buildmode=plugin -o $@ $^

plugin_sqlite_linux_amd64:
	PLUGIN_OS=linux PLUGIN_ARCH=amd64 make plugin_sqlite

plugin_sqlite_linux_arm64:
	PLUGIN_OS=linux PLUGIN_ARCH=arm64 make plugin_sqlite

plugin_sqlite_darwin_amd64:
	PLUGIN_OS=darwin PLUGIN_ARCH=amd64 make plugin_sqlite

plugin_sqlite_darwin_arm64:
	PLUGIN_OS=darwin PLUGIN_ARCH=arm64 make plugin_sqlite

release-glauth-sqlite:
	@P=sqlite M=pkg/plugins/glauth-sqlite make releaseplugin
