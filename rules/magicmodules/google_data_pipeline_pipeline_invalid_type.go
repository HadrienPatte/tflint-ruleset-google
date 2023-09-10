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

// GoogleDataPipelinePipelineInvalidTypeRule checks the pattern is valid
type GoogleDataPipelinePipelineInvalidTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleDataPipelinePipelineInvalidTypeRule returns new rule with default attributes
func NewGoogleDataPipelinePipelineInvalidTypeRule() *GoogleDataPipelinePipelineInvalidTypeRule {
	return &GoogleDataPipelinePipelineInvalidTypeRule{
		resourceType:  "google_data_pipeline_pipeline",
		attributeName: "type",
	}
}

// Name returns the rule name
func (r *GoogleDataPipelinePipelineInvalidTypeRule) Name() string {
	return "google_data_pipeline_pipeline_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleDataPipelinePipelineInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleDataPipelinePipelineInvalidTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleDataPipelinePipelineInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleDataPipelinePipelineInvalidTypeRule) Check(runner tflint.Runner) error {
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
			validateFunc := validation.StringInSlice([]string{"PIPELINE_TYPE_UNSPECIFIED", "PIPELINE_TYPE_BATCH", "PIPELINE_TYPE_STREAMING"}, false)

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