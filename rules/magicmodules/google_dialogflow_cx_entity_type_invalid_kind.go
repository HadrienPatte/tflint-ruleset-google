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

// GoogleDialogflowCxEntityTypeInvalidKindRule checks the pattern is valid
type GoogleDialogflowCxEntityTypeInvalidKindRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleDialogflowCxEntityTypeInvalidKindRule returns new rule with default attributes
func NewGoogleDialogflowCxEntityTypeInvalidKindRule() *GoogleDialogflowCxEntityTypeInvalidKindRule {
	return &GoogleDialogflowCxEntityTypeInvalidKindRule{
		resourceType:  "google_dialogflow_cx_entity_type",
		attributeName: "kind",
	}
}

// Name returns the rule name
func (r *GoogleDialogflowCxEntityTypeInvalidKindRule) Name() string {
	return "google_dialogflow_cx_entity_type_invalid_kind"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleDialogflowCxEntityTypeInvalidKindRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleDialogflowCxEntityTypeInvalidKindRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleDialogflowCxEntityTypeInvalidKindRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleDialogflowCxEntityTypeInvalidKindRule) Check(runner tflint.Runner) error {
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

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"KIND_MAP", "KIND_LIST", "KIND_REGEXP"}, false)

		err = runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssue(r, err.Error(), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
