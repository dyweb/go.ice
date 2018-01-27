# Cobra

- `cmd.Help()` shows the `Long` text at top
- `cmd.Usage()` is same as help except no description

````go
var rootCmd = &cobra.Command{
	Use:   myname,
	Short: "icehub daemon",
	Long:  "IceHub is an example GitHub integration service using go.ice",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
		//IceHub is an example GitHub integration service using go.ice
		//
		//Usage:
		//	icehubd [flags]
		//	icehubd [command]
		//
		//	Available Commands:
		//	help        Help about any command
		//	version     print version
		//
		//Flags:
		//	-h, --help   help for icehubd
		//
		//Use "icehubd [command] --help" for more information about a command.

		// usage does not have the long description like help
		//cmd.Usage()
		//Usage:
		//	icehubd [flags]
		//	icehubd [command]
		//
		//	Available Commands:
		//	help        Help about any command
		//	version     print version
		//
		//Flags:
		//	-h, --help   help for icehubd
		//
		//	Use "icehubd [command] --help" for more information about a command.
	},
}
````