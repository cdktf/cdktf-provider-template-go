package dir

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DirConfig struct {
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
	// Path to the directory where the templated files will be written.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/template/r/dir#destination_dir Dir#destination_dir}
	DestinationDir *string `field:"required" json:"destinationDir" yaml:"destinationDir"`
	// Path to the directory where the files to template reside.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/template/r/dir#source_dir Dir#source_dir}
	SourceDir *string `field:"required" json:"sourceDir" yaml:"sourceDir"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/template/r/dir#id Dir#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Variables to substitute.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/template/r/dir#vars Dir#vars}
	Vars *map[string]*string `field:"optional" json:"vars" yaml:"vars"`
}

