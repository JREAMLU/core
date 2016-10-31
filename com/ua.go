package com

import "github.com/JREAMLU/core/user_agent"

func ParseUserAgent(ual string) map[string]interface{} {
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
