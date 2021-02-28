package httphandler

import (
	"fmt"
    "net/http"
	"github.com/fengjijiao/dingding-push-restful-api/pkg/conf"
	"github.com/fengjijiao/dingding-push-restful-api/pkg/logio"
	"go.uber.org/zap"
	//"path"
	"github.com/fengjijiao/dingding-push-restful-api/pkg/sqlhandler"
	"encoding/json"
	"github.com/imroc/req"
)

type SendInfo struct {
	Msg struct {
		Markdown struct {
			Text  string `json:"text"`
			Title string `json:"title"`
		} `json:"markdown"`
		Text struct {
			Content string `json:"content"`
		} `json:"text"`
		MsgType string `json:"msgtype"`
	} `json:"msg"` //不超过2048字节
	ToAllUser  string `json:"to_all_user"`
	AgentId    string `json:"agent_id"`
	DeptIdList string `json:"dept_id_list"`
	UserIdList string `json:"userid_list"`
}

type ErrorInfo struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

func sendMarkDownHttpHandler(w http.ResponseWriter, hr *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	hr.ParseForm()
	title := hr.FormValue("title")
	body := hr.FormValue("body")
	touser := hr.FormValue("touser")
	toparty := hr.FormValue("toparty")
	if len(title) <= 0 || len(body) <= 0 {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, "send message failed, missing required parameters!"})
		return
	}
	var res ErrorInfo
	var sendInfo SendInfo
	if len(touser) > 0 {
		sendInfo.UserIdList = touser
	}else if len(toparty) > 0 {
		sendInfo.DeptIdList = toparty
	}else {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, "send message failed, missing required parameters!"})
		return
	}
	sendInfo.Msg.MsgType = "markdown"
	sendInfo.Msg.Title = title
	sendInfo.Msg.Text = body
	sendInfo.AgentId = conf.Config.DingdingAgentId
	param, err := json.Marshal(&sendInfo)
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, err.Error()})
		logio.Logger.Error("sendHttpHandler: ", zap.Error(err))
		return
	}
	token, err := sqlhandler.GetToken()
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, err.Error()})
		logio.Logger.Error("sendHttpHandler: ", zap.Error(err))
		return
	}
	r, err := req.Post(fmt.Sprintf(`https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=%s`, token), param)
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, err.Error()})
		logio.Logger.Error("sendHttpHandler: ", zap.Error(err))
		return
	}
	r.ToJSON(&res)
	if res.ErrCode == 0 {
		json.NewEncoder(w).Encode(&ErrorInfo{0, "send message success!"})
	}else {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, "send message failed!"+res.ErrMsg})
	}
}

func sendTextHttpHandler(w http.ResponseWriter, hr *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	hr.ParseForm()
	context := hr.FormValue("context")
	touser := hr.FormValue("touser")
	toparty := hr.FormValue("toparty")
	if len(context) <= 0 {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, "send message failed, missing required parameters!"})
		return
	}
	var res ErrorInfo
	var sendInfo SendInfo
	if len(touser) > 0 {
		sendInfo.UserIdList = touser
	}else if len(toparty) > 0 {
		sendInfo.DeptIdList = toparty
	}else {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, "send message failed, missing required parameters!"})
		return
	}
	sendInfo.Msg.MsgType = "text"
	sendInfo.AgentId = conf.Config.DingdingAgentId
	sendInfo.Msg.Text.Content = context
	param, err := json.Marshal(&sendInfo)
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, err.Error()})
		logio.Logger.Error("sendHttpHandler: ", zap.Error(err))
		return
	}
	token, err := sqlhandler.GetToken()
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, err.Error()})
		logio.Logger.Error("sendHttpHandler: ", zap.Error(err))
		return
	}
	r, err := req.Post(fmt.Sprintf(`https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=%s`, token), param)
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, err.Error()})
		logio.Logger.Error("sendHttpHandler: ", zap.Error(err))
		return
	}
	r.ToJSON(&res)
	if res.ErrCode == 0 {
		json.NewEncoder(w).Encode(&ErrorInfo{0, "send message success!"})
	}else {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, "send message failed!"+res.ErrMsg})
	}
}