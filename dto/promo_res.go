package dto

type PromoResponse struct {
	Status  bool        `json:"status"`
	Title   interface{} `json:"title"`
	Message string      `json:"message"`
	Code    interface{} `json:"code"`
	Data    struct {
		ID                  int         `json:"id"`
		ID2                 int         `json:"id2"`
		Name                string      `json:"name"`
		TagDetail           interface{} `json:"tagDetail"`
		Child               interface{} `json:"child"`
		Metadata            interface{} `json:"metadata"`
		LayoutID            interface{} `json:"layoutId"`
		MoreID              interface{} `json:"moreId"`
		IsCard              interface{} `json:"isCard"`
		IsAutorotate        interface{} `json:"isAutorotate"`
		TimePeriod          interface{} `json:"timePeriod"`
		Survey              interface{} `json:"survey"`
		MainCategoryID      int         `json:"mainCategoryId"`
		FlagType            interface{} `json:"flagType"`
		FlagURL             interface{} `json:"flagUrl"`
		Color1              string      `json:"color1"`
		Color2              string      `json:"color2"`
		Color3              string      `json:"color3"`
		Image               string      `json:"image"`
		TeaserImage         interface{} `json:"teaserImage"`
		TeaserStartDateTime interface{} `json:"teaserStartDateTime"`
		TeaserEndDateTime   interface{} `json:"teaserEndDateTime"`
		PromoStartDateTime  string      `json:"promoStartDateTime"`
		PromoEndDateTime    string      `json:"promoEndDateTime"`
		PromoActiveFlag     bool        `json:"promoActiveFlag"`
		PromoType           string      `json:"promoType"`
		ProductList         []struct {
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
				IMGSMALL              string `json:"IMG_SMALL"`
				IMGRFUBIG             string `json:"IMG_RFU_BIG"`
				IMGPRODUCTSQUAREBIG   string `json:"IMG_PRODUCT_SQUARE_BIG"`
				IMGPRODUCTPORTRAIT    string `json:"IMG_PRODUCT_PORTRAIT"`
				IMGRFU                string `json:"IMG_RFU"`
				IMGBREAKERBIG         string `json:"IMG_BREAKER_BIG"`
				IMGSMALLBIG           string `json:"IMG_SMALL_BIG"`
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
			StartDateTime                   string      `json:"startDateTime"`
			EndDateTime                     string      `json:"endDateTime"`
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
		} `json:"productList"`
	} `json:"data"`
}

type PromoErrorResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
