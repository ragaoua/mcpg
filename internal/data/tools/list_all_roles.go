package data

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mark3labs/mcp-go/mcp"
)

func ListAllRolesHandler(ctx context.Context, request mcp.CallToolRequest, db_url string) (*mcp.CallToolResult, error) {
	roles, err := listAllRoles(db_url)
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
