package config

import (
	"github.com/chenyuIT/framework/facades"
)

func init() {
	config := facades.Config
	config.Add("sms", map[string]any{
		"default": config.Env("SMS_VENDOR", "alidayu"),
		"vendors": map[string]any{
			//阿里云短信配置
			"alidayu": map[string]any{
				"appkey":    config.Env("SMS_APPKEY"),
				"appSecret": config.Env("SMS_APPSECRET"),
				"issendbox": false,
				"servicelist": map[string]any{
					"register": map[string]any{
						"vendor":      "alidayu",       //短信通道供应商
						"group":       "db1",           //相同组内的uid数据共享
						"smstpl":      "SMS_154950909", //阿里云短信短信模板id
						"signname":    "xx",            //阿里云短信短信签名
						"callback":    "http://127.0.01/test9.php",
						"allowcitys":  []string{"0575", "0571", "0574"}, //仅限如下的手机号归属区接收验证码
						"maxsendnums": 4,                                //一个手机号每天发送限额,这个受短信运营商的限制。
						"validtime":   600,                              //单位：秒 。 收到的手机验证码x秒内有效，超过后验证无效；
						"mode":        3,                                //模式  1：只有手机号对应的uid存在时才能发送，2：只有uid不存在时才能发送，3：不管uid是否存在都发送
						"outformat":   "mobcent",                        //RestAPi接口输出样式（mobcent,default）
					},
					"restpwd": map[string]any{
						"vendor":      "alidayu",
						"group":       "db1",
						"smstpl":      "SMS_39190087",
						"signname":    "xxx",
						"callback":    "",
						"allowcitys":  []string{"0578"},
						"maxsendnums": 2,
						"validtime":   360,
						"mode":        3,
						"outformat":   "mobcent",
					},
					"getpwd": map[string]any{
						"vendor":      "yuntongxun",
						"group":       "db1",
						"smstpl":      149350,
						"signname":    "",
						"callback":    "",
						"allowcitys":  []string{"0578"},
						"maxsendnums": 2,
						"validtime":   360,
						"mode":        3,
						"outformat":   "mobcent",
					},
				},
			},
			//云通讯配置  http://www.yuntongxun.com/
			"yuntongxun": map[string]any{
				"AccountSid":   "8a48b55e434514c9c31921a039b",
				"AccountToken": "61434dc2b245435eadf82d381fa3f",
				"AppId":        "aaf98f8fsdafd2678c9d07875040f",
				"SoftVersion":  "2013-12-26",
				"RestURL":      "https://app.cloopen.com:8883",
			},
			//互亿无线  http://www.ihuyi.cn/
			"hywx": map[string]any{
				"account":  "6666666666666666",
				"password": "88888888888888888888888",
				"RestURL":  "http://106.ihuyi.cn/webservice/sms.php?method=Submit",
			},
		},
		"timezone": "PRC",      //时区设置
		"timeout":  "5",        //短信供应商网关响应超时时间（秒）
		"dbPath":   "level.db", //数据库保存路径
		"errormsg": map[string]any{
			"err_model_not_ok1":         "当前用户(%s)不存在，不能发送手机验证码",
			"err_model_not_ok2":         "当前用户(%s)存在，不能发送手机验证码",
			"err_code_not_ok":           "手机验证码:%s 不正确，请重新输入",
			"err_vailtime_not_ok":       "手机验证码已超时:%s，请重新获取",
			"err_per_day_max_send_nums": "一个手机号每天仅限发送%d条验证码",
			"err_per_minute_send_num":   "一个手机号每分钟仅限发送1条验证码",
			"err_allow_areacode":        "手机号归属地不允许，仅限于:%s",
			"err_not_uid":               "手机号%s对应的%s不存在或者不匹配",
		},
	})
}
