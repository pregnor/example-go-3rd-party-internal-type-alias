package alias

import (
	"github.com/pregnor/example-go-type-alias/internal"
)

// InternalAliasType is used as an example exported external type aliased to an
// exported internal type.
type InternalAliasType = internal.InternalType
