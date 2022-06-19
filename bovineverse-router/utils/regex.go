package utils

import "regexp"

var (
	specialRegex    = regexp.MustCompile("[`~!@#$%^*()+=|{}':;',\\[\\].;·/?~！@#￥%……*（）——+|{}【】‘；：”“’。，、？]")
	chineseNameRegx = regexp.MustCompile("[\u4e00-\u9fa5]")
)

func SpecialRegexCheck(c string) bool {
	return specialRegex.MatchString(c)
}

func ChineseNameCheck(c string) bool {
	return chineseNameRegx.MatchString(c)
}
