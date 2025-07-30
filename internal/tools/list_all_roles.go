package tools

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/mark3labs/mcp-go/mcp"
)

var ListAllRoles Tool

func init() {
	ListAllRoles = Tool{
		Spec: mcp.NewTool(
			"list_all_roles",
			mcp.WithDescription("List all roles in the cluster"),
		),
		Handler: listAllRolesHandler,
	}
}

func listAllRolesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Printf("db url : %v", ListAllRoles.DbUrl)
	if ListAllRoles.DbUrl == "" {
		return nil, errors.New("Field DbUrl unset")
	}

	roles, err := listAllRoles(ListAllRoles.DbUrl)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("%v", roles)), nil
}

func listAllRoles(db_url string) ([]string, error) {
	conn, err := pgx.Connect(context.Background(), db_url)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT rolname FROM pg_roles")
	if err != nil {
		return nil, fmt.Errorf("Query error : %v\n", err)
	}

	var roles []string
	for rows.Next() {
		var role string

		err := rows.Scan(&role)
		if err != nil {
			return nil, fmt.Errorf("Fetching error : %v\n", err)
		}

		roles = append(roles, role)
	}

	return roles, nil
}
