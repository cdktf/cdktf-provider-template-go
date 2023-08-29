// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datatemplatecloudinitconfig


type DataTemplateCloudinitConfigPart struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/template/2.2.0/docs/data-sources/cloudinit_config#content DataTemplateCloudinitConfig#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/template/2.2.0/docs/data-sources/cloudinit_config#content_type DataTemplateCloudinitConfig#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/template/2.2.0/docs/data-sources/cloudinit_config#filename DataTemplateCloudinitConfig#filename}.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/template/2.2.0/docs/data-sources/cloudinit_config#merge_type DataTemplateCloudinitConfig#merge_type}.
	MergeType *string `field:"optional" json:"mergeType" yaml:"mergeType"`
}

