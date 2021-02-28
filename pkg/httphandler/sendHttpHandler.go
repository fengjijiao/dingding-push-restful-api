package httphandler

import (
	"fmt"
    "net/http"
	"github.com/fengjijiao/dingding-push-restful-api/pkg/conf"
	"github.com/fengjijiao/dingding-push-restful-api/pkg/logio"
	"go.uber.org/zap"
	"github.com/fengjijiao/dingding-push-restful-api/pkg/sqlhandler"
	"encoding/json"
	"github.com/imroc/req"
)

type SendInfo struct {
	Msg struct {
		Markdown struct {
			Text  string `json:"text,omitempty"`
			Title string `json:"title,omitempty"`
		} `json:"markdown,omitempty"`
		Text struct {
			Content string `json:"content,omitempty"`
		} `json:"text,omitempty"`
		MsgType string `json:"msgtype,omitempty"`
	} `json:"msg,omitempty"` //不超过2048字节
	ToAllUser  bool `json:"to_all_user,omitempty"`
	AgentId    int `json:"agent_id"`
	DeptIdList string `json:"dept_id_list,omitempty"`
	UserIdList string `json:"userid_list,omitempty"`
}

type ErrorInfo struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

func sendMessageHttpHandler(w http.ResponseWriter, hr *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	msgtype := hr.URL.Query().Get("msgtype")
	markdownFlag := true
	if len(msgtype) <= 0 {
		msgtype = "text"
		markdownFlag = false
	}else {
		markdownFlag = (msgtype == "markdown")
	}
	var title string
	var body string
	var context string
	hr.ParseForm()
	if markdownFlag {
		title = hr.FormValue("title")
		body = hr.FormValue("body")
		if len(title) <= 0 || len(body) <= 0 {
			json.NewEncoder(w).Encode(&ErrorInfo{-1, "send message failed, missing required parameters!"})
			return
		}
	}else {
		context = hr.FormValue("context")
		if len(context) <= 0 {
			json.NewEncoder(w).Encode(&ErrorInfo{-1, "send message failed, missing required parameters!"})
			return
		}
	}
	toDeptFlag := false
	touser := hr.FormValue("touser")
	toparty := hr.FormValue("toparty")
	if len(touser) <= 0 {
		toDeptFlag = true
	}
	if len(touser) <= 0 && len(toparty) <= 0 {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, "send message failed, missing required parameters!"})
		return
	}
	var res ErrorInfo
	var sendInfo SendInfo
	sendInfo.AgentId = conf.Config.DingdingAgentId
	if toDeptFlag {
		sendInfo.DeptIdList = toparty
	}else {
		sendInfo.UserIdList = touser
	}
	sendInfo.Msg.MsgType = msgtype
	if markdownFlag {
		sendInfo.Msg.Markdown.Text = body
		sendInfo.Msg.Markdown.Title = title
	}else {
		sendInfo.Msg.Text.Content = context
	}
	param, err := json.Marshal(&sendInfo)
	fmt.Printf("param: %s\n", string(param))
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