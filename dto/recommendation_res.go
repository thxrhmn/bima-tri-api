package dto

type RecommendationResponse struct {
	Status   bool `json:"status"`
	Products []struct {
		ProductID            string      `json:"productId"`
		ProductName          string      `json:"productName"`
		ProductPrice         string      `json:"productPrice"`
		ProductValidity      interface{} `json:"productValidity"`
		ProductOriginalPrice string      `json:"productOriginalPrice"`
		ProductDescription   string      `json:"productDescription"`
		ProductOtherInfo     interface{} `json:"productOtherInfo"`
		ProductHowTo         string      `json:"productHowTo"`
		ProductPricing       string      `json:"productPricing"`
		ProductDetailsType   string      `json:"productDetailsType"`
		ProductDetails       struct {
		} `json:"productDetails"`
		LabelDescription       string      `json:"labelDescription"`
		LabelHowTo             string      `json:"labelHowTo"`
		LabelPricing           string      `json:"labelPricing"`
		LabelDetails           string      `json:"labelDetails"`
		LabelOtherInfo         string      `json:"labelOtherInfo"`
		ButtonBuy              string      `json:"buttonBuy"`
		BonstriPoints          interface{} `json:"bonstriPoints"`
		ProductRating          float64     `json:"productRating"`
		ProductRatingTotalUser int         `json:"productRatingTotalUser"`
		CategoryID             interface{} `json:"categoryId"`
		PaymentMatrix          []int       `json:"paymentMatrix"`
		IsRedirectToLink       bool        `json:"isRedirectToLink"`
		RedirectLink           string      `json:"redirectLink"`
		IsDownloadLink         bool        `json:"isDownloadLink"`
		DownloadLink           string      `json:"downloadLink"`
		Metadata               struct {
			IMGBANNERBIG          string `json:"IMG_BANNER_BIG"`
			IMGPRODUCTSQUARE      string `json:"IMG_PRODUCT_SQUARE"`
			IMGRFUBIG             string `json:"IMG_RFU_BIG"`
			IMGPRODUCTSQUAREBIG   string `json:"IMG_PRODUCT_SQUARE_BIG"`
			IMGPRODUCTPORTRAIT    string `json:"IMG_PRODUCT_PORTRAIT"`
			IMGRFU                string `json:"IMG_RFU"`
			IMGBREAKERBIG         string `json:"IMG_BREAKER_BIG"`
			IMGBANNER             string `json:"IMG_BANNER"`
			IMGBREAKER            string `json:"IMG_BREAKER"`
			IMGPRODUCTPORTRAITBIG string `json:"IMG_PRODUCT_PORTRAIT_BIG"`
		} `json:"metadata"`
		VendorList []struct {
			VendorID   int    `json:"vendorId"`
			VendorName string `json:"vendorName"`
			PriceList  []struct {
				PlanName string `json:"planName"`
				Price    string `json:"price"`
			} `json:"priceList"`
		} `json:"vendorList"`
		PaymentList []struct {
			MethodCode string `json:"methodCode"`
			MethodName string `json:"methodName"`
			Rest       bool   `json:"rest"`
		} `json:"paymentList"`
		IsInappBrowser                  interface{} `json:"isInappBrowser"`
		IsTransferable                  int         `json:"isTransferable"`
		IsShareable                     int         `json:"isShareable"`
		IsAutoRenewable                 int         `json:"isAutoRenewable"`
		Deeplink                        interface{} `json:"deeplink"`
		Campaign                        int         `json:"campaign"`
		TagDetail                       interface{} `json:"tagDetail"`
		FlagURL                         interface{} `json:"flagUrl"`
		FlagType                        interface{} `json:"flagType"`
		IsArticle                       bool        `json:"isArticle"`
		IsBanner                        bool        `json:"isBanner"`
		SubMenuCategoryID               interface{} `json:"subMenuCategoryId"`
		ArticleTitle                    interface{} `json:"articleTitle"`
		ArticleBody                     interface{} `json:"articleBody"`
		ArticleLayout                   interface{} `json:"articleLayout"`
		IsSquareBanner                  interface{} `json:"isSquareBanner"`
		ProductLayout                   interface{} `json:"productLayout"`
		IsInstallment                   interface{} `json:"isInstallment"`
		IsRedemption                    bool        `json:"isRedemption"`
		BasePackage                     interface{} `json:"basePackage"`
		PartnerID                       interface{} `json:"partnerId"`
		IsProductRedirection            bool        `json:"isProductRedirection"`
		IsGamePartner                   bool        `json:"isGamePartner"`
		IsActive                        bool        `json:"isActive"`
		StartDateTime                   interface{} `json:"startDateTime"`
		EndDateTime                     interface{} `json:"endDateTime"`
		OfferDetail                     interface{} `json:"offerDetail"`
		GameProductCategory             interface{} `json:"gameProductCategory"`
		GamePurchaseType                interface{} `json:"gamePurchaseType"`
		ParentLayoutID                  interface{} `json:"parentLayoutId"`
		CustomData                      interface{} `json:"customData"`
		CustomDataID                    interface{} `json:"customDataId"`
		CustomDataQuota                 interface{} `json:"customDataQuota"`
		Zone                            interface{} `json:"zone"`
		Active                          bool        `json:"active"`
		IsRequireGameIDVerification     bool        `json:"isRequireGameIdVerification"`
		IsRequireGameZoneIDVerification bool        `json:"isRequireGameZoneIdVerification"`
		Parent                          bool        `json:"parent"`
	} `json:"products"`
}

type RecommendationErrorResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
