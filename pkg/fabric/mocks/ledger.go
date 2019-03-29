// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric-cli/pkg/fabric"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
)

type Ledger struct {
	QueryBlockStub        func(uint64, ...ledger.RequestOption) (*common.Block, error)
	queryBlockMutex       sync.RWMutex
	queryBlockArgsForCall []struct {
		arg1 uint64
		arg2 []ledger.RequestOption
	}
	queryBlockReturns struct {
		result1 *common.Block
		result2 error
	}
	queryBlockReturnsOnCall map[int]struct {
		result1 *common.Block
		result2 error
	}
	QueryBlockByHashStub        func([]byte, ...ledger.RequestOption) (*common.Block, error)
	queryBlockByHashMutex       sync.RWMutex
	queryBlockByHashArgsForCall []struct {
		arg1 []byte
		arg2 []ledger.RequestOption
	}
	queryBlockByHashReturns struct {
		result1 *common.Block
		result2 error
	}
	queryBlockByHashReturnsOnCall map[int]struct {
		result1 *common.Block
		result2 error
	}
	QueryBlockByTxIDStub        func(fab.TransactionID, ...ledger.RequestOption) (*common.Block, error)
	queryBlockByTxIDMutex       sync.RWMutex
	queryBlockByTxIDArgsForCall []struct {
		arg1 fab.TransactionID
		arg2 []ledger.RequestOption
	}
	queryBlockByTxIDReturns struct {
		result1 *common.Block
		result2 error
	}
	queryBlockByTxIDReturnsOnCall map[int]struct {
		result1 *common.Block
		result2 error
	}
	QueryConfigStub        func(...ledger.RequestOption) (fab.ChannelCfg, error)
	queryConfigMutex       sync.RWMutex
	queryConfigArgsForCall []struct {
		arg1 []ledger.RequestOption
	}
	queryConfigReturns struct {
		result1 fab.ChannelCfg
		result2 error
	}
	queryConfigReturnsOnCall map[int]struct {
		result1 fab.ChannelCfg
		result2 error
	}
	QueryInfoStub        func(...ledger.RequestOption) (*fab.BlockchainInfoResponse, error)
	queryInfoMutex       sync.RWMutex
	queryInfoArgsForCall []struct {
		arg1 []ledger.RequestOption
	}
	queryInfoReturns struct {
		result1 *fab.BlockchainInfoResponse
		result2 error
	}
	queryInfoReturnsOnCall map[int]struct {
		result1 *fab.BlockchainInfoResponse
		result2 error
	}
	QueryTransactionStub        func(fab.TransactionID, ...ledger.RequestOption) (*peer.ProcessedTransaction, error)
	queryTransactionMutex       sync.RWMutex
	queryTransactionArgsForCall []struct {
		arg1 fab.TransactionID
		arg2 []ledger.RequestOption
	}
	queryTransactionReturns struct {
		result1 *peer.ProcessedTransaction
		result2 error
	}
	queryTransactionReturnsOnCall map[int]struct {
		result1 *peer.ProcessedTransaction
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Ledger) QueryBlock(arg1 uint64, arg2 ...ledger.RequestOption) (*common.Block, error) {
	fake.queryBlockMutex.Lock()
	ret, specificReturn := fake.queryBlockReturnsOnCall[len(fake.queryBlockArgsForCall)]
	fake.queryBlockArgsForCall = append(fake.queryBlockArgsForCall, struct {
		arg1 uint64
		arg2 []ledger.RequestOption
	}{arg1, arg2})
	fake.recordInvocation("QueryBlock", []interface{}{arg1, arg2})
	fake.queryBlockMutex.Unlock()
	if fake.QueryBlockStub != nil {
		return fake.QueryBlockStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryBlockReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Ledger) QueryBlockCallCount() int {
	fake.queryBlockMutex.RLock()
	defer fake.queryBlockMutex.RUnlock()
	return len(fake.queryBlockArgsForCall)
}

func (fake *Ledger) QueryBlockCalls(stub func(uint64, ...ledger.RequestOption) (*common.Block, error)) {
	fake.queryBlockMutex.Lock()
	defer fake.queryBlockMutex.Unlock()
	fake.QueryBlockStub = stub
}

func (fake *Ledger) QueryBlockArgsForCall(i int) (uint64, []ledger.RequestOption) {
	fake.queryBlockMutex.RLock()
	defer fake.queryBlockMutex.RUnlock()
	argsForCall := fake.queryBlockArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Ledger) QueryBlockReturns(result1 *common.Block, result2 error) {
	fake.queryBlockMutex.Lock()
	defer fake.queryBlockMutex.Unlock()
	fake.QueryBlockStub = nil
	fake.queryBlockReturns = struct {
		result1 *common.Block
		result2 error
	}{result1, result2}
}

func (fake *Ledger) QueryBlockReturnsOnCall(i int, result1 *common.Block, result2 error) {
	fake.queryBlockMutex.Lock()
	defer fake.queryBlockMutex.Unlock()
	fake.QueryBlockStub = nil
	if fake.queryBlockReturnsOnCall == nil {
		fake.queryBlockReturnsOnCall = make(map[int]struct {
			result1 *common.Block
			result2 error
		})
	}
	fake.queryBlockReturnsOnCall[i] = struct {
		result1 *common.Block
		result2 error
	}{result1, result2}
}

func (fake *Ledger) QueryBlockByHash(arg1 []byte, arg2 ...ledger.RequestOption) (*common.Block, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.queryBlockByHashMutex.Lock()
	ret, specificReturn := fake.queryBlockByHashReturnsOnCall[len(fake.queryBlockByHashArgsForCall)]
	fake.queryBlockByHashArgsForCall = append(fake.queryBlockByHashArgsForCall, struct {
		arg1 []byte
		arg2 []ledger.RequestOption
	}{arg1Copy, arg2})
	fake.recordInvocation("QueryBlockByHash", []interface{}{arg1Copy, arg2})
	fake.queryBlockByHashMutex.Unlock()
	if fake.QueryBlockByHashStub != nil {
		return fake.QueryBlockByHashStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryBlockByHashReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Ledger) QueryBlockByHashCallCount() int {
	fake.queryBlockByHashMutex.RLock()
	defer fake.queryBlockByHashMutex.RUnlock()
	return len(fake.queryBlockByHashArgsForCall)
}

func (fake *Ledger) QueryBlockByHashCalls(stub func([]byte, ...ledger.RequestOption) (*common.Block, error)) {
	fake.queryBlockByHashMutex.Lock()
	defer fake.queryBlockByHashMutex.Unlock()
	fake.QueryBlockByHashStub = stub
}

func (fake *Ledger) QueryBlockByHashArgsForCall(i int) ([]byte, []ledger.RequestOption) {
	fake.queryBlockByHashMutex.RLock()
	defer fake.queryBlockByHashMutex.RUnlock()
	argsForCall := fake.queryBlockByHashArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Ledger) QueryBlockByHashReturns(result1 *common.Block, result2 error) {
	fake.queryBlockByHashMutex.Lock()
	defer fake.queryBlockByHashMutex.Unlock()
	fake.QueryBlockByHashStub = nil
	fake.queryBlockByHashReturns = struct {
		result1 *common.Block
		result2 error
	}{result1, result2}
}

func (fake *Ledger) QueryBlockByHashReturnsOnCall(i int, result1 *common.Block, result2 error) {
	fake.queryBlockByHashMutex.Lock()
	defer fake.queryBlockByHashMutex.Unlock()
	fake.QueryBlockByHashStub = nil
	if fake.queryBlockByHashReturnsOnCall == nil {
		fake.queryBlockByHashReturnsOnCall = make(map[int]struct {
			result1 *common.Block
			result2 error
		})
	}
	fake.queryBlockByHashReturnsOnCall[i] = struct {
		result1 *common.Block
		result2 error
	}{result1, result2}
}

func (fake *Ledger) QueryBlockByTxID(arg1 fab.TransactionID, arg2 ...ledger.RequestOption) (*common.Block, error) {
	fake.queryBlockByTxIDMutex.Lock()
	ret, specificReturn := fake.queryBlockByTxIDReturnsOnCall[len(fake.queryBlockByTxIDArgsForCall)]
	fake.queryBlockByTxIDArgsForCall = append(fake.queryBlockByTxIDArgsForCall, struct {
		arg1 fab.TransactionID
		arg2 []ledger.RequestOption
	}{arg1, arg2})
	fake.recordInvocation("QueryBlockByTxID", []interface{}{arg1, arg2})
	fake.queryBlockByTxIDMutex.Unlock()
	if fake.QueryBlockByTxIDStub != nil {
		return fake.QueryBlockByTxIDStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryBlockByTxIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Ledger) QueryBlockByTxIDCallCount() int {
	fake.queryBlockByTxIDMutex.RLock()
	defer fake.queryBlockByTxIDMutex.RUnlock()
	return len(fake.queryBlockByTxIDArgsForCall)
}

func (fake *Ledger) QueryBlockByTxIDCalls(stub func(fab.TransactionID, ...ledger.RequestOption) (*common.Block, error)) {
	fake.queryBlockByTxIDMutex.Lock()
	defer fake.queryBlockByTxIDMutex.Unlock()
	fake.QueryBlockByTxIDStub = stub
}

func (fake *Ledger) QueryBlockByTxIDArgsForCall(i int) (fab.TransactionID, []ledger.RequestOption) {
	fake.queryBlockByTxIDMutex.RLock()
	defer fake.queryBlockByTxIDMutex.RUnlock()
	argsForCall := fake.queryBlockByTxIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Ledger) QueryBlockByTxIDReturns(result1 *common.Block, result2 error) {
	fake.queryBlockByTxIDMutex.Lock()
	defer fake.queryBlockByTxIDMutex.Unlock()
	fake.QueryBlockByTxIDStub = nil
	fake.queryBlockByTxIDReturns = struct {
		result1 *common.Block
		result2 error
	}{result1, result2}
}

func (fake *Ledger) QueryBlockByTxIDReturnsOnCall(i int, result1 *common.Block, result2 error) {
	fake.queryBlockByTxIDMutex.Lock()
	defer fake.queryBlockByTxIDMutex.Unlock()
	fake.QueryBlockByTxIDStub = nil
	if fake.queryBlockByTxIDReturnsOnCall == nil {
		fake.queryBlockByTxIDReturnsOnCall = make(map[int]struct {
			result1 *common.Block
			result2 error
		})
	}
	fake.queryBlockByTxIDReturnsOnCall[i] = struct {
		result1 *common.Block
		result2 error
	}{result1, result2}
}

func (fake *Ledger) QueryConfig(arg1 ...ledger.RequestOption) (fab.ChannelCfg, error) {
	fake.queryConfigMutex.Lock()
	ret, specificReturn := fake.queryConfigReturnsOnCall[len(fake.queryConfigArgsForCall)]
	fake.queryConfigArgsForCall = append(fake.queryConfigArgsForCall, struct {
		arg1 []ledger.RequestOption
	}{arg1})
	fake.recordInvocation("QueryConfig", []interface{}{arg1})
	fake.queryConfigMutex.Unlock()
	if fake.QueryConfigStub != nil {
		return fake.QueryConfigStub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryConfigReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Ledger) QueryConfigCallCount() int {
	fake.queryConfigMutex.RLock()
	defer fake.queryConfigMutex.RUnlock()
	return len(fake.queryConfigArgsForCall)
}

func (fake *Ledger) QueryConfigCalls(stub func(...ledger.RequestOption) (fab.ChannelCfg, error)) {
	fake.queryConfigMutex.Lock()
	defer fake.queryConfigMutex.Unlock()
	fake.QueryConfigStub = stub
}

func (fake *Ledger) QueryConfigArgsForCall(i int) []ledger.RequestOption {
	fake.queryConfigMutex.RLock()
	defer fake.queryConfigMutex.RUnlock()
	argsForCall := fake.queryConfigArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Ledger) QueryConfigReturns(result1 fab.ChannelCfg, result2 error) {
	fake.queryConfigMutex.Lock()
	defer fake.queryConfigMutex.Unlock()
	fake.QueryConfigStub = nil
	fake.queryConfigReturns = struct {
		result1 fab.ChannelCfg
		result2 error
	}{result1, result2}
}

func (fake *Ledger) QueryConfigReturnsOnCall(i int, result1 fab.ChannelCfg, result2 error) {
	fake.queryConfigMutex.Lock()
	defer fake.queryConfigMutex.Unlock()
	fake.QueryConfigStub = nil
	if fake.queryConfigReturnsOnCall == nil {
		fake.queryConfigReturnsOnCall = make(map[int]struct {
			result1 fab.ChannelCfg
			result2 error
		})
	}
	fake.queryConfigReturnsOnCall[i] = struct {
		result1 fab.ChannelCfg
		result2 error
	}{result1, result2}
}

func (fake *Ledger) QueryInfo(arg1 ...ledger.RequestOption) (*fab.BlockchainInfoResponse, error) {
	fake.queryInfoMutex.Lock()
	ret, specificReturn := fake.queryInfoReturnsOnCall[len(fake.queryInfoArgsForCall)]
	fake.queryInfoArgsForCall = append(fake.queryInfoArgsForCall, struct {
		arg1 []ledger.RequestOption
	}{arg1})
	fake.recordInvocation("QueryInfo", []interface{}{arg1})
	fake.queryInfoMutex.Unlock()
	if fake.QueryInfoStub != nil {
		return fake.QueryInfoStub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryInfoReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Ledger) QueryInfoCallCount() int {
	fake.queryInfoMutex.RLock()
	defer fake.queryInfoMutex.RUnlock()
	return len(fake.queryInfoArgsForCall)
}

func (fake *Ledger) QueryInfoCalls(stub func(...ledger.RequestOption) (*fab.BlockchainInfoResponse, error)) {
	fake.queryInfoMutex.Lock()
	defer fake.queryInfoMutex.Unlock()
	fake.QueryInfoStub = stub
}

func (fake *Ledger) QueryInfoArgsForCall(i int) []ledger.RequestOption {
	fake.queryInfoMutex.RLock()
	defer fake.queryInfoMutex.RUnlock()
	argsForCall := fake.queryInfoArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Ledger) QueryInfoReturns(result1 *fab.BlockchainInfoResponse, result2 error) {
	fake.queryInfoMutex.Lock()
	defer fake.queryInfoMutex.Unlock()
	fake.QueryInfoStub = nil
	fake.queryInfoReturns = struct {
		result1 *fab.BlockchainInfoResponse
		result2 error
	}{result1, result2}
}

func (fake *Ledger) QueryInfoReturnsOnCall(i int, result1 *fab.BlockchainInfoResponse, result2 error) {
	fake.queryInfoMutex.Lock()
	defer fake.queryInfoMutex.Unlock()
	fake.QueryInfoStub = nil
	if fake.queryInfoReturnsOnCall == nil {
		fake.queryInfoReturnsOnCall = make(map[int]struct {
			result1 *fab.BlockchainInfoResponse
			result2 error
		})
	}
	fake.queryInfoReturnsOnCall[i] = struct {
		result1 *fab.BlockchainInfoResponse
		result2 error
	}{result1, result2}
}

func (fake *Ledger) QueryTransaction(arg1 fab.TransactionID, arg2 ...ledger.RequestOption) (*peer.ProcessedTransaction, error) {
	fake.queryTransactionMutex.Lock()
	ret, specificReturn := fake.queryTransactionReturnsOnCall[len(fake.queryTransactionArgsForCall)]
	fake.queryTransactionArgsForCall = append(fake.queryTransactionArgsForCall, struct {
		arg1 fab.TransactionID
		arg2 []ledger.RequestOption
	}{arg1, arg2})
	fake.recordInvocation("QueryTransaction", []interface{}{arg1, arg2})
	fake.queryTransactionMutex.Unlock()
	if fake.QueryTransactionStub != nil {
		return fake.QueryTransactionStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryTransactionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Ledger) QueryTransactionCallCount() int {
	fake.queryTransactionMutex.RLock()
	defer fake.queryTransactionMutex.RUnlock()
	return len(fake.queryTransactionArgsForCall)
}

func (fake *Ledger) QueryTransactionCalls(stub func(fab.TransactionID, ...ledger.RequestOption) (*peer.ProcessedTransaction, error)) {
	fake.queryTransactionMutex.Lock()
	defer fake.queryTransactionMutex.Unlock()
	fake.QueryTransactionStub = stub
}

func (fake *Ledger) QueryTransactionArgsForCall(i int) (fab.TransactionID, []ledger.RequestOption) {
	fake.queryTransactionMutex.RLock()
	defer fake.queryTransactionMutex.RUnlock()
	argsForCall := fake.queryTransactionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Ledger) QueryTransactionReturns(result1 *peer.ProcessedTransaction, result2 error) {
	fake.queryTransactionMutex.Lock()
	defer fake.queryTransactionMutex.Unlock()
	fake.QueryTransactionStub = nil
	fake.queryTransactionReturns = struct {
		result1 *peer.ProcessedTransaction
		result2 error
	}{result1, result2}
}

func (fake *Ledger) QueryTransactionReturnsOnCall(i int, result1 *peer.ProcessedTransaction, result2 error) {
	fake.queryTransactionMutex.Lock()
	defer fake.queryTransactionMutex.Unlock()
	fake.QueryTransactionStub = nil
	if fake.queryTransactionReturnsOnCall == nil {
		fake.queryTransactionReturnsOnCall = make(map[int]struct {
			result1 *peer.ProcessedTransaction
			result2 error
		})
	}
	fake.queryTransactionReturnsOnCall[i] = struct {
		result1 *peer.ProcessedTransaction
		result2 error
	}{result1, result2}
}

func (fake *Ledger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.queryBlockMutex.RLock()
	defer fake.queryBlockMutex.RUnlock()
	fake.queryBlockByHashMutex.RLock()
	defer fake.queryBlockByHashMutex.RUnlock()
	fake.queryBlockByTxIDMutex.RLock()
	defer fake.queryBlockByTxIDMutex.RUnlock()
	fake.queryConfigMutex.RLock()
	defer fake.queryConfigMutex.RUnlock()
	fake.queryInfoMutex.RLock()
	defer fake.queryInfoMutex.RUnlock()
	fake.queryTransactionMutex.RLock()
	defer fake.queryTransactionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Ledger) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ fabric.Ledger = new(Ledger)