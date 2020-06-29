package sysinit

import (
	"fmt"
	"github.com/olebedev/config"
)

var GCF *config.Config //global config

func InitConf() {
	configPath := "./configs/application.yml"
	//configPath := `G:\坚果云同步\github\1_projects\hot-search-back\configs\application.yml`
	var err error
	GCF, err = config.ParseYamlFile(configPath)
	if err != nil {
		fmt.Println(err)
	}
}
