package commonfunctions

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func TranslateError(err error, validate *validator.Validate) map[string]interface{} {

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	if err == nil {
		return nil
	}
	errorResponse := map[string]interface{}{}
	validatorErrs := err.(validator.ValidationErrors)
	for _, validationErr := range validatorErrs {
		fieldName := validationErr.Field()
		fieldJsonName := fieldName
		errorResponse[fieldJsonName] = fmt.Sprint(validationErr.Translate(trans))
	}
	return errorResponse
}
