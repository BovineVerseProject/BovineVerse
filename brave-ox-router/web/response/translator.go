package response

import (
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var _trans ut.Translator

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		enT := en.New()

		uni := ut.New(enT, enT)

		var ok bool
		_trans, ok = uni.GetTranslator("en")
		if !ok {
			return
		}

		enTranslations.RegisterDefaultTranslations(v, _trans)
	}
	return
}

func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		key := field[strings.Index(field, ".")+1:]
		res[strings.ToLower(key)] = err[len(key):]
	}
	return res
}

func Trans() ut.Translator {
	return _trans
}
