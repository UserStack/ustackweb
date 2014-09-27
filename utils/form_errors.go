package utils

import (
  "github.com/astaxie/beego/validation"
)

func HasFormError(errors []*validation.ValidationError, errorKey string) string {
  for _, error := range errors {
    if error.Key == errorKey {
      return "has-error"
    }
  }
  return ""
}
