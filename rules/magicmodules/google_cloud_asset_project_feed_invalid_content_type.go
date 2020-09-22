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

// GoogleCloudAssetProjectFeedInvalidContentTypeRule checks the pattern is valid
type GoogleCloudAssetProjectFeedInvalidContentTypeRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleCloudAssetProjectFeedInvalidContentTypeRule returns new rule with default attributes
func NewGoogleCloudAssetProjectFeedInvalidContentTypeRule() *GoogleCloudAssetProjectFeedInvalidContentTypeRule {
	return &GoogleCloudAssetProjectFeedInvalidContentTypeRule{
		resourceType:  "google_cloud_asset_project_feed",
		attributeName: "content_type",
	}
}

// Name returns the rule name
func (r *GoogleCloudAssetProjectFeedInvalidContentTypeRule) Name() string {
	return "google_cloud_asset_project_feed_invalid_content_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleCloudAssetProjectFeedInvalidContentTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleCloudAssetProjectFeedInvalidContentTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleCloudAssetProjectFeedInvalidContentTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleCloudAssetProjectFeedInvalidContentTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		validateFunc := validation.StringInSlice([]string{"CONTENT_TYPE_UNSPECIFIED", "RESOURCE", "IAM_POLICY", "ORG_POLICY", "ACCESS_POLICY", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
