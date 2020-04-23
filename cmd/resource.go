package cmd

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/micro/platform/infra"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	resourceCommand = &cobra.Command{
		Use:   "resources",
		Short: "Deploy shared resources",
		Long:  `Deploy shared resources`,
	}

	resourceCreateCommand = &cobra.Command{
		Use:   "create",
		Short: "Create a shared resources",
		Long:  "Create a shared resources",
		Run: func(cmd *cobra.Command, args []string) {
			r, err := makeResources()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%+v\n", err)
				os.Exit(1)
			}
			err = infra.ExecuteApply(r)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%+v\n", err)
				os.Exit(1)
			}
		},
	}

	resourceDestroyCommand = &cobra.Command{
		Use:   "destroy",
		Short: "Destroy shared resources",
		Long:  "Destroy shared resources",
		Run: func(cmd *cobra.Command, args []string) {
			r, err := makeResources()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%+v\n", err)
				os.Exit(1)
			}
			err = infra.ExecuteDestroy(r)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%+v\n", err)
				os.Exit(1)
			}
		},
	}
)

func makeResources() ([]infra.Step, error) {
	r := &infra.Resource{
		Namespace: viper.GetString("resource-namespace"),
	}
	return r.Steps(rand.Int31())
}

func init() {
	rootCmd.AddCommand(resourceCommand)
	resourceCommand.AddCommand(resourceCreateCommand, resourceDestroyCommand)
	resourceCommand.PersistentFlags().String("resource-namespace", "resource", "namespace to deploy shared resources")
	viper.BindPFlag("resource-namespace", resourceCommand.PersistentFlags().Lookup("resource-namespace"))
	resourceCommand.PersistentFlags().String("platform-namespace", "network", "namespace to deploy micro platform")
	viper.BindPFlag("platform-namespace", resourceCommand.PersistentFlags().Lookup("platform-namespace"))
}
