
import requests
import os
import hashlib
import time
import random
import configparser


nowtime = str(int(time.time()))


proDir = os.path.split(os.path.realpath(__file__))[0]
configPath = os.path.join(proDir, "info.conf")
path = os.path.abspath(configPath)
conf = configparser.ConfigParser()
conf.read(path, encoding="utf-8-sig")

marketURL = conf.get('Key', 'MarketURL')
tradeURL = conf.get('Key', 'TradeURL')


class bitzSwap:
    def __init__(self, market_url=None, trade_url=None):
        if(market_url == None):
            self.__market_url = marketURL
        else:
            self.__market_url = market_url

        if(trade_url == None):
            self.__trade_url = tradeURL
        else:
            self.__trade_url = trade_url

    '''
    获取合约交易市场列表
    '''

    def getContractCoin(self, contractId=None):
        param = {'contractId': contractId}
        response = requests.get(self.__market_url + 'getContractCoin', param)
        return response.json()

    '''
    获取合约K线数据
    '''

    def getContractKline(self, contractId, type, size):
        param = {'contractId': contractId,
                 'type': type,
                 'size': size
                 }
        response = requests.get(self.__market_url + 'getContractKline', param)
        return response.json()

    '''
    获取合约交易的市场深度
    '''

    def getContractOrderBook(self, contractId):
        param = {'contractId': contractId}
        response = requests.get(
            self.__market_url + 'getContractOrderBook', param)
        return response.json()

    '''
    获取合约交易的成交历史
    '''

    def getContractTradesHistory(self, contractId, pageSize=None):
        param = {'contractId': contractId,
                 'pageSize': pageSize
                 }
        response = requests.get(
            self.__market_url + 'getContractTradesHistory', param)
        return response.json()

    '''
    获取合约交易最新行情
    '''

    def getContractTickers(self, contractId=None):
        param = {'contractId': contractId}
        response = requests.get(
            self.__market_url + 'getContractTickers', param)
        return response.json()

    '''
    合约交易下单
    '''

    def addContractTrade(self, contractId, price, amount, leverage, direction, type, is_cross):
        param = {'contractId': contractId,
                 'price': price,
                 'amount': amount,
                 'leverage': leverage,
                 'direction': direction,
                 'type': type,
                 'isCross': is_cross}
        param['sign'] = self.build_sign(param)
        response = requests.post(
            self.__trade_url + 'addContractTrade', param)
        return response.json()

    '''
    合约交易取消委托
    '''

    def cancelContractTrade(self, entrustSheetId):
        param = {'entrustSheetId': entrustSheetId}
        param['sign'] = self.build_sign(param)
        response = requests.post(
            self.__trade_url + 'cancelContractTrade', param)
        return response.json()

    '''
    获取当前未平仓位
    '''

    def getContractActivePositions(self, contractId):
        param = {'contractId': contractId}
        param['sign'] = self.build_sign(param)
        response = requests.post(
            self.__trade_url + 'getContractActivePositions', param)
        return response.json()

    '''
    获取合约账户权益（资产）
    '''

    def getContractAccountInfo(self):
        param = {}
        param['sign'] = self.build_sign(param)
        response = requests.post(
            self.__trade_url + 'getContractAccountInfo', param)
        return response.json()

    '''
    获取我的已平仓仓位
    '''

    def getContractMyPositions(self, contractId, page, pageSize):
        param = {'contractId': contractId,
                 'page': page,
                 'pageSize': pageSize
                 }
        param['sign'] = self.build_sign(param)
        response = requests.post(
            self.__trade_url + 'getContractMyPositions', param)
        return response.json()

    '''
    单个或多个委托单明细
    '''

    def getContractOrderResult(self, entrustSheetId):
        param = {'entrustSheetId': entrustSheetId}
        param['sign'] = self.build_sign(param)
        response = requests.post(
            self.__trade_url + 'getContractOrderResult', param)
        return response.json()

    '''
    获取我的活动委托
    '''

    def getContractOrder(self, contractId):
        param = {'contractId': contractId}
        param['sign'] = self.build_sign(param)
        response = requests.post(
            self.__trade_url + 'getContractOrder', param)
        return response.json()

    '''
    获取某个委托的成交明细
    '''

    def getContractTradeResult(self, entrustSheetId):
        param = {'entrustSheetId': entrustSheetId}
        param['sign'] = self.build_sign(param)
        response = requests.post(
            self.__trade_url + 'getContractTradeResult', param)
        return response.json()

    '''
    获取我的委托历史
    '''

    def getContractMyHistoryTrade(self, contractId, page, pageSize):
        param = {'contractId': contractId,
                 'page': page,
                 'pageSize': pageSize
                 }
        param['sign'] = self.build_sign(param)
        response = requests.post(
            self.__trade_url + 'getContractMyHistoryTrade', param)
        return response.json()

    '''
    获取我的成交历史
    '''

    def getContractMyTrades(self, contractId=None):
        param = {'contractId': contractId}
        param['sign'] = self.build_sign(param)
        response = requests.post(
            self.__trade_url + 'getContractMyTrades', param)
        return response.json()

    '''
    签名排序方法
    '''

    def build_sign(self, param):
        sign = ''
        param['apiKey'] = conf.get('Key', 'ApiKey')
        param['timeStamp'] = nowtime
        param['nonce'] = str(random.randint(100000, 999999))
        for key in sorted(param.keys()):
            sign += key + '=' + str(param[key]) + '&'
        data = sign[:-1] + conf.get('Key', 'Secret')

        return hashlib.md5(data.encode("utf8")).hexdigest()
