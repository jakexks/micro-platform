package infra

import (
	"fmt"

	"github.com/spf13/viper"
)

// Resource is a deployment of shared resources
type Resource struct {
	Namespace string
}

// Steps generates a terraform plan to create shared resources
func (r *Resource) Steps(runID int32) ([]Step, error) {
	var s []Step
	vars := make(map[string]string)
	env := make(map[string]string)
	vars["resource_namespace"] = r.Namespace
	env["KUBECONFIG"] = viper.GetString("kube-config-path")
	s = append(s, Step{
		&TerraformModule{
			ID:        r.configName(),
			Name:      r.configName(),
			Env:       env,
			Variables: vars,
			Path:      fmt.Sprintf("/tmp/%s-%d", r.configName(), runID),
			Source:    "./infra/resource",
		},
	})
	return s, nil
}

func (r *Resource) configName() string {
	return fmt.Sprintf("%s-%s-%s-%s", viper.GetString("cluster-name"), viper.GetString("cluster-region"), viper.GetString("cloud-provider"), r.Namespace)
}
