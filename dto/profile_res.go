package dto

type ProfileResponse struct {
	Status  bool        `json:"status"`
	Title   interface{} `json:"title"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    struct {
		Status                  bool          `json:"status"`
		Msisdn                  string        `json:"msisdn"`
		Balance                 string        `json:"balance"`
		CreditLimit             string        `json:"creditLimit"`
		CreditLimitHybrid       string        `json:"creditLimitHybrid"`
		Validity                string        `json:"validity"`
		BalanceTrims            string        `json:"balanceTrims"`
		StatusSubscription      bool          `json:"statusSubscription"`
		SubscriberType          string        `json:"subscriberType"`
		Email                   string        `json:"email"`
		EmailHybrid             string        `json:"emailHybrid"`
		PackageProductList      []interface{} `json:"packageProductList"`
		PackageProductGroupList []struct {
			IsTransferable      bool          `json:"isTransferable"`
			IsShareable         bool          `json:"isShareable"`
			IsStoppable         bool          `json:"isStoppable"`
			ProductID           interface{}   `json:"productId"`
			ShareQuota          []interface{} `json:"shareQuota"`
			MaxChildCount       int           `json:"maxChildCount"`
			CurrChildCount      int           `json:"currChildCount"`
			ListOfSharableQuota []interface{} `json:"listOfSharableQuota"`
			Name                string        `json:"name"`
			Content             []struct {
				Title        string      `json:"title"`
				Validity     string      `json:"validity"`
				ValidityRaw  string      `json:"validityRaw"`
				Consumed     string      `json:"consumed"`
				ConsumedRaw  string      `json:"consumedRaw"`
				Allocated    string      `json:"allocated"`
				AllocatedRaw string      `json:"allocatedRaw"`
				Remaining    string      `json:"remaining"`
				RemainingRaw string      `json:"remainingRaw"`
				BonusType    string      `json:"bonusType"`
				IsShareable  bool        `json:"isShareable"`
				ProductID    interface{} `json:"productId"`
				LocalQuota   bool        `json:"localQuota"`
			} `json:"content"`
		} `json:"packageProductGroupList"`
		PackageList []struct {
			IsRenewable         bool        `json:"isRenewable"`
			IsProductBuy        bool        `json:"isProductBuy"`
			IsMore              bool        `json:"isMore"`
			IsShareQuota        bool        `json:"isShareQuota"`
			Name                string      `json:"name"`
			LinkProductBuyType  string      `json:"linkProductBuyType"`
			LinkProductBuyValue string      `json:"linkProductBuyValue"`
			ProductID           interface{} `json:"productId"`
			Detail              []struct {
				Validity      string      `json:"validity"`
				Value         string      `json:"value"`
				ValueOriginal interface{} `json:"valueOriginal"`
			} `json:"detail"`
			GroupID string `json:"groupId"`
		} `json:"packageList"`
		DueDateHybrid        string `json:"dueDateHybrid"`
		DueDate              string `json:"dueDate"`
		BalanceOnNet         string `json:"balanceOnNet"`
		BalanceOffNet        string `json:"balanceOffNet"`
		BalanceTotal         string `json:"balanceTotal"`
		BalanceOnAndOffNet   string `json:"balanceOnAndOffNet"`
		ProfilePicture       string `json:"profilePicture"`
		CallPlan             string `json:"callPlan"`
		ProfileColor         string `json:"profileColor"`
		ProfileTime          int    `json:"profileTime"`
		CustomerSegmentation struct {
			SegmentationName string      `json:"segmentationName"`
			Title            string      `json:"title"`
			Description      string      `json:"description"`
			Deeplink         interface{} `json:"deeplink"`
		} `json:"customerSegmentation"`
		ProfileName    string      `json:"profileName"`
		ProfileHobby   interface{} `json:"profileHobby"`
		DukcapilName   string      `json:"dukcapilName"`
		DukcapilURL    string      `json:"dukcapilUrl"`
		Cashback       string      `json:"cashback"`
		BalanceDetails struct {
			Balance struct {
				Title string `json:"title"`
				Value string `json:"value"`
			} `json:"balance"`
			BalanceOnNet       interface{} `json:"balanceOnNet"`
			BalanceOffnet      interface{} `json:"balanceOffnet"`
			BalanceOnAndOffNet struct {
				Detail struct {
					MasaBerlaku string `json:"Masa Berlaku"`
					SisaPulsa   string `json:"Sisa Pulsa"`
				} `json:"detail"`
				Title string `json:"title"`
				Value string `json:"value"`
			} `json:"balanceOnAndOffNet"`
		} `json:"balanceDetails"`
		FirstLogin   string `json:"firstLogin"`
		LastLogin    string `json:"lastLogin"`
		SumOfBonuses struct {
			Data  string `json:"Data"`
			Voice string `json:"Voice"`
			SMS   string `json:"SMS"`
		} `json:"sumOfBonuses"`
		MaxPullNotification int         `json:"maxPullNotification"`
		BirthDate           interface{} `json:"birthDate"`
		ActivePackages      interface{} `json:"activePackages"`
		BonstriDetail       struct {
			TotalPoint           int `json:"totalPoint"`
			LoyaltyPoint         int `json:"loyaltyPoint"`
			ValidityLoyaltyPoint int `json:"validityLoyaltyPoint"`
			BonusPoint           int `json:"bonusPoint"`
		} `json:"bonstriDetail"`
		MgmReferralDetails struct {
			CountOfReferrals int    `json:"countOfReferrals"`
			TotalPoints      int    `json:"totalPoints"`
			ReferralCode     string `json:"referralCode"`
			TncURL           string `json:"tncURL"`
		} `json:"mgmReferralDetails"`
	} `json:"data"`
}

type ProfileErrorResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
