// Prebuilt template Provider for Terraform CDK (cdktf)
package template


type CloudinitConfigPart struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/template/r/cloudinit_config#content CloudinitConfig#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/template/r/cloudinit_config#content_type CloudinitConfig#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/template/r/cloudinit_config#filename CloudinitConfig#filename}.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/template/r/cloudinit_config#merge_type CloudinitConfig#merge_type}.
	MergeType *string `field:"optional" json:"mergeType" yaml:"mergeType"`
}

