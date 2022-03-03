package response

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"project_we/app/constant"
	"project_we/app/pkg/errors"
	"time"

	"github.com/beego/beego/v2/server/web"
)

// Response is response model
type Response struct {
	Meta  Meta        `json:"meta"`
	Error interface{} `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

// Meta is meta model
type Meta struct {
	Latency     string `json:"latency"`
	RequestUUID string `json:"request_uuid"`
}

// Error is error model
type Error struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

// WriteResponse is function to write the response
func WriteResponse(controller *web.Controller, errs errors.IError, response interface{}) {
	controller.Ctx.ResponseWriter.Header().Add("Content-Type", "application/json")
	var errResponse Error

	if errs != nil {
		errResponse.StatusCode = errs.GetCustomCode()
		errResponse.Message = errs.GetCustomMessage()

		controller.Ctx.ResponseWriter.WriteHeader(errs.GetHTTPCode())
	}

	//set latency
	birthTime := controller.Ctx.Input.GetData(constant.ContextBirthTime).(time.Time)
	latency := time.Since(birthTime).Seconds() * 1000
	latencyFmt := fmt.Sprintf("%.2f ms", latency)

	responseModel := Response{
		Meta: Meta{
			Latency:     latencyFmt,
			RequestUUID: controller.Ctx.Input.GetData(constant.ContextUUID).(string),
		},
		Data: response,
	}

	resp, _ := json.Marshal(responseModel)
	controller.Ctx.ResponseWriter.Write(resp)
}

// WriteResponseFilter is function to write the response for filter access
func WriteResponseFilter(ctx *context.Context, errs errors.IError, response interface{}) {
	ctx.ResponseWriter.Header().Add("Content-Type", "application/json")

	var errResponse Error
	if errs != nil {
		errResponse.StatusCode = errs.GetCustomCode()
		errResponse.Message = errs.GetCustomMessage()
	}

	//set latency
	birthTime := ctx.Input.GetData(constant.ContextBirthTime).(time.Time)
	latency := time.Since(birthTime).Seconds() * 1000
	latencyFmt := fmt.Sprintf("%.2f ms", latency)

	responseModel := Response{
		Meta: Meta{
			Latency:     latencyFmt,
			RequestUUID: ctx.Input.GetData(constant.ContextUUID).(string),
		},
		Error: errResponse,
		Data:  response,
	}

	resp, _ := json.Marshal(responseModel)
	ctx.ResponseWriter.Write(resp)
}
