//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc/gen"
	"github.com/datumforge/fgax"
	"github.com/datumforge/fgax/entfga"
	"go.uber.org/zap"

	"entgo.io/ent/entc"
)

func main() {
	gqlExt, err := entgql.NewExtension(
		// Tell Ent to generate a GraphQL schema for
		// the Ent schema in a file named ent.graphql.
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("../schema/ent.graphql"),
		entgql.WithConfigPath("../gqlgen.yml"),
		entgql.WithWhereInputs(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	if err := entc.Generate("./schema",
		&gen.Config{
			Features: []gen.Feature{gen.FeaturePrivacy},
		},
		entc.Dependency(
			entc.DependencyName("Authz"),
			entc.DependencyType(fgax.Client{}),
		),
		entc.Dependency(
			entc.DependencyName("Logger"),
			entc.DependencyType(zap.SugaredLogger{}),
		),
		entc.Extensions(
			gqlExt,
			entfga.NewFGAExtension(),
		),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
