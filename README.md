# example-go-type-alias

This is an example repository for a 3rd party module which defines type aliases
such as
1. a generic type alias (`AliasType`) as a reference example.
2. an unexported type (`unexportedType`) that is type aliased to an exported
   type (`UnexportedAliasType`),
3. an exported internal type (`internal/InternalType`) that is type aliased in the same
   module in an external package to an exported type (`InternalAliasType`),

The examples are used in code generation context as use-cases.
