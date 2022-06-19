package vo

type NftVOAttr struct {
	TraitType string `json:"trait_type,omitempty"`
	//DisplayType string      `json:"display_type"`
	Value interface{} `json:"value,omitempty"`
	//MaxValue    int         `json:"max_value"`
}

type NftVO struct {
	Name        string `json:"name"`
	ExternalUrl string `json:"external_url"`
	Image       string `json:"image"`
	Description string `json:"description"`
	//BackgroundColor string       `json:"background_color"`
	Attributes []*NftVOAttr `json:"attributes"`
}

func NewNFTVO(name string, image string, desc string) *NftVO {
	return &NftVO{
		Name:        name,
		Image:       image,
		Description: desc,
	}
}
