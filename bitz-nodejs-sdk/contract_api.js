var CryptoJS = require('crypto-js');
var Promise = require('bluebird');
const request = require('request');
const DEFAULT_HEADERS = {
    // "Content-Type": "application/json",
    "Content-Type": "application/x-www-form-urlencoded",
    "User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36"
}
//
var WEB_BASE = ''; //
// 
var config = {
    'apiKey': '',
    'secret_key': ''
};
//
function log(msg) {
    console.log(msg);
}

function market_api(path, data) {
    var url = `${WEB_BASE}${path}`; 
    var pars = [];
    for (let key in data) {
        let v = data[key];
        pars.push(key + "=" + v);
    }
    var p = pars.join("&");
    url = url + "?" + p;
    return new Promise((resolve, reject) => {
        try{
            var httpOptions = {
                url: url,
                method: 'get',
                timeout: 3000,
                headers: DEFAULT_HEADERS,
            }
            log('-------------httpOptions-----------');
            log(httpOptions);
            request.get(httpOptions, function(err, res, data) {
                if (err) {
                    reject(err);
                } else {
                    if (res.statusCode == 200) {
                        resolve(data);
                    } else {
                        reject(res.statusCode);
                    }
                }
            }).on('error', function(err){
                console.log('http get err:'+url);
                reject(null);
            });
        }catch(err){
            console.log('http get err:'+url);
            reject(null);    
        }
    });
}

function sign_api(path, data) {
    var url = `${WEB_BASE}${path}`; 
    data = setSign(data);
    return new Promise((resolve, reject) => {
        try{
            var httpOptions = {
                url: url,
                form: data,
                method: 'post',
                timeout: 3000,
                headers: DEFAULT_HEADERS,
            };
            log('-------------httpOptions-----------');
            log(httpOptions);
            request(httpOptions, function(err, res, data) {
                if (err) {
                    reject(err);
                } else {
                    if (res.statusCode == 200) {
                        resolve(data);
                    } else {
                        reject(res.statusCode);
                    }
                }
            }).on('error', function(err){
                console.log('http form_post err:'+err);
                reject(null);
            });
        }catch(err){
            console.log('http form_post err:'+err);
            reject(null);    
        }
    });
}

function setSign(params) {
    var pars = [];
    let keys = Object.keys(params);
    let n = keys.length;
    keys = keys.sort();
    let sign = '';
    for (let i = 0; i < n; i++) {
        if (sign != '') sign = sign + "&";
        sign = sign + keys[i] + "=" + params[keys[i]];
    }
    //
    sign = sign + config.secret_key;
    sign = CryptoJS.MD5(sign).toString().toLowerCase();
    params.sign = sign;
    return params;
}

function getSignBaseParams() {
    let timestamp = Math.round(new Date().getTime() / 1000) + "";
    return {
        "apiKey": config.apiKey,
        "timeStamp": timestamp,
        "nonce": timestamp.substr(-6)
    };
}

var CONTRACT_API = {
    setConfig : function(apiServer,apiKey,secretKey,tradePwd){
        WEB_BASE = apiServer;
        config = {
            'apiKey': apiKey ,
            'secret_key': secretKey ,
            'tradePwd': tradePwd || '',
        };
    },

    /**
     * ------------------------------------------------------------------
     * 合约账户权益
     * ------------------------------------------------------------------
     */
    getContractAccountInfo: function() {
        var data = getSignBaseParams();
        return sign_api("/Contract/getContractAccountInfo", data);
    },
    /**
     * ------------------------------------------------------------------
     * 我的活动委托
     * ------------------------------------------------------------------
     */
    getContractOrder: function(contractId, page, pageSize) {
        var data = getSignBaseParams();
        data.contractId = contractId;
        if(page) data.page = page;
        if(pageSize) data.pageSize = pageSize;
        return sign_api("/Contract/getContractOrder", data);
    },
    /**
     * ------------------------------------------------------------------
     * 获取当前仓位 
     * @param contractId   101
     * ------------------------------------------------------------------
     */
    getContractActivePositions: function(contractId) {
        var data = getSignBaseParams();
        data.contractId = contractId;
        return sign_api("/Contract/getContractActivePositions", data);
    },
    /**
     * ------------------------------------------------------------------
     * 提交委托单 (Place an order)
     * @param contractId      101
     * @param price        float
     * @param amount         float
     * @param leverage         float
     * @param direction         float
     * @param isCross         float
     * @param type          string  "limit"  "market"
     * ------------------------------------------------------------------
     */
    addContractTrade: function(contractId, price, amount, leverage, direction, type, isCross) {
        var data = getSignBaseParams();
        data.contractId = contractId;
        data.price = price;
        data.amount = amount;
        data.leverage = leverage;
        data.direction = direction;
        data.type = type;
        data.isCross = isCross;
        return sign_api("/Contract/addContractTrade", data);
    },
    /**
     * ------------------------------------------------------------------
     * 取消委托单
     * @param entrustSheetId   float
     * ------------------------------------------------------------------
     */
    cancelContractTrade: function(entrustSheetId) {
        var data = getSignBaseParams();
        data.entrustSheetId = entrustSheetId;
        return sign_api("/Contract/cancelContractTrade", data);
    },
    /**
     * ------------------------------------------------------------------
     * 获取我的委托历史
     * @param contractId      101
     * ------------------------------------------------------------------
     */
    getContractMyHistoryTrade: function(contractId, page, pageSize) {
        var data = getSignBaseParams();
        data.contractId = contractId;
        if(page) data.page = page;
        if(pageSize) data.pageSize = pageSize;
        return sign_api("/Contract/getContractMyHistoryTrade", data);
    },
    /**
     * ------------------------------------------------------------------
     * 交易历史
     * @param contractId    101
     * ------------------------------------------------------------------
     */
    getContractMyTrades: function(contractId) {
        var data = getSignBaseParams();
        data.contractId = contractId;
        return sign_api("/Contract/getContractMyTrades", data);
    },
    /**
     * ------------------------------------------------------------------
     * 单个委托成交明细
     * @param entrustSheetId  float  6235
     * ------------------------------------------------------------------
     */
    getContractTradeResult: function(entrustSheetId) {
        var data = getSignBaseParams();
        data.entrustSheetId = entrustSheetId;
        return sign_api("/Contract/getContractTradeResult", data);
    },
    /**
     * ------------------------------------------------------------------
     * 已平仓仓位
     * @param contractId    101
     * ------------------------------------------------------------------
     */
    getContractMyPositions: function(contractId, page, pageSize) {
        var data = getSignBaseParams();
        data.contractId = contractId;
        if(page) data.page = page;
        if(pageSize) data.pageSize = pageSize;
        return sign_api("/Contract/getContractMyPositions", data);
    },
    


    /**
     * ------------------------------------------------------------------
     * 合约交易市场列表
     * ------------------------------------------------------------------
     */
    getContractCoin: function() {
        var data = {};
        return market_api("/Market/getContractCoin", data);
    },
    /**
     * ------------------------------------------------------------------
     * 单个合约交易市场
     * ------------------------------------------------------------------
     */
    getSingleContractCoin: function(contracId) {
        var data = {};
        data.contracId = contracId;
        return market_api("/Market/getContractCoin", data);
    },
    /**
     * ------------------------------------------------------------------
     * 获取最新行情
     * @param contracId    101
     * ------------------------------------------------------------------
     */
    getContractTickers: function(contracId) {
        var data = {};
        data.contracId = contracId;
        return market_api("/Market/getContractTickers", data);
    },
    /**
     * ------------------------------------------------------------------
     * 获取K线数据
     * @param contractId        101
     * @param type    string [1min 、5min 、15min 、30min 、60min、 4hour 、 1day 、5day 、1week、 1mon]
     * @param size          number 1 ~ 300 (can empty)
     * ------------------------------------------------------------------
     */
    getContractKline: function(contractId, type, size) {
        var data = {};
        data.contractId = contractId;
        data['type'] = type;
        if(size) data['size'] = size;
        return market_api("/Market/getContractKline", data);
    },
    /**
     * ------------------------------------------------------------------
     * 获取深度数据
     * @param contractId    101
     * ------------------------------------------------------------------
     */
    getContractOrderBook: function(contractId) {
        var data = {};
        data.contractId = contractId;
        return market_api("/Market/getContractOrderBook", data);
    },
    /**
     * ------------------------------------------------------------------
     * 获取某个合约交易历史 (Get contract trade history)
     * @param contractId    101
     * ------------------------------------------------------------------
     */
    getContractTradesHistory: function(contractId, pageSize) {
        var data = {};
        data.contractId = contractId;
        if(pageSize) data['pageSize'] = pageSize;
        return market_api("/Market/getContractTradesHistory", data);
    }


}


module.exports = CONTRACT_API;