package validate

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func Load() {
	zh := zh.New()
	uni = ut.New(zh, zh)

	trans, _ = uni.GetTranslator("zh")

	//获取gin的校验器
	validate := binding.Validator.Engine().(*validator.Validate)
	//注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

func GetError(errs validator.ValidationErrors, r interface{}) string {
	s := reflect.TypeOf(r)
	for _, fieldError := range errs {
		filed, _ := s.FieldByName(fieldError.Field())
		errTag := fieldError.Tag() + "_msg"
		errTagText := filed.Tag.Get(errTag)
		errText := filed.Tag.Get("msg")
		if errTagText != "" {
			return errTagText
		}
		if errText != "" {
			return errText
		}
	}
	return ""
}

func Translate(err error, r interface{}) string {
	var result string
	self := GetError(err.(validator.ValidationErrors), r)
	if self != "" {
		return self
	}
	for _, err := range err.(validator.ValidationErrors) {
		result = result + err.Translate(trans)
	}
	return result
}
