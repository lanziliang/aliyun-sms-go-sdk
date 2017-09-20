package aliyunsms

import (
	"net/url"
	"testing"
)

const (
	AccessKeyID = "AccessKeyID"
	AccessKeySecret = "AccessKeySecret"

	SignName = "阿里云短信测试专用"
)

func Test_signature_method(t *testing.T) {
	string_to_sign := `POST&%2F&AccessKeyId%3Dtestid%26Action%3DSingleSendSms%26Format%3DXML%26ParamString%3D%257B%2522name%2522%253A%2522d%2522%252C%2522name1%2522%253A%2522d%2522%257D%26RecNum%3D13098765432%26RegionId%3Dcn-hangzhou%26SignName%3D%25E6%25A0%2587%25E7%25AD%25BE%25E6%25B5%258B%25E8%25AF%2595%26SignatureMethod%3DHMAC-SHA1%26SignatureNonce%3D9e030f6b-03a2-40f0-a6ba-157d44532fd0%26SignatureVersion%3D1.0%26TemplateCode%3DSMS_1650053%26Timestamp%3D2016-10-20T05%253A37%253A52Z%26Version%3D2016-09-27`
	signature := signature_method(`testsecret`, string_to_sign)
	if url.QueryEscape(signature) != `ka8PDlV7S9sYqxEMRnmlBv%2FDoAE%3D` {
		t.Error("signature_method failed")
	}
}


func TestSMSClient_SendOne(t *testing.T) {
	cli := New(AccessKeyID, AccessKeySecret)

	e, err := cli.SendOne("1805****318", SignName, "SMS_95600259", "{\"code\":\"123456\"}", "")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(e)
	}
}