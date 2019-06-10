package validator

import (
	"errors"
	"reflect"
	"strings"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"

	"github.com/go-playground/locales/zh"

	ut "github.com/go-playground/universal-translator"

	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	trans    ut.Translator
	validate = validator.New()
)

func init() {
	zh := zh.New()
	uni = ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")
	v := &defaultValidator{}
	v.lazyinit()
	binding.Validator = v
}

func ValidateStruct(obj interface{}) error {
	return binding.Validator.ValidateStruct(obj)
}

func Var(field interface{}, tag string) error {
	if err := validate.Var(field, tag); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			return errors.New(e.Translate(trans))
		}
	}
	return nil
}

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) != reflect.Struct {
		return nil
	}
	v.lazyinit()
	if err := v.validate.Struct(obj); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			return errors.New(e.Translate(trans))
		}
	}
	return nil
}

func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		zh_translations.RegisterDefaultTranslations(v.validate, trans)
		v.validate.SetTagName("binding")

		//domain
		v.validate.RegisterValidation("domain", domain)
		v.validate.RegisterTranslation("domain", trans, func(ut ut.Translator) error {
			return ut.Add("domain", "{0} 不是一个域名格式", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("domain", fe.Field())
			return t
		})
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func domain(f validator.FieldLevel) bool {
	val := f.Field().String()
	if !strings.Contains(val, ".") {
		return false
	}
	return validate.Var(val, "hostname_rfc1123") == nil
}
