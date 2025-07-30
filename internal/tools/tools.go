package tools

func GetTools(dbUrl string) []*Tool {
	tools := []*Tool{
		&ListAllRoles,
	}
	for i := range tools {
		tools[i].DbUrl = dbUrl
	}
	return tools
}
