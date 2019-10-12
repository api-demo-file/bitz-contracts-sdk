
import java.io.IOException;
import java.util.*;

public class BitzApi {

    private String apiKey;
    private String secretKey;
    private String tradePwd;
    private String WEB_BASE; //


    public BitzApi(String api_server, String api_key, String secret_key) {
        this.WEB_BASE = api_server;
        this.apiKey = api_key;
        this.secretKey = secret_key;
        this.tradePwd = "";
    }


    /**
     * 获取合约账户信息
     *
     * @return String
     */
    public String getAccountInfo() throws IOException {

        HashMap<String, String> params = new HashMap<String, String>();
        return this.sign_api("/Contract/getContractAccountInfo", params);
    }


    /**
     * 下单(做多/做空)
     *
     * @param contract_id 合约ID
     * @param price       价格
     * @param amount      合约张数
     * @param leverage    杠杆倍数x
     * @param direction   委托方向
     * @param type        委托类型
     * @param is_cross    是否全仓
     */

    public String getPlaceOrder(String contract_id, String price, String amount, String leverage, String direction, String type, String is_cross) throws IOException {

        HashMap<String, String> params = new HashMap<String, String>();
        params.put("contractId", contract_id);
        params.put("amount", amount);
        params.put("leverage", leverage);
        params.put("direction", direction);
        params.put("type", type);
        params.put("isCross", is_cross);
        if (!"market".equals(type)) {
            params.put("price", price);
        }
        return this.sign_api("/Contract/addContractTrade", params);
    }


    /**
     * 取消委托
     * order_id  委托单id
     */

    public String getCancelOrder(String order_id) throws IOException {

        HashMap<String, String> params = new HashMap<String, String>();
        params.put("entrustSheetId", order_id);
        return this.sign_api("/Contract/cancelContractTrade", params);

    }


    /**
     * 获取我的活动委托
     * contract_id  合约ID
     */

    public String getActiveOrders(String contract_id) throws IOException {

        HashMap<String, String> params = new HashMap<String, String>();
        params.put("contractId", contract_id);
        return this.sign_api("/Contract/getContractOrder", params);

    }


    /**
     * 获取当前未平仓位
     * contract_id  合约ID
     */

    public String getMyActivePositions(String contract_id) throws IOException {

        HashMap<String, String> params = new HashMap<String, String>();
        params.put("contractId", contract_id);
        return this.sign_api("/Contract/getContractActivePositions", params);
    }


    /**
     * 获取某个合约交易的市场深度
     *
     * @param contract_id 合约ID
     * @return String
     */

    public String getMxDataOrderBook(String contract_id) throws IOException {

        Map<String, String> params = new HashMap<String, String>();
        params.put("contractId", contract_id);
        return this.market_api("/Market/getContractOrderBook", params);
    }


    /**
     * 获取某个合约交易最新行情
     *
     * @param contract_id (合约ID)
     * @return String
     */

    public String getMxDataTickers(String contract_id) throws IOException {

        Map<String, String> params = new HashMap<String, String>();
        params.put("contractId", contract_id);
        return this.market_api("/Market/getContractTickers", params);
    }


    /**
     * 获取某个合约交易历史
     *
     * @param contract_id 合约ID
     * @param page_size
     * @return
     * @throws IOException
     */
    public String getMxDataTrades(String contract_id, String page_size) throws IOException {

        Map<String, String> params = new HashMap<String, String>();
        params.put("contractId", contract_id);
        params.put("page_size", page_size);
        return this.market_api("/Market/getContractTradesHistory", params);

    }

    /**
     * 获取某个合约交易k线
     *
     * @param contract_id 合约ID
     * @param type        K线类型(5m, 15m, 30m, 1h, 4h, 8h, 1d)
     * @param amount      获取数据的数量
     * @return
     * @throws IOException
     */

    public String getMxDatakData(String contract_id, String type, String amount) throws IOException {

        Map<String, String> params = new HashMap<String, String>();
        params.put("contractId", contract_id);
        params.put("type", type);
        params.put("size", amount);
        return this.market_api("/Market/getContractKline", params);

    }


    /**
     * 获取我的委托历史
     * contract_id  合约ID
     */

    public String getMyOrders(String contract_id) throws IOException {

        HashMap<String, String> params = new HashMap<String, String>();
        params.put("contractId", contract_id);
        return this.sign_api("/Contract/getContractMyHistoryTrade", params);

    }


    /**
     * 获取我的已平仓位
     * contract_id  合约ID
     */

    public String getMyPositions(String contract_id, String page, String pageSize) throws IOException {

        HashMap<String, String> params = new HashMap<String, String>();
        params.put("contractId", contract_id);
        params.put("page", page);
        params.put("pageSize", pageSize);
        return this.sign_api("/Contract/getContractMyPositions", params);

    }


    /**
     * 获取某个委托的成交明细
     * order_id  委托单ID
     */

    public String getOrderMatchResults(String order_id) throws IOException {

        HashMap<String, String> params = new HashMap<String, String>();
        params.put("entrustSheetId", order_id);
        return this.sign_api("/Contract/getContractTradeResult", params);

    }


    /**
     * 获取合约交易市场列表
     *
     * @param contract_id (合约id)  必传 false
     */
    public String getMxDataAontracts(String contract_id) throws IOException {

        Map<String, String> params = new HashMap<String, String>();
        params.put("contractId", contract_id);
        return this.market_api("/Market/getContractCoin", params);
    }


    /**
     * 获取合约交易的成交历史
     * contract_id  合约ID
     */

    public String getMyTrades(String contract_id) throws IOException {

        Map<String, String> params = new HashMap<String, String>();
        params.put("contractId", contract_id);
        return this.market_api("/Market/getContractTradesHistory", params);
    }


    /**
     * @param path
     * @param params
     * @return
     */
    private String market_api(String path, Map<String, String> params) {
        String result = "";
        String strParams = SignUtil.createLinkString(params);
        try {
            result = HttpUtilManager.getInstance().requestHttpGet(WEB_BASE, path, strParams);
        } catch (Exception e) {
            result = "";
        }
        return result;
    }

    private String sign_api(String path, Map<String, String> params) {
        String result = "";

        long timeStamp = System.currentTimeMillis();
        timeStamp = timeStamp / 1000;
        params.put("apiKey", this.apiKey);
        params.put("timeStamp", String.valueOf(timeStamp));
        params.put("nonce", String.valueOf(timeStamp).substring(0, 6));
        String sign = SignUtil.buildSign(params, this.secretKey);
        params.put("sign", sign);
        //
        try {
            result = HttpUtilManager.getInstance().requestHttpPost(WEB_BASE, path, params);
        } catch (Exception e) {

        }
        return result;
    }


}
