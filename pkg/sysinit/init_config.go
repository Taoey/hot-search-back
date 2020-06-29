package sysinit

import (
	"fmt"
	. "github.com/Taoey/hot-search-back/pkg"
	"github.com/olebedev/config"
)

func InitConf() {
	configPath := "./configs/application.yml"
	var err error
	GCF, err = config.ParseYamlFile(configPath)
	if err != nil {
		fmt.Println(err)
	}
}
