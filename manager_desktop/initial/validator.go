package initial

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
	"simple_gateway/global"
	"strings"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// ValidateAdmin implements validator.Func
func ValidateAdmin(fl validator.FieldLevel) bool {
	return fl.Field().String() == "admin"
}

func ReplaceGinBinding(language string) {
	RegisterCustomValidator()
	RegisterTranslator(language)
}

func RegisterCustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("is-validuser", ValidateAdmin)
	}
}

func RegisterTranslator(language string) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		zh := zh.New()
		uni := ut.New(zh, en, zh)
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			jTag := field.Tag.Get("json")
			return jTag
		})
		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		global.Trans, _ = uni.GetTranslator("zh")
		switch language {
		case "en":
			en_translations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, global.Trans)
		default:
			zh_translations.RegisterDefaultTranslations(v, global.Trans)
		}
		//自定义验证方法
		//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
		v.RegisterValidation("is_valid_ip_list", func(fl validator.FieldLevel) bool {
			for _, rule := range strings.Split(fl.Field().String(), "\n") {
				match, _ := regexp.Match(`^\S+:\d+$`, []byte(rule))
				if !match {
					return false
				}
			}
			return true
		})

		//自定义翻译器
		//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
		v.RegisterTranslation("is_valid_ip_list", global.Trans, func(ut ut.Translator) error {
			return ut.Add("is_valid_ip_list", "{0} 不符合规则", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("is_valid_ip_list", fe.Field())
			return t
		})
		//自定义验证方法
		//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
		v.RegisterValidation("is_valid_weight_list", func(fl validator.FieldLevel) bool {
			for _, rule := range strings.Split(fl.Field().String(), ",") {
				match, _ := regexp.Match(`^\d+$`, []byte(rule))
				if !match {
					return false
				}
			}
			return true
		})

		//自定义翻译器
		//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
		v.RegisterTranslation("is_valid_weight_list", global.Trans, func(ut ut.Translator) error {
			return ut.Add("is_valid_weight_list", "{0} 不符合规则", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("is_valid_weight_list", fe.Field())
			return t
		})
		//自定义验证方法
		//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
		v.RegisterValidation("is_valid_header_transfor", func(fl validator.FieldLevel) bool {
			if fl.Field().String() == "" {
				return true
			}
			for _, rule := range strings.Split(fl.Field().String(), "\n") {
				if len(strings.Split(rule, " ")) != 3 {
					return false
				}
			}
			return true
		})

		//自定义翻译器
		//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
		v.RegisterTranslation("is_valid_header_transfor", global.Trans, func(ut ut.Translator) error {
			return ut.Add("is_valid_header_transfor", "{0} 不符合规则", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("is_valid_header_transfor", fe.Field())
			return t
		})
		//自定义验证方法
		//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
		v.RegisterValidation("is_valid_url_rewrite", func(fl validator.FieldLevel) bool {
			if fl.Field().String() == "" {
				return true
			}
			for _, rule := range strings.Split(fl.Field().String(), "\n") {
				if len(strings.Split(rule, " ")) != 2 {
					return false
				}
			}
			return true
		})

		//自定义翻译器
		//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
		v.RegisterTranslation("is_valid_url_rewrite", global.Trans, func(ut ut.Translator) error {
			return ut.Add("is_valid_url_rewrite", "{0} 格式错误", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("is_valid_url_rewrite", fe.Field())
			return t
		})
		//自定义验证方法
		//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
		v.RegisterValidation("is_valid_rule", func(fl validator.FieldLevel) bool {
			match, _ := regexp.Match(`^\S+$`, []byte(fl.Field().String()))
			return match
		})

		//自定义翻译器
		//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
		v.RegisterTranslation("is_valid_rule", global.Trans, func(ut ut.Translator) error {
			return ut.Add("is_valid_rule", "{0} 需要非空", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("is_valid_rule", fe.Field())
			return t
		})

		//自定义验证方法
		//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
		v.RegisterValidation("match_service_name", func(fl validator.FieldLevel) bool {
			match, _ := regexp.Match(`^[a-zA-z0-9_]{6,128}$`, []byte(fl.Field().String()))
			return match
		})

		//自定义翻译器
		//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
		v.RegisterTranslation("match_service_name", global.Trans, func(ut ut.Translator) error {
			return ut.Add("match_service_name", "{0} 需要满足6-128位字母数字下划线", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("match_service_name", fe.Field())
			return t
		})

		//自定义验证方法
		//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
		v.RegisterValidation("is-validuser", func(fl validator.FieldLevel) bool {
			return fl.Field().String() == "admin"
		})

		//自定义翻译器
		//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
		v.RegisterTranslation("is-validuser", global.Trans, func(ut ut.Translator) error {
			return ut.Add("is-validuser", "{0} 填写不正确哦", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("is-validuser", fe.Field())
			return t
		})
	}
}
