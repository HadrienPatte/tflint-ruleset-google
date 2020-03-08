package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-google/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: tflint.RuleSet{
			Name:    "google",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewGoogleComputeInstanceExampleTypeRule(),
			},
		},
	})
}
