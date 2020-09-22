package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-google/project"
	"github.com/terraform-linters/tflint-ruleset-google/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: tflint.RuleSet{
			Name:    "google",
			Version: project.Version,
			Rules:   rules.Rules,
		},
	})
}
