package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var _braveox *config

func init() {
	_braveox = new(config)
}

type config struct {
	Chain         string
	SensitiveWord string   `toml:"sensitive_word"`
	AllowIp       []string `toml:"allow_ip"`
	ImageHosting  string

	Banker    map[string]*AddressInfo
	Contracts map[string]map[string]*AddressInfo
	Oss       *Oss
	OtaSecret string `toml:"ota_secret"`
}

type Database struct {
	Host   string
	Port   uint32
	User   string
	Pwd    string
	DBName string `toml:"db_name"`
	Param  string
}

type AddressInfo struct {
	Address string
	Key     string
}

type Oss struct {
	Endpoint  string `toml:"endpoint"`
	AccessID  string `toml:"access_id"`
	AccessKey string `toml:"access_key"`
	Bucket    string `toml:"bucket"`
}

func Parse(path string) {
	_, err := toml.DecodeFile(path, _braveox)
	if err != nil {
		panic(err)
	}

	_braveox.ImageHosting = "https://" + _braveox.Oss.Bucket + "." + _braveox.Oss.Endpoint
}

func ChainContract(chain string) []string {
	d, ok := _braveox.Contracts[chain]
	if !ok {
		panic(fmt.Errorf("Invalid chain contract config, chain:%s", chain))
	}

	var contracts []string
	for _, v := range d {
		contracts = append(contracts, v.Address)
	}

	return contracts
}

func ImageHosting() string {
	return _braveox.ImageHosting
}

func GetOssConf() *Oss {
	return _braveox.Oss
}

func SensitivePath() string {
	return _braveox.SensitiveWord
}

func IPWhiteList() []string {
	return _braveox.AllowIp
}

func OTASecret() string {
	return _braveox.OtaSecret
}

func AirdropBanker() *AddressInfo {
	return _braveox.Banker["airdrop"]
}

func CattleContract() string {
	return _braveox.Contracts[_braveox.Chain]["cattle"].Address
}

func QueryAPIContract() string {
	return _braveox.Contracts[_braveox.Chain]["query"].Address
}

func PlanetContract() string {
	return _braveox.Contracts[_braveox.Chain]["planet"].Address
}

func Planet721Contract() string {
	return _braveox.Contracts[_braveox.Chain]["planet721"].Address
}

func Item1155Contract() string {
	return _braveox.Contracts[_braveox.Chain]["item"].Address
}

func BreedCenterContract() string {
	return _braveox.Contracts[_braveox.Chain]["breed"].Address
}

// 牛栏
func StableContract() string {
	return _braveox.Contracts[_braveox.Chain]["stable"].Address
}

func Market721() string {
	return _braveox.Contracts[_braveox.Chain]["market721"].Address
}

func MainChain() string {
	return _braveox.Chain
}

func AirdropContract() string {
	return _braveox.Contracts["bsc"]["airdrop"].Address
}

func InvitationContract() string {
	return _braveox.Contracts[_braveox.Chain]["invite"].Address
}

func CompoundContract() string {
	return _braveox.Contracts[_braveox.Chain]["compound"].Address
}

func MailContract() string {
	return _braveox.Contracts[_braveox.Chain]["mail"].Address
}
