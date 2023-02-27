package structs

//type OrderCancelResp struct {
//	Code    int    `json:"code"`
//	Message string `json:"message"`
//	Data    struct {
//		CodeType int    `json:"codeType"`
//		CodeMsg  string `json:"codeMsg"`
//		Vo       int    `json:"vo"`
//	} `json:"data"`
//	Errtag int `json:"errtag"`
//}

type OrderResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		CartTotalAmount        string      `json:"cartTotalAmount"`
		CartTotalAmountAll     string      `json:"cartTotalAmountAll"`
		ItemsTotalAmountAll    string      `json:"itemsTotalAmountAll"`
		ExpressTotalAmountAll  string      `json:"expressTotalAmountAll"`
		PayTotalAmountAll      string      `json:"payTotalAmountAll"`
		BenefitAmountAll       string      `json:"benefitAmountAll"`
		DiscountTotalAmountAll string      `json:"discountTotalAmountAll"`
		TaxTotalAmountAll      string      `json:"taxTotalAmountAll"`
		ItemsTaxTotalAmount    string      `json:"itemsTaxTotalAmount"`
		ExpressTaxTotalAmount  string      `json:"expressTaxTotalAmount"`
		PayErrorUrl            interface{} `json:"payErrorUrl"`
		Mid                    int         `json:"mid"`
		OrderList              []int64     `json:"orderList"`
		ValidList              []struct {
			Amount              string      `json:"amount"`
			MarketAmount        interface{} `json:"marketAmount"`
			FrontAmount         string      `json:"frontAmount"`
			PreDepositAmount    string      `json:"preDepositAmount"`
			TaxAmount           string      `json:"taxAmount"`
			AdditionText        interface{} `json:"additionText"`
			OrderId             interface{} `json:"orderId"`
			FinalMoney          interface{} `json:"finalMoney"`
			FinalAmount         interface{} `json:"finalAmount"`
			FinalDiscount       interface{} `json:"finalDiscount"`
			FinalDiscountAmount interface{} `json:"finalDiscountAmount"`
			Valid               string      `json:"valid"`
			Labels              []struct {
				ActType             int         `json:"actType"`
				Text                string      `json:"text"`
				TextColor           interface{} `json:"textColor"`
				TextBackgroundColor interface{} `json:"textBackgroundColor"`
			} `json:"labels"`
			ActivityInfos             interface{}   `json:"activityInfos"`
			AdvancePublishTime        interface{}   `json:"advancePublishTime"`
			AdvanceCountdownTime      interface{}   `json:"advanceCountdownTime"`
			ActivityType              interface{}   `json:"activityType"`
			Type                      interface{}   `json:"type"`
			PresaleStartOrderTime     interface{}   `json:"presaleStartOrderTime"`
			PresaleEndOrderTime       interface{}   `json:"presaleEndOrderTime"`
			ShowPromotionTag          interface{}   `json:"showPromotionTag"`
			PurchaseSaleMode          interface{}   `json:"purchaseSaleMode"`
			ExtraItemsType            interface{}   `json:"extraItemsType"`
			SkuNotice                 interface{}   `json:"skuNotice"`
			SkuPromotionText          interface{}   `json:"skuPromotionText"`
			AsyncSku                  bool          `json:"asyncSku"`
			CartId                    int           `json:"cartId"`
			SkuNum                    int           `json:"skuNum"`
			SkuId                     int           `json:"skuId"`
			ItemsId                   int           `json:"itemsId"`
			CateId                    int           `json:"cateId"`
			ItemsName                 string        `json:"itemsName"`
			ItemsType                 []interface{} `json:"itemsType"`
			CartIsCheck               int           `json:"cartIsCheck"`
			CartIsDelete              int           `json:"cartIsDelete"`
			ShopId                    int           `json:"shopId"`
			ShopName                  string        `json:"shopName"`
			Sn                        string        `json:"sn"`
			Price                     float64       `json:"price"`
			MarketPrice               int           `json:"marketPrice"`
			FrontPrice                float64       `json:"frontPrice"`
			PreDepositPrice           float64       `json:"preDepositPrice"`
			Storage                   int           `json:"storage"`
			SpuLimitNum               int           `json:"spuLimitNum"`
			WhiteLimitNum             interface{}   `json:"whiteLimitNum"`
			VipLevel                  interface{}   `json:"vipLevel"`
			MemberLevel               interface{}   `json:"memberLevel"`
			StorageAlert              int           `json:"storageAlert"`
			StorageStatus             interface{}   `json:"storageStatus"`
			SkuSpec                   string        `json:"skuSpec"`
			ItemsImg                  string        `json:"itemsImg"`
			ItemsThumbImg             string        `json:"itemsThumbImg"`
			ItemsVersion              int           `json:"itemsVersion"`
			ItemsState                int           `json:"itemsState"`
			ItemsIsPresale            int           `json:"itemsIsPresale"`
			ItemsIsOversea            int           `json:"itemsIsOversea"`
			TaxPrice                  float64       `json:"taxPrice"`
			SpikeStatus               int           `json:"spikeStatus"`
			SaleType                  int           `json:"saleType"`
			WarehouseName             interface{}   `json:"warehouseName"`
			WareHouseId               interface{}   `json:"wareHouseId"`
			SeckillLimit              interface{}   `json:"seckillLimit"`
			SecKillSoldStatus         interface{}   `json:"secKillSoldStatus"`
			ItemsInfoUrl              interface{}   `json:"itemsInfoUrl"`
			HasItemNum                interface{}   `json:"hasItemNum"`
			BlindBoxId                interface{}   `json:"blindBoxId"`
			PricePrefix               interface{}   `json:"pricePrefix"`
			PriceSymbol               string        `json:"priceSymbol"`
			PriceCyberMoney           interface{}   `json:"priceCyberMoney"`
			ShowContent               []interface{} `json:"showContent"`
			DiscountText              interface{}   `json:"discountText"`
			PresaleSupplyMoneyEndTime interface{}   `json:"presaleSupplyMoneyEndTime"`
			CartOrderType             int           `json:"cartOrderType"`
			PriceRestShowText         string        `json:"priceRestShowText"`
			PriceShowReal             float64       `json:"priceShowReal"`
			AmountShowReal            string        `json:"amountShowReal"`
			PresaleDeliveryTime       interface{}   `json:"presaleDeliveryTime"`
			CartItemsType             interface{}   `json:"cartItemsType"`
			EarlyBuyMaxDiscount       interface{}   `json:"earlyBuyMaxDiscount"`
			EarlyPromotion            interface{}   `json:"earlyPromotion"`
			EarlyPromotionFlag        bool          `json:"earlyPromotionFlag"`
			MoneyShows                interface{}   `json:"moneyShows"`
			Status                    interface{}   `json:"status"`
			SubType                   int           `json:"subType"`
			ResourceId                interface{}   `json:"resourceId"`
			ResourceType              interface{}   `json:"resourceType"`
			ExtraData                 interface{}   `json:"extraData"`
			DepositAmount             interface{}   `json:"depositAmount"`
			GoodsIsPromotionTag       int           `json:"goodsIsPromotionTag"`
			PromotionShowText         string        `json:"promotionShowText"`
		} `json:"validList"`
		InvalidList          []interface{} `json:"invalidList"`
		CartTotalMoney       float64       `json:"cartTotalMoney"`
		CartTotalMoneyAll    float64       `json:"cartTotalMoneyAll"`
		ItemsNumAll          int           `json:"itemsNumAll"`
		ItemsTotalMoneyAll   float64       `json:"itemsTotalMoneyAll"`
		ExpressTotalMoneyAll float64       `json:"expressTotalMoneyAll"`
		PayTotalMoneyAll     float64       `json:"payTotalMoneyAll"`
		BenifitTotalMoneyAll float64       `json:"benifitTotalMoneyAll"`
		PayInfo              struct {
			CustomerId      int    `json:"customerId"`
			ServiceType     int    `json:"serviceType"`
			OrderId         string `json:"orderId"`
			OrderCreateTime int64  `json:"orderCreateTime"`
			OrderExpire     int    `json:"orderExpire"`
			FeeType         string `json:"feeType"`
			PayAmount       int    `json:"payAmount"`
			OriginalAmount  int    `json:"originalAmount"`
			DeviceType      int    `json:"deviceType"`
			DeviceInfo      string `json:"deviceInfo"`
			NotifyUrl       string `json:"notifyUrl"`
			ProductId       string `json:"productId"`
			ProductUrl      string `json:"productUrl"`
			ShowTitle       string `json:"showTitle"`
			ShowContent     string `json:"showContent"`
			CreateIp        string `json:"createIp"`
			CreateUa        string `json:"createUa"`
			ReturnUrl       string `json:"returnUrl"`
			FailUrl         string `json:"failUrl"`
			ExtData         string `json:"extData"`
			TraceId         string `json:"traceId"`
			Timestamp       int64  `json:"timestamp"`
			Version         string `json:"version"`
			SignType        string `json:"signType"`
			Sign            string `json:"sign"`
			DefaultChoose   string `json:"defaultChoose"`
			Uid             int    `json:"uid"`
			GrayChannel     int    `json:"grayChannel"`
		} `json:"payInfo"`
		CodeType       int           `json:"codeType"`
		CodeMsg        interface{}   `json:"codeMsg"`
		CommonDialogVO interface{}   `json:"commonDialogVO"`
		ErrorList      []interface{} `json:"errorList"`
		ExtData        struct {
			FrontDiscount      interface{} `json:"frontDiscount"`
			FinalDiscount      interface{} `json:"finalDiscount"`
			BaotuangouDiscount interface{} `json:"baotuangouDiscount"`
			CouponTypes        interface{} `json:"couponTypes"`
			BookMoney          interface{} `json:"bookMoney"`
			BookDiscount       interface{} `json:"bookDiscount"`
			NotifyPhone        interface{} `json:"notifyPhone"`
			FrontMoney         interface{} `json:"frontMoney"`
			FinalMoney         interface{} `json:"finalMoney"`
			FrontOrderNo       interface{} `json:"frontOrderNo"`
			FinalOrderNo       interface{} `json:"finalOrderNo"`
		} `json:"extData"`
		CartOrderType         int           `json:"cartOrderType"`
		SubStatus             interface{}   `json:"subStatus"`
		ShowContent           []interface{} `json:"showContent"`
		CouponListIsShow      int           `json:"couponListIsShow"`
		CouponCodeId          string        `json:"couponCodeId"`
		CouponDesc            interface{}   `json:"couponDesc"`
		CouponCodeList        []interface{} `json:"couponCodeList"`
		DiscountTotalMoneyAll int           `json:"discountTotalMoneyAll"`
		DiscountList          []interface{} `json:"discountList"`
		OverseaIsShow         interface{}   `json:"overseaIsShow"`
		NotifyText            interface{}   `json:"notifyText"`
		TaxTotalMoneyAll      float64       `json:"taxTotalMoneyAll"`
		ItemsTaxTotalMoney    float64       `json:"itemsTaxTotalMoney"`
		ExpressTaxTotalMoney  float64       `json:"expressTaxTotalMoney"`
		MoneyList             []struct {
			NameDetail interface{} `json:"nameDetail"`
			Nums       interface{} `json:"nums"`
			Sort       interface{} `json:"sort"`
			Name       string      `json:"name"`
			Value      string      `json:"value"`
			Detail     interface{} `json:"detail"`
		} `json:"moneyList"`
		SecKill      interface{}   `json:"secKill"`
		SecKillList  []interface{} `json:"secKillList"`
		ActivityInfo struct {
			ActivityId           int         `json:"activityId"`
			Type                 int         `json:"type"`
			AdvancePublishTime   interface{} `json:"advancePublishTime"`
			AdvanceCountdownTime interface{} `json:"advanceCountdownTime"`
			MarketingId          interface{} `json:"marketingId"`
			SoldOut              interface{} `json:"soldOut"`
			MaxDiscount          interface{} `json:"maxDiscount"`
		} `json:"activityInfo"`
		OrderTitle        string      `json:"orderTitle"`
		ExpressTitle      string      `json:"expressTitle"`
		PayType           int         `json:"payType"`
		PayChannelId      interface{} `json:"payChannelId"`
		PayChannel        interface{} `json:"payChannel"`
		RealChannel       interface{} `json:"realChannel"`
		PayChannels       interface{} `json:"payChannels"`
		VerfyConf         interface{} `json:"verfyConf"`
		SubsidyVo         interface{} `json:"subsidyVo"`
		SubsidyIsShow     int         `json:"subsidyIsShow"`
		PromotionAreaVO   interface{} `json:"promotionAreaVO"`
		OrderCtime        int64       `json:"orderCtime"`
		UseIsEneygy       interface{} `json:"useIsEneygy"`
		EnergyConsumption interface{} `json:"energyConsumption"`
		FreightCouponVo   interface{} `json:"freightCouponVo"`
	} `json:"data"`
	Errtag int `json:"errtag"`
}

type GetBlindBoxInfoResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ItemsId     int         `json:"itemsId"`
		SkuId       int         `json:"skuId"`
		Name        string      `json:"name"`
		NameLeftTag interface{} `json:"nameLeftTag"`
		NameUserTag interface{} `json:"nameUserTag"`
		NameTagType interface{} `json:"nameTagType"`
		TipBoxInfo  struct {
			TipImg     string `json:"tipImg"`
			TipDesc    string `json:"tipDesc"`
			TipJumpUrl string `json:"tipJumpUrl"`
		} `json:"tipBoxInfo"`
		ItemsType  int `json:"itemsType"`
		SubSkuList []struct {
			ImageUrl                string      `json:"imageUrl"`
			Type                    int         `json:"type"`
			Name                    string      `json:"name"`
			SubSkuId                int         `json:"subSkuId"`
			SaleStatus              interface{} `json:"saleStatus"`
			WishedSku               bool        `json:"wishedSku"`
			ItemsType               int         `json:"itemsType"`
			PresaleDeliveryTimeStr  *string     `json:"presaleDeliveryTimeStr"`
			OverseasPreShippingTime *string     `json:"overseasPreShippingTime"`
			WhiteListSku            interface{} `json:"whiteListSku"`
			SubSkuPrice             string      `json:"subSkuPrice"`
			SubSkuItemsId           int         `json:"subSkuItemsId"`
			Url                     *string     `json:"url"`
			SliceNum                interface{} `json:"sliceNum"`
		} `json:"subSkuList"`
		Price                float64       `json:"price"`
		DiscountedPrice      interface{}   `json:"discountedPrice"`
		MultiDiscountedPrice interface{}   `json:"multiDiscountedPrice"`
		BlindBoxCoin         float64       `json:"blindBoxCoin"`
		ServerTime           interface{}   `json:"serverTime"`
		Sales                interface{}   `json:"sales"`
		MobileDesc           string        `json:"mobileDesc"`
		PcDesc               string        `json:"pcDesc"`
		Status               int           `json:"status"`
		AutoOnSaleTime       int           `json:"autoOnSaleTime"`
		ItemsStatus          int           `json:"itemsStatus"`
		SaleStatus           int           `json:"saleStatus"`
		AttrList             []interface{} `json:"attrList"`
		WholeBoxVO           interface{}   `json:"wholeBoxVO"`
		UserBlindBoxCoin     interface{}   `json:"userBlindBoxCoin"`
		PayType              int           `json:"payType"`
		MultiNum             int           `json:"multiNum"`
		IsStockSufficient    bool          `json:"isStockSufficient"`
		HasRecharge          bool          `json:"hasRecharge"`
		IsStoneOffline       bool          `json:"isStoneOffline"`
		BlindBoxWish         struct {
			HasWishedItems bool   `json:"hasWishedItems"`
			JumpUrl        string `json:"jumpUrl"`
		} `json:"blindBoxWish"`
		NeedLogin           bool        `json:"needLogin"`
		ActivityInfoVO      interface{} `json:"activityInfoVO"`
		MojinActivityInfoVO interface{} `json:"mojinActivityInfoVO"`
		EuroActivityInfoVO  struct {
			TagName  interface{} `json:"tagName"`
			Desc     interface{} `json:"desc"`
			JumpUrl  interface{} `json:"jumpUrl"`
			IsEnable bool        `json:"isEnable"`
		} `json:"euroActivityInfoVO"`
		PresaleDeliveryTime   string `json:"presaleDeliveryTime"`
		AssociatedTaoBaoSpuId string `json:"associatedTaoBaoSpuId"`
		DrawSliceInfo         struct {
			Status           int         `json:"status"`
			Slice            int         `json:"slice"`
			IsLastSlice      interface{} `json:"isLastSlice"`
			Total            interface{} `json:"total"`
			ProbabilityInfos []struct {
				Type        int    `json:"type"`
				Probability string `json:"probability"`
			} `json:"probabilityInfos"`
		} `json:"drawSliceInfo"`
		NewCustomer    bool        `json:"newCustomer"`
		ThemeId        int         `json:"themeId"`
		PreType        interface{} `json:"preType"`
		NewUserLimit   interface{} `json:"newUserLimit"`
		IsInWhiteList  interface{} `json:"isInWhiteList"`
		ProtocolConfig struct {
			ImageUrl     string `json:"image_url"`
			SceneId      int    `json:"scene_id"`
			Id           int    `json:"id"`
			NeedRecheck  int    `json:"need_recheck"`
			Mtime        string `json:"mtime"`
			Url          string `json:"url"`
			ProtocolName string `json:"protocol_name"`
		} `json:"protocolConfig"`
		IllustrationActivityInfoVO interface{} `json:"illustrationActivityInfoVO"`
		UserTaskGuideVO            struct {
			Uid            interface{} `json:"uid"`
			HasActivity    bool        `json:"hasActivity"`
			ActivityId     string      `json:"activityId"`
			ActivityName   string      `json:"activityName"`
			SubTitle       interface{} `json:"subTitle"`
			UserTaskStatus int         `json:"userTaskStatus"`
			WaitReceiveNum int         `json:"waitReceiveNum"`
			HomeLink       interface{} `json:"homeLink"`
			GuideTasks     []struct {
				TaskName string `json:"taskName"`
				Prize    struct {
					PassIndex    interface{} `json:"passIndex"`
					PrizeName    string      `json:"prizeName"`
					PrizeImg     string      `json:"prizeImg"`
					PrizeNum     int         `json:"prizeNum"`
					PrizeType    interface{} `json:"prizeType"`
					RightId      interface{} `json:"rightId"`
					RightType    interface{} `json:"rightType"`
					RandomType   interface{} `json:"randomType"`
					IsMust       interface{} `json:"isMust"`
					RewardStatus interface{} `json:"rewardStatus"`
					RewardId     interface{} `json:"rewardId"`
					FaceValue    interface{} `json:"faceValue"`
					AmountType   interface{} `json:"amountType"`
				} `json:"prize"`
			} `json:"guideTasks"`
		} `json:"userTaskGuideVO"`
		RecyclableStatus int `json:"recyclableStatus"`
		ActivityLegoData struct {
			GeneralProgress   interface{} `json:"generalProgress"`
			MlsPmActivityInfo struct {
				MlsPmEntranceInfo struct {
					EntranceName   string `json:"entranceName"`
					WaitReceiveNum int    `json:"waitReceiveNum"`
					SignStatus     int    `json:"signStatus"`
					SignText       string `json:"signText"`
					Mid            int    `json:"mid"`
					JumpUrl        string `json:"jumpUrl"`
					IconUrl        string `json:"iconUrl"`
				} `json:"mlsPmEntranceInfo"`
				TaskGuideRes struct {
					Uid            interface{} `json:"uid"`
					HasActivity    bool        `json:"hasActivity"`
					ActivityId     string      `json:"activityId"`
					ActivityName   string      `json:"activityName"`
					SubTitle       interface{} `json:"subTitle"`
					UserTaskStatus int         `json:"userTaskStatus"`
					WaitReceiveNum int         `json:"waitReceiveNum"`
					HomeLink       interface{} `json:"homeLink"`
					GuideTasks     []struct {
						TaskName string `json:"taskName"`
						Prize    struct {
							PassIndex    interface{} `json:"passIndex"`
							PrizeName    string      `json:"prizeName"`
							PrizeImg     string      `json:"prizeImg"`
							PrizeNum     int         `json:"prizeNum"`
							PrizeType    interface{} `json:"prizeType"`
							RightId      interface{} `json:"rightId"`
							RightType    interface{} `json:"rightType"`
							RandomType   interface{} `json:"randomType"`
							IsMust       interface{} `json:"isMust"`
							RewardStatus interface{} `json:"rewardStatus"`
							RewardId     interface{} `json:"rewardId"`
							FaceValue    interface{} `json:"faceValue"`
							AmountType   interface{} `json:"amountType"`
						} `json:"prize"`
					} `json:"guideTasks"`
				} `json:"taskGuideRes"`
			} `json:"mlsPmActivityInfo"`
			MagicBubbleDetailPageInfo struct {
				ActivityInfoList []struct {
					MagicBubbleTaskType       int         `json:"magicBubbleTaskType"`
					TaskGroupType             int         `json:"taskGroupType"`
					MainTitle                 string      `json:"mainTitle"`
					SubTitle                  string      `json:"subTitle"`
					ActivityStatusText        string      `json:"activityStatusText"`
					ActivityStatusSubText     string      `json:"activityStatusSubText"`
					ActivityStatusPrizeText   string      `json:"activityStatusPrizeText"`
					ActivityStatusImg         interface{} `json:"activityStatusImg"`
					GuideLink                 interface{} `json:"guideLink"`
					GuideButtonText           string      `json:"guideButtonText"`
					ActivityImg               string      `json:"activityImg"`
					PicIcon                   interface{} `json:"picIcon"`
					ActivityStatus            int         `json:"activityStatus"`
					BubbleNextGradePrefixText interface{} `json:"bubbleNextGradePrefixText"`
					BubbleNextGradePrizeText  interface{} `json:"bubbleNextGradePrizeText"`
					ReceivePrizeId            interface{} `json:"receivePrizeId"`
					ReceivePrizeIdList        interface{} `json:"receivePrizeIdList"`
					PrizeImg                  string      `json:"prizeImg"`
					FinalPrize                bool        `json:"finalPrize"`
				} `json:"activityInfoList"`
				ExistAwardingActivity bool   `json:"existAwardingActivity"`
				JumpUrl               string `json:"jumpUrl"`
				RotationInfoList      []struct {
					MainTitle                 string      `json:"mainTitle"`
					SubTitle                  string      `json:"subTitle"`
					PreText                   string      `json:"preText"`
					PrizeText                 string      `json:"prizeText"`
					Priority                  int         `json:"priority"`
					GuideLink                 interface{} `json:"guideLink"`
					GuideButtonText           string      `json:"guideButtonText"`
					ActivityStatus            int         `json:"activityStatus"`
					MagicBubbleTaskType       int         `json:"magicBubbleTaskType"`
					ReceivePrizeIds           interface{} `json:"receivePrizeIds"`
					PrizeImg                  string      `json:"prizeImg"`
					BubbleNextGradePrefixText interface{} `json:"bubbleNextGradePrefixText"`
					BubbleNextGradePrizeText  interface{} `json:"bubbleNextGradePrizeText"`
				} `json:"rotationInfoList"`
			} `json:"magicBubbleDetailPageInfo"`
		} `json:"activityLegoData"`
		DeliveryNotices interface{} `json:"deliveryNotices"`
		PageType        string      `json:"pageType"`
		DetailGuideInfo interface{} `json:"detailGuideInfo"`
	} `json:"data"`
	Errtag int `json:"errtag"`
}
