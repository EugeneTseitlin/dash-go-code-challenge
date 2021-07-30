package validation

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"

	"github.com/EugeneTseitlin/dash-go-code-challenge/internal/app/model"
)

func IsISO8601(fl validator.FieldLevel) bool {
	ISO8601DateRegexString := `^(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(.[0-9]+)?(Z)?$`
	ISO8601DateRegex := regexp.MustCompile(ISO8601DateRegexString)

	return ISO8601DateRegex.MatchString(fl.Field().String())
}

func IsAplhaNumDotUnderscore(fl validator.FieldLevel) bool {
	var usernamePattern = regexp.MustCompile(`^[a-zA-Z0-9._]+$`)
	return usernamePattern.MatchString(fl.Field().String())
}

func CreateValidator() *validator.Validate {
	v := validator.New()
	_ = v.RegisterValidation("alphanumdotunderscore", IsAplhaNumDotUnderscore)
	_ = v.RegisterValidation("iso8601", IsISO8601)
	return v
}

func ValidateData(v *validator.Validate, items []map[string]interface{}) error {

	for _, item := range items {
		itemType := item["type"].(string)
		switch itemType {
		case "user":
			var user model.User
			mapstructure.Decode(item, &user)

			err := v.Struct(user)
			if err != nil {
				return err
			}
		case "merchant":
			var merchant model.Merchant
			mapstructure.Decode(item, &merchant)

			err := v.Struct(merchant)
			if err != nil {
				return err
			}
		case "payment":
			var payment model.Payment
			mapstructure.Decode(item, &payment)

			err := v.Struct(payment)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported data type")
		}
	}

	return nil
}
