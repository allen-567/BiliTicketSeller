package bilibili

import (
	"BiliTicketSeller/bilibili/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func CancelOrder(sessdata string, orderId string) {

	url := "https://mall.bilibili.com/mall-c/order/detail/cancel?orderId=" + orderId
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Host", "mall.bilibili.com")
	req.Header.Add("Connection", "close")
	req.Header.Add("APP-KEY", "iphone")
	req.Header.Add("ENV", "prod")
	req.Header.Add("User-Agent", "bili-universal/71801100 CFNetwork/1404.0.5 Darwin/22.3.0 os/ios model/iPhone XS mobi_app/iphone build/71801100 osVer/16.3.1 network/2 channel/AppStore mallVersion/7181000 mNative/1")
	req.Header.Add("Accept-Encoding", "deflate")
	req.Header.Add("Cookie", "SESSDATA="+sessdata+";")

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	return
	//}
	//orderCancelResp := &OrderCancelResp{}
	//json.Unmarshal(body, orderCancelResp)
}

// {"code":0,"message":null,"data":{"codeType":1,"codeMsg":"操作成功","vo":1},"errtag":0}
type DeleteOrderResp struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    struct {
		CodeType int    `json:"codeType"`
		CodeMsg  string `json:"codeMsg"`
		Vo       int    `json:"vo"`
	} `json:"data"`
	Errtag int `json:"errtag"`
}

func DeleteOrder(sessdata string, orderId string) *DeleteOrderResp {
	url := "https://mall.bilibili.com/mall-c/order/delete?orderId=" + orderId
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil
	}
	req.Header.Add("Host", "mall.bilibili.com")
	req.Header.Add("Connection", "close")
	req.Header.Add("APP-KEY", "iphone")
	req.Header.Add("ENV", "prod")
	req.Header.Add("User-Agent", "bili-universal/71801100 CFNetwork/1404.0.5 Darwin/22.3.0 os/ios model/iPhone XS mobi_app/iphone build/71801100 osVer/16.3.1 network/2 channel/AppStore mallVersion/7181000 mNative/1")
	req.Header.Add("Accept-Encoding", "deflate")
	req.Header.Add("Cookie", "SESSDATA="+sessdata+";")

	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	deleteOrderResp := &DeleteOrderResp{}
	json.Unmarshal(body, deleteOrderResp)
	return deleteOrderResp
}

func CreateOrder(sessData string, itemsId string, skuId string, num string, payCount string) *structs.OrderResult {

	url := "https://mall.bilibili.com/magic-c/cart/na/ordercreate"
	method := "POST"
	payloadStr := `{"vtoken":"","activityInfo":{"activityId":0,"type":0,"marketingId":null},"secKill":0,"crystalDiscountCnt":-1,"benefitAmountAll":"0","buyerId":0,"cartOrderType":6,"deviceInfo":"WEB","deviceType":3,"distId":0,"expressTotalAmountAll":"0","from":"mall_home_search","recId":"","source":"category_mls","activityId":"","invoiceId":0,"itemsTotalAmountAll":"` + payCount + `","orders":[{"buyerComment":"","items":[{"cartId":0,"itemsId":` + itemsId + `,"amount":"0","skuId":` + skuId + `,"skuNum":` + num + `}],"shopId":2233,"shopIsNotice":1}],"payTotalAmountAll":"` + payCount + `","couponCodeId":"","payType":0,"payTypeStatus":1}`

	payload := strings.NewReader(payloadStr)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Host", "mall.bilibili.com")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Origin", "https://mall.bilibili.com")
	req.Header.Add("Accept-Encoding", "deflate")
	req.Header.Add("Cookie", "SESSDATA="+sessData+";")
	req.Header.Add("Connection", "close")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/614.4.6.0.6 (KHTML, like Gecko) Mobile/20D67 BiliApp/71800100 os/ios model/iPhone XS mobi_app/iphone build/71800100 osVer/16.3.1 network/2 channel/AppStore c_locale/zh-Hans_CN s_locale/zh-Hans_CN disable_rcmd/0 mallVersion/7180000 mVersion/179 flutterNotch/1")
	req.Header.Add("Referer", "https://mall.bilibili.com/neul-next/index.html?activeType=1&from=mall_home_search&outsideMall=no&shopId=2233&noTitleBar=1&themeId=0&page=magic-detail_detail")
	req.Header.Add("Accept-Language", "zh-CN,zh-Hans;q=0.9")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	orderResult := &structs.OrderResult{}
	json.Unmarshal(body, orderResult)
	return orderResult
}

func GetBlindBoxInfo(itemsId string) *structs.GetBlindBoxInfoResp {

	url := "https://mall.bilibili.com/magic-c-search/blind_box/info?itemsId=" + itemsId
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Accept-Language", "zh-CN,zh-Hans;q=0.9")
	req.Header.Add("Host", "mall.bilibili.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Safari/605.1.15")
	req.Header.Add("Referer", "https://mall.bilibili.com/neul-next/index.html?activeType=1&noTitleBar=1&from=history&page=magic-detail_detail&msource=history&outsideMall=yes")
	req.Header.Add("Accept-Encoding", "deflate")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	result := &structs.GetBlindBoxInfoResp{}
	json.Unmarshal(body, result)
	return result
}

func GetBlindBoxItemsId(url string) string {
	re := regexp.MustCompile(`itemsId=(\d+)`)
	match := re.FindStringSubmatch(url)

	if len(match) > 1 {
		itemsId := match[1]
		return itemsId
	}
	if strings.Contains(url, "b23.tv") {
		url = GetUrlFromSort(url)
		match = re.FindStringSubmatch(url)
		if len(match) > 1 {
			itemsId := match[1]
			return itemsId
		}
	}
	return ""
}
func GetUrlFromSort(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ""
	}

	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	return resp.Request.URL.String()
}
