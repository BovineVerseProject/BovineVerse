package metadata

import "brave-ox-web/web/vo"

var metadataSvc MetadataService

func Init() {
	metadataSvc = &metadataService{}
}

func Service() MetadataService {
	return metadataSvc
}

type MetadataService interface {
	Cattle(id uint32) *vo.NftVO
	Skin(id uint32) *vo.NftVO
	Box(types uint32, address string) *vo.NftVO
	Planet(types uint32) *vo.NftVO
}
