package validations

import (
	"context"
	"regexp"
	"rowing-market-api/api/models"
	"rowing-market-api/pkg/constants"
	"rowing-market-api/pkg/translator"
	"strings"
)

type CreatePostError struct {
	HasError    bool   `json:"-"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Price       uint64 `json:"price,omitempty"`
}

func CreatePostValidation(ctx context.Context, c models.Post, lang string) CreatePostError {
	err := CreatePostError{}

	if len(strings.Trim(c.Title, " ")) == 0 {
		err.HasError = true
		err.Title = translator.Trans("fieldIsMissing", lang, map[string]interface{}{"Field": "title"})
	} else {
		regxExclusion := regexp.MustCompile(constants.RegexExclusion)
		if !regxExclusion.MatchString(c.Title) {
			err.HasError = true
			err.Title = translator.Trans("fieldIsNotValid", lang, map[string]interface{}{"Field": "post_code"})
		}
	}

	if len(strings.Trim(c.Description, " ")) == 0 {
		err.HasError = true
		err.Description = translator.Trans("fieldIsMissing", lang, map[string]interface{}{"Field": "description"})
	} else {
		regxExclusion := regexp.MustCompile(constants.RegexExclusion)
		if !regxExclusion.MatchString(c.Description) {
			err.HasError = true
			err.Title = translator.Trans("fieldIsNotValid", lang, map[string]interface{}{"Field": "description"})
		}
	}
	if c.Price == 0 {
		err.HasError = true
		err.Description = translator.Trans("fieldIsMissing", lang, map[string]interface{}{"Field": "price"})
	}

	return err
}
