package test

import (
	"fmt"
	"go_gorm/internal/config"
	"reflect"
	"testing"
)

func TestConfigServer(t *testing.T) {
	serverValue := reflect.ValueOf(config.ServerConf).Elem()

	// Mendapatkan tipe refleksi dari struct.
	serverType := serverValue.Type()

	// Iterasi melalui setiap bidang dari struct.
	for i := 0; i < serverType.NumField(); i++ {
		// Mendapatkan informasi tentang bidang (field) saat ini.
		field := serverType.Field(i)

		// Mendapatkan nilai dari bidang menggunakan refleksi.
		fieldValue := serverValue.Field(i)

		// Cetak informasi bidang.
		fmt.Printf("Field Name: %s\n", field.Name)
		fmt.Printf("Field Type: %s\n", fieldValue.Type())
		fmt.Printf("Field Value: %v\n", fieldValue.Interface())
		fmt.Println("----------------------")
	}
}
func TestConfigDatabase(t *testing.T) {
	databaseValue := reflect.ValueOf(config.DatabaseConf).Elem()

	// Mendapatkan tipe refleksi dari struct.
	databaseType := databaseValue.Type()

	// Iterasi melalui setiap bidang dari struct.
	for i := 0; i < databaseType.NumField(); i++ {
		// Mendapatkan informasi tentang bidang (field) saat ini.
		field := databaseType.Field(i)

		// Mendapatkan nilai dari bidang menggunakan refleksi.
		fieldValue := databaseValue.Field(i)

		// Cetak informasi bidang.
		fmt.Printf("Field Name: %s\n", field.Name)
		fmt.Printf("Field Type: %s\n", fieldValue.Type())
		fmt.Printf("Field Value: %v\n", fieldValue.Interface())
		fmt.Println("----------------------")
	}
}
