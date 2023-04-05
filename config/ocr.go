package config

import "github.com/chenyuIT/framework/facades"

func init() {
	config := facades.Config
	config.Add("ocr", map[string]any{
		"default": config.Env("OCR_VENDOR", "baidu"),

		"vendors": map[string]any{
			"baidu": map[string]any{
				"tokenUrl":       "https://aip.baidubce.com/oauth/2.0/token",
				"ocrUrl":         "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic",       //基础ocr
				"ocrAccurateUrl": "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic",      //高精度
				"iOcrUrl":        "https://aip.baidubce.com/rest/2.0/solution/v1/iocr/recognise", //自定义
				"granttype":      "client_credentials",
				"clientid":       "pk7KNGeeMrkpYSgs5BkE8IZj",         // 填写百度应用的API_KEY
				"clientsecret":   "IXdtgutnPlxssBgDsTeGdwv1qbg0XW46", // 填写百度应用的Secret_KEY
			},
			"ali": map[string]any{
				"tokenUrl":       config.Env("OCR_TOKENURL"),
				"ocrUrl":         config.Env("OCR_OCRURL"),
				"ocrAccurateUrl": config.Env("OCR_ACCURATEURL"),
				"iOcrUrl":        config.Env("OCR_IOCRURL"),
				"granttype":      config.Env("OCR_GRANTTYPE"),
				"clientid":       config.Env("OCR_CLIENTID"),
				"clientsecret":   config.Env("OCR_CLIENTSECRET"),
			},
		},
	})
}
