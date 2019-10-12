/**
 * @Time : 2019/7/17 17:54 
 * @Author : Archmage
 * @File : CommonRes
 * @Intro:
**/
package dao

type CommonRes struct{
	Status int `json:status`
	Msg string `json:msg`
	Time int64 `json:time`
	Microtime string `json:microtime`
	Source string `json:source`
}
