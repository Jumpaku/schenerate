package spanner

import (
	"context"
	"flag"
	"fmt"
	"github.com/Jumpaku/sqanner/tokenize"
	"github.com/samber/lo"
	"strings"
	"testing"

	spanner_admin "cloud.google.com/go/spanner/admin/database/apiv1"
	spanner_adminpb "cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
)

var spannerInstance = flag.String("instance", "", "-instance=<spanner instance>")
var spannerProject = flag.String("project", "", "-project=<GCP project>")

func Setup(t *testing.T, database string, ddls []string) (q queryer, teardown func()) {
	t.Helper()

	if *spannerProject == "" || *spannerInstance == "" {
		t.Skip(`spanner project and instance are required`)
		return queryer{}, nil
	}

	project := *spannerProject
	instance := *spannerInstance
	dataSource := fmt.Sprintf(`projects/%s/instances/%s/databases/%s`, project, instance, database)

	ctx := context.Background()
	{
		c, err := spanner_admin.NewDatabaseAdminClient(ctx)
		if err != nil {
			t.Fatalf(`fail to create spanner admin client: %v`, err)
		}
		defer c.Close()
		{
			parent := fmt.Sprintf(`projects/%s/instances/%s`, project, instance)
			op, err := c.CreateDatabase(ctx, &spanner_adminpb.CreateDatabaseRequest{
				Parent:          parent,
				CreateStatement: fmt.Sprintf("CREATE DATABASE %s", database),
			})
			if err != nil {
				t.Fatalf(`fail to create spanner database in %q: %v`, parent, err)
			}
			if _, err := op.Wait(ctx); err != nil {
				t.Fatalf(`fail to wait create spanner database: %v`, err)
			}
		}
		{
			ddls = lo.FlatMap(ddls, func(s string, _ int) []string {
				tokens, err := tokenize.Tokenize([]rune(s))
				if err != nil {
					t.Fatalf(`fail to tokenize: %v`, err)
				}
				ddls := []string{""}
				for _, token := range tokens {
					switch {
					case token.Kind == tokenize.TokenSpecialChar && string(token.Content) == ";":
						ddls = append(ddls, "")
					case token.Kind == tokenize.TokenComment || token.Kind == tokenize.TokenSpace:
						ddls[len(ddls)-1] += " "
					default:
						ddls[len(ddls)-1] += string(token.Content)
					}
				}
				return ddls
			})
			ddls = lo.Map(ddls, func(s string, _ int) string { return strings.TrimSpace(s) })
			ddls = lo.Filter(ddls, func(s string, _ int) bool { return s != "" })
			if len(ddls) > 0 {
				op, err := c.UpdateDatabaseDdl(ctx, &spanner_adminpb.UpdateDatabaseDdlRequest{
					Database:   dataSource,
					Statements: ddls,
				})
				if err != nil {
					t.Fatalf(`fail to update schema: %v`, err)
				}
				if err := op.Wait(ctx); err != nil {
					t.Fatalf(`fail to wait update schema: %v`, err)
				}
			}
		}
	}

	q, err := Open(ctx, project, instance, database)
	if err != nil {
		t.Fatalf(`fail to create spanner client with %q %q %q: %v`, project, instance, database, err)
	}
	teardown = func() {
		q.Close()
		q.Close()
	}
	return q, teardown
}
