module github.com/finos/morphir-go/pkg/bindings/wit

go 1.26.0

require (
	github.com/Masterminds/semver/v3 v3.5.0
	github.com/finos/morphir-go/pkg/bindings/typemap v0.4.0-alpha.4
	github.com/finos/morphir-go/pkg/models v0.4.0-alpha.4
	github.com/finos/morphir-go/pkg/pipeline v0.4.0-alpha.4
	github.com/finos/morphir-go/pkg/toolchain v0.4.0-alpha.4
	github.com/finos/morphir-go/pkg/vfs v0.4.0-alpha.4
	github.com/stretchr/testify v1.12.1
	go.bytecodealliance.org v0.7.0
)

require (
	github.com/bmatcuk/doublestar/v4 v4.9.2 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/finos/morphir-go/pkg/config v0.3.3 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/rs/zerolog v1.34.0 // indirect
	github.com/tetratelabs/wazero v1.9.0 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/sys v0.33.0 // indirect
)

replace (
	github.com/finos/morphir-go/pkg/bindings/golang => ../golang
	github.com/finos/morphir-go/pkg/bindings/morphir-elm => ../morphir-elm
	github.com/finos/morphir-go/pkg/bindings/typemap => ../typemap
	github.com/finos/morphir-go/pkg/bindings/wit => ../wit
	github.com/finos/morphir-go/pkg/config => ../../config
	github.com/finos/morphir-go/pkg/docling-doc => ../../docling-doc
	github.com/finos/morphir-go/pkg/logging => ../../logging
	github.com/finos/morphir-go/pkg/models => ../../models
	github.com/finos/morphir-go/pkg/nbformat => ../../nbformat
	github.com/finos/morphir-go/pkg/pipeline => ../../pipeline
	github.com/finos/morphir-go/pkg/sdk => ../../sdk
	github.com/finos/morphir-go/pkg/task => ../../task
	github.com/finos/morphir-go/pkg/toolchain => ../../toolchain
	github.com/finos/morphir-go/pkg/tooling => ../../tooling
	github.com/finos/morphir-go/pkg/vfs => ../../vfs
)
