module github.com/finos/morphir-go/pkg/vfs

go 1.25.5

require (
	github.com/bmatcuk/doublestar/v4 v4.9.2
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
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
