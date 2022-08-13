// Prebuilt template Provider for Terraform CDK (cdktf)
package template


type TemplateProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/template#alias TemplateProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

