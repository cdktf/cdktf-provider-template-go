package cloudinitconfig


type CloudinitConfigPart struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/template/2.2.0/docs/resources/cloudinit_config#content CloudinitConfig#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/template/2.2.0/docs/resources/cloudinit_config#content_type CloudinitConfig#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/template/2.2.0/docs/resources/cloudinit_config#filename CloudinitConfig#filename}.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/template/2.2.0/docs/resources/cloudinit_config#merge_type CloudinitConfig#merge_type}.
	MergeType *string `field:"optional" json:"mergeType" yaml:"mergeType"`
}

