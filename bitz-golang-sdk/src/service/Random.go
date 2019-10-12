/**
 * @Time : 2019/7/16 14:03 
 * @Author : Archmage
 * @File : Random
 * @Intro:
**/
package service

import (
	"time"
	"strings"
	"fmt"
	"math/rand"
)

func GenValidateCode(width int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	return sb.String()
}
