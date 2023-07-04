package models

import (
	"github.com/amerta-teknologi/az-finance/config"
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id             int    `json:"id"`
	ProductId      int    `json:"parent_id"`
	UserId         int    `json:"user_id"`
	TypeId         int    `json:"type_id"`
	Description    string `json:"description"`
	Amount         int    `json:"amount"`
	Balance        int    `json:"balance"`
	GainPercentage int    `json:"gain_percentage"`
	CreatedAt      string `json:"created_at"`
	Type           Type   `json:"type"`
}

type TransactionsResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []Transaction `json:"data"`
}

type TransactionResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    Transaction `json:"data"`
}

func (model *Transaction) FindAll(c *gin.Context) *TransactionsResponse {
	var queries string = "?"

	if c.Param("product-slug") != "" {
		queries = queries + "slug=" + c.Param("product-slug") + "&"
	}

	if c.Query("page") == "" {
		queries = queries + "page=1&"
	}

	if c.Query("per_page") == "" {
		queries = queries + "per_page=12&"
	}

	queries = queries + removeEmptyQueries(c.Request.URL.Query())

	response := Get(config.Data.WhaleAddress + "/api/v1/catalog/transactions" + queries)
	body := GetBody(response, &TransactionsResponse{})

	return body.(*TransactionsResponse)
}
