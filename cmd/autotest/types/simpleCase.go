package types

import (
	"strings"

	"errors"
)

//simple case just executes without checking, suitable for init situation

type SimpleCase struct {
	BaseCase
}

type SimplePack struct {
	BaseCasePack
}

func (testCase *SimpleCase) SendCommand(packID string) (PackFunc, error) {


	output, err := RunChain33Cli(strings.Fields(testCase.GetCmd()))
	if err != nil {
		return nil, errors.New(output)
	}
	testPack := &SimplePack{}
	pack := testPack.GetBasePack()
	pack.TxHash = output
	pack.TCase = testCase

	pack.PackID = packID
	pack.CheckTimes = 0
	return testPack, nil


}

//simple case needn't check
func (pack *SimplePack) CheckResult(handlerMap interface{}) (bCheck bool, bSuccess bool) {

	bCheck = true
	bSuccess = true
	if strings.Contains(pack.TxHash, "Err") || strings.Contains(pack.TxHash, "connection refused") {

		bSuccess = false
	}

	return bCheck, bSuccess
}
