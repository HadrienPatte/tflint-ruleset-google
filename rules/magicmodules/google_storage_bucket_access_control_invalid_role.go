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

// GoogleStorageBucketAccessControlInvalidRoleRule checks the pattern is valid
type GoogleStorageBucketAccessControlInvalidRoleRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleStorageBucketAccessControlInvalidRoleRule returns new rule with default attributes
func NewGoogleStorageBucketAccessControlInvalidRoleRule() *GoogleStorageBucketAccessControlInvalidRoleRule {
	return &GoogleStorageBucketAccessControlInvalidRoleRule{
		resourceType:  "google_storage_bucket_access_control",
		attributeName: "role",
	}
}

// Name returns the rule name
func (r *GoogleStorageBucketAccessControlInvalidRoleRule) Name() string {
	return "google_storage_bucket_access_control_invalid_role"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleStorageBucketAccessControlInvalidRoleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleStorageBucketAccessControlInvalidRoleRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleStorageBucketAccessControlInvalidRoleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleStorageBucketAccessControlInvalidRoleRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		validateFunc := validation.StringInSlice([]string{"OWNER", "READER", "WRITER", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
