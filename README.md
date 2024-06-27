[![Build status](https://badge.buildkite.com/9884d4aae19682480e179efef95198c4db6fc1a26d9bcce22f.svg?branch=main)](https://buildkite.com/datum/fgax) 
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=datumforge_fgax&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=datumforge_fgax)
[![Go Report Card](https://goreportcard.com/badge/github.com/datumforge/fgax)](https://goreportcard.com/report/github.com/datumforge/fgax)
[![Go Reference](https://pkg.go.dev/badge/github.com/datumforge/fgax.svg)](https://pkg.go.dev/github.com/datumforge/fgax)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache2.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)

# fgax

A go library for interacting with [OpenFGA](https://openfga.dev/) - it is comprised of 2 packages, `fgax` and `entfga`.
- fgax: wrapper to interact with the [OpenFGA go-sdk](https://github.com/openfga/go-sdk) and client libraries
- entfga: an [ent extension](https://entgo.io/docs/extensions/) to create relationship tuples using [ent Hooks](https://entgo.io/docs/hooks/)

## fgax package

You can install `fgax` by running the following command:

```shell
go get github.com/datumforge/fgax@latest
```

## entfga

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

### Usage

When creating the `*ent.Client` add the following to enable the authz hooks and policies:

```
	client.WithAuthz()
```

The `privacy` feature **must** be turned on:

```
	Features: []gen.Feature{gen.FeaturePrivacy},
```

## Generate Hooks and Policies

In the `ent` schema, provide the following annotation:

```go
// Annotations of the OrgMembership
func (OrgMembership) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entfga.Annotations{
			ObjectType:   "organization",
			IncludeHooks: true,
			IDField:      "OrganizationID", // Defaults to ID, override to object ID field
		},
	}
}
```

The `ObjectType` **must** be the same between the ID field name in the schema and the object type in the FGA relationship. In the example above
the field in the schema is `OrganizationID` and the object in FGA is `organization`.

If the `ID` field is `Optional()`, you'll need to set `NillableIDField: true,` on the annotation to ensure the `string` value is used instead of the `pointer` on the `CreateInput`.


## Generate Policies Only

In the `ent` schema, provide the following annotation:

```go
// Annotations of the Organization
func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entfga.Annotations{
			ObjectType:   "organization",
			IncludeHooks: false,
		},
	}
}
```

## Using Policies

A policy check function will be created per mutation and query type when the annotation is used, these can be set on the policy of the schema.
They must be wrapped in the `privacy` `MutationRuleFunc`, as seen the example below:

```go
// Policy of the Organization
func (Organization) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoSubject(),
			privacy.OrganizationMutationRuleFunc(func(ctx context.Context, m *generated.OrganizationMutation) error {
				return m.CheckAccessForEdit(ctx)
			}),
			// Add a separate delete policy if permissions for delete of the object differ from normal edit permissions
			privacy.OrganizationMutationRuleFunc(func(ctx context.Context, m *generated.OrganizationMutation) error {
				return m.CheckAccessForDelete(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.OrganizationQueryRuleFunc(func(ctx context.Context, q *generated.OrganizationQuery) error {
				return q.CheckAccess(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
```


### Soft Deletes

If you are using the soft delete mixin provided by [entx](https://github.com/datumforge/datum/blob/authz-hooks/internal/entx/softdeletes.go), add
the following option to the extension:

```go
    entfga.WithSoftDeletes(),
```

This will allow the hooks to delete tuples correctly after the `ent.Op` is updated to a `UpdateOne` from a `DeleteOne`

## Contributing

Please read the [contributing](.github/CONTRIBUTING.md) guide as well as the [Developer Certificate of Origin](https://developercertificate.org/). You will be required to sign all commits to the Datum project, so if you're unfamiliar with how to set that up, see [github's documentation](https://docs.github.com/en/authentication/managing-commit-signature-verification/about-commit-signature-verification).

## Security

We take the security of our software products and services seriously, including all of the open source code repositories managed through our Github Organizations, such as [datumforge](https://github.com/datumforge). If you believe you have found a security vulnerability in any of our repositories, please report it to us through coordinated disclosure.

**Please do NOT report security vulnerabilities through public github issues, discussions, or pull requests!**

Instead, please send an email to `security@datum.net` with as much information as possible to best help us understand and resolve the issues. See the security policy attached to this repository for more details.

## Questions?

Open a github issue on this repository and we'll respond as soon as we're able!
