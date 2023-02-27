package main

import (
	"BiliTicketSeller/bilibili"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

var orderIds []string

func main() {
	sessData := ""
	for {
		fmt.Print("请输入您的Sessdata:")
		i, err := fmt.Scanf("%s\n", &sessData)
		if i > 0 && err == nil {
			break
		}
	}
	sessData = strings.TrimSuffix(sessData, "\n")
	for {
		itemsId := ""
		for {
			url := ""
			fmt.Print("请输入池子链接：")
			i, err := fmt.Scanf("%s\n", &url)
			if i == 0 || err != nil {
				fmt.Println("池子链接不正确！")
				continue
			}
			url = strings.TrimSuffix(url, "\n")
			itemsId = bilibili.GetBlindBoxItemsId(url)
			if itemsId == "" {
				fmt.Println("池子链接不正确！")
				continue
			}
			break
		}
		blindBoxInfo := bilibili.GetBlindBoxInfo(itemsId)
		maxDraw := 0
		for _, item := range blindBoxInfo.Data.SubSkuList {
			if item.Type == 0 {
				maxDraw++
			}
		}
		fmt.Println("池子:" + blindBoxInfo.Data.Name)
		fmt.Println("最大抽数:" + strconv.Itoa(maxDraw))
		fmt.Println("单抽价格:" + strconv.FormatFloat(blindBoxInfo.Data.Price, 'f', 2, 64))

		draw := 0
		for {
			for {
				fmt.Print("请输入需要的抽数（输入-1为释放并删除之前的订单，输入0为更换池子）：")
				fmt.Scanf("%d\n", &draw)
				if draw == 0 || draw == -1 {
					break
				}
				if draw < 1 || draw > maxDraw {
					fmt.Println("抽数输入错误！")
					continue
				}
				break
			}
			if draw == 0 {
				break
			} else if draw == -1 {
				if len(orderIds) == 0 {
					fmt.Println("目前没有可以释放订单！")
					continue
				}
				fmt.Println("开始释放并删除订单...")
				for _, oid := range orderIds {
					fmt.Printf("释放订单：%s ... ", oid)
					bilibili.CancelOrder(sessData, oid)
					deleteResp := bilibili.DeleteOrder(sessData, oid)
					if deleteResp == nil {
						fmt.Println("失败")
						continue
					}
					if deleteResp.Code != 0 {
						fmt.Printf("失败(%s)\n", deleteResp.Message)
						continue
					}
					fmt.Println("成功")
				}
				orderIds = []string{}
				continue
			}
			price := draw * int(blindBoxInfo.Data.Price)
			fmt.Printf("正在为池子<%s>创建订单，抽数：%d,总价：%d元。\n", blindBoxInfo.Data.Name, draw, price)
			skuId := strconv.Itoa(blindBoxInfo.Data.SkuId)
			result := bilibili.CreateOrder(sessData, itemsId, skuId, strconv.Itoa(draw), strconv.Itoa(price))
			if result.Code == 83001002 {
				fmt.Println("登录无效，请检查Sessdata。")
				return
			}
			if result.Data.CodeMsg == nil && result.Data.PayInfo.ExtData != "" {
				type orderIdStruct struct {
					OrderId int64 `json:"orderId"`
				}
				oidS := &orderIdStruct{}
				json.Unmarshal([]byte(result.Data.PayInfo.ExtData), oidS)
				orderIds = append(orderIds, strconv.FormatInt(oidS.OrderId, 10))
				fmt.Println("订单创建成功，请在客户端完成付款。")
			} else {
				fmt.Printf("订单创建失败，原因：%s。\n", result.Data.CodeMsg)
			}
		}
	}
}
