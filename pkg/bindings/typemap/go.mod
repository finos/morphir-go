module github.com/finos/morphir-go/pkg/bindings/typemap

go 1.25.5

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
