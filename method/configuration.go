package method

import (
	"github.com/philiphil/apiman/method/MethodType"
)

type ApiMethodConfiguration struct {
	Method              method_type.ApiMethod
	SerializationGroups []string
}
