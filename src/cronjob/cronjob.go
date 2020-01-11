package cronjob

import (
	"fmt"

	"github.com/user/go_cronjob/src/config"
)

var jobConfig *config.TomlConfig

// var bgAll [string]

// StartCronjob get config
func StartCronjob() {
	// 取得config
	jobConfig = config.GetConfig()

	// 依照config逐行建立bg監控物件
	for _, element := range jobConfig.JobList {
		fmt.Println("element = ", element)
	}
}
