module github.com/finos/morphir-go/pkg/task

go 1.26.0

require (
	github.com/finos/morphir-go/pkg/pipeline v0.4.0-alpha.4
	github.com/finos/morphir-go/pkg/tooling v0.3.3
	github.com/finos/morphir-go/pkg/vfs v0.4.0-alpha.4
	github.com/stretchr/testify v1.12.1
)

require (
	github.com/bmatcuk/doublestar/v4 v4.9.2 // indirect
	github.com/finos/morphir-go/pkg/models v0.4.0-alpha.4 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/rs/zerolog v1.35.1 // indirect
	github.com/santhosh-tekuri/jsonschema/v6 v6.0.2 // indirect
	go.yaml.in/yaml/v2 v2.4.2 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/sys v0.36.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	sigs.k8s.io/yaml v1.6.0 // indirect
)

replace (
	github.com/finos/morphir-go/pkg/bindings/golang => ../bindings/golang
	github.com/finos/morphir-go/pkg/bindings/morphir-elm => ../bindings/morphir-elm
	github.com/finos/morphir-go/pkg/bindings/typemap => ../bindings/typemap
	github.com/finos/morphir-go/pkg/bindings/wit => ../bindings/wit
	github.com/finos/morphir-go/pkg/config => ../config
	github.com/finos/morphir-go/pkg/docling-doc => ../docling-doc
	github.com/finos/morphir-go/pkg/logging => ../logging
	github.com/finos/morphir-go/pkg/models => ../models
	github.com/finos/morphir-go/pkg/nbformat => ../nbformat
	github.com/finos/morphir-go/pkg/pipeline => ../pipeline
	github.com/finos/morphir-go/pkg/sdk => ../sdk
	github.com/finos/morphir-go/pkg/task => ../task
	github.com/finos/morphir-go/pkg/toolchain => ../toolchain
	github.com/finos/morphir-go/pkg/tooling => ../tooling
	github.com/finos/morphir-go/pkg/vfs => ../vfs
)
