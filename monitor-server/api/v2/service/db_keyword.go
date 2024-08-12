package service

import (
	"fmt"
	"github.com/WeBankPartners/open-monitor/monitor-server/middleware"
	"github.com/WeBankPartners/open-monitor/monitor-server/models"
	"github.com/WeBankPartners/open-monitor/monitor-server/services/db"
	"github.com/gin-gonic/gin"
)

func ListDBKeywordConfig(c *gin.Context) {
	listType := c.Query("type")
	listGuid := c.Query("guid")
	result, err := db.ListDBKeywordConfig(listType, listGuid)
	if err != nil {
		middleware.ReturnHandleError(c, err.Error(), err)
	} else {
		middleware.ReturnSuccessData(c, result)
	}
}

func CreateDBKeywordConfig(c *gin.Context) {
	var param models.DbKeywordConfigObj
	var list []*models.DbKeywordMonitor
	var err error
	if err = c.ShouldBindJSON(&param); err != nil {
		middleware.ReturnValidateError(c, err.Error())
		return
	}
	if list, err = db.GetDbKeywordMonitorByName("", param.Name); err != nil {
		middleware.ReturnServerHandleError(c, err)
		return
	}
	if len(list) > 0 {
		middleware.ReturnServerHandleError(c, fmt.Errorf(middleware.GetMessageMap(c).AlertNameRepeatError))
		return
	}
	err = db.CreateDBKeywordConfig(&param, middleware.GetOperateUser(c))
	if err != nil {
		middleware.ReturnHandleError(c, err.Error(), err)
	} else {
		err = db.SyncDbMetric(false)
		if err != nil {
			middleware.ReturnHandleError(c, err.Error(), err)
		} else {
			respData := make(map[string]string)
			respData["guid"] = param.Guid
			middleware.ReturnSuccessData(c, respData)
		}
	}
}

func UpdateDBKeywordConfig(c *gin.Context) {
	var param models.DbKeywordConfigObj
	var list []*models.DbKeywordMonitor
	var err error
	if err = c.ShouldBindJSON(&param); err != nil {
		middleware.ReturnValidateError(c, err.Error())
		return
	}
	if list, err = db.GetDbKeywordMonitorByName(param.Guid, param.Name); err != nil {
		middleware.ReturnServerHandleError(c, err)
		return
	}
	if len(list) > 0 {
		middleware.ReturnServerHandleError(c, fmt.Errorf(middleware.GetMessageMap(c).AlertNameRepeatError))
		return
	}
	err = db.UpdateDBKeywordConfig(&param, middleware.GetOperateUser(c))
	if err != nil {
		middleware.ReturnHandleError(c, err.Error(), err)
	} else {
		err = db.SyncDbMetric(false)
		if err != nil {
			middleware.ReturnHandleError(c, err.Error(), err)
		} else {
			respData := make(map[string]string)
			respData["guid"] = param.Guid
			middleware.ReturnSuccessData(c, respData)
		}
	}
}

func DeleteDBKeywordConfig(c *gin.Context) {
	dbConfigGuid := c.Query("guid")
	err := db.DeleteDBKeywordConfig(dbConfigGuid)
	if err != nil {
		middleware.ReturnHandleError(c, err.Error(), err)
	} else {
		err = db.SyncDbMetric(false)
		if err != nil {
			middleware.ReturnHandleError(c, err.Error(), err)
		} else {
			middleware.ReturnSuccess(c)
		}
	}
}