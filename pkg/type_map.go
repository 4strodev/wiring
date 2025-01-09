package pkg

import "reflect"

// typeMap is a concurrently save map that holds dependency specs associated to [reflect.Type]
type typeMap map[reflect.Type]*dependencySpec
