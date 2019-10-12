/**
 * @Time : 2019/7/16 10:43 
 * @Author : Archmage
 * @File : main
 * @Intro:
**/
package main
import (
	"bitz-contract-sdk/src/model/api"
	"fmt"
)

func main() {

	var ApiKey string = ""
	var SecretKey string = ""
	api := api.NewBitzApi(ApiKey, SecretKey)
	fmt.Print(api)

	/*--
	param := make(map[string]string)
	param["timeStamp"] = "1510235730"
	param["nonce"] = "309127"
	param["contractId"] = "101"
	param["price"] = "1"
	param["amount"] = "1"
	param["leverage"] = "10"
	param["direction"] = "1"
	param["type"] = "1"
	param["isCross"] = "1"

	sign := service.BitzSign(param)

	str := "apiKey=376892265asdad5d12726d8bbfbd8912b3&timeStamp=1510235730&nonce=309127&contractId=101&price=1&amount=1&leverage=10&direction=1&type=limit&isCross=1aQmE8U7bxj16KdJcSd3yX8F8Sakd8aO6LopnHXh27d4kWyb28PxcaTvGrajLDvAw"
	fmt.Println(str)

	h := md5.New()
	h.Write([]byte(str))
	sign2 := hex.EncodeToString(h.Sum(nil))


	fmt.Println(sign)
	fmt.Println(sign2)
	//----*/
	//api.GetContractCoin(101)
	//api.GetContractKline(101, "1d", 300)
	//api.GetContractTradesHistory(101, 10)
	//api.GetContractTickers(101)
	//_, bids, asks := api.GetContractOrderBook(101, 10)
	//fmt.Println(bids)
	//fmt.Println(asks)

	//1. 下单
	//contractId int, price float64, amount int, leverage int, direction int, t string, isCross int

	/*---
	cid := 101
	var price float64 = 9000.0
	amount := 1
	l := 15
	d := 1
	t := "limit"
	isCross := -1

	api.AddContractTrade(cid, price, amount, l, d, t, isCross)

	//---*/

	//cancel order
	//orderId 6050879,6052108, 6052235
	//2. 取消委托 6183813
	//api.CancelContractTrade(6183813)
	//3. 未平仓位
	//api.GetContractActivePositions(101)
	//4. 我的账户权益
	//_, accountInfo := api.GetContractAccountInfo()
	//fmt.Print(accountInfo.Balances)
	//5. 已平仓位
	//api.GetContractMyPositions(101, 1, 10)
	//6. 我的活动委托
	//api.GetContractOrder(101)
	//7. 单个或多个委托明细
	//api.GetContractOrderResult("6064638,6064632")
	//8. 某个委托的成交明细
	//api.GetContractTradeResult(6063459)
	//9. 我的委托历史
	//api.GetContractMyHistoryTrade(101, 1, 10)
	//10. 成交历史
	//api.GetContractMyTrades(101, 1, 10)
}