package conf

import (
	"fmt"
	"testing"
)

func TestReadConfig(t *testing.T) {
	if ReadConfig(".") != nil {
		fmt.Println("fail")
	}
	fmt.Println("success,config:", Config)
}
