module github.com/finos/morphir-go/tests/bdd

go 1.25.5

require (
	github.com/bmatcuk/doublestar/v4 v4.9.2
	github.com/cucumber/godog v0.15.0
	github.com/finos/morphir-go/pkg/docling-doc v0.4.0-alpha.4
	github.com/finos/morphir-go/pkg/models v0.4.0-alpha.4
	github.com/pelletier/go-toml/v2 v2.2.4
	github.com/stretchr/testify v1.11.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/cucumber/gherkin/go/v26 v26.2.0 // indirect
	github.com/cucumber/messages/go/v21 v21.0.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gofrs/uuid v4.3.1+incompatible // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-memdb v1.3.4 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
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
