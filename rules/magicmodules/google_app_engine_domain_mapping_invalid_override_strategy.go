// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package magicmodules

import (
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleAppEngineDomainMappingInvalidOverrideStrategyRule checks the pattern is valid
type GoogleAppEngineDomainMappingInvalidOverrideStrategyRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleAppEngineDomainMappingInvalidOverrideStrategyRule returns new rule with default attributes
func NewGoogleAppEngineDomainMappingInvalidOverrideStrategyRule() *GoogleAppEngineDomainMappingInvalidOverrideStrategyRule {
	return &GoogleAppEngineDomainMappingInvalidOverrideStrategyRule{
		resourceType:  "google_app_engine_domain_mapping",
		attributeName: "override_strategy",
	}
}

// Name returns the rule name
func (r *GoogleAppEngineDomainMappingInvalidOverrideStrategyRule) Name() string {
	return "google_app_engine_domain_mapping_invalid_override_strategy"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleAppEngineDomainMappingInvalidOverrideStrategyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleAppEngineDomainMappingInvalidOverrideStrategyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleAppEngineDomainMappingInvalidOverrideStrategyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleAppEngineDomainMappingInvalidOverrideStrategyRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		validateFunc := validation.StringInSlice([]string{"STRICT", "OVERRIDE", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
