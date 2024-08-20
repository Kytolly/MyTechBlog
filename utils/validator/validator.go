package validator

import (
	"github.com/go-playground/validator/v10"
	unTrans "github.com/go-playground/universal-translator"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/go-playground/locales/zh_Hans_CN"
	//"mytechblog/model"
	"log/slog"
	msg "mytechblog/utils/errormsg"
	"reflect"
)

func Validate(data interface{})(string, int){
	validate := validator.New()
	uni := unTrans.New(zh_Hans_CN.New())
	trans,_ := uni.GetTranslator("zh_Hans_CN")

	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		slog.Error(err.Error())
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField)string{
		return field.Tag.Get("label")
	})
	err = validate.Struct(data)
	if err != nil{
		for _, v := range err.(validator.ValidationErrors){
			return v.Translate(trans), msg.ERROR
		}
	}
	return "", msg.SUCCESS
}