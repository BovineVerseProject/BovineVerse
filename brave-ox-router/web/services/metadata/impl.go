package metadata

import (
	"brave-ox-web/conf"
	"brave-ox-web/log"
	"brave-ox-web/web/cache"
	"brave-ox-web/web/services/cattle"
	"brave-ox-web/web/vo"
	"fmt"
)

type metadataService struct{}

func (*metadataService) Cattle(id uint32) *vo.NftVO {
	info, err := cattle.Service().GetCattleById(id)
	if err != nil {
		log.Errorf("Get cattle err:%v", err)
		return nil
	}

	if info == nil {
		return nil
	}

	tmp, _ := cache.GetCattleTemplate(info.Class, cache.CattleTemplateKey(info.Class, info.Gender, info.GenderSeq, info.IsAdult))
	if tmp == nil {
		return nil
	}
	if info.Class == conf.TypeBovineHero {
		tmp.NameEn = fmt.Sprintf("Bovine Hero #%d", info.Id)
	}

	return vo.NewNFTVO(tmp.NameEn, cache.GetCattleImage(info.Class, info.Gender, info.GenderSeq, info.IsAdult), tmp.Description)
}

func (*metadataService) Skin(id uint32) *vo.NftVO {
	skin, err := cache.GetSkinTemplate(id)
	if err != nil {
		log.Errorf("skin config err:%v", err)
		return nil
	}

	attrs := make([]*vo.NftVOAttr, 1)
	attrs[0] = &vo.NftVOAttr{
		TraitType: "level",
		Value:     skin.Level,
	}
	if skin.Attack > 0 {
		attrs = append(attrs, &vo.NftVOAttr{
			TraitType: "attack",
			Value:     skin.Attack,
		})
	}

	if skin.Defense > 0 {
		attrs = append(attrs, &vo.NftVOAttr{
			TraitType: "defense",
			Value:     skin.Defense,
		})
	}

	if skin.Stamina > 0 {
		attrs = append(attrs, &vo.NftVOAttr{
			TraitType: "stamina",
			Value:     skin.Stamina,
		})
	}

	if skin.Life > 0 {
		attrs = append(attrs, &vo.NftVOAttr{
			TraitType: "life",
			Value:     skin.Life,
		})
	}

	nft := &vo.NftVO{
		Name:        skin.NameEn,
		Image:       conf.ImageHosting() + skin.Image,
		Description: skin.Description,
		Attributes:  attrs,
	}
	return nft
}

func (*metadataService) Planet(types uint32) *vo.NftVO {
	typesConf, err := cache.GetPlanet(uint8(types))
	if err != nil {
		log.Errorf("planet config err:%v", err)
		return nil
	}

	meta := &vo.NftVO{
		Name:        typesConf.NameEn,
		Description: typesConf.Description,
		Image:       typesConf.Image,
	}

	return meta
}

func (this *metadataService) Box(types uint32, address string) *vo.NftVO {
	b, err := cache.GetBox(types, address)
	if err != nil {
		log.Errorf("box config err:%v", err)
		return nil
	}

	return vo.NewNFTVO(b.Name, b.Image, b.Description)
}
