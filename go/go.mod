module github.com/mojo-lang/http/go

go 1.16

replace (
	github.com/mojo-lang/core/go => ../../core/go
	github.com/mojo-lang/lang/go => ../../lang/go
)

require (
	github.com/golang/protobuf v1.5.2
	github.com/json-iterator/go v1.1.12
	github.com/mojo-lang/core/go v0.0.0-20220316072623-f4cf14310a3a
	github.com/mojo-lang/lang/go v0.0.0-20220316075324-fd2be24ebd68
)
