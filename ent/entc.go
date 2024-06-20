//go:build ignore

package main

import (
	"log"
	"strings"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/vektah/gqlparser/v2/ast"
)

func snake(s string) string {
	var result string

	for i, v := range s {
		if i > 0 && v >= 'A' && v <= 'Z' {
			result += "_"
		}
		result += string(v)
	}

	return strings.ToLower(result)
}

func generateSnakeCasedNames(g *gen.Graph, s *ast.Schema) error {
	for _, s := range s.Types {
		for _, f := range s.Fields {
			f.Name = snake(f.Name)
		}
	}
	return nil
}

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaHook(generateSnakeCasedNames),
		entgql.WithConfigPath("../gqlgen.yml"),
		entgql.WithSchemaPath("ent.graphql"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	if err := entc.Generate("./schema", &gen.Config{Features: []gen.Feature{gen.FeatureUpsert}}, entc.Extensions(ex)); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
