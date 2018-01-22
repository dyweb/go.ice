package db

import (
	"fmt"
	"os"
	"time"
	"database/sql"

	"github.com/at15/go.ice/ice/db/migration"
	"github.com/spf13/cobra"
)

// cobra command for database related operations

// Command is a wrapper to allow user update manager after cobra commands have been created
type Command struct {
	Root   *cobra.Command
	PreRun func(dbc *Command, cmd *cobra.Command, args []string)
	Mgr    *Manager
	db     string // database selected by user
}

func (dbc *Command) MustWrapper() *Wrapper {
	var (
		w    *Wrapper
		name string
		err  error
	)
	if dbc.db != "" {
		name = dbc.db
	} else if name, err = dbc.Mgr.DefaultName(); err != nil {
		log.Fatal(err)
	}
	if w, err = dbc.Mgr.Wrapper(name); err != nil {
		log.Fatal(err)
	}
	return w
}

var rootCmd = &cobra.Command{
	Use:   "db",
	Short: "database maintenance",
	Long:  "Database drivers, migration, status, REPL",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

var driverCmd = &cobra.Command{
	Use:   "drivers",
	Short: "registered database drivers",
	Long:  "Show registered database drivers",
	Run: func(cmd *cobra.Command, args []string) {
		drivers := Drivers()
		if len(drivers) == 0 {
			fmt.Println("not database/sql driver registered")
			return
		}
		fmt.Printf("%d drivers registered\n", len(drivers))
		for _, d := range drivers {
			fmt.Println(d)
		}
	},
}

var adapterCmd = &cobra.Command{
	Use:   "adapters",
	Short: "registered database adapters",
	Long:  "Show registered ice/db/adapters",
	Run: func(cmd *cobra.Command, args []string) {
		adapters := Adapters()
		if len(adapters) == 0 {
			fmt.Println("not ice/db/adapters adapter registered")
			return
		}
		fmt.Printf("%d adapters registered\n", len(adapters))
		for _, a := range adapters {
			fmt.Println(a)
		}
	},
}

func makeConfigCmd(dbc *Command) *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "print configuration",
		Long:  "Print configuration of manager and databases",
		Run: func(cmd *cobra.Command, args []string) {
			dbc.PreRun(dbc, cmd, args)
			dbc.Mgr.PrintConfig()
		},
	}
}

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

// TODO: create database (the user need to be root ... but this is normally the case in local dev ...)

func makePingCmd(dbc *Command) *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "check database connectivity",
		Long:  "Check if database is reachable",
		Run: func(cmd *cobra.Command, args []string) {
			dbc.PreRun(dbc, cmd, args)
			w := dbc.MustWrapper()
			if duration, err := w.Ping(5 * time.Second); err != nil {
				log.Fatal(err)
			} else {
				log.Infof("ping took %s", duration)
			}
			// TODO: dbc should have cleanup
		},
	}
}

func makeMigrationCmd(dbc *Command) *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database",
		Long:  "Run registered migration tasks to update schema and feed fixture",
		Run: func(cmd *cobra.Command, args []string) {
			dbc.PreRun(dbc, cmd, args)
			var (
				tx  *sql.Tx
				err error
			)
			w := dbc.MustWrapper()
			if tx, err = w.Transaction(); err != nil {
				log.Fatal(err)
			}
			init := migration.InitTask()
			if err = init.Up(tx); err != nil {
				log.Fatal(err)
			} else {
				if err = tx.Commit(); err != nil {
					log.Fatalf("failed to commit %v", err)
				}
			}
			log.Info("migration finished")
			// TODO: dbc should have cleanup
			// Aborted connection 6 to db: 'icehub' user: 'root' host: '172.19.0.1' (Got an error reading communication packets)
		},
	}
}

// TODO: command for migrating database (create table, fill in dummy data)
// TODO: dbshell https://docs.djangoproject.com/en/2.0/ref/django-admin/#dbshell
// - also consider support docker container ...
func NewCommand(preRun func(dbc *Command, cmd *cobra.Command, args []string)) *Command {
	dbc := &Command{Mgr: nil, PreRun: preRun}
	root := *rootCmd
	// flags
	root.PersistentFlags().StringVar(&dbc.db, "db", "", "database to run command on, ping/migrate etc.")
	// sub commands
	root.AddCommand(driverCmd)
	root.AddCommand(adapterCmd)
	root.AddCommand(makeConfigCmd(dbc))
	root.AddCommand(makePingCmd(dbc))
	root.AddCommand(makeMigrationCmd(dbc))
	dbc.Root = &root
	return dbc
}
