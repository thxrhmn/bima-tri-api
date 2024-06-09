package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"thxrhmn-bimatri-api/bima"
	"thxrhmn-bimatri-api/config"
	"thxrhmn-bimatri-api/docs"
	"thxrhmn-bimatri-api/dto"

	"github.com/common-nighthawk/go-figure"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// swagger:meta

// @title Bima TRI API
// @version 1.0
// @description This is bima tri api
// @termsOfService https://google.com/terms/

// @contact.name API Support
// @contact.url https://www.google.com/contact
// @contact.email support@google.com
// @BasePath /

// Ping godoc
// @Summary Ping
// @Schemes
// @Description do ping
// @Tags v1
// @Accept json
// @Produce json
// @Success 200 {string} Ping
// @Router / [get]
func Ping(g *gin.Context) {
	response := map[string]interface{}{"status": 200, "message": "hello world!"}
	g.JSON(http.StatusOK, response)
}

// Login godoc
// @Summary Login
// @Schemes
// @Description Login request OTP
// @Tags v1
// @Accept multipart/form-data
// @Produce json
// @Param msisdn formData string true "MSISDN"
// @Success 200 {object} dto.LoginResponse
// @Router /auth/login [post]
func Login(g *gin.Context) {
	msisdn := g.PostForm("msisdn")
	if msisdn == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "msisdn is required"})
		return
	}

	responseData := bima.LoginOtpRequest(config.EndpointInfo{
		URL:    config.Endpoints.LoginOtpRequestRequest,
		Method: "POST",
	}, msisdn)

	var loginResponse dto.LoginResponse
	if err := mapToStruct(responseData, &loginResponse); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
		return
	}

	g.JSON(http.StatusOK, loginResponse)
}

// Login With Otp godoc
// @Summary Login With OTP
// @Schemes
// @Description Login with OTP
// @Tags v1
// @Accept multipart/form-data
// @Produce json
// @Param msisdn formData string true "MSISDN"
// @Param otp formData string true "OTP"
// @Success 200 {object} dto.LoginWithOtpResponse
// @Failure 400 {object} dto.LoginWithOtpErrorResponse
// @Router /auth/login-otp [post]
func LoginWithOtp(g *gin.Context) {
	var loginWithOtpErrorResponse dto.LoginWithOtpErrorResponse

	msisdn := g.PostForm("msisdn")
	otp := g.PostForm("otp")

	if msisdn == "" {
		errorMsisdn := dto.LoginWithOtpErrorResponse{
			Status:  false,
			Message: "msisdn is required",
		}
		g.JSON(http.StatusBadRequest, errorMsisdn)
		return
	}

	if otp == "" {
		errorOtp := dto.LoginWithOtpErrorResponse{
			Status:  false,
			Message: "otp is required",
		}
		g.JSON(http.StatusBadRequest, errorOtp)
		return
	}

	responseData := bima.LoginWithOtp(config.EndpointInfo{
		URL:    config.Endpoints.LoginWithOtp,
		Method: "POST",
	}, msisdn, otp)

	if responseData["status"] == false {
		if err := mapToStruct(responseData, &loginWithOtpErrorResponse); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
			return
		}
		g.JSON(http.StatusBadRequest, loginWithOtpErrorResponse)
		return
	}

	var loginWithOtpResponse dto.LoginWithOtpResponse
	if err := mapToStruct(responseData, &loginWithOtpResponse); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response"})
		return
	}

	g.JSON(http.StatusOK, loginWithOtpResponse)
}

// Profile godoc
// @Summary Profile
// @Description Profile API endpoint to retrieve user profile information.
// @Tags v1
// @Accept json
// @Produce json
// @Param profileRequest body dto.ProfileRequest true "Profile Request"
// @Success 200 {object} dto.ProfileResponse "Successfully retrieved profile information."
// @Failure 400 {object} dto.ProfileErrorResponse "Invalid request or authentication failure."
// @Router /profile [post]
func Profile(g *gin.Context) {
	var profileReq dto.AllRequest
	var profileResponse dto.ProfileResponse

	if err := g.ShouldBindJSON(&profileReq); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	endpointInfo := config.EndpointInfo{
		URL:    config.Endpoints.ProfileAccount,
		Method: "POST",
	}

	responseData, err := bima.Get(endpointInfo, profileReq)

	if err != nil {
		if errStr := err.Error(); strings.Contains(errStr, "unauthorized:") {
			errResponse := dto.ProfileErrorResponse{
				Status:  false,
				Message: err.Error(),
			}
			g.JSON(http.StatusUnauthorized, errResponse)
			return
		}

		errResponse := dto.ProfileErrorResponse{
			Status:  false,
			Message: err.Error(),
		}
		g.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	if err := mapToStruct(responseData, &profileResponse); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	g.JSON(http.StatusOK, profileResponse)
}

// Promo godoc
// @Summary Promo
// @Description Promo API endpoint to retrieve promo information.
// @Tags v1
// @Accept json
// @Produce json
// @Param promoRequest body dto.AllRequest true "Promo Request"
// @Success 200 {object} dto.PromoResponse "Successfully retrieved promo information."
// @Failure 400 {object} dto.PromoErrorResponse "Invalid request or authentication failure."
// @Router /promo [post]
func Promo(g *gin.Context) {
	var promoReq dto.AllRequest
	var promoResponse dto.PromoResponse

	if err := g.ShouldBindJSON(&promoReq); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	endpointInfo := config.EndpointInfo{
		URL:    config.Endpoints.Promo,
		Method: "POST",
	}

	responseData, err := bima.Get(endpointInfo, promoReq)

	if err != nil {
		if errStr := err.Error(); strings.Contains(errStr, "unauthorized:") {
			errResponse := dto.PromoErrorResponse{
				Status:  false,
				Message: err.Error(),
			}
			g.JSON(http.StatusUnauthorized, errResponse)
			return
		}

		errResponse := dto.PromoErrorResponse{
			Status:  false,
			Message: err.Error(),
		}
		g.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	if err := mapToStruct(responseData, &promoResponse); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	g.JSON(http.StatusOK, promoResponse)
}

// Transaction godoc
// @Summary Transaction
// @Description Transaction API endpoint to retrieve transaction information.
// @Tags v1
// @Accept json
// @Produce json
// @Param category query string false "Transaction Category" Enums(PULSA, PAKET, TRANSFER, HIBURAN, GAME) default(PULSA)
// @Param status query string false "Transaction Status" Enums(ALL, SUCCESS, PENDING, FAILED) default(ALL)
// @Param transactionRequest body dto.AllRequest true "Transaction Request"
// @Success 200 {object} dto.TransactionResponse "Successfully retrieved transaction information."
// @Failure 400 {object} dto.TransactionErrorResponse "Invalid request or authentication failure."
// @Router /transaction [post]
func Transaction(g *gin.Context) {
	var transactionReq dto.AllRequest
	var transactionRes dto.TransactionResponse

	category := g.DefaultQuery("category", "PULSA")
	status := g.DefaultQuery("status", "ALL")

	var endpointURL string
	switch category {
	case "PULSA":
		endpointURL = config.Endpoints.TransactionHistoryPulsa + status
	case "PAKET":
		endpointURL = config.Endpoints.TransactionHistoryPaket + status
	case "TRANSFER":
		endpointURL = config.Endpoints.TransactionHistoryTransfer + status
	case "HIBURAN":
		endpointURL = config.Endpoints.TransactionHistoryHiburan + status
	case "GAME":
		endpointURL = config.Endpoints.TransactionHistoryGames + status
	default:
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
		return
	}

	if err := g.ShouldBindJSON(&transactionReq); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	endpointInfo := config.EndpointInfo{
		URL:    endpointURL,
		Method: "POST",
	}

	responseData, err := bima.Get(endpointInfo, transactionReq)

	if err != nil {
		if strings.Contains(err.Error(), "unauthorized:") {
			errResponse := dto.TransactionErrorResponse{
				Status:  false,
				Message: err.Error(),
			}
			g.JSON(http.StatusUnauthorized, errResponse)
			return
		}

		errResponse := dto.TransactionErrorResponse{
			Status:  false,
			Message: err.Error(),
		}
		g.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	if err := mapToStruct(responseData, &transactionRes); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	g.JSON(http.StatusOK, transactionRes)
}

func mapToStruct(data map[string]interface{}, result interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, result)
}

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config.Init()

	// Logging to a file.
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.SetMode(gin.ReleaseMode)

	figure := figure.NewFigure("Bima TRI API", "smslant", false)
	figure.Print()
	fmt.Print("\nCreated by: @thxrhmn")
	fmt.Print("\nCreated at: 2024/06/09\n\n")

	SERVER_URL := viper.GetString("SERVER_URL")
	SERVER_PORT := viper.GetString("SERVER_PORT")

	if SERVER_URL == "" {
		SERVER_URL = "http://localhost:8787"
	}

	if SERVER_PORT == "" {
		SERVER_PORT = ":8787"
	}

	r := gin.Default()

	// Setup CORS
	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = []string{"*"} // Use the same origin as your Swagger UI
	cfg.AllowCredentials = true
	cfg.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	cfg.AllowHeaders = []string{"*"}

	r.Use(cors.New(cfg))

	docs.SwaggerInfo.BasePath = "/api/v1/"

	swaggerUrl := ginSwagger.URL(SERVER_URL + "/swagger/doc.json")

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/", Ping)
		apiV1.POST("/auth/login", Login)
		apiV1.POST("/auth/login-otp", LoginWithOtp)
		apiV1.POST("/profile", Profile)
		apiV1.POST("/promo", Promo)
		apiV1.POST("/transaction", Transaction)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, swaggerUrl))

	r.Run(SERVER_PORT)
}
