package models

import (
	"github.com/amerta-teknologi/az-finance/config"
	"github.com/gin-gonic/gin"
)

type Projection struct {
	Id         int     `json:"id"`
	UserId     int     `json:"user_id"`
	Projection float32 `json:"projection"`
	Actual     float32 `json:"actual"`
	Year       int     `json:"year"`
	Month      int     `json:"month"`
}

type ProjectionsResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    []Projection `json:"data"`
}

func (model *Projection) FindAll(c *gin.Context) *ProjectionsResponse {
	var queries string = "?"

	queries = queries + removeEmptyQueries(c.Request.URL.Query())

	response := Get(config.Data.WhaleAddress + "/api/v1/catalog/projections" + queries)
	body := GetBody(response, &ProjectionsResponse{})

	return body.(*ProjectionsResponse)
}
