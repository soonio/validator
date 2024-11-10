package validator

import "github.com/go-playground/validator/v10"

// 配置项

// WithDefaultLanguage 设置默认配置
// lang 当前可选值 zh | en
func WithDefaultLanguage(lang string) func(v *Validator) *Validator {
	return func(v *Validator) *Validator {
		v.defaultLang = lang
		return v
	}
}

// WithValidator 外部注入validator(兼容gin框架默认初始化了validator)
func WithValidator(validator *validator.Validate) func(v *Validator) *Validator {
	return func(v *Validator) *Validator {
		v.validator = validator
		return v
	}
}
