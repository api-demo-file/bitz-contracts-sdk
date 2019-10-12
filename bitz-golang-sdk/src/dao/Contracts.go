/**
 * @Time : 2019/7/17 17:47 
 * @Author : Archmage
 * @File : MxTicker
 * @Intro:
**/
package dao

type MxContractInfo struct {
	AllowCross        string `json:"allowCross"`
	AllowLeverages    string `json:"allowLeverages"`
	AnchorDec         string `json:"anchorDec"`
	ContractAnchor    string `json:"contractAnchor"`
	ContractID        int `json:"contractId,string"`
	ContractValue     float64 `json:"contractValue,string"`
	Expiry            string `json:"expiry"`
	Isreverse         int `json:"isreverse,string"`
	MaintanceMargin   float64 `json:"maintanceMargin,string"`
	MakerFee          float64 `json:"makerFee,string"`
	MaxAmount         int64 `json:"maxAmount,string"`
	MaxLeverage       int `json:"maxLeverage,string"`
	MaxOrderNum       int64 `json:"maxOrderNum,string"`
	MaxPositionAmount int64 `json:"maxPositionAmount,string"`
	MinAmount         int64 `json:"minAmount,string"`
	Pair              string `json:"pair"`
	PriceDec          int `json:"priceDec,string"`
	QuoteAnchor       string `json:"quoteAnchor"`
	SettleAnchor      string `json:"settleAnchor"`
	SettleFee         float64 `json:"settleFee,string"`
	Status            int `json:"status,string"`
	Symbol            string `json:"symbol"`
	TakerFee          float64 `json:"takerFee,string"`
	Type              string `json:"type"`
}

type KData []interface{}

type KDataLists struct{
	Lists KData `json:"lists"`
}

type MxTradeInfo struct {
	Time int64 `json:"time"`
	Price float64 `json:"price,string"`
	Num int64 `json:"num,string"`
	TradeType string `json:"type"`
	TradeId  string `json:"tradeid"`
	ContractId int `json:"contractId,string"`
	Pair string `json:"pair"`
	TradeFee float64 `json:"tradeFee,string"`
	Leverage int `json:"leverage,string"`
	IsCross int `json:"isCross,string"`
}

type MxTicker struct{
	Amount          float64 `json:"amount,string"`
	Change24h       float64 `json:"change24h,string"`
	ContractID      int `json:"contractId,string"`
	FairPrice       float64 `json:"fairPrice,string"`
	IndexPrice      float64 `json:"indexPrice,string"`
	Latest          float64 `json:"latest,string"`
	Max             float64 `json:"max,string"`
	Min             float64 `json:"min,string"`
	NextFundingRate float64 `json:"nextFundingRate,string"`
	OpenInterest    float64 `json:"openInterest,string"`
	Pair            string `json:"pair"`
	Volumn          float64 `json:"volumn,string"`
}

type MxOrder struct {
	Price float64 `json:"price,string"`
	Amount int64 `json:"amount,string"`
	OrderId int64 `json:"orderId,string"`
	ContractId int `json:"contractId,string"`
	OrderType string `json:"type"`
	Leverage int `json:"leverage,string"`
	Direction int `json:"direction,string"`
	OrderStatus int `json:"orderStatus,string"`
	IsCross int `json:"isCross,string"`
	Available int `json:"available,string"`
	Time int64 `json:"time,int"`
	Pair string `json:"pair"`
}

type MxOrderBook struct{
	Bids []MxOrder `json:"bids"`
	Asks []MxOrder `json:"asks"`
}

type MxPositionInfo struct{
	Amount           int64 `json:"amount,string"`
	ContractID       int `json:"contractId,string"`
	Direction        int `json:"direction,string"`
	IsCross          int `json:"isCross,string"`
	Leverage         float64 `json:"leverage,string"`
	LiquidationPrice float64 `json:"liquidationPrice,string"`
	Margin           float64 `json:"margin,string"`
	Pair             string `json:"pair"`
	PositionID       int64 `json:"positionId,string"`
	Price            float64 `json:"price,string"`
	RlzPnl           float64 `json:"rlzPnl,string"`
	UnrlzPnl         float64 `json:"unrlzPnl,string"`
	StartTime		 int64 `json:"startTime,int"`
	EndTime  int64 `json:"endTime,int"`
}

type MxAccountInfo struct {
	Time int64 `json:"time,int"`
	EstimateBTC float64 `json:"estimate_BTC,string"`
	EstimateUSD float64 `json:"estimate_USD,string"`
	EstimateCNY float64 `json:"estimate_CNY,string"`
	Balances []MxCoin `json:"balances"`
}

type MxCoin struct{
	Balance        float64 `json:"balance,string"`
	Coin           string `json:"coin"`
	EstimateBTC    float64 `json:"estimate_BTC,string"`
	EstimateCNY    float64 `json:"estimate_CNY,string"`
	EstimateUSD    float64 `json:"estimate_USD,string"`
	OrderMargin    float64 `json:"orderMargin,string"`
	PositionMargin float64 `json:"positionMargin,string"`
	Total          float64 `json:"total,string"`
	UnrlzPnl       float64 `json:"unrlzPnl,string"`
}

type MxTradeDetail struct{
	ContractID int `json:"contractId,string"`
	IsCross    int `json:"isCross,string"`
	Leverage   int `json:"leverage,string"`
	OrderID    int64 `json:"orderId,string"`
	OrderPrice float64 `json:"orderPrice,string"`
	Pair       string `json:"pair"`
	Price      float64 `json:"price,string"`
	Time       int    `json:"time"`
	TradeFee   float64 `json:"tradeFee,string"`
	TradeID    string `json:"tradeId"`
	TradeNum   int64 `json:"tradeNum,string"`
	TradeType  string `json:"type"`
}
