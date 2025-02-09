package mocks

import (
	"encoding/hex"
	"math/big"

	cstaskmanager "github.com/ehsueh/trade-algo-avs-avs/contracts/bindings/TradeAlgoTaskManager"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

func MockTransactionByHash() *gethtypes.Transaction {
	// This is the calldata that is obtained from having only one operator participating in the AVS with a pubkey we choose
	// in our setting and is behaving maliciously.
	// This calldata was obtained by printing the "RespondToTask" transation and copying the calldata from the logs.
	input, err := hex.DecodeString("5baec9a00000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000090000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000001c00bcb2b68b6c68a5aea7fe75b5446c4ca410461fa226c2487d07eb2c504639cb52a3b3f7ef4f62985af31809fdc531483e5f1cd67aa1bcf0f8ac0d17e158aa9670940e64478db51fe630cc540dbeabea34d072a54fd7c743056e18174f9a1b64e00c874e4fcfb88d5c98a0240bc6f7f37d45f2226ca147317b3a2b7243ddb6c1b18b8fa707217c86c152f2ec3824f7de662a7c51d3bdf1c20ad574282dde7e6b30ddcb5a696413e8e02e318ae10dadd07c53822ed6832c97f8923d9f38f55b5170000000000000000000000000000000000000000000000000000000000000220000000000000000000000000000000000000000000000000000000000000026000000000000000000000000000000000000000000000000000000000000002a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010b5c664a181b916702c1acbbe4617a673cf9fffe733ff902353e315e11f9a7c11bd3ff011ee7d19e9d2ceef4475aac85313e8f31e4b8c66114d32e4ef0ee99590000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000")
	if err != nil {
		panic(err)
	}
	baseTx := &gethtypes.DynamicFeeTx{
		Data: input,
	}

	return gethtypes.NewTx(baseTx)
}

func MockRaiseAndResolveChallengeCall(blockNum, taskNum uint32) *gethtypes.Receipt {
	return &gethtypes.Receipt{
		Type:              1,
		PostState:         []byte{},
		Status:            0,
		CumulativeGasUsed: 1000000,
		Bloom:             gethtypes.Bloom{},
		Logs: []*gethtypes.Log{
			{
				Address: common.HexToAddress("0xeb61b2307b5846f9e0a72eece51d9cba88256683"),
				Topics: []common.Hash{
					// this is the actual TaskChallengedSuccessfully event from calling the func "raiseAndResolveChallenge"
					common.HexToHash("0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec"),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
					common.HexToHash("0x0000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc"),
				},
				Data:        []byte{},
				BlockNumber: uint64(blockNum),
				TxHash:      common.HexToHash("0x0"),
				TxIndex:     1,
				BlockHash:   common.HexToHash("0x0"),
				Index:       1,
				Removed:     false,
			},
		},
		TxHash:            common.HexToHash("0x0"),
		ContractAddress:   common.HexToAddress("0x0"),
		GasUsed:           0,
		EffectiveGasPrice: big.NewInt(int64(blockNum)),
		BlockHash:         common.HexToHash("0x0"),
		BlockNumber:       big.NewInt(int64(blockNum)),
		TransactionIndex:  1,
	}
}

func MockTaskResponseLog() gethtypes.Log {
	return gethtypes.Log{
		Address: common.HexToAddress("0x9e545e3c0baab3e08cdfd552c960a1050f373042"),
		Topics: []common.Hash{
			// this is the actual TaskChallengedSuccessfully event from calling the func "raiseAndResolveChallenge"
			common.HexToHash("0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a"),
		},
		Data:        common.Hex2Bytes("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a50fe07922f57ae3b4553201bfd7c11aca85e1541f91db8e62dca9c418dc5feae"),
		BlockNumber: uint64(100),
		TxHash:      common.HexToHash("0x0"),
		TxIndex:     1,
		BlockHash:   common.HexToHash("0x0"),
		Index:       1,
		Removed:     false,
	}
}

func MockTaskResponseReferenceTaskIndex(taskResponse cstaskmanager.ITradeAlgoTaskManagerTaskResponse) uint32 {
	return taskResponse.ReferenceTaskIndex
}
