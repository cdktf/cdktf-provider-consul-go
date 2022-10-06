package catalogentry


type CatalogEntryService struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/catalog_entry#name CatalogEntry#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/catalog_entry#address CatalogEntry#address}.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/catalog_entry#id CatalogEntry#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/catalog_entry#port CatalogEntry#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/consul/r/catalog_entry#tags CatalogEntry#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

