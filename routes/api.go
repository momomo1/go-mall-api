package routes

import (
	"github.com/gin-gonic/gin"
	controllers "go-mall-api/app/http/controllers/api/v1"
	"go-mall-api/app/http/controllers/api/v1/auth"
	"go-mall-api/app/http/controllers/api/v1/mqdemo"
	"go-mall-api/app/http/middlewares"
	"go-mall-api/pkg/config"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	var v1 *gin.RouterGroup
	if len(config.Get("app.api_domain")) == 0 {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 API 每小时最多 60 个请求（根据 IP）。
	v1.Use(middlewares.LimitIP("208-H"))
	{
		authGroup := v1.Group("/auth")
		{
			// 登录
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)

			// 注册用户
			suc := new(auth.SignupController)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)
			// 图片验证码
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
		}

		uc := new(controllers.UsersController)
		// 获取当前用户
		v1.GET("/user", middlewares.AuthJWT(), uc.Index)

		mqGroup := v1.Group("/mq")
		{
			suc := new(mqdemo.MqController)
			//判断手机是否已注册
			mqGroup.POST("/push/demo", suc.PushMsg)
		}
	}
}
