# example-go-type-alias

This is an example repository for a 3rd party module which defines type aliases
such as
1. a generic type alias (`alias.Alias`) as a reference example.
2. an unexported type (`alias.unexported`) that is type aliased to an exported
   type (`alias.UnexportedAlias`),
3. an exported internal type (`internal.Internal`) that is type aliased in the same
   module in an external package to an exported type (`alias.InternalAlias`),

The examples are used in code generation context as use-cases.
