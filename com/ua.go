package com

import (
	"github.com/JREAMLU/core/useragent"
	"github.com/davecgh/go-spew/spew"
	"github.com/mssola/user_agent"
)

func ParseUserAgent_(ual string) map[string]interface{} {
	res := make(map[string]interface{})
	ua := user_agent.New(ual)

	ename, eversion := ua.Engine()
	bname, bversion := ua.Browser()

	mu.Lock()
	res["bot"] = ua.Bot()
	res["localization"] = ua.Localization()
	res["mobile"] = ua.Mobile()
	res["mozilla"] = ua.Mozilla()
	res["platform"] = ua.Platform()
	res["os"] = ua.OS()
	res["engine_name"] = ename
	res["engine_version"] = eversion
	res["browser_name"] = bname
	res["browser_version"] = bversion
	mu.Unlock()

	return res
}

func ParseUserAgent(ual string) map[string]interface{} {
	res := make(map[string]interface{})
	agent := useragent.ParseByString(ual)
	spew.Dump(agent)
	bot := false
	mobile := false

	if agent.Type == "robot" {
		bot = true
	}

	if agent.Device.Type == "mobile" {
		mobile = true
	}

	mu.Lock()
	res["browser_name"] = agent.Client["name"]
	res["browser_version"] = agent.Client["version"]
	res["platform"] = agent.OS.Name
	res["os"] = agent.OS.Version
	res["bot"] = bot
	res["mobile"] = mobile
	mu.Unlock()

	return res
}
