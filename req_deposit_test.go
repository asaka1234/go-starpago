package go_starpago

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

// 巴基斯坦 代收
func TestPKDeposit(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &StarPagoInitParams{MERCHANT_ID, ACCESS_SECRET, IP, DEPOSIT_URL, WITHDRAW_URL, DEPOSIT_BACK_URL, WITHDRAW_BACK_URL})

	//发请求
	resp, err := cli.Deposit(GenPKDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

// 巴基斯坦
func GenPKDepositRequestDemo() StarPagoDepositReq {
	return StarPagoDepositReq{
		MerOrderNo: "323231224", //商户id
		Currency:   "PKR",
		Amount:     "5000",
		PayMethod:  PKPayInMethodWALLET,
		Extra: StarPagoPKDepositReqExtra{
			PayType: "PK_WALLET_EASYPAISA", //这个不需要其他参数
			Mobile:  "03123456789",
		},
	}
}

//===================

// 菲律宾 代收
func TestPHDeposit(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &StarPagoInitParams{MERCHANT_ID, ACCESS_SECRET, IP, DEPOSIT_URL, WITHDRAW_URL, DEPOSIT_BACK_URL, WITHDRAW_BACK_URL})

	//发请求
	resp, err := cli.Deposit(GenPHDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

// 巴基斯坦
func GenPHDepositRequestDemo() StarPagoDepositReq {
	return StarPagoDepositReq{
		MerOrderNo: "323231224", //商户id
		Currency:   "PHP",
		Amount:     "5000",
		PayMethod:  PHPayInMethodMAYA,
		Extra: StarPagoPHDepositReqExtra{
			Email:       "demo@gmail.com", //这个不需要其他参数
			Phone:       "03123456789",
			AccountName: "john",
		},
	}
}
