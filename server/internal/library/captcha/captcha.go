package captcha

import (
	"context"
	"image/color"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mojocn/base64Captcha"
)

// Generate 生成验证码
func Generate(ctx context.Context) (id string, base64 string) {
	// 字符
	//	driver := &base64Captcha.DriverString{
	//		Height: 42,
	//		Width:  100,
	//		//NoiseCount:      50,
	//		//ShowLineOptions: 20,
	//		Length: 4,
	//		BgColor: &color.RGBA{
	//			R: 255,
	//			G: 250,
	//			B: 250,
	//			A: 250,
	//		},
	//		Source: "0123456789", // abcdefghjkmnpqrstuvwxyz23456789
	//		Fonts:  []string{"chromohv.ttf"},
	//	}

	// 算数
	driver := &base64Captcha.DriverMath{
		Height:          42,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 0,
		BgColor: &color.RGBA{
			R: 255,
			G: 250,
			B: 250,
			A: 250,
		},
		Fonts: []string{"chromohv.ttf"},
	}

	c := base64Captcha.NewCaptcha(driver.ConvertFonts(), base64Captcha.DefaultMemStore)
	id, base64, _, err := c.Generate()
	if err != nil {
		g.Log().Errorf(ctx, "captcha.Generate err:%+v", err)
	}
	return
}

// Verify 验证输入的验证码是否正确
func Verify(id, answer string) bool {
	if id == "" || answer == "" {
		return false
	}
	c := base64Captcha.NewCaptcha(new(base64Captcha.DriverString), base64Captcha.DefaultMemStore)
	return c.Verify(id, gstr.ToLower(answer), true)
}
