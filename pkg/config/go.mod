module github.com/finos/morphir-go/pkg/config

go 1.25.5

require (
	github.com/finos/morphir-go/pkg/bindings/typemap v0.4.0-alpha.4
	github.com/joho/godotenv v1.5.1
	github.com/pelletier/go-toml/v2 v2.2.4
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
