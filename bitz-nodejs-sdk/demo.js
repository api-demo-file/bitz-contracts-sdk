
const _API = require('./contract_api');

/**
 * ------------------------------------------------------------------
 * @param apiServer     必须
 * @param apiKey        必须
 * @param secretKey     必须
 * ------------------------------------------------------------------
 */
var apiServer   = 'https://';
var apiKey 		= '';
var secretKey 	= '';
_API.setConfig( apiServer, apiKey, secretKey );


// 获取资产
_API.getContractMyPositions(101).then(function(data){
	console.log(data);
}).catch(function(err){
	console.log(err);
});





// ------------------------------------------------

var demo_cmd = `
	// 行情接口
	# node demo.js getContractCoin
	# node demo.js getContractTickers
	# node demo.js getContractTradesHistory
	# node demo.js getContractOrderBook 
	# node demo.js getContractKline 
	
	// 合约账户权益
	#	node demo.js getContractAccountInfo
	// 做空
	#	node demo.js addContractTrade 
	// 做多
	#	node demo.js addContractTrade 
	// 取消订单
	#	node demo.js cancelContractTrade 
	// 活动委托
	#	node demo.js getContractOrder 
	// 当前仓位
	#	node demo.js getContractActivePositions 
	// 委托历史
	#	node demo.js getContractMyHistoryTrade 
	// 交易历史
	#	node demo.js getContractMyTrades 
	// 当个委托成交明细
	#	node demo.js getContractTradeResult 
	// 已平仓仓位
	#	node demo.js getContractMyPositions 
`;

function commandParser(){

	var argv = process.argv.splice(2);
	var func = argv[0];
	var params = argv.splice(1);
	var str_params = '';
	if( typeof(_API[func]) == 'function' ){
		if(params.length>0){
			str_params = params[0];
		}
		//	
		var script = `
			_API.${func}(${str_params})
			.then(
				function(d){
					log('-------------Response-----------');
					var data = JSON.parse(d);
					log(data);
					log(data.data);
				}
			)
			.catch(
				function(err){
					log('-------------Error-----------');
					log(err)
				}
			);`;
		log("----start-----"+func);
		log(script);
		eval(script);
	}else if(func == 'help'){
		log(demo_cmd);
	}
}
function log(msg){
	console.log(msg);
	console.log(" ");
}

commandParser();