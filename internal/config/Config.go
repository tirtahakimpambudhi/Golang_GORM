package config

import (
	"github.com/spf13/viper"
	"go_gorm/internal/domain/model"
	"go_gorm/pkg/common"
	"log"
	"os"
	"path/filepath"
)

var ServerConf *model.Server
var DatabaseConf *model.Database

func init() {
	workdir, _ := os.Getwd()
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join(workdir, "config"))    //For Main workdir
	viper.AddConfigPath(filepath.Join(workdir, "../config")) //For Test Workdir

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("File Configuration Not Found \n")
		} else {
			log.Fatalf("Another Error : %v\n", err.Error())
		}
	}

	ServerConf = &model.Server{}
	common.GetValueConfigByTag("config", ServerConf)
	DatabaseConf = &model.Database{}
	common.GetValueConfigByTag("config", DatabaseConf)
}
