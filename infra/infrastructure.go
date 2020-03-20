// Package infrastructure provides functions for orchestrating a Micro platform
package infra

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Task describes an individual task
type Task interface {
	Validate() error
	Plan() error
	Apply() error
	Finalise() error
	Destroy() error
}

// Step is a list of parallisable tasks
type Step []Task

// Platform defines a complete platform
type Platform struct {
	Name    string
	Domain  string
	Gslb    string
	Kv      string
	Regions []struct {
		Provider string
		Region   string
		Control  []string
		Resource []string
		Network  []string
	}
}

// Steps generates an action plan from a Platform description
func (p *Platform) Steps() ([]Step, error) {
	// Not secure random, it doesn't matter as it's only to generate non colliding directory names
	rand.Seed(time.Now().UnixNano())
	dirSuffix := rand.Int31()
	var steps []Step
	// 1: Ensure Remote state is available
	steps = append(steps, Step{&RemoteState{ID: p.Name + "-check-remote-state", Name: p.Name + "-check-remote-state"}})

	// 2: Set up KV namespace
	steps = append(steps, Step{
		&TerraformModule{
			ID:     p.Name + "-global-kv",
			Name:   p.Name + "-global-kv",
			Source: "./infra/kv/" + p.Kv,
			Path:   fmt.Sprintf("/tmp/%s-%d", p.Name+"-kv", dirSuffix),
		},
	})

	for _, r := range p.Regions {
		// 2.1 Create Kubernetes cluster
		steps = append(steps, Step{
			&TerraformModule{
				ID:     p.Name + "-" + r.Region + "-" + r.Provider + "-k8s",
				Name:   p.Name + "-" + r.Region + "-" + r.Provider + "-k8s",
				Source: "./infra/kubernetes/" + r.Provider,
				Path:   fmt.Sprintf("/tmp/%s-%s-%s-%d", p.Name, r.Region, r.Provider, dirSuffix),
			},
		})

		// 2.2 Grab Kubernetes config from the configured cluster
		vars := make(map[string]string)
		vars["kubernetes"] = r.Provider
		vars["args"] = fmt.Sprintf(`["%s", "%s"]`, p.Name+"-"+r.Region+"-"+r.Provider+"-k8s", "eu-west-2")
		steps = append(steps, Step{
			&TerraformModule{
				ID:        p.Name + "-" + r.Region + "-" + r.Provider + "-kubeconfig",
				Name:      p.Name + "-" + r.Region + "-" + r.Provider + "-kubeconfig",
				Source:    "./infra/kubernetes/kubeconfig",
				Path:      fmt.Sprintf("/tmp/%s-%s-%s-kubeconfig-%d", p.Name, r.Region, r.Provider, dirSuffix),
				Variables: vars,
			},
		})

		// 2.3 Create namespaces
		vars = make(map[string]string)
		env := make(map[string]string)
		vars["control_namespace"] = strings.ToLower(fmt.Sprintf("%s-control", p.Name))
		vars["resource_namespace"] = strings.ToLower(fmt.Sprintf("%s-resource", p.Name))
		vars["network_namespace"] = strings.ToLower(fmt.Sprintf("%s-network", p.Name))
		env["KUBECONFIG"] = fmt.Sprintf("/tmp/%s-%s-%s-kubeconfig-%d/kubeconfig", p.Name, r.Region, r.Provider, dirSuffix)
		steps = append(steps, Step{
			&TerraformModule{
				ID:        p.Name + "-" + r.Region + "-" + r.Provider + "-namespaces",
				Name:      p.Name + "-" + r.Region + "-" + r.Provider + "-namespaces",
				Source:    "./infra/kubernetes/namespaces",
				Path:      fmt.Sprintf("/tmp/%s-%s-%s-namespaces-%d", p.Name, r.Region, r.Provider, dirSuffix),
				Variables: vars,
				Env:       env,
			},
		})

		// 2.4 Create shared resources
		vars = make(map[string]string)
		env = make(map[string]string)
		remoteStates := make(map[string]string)
		if r.Provider == "aws" {
			vars["in_aws"] = "true"
		} else {
			vars["in_aws"] = "false"
		}
		env["KUBECONFIG"] = fmt.Sprintf("/tmp/%s-%s-%s-kubeconfig-%d/kubeconfig", p.Name, r.Region, r.Provider, dirSuffix)
		remoteStates["namespaces"] = p.Name + "-" + r.Region + "-" + r.Provider + "-namespaces"
		steps = append(steps, Step{
			&TerraformModule{
				ID:           p.Name + "-" + r.Region + "-" + r.Provider + "-resource",
				Name:         p.Name + "-" + r.Region + "-" + r.Provider + "-resource",
				Source:       "./infra/resource",
				Path:         fmt.Sprintf("/tmp/%s-%s-%s-resource-%d", p.Name, r.Region, r.Provider, dirSuffix),
				Variables:    vars,
				Env:          env,
				RemoteStates: remoteStates,
			},
		})

		// 2.5 Create control plane
		vars = make(map[string]string)
		env = make(map[string]string)
		remoteStates = make(map[string]string)
		vars["domain_name"] = p.Domain
		env["KUBECONFIG"] = fmt.Sprintf("/tmp/%s-%s-%s-kubeconfig-%d/kubeconfig", p.Name, r.Region, r.Provider, dirSuffix)
		remoteStates["namespaces"] = p.Name + "-" + r.Region + "-" + r.Provider + "-namespaces"
		steps = append(steps, Step{
			&TerraformModule{
				ID:           p.Name + "-" + r.Region + "-" + r.Provider + "-control",
				Name:         p.Name + "-" + r.Region + "-" + r.Provider + "-control",
				Source:       "./infra/control",
				Path:         fmt.Sprintf("/tmp/%s-%s-%s-control-%d", p.Name, r.Region, r.Provider, dirSuffix),
				Variables:    vars,
				Env:          env,
				RemoteStates: remoteStates,
			},
		})

		// 2.6 Create network
		vars = make(map[string]string)
		env = make(map[string]string)
		remoteStates = make(map[string]string)
		vars["domain_name"] = p.Domain
		vars["cloudflare_account_id"] = "TODO"
		vars["cloudflare_dns_zone_id"] = "TODO"
		vars["cloudflare_api_token"] = "TODO"
		vars["region_slug"] = r.Region + "-" + r.Provider
		env["KUBECONFIG"] = fmt.Sprintf("/tmp/%s-%s-%s-kubeconfig-%d/kubeconfig", p.Name, r.Region, r.Provider, dirSuffix)
		remoteStates["namespaces"] = p.Name + "-" + r.Region + "-" + r.Provider + "-namespaces"
		remoteStates["kv"] = p.Name + "-global-kv"
		steps = append(steps, Step{
			&TerraformModule{
				ID:           p.Name + "-" + r.Region + "-" + r.Provider + "-network",
				Name:         p.Name + "-" + r.Region + "-" + r.Provider + "-network",
				Source:       "./infra/network",
				Path:         fmt.Sprintf("/tmp/%s-%s-%s-network-%d", p.Name, r.Region, r.Provider, dirSuffix),
				Variables:    vars,
				Env:          env,
				RemoteStates: remoteStates,
			},
		})
	}

	return steps, nil
}

// ExecutePlan carries out a plan on steps
func ExecutePlan(steps []Step) error {
	for _, step := range steps {
		for _, t := range step {
			defer t.Finalise()
			if err := t.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

// ExecuteApply carries out an apply on steps
func ExecuteApply(steps []Step) error {
	for _, step := range steps {
		for _, t := range step {
			defer t.Finalise()
			if err := t.Validate(); err != nil {
				return err
			}
			if err := t.Apply(); err != nil {
				return err
			}
		}
	}
	return nil
}

// ExecuteDestroy destroys steps
func ExecuteDestroy(steps []Step) error {
	// Find any kubeconfig steps; we need them to destroy the resources
	for _, s := range steps {
		for _, task := range s {
			switch t := task.(type) {
			case *TerraformModule:
				if strings.Contains(t.Source, "kubeconfig") {
					defer t.Finalise()
					if err := t.Validate(); err != nil {
						return err
					}
					if err := t.Apply(); err != nil {
						return err
					}
					t.Variables["kubernetes"] = "none"
					defer t.Destroy()
				}
			}
		}
	}
	for i := len(steps) - 1; i >= 0; i-- {
		for _, task := range steps[i] {
			switch t := task.(type) {
			case *TerraformModule:
				// Skip any kubeconfig steps
				if !strings.Contains(t.Source, "kubeconfig") {
					defer t.Finalise()
					if err := t.Validate(); err != nil {
						return err
					}
					if err := t.Destroy(); err != nil {
						return err
					}
				}
			default:
				defer t.Finalise()
				if err := t.Validate(); err != nil {
					return err
				}
				if err := t.Destroy(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}