package control

import (
	"strconv"
)

type SafariCompatible struct {
}

//跳转到cn
func (a *SafariCompatible) ComRedirectToCn(c *gin.Context) {
	callbackUrl := c.Query("callbackUrl")
	challengeIdStr := c.Query("challengeId")
	challengeId, err := strconv.Atoi(challengeIdStr)
	if callbackUrl == "" || err != nil {
		logkit.Warnf(" 参数错误 callbackUrl %s challengeId %s err is %s", callbackUrl, challengeId, err)
		ginext.JSON(c, 40000, "callbackUrl challengeId 参数空", nil)
		return
	}
	logkit.Debugf("ComRedirectToCn request %#v", c.Request.URL)
	defaultSafariCompatibleDao := dao.GetDefaultSafariCompatibleDao()
	challenge, err := defaultSafariCompatibleDao.GetChallenge(challengeId)
	if err != nil {
		logkit.Warnf(" defaultSafariCompatibleDao.GetChallenge err challengeId is %d", challengeId)
		ginext.JSON(c, 50000, "查询错", nil)
		return
	}
	userId := c.GetInt(GinUidKey)
	userIdStr := fmt.Sprintf("%d", userId)
	redirectUrl := config.GetConfig().SafariCompatible.ComRedirectToCn +
		"?path=" + challenge.Path +
		"&userId=" + userIdStr +
		"&callbackUrl=" + url.QueryEscape(callbackUrl)
	logkit.Debugf("ComRedirectToCn redirectUrl %s", redirectUrl)
	c.Redirect(http.StatusFound, redirectUrl)
	return
}

//获取token
func (a *SafariCompatible) GetIsraelCookie(c *gin.Context) {
	callbackUrl := c.Query("callbackUrl")
	path := c.Query("path")
	userIdStr := c.Query("userId")
	logkit.Debugf("GetIsraelCookie request %#v", c.Request.URL)
	if callbackUrl == "" || path == "" || userIdStr == "" {
		logkit.Warnf(" 参数错误 callbackUrl %s path %s userIdStr %s ", callbackUrl, path, userIdStr)
		ginext.JSON(c, 40000, "callbackUrl path userId 参数空", nil)
		return
	}
	token, err := util.GenTokenWithMap(config.GetConfig().Isy, map[string]interface{}{"uid": userIdStr, "redirect_to": path})
	//token, err := util.GenTokenWithKeyV2(ur.Cfg.IsrtKey, uidStr, nickName, role, path, "")
	if err != nil {
		BuildSysErrResp(c, err)
		return
	}
	partUrl := config.GetConfig().CodeRush.CodeRushJumpUrl + token

	resp, err := http.Get(partUrl)
	if err != nil {
		logkit.Warnf(" 参数错误 callbackUrl %s path %s userIdStr %s ", callbackUrl, path, userIdStr)
		ginext.JSON(c, 40000, "callbackUrl path userId 参数空", nil)
		return
	}
	for k, values := range resp.Header {
		if k == "Set-Cookie" {
			for _, value := range values {
				c.Header(k, value)
			}
		}
	}
	logkit.Debugf("GetIsraelCookie callbackUrl %s", callbackUrl)
	c.Redirect(http.StatusFound, callbackUrl)
	return

}

