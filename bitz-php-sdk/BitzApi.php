<?php
/**
 * Created by Bitz
 * Date 2019-10-12
 */

class BitzApi
{

	protected $apiKey = '';
	protected $secretKey = '';
	protected $tradePwd = '';
	protected $url = '';


	public function __construct($options = null)
	{
		try {
			if (is_array($options))
			{
				foreach ($options as $option => $value)
				{
					$this->$option = $value;
				}
			}
			else
			{
				return false;
			}
		}
		catch (PDOException $e) {
			throw new Exception($e->getMessage());
		}
	}


	/**
	 * 获取合约交易市场列表
	 * @param contractId 选填 int
	 * @return array
	 */
	public function getContractCoin(int $contractId = 0){
		$url = $this->url.'/Market/getContractCoin?contractId='.intval($contractId);
		return $this->httpRequest($url);
	}


	/**
	 * 获取合约交易最新行情
	 * @param contractId 选填 int
	 * @return array
	 */
	public function getContractTickers(int $contractId = 0){
		$url = $this->url.'/Market/getContractTickers?contractId='.intval($contractId);
		return $this->httpRequest($url);
	}


	/**
	 * 获取合约交易的市场深度
	 * @param contractId 必填 int
	 * @param depth 选填 string
	 * @return array
	 */
	public function getContractOrderBook(array $parms){
		$url = $this->url.'/Market/getContractOrderBook?'.http_build_query($parms);
		return $this->httpRequest($url);
	}


	/**
	 * 获取某个合约的成交历史
	 * @param contractId 必填 int
	 * @param pageSize 选填 int
	 * @return array
	 */
	public function getContractTradesHistory(array $parms){
		$url = $this->url.'/Market/getContractTradesHistory?'.http_build_query($parms);
		return $this->httpRequest($url);
	}


	/**
	 * 获取某个合约市场的K线数据
	 * @param contractId 必填 int
	 * @param type 必填 string
	 * @param size 选填 int
	 * @return array
	 */
	public function getContractKline(array $parms){
		$url = $this->url.'/Market/getContractKline?'.http_build_query($parms);
		return $this->httpRequest($url);
	}


	/**
	 * 获取合约账户权益(资产)
	 * @return array
	 */
	public function getContractAccountInfo(){
		$sendData  = $this->getData();
		$url = $this->url.'/Contract/getContractAccountInfo';
		return $this->httpRequest($url,$sendData);
	}


	/**
	 * 下单 做多/做空
	 * @param contractId 必填  int
	 * @param price      选填  float
	 * @param amount     必填  int
	 * @param leverage   必填  float
	 * @param direction  必填  int
	 * @param type       必填  string
	 * @param isCross    必填  int
	 * @return array
	 */
	public function addContractTrade(array $parms){
		$sendData  = $this->getData($parms);
		$url = $this->url.'/Contract/addContractTrade';
		return $this->httpRequest($url,$sendData);
	}


	/**
	 * 撤销委托单
	 * @param entrustSheetId 必填 int
	 * @return array
	 */
	public function cancelContractTrade(array $parms){
		$sendData  = $this->getData($parms);
		$url = $this->url.'/Contract/cancelContractTrade';
		return $this->httpRequest($url,$sendData);
	}


	/**
	 * 获取当前委托单
	 * @param contractId 必填 int
	 * @return array
	 */
	public function getContractOrder(array $parms){
		$sendData  = $this->getData($parms);
		$url = $this->url.'/Contract/getContractOrder';
		return $this->httpRequest($url,$sendData);
	}


	/**
	 * 获取当前未平仓位
	 * @param contractId 必填 int
	 * @return array
	 */
	public function getContractActivePositions(array $parms){
		$sendData  = $this->getData($parms);
		$url = $this->url.'/Contract/getContractActivePositions';
		return $this->httpRequest($url,$sendData);
	}


	/**
	 * 获取历史委托单详情
	 * @param contractId   必填  int
	 * @param pageSize     选填  int
	 * @param page         选填  int
	 * @return array
	 */
	public function getContractMyHistoryTrade(array $parms){
		$sendData  = $this->getData($parms);
		$url = $this->url.'/Contract/getContractMyHistoryTrade';
		return $this->httpRequest($url,$sendData);
	}


	/**
	 * 获取我的已平仓位
	 * @param contractId   必填  int
	 * @param pageSize     选填  int
	 * @param page         选填  int
	 * @return array
	 */
	public function getContractMyPositions(array $parms){
		$sendData  = $this->getData($parms);
		$url = $this->url.'/Contract/getContractMyPositions';
		return $this->httpRequest($url,$sendData);
	}


	/**
	 * 获取单个或多个委托单明细
	 * @param entrustSheetIds  必填  string
	 * @return array
	 */
	public function getContractOrderResult(array $parms){
		$sendData  = $this->getData($parms);
		$url = $this->url.'/Contract/getContractOrderResult';
		return $this->httpRequest($url,$sendData);
	}


	/**
	 * 获取某个委托的成交明细
	 * @param entrustSheetId   必填  int
	 * @return array
	 */
	public function getContractTradeResult(array $parms){
		$sendData  = $this->getData($parms);
		$url = $this->url.'/Contract/getContractTradeResult';
		return $this->httpRequest($url,$sendData);
	}


	/**
	 * 获取我的成交历史
	 * @param contractId   必填  int
	 * @param pageSize     选填  int
	 * @param page         选填  int
	 * @return array
	 */
	public function getContractMyTrades(array $parms){
		$sendData  = $this->getData($parms);
		$url = $this->url.'/Contract/getContractMyTrades';
		return $this->httpRequest($url,$sendData);
	}


	//获取带签名的参数
	protected function getData($data = null){
		$baseArr = array(
				'apiKey'      =>  $this->apiKey,
				'timeStamp'   =>  time(),
				'nonce'		  =>  $this->getRandomString(6),
				);

		if(isset($data)){
			$sendData = array_merge($baseArr,$data);
		}else{
			$sendData = $baseArr;
		}
		$sendData['sign'] = $this->getSign($sendData);
		return $sendData;
	}


	//参数签名
	protected function getSign($data){
		ksort($data);
		foreach( $data as $k=>$v)
		{
			if(!$v){
				unset($data[$k]);
			}
		}

		$dataStr = '';
		foreach ($data as $k => $v)
		{

			$dataStr .= $k.'='.$v."&";
		}
		return md5(trim($dataStr,"&").$this->secretKey);
	}


	//获取随机字符串nonce
	protected function getRandomString($len, $chars=null){
		if (is_null($chars)){
			$chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
		}

		for ($i = 0, $str = '', $lc = strlen($chars)-1; $i < $len; $i++){
			$str .= $chars[mt_rand(0, $lc)];
		}
		return $str;
	}
	

	//HTTP请求
	protected function httpRequest($url,$data = null){
		$curl = curl_init();
		curl_setopt($curl, CURLOPT_URL, $url);
		curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, FALSE);
		curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, FALSE);
		curl_setopt($curl, CURLOPT_TIMEOUT, 100);
		if(!empty($data)){
			curl_setopt($curl,CURLOPT_POST,1);
			curl_setopt($curl,CURLOPT_POSTFIELDS,$data);
		}
		curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);
		$output = curl_exec($curl);
		curl_close($curl);
		return json_decode($output,1);
	}
}

// $bitz = new BitzApi();
//获取合约所有市场
// $res = $bitz->getContractCoin();
//获取合约单个市场数据
// $contractId = 1;
// $res = $bitz->getContractCoin($contractId);

//获取合约市场最新行情
// $contractId = 101;
// $res = $bitz->getContractTickers($contractId);

//获取合约市场深度
// $parms = array('contractId'=>101,'depth'=>'5');
// $res = $bitz->getContractOrderBook($parms);

//获取合约市场成交历史
// $parms = array('contractId'=>101,'pageSize'=>'10');
// $res = $bitz->getContractTradesHistory($parms);

//获取合约市场k线数据
// $parms = array('contractId'=>101,'type'=>'5m','size'=>100);
// $res = $bitz->getContractKline($parms);

//获取我的账户权益
// $res = $bitz->getContractAccountInfo();

//下单委托
// $parms = array('contractId'=>101,'price'=>8200.01,'amount'=>100,'leverage'=>10,'direction'=>1,'type'=>'limit','isCross'=>-1);
// $res = $bitz->addContractTrade($parms);

//取消委托
// $parms = array('entrustSheetId'=>15484932);
// $res = $bitz->cancelContractTrade($parms);

//获取我的当前委托
// $parms = array('contractId'=>101);
// $res = $bitz->getContractOrder($parms);

//获取我的未平仓位
// $parms = array('contractId'=>101);
// $res = $bitz->getContractActivePositions($parms);

//获取我的委托历史
// $parms = array('contractId'=>101,'page'=>1,'pageSize'=>50);
// $res = $bitz->getContractMyHistoryTrade($parms);

//获取我的已平仓位
// $parms = array('contractId'=>101,'page'=>1,'pageSize'=>50);
// $res = $bitz->getContractMyPositions($parms);

//获取单个或多个委托单的委托明细
// $parms = array('entrustSheetIds'=>'15485465,15484932,14023097');
// $res = $bitz->getContractOrderResult($parms);

//获取某个委托单的成交明细
// $parms = array('entrustSheetId'=>14023097);
// $res = $bitz->getContractTradeResult($parms);

//获取我的成交历史
// $parms = array('contractId'=>101,'page'=>1,'pageSize'=>50);
// $res = $bitz->getContractMyTrades($parms);

// var_dump($res);die;



















