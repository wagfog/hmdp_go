package test

import (
	"fmt"
	"testing"

	"github.com/wagfog/hmdp_go/utils"
)

func TestJWT(t *testing.T) {
	stoken, err := utils.GenerateToken("15359941669")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(stoken)

	// c.defa

}
