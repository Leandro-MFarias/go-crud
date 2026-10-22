package validation

import (
	"encoding/json"
	"errors"

	"github.com/Leandro-MFarias/go-crud/src/config/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {

		en := en.New()
		unt := ut.New(en, en)

		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	// Unmarshal não tem como fazer bind nele então vamos só retornar ele. Acontece quando dá erro de tipagem em algum campo
	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		// Popular as causas
		errorsCauses := []rest_err.Causes{}

		// Usar um for ignorar o primeiro parametro que é um index. Usar o range pq pode ter dado erro em 20 campos.
		for _, e := range validation_err.(validator.ValidationErrors) {
			// Cria um objeto de cause do tipo rest_err.Causes que precisa enviar Message e Field.
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			// A cada iteração ele adiciona a cause no errorsCauses.
			errorsCauses = append(errorsCauses, cause)
		}
		// Retorna no Metodo que criamos de Validation erros uma mensagem e envia o errorsCauses.
		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Errors trying to convert fields")
	}
}
