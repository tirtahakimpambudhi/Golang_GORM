package common

import (
	"fmt"
	"github.com/spf13/viper"
	"reflect"
)

func GetValueConfigByTag(keyTags string, conf interface{}) {
	confValue := reflect.ValueOf(conf).Elem()
	confType := confValue.Type()

	for i := 0; i < confType.NumField(); i++ {
		field := confType.Field(i)
		tagValue := field.Tag.Get(keyTags)

		// Mengakses nilai konfigurasi dari Viper berdasarkan tag.
		value := viper.Get(tagValue)

		// Mengeset nilai bidang di struct menggunakan reflect.
		if confValue.Field(i).CanSet() {
			fieldType := confValue.Field(i).Type()

			if reflect.TypeOf(value) == fieldType {
				confValue.Field(i).Set(reflect.ValueOf(value))
			} else {
				// Penanganan jika tipe data tidak cocok.
				fmt.Printf("Type mismatch for field %s: expected %v, got %v\n", field.Name, fieldType, reflect.TypeOf(value))
			}
		}
	}
}
