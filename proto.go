package nebula

// these imports to to work around Go modules not being able to resolve the imports in the protos
import (
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf"
	_ "github.com/gogo/protobuf/gogoproto"
)
