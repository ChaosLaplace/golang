package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ExchangeRates struct {
	Currencies map[string]map[string]float64 `json:"currencies"`
}

type ConversionResponse struct {
	Msg    string `json:"msg"`
	Amount string `json:"amount"`
}

func AsiaYo(c *gin.Context) {
	source := c.Query("source")
	target := c.Query("target")
	amountStr := c.Query("amount")

	if source == "" || target == "" || amountStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
		return
	}
	// 移除符號
	amountStr = strings.ReplaceAll(amountStr, "$", "")
	amountStr = strings.ReplaceAll(amountStr, ",", "")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount = " + amountStr})
		return
	}

	rates := ExchangeRates{
		Currencies: map[string]map[string]float64{
			"TWD": {
				"TWD": 1,
				"JPY": 3.669,
				"USD": 0.03281,
			},
			"JPY": {
				"TWD": 0.26956,
				"JPY": 1,
				"USD": 0.00885,
			},
			"USD": {
				"TWD": 30.444,
				"JPY": 111.801,
				"USD": 1,
			},
		},
	}

	rate := rates.Currencies[source][target]
	convertedAmount := amount * rate
	// 四捨五入
	amountTemp := math.Round(convertedAmount*100) / 100

	response := ConversionResponse{
		Msg:    "success",
		Amount: fmt.Sprintf("$%.2f", amountTemp),
	}

	c.JSON(http.StatusOK, response)
}
