package cmd

import "github.com/spf13/cobra"

func makeShellCmd(dbc *Command) *cobra.Command {
	// TODO: leave it to adapter
	// TODO: we can also use docker exec to use container shell ....
	// mysql -u user --password -h database_host database_name
	// https://dev.mysql.com/doc/refman/5.7/en/multiple-server-clients.html need to use 127.0.0.1 to avoid using sock
	// mysql -u root -pmysqlpassword -h 127.0.0.1
	// pg TODO: how to pass password in command option ...
	// psql -h 127.0.0.1 -U pguser -W icehub
	return nil
}
