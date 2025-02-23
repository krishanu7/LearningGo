package modules

import (
	"fmt"
	"time"
)

func GetTime() {
	fmt.Println("The time is", time.Now().Format(time.Kitchen));
}