package validators

import (
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v8"
	"gorm.io/gorm"
)

var CustomValidator *cValidator

type cValidator struct {
	db *gorm.DB
}

func NewValidator(db *gorm.DB) {
	CustomValidator = &cValidator{
		db: db,
	}
}

// check if value of request is unique in database
// tag format : unique=tablename.columnname
func (cv *cValidator) Unique() func(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	return func(
		v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
		field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
	) bool {
		arr := strings.Split(param, ".")
		rows, err := cv.db.Table(arr[0]).Select("*").Where(arr[1]+" = ?", field.Interface().(string)).Rows()
		if err != nil || rows.Next() {
			return false
		}
		return true
	}
}
