/**
 * @Time : 2019/7/16 11:59 
 * @Author : Archmage
 * @File : Trade
 * @Intro:
**/
package api

import (
	"strconv"
	"bitz-contract-sdk/src/config"
	"time"
	"bitz-contract-sdk/src/service"
	"bitz-contract-sdk/src/dao"
	"encoding/json"
	"fmt"
)

type BitzContractApi struct {
	ApiKey string
	SecretKey string
	BaseUrl string
	TradeUrl string
	MarketUrl string
}

func NewBitzApi(apiKey string, secretKey string) BitzContractApi{
	return BitzContractApi{
		ApiKey: apiKey,
		SecretKey:secretKey,
		BaseUrl:config.BaseUrl,
		TradeUrl:config.TradeUrl,
		MarketUrl:config.MarketUrl,
	}
}

func (api BitzContractApi) AddContractTrade(contractId int, price float64, amount int, leverage int, direction int, t string, isCross int) (int, int64) {
	//apiKey, timeStamp, nonce, sign
	param := map[string] string{
		"contractId" : strconv.Itoa(contractId),
		"price" : strconv.FormatFloat(price,'f',-1, 64),
		"amount" : strconv.Itoa(amount),
		"leverage" : strconv.Itoa(leverage),
		"direction" : strconv.Itoa(direction),
		"type" : t,
		"isCross" : strconv.Itoa(isCross),
	}

	param["apiKey"] = api.ApiKey

	param["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10)

	param["nonce"] = service.GenValidateCode(6)

	param["sign"] = service.BitzSign(param, api.SecretKey)

	url := api.TradeUrl + "addContractTrade"

	res := service.HttpPostRequest(url, param)

	header := dao.CommonRes{}

	type D struct{
		OrderId int64 `json:orderId`
	}

	data := D{}
	rs := struct {
		* dao.CommonRes
		Data * D `json:data`
	}{ &header, &data}

	json.Unmarshal([]byte(res), &rs)

	if rs.Status == 200 {
		return rs.Status, rs.Data.OrderId
	} else {
		return rs.Status, -1
	}
}

func (api BitzContractApi) CancelContractTrade(entrustSheetId int) (int, int64) {
	param := make(map[string]string)
	param["apiKey"] = api.ApiKey
	param["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	param["nonce"] = service.GenValidateCode(6)
	param["entrustSheetId"] = strconv.Itoa(entrustSheetId)
	param["sign"] = service.BitzSign(param, api.SecretKey)

	url := api.TradeUrl + "cancelContractTrade"

	res := service.HttpPostRequest(url, param)

	header := dao.CommonRes{}

	type D struct{
		OrderId int64 `json:orderId`
	}

	data := D{}
	rs := struct {
		* dao.CommonRes
		Data * D `json:data`
	}{ &header, &data}

	json.Unmarshal([]byte(res), &rs)

	if rs.Status == 200 {
		return rs.Status, rs.Data.OrderId
	} else {
		return rs.Status, -1
	}
}

func (api BitzContractApi) GetContractAccountInfo() (int, *dao.MxAccountInfo) {
	param := make(map[string] string)
	param["apiKey"] = api.ApiKey
	param["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	param["nonce"] = service.GenValidateCode(6)
	param["sign"] = service.BitzSign(param, api.SecretKey)

	url := api.TradeUrl + "getContractAccountInfo"

	res := service.HttpPostRequest(url, param)

	rs := struct {
		* dao.CommonRes
		Data *dao.MxAccountInfo `json:"data"`
	}{
		&dao.CommonRes{},
		&dao.MxAccountInfo{},
	}

	json.Unmarshal([]byte(res),&rs)

	if rs.Status == 200 {
		return rs.Status, rs.Data
	} else {
		return rs.Status, nil
	}
}

func (api BitzContractApi) GetContractActivePositions(contractId int) (int, []dao.MxPositionInfo){
	param := map[string]string{
		"contractId" : strconv.Itoa(contractId),
	}
	param["apiKey"] = api.ApiKey
	param["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	param["nonce"] = service.GenValidateCode(6)
	param["sign"] = service.BitzSign(param, api.SecretKey)

	url := api.TradeUrl + "getContractActivePositions"

	res := service.HttpPostRequest(url, param)

	header := dao.CommonRes{}

	rs := struct {
		* dao.CommonRes
		Data []dao.MxPositionInfo `json:"data"`
	}{
		&header,
		[]dao.MxPositionInfo{},
	}

	json.Unmarshal([]byte(res), &rs)

	if rs.Status == 200 {
		return rs.Status, rs.Data
	} else {
		return rs.Status, nil
	}
}

func (api BitzContractApi) GetContractMyPositions(contractId int, page int, pageSize int) (int, []dao.MxPositionInfo) {
	param := map[string]string{
		"contractId" : strconv.Itoa(contractId),
		"page" : strconv.Itoa(page),
		"pageSize" : strconv.Itoa(pageSize),
	}

	param["apiKey"] = api.ApiKey
	param["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	param["nonce"] = service.GenValidateCode(6)
	param["sign"] = service.BitzSign(param, api.SecretKey)

	url := api.TradeUrl + "getContractMyPositions"

	res := service.HttpPostRequest(url, param)

	header := dao.CommonRes{}

	rs := struct {
		* dao.CommonRes
		Data []dao.MxPositionInfo `json:"data"`
	}{
		&header,
		[]dao.MxPositionInfo{},
	}

	json.Unmarshal([]byte(res), &rs)

	if rs.Status == 200 {
		return rs.Status, rs.Data
	} else {
		return rs.Status, nil
	}
}

func (api BitzContractApi) GetContractOrderResult(entrustSheetIds string) (int, []dao.MxOrder){
	param := map[string]string {
		"entrustSheetIds" : entrustSheetIds,
	}
	param["apiKey"] = api.ApiKey
	param["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	param["nonce"] = service.GenValidateCode(6)
	param["sign"] = service.BitzSign(param, api.SecretKey)

	url := api.TradeUrl + "getContractOrderResult"

	res := service.HttpPostRequest(url, param)

	rs := struct {
		* dao.CommonRes
		Data []dao.MxOrder
	}{
		& dao.CommonRes{},
		[]dao.MxOrder{},
	}

	json.Unmarshal([]byte(res), &rs)

	if rs.Status == 200 {
		return rs.Status, rs.Data
	} else {
		return rs.Status, nil
	}
}

func (api BitzContractApi) GetContractOrder(contractId int) (int, []dao.MxOrder) {
	param := map[string]string{
		"contractId": strconv.Itoa(contractId),
	}
	param["apiKey"] = api.ApiKey
	param["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	param["nonce"] = service.GenValidateCode(6)
	param["sign"] = service.BitzSign(param, api.SecretKey)

	url := api.TradeUrl + "getContractOrder"

	res := service.HttpPostRequest(url, param)

	header := dao.CommonRes{}

	orders := []dao.MxOrder{}

	rs := struct {
		*dao.CommonRes
		Data []dao.MxOrder `json:data`
	}{
		&header,
		orders,
	}

	json.Unmarshal([]byte(res), &rs)

	if rs.Status == 200 {
		return rs.Status, rs.Data
	} else {
		return rs.Status, nil
	}
}

func (api BitzContractApi) GetContractTradeResult(entrustSheetId int) (int, []dao.MxTradeDetail){
	param := map[string]string {
		"entrustSheetId" : strconv.Itoa(entrustSheetId),
	}
	param["apiKey"] = api.ApiKey
	param["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	param["nonce"] = service.GenValidateCode(6)
	param["sign"] = service.BitzSign(param, api.SecretKey)

	url := api.TradeUrl + "getContractTradeResult"

	res := service.HttpPostRequest(url, param)

	rs := struct {
		 * dao.CommonRes
		 Data []dao.MxTradeDetail
	}{
		& dao.CommonRes{},
		[] dao.MxTradeDetail{},
	}

	json.Unmarshal([]byte(res), &rs)

	if rs.Status == 200 {
		return rs.Status, rs.Data
	} else {
		return rs.Status, nil
	}
}

func (api BitzContractApi) GetContractMyHistoryTrade(contractId int, page int, pageSize int) (int, []dao.MxOrder) {
	param := map[string]string {
		"contractId" : strconv.Itoa(contractId),
		"page" : strconv.Itoa(page),
		"pageSize" : strconv.Itoa(pageSize),
	}
	param["apiKey"] = api.ApiKey
	param["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	param["nonce"] = service.GenValidateCode(6)
	param["sign"] = service.BitzSign(param, api.SecretKey)

	url := api.TradeUrl + "getContractMyHistoryTrade"

	res := service.HttpPostRequest(url, param)

	header := dao.CommonRes{}

	orders := []dao.MxOrder{}

	rs := struct {
		*dao.CommonRes
		Data []dao.MxOrder `json:data`
	}{
		&header,
		orders,
	}

	json.Unmarshal([]byte(res), &rs)

	if rs.Status == 200 {
		return rs.Status, rs.Data
	} else {
		return rs.Status, nil
	}
}

func (api BitzContractApi) GetContractMyTrades(contractId int, page int, pageSize int) (int, []dao.MxTradeInfo) {
	param := map[string]string {
		"contractId" : strconv.Itoa(contractId),
		"page" : strconv.Itoa(page),
		"pageSize" : strconv.Itoa(pageSize),
	}
	param["apiKey"] = api.ApiKey
	param["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	param["nonce"] = service.GenValidateCode(6)
	param["sign"] = service.BitzSign(param, api.SecretKey)

	url := api.TradeUrl + "getContractMyTrades"

	res := service.HttpPostRequest(url, param)

	rs := struct {
		* dao.CommonRes
		Data []dao.MxTradeInfo
	}{
		& dao.CommonRes{},
		[] dao.MxTradeInfo{},
	}

	json.Unmarshal([]byte(res), &rs)

	if rs.Status == 200 {
		return rs.Status, rs.Data
	} else {
		return rs.Status, nil
	}
}

//-----market apis-----
func (api BitzContractApi) GetContractCoin(contractId int) (int, []dao.MxContractInfo){
	url := config.MarketUrl + "getContractCoin"
	res := service.HttpGetRequest(url,nil)

	header := dao.CommonRes{}
	cInfos := []dao.MxContractInfo{}
	rs := struct {
		*dao.CommonRes
		Data * [] dao.MxContractInfo `json:Data`
	}{
		&header,
		&cInfos,
	}

	json.Unmarshal([]byte(res), &rs)

	if rs.Status == 200 {
		return rs.Status, cInfos
	} else{
		return rs.Status, nil
	}
}

func (api BitzContractApi) GetContractKline(contractId int, t string, size int) (int, dao.KData){
	if size <= 0 {
		size = 300
	}

	param := map[string]string{
		"contractId" : strconv.Itoa(contractId),
		"type" : t,
		"size" : strconv.Itoa(size),
	}

	url := config.MarketUrl + "getContractKline"
	res := service.HttpGetRequest(url, param)

	header := dao.CommonRes{}
	lists := dao.KDataLists{}

	rs := struct {
		* dao.CommonRes
		Data dao.KDataLists `json:data`
	}{
		&header,
		lists,
	}

	json.Unmarshal([]byte(res), &rs)

	if rs.Status == 200 {
		return rs.Status, rs.Data.Lists
	}else{
		return rs.Status, nil
	}
}

func (api BitzContractApi) GetContractOrderBook(contractId int, depth int) (int, []dao.MxOrder, []dao.MxOrder) {
	if depth <= 0 {
		depth = 10
	}

	param := map[string]string{
		"contractId" : strconv.Itoa(contractId),
		"depth" : strconv.Itoa(depth),
	}

	url := config.MarketUrl + "getContractOrderBook"

	res := service.HttpGetRequest(url, param)

	header := dao.CommonRes{}
	orders := dao.MxOrderBook{}

	rs := struct {
		* dao.CommonRes
		Data * dao.MxOrderBook `json:data`
	}{
		&header,
		&orders,
	}

	json.Unmarshal([]byte(res), &rs)

	if rs.Status == 200 {
		return rs.Status, rs.Data.Bids, rs.Data.Asks
	}else{
		return rs.Status, nil, nil
	}
}

func (api BitzContractApi) GetContractTradesHistory(contractId int, pageSize int) (int, []dao.MxTradeInfo) {
	if pageSize <= 0 {
		pageSize = 10
	}

	param := map[string] string{
		"contractId" : strconv.Itoa(contractId),
		"pageSize" : strconv.Itoa(pageSize),
	}

	url := config.MarketUrl + "getContractTradesHistory"

	res := service.HttpGetRequest(url, param)

	header := dao.CommonRes{}

	type D struct {
		Lists []dao.MxTradeInfo `json:lists`
	}

	d := D{
	}

	rs := struct {
		* dao.CommonRes
		Data * D `json:data`
	}{
		&header,
		&d,
	}

	json.Unmarshal([]byte(res), &rs)

	if rs.Status == 200{
		return rs.Status, rs.Data.Lists
	} else {
		return rs.Status, nil
	}
}

func (api BitzContractApi) GetContractTickers(contractId int) (int, []dao.MxTicker){
	param := map[string] string{
		"contractId" : strconv.Itoa(contractId),
	}

	url := config.MarketUrl + "getContractTickers"

	res := service.HttpGetRequest(url, param)
	fmt.Println(res)

	header := dao.CommonRes{}


	rs := struct {
		* dao.CommonRes
		Data []dao.MxTicker `json:data`
	}{
		&header,
		[]dao.MxTicker{},
	}

	json.Unmarshal([]byte(res), &rs)

	if rs.Status == 200 {
		return rs.Status, rs.Data
	}else{
		return rs.Status, nil
	}
}