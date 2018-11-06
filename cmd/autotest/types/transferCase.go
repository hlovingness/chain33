package types

import (
	"strconv"
)

type TransferCase struct {
	BaseCase
	From   string `toml:"from"`
	To     string `toml:"to"`
	Amount string `toml:"amount"`
}

type TransferPack struct {
	BaseCasePack
}

func (testCase *TransferCase) SendCommand(packID string) (PackFunc, error) {

	return DefaultSend(testCase, &TransferPack{}, packID)
}

func (pack *TransferPack) GetCheckHandlerMap() interface{} {

	funcMap := make(CheckHandlerMapDiscard, 2)
	funcMap["balance"] = pack.checkBalance

	return funcMap
}

func (pack *TransferPack) checkBalance(txInfo map[string]interface{}) bool {
	/*fromAddr := txInfo["tx"].(map[string]interface{})["from"].(string)
	toAddr := txInfo["tx"].(map[string]interface{})["to"].(string)*/
	interCase := pack.TCase.(*TransferCase)
	feeStr := txInfo["tx"].(map[string]interface{})["fee"].(string)
	logArr := txInfo["receipt"].(map[string]interface{})["logs"].([]interface{})
	logFee := logArr[0].(map[string]interface{})["log"].(map[string]interface{})
	logSend := logArr[1].(map[string]interface{})["log"].(map[string]interface{})
	logRecv := logArr[2].(map[string]interface{})["log"].(map[string]interface{})
	fee, _ := strconv.ParseFloat(feeStr, 64)
	Amount, _ := strconv.ParseFloat(interCase.Amount, 64)

	pack.FLog.Info("TransferBalanceDetails", "TestID", pack.PackID,
		"Fee", feeStr, "Amount", interCase.Amount,
		"FromPrev", logSend["prev"].(map[string]interface{})["balance"].(string),
		"FromCurr", logSend["current"].(map[string]interface{})["balance"].(string),
		"ToPrev", logRecv["prev"].(map[string]interface{})["balance"].(string),
		"ToCurr", logRecv["current"].(map[string]interface{})["balance"].(string))

	depositCheck := true
	//transfer to contract, deposit
	if len(logArr) == 4 {
		logDeposit := logArr[3].(map[string]interface{})["log"].(map[string]interface{})
		depositCheck = CheckBalanceDeltaWithAddr(logDeposit, interCase.From, Amount)
	}

	return CheckBalanceDeltaWithAddr(logFee, interCase.From, -fee) &&
		CheckBalanceDeltaWithAddr(logSend, interCase.From, -Amount) &&
		CheckBalanceDeltaWithAddr(logRecv, interCase.To, Amount) && depositCheck
}
