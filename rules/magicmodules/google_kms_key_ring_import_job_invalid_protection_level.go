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

// GoogleKmsKeyRingImportJobInvalidProtectionLevelRule checks the pattern is valid
type GoogleKmsKeyRingImportJobInvalidProtectionLevelRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleKmsKeyRingImportJobInvalidProtectionLevelRule returns new rule with default attributes
func NewGoogleKmsKeyRingImportJobInvalidProtectionLevelRule() *GoogleKmsKeyRingImportJobInvalidProtectionLevelRule {
	return &GoogleKmsKeyRingImportJobInvalidProtectionLevelRule{
		resourceType:  "google_kms_key_ring_import_job",
		attributeName: "protection_level",
	}
}

// Name returns the rule name
func (r *GoogleKmsKeyRingImportJobInvalidProtectionLevelRule) Name() string {
	return "google_kms_key_ring_import_job_invalid_protection_level"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleKmsKeyRingImportJobInvalidProtectionLevelRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleKmsKeyRingImportJobInvalidProtectionLevelRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleKmsKeyRingImportJobInvalidProtectionLevelRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleKmsKeyRingImportJobInvalidProtectionLevelRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"SOFTWARE", "HSM", "EXTERNAL"}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
