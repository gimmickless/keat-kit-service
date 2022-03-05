package config

import (
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

func BindEnvs(v *viper.Viper, ci interface{}) {
	iv := reflect.ValueOf(ci)
	it := reflect.TypeOf(ci)

	for i := 0; i < it.NumField(); i++ {
		fv := iv.Field(i)
		ft := it.Field(i)
		name := strings.ToLower(ft.Name)

		if tag, ok := ft.Tag.Lookup("mapstructure"); ok {
			name = tag
		}

		if k := fv.Kind(); k == reflect.Struct {
			BindEnvs(v, fv.Interface())
		} else {
			_ = v.BindEnv(name)
		}
	}
}
