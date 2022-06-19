package cache

import (
	"fmt"
	"strings"
	"sync"

	"brave-ox-web/common"
	"brave-ox-web/conf"
	"brave-ox-web/db"
	"brave-ox-web/log"
	"brave-ox-web/model"
)

var (
	_cattle       *sync.Map
	_profilePhoto *sync.Map
	_cattleSkin   *sync.Map
	_planetTypes  *sync.Map
	_box          *sync.Map
)

func init() {
	_cattle = new(sync.Map)
	_profilePhoto = new(sync.Map)
	_cattleSkin = new(sync.Map)
	_planetTypes = new(sync.Map)
	_box = new(sync.Map)
}

func LoadTemplateData() {
	// 加载静态配置
	loadCattleData()
	loadProfilePhotoData()
	loadCattleSkinData()
	loadPlanetData()
	loadBoxData()
}

func loadCattleData() {
	var cattleTemp []*model.CattleTemplate
	err := db.QueryAll(&cattleTemp)
	common.MustNoError(err)

	cattle := make(map[uint8]map[string]*model.CattleTemplate)
	for _, v := range cattleTemp {
		typeMap, ok := cattle[v.Class]
		if !ok {
			typeMap = make(map[string]*model.CattleTemplate)
			cattle[v.Class] = typeMap
		}
		typeMap[CattleTemplateKey(v.Class, v.Gender, v.Gsq, v.IsAdult)] = v
	}

	_cattle.Store("cattle", cattle)
}

func CattleTemplateKey(class uint8, gender uint8, gsq uint32, isAdult bool) string {
	var key string
	switch class {
	case conf.TypeBovineHero:
		key = fmt.Sprintf("%d-%d", gender, gsq)
	case conf.TypeCattle:
		if isAdult {
			key = fmt.Sprintf("%d-1", gender)
		} else {
			key = fmt.Sprintf("%d-0", gender)
		}
	default:
		panic("invalid cattle type")
	}
	return key
}

func GetCattleTemplate(class uint8, key string) (*model.CattleTemplate, error) {
	data, _ := _cattle.Load("cattle")
	cattle, ok := data.(map[uint8]map[string]*model.CattleTemplate)[class][key]
	if !ok {
		return nil, fmt.Errorf("cattten key:%s not found", key)
	}

	return cattle, nil
}

func loadProfilePhotoData() {
	var cattleTemp []*model.ProfilePhotoTemplate
	err := db.QueryAll(&cattleTemp)
	common.MustNoError(err)

	profilePhoto := make(map[uint32]*model.ProfilePhotoTemplate)
	for _, v := range cattleTemp {
		profilePhoto[v.Id] = v
	}

	_profilePhoto.Store("profile-photo", profilePhoto)
}

func GetProfileTemplate(id uint32) (*model.ProfilePhotoTemplate, error) {
	data, _ := _profilePhoto.Load("profile-photo")
	photo, ok := data.(map[uint32]*model.ProfilePhotoTemplate)[id]
	if !ok {
		return nil, fmt.Errorf("profile-photo :%d not found", id)
	}

	return photo, nil
}

func GetCattleImage(class uint8, gender uint8, gsq uint32, isAdult bool) string {
	cattleCfg, err := GetCattleTemplate(class, CattleTemplateKey(class, gender, gsq, isAdult))
	if err != nil {
		log.Error(err)
		return ""
	}
	return conf.ImageHosting() + cattleCfg.Image
}

func CattleName(class uint8, gender uint8, gsq uint32, isAdult bool) string {
	if class != conf.TypeBovineHero {
		return ""
	}
	cattleCfg, err := GetCattleTemplate(class, CattleTemplateKey(class, gender, gsq, isAdult))
	if err != nil {
		log.Error(err)
		return ""
	}

	arr := strings.Split(cattleCfg.NameCN, "#")
	if len(arr) == 2 {
		return arr[1]
	}
	return ""

}

func loadCattleSkinData() {
	var skinTemp []*model.CattleSkinTemplate
	err := db.QueryAll(&skinTemp)
	common.MustNoError(err)

	skin := make(map[uint32]*model.CattleSkinTemplate)
	for _, v := range skinTemp {
		skin[v.Id] = v
	}
	_cattleSkin.Store("skin", skin)
}

func GetSkinTemplate(id uint32) (*model.CattleSkinTemplate, error) {
	data, _ := _cattleSkin.Load("skin")
	skin, ok := data.(map[uint32]*model.CattleSkinTemplate)[id]
	if !ok {
		return nil, fmt.Errorf("skin:%d not found", id)
	}
	return skin, nil
}

func loadPlanetData() {
	var planet []*model.PlanetTemplate
	err := db.QueryAll(&planet)
	common.MustNoError(err)

	planetMap := make(map[uint8]*model.PlanetTemplate)
	for _, v := range planet {
		planetMap[v.Types] = v
	}
	_planetTypes.Store("planet", planetMap)
}

func GetPlanet(types uint8) (*model.PlanetTemplate, error) {
	data, _ := _planetTypes.Load("planet")
	planet, ok := data.(map[uint8]*model.PlanetTemplate)[types]
	if !ok {
		return nil, fmt.Errorf("planet type:%d not found", types)
	}
	return planet, nil
}

func loadBoxData() {
	var box []*model.BoxTemplate
	err := db.QueryAll(&box)
	common.MustNoError(err)

	boxMap := make(map[string]*model.BoxTemplate)
	for _, v := range box {
		boxMap[boxKey(v.Types, v.Contract)] = v
	}
	_box.Store("box", boxMap)
}

func boxKey(types uint32, address string) string {
	return fmt.Sprintf("%s-%d", address, types)
}

func GetBox(types uint32, address string) (*model.BoxTemplate, error) {
	data, _ := _box.Load("box")
	box, ok := data.(map[string]*model.BoxTemplate)[boxKey(types, address)]
	if !ok {
		return nil, fmt.Errorf("box:%d not found", types)
	}
	return box, nil
}
