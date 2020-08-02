package alias

import (
	"github.com/pregnor/example-go-type-alias/internal"
)

// InternalAlias is used as an example exported external type aliased to an
// exported internal type.
type InternalAlias = internal.Internal
