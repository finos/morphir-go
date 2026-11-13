module github.com/finos/morphir-go/tests/bdd

go 1.26.0

require (
	github.com/bmatcuk/doublestar/v4 v4.10.4
	github.com/cucumber/godog v0.16.0
	github.com/finos/morphir-go/pkg/bindings/wit v0.0.0-00010101000000-000000000000
	github.com/finos/morphir-go/pkg/config v0.4.0-alpha.4
	github.com/finos/morphir-go/pkg/docling-doc v0.4.0-alpha.4
	github.com/finos/morphir-go/pkg/models v0.4.0-alpha.4
	github.com/finos/morphir-go/pkg/tooling v0.0.0-00010101000000-000000000000
	github.com/pelletier/go-toml/v2 v2.4.3
	github.com/stretchr/testify v1.12.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/BurntSushi/toml v1.6.0 // indirect
	github.com/Masterminds/semver/v3 v3.5.0 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/cucumber/gherkin/go/v42 v42.0.0 // indirect
	github.com/cucumber/messages/go/v34 v34.2.0 // indirect
	github.com/finos/morphir-go/pkg/bindings/typemap v0.4.0-alpha.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-memdb v1.3.5 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/tetratelabs/wazero v1.9.0 // indirect
	go.bytecodealliance.org v0.7.0 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
)

replace (
	github.com/finos/morphir-go/pkg/bindings/golang => ../../pkg/bindings/golang
	github.com/finos/morphir-go/pkg/bindings/morphir-elm => ../../pkg/bindings/morphir-elm
	github.com/finos/morphir-go/pkg/bindings/typemap => ../../pkg/bindings/typemap
	github.com/finos/morphir-go/pkg/bindings/wit => ../../pkg/bindings/wit
	github.com/finos/morphir-go/pkg/config => ../../pkg/config
	github.com/finos/morphir-go/pkg/docling-doc => ../../pkg/docling-doc
	github.com/finos/morphir-go/pkg/logging => ../../pkg/logging
	github.com/finos/morphir-go/pkg/models => ../../pkg/models
	github.com/finos/morphir-go/pkg/nbformat => ../../pkg/nbformat
	github.com/finos/morphir-go/pkg/pipeline => ../../pkg/pipeline
	github.com/finos/morphir-go/pkg/sdk => ../../pkg/sdk
	github.com/finos/morphir-go/pkg/task => ../../pkg/task
	github.com/finos/morphir-go/pkg/toolchain => ../../pkg/toolchain
	github.com/finos/morphir-go/pkg/tooling => ../../pkg/tooling
	github.com/finos/morphir-go/pkg/vfs => ../../pkg/vfs
)
