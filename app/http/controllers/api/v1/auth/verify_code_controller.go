package auth

import (
	"github.com/gin-gonic/gin"
	requests2 "go-mall-api/app/http/controllers/api/requests"
	v1 "go-mall-api/app/http/controllers/api/v1"
	"go-mall-api/pkg/captcha"
	"go-mall-api/pkg/logger"
	"go-mall-api/pkg/response"
	"go-mall-api/pkg/verifycode"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	//生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	//记录错误日志,因为验证码是用户的入口, 出错是应该记error等级的日志
	logger.LogIf(err)
	//返回给用户
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

// SendUsingPhone 发送手机验证码
func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {
	//1 验证表单
	request := requests2.VerifyCodePhoneRequest{}
	if ok := requests2.Validate(c, &request, requests2.VerifyCodePhone); !ok {
		return
	}

	//2 发送SMS
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送短信失败~")
	} else {
		response.Success(c)
	}
}

// SendUsingEmail 发送 Email 验证码
func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {
	// 1. 验证表单
	request := requests2.VerifyCodeEmailRequest{}
	if ok := requests2.Validate(c, &request, requests2.VerifyCodeEmail); !ok {
		return
	}

	// 2. 发送SMS
	err := verifycode.NewVerifyCode().SendEmail(request.Email)
	if err != nil {
		response.Abort500(c, "发送Email验证码失败~")
	} else {
		response.Success(c)
	}
}
