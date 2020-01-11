package main

import (
	"fmt"

	"github.com/user/go_cronjob/src/cronjob"
)

func main() {
	fmt.Println("start process")
	cronjob.StartCronjob()
	fmt.Println("end process")
}
