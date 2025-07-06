package go_starpago

import (
	"crypto/tls"
	"github.com/asaka1234/go-starpago/utils"
	"github.com/mitchellh/mapstructure"
)

// withdraw
func (cli *Client) Withdraw(req StarPagoWithdrawReq) (*StarPagoWithdrawResponse, error) {

	rawURL := cli.Params.WithdrawUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["appId"] = cli.Params.MerchantId
	params["notifyUrl"] = cli.Params.WithdrawBackUrl
	params["payMethod"] = "PK_DF" //fixed 巴基斯坦代付

	//签名
	signStr := utils.Sign(params, cli.Params.AccessKey)
	params["sign"] = signStr

	//返回值会放到这里
	var result StarPagoWithdrawResponse

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetBody(params).
		SetDebug(cli.debugMode).
		SetResult(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	return &result, err
}
