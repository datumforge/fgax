[![Build status](https://badge.buildkite.com/9884d4aae19682480e179efef95198c4db6fc1a26d9bcce22f.svg)](https://buildkite.com/datum/fgax)

# fgax

Go libraries to interact with [OpenFGA](https://openfga.dev/)

## Packages

### fgax

Wrappers to interact with the [OpenFGA go-sdk](https://github.com/openfga/go-sdk) and client libraries

#### Installation

You can install `fgax` by running the following command:

```shell
go get github.com/datumforge/fgax@latest
```

### entfga

[Ent extension](https://entgo.io/docs/extensions/) to create relationship tuples using [Ent Hooks](https://entgo.io/docs/hooks/)

#### Installation

You can install `entfga` by running the following command:

```shell
go get github.com/datumforge/fgax/entfga@latest
```

In addition to installing `entfga`, you need to create two files in your `ent` directory: `entc.go` and `generate.go`.
The `entc.go` file should contain the following code:

```go
//go:build ignore

package main

import (
	"log"
	"github.com/datumforge/fgax/entfga"
	"entgo.io/ent/entc"
)

func main() {
	if err := entc.Generate("./schema",
		&gen.Config{},
		entc.Extensions(
            entfga.NewFGAExtension(
                entfga.WithSoftDeletes(),
            ),
		),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
```

The `generate.go` file should contain the following code:

```go
package ent

//go:generate go run -mod=mod entc.go
```

#### Usage

In the `ent` schema, provide the following annotation:

```go 
// Annotations of the OrgMembership
func (OrgMembership) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entfga.Annotations{
			ObjectType: "organization",
		},
	}
}
```

The `ObjectType` **must** be the same between the ID field name in the schema and the object type in the FGA relationship. In the example above
the field in the schema is `OrganizationID` and the object in FGA is `organization`. 

#### Soft Deletes

If you are using the soft delete mixin provided by [entx](https://github.com/datumforge/datum/blob/authz-hooks/internal/entx/softdeletes.go), add 
the following option to the extension:

```go
    entfga.WithSoftDeletes(),
```

This will allow the hooks to delete tuples correctly after the `ent.Op` is updated to a `UpdateOne` from a `DeleteOne`