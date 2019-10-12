# NODE DEMO 

## 安装依赖

```
	npm install 

	node test.js help

```

## 使用方法

```
	const _API = require('./api');
	/**
	 * @param apiServer     必须
	 * @param apiKey        必须
	 * @param secretKey     必须
	 */
	var apiServer   = '';
	var apiKey 		= '';
	var secretKey 	= '';
	_API.setConfig( apiServer, apiKey, secretKey, tradePwd );

	// 获取合约账户权益
	_API.getContractAccountInfo().then(function(data){
		console.log(data);
	}).catch(function(err){
		console.log(err);
	});

	// 获取合约市场
	_API.getContractCoin().then(function(data){
		console.log(data);
	}).catch(function(err){
		console.log(err);
	});

```

