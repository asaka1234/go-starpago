package go_starpago

import (
	"fmt"
	"testing"
)

// 巴基斯坦 代付
func TestPKWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &StarPagoInitParams{MERCHANT_ID, ACCESS_SECRET, IP, DEPOSIT_URL, WITHDRAW_URL, DEPOSIT_BACK_URL, WITHDRAW_BACK_URL})

	//发请求
	resp, err := cli.Withdraw(GenPKWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenPKWithdrawRequestDemo() StarPagoWithdrawReq {
	return StarPagoWithdrawReq{
		MerOrderNo: "111",
		Currency:   "PKR",
		Amount:     "6000", //100-500000
		PayMethod:  PKPayInMethodDF,
		Extra: StarPagoPKWithdrawReqExtra{
			BankCode:  "BIPL",
			AccountNo: "129171971",
			Cnic:      "1891917917912",
			Mobile:    "03123456789",
		},
	}
}

//=======

// 菲律宾 代付
func TestPHWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &StarPagoInitParams{MERCHANT_ID, ACCESS_SECRET, IP, DEPOSIT_URL, WITHDRAW_URL, DEPOSIT_BACK_URL, WITHDRAW_BACK_URL})

	//发请求
	resp, err := cli.Withdraw(GenPHWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenPHWithdrawRequestDemo() StarPagoWithdrawReq {
	return StarPagoWithdrawReq{
		MerOrderNo: "2222",
		Currency:   "PHP",
		Amount:     "200", //100-500000
		PayMethod:  PHPayOutMethodWallet,
		Extra: StarPagoPHWithdrawReqExtra{
			BankCode:    "CRBRB",
			AccountNo:   "129171971",
			AccountName: "john",
			Email:       "1891917917912@qq.com",
			Mobile:      "03123456789",
		},
	}
}
