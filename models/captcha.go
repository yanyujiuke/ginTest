package models

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// 设置自带的 store
// var store = base64Captcha.DefaultMemStore
var store base64Captcha.Store = RedisStore{}

// 获取验证码
func MakeCaptcha() (lid, lb64s string, lerr error) {
	var driver base64Captcha.Driver
	// 配置验证码信息
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	// ConvertFonts 按名称加载字体
	driver = driverString.ConvertFonts()
	// 创建 Captcha
	captcha := base64Captcha.NewCaptcha(driver, store)
	// Generate 生成随机 id、base64 图像字符串
	lid, lb64s, lerr = captcha.Generate()
	return lid, lb64s, lerr
}

// 验证验证码
func VerifyCaptcha(id, verifyValue string) bool {
	if store.Verify(id, verifyValue, true) {
		return true
	}
	return false
}
