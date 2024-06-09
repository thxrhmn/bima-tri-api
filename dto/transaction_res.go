package dto

type TransactionResponse struct {
	Status  bool        `json:"status"`
	Title   interface{} `json:"title"`
	Message string      `json:"message"`
	Code    interface{} `json:"code"`
	Data    []struct {
		TransactionID   string      `json:"transactionId"`
		DateReg         string      `json:"dateReg"`
		Name            string      `json:"name"`
		Price           string      `json:"price"`
		ProductID       string      `json:"productId"`
		Category        string      `json:"category"`
		Status          string      `json:"status"`
		Icon            string      `json:"icon"`
		Addon           interface{} `json:"addon"`
		IsAddon         bool        `json:"isAddon"`
		Target          interface{} `json:"target"`
		ParentProductID string      `json:"parentProductId"`
	} `json:"data"`
}

type TransactionErrorResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
