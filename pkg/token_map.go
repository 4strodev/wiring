package pkg

// tokenMap is a concurrently save map that holds a dependency spec associated to a token
type tokenMap map[string]*dependencySpec
