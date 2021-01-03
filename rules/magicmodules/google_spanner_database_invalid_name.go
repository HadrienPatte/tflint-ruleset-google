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
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleSpannerDatabaseInvalidNameRule checks the pattern is valid
type GoogleSpannerDatabaseInvalidNameRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleSpannerDatabaseInvalidNameRule returns new rule with default attributes
func NewGoogleSpannerDatabaseInvalidNameRule() *GoogleSpannerDatabaseInvalidNameRule {
	return &GoogleSpannerDatabaseInvalidNameRule{
		resourceType:  "google_spanner_database",
		attributeName: "name",
	}
}

// Name returns the rule name
func (r *GoogleSpannerDatabaseInvalidNameRule) Name() string {
	return "google_spanner_database_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleSpannerDatabaseInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleSpannerDatabaseInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleSpannerDatabaseInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleSpannerDatabaseInvalidNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validateRegexp(`^[a-z][a-z0-9_\-]*[a-z0-9]$`)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
