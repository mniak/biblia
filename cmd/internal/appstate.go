package internal

import (
	"log"
	"os"

	"github.com/mniak/biblia/pkg/biblehub"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var AppState AppStateStruct

type AppStateStruct struct {
	LastInterlinearChapter *biblehub.ChapterID `mapstructure:"last_interlinear_chapter" yaml:"last_interlinear_chapter"`
}

func (state AppStateStruct) Save() error {

	yamlbytes, err := yaml.Marshal(&state)
	if err != nil {
		return err
	}
	err = os.WriteFile(".biblia.yaml", yamlbytes, 0655)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	viper.SetConfigName(".biblia")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("failed to read the config file: %s", err)
	}
	err = viper.Unmarshal(&AppState)
	if err != nil {
		log.Printf("failed to parse the config file: %s", err)
	}

}
