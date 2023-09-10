// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeRegionSslPolicyInvalidProfileRule checks the pattern is valid
type GoogleComputeRegionSslPolicyInvalidProfileRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleComputeRegionSslPolicyInvalidProfileRule returns new rule with default attributes
func NewGoogleComputeRegionSslPolicyInvalidProfileRule() *GoogleComputeRegionSslPolicyInvalidProfileRule {
	return &GoogleComputeRegionSslPolicyInvalidProfileRule{
		resourceType:  "google_compute_region_ssl_policy",
		attributeName: "profile",
	}
}

// Name returns the rule name
func (r *GoogleComputeRegionSslPolicyInvalidProfileRule) Name() string {
	return "google_compute_region_ssl_policy_invalid_profile"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeRegionSslPolicyInvalidProfileRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeRegionSslPolicyInvalidProfileRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeRegionSslPolicyInvalidProfileRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleComputeRegionSslPolicyInvalidProfileRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(val string) error {
			validateFunc := validation.StringInSlice([]string{"COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM", ""}, false)

			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				if err := runner.EmitIssue(r, err.Error(), attribute.Expr.Range()); err != nil {
					return err
				}
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}