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
	networkCommand = &cobra.Command{
		Use:   "network",
		Short: "Deploy a micro network",
		Long:  `Deploy a micro network`,
	}

	networkCreateCommand = &cobra.Command{
		Use:   "create",
		Short: "Create a micro network",
		Long:  "Create a micro network",
		Run: func(cmd *cobra.Command, args []string) {
			r, err := makeNetwork()
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

	networkDestroyCommand = &cobra.Command{
		Use:   "destroy",
		Short: "Destroy micro network",
		Long:  "Destroy micro network",
		Run: func(cmd *cobra.Command, args []string) {
			r, err := makeNetwork()
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

func makeNetwork() ([]infra.Step, error) {
	n := &infra.Network{
		Namespace: viper.GetString("platform-namespace"),
	}
	return n.Steps(rand.Int31())
}

func init() {
	rootCmd.AddCommand(networkCommand)
	networkCommand.AddCommand(networkCreateCommand, networkDestroyCommand)
	networkCommand.PersistentFlags().String("resource-namespace", "resource", "namespace to deploy micro network")
	viper.BindPFlag("resource-namespace", networkCommand.PersistentFlags().Lookup("resource-namespace"))
	networkCommand.PersistentFlags().String("platform-namespace", "network", "namespace to deploy micro platform")
	viper.BindPFlag("platform-namespace", networkCommand.PersistentFlags().Lookup("platform-namespace"))
	networkCommand.PersistentFlags().String("domain-name", "m30.dev", "platform domain name")
	viper.BindPFlag("domain-name", networkCommand.PersistentFlags().Lookup("domain-name"))
}
