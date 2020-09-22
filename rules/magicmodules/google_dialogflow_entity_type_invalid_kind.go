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

// GoogleDialogflowEntityTypeInvalidKindRule checks the pattern is valid
type GoogleDialogflowEntityTypeInvalidKindRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleDialogflowEntityTypeInvalidKindRule returns new rule with default attributes
func NewGoogleDialogflowEntityTypeInvalidKindRule() *GoogleDialogflowEntityTypeInvalidKindRule {
	return &GoogleDialogflowEntityTypeInvalidKindRule{
		resourceType:  "google_dialogflow_entity_type",
		attributeName: "kind",
	}
}

// Name returns the rule name
func (r *GoogleDialogflowEntityTypeInvalidKindRule) Name() string {
	return "google_dialogflow_entity_type_invalid_kind"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleDialogflowEntityTypeInvalidKindRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleDialogflowEntityTypeInvalidKindRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleDialogflowEntityTypeInvalidKindRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleDialogflowEntityTypeInvalidKindRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		validateFunc := validation.StringInSlice([]string{"KIND_MAP", "KIND_LIST", "KIND_REGEXP"}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}