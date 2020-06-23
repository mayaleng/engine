package web

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var validate = validator.New()
var enLocale = en.New()
var translator = ut.New(enLocale, enLocale)

func init() {
	lang, _ := translator.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, lang)

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// Decode tries to parse the input body and validate its content
func Decode(r *http.Request, val interface{}) []Error {
	var decoder = json.NewDecoder(r.Body)
	lang, _ := translator.GetTranslator("en")

	error := decoder.Decode(val)

	if error != nil {
		error := Error{
			Status: 400,
			Detail: "Bad request",
		}
		return []Error{error}
	}

	error = validate.Struct(val)

	if error == nil {
		return []Error{}
	}

	validationErrors, _ := error.(validator.ValidationErrors)

	if error != nil {
		var errors = make([]Error, 0)

		for _, validationError := range validationErrors {
			webError := Error{
				Detail: validationError.Translate(lang),
				Status: 422,
				Source: ErrorSource{
					Pointer: validationError.Field(),
				},
			}

			errors = append(errors, webError)
		}

		return errors
	}

	return []Error{}
}
