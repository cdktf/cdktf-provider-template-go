// Prebuilt template Provider for Terraform CDK (cdktf)
package template

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudinitConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// part block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/template/r/cloudinit_config#part CloudinitConfig#part}
	Part interface{} `field:"required" json:"part" yaml:"part"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/template/r/cloudinit_config#base64_encode CloudinitConfig#base64_encode}.
	Base64Encode interface{} `field:"optional" json:"base64Encode" yaml:"base64Encode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/template/r/cloudinit_config#gzip CloudinitConfig#gzip}.
	Gzip interface{} `field:"optional" json:"gzip" yaml:"gzip"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/template/r/cloudinit_config#id CloudinitConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

