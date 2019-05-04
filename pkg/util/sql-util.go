package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func MatchAggregate(aggregate string) string {
	if aggregate == "" {
		return "LEFT"
	}
	if aggregate == "all" {
		return "FULL OUTER"
	}
	if aggregate == "inner" {
		return "LEFT INNER"
	}
	if aggregate == "right" {
		return "RIGHT"
	}
	if aggregate == "ri" {
		return "RIGHT INNER"
	}
	if aggregate == "i" {
		return "INNER"
	}
	if aggregate == "lo" {
		return "LEFT OUTER"
	}
	if aggregate == "ro" {
		return "RIGHT OUTER"
	}
	return "LEFT"
}

func getParamInQueryOrHeader(context *gin.Context, param string) (string, bool) {
	var stringData = ""
	var boolData = false
	if context.Query(param) != "" {
		stringData = context.Query(param)
		boolData = true
	} else {
		stringData = context.Request.Header.Get(param)
		boolData = true
	}
	return stringData, boolData
}

func GetQueryOptionFromContext(context *gin.Context) {
	var requestOption RequestOption
	order, _ := getParamInQueryOrHeader(context, "order")
	page, _ := getParamInQueryOrHeader(context, "page")
	pageSize, _ := getParamInQueryOrHeader(context, "pagesize")
	aggregate, _ := getParamInQueryOrHeader(context, "aggregate")
	_, strict := getParamInQueryOrHeader(context, "strict")
	state, _ := getParamInQueryOrHeader(context, "state")
	_, specify := getParamInQueryOrHeader(context, "specify")
	_, autoRank := getParamInQueryOrHeader(context, "auto_rank")
	requestOption.Order = order
	requestOption.Aggregate = aggregate
	requestOption.Strict = strict
	requestOption.State = state
	requestOption.Specify = specify
	requestOption.AutoRank = autoRank
	if page != "" {
		requestOption.Page, _ = strconv.Atoi(page)
	}
	if pageSize != "" {
		requestOption.PageSize, _ = strconv.Atoi(pageSize)
	}
}
