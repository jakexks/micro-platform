package infra

import (
	"fmt"

	"github.com/spf13/viper"
)

// Network is a deployment of a micro platform
type Network struct {
	Namespace string
}

// Steps generates a terraform plan to create a network
func (n *Network) Steps(runID int32) ([]Step, error) {
	var s []Step
	vars := make(map[string]string)
	env := make(map[string]string)
	vars["network_namespace"] = n.Namespace
	vars["resource_namespace"] = viper.GetString("resource-namespace")
	vars["domain_name"] = viper.GetString("domain-name")
	env["KUBECONFIG"] = viper.GetString("kube-config-path")
	s = append(s, Step{
		&TerraformModule{
			ID:        n.configName(),
			Name:      n.configName(),
			Env:       env,
			Variables: vars,
			Path:      fmt.Sprintf("/tmp/%s-%d", n.configName(), runID),
			Source:    "./infra/network",
		},
	})
	return s, nil
}

func (n *Network) configName() string {
	return fmt.Sprintf("%s-%s-%s-%s", viper.GetString("cluster-name"), viper.GetString("cluster-region"), viper.GetString("cloud-provider"), n.Namespace)
}
