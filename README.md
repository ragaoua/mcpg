An collection of MCP servers providing tools for interacting with a PostgreSQL database.

Every server provides tools for a given "realm" :

- General : get information about the postgresql version, data directory/logs/wal/binary location...
- Data : tools in this realm provide an interface to access user data,
    but also metadata (cluster/database/schema information, roles...)
- System : interact with the underlying filesystems, get information about the memory and CPU usage...
- Backups : tools that interact with the backups
- Documentation : tools that access the PostgreSQL official documentation

Each realm can be seen as a set of tools that can be provided to an agent that would be specialized in that realm.
This would allow for creating a powerful agentic AI workflow.
