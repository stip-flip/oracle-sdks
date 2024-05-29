// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ClaimLogicClaimPositionState is an auto generated low-level Go binding around an user-defined struct.
type ClaimLogicClaimPositionState struct {
	AmountToSwap *big.Int
	BotFees      *big.Int
	TickAndOwner [32]byte
	LastRound    uint64
	Round        uint64
}

// EnterPreviewResult is an auto generated low-level Go binding around an user-defined struct.
type EnterPreviewResult struct {
	SwapOut   *big.Int
	FeeAmount *big.Int
	FrAfter   *big.Int
}

// ExitPreviewResult is an auto generated low-level Go binding around an user-defined struct.
type ExitPreviewResult struct {
	SwapOut   *big.Int
	FeeAmount *big.Int
	FrAfter   *big.Int
}

// PositionInfo is an auto generated low-level Go binding around an user-defined struct.
type PositionInfo struct {
	Shares      *big.Int
	SharesRatio *big.Int
}

// PoolMetaData contains all meta data concerning the Pool contract.
var PoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"positionTick\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"sharesBurned\",\"type\":\"uint128\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"positionTick\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ClaimedBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"ClaimedEnter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"exitee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"ClaimedExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"positionTick\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ClaimedMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"positionTick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"Collect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"amountSent\",\"type\":\"uint96\"}],\"name\":\"Entered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"exitee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"sharesLocked\",\"type\":\"uint128\"}],\"name\":\"Exited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int80\",\"name\":\"fr\",\"type\":\"int80\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"positionTick\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"amountSent\",\"type\":\"uint96\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"liquidityMoved\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ORACLE_FEE\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_transferShares\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_transferSharesFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"positionTick\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"shares_\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"mints_\",\"type\":\"uint64[]\"},{\"internalType\":\"int24[]\",\"name\":\"mintTicks\",\"type\":\"int24[]\"},{\"internalType\":\"uint64[]\",\"name\":\"burns_\",\"type\":\"uint64[]\"},{\"internalType\":\"int24[]\",\"name\":\"burnTicks\",\"type\":\"int24[]\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"claimAllPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"mintees\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"burnees\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"claimFee\",\"type\":\"uint96\"}],\"name\":\"claimAllPosition\",\"outputs\":[{\"components\":[{\"internalType\":\"int96\",\"name\":\"amountToSwap\",\"type\":\"int96\"},{\"internalType\":\"uint96\",\"name\":\"botFees\",\"type\":\"uint96\"},{\"internalType\":\"bytes32\",\"name\":\"tickAndOwner\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastRound\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"internalType\":\"structClaimLogic.ClaimPositionState\",\"name\":\"state\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"entries_\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"exits_\",\"type\":\"uint64[]\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"claimAllSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"enterees\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"exitees\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"claimFee\",\"type\":\"uint96\"}],\"name\":\"claimAllSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"tickAndFrom\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"claimBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"claimEnter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"claimExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"tickAndFrom\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"claimMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"enter\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee_\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"oracle_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"oracleSlot_\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"long_\",\"type\":\"bool\"},{\"internalType\":\"enumIOracleView.Leverage\",\"name\":\"leverage_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leverage\",\"outputs\":[{\"internalType\":\"enumIOracleView.Leverage\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"long\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"positionTick\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractIOracleView\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleSlot\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_int24\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"pack\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolDebt\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"positionTick\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"position\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"shares\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"sharesRatio\",\"type\":\"uint128\"}],\"internalType\":\"structPosition.Info\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"positionTick\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"positionValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"swapIn\",\"type\":\"int256\"}],\"name\":\"previewEnter\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"swapOut\",\"type\":\"int256\"},{\"internalType\":\"uint96\",\"name\":\"feeAmount\",\"type\":\"uint96\"},{\"internalType\":\"int80\",\"name\":\"frAfter\",\"type\":\"int80\"}],\"internalType\":\"structEnter.PreviewResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"swapIn\",\"type\":\"int256\"}],\"name\":\"previewExit\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"swapOut\",\"type\":\"int256\"},{\"internalType\":\"uint96\",\"name\":\"feeAmount\",\"type\":\"uint96\"},{\"internalType\":\"int80\",\"name\":\"frAfter\",\"type\":\"int80\"}],\"internalType\":\"structExit.PreviewResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"shares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"name\":\"sharesValueWithRebalance\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slot0\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"pnl\",\"type\":\"uint128\"},{\"internalType\":\"uint96\",\"name\":\"totalLiquidities\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slot1\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"tickRatio\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightMostInitializedTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"leftMostInitializedTick\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slot2\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"totalShares\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"lastUpdate\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastPrice\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"tickValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes\",\"type\":\"bytes32\"}],\"name\":\"unpack\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolMetaData.ABI instead.
var PoolABI = PoolMetaData.ABI

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// ORACLEFEE is a free data retrieval call binding the contract method 0x31d8a985.
//
// Solidity: function ORACLE_FEE() view returns(uint24)
func (_Pool *PoolCaller) ORACLEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "ORACLE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ORACLEFEE is a free data retrieval call binding the contract method 0x31d8a985.
//
// Solidity: function ORACLE_FEE() view returns(uint24)
func (_Pool *PoolSession) ORACLEFEE() (*big.Int, error) {
	return _Pool.Contract.ORACLEFEE(&_Pool.CallOpts)
}

// ORACLEFEE is a free data retrieval call binding the contract method 0x31d8a985.
//
// Solidity: function ORACLE_FEE() view returns(uint24)
func (_Pool *PoolCallerSession) ORACLEFEE() (*big.Int, error) {
	return _Pool.Contract.ORACLEFEE(&_Pool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pool *PoolCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pool *PoolSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Pool.Contract.Allowance(&_Pool.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pool *PoolCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Pool.Contract.Allowance(&_Pool.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Pool *PoolCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Pool *PoolSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Pool.Contract.BalanceOf(&_Pool.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Pool *PoolCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Pool.Contract.BalanceOf(&_Pool.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pool *PoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pool *PoolSession) Decimals() (uint8, error) {
	return _Pool.Contract.Decimals(&_Pool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pool *PoolCallerSession) Decimals() (uint8, error) {
	return _Pool.Contract.Decimals(&_Pool.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Pool *PoolCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Pool *PoolSession) Description() (string, error) {
	return _Pool.Contract.Description(&_Pool.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Pool *PoolCallerSession) Description() (string, error) {
	return _Pool.Contract.Description(&_Pool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_Pool *PoolCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_Pool *PoolSession) Fee() (*big.Int, error) {
	return _Pool.Contract.Fee(&_Pool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_Pool *PoolCallerSession) Fee() (*big.Int, error) {
	return _Pool.Contract.Fee(&_Pool.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint64)
func (_Pool *PoolCaller) GetPrice(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getPrice")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint64)
func (_Pool *PoolSession) GetPrice() (uint64, error) {
	return _Pool.Contract.GetPrice(&_Pool.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint64)
func (_Pool *PoolCallerSession) GetPrice() (uint64, error) {
	return _Pool.Contract.GetPrice(&_Pool.CallOpts)
}

// Leverage is a free data retrieval call binding the contract method 0x2c86d98e.
//
// Solidity: function leverage() view returns(uint8)
func (_Pool *PoolCaller) Leverage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "leverage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Leverage is a free data retrieval call binding the contract method 0x2c86d98e.
//
// Solidity: function leverage() view returns(uint8)
func (_Pool *PoolSession) Leverage() (uint8, error) {
	return _Pool.Contract.Leverage(&_Pool.CallOpts)
}

// Leverage is a free data retrieval call binding the contract method 0x2c86d98e.
//
// Solidity: function leverage() view returns(uint8)
func (_Pool *PoolCallerSession) Leverage() (uint8, error) {
	return _Pool.Contract.Leverage(&_Pool.CallOpts)
}

// Long is a free data retrieval call binding the contract method 0x07bfce37.
//
// Solidity: function long() view returns(bool)
func (_Pool *PoolCaller) Long(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "long")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Long is a free data retrieval call binding the contract method 0x07bfce37.
//
// Solidity: function long() view returns(bool)
func (_Pool *PoolSession) Long() (bool, error) {
	return _Pool.Contract.Long(&_Pool.CallOpts)
}

// Long is a free data retrieval call binding the contract method 0x07bfce37.
//
// Solidity: function long() view returns(bool)
func (_Pool *PoolCallerSession) Long() (bool, error) {
	return _Pool.Contract.Long(&_Pool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pool *PoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pool *PoolSession) Name() (string, error) {
	return _Pool.Contract.Name(&_Pool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pool *PoolCallerSession) Name() (string, error) {
	return _Pool.Contract.Name(&_Pool.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Pool *PoolCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Pool *PoolSession) Oracle() (common.Address, error) {
	return _Pool.Contract.Oracle(&_Pool.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Pool *PoolCallerSession) Oracle() (common.Address, error) {
	return _Pool.Contract.Oracle(&_Pool.CallOpts)
}

// OracleSlot is a free data retrieval call binding the contract method 0x63730345.
//
// Solidity: function oracleSlot() view returns(uint8)
func (_Pool *PoolCaller) OracleSlot(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "oracleSlot")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OracleSlot is a free data retrieval call binding the contract method 0x63730345.
//
// Solidity: function oracleSlot() view returns(uint8)
func (_Pool *PoolSession) OracleSlot() (uint8, error) {
	return _Pool.Contract.OracleSlot(&_Pool.CallOpts)
}

// OracleSlot is a free data retrieval call binding the contract method 0x63730345.
//
// Solidity: function oracleSlot() view returns(uint8)
func (_Pool *PoolCallerSession) OracleSlot() (uint8, error) {
	return _Pool.Contract.OracleSlot(&_Pool.CallOpts)
}

// Pack is a free data retrieval call binding the contract method 0x34a55cfa.
//
// Solidity: function pack(int24 _int24, address _address) pure returns(bytes32)
func (_Pool *PoolCaller) Pack(opts *bind.CallOpts, _int24 *big.Int, _address common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "pack", _int24, _address)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Pack is a free data retrieval call binding the contract method 0x34a55cfa.
//
// Solidity: function pack(int24 _int24, address _address) pure returns(bytes32)
func (_Pool *PoolSession) Pack(_int24 *big.Int, _address common.Address) ([32]byte, error) {
	return _Pool.Contract.Pack(&_Pool.CallOpts, _int24, _address)
}

// Pack is a free data retrieval call binding the contract method 0x34a55cfa.
//
// Solidity: function pack(int24 _int24, address _address) pure returns(bytes32)
func (_Pool *PoolCallerSession) Pack(_int24 *big.Int, _address common.Address) ([32]byte, error) {
	return _Pool.Contract.Pack(&_Pool.CallOpts, _int24, _address)
}

// PoolDebt is a free data retrieval call binding the contract method 0xade1a170.
//
// Solidity: function poolDebt() view returns(uint96)
func (_Pool *PoolCaller) PoolDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "poolDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolDebt is a free data retrieval call binding the contract method 0xade1a170.
//
// Solidity: function poolDebt() view returns(uint96)
func (_Pool *PoolSession) PoolDebt() (*big.Int, error) {
	return _Pool.Contract.PoolDebt(&_Pool.CallOpts)
}

// PoolDebt is a free data retrieval call binding the contract method 0xade1a170.
//
// Solidity: function poolDebt() view returns(uint96)
func (_Pool *PoolCallerSession) PoolDebt() (*big.Int, error) {
	return _Pool.Contract.PoolDebt(&_Pool.CallOpts)
}

// Position is a free data retrieval call binding the contract method 0x97fc1326.
//
// Solidity: function position(int24 positionTick, address owner) view returns((uint128,uint128))
func (_Pool *PoolCaller) Position(opts *bind.CallOpts, positionTick *big.Int, owner common.Address) (PositionInfo, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "position", positionTick, owner)

	if err != nil {
		return *new(PositionInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PositionInfo)).(*PositionInfo)

	return out0, err

}

// Position is a free data retrieval call binding the contract method 0x97fc1326.
//
// Solidity: function position(int24 positionTick, address owner) view returns((uint128,uint128))
func (_Pool *PoolSession) Position(positionTick *big.Int, owner common.Address) (PositionInfo, error) {
	return _Pool.Contract.Position(&_Pool.CallOpts, positionTick, owner)
}

// Position is a free data retrieval call binding the contract method 0x97fc1326.
//
// Solidity: function position(int24 positionTick, address owner) view returns((uint128,uint128))
func (_Pool *PoolCallerSession) Position(positionTick *big.Int, owner common.Address) (PositionInfo, error) {
	return _Pool.Contract.Position(&_Pool.CallOpts, positionTick, owner)
}

// PositionValue is a free data retrieval call binding the contract method 0x1f7a7002.
//
// Solidity: function positionValue(int24 positionTick, address owner) view returns(uint256)
func (_Pool *PoolCaller) PositionValue(opts *bind.CallOpts, positionTick *big.Int, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "positionValue", positionTick, owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PositionValue is a free data retrieval call binding the contract method 0x1f7a7002.
//
// Solidity: function positionValue(int24 positionTick, address owner) view returns(uint256)
func (_Pool *PoolSession) PositionValue(positionTick *big.Int, owner common.Address) (*big.Int, error) {
	return _Pool.Contract.PositionValue(&_Pool.CallOpts, positionTick, owner)
}

// PositionValue is a free data retrieval call binding the contract method 0x1f7a7002.
//
// Solidity: function positionValue(int24 positionTick, address owner) view returns(uint256)
func (_Pool *PoolCallerSession) PositionValue(positionTick *big.Int, owner common.Address) (*big.Int, error) {
	return _Pool.Contract.PositionValue(&_Pool.CallOpts, positionTick, owner)
}

// PreviewEnter is a free data retrieval call binding the contract method 0xac205b43.
//
// Solidity: function previewEnter(int256 swapIn) view returns((int256,uint96,int80))
func (_Pool *PoolCaller) PreviewEnter(opts *bind.CallOpts, swapIn *big.Int) (EnterPreviewResult, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "previewEnter", swapIn)

	if err != nil {
		return *new(EnterPreviewResult), err
	}

	out0 := *abi.ConvertType(out[0], new(EnterPreviewResult)).(*EnterPreviewResult)

	return out0, err

}

// PreviewEnter is a free data retrieval call binding the contract method 0xac205b43.
//
// Solidity: function previewEnter(int256 swapIn) view returns((int256,uint96,int80))
func (_Pool *PoolSession) PreviewEnter(swapIn *big.Int) (EnterPreviewResult, error) {
	return _Pool.Contract.PreviewEnter(&_Pool.CallOpts, swapIn)
}

// PreviewEnter is a free data retrieval call binding the contract method 0xac205b43.
//
// Solidity: function previewEnter(int256 swapIn) view returns((int256,uint96,int80))
func (_Pool *PoolCallerSession) PreviewEnter(swapIn *big.Int) (EnterPreviewResult, error) {
	return _Pool.Contract.PreviewEnter(&_Pool.CallOpts, swapIn)
}

// PreviewExit is a free data retrieval call binding the contract method 0xb1eda5dc.
//
// Solidity: function previewExit(int256 swapIn) view returns((int256,uint96,int80))
func (_Pool *PoolCaller) PreviewExit(opts *bind.CallOpts, swapIn *big.Int) (ExitPreviewResult, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "previewExit", swapIn)

	if err != nil {
		return *new(ExitPreviewResult), err
	}

	out0 := *abi.ConvertType(out[0], new(ExitPreviewResult)).(*ExitPreviewResult)

	return out0, err

}

// PreviewExit is a free data retrieval call binding the contract method 0xb1eda5dc.
//
// Solidity: function previewExit(int256 swapIn) view returns((int256,uint96,int80))
func (_Pool *PoolSession) PreviewExit(swapIn *big.Int) (ExitPreviewResult, error) {
	return _Pool.Contract.PreviewExit(&_Pool.CallOpts, swapIn)
}

// PreviewExit is a free data retrieval call binding the contract method 0xb1eda5dc.
//
// Solidity: function previewExit(int256 swapIn) view returns((int256,uint96,int80))
func (_Pool *PoolCallerSession) PreviewExit(swapIn *big.Int) (ExitPreviewResult, error) {
	return _Pool.Contract.PreviewExit(&_Pool.CallOpts, swapIn)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address ) view returns(uint256)
func (_Pool *PoolCaller) Shares(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "shares", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address ) view returns(uint256)
func (_Pool *PoolSession) Shares(arg0 common.Address) (*big.Int, error) {
	return _Pool.Contract.Shares(&_Pool.CallOpts, arg0)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address ) view returns(uint256)
func (_Pool *PoolCallerSession) Shares(arg0 common.Address) (*big.Int, error) {
	return _Pool.Contract.Shares(&_Pool.CallOpts, arg0)
}

// SharesValueWithRebalance is a free data retrieval call binding the contract method 0xe71a4dd0.
//
// Solidity: function sharesValueWithRebalance(uint256 shares_) view returns(uint96)
func (_Pool *PoolCaller) SharesValueWithRebalance(opts *bind.CallOpts, shares_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "sharesValueWithRebalance", shares_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesValueWithRebalance is a free data retrieval call binding the contract method 0xe71a4dd0.
//
// Solidity: function sharesValueWithRebalance(uint256 shares_) view returns(uint96)
func (_Pool *PoolSession) SharesValueWithRebalance(shares_ *big.Int) (*big.Int, error) {
	return _Pool.Contract.SharesValueWithRebalance(&_Pool.CallOpts, shares_)
}

// SharesValueWithRebalance is a free data retrieval call binding the contract method 0xe71a4dd0.
//
// Solidity: function sharesValueWithRebalance(uint256 shares_) view returns(uint96)
func (_Pool *PoolCallerSession) SharesValueWithRebalance(shares_ *big.Int) (*big.Int, error) {
	return _Pool.Contract.SharesValueWithRebalance(&_Pool.CallOpts, shares_)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint128 pnl, uint96 totalLiquidities)
func (_Pool *PoolCaller) Slot0(opts *bind.CallOpts) (struct {
	Pnl              *big.Int
	TotalLiquidities *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "slot0")

	outstruct := new(struct {
		Pnl              *big.Int
		TotalLiquidities *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pnl = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalLiquidities = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint128 pnl, uint96 totalLiquidities)
func (_Pool *PoolSession) Slot0() (struct {
	Pnl              *big.Int
	TotalLiquidities *big.Int
}, error) {
	return _Pool.Contract.Slot0(&_Pool.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint128 pnl, uint96 totalLiquidities)
func (_Pool *PoolCallerSession) Slot0() (struct {
	Pnl              *big.Int
	TotalLiquidities *big.Int
}, error) {
	return _Pool.Contract.Slot0(&_Pool.CallOpts)
}

// Slot1 is a free data retrieval call binding the contract method 0x1f457cb5.
//
// Solidity: function slot1() view returns(uint128 tickRatio, int24 tick, int24 rightMostInitializedTick, int24 leftMostInitializedTick)
func (_Pool *PoolCaller) Slot1(opts *bind.CallOpts) (struct {
	TickRatio                *big.Int
	Tick                     *big.Int
	RightMostInitializedTick *big.Int
	LeftMostInitializedTick  *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "slot1")

	outstruct := new(struct {
		TickRatio                *big.Int
		Tick                     *big.Int
		RightMostInitializedTick *big.Int
		LeftMostInitializedTick  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TickRatio = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RightMostInitializedTick = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LeftMostInitializedTick = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Slot1 is a free data retrieval call binding the contract method 0x1f457cb5.
//
// Solidity: function slot1() view returns(uint128 tickRatio, int24 tick, int24 rightMostInitializedTick, int24 leftMostInitializedTick)
func (_Pool *PoolSession) Slot1() (struct {
	TickRatio                *big.Int
	Tick                     *big.Int
	RightMostInitializedTick *big.Int
	LeftMostInitializedTick  *big.Int
}, error) {
	return _Pool.Contract.Slot1(&_Pool.CallOpts)
}

// Slot1 is a free data retrieval call binding the contract method 0x1f457cb5.
//
// Solidity: function slot1() view returns(uint128 tickRatio, int24 tick, int24 rightMostInitializedTick, int24 leftMostInitializedTick)
func (_Pool *PoolCallerSession) Slot1() (struct {
	TickRatio                *big.Int
	Tick                     *big.Int
	RightMostInitializedTick *big.Int
	LeftMostInitializedTick  *big.Int
}, error) {
	return _Pool.Contract.Slot1(&_Pool.CallOpts)
}

// Slot2 is a free data retrieval call binding the contract method 0xd987e6b5.
//
// Solidity: function slot2() view returns(uint128 totalShares, uint64 lastUpdate, uint64 lastPrice)
func (_Pool *PoolCaller) Slot2(opts *bind.CallOpts) (struct {
	TotalShares *big.Int
	LastUpdate  uint64
	LastPrice   uint64
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "slot2")

	outstruct := new(struct {
		TotalShares *big.Int
		LastUpdate  uint64
		LastPrice   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalShares = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastUpdate = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.LastPrice = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Slot2 is a free data retrieval call binding the contract method 0xd987e6b5.
//
// Solidity: function slot2() view returns(uint128 totalShares, uint64 lastUpdate, uint64 lastPrice)
func (_Pool *PoolSession) Slot2() (struct {
	TotalShares *big.Int
	LastUpdate  uint64
	LastPrice   uint64
}, error) {
	return _Pool.Contract.Slot2(&_Pool.CallOpts)
}

// Slot2 is a free data retrieval call binding the contract method 0xd987e6b5.
//
// Solidity: function slot2() view returns(uint128 totalShares, uint64 lastUpdate, uint64 lastPrice)
func (_Pool *PoolCallerSession) Slot2() (struct {
	TotalShares *big.Int
	LastUpdate  uint64
	LastPrice   uint64
}, error) {
	return _Pool.Contract.Slot2(&_Pool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pool *PoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pool *PoolSession) Symbol() (string, error) {
	return _Pool.Contract.Symbol(&_Pool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pool *PoolCallerSession) Symbol() (string, error) {
	return _Pool.Contract.Symbol(&_Pool.CallOpts)
}

// TickValue is a free data retrieval call binding the contract method 0x3f63254d.
//
// Solidity: function tickValue(int24 tick) view returns(uint256)
func (_Pool *PoolCaller) TickValue(opts *bind.CallOpts, tick *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "tickValue", tick)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickValue is a free data retrieval call binding the contract method 0x3f63254d.
//
// Solidity: function tickValue(int24 tick) view returns(uint256)
func (_Pool *PoolSession) TickValue(tick *big.Int) (*big.Int, error) {
	return _Pool.Contract.TickValue(&_Pool.CallOpts, tick)
}

// TickValue is a free data retrieval call binding the contract method 0x3f63254d.
//
// Solidity: function tickValue(int24 tick) view returns(uint256)
func (_Pool *PoolCallerSession) TickValue(tick *big.Int) (*big.Int, error) {
	return _Pool.Contract.TickValue(&_Pool.CallOpts, tick)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pool *PoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pool *PoolSession) TotalSupply() (*big.Int, error) {
	return _Pool.Contract.TotalSupply(&_Pool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pool *PoolCallerSession) TotalSupply() (*big.Int, error) {
	return _Pool.Contract.TotalSupply(&_Pool.CallOpts)
}

// Unpack is a free data retrieval call binding the contract method 0x71516dd9.
//
// Solidity: function unpack(bytes32 _bytes) pure returns(int24, address)
func (_Pool *PoolCaller) Unpack(opts *bind.CallOpts, _bytes [32]byte) (*big.Int, common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "unpack", _bytes)

	if err != nil {
		return *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// Unpack is a free data retrieval call binding the contract method 0x71516dd9.
//
// Solidity: function unpack(bytes32 _bytes) pure returns(int24, address)
func (_Pool *PoolSession) Unpack(_bytes [32]byte) (*big.Int, common.Address, error) {
	return _Pool.Contract.Unpack(&_Pool.CallOpts, _bytes)
}

// Unpack is a free data retrieval call binding the contract method 0x71516dd9.
//
// Solidity: function unpack(bytes32 _bytes) pure returns(int24, address)
func (_Pool *PoolCallerSession) Unpack(_bytes [32]byte) (*big.Int, common.Address, error) {
	return _Pool.Contract.Unpack(&_Pool.CallOpts, _bytes)
}

// TransferShares is a paid mutator transaction binding the contract method 0x2cf036ce.
//
// Solidity: function _transferShares(address to, uint256 amount) returns(bool)
func (_Pool *PoolTransactor) TransferShares(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "_transferShares", to, amount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x2cf036ce.
//
// Solidity: function _transferShares(address to, uint256 amount) returns(bool)
func (_Pool *PoolSession) TransferShares(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TransferShares(&_Pool.TransactOpts, to, amount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x2cf036ce.
//
// Solidity: function _transferShares(address to, uint256 amount) returns(bool)
func (_Pool *PoolTransactorSession) TransferShares(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TransferShares(&_Pool.TransactOpts, to, amount)
}

// TransferSharesFrom is a paid mutator transaction binding the contract method 0xf880c8d8.
//
// Solidity: function _transferSharesFrom(address from, address to, uint256 amount) returns(bool)
func (_Pool *PoolTransactor) TransferSharesFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "_transferSharesFrom", from, to, amount)
}

// TransferSharesFrom is a paid mutator transaction binding the contract method 0xf880c8d8.
//
// Solidity: function _transferSharesFrom(address from, address to, uint256 amount) returns(bool)
func (_Pool *PoolSession) TransferSharesFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TransferSharesFrom(&_Pool.TransactOpts, from, to, amount)
}

// TransferSharesFrom is a paid mutator transaction binding the contract method 0xf880c8d8.
//
// Solidity: function _transferSharesFrom(address from, address to, uint256 amount) returns(bool)
func (_Pool *PoolTransactorSession) TransferSharesFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TransferSharesFrom(&_Pool.TransactOpts, from, to, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Pool *PoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Pool *PoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Approve(&_Pool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Pool *PoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Approve(&_Pool.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x062df7a3.
//
// Solidity: function burn(int24 positionTick, uint128 shares_, address claimer) returns()
func (_Pool *PoolTransactor) Burn(opts *bind.TransactOpts, positionTick *big.Int, shares_ *big.Int, claimer common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "burn", positionTick, shares_, claimer)
}

// Burn is a paid mutator transaction binding the contract method 0x062df7a3.
//
// Solidity: function burn(int24 positionTick, uint128 shares_, address claimer) returns()
func (_Pool *PoolSession) Burn(positionTick *big.Int, shares_ *big.Int, claimer common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Burn(&_Pool.TransactOpts, positionTick, shares_, claimer)
}

// Burn is a paid mutator transaction binding the contract method 0x062df7a3.
//
// Solidity: function burn(int24 positionTick, uint128 shares_, address claimer) returns()
func (_Pool *PoolTransactorSession) Burn(positionTick *big.Int, shares_ *big.Int, claimer common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Burn(&_Pool.TransactOpts, positionTick, shares_, claimer)
}

// ClaimAllPosition is a paid mutator transaction binding the contract method 0x009de35e.
//
// Solidity: function claimAllPosition(uint64[] mints_, int24[] mintTicks, uint64[] burns_, int24[] burnTicks, address recipient) returns()
func (_Pool *PoolTransactor) ClaimAllPosition(opts *bind.TransactOpts, mints_ []uint64, mintTicks []*big.Int, burns_ []uint64, burnTicks []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "claimAllPosition", mints_, mintTicks, burns_, burnTicks, recipient)
}

// ClaimAllPosition is a paid mutator transaction binding the contract method 0x009de35e.
//
// Solidity: function claimAllPosition(uint64[] mints_, int24[] mintTicks, uint64[] burns_, int24[] burnTicks, address recipient) returns()
func (_Pool *PoolSession) ClaimAllPosition(mints_ []uint64, mintTicks []*big.Int, burns_ []uint64, burnTicks []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ClaimAllPosition(&_Pool.TransactOpts, mints_, mintTicks, burns_, burnTicks, recipient)
}

// ClaimAllPosition is a paid mutator transaction binding the contract method 0x009de35e.
//
// Solidity: function claimAllPosition(uint64[] mints_, int24[] mintTicks, uint64[] burns_, int24[] burnTicks, address recipient) returns()
func (_Pool *PoolTransactorSession) ClaimAllPosition(mints_ []uint64, mintTicks []*big.Int, burns_ []uint64, burnTicks []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ClaimAllPosition(&_Pool.TransactOpts, mints_, mintTicks, burns_, burnTicks, recipient)
}

// ClaimAllPosition0 is a paid mutator transaction binding the contract method 0xa3055dca.
//
// Solidity: function claimAllPosition(bytes32[] mintees, bytes32[] burnees, uint64 round, uint96 claimFee) returns((int96,uint96,bytes32,uint64,uint64) state)
func (_Pool *PoolTransactor) ClaimAllPosition0(opts *bind.TransactOpts, mintees [][32]byte, burnees [][32]byte, round uint64, claimFee *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "claimAllPosition0", mintees, burnees, round, claimFee)
}

// ClaimAllPosition0 is a paid mutator transaction binding the contract method 0xa3055dca.
//
// Solidity: function claimAllPosition(bytes32[] mintees, bytes32[] burnees, uint64 round, uint96 claimFee) returns((int96,uint96,bytes32,uint64,uint64) state)
func (_Pool *PoolSession) ClaimAllPosition0(mintees [][32]byte, burnees [][32]byte, round uint64, claimFee *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ClaimAllPosition0(&_Pool.TransactOpts, mintees, burnees, round, claimFee)
}

// ClaimAllPosition0 is a paid mutator transaction binding the contract method 0xa3055dca.
//
// Solidity: function claimAllPosition(bytes32[] mintees, bytes32[] burnees, uint64 round, uint96 claimFee) returns((int96,uint96,bytes32,uint64,uint64) state)
func (_Pool *PoolTransactorSession) ClaimAllPosition0(mintees [][32]byte, burnees [][32]byte, round uint64, claimFee *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ClaimAllPosition0(&_Pool.TransactOpts, mintees, burnees, round, claimFee)
}

// ClaimAllSwap is a paid mutator transaction binding the contract method 0x809a89c1.
//
// Solidity: function claimAllSwap(uint64[] entries_, uint64[] exits_, address recipient) returns()
func (_Pool *PoolTransactor) ClaimAllSwap(opts *bind.TransactOpts, entries_ []uint64, exits_ []uint64, recipient common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "claimAllSwap", entries_, exits_, recipient)
}

// ClaimAllSwap is a paid mutator transaction binding the contract method 0x809a89c1.
//
// Solidity: function claimAllSwap(uint64[] entries_, uint64[] exits_, address recipient) returns()
func (_Pool *PoolSession) ClaimAllSwap(entries_ []uint64, exits_ []uint64, recipient common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ClaimAllSwap(&_Pool.TransactOpts, entries_, exits_, recipient)
}

// ClaimAllSwap is a paid mutator transaction binding the contract method 0x809a89c1.
//
// Solidity: function claimAllSwap(uint64[] entries_, uint64[] exits_, address recipient) returns()
func (_Pool *PoolTransactorSession) ClaimAllSwap(entries_ []uint64, exits_ []uint64, recipient common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ClaimAllSwap(&_Pool.TransactOpts, entries_, exits_, recipient)
}

// ClaimAllSwap0 is a paid mutator transaction binding the contract method 0xd1f1be0c.
//
// Solidity: function claimAllSwap(address[] enterees, address[] exitees, uint64 round, uint96 claimFee) returns()
func (_Pool *PoolTransactor) ClaimAllSwap0(opts *bind.TransactOpts, enterees []common.Address, exitees []common.Address, round uint64, claimFee *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "claimAllSwap0", enterees, exitees, round, claimFee)
}

// ClaimAllSwap0 is a paid mutator transaction binding the contract method 0xd1f1be0c.
//
// Solidity: function claimAllSwap(address[] enterees, address[] exitees, uint64 round, uint96 claimFee) returns()
func (_Pool *PoolSession) ClaimAllSwap0(enterees []common.Address, exitees []common.Address, round uint64, claimFee *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ClaimAllSwap0(&_Pool.TransactOpts, enterees, exitees, round, claimFee)
}

// ClaimAllSwap0 is a paid mutator transaction binding the contract method 0xd1f1be0c.
//
// Solidity: function claimAllSwap(address[] enterees, address[] exitees, uint64 round, uint96 claimFee) returns()
func (_Pool *PoolTransactorSession) ClaimAllSwap0(enterees []common.Address, exitees []common.Address, round uint64, claimFee *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ClaimAllSwap0(&_Pool.TransactOpts, enterees, exitees, round, claimFee)
}

// ClaimBurn is a paid mutator transaction binding the contract method 0x9fb90520.
//
// Solidity: function claimBurn(uint64 round, bytes32 tickAndFrom, address recipient) returns()
func (_Pool *PoolTransactor) ClaimBurn(opts *bind.TransactOpts, round uint64, tickAndFrom [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "claimBurn", round, tickAndFrom, recipient)
}

// ClaimBurn is a paid mutator transaction binding the contract method 0x9fb90520.
//
// Solidity: function claimBurn(uint64 round, bytes32 tickAndFrom, address recipient) returns()
func (_Pool *PoolSession) ClaimBurn(round uint64, tickAndFrom [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ClaimBurn(&_Pool.TransactOpts, round, tickAndFrom, recipient)
}

// ClaimBurn is a paid mutator transaction binding the contract method 0x9fb90520.
//
// Solidity: function claimBurn(uint64 round, bytes32 tickAndFrom, address recipient) returns()
func (_Pool *PoolTransactorSession) ClaimBurn(round uint64, tickAndFrom [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ClaimBurn(&_Pool.TransactOpts, round, tickAndFrom, recipient)
}

// ClaimEnter is a paid mutator transaction binding the contract method 0x02f1b248.
//
// Solidity: function claimEnter(uint64 round, address recipient) returns()
func (_Pool *PoolTransactor) ClaimEnter(opts *bind.TransactOpts, round uint64, recipient common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "claimEnter", round, recipient)
}

// ClaimEnter is a paid mutator transaction binding the contract method 0x02f1b248.
//
// Solidity: function claimEnter(uint64 round, address recipient) returns()
func (_Pool *PoolSession) ClaimEnter(round uint64, recipient common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ClaimEnter(&_Pool.TransactOpts, round, recipient)
}

// ClaimEnter is a paid mutator transaction binding the contract method 0x02f1b248.
//
// Solidity: function claimEnter(uint64 round, address recipient) returns()
func (_Pool *PoolTransactorSession) ClaimEnter(round uint64, recipient common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ClaimEnter(&_Pool.TransactOpts, round, recipient)
}

// ClaimExit is a paid mutator transaction binding the contract method 0x2a9e4900.
//
// Solidity: function claimExit(uint64 round, address recipient) returns()
func (_Pool *PoolTransactor) ClaimExit(opts *bind.TransactOpts, round uint64, recipient common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "claimExit", round, recipient)
}

// ClaimExit is a paid mutator transaction binding the contract method 0x2a9e4900.
//
// Solidity: function claimExit(uint64 round, address recipient) returns()
func (_Pool *PoolSession) ClaimExit(round uint64, recipient common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ClaimExit(&_Pool.TransactOpts, round, recipient)
}

// ClaimExit is a paid mutator transaction binding the contract method 0x2a9e4900.
//
// Solidity: function claimExit(uint64 round, address recipient) returns()
func (_Pool *PoolTransactorSession) ClaimExit(round uint64, recipient common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ClaimExit(&_Pool.TransactOpts, round, recipient)
}

// ClaimMint is a paid mutator transaction binding the contract method 0x28ab1a2e.
//
// Solidity: function claimMint(uint64 round, bytes32 tickAndFrom, address recipient) returns()
func (_Pool *PoolTransactor) ClaimMint(opts *bind.TransactOpts, round uint64, tickAndFrom [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "claimMint", round, tickAndFrom, recipient)
}

// ClaimMint is a paid mutator transaction binding the contract method 0x28ab1a2e.
//
// Solidity: function claimMint(uint64 round, bytes32 tickAndFrom, address recipient) returns()
func (_Pool *PoolSession) ClaimMint(round uint64, tickAndFrom [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ClaimMint(&_Pool.TransactOpts, round, tickAndFrom, recipient)
}

// ClaimMint is a paid mutator transaction binding the contract method 0x28ab1a2e.
//
// Solidity: function claimMint(uint64 round, bytes32 tickAndFrom, address recipient) returns()
func (_Pool *PoolTransactorSession) ClaimMint(round uint64, tickAndFrom [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ClaimMint(&_Pool.TransactOpts, round, tickAndFrom, recipient)
}

// Enter is a paid mutator transaction binding the contract method 0xd014c01f.
//
// Solidity: function enter(address claimer) payable returns()
func (_Pool *PoolTransactor) Enter(opts *bind.TransactOpts, claimer common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "enter", claimer)
}

// Enter is a paid mutator transaction binding the contract method 0xd014c01f.
//
// Solidity: function enter(address claimer) payable returns()
func (_Pool *PoolSession) Enter(claimer common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Enter(&_Pool.TransactOpts, claimer)
}

// Enter is a paid mutator transaction binding the contract method 0xd014c01f.
//
// Solidity: function enter(address claimer) payable returns()
func (_Pool *PoolTransactorSession) Enter(claimer common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Enter(&_Pool.TransactOpts, claimer)
}

// Exit is a paid mutator transaction binding the contract method 0xcff40759.
//
// Solidity: function exit(uint256 amount, address claimer) returns()
func (_Pool *PoolTransactor) Exit(opts *bind.TransactOpts, amount *big.Int, claimer common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "exit", amount, claimer)
}

// Exit is a paid mutator transaction binding the contract method 0xcff40759.
//
// Solidity: function exit(uint256 amount, address claimer) returns()
func (_Pool *PoolSession) Exit(amount *big.Int, claimer common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Exit(&_Pool.TransactOpts, amount, claimer)
}

// Exit is a paid mutator transaction binding the contract method 0xcff40759.
//
// Solidity: function exit(uint256 amount, address claimer) returns()
func (_Pool *PoolTransactorSession) Exit(amount *big.Int, claimer common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Exit(&_Pool.TransactOpts, amount, claimer)
}

// Initialize is a paid mutator transaction binding the contract method 0x18920773.
//
// Solidity: function initialize(uint24 fee_, address oracle_, uint8 oracleSlot_, string name_, string symbol_, string description_, bool long_, uint8 leverage_) returns()
func (_Pool *PoolTransactor) Initialize(opts *bind.TransactOpts, fee_ *big.Int, oracle_ common.Address, oracleSlot_ uint8, name_ string, symbol_ string, description_ string, long_ bool, leverage_ uint8) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "initialize", fee_, oracle_, oracleSlot_, name_, symbol_, description_, long_, leverage_)
}

// Initialize is a paid mutator transaction binding the contract method 0x18920773.
//
// Solidity: function initialize(uint24 fee_, address oracle_, uint8 oracleSlot_, string name_, string symbol_, string description_, bool long_, uint8 leverage_) returns()
func (_Pool *PoolSession) Initialize(fee_ *big.Int, oracle_ common.Address, oracleSlot_ uint8, name_ string, symbol_ string, description_ string, long_ bool, leverage_ uint8) (*types.Transaction, error) {
	return _Pool.Contract.Initialize(&_Pool.TransactOpts, fee_, oracle_, oracleSlot_, name_, symbol_, description_, long_, leverage_)
}

// Initialize is a paid mutator transaction binding the contract method 0x18920773.
//
// Solidity: function initialize(uint24 fee_, address oracle_, uint8 oracleSlot_, string name_, string symbol_, string description_, bool long_, uint8 leverage_) returns()
func (_Pool *PoolTransactorSession) Initialize(fee_ *big.Int, oracle_ common.Address, oracleSlot_ uint8, name_ string, symbol_ string, description_ string, long_ bool, leverage_ uint8) (*types.Transaction, error) {
	return _Pool.Contract.Initialize(&_Pool.TransactOpts, fee_, oracle_, oracleSlot_, name_, symbol_, description_, long_, leverage_)
}

// Mint is a paid mutator transaction binding the contract method 0xa48d5ea5.
//
// Solidity: function mint(int24 positionTick, address claimer) payable returns()
func (_Pool *PoolTransactor) Mint(opts *bind.TransactOpts, positionTick *big.Int, claimer common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "mint", positionTick, claimer)
}

// Mint is a paid mutator transaction binding the contract method 0xa48d5ea5.
//
// Solidity: function mint(int24 positionTick, address claimer) payable returns()
func (_Pool *PoolSession) Mint(positionTick *big.Int, claimer common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Mint(&_Pool.TransactOpts, positionTick, claimer)
}

// Mint is a paid mutator transaction binding the contract method 0xa48d5ea5.
//
// Solidity: function mint(int24 positionTick, address claimer) payable returns()
func (_Pool *PoolTransactorSession) Mint(positionTick *big.Int, claimer common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Mint(&_Pool.TransactOpts, positionTick, claimer)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pool *PoolTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pool *PoolSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pool.Contract.Permit(&_Pool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pool *PoolTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pool.Contract.Permit(&_Pool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Rebalance is a paid mutator transaction binding the contract method 0x7d7c2a1c.
//
// Solidity: function rebalance() returns()
func (_Pool *PoolTransactor) Rebalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "rebalance")
}

// Rebalance is a paid mutator transaction binding the contract method 0x7d7c2a1c.
//
// Solidity: function rebalance() returns()
func (_Pool *PoolSession) Rebalance() (*types.Transaction, error) {
	return _Pool.Contract.Rebalance(&_Pool.TransactOpts)
}

// Rebalance is a paid mutator transaction binding the contract method 0x7d7c2a1c.
//
// Solidity: function rebalance() returns()
func (_Pool *PoolTransactorSession) Rebalance() (*types.Transaction, error) {
	return _Pool.Contract.Rebalance(&_Pool.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Pool *PoolTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Pool *PoolSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Transfer(&_Pool.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Pool *PoolTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Transfer(&_Pool.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Pool *PoolTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Pool *PoolSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TransferFrom(&_Pool.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Pool *PoolTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TransferFrom(&_Pool.TransactOpts, sender, recipient, amount)
}

// PoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Pool contract.
type PoolApprovalIterator struct {
	Event *PoolApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolApproval represents a Approval event raised by the Pool contract.
type PoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pool *PoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PoolApprovalIterator{contract: _Pool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pool *PoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolApproval)
				if err := _Pool.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pool *PoolFilterer) ParseApproval(log types.Log) (*PoolApproval, error) {
	event := new(PoolApproval)
	if err := _Pool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Pool contract.
type PoolBurnIterator struct {
	Event *PoolBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBurn represents a Burn event raised by the Pool contract.
type PoolBurn struct {
	Owner        common.Address
	PositionTick *big.Int
	Round        uint64
	Claimer      common.Address
	SharesBurned *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x368cba1582afce51c89331bd9da0af5f43f52ff2224dbb03744beb7d22e08edb.
//
// Solidity: event Burn(address indexed owner, int24 indexed positionTick, uint64 indexed round, address claimer, uint128 sharesBurned)
func (_Pool *PoolFilterer) FilterBurn(opts *bind.FilterOpts, owner []common.Address, positionTick []*big.Int, round []uint64) (*PoolBurnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var positionTickRule []interface{}
	for _, positionTickItem := range positionTick {
		positionTickRule = append(positionTickRule, positionTickItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Burn", ownerRule, positionTickRule, roundRule)
	if err != nil {
		return nil, err
	}
	return &PoolBurnIterator{contract: _Pool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x368cba1582afce51c89331bd9da0af5f43f52ff2224dbb03744beb7d22e08edb.
//
// Solidity: event Burn(address indexed owner, int24 indexed positionTick, uint64 indexed round, address claimer, uint128 sharesBurned)
func (_Pool *PoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *PoolBurn, owner []common.Address, positionTick []*big.Int, round []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var positionTickRule []interface{}
	for _, positionTickItem := range positionTick {
		positionTickRule = append(positionTickRule, positionTickItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Burn", ownerRule, positionTickRule, roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBurn)
				if err := _Pool.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurn is a log parse operation binding the contract event 0x368cba1582afce51c89331bd9da0af5f43f52ff2224dbb03744beb7d22e08edb.
//
// Solidity: event Burn(address indexed owner, int24 indexed positionTick, uint64 indexed round, address claimer, uint128 sharesBurned)
func (_Pool *PoolFilterer) ParseBurn(log types.Log) (*PoolBurn, error) {
	event := new(PoolBurn)
	if err := _Pool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolClaimedBurnIterator is returned from FilterClaimedBurn and is used to iterate over the raw logs and unpacked data for ClaimedBurn events raised by the Pool contract.
type PoolClaimedBurnIterator struct {
	Event *PoolClaimedBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolClaimedBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolClaimedBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolClaimedBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolClaimedBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolClaimedBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolClaimedBurn represents a ClaimedBurn event raised by the Pool contract.
type PoolClaimedBurn struct {
	Owner        common.Address
	PositionTick *big.Int
	Round        uint64
	Recipient    common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClaimedBurn is a free log retrieval operation binding the contract event 0xb7724ad1d7a15a6494cb008af8b8b7860a8b97f5c6dc1b75eda9135bbb1e857f.
//
// Solidity: event ClaimedBurn(address indexed owner, int24 indexed positionTick, uint64 indexed round, address recipient)
func (_Pool *PoolFilterer) FilterClaimedBurn(opts *bind.FilterOpts, owner []common.Address, positionTick []*big.Int, round []uint64) (*PoolClaimedBurnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var positionTickRule []interface{}
	for _, positionTickItem := range positionTick {
		positionTickRule = append(positionTickRule, positionTickItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "ClaimedBurn", ownerRule, positionTickRule, roundRule)
	if err != nil {
		return nil, err
	}
	return &PoolClaimedBurnIterator{contract: _Pool.contract, event: "ClaimedBurn", logs: logs, sub: sub}, nil
}

// WatchClaimedBurn is a free log subscription operation binding the contract event 0xb7724ad1d7a15a6494cb008af8b8b7860a8b97f5c6dc1b75eda9135bbb1e857f.
//
// Solidity: event ClaimedBurn(address indexed owner, int24 indexed positionTick, uint64 indexed round, address recipient)
func (_Pool *PoolFilterer) WatchClaimedBurn(opts *bind.WatchOpts, sink chan<- *PoolClaimedBurn, owner []common.Address, positionTick []*big.Int, round []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var positionTickRule []interface{}
	for _, positionTickItem := range positionTick {
		positionTickRule = append(positionTickRule, positionTickItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "ClaimedBurn", ownerRule, positionTickRule, roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolClaimedBurn)
				if err := _Pool.contract.UnpackLog(event, "ClaimedBurn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimedBurn is a log parse operation binding the contract event 0xb7724ad1d7a15a6494cb008af8b8b7860a8b97f5c6dc1b75eda9135bbb1e857f.
//
// Solidity: event ClaimedBurn(address indexed owner, int24 indexed positionTick, uint64 indexed round, address recipient)
func (_Pool *PoolFilterer) ParseClaimedBurn(log types.Log) (*PoolClaimedBurn, error) {
	event := new(PoolClaimedBurn)
	if err := _Pool.contract.UnpackLog(event, "ClaimedBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolClaimedEnterIterator is returned from FilterClaimedEnter and is used to iterate over the raw logs and unpacked data for ClaimedEnter events raised by the Pool contract.
type PoolClaimedEnterIterator struct {
	Event *PoolClaimedEnter // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolClaimedEnterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolClaimedEnter)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolClaimedEnter)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolClaimedEnterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolClaimedEnterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolClaimedEnter represents a ClaimedEnter event raised by the Pool contract.
type PoolClaimedEnter struct {
	Sender    common.Address
	Recipient common.Address
	Round     uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimedEnter is a free log retrieval operation binding the contract event 0x6d4f85f18c964c2bcc745c088b8daf077c5fc496dbd1f11c9126d3866a94539b.
//
// Solidity: event ClaimedEnter(address indexed sender, address indexed recipient, uint64 indexed round)
func (_Pool *PoolFilterer) FilterClaimedEnter(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, round []uint64) (*PoolClaimedEnterIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "ClaimedEnter", senderRule, recipientRule, roundRule)
	if err != nil {
		return nil, err
	}
	return &PoolClaimedEnterIterator{contract: _Pool.contract, event: "ClaimedEnter", logs: logs, sub: sub}, nil
}

// WatchClaimedEnter is a free log subscription operation binding the contract event 0x6d4f85f18c964c2bcc745c088b8daf077c5fc496dbd1f11c9126d3866a94539b.
//
// Solidity: event ClaimedEnter(address indexed sender, address indexed recipient, uint64 indexed round)
func (_Pool *PoolFilterer) WatchClaimedEnter(opts *bind.WatchOpts, sink chan<- *PoolClaimedEnter, sender []common.Address, recipient []common.Address, round []uint64) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "ClaimedEnter", senderRule, recipientRule, roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolClaimedEnter)
				if err := _Pool.contract.UnpackLog(event, "ClaimedEnter", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimedEnter is a log parse operation binding the contract event 0x6d4f85f18c964c2bcc745c088b8daf077c5fc496dbd1f11c9126d3866a94539b.
//
// Solidity: event ClaimedEnter(address indexed sender, address indexed recipient, uint64 indexed round)
func (_Pool *PoolFilterer) ParseClaimedEnter(log types.Log) (*PoolClaimedEnter, error) {
	event := new(PoolClaimedEnter)
	if err := _Pool.contract.UnpackLog(event, "ClaimedEnter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolClaimedExitIterator is returned from FilterClaimedExit and is used to iterate over the raw logs and unpacked data for ClaimedExit events raised by the Pool contract.
type PoolClaimedExitIterator struct {
	Event *PoolClaimedExit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolClaimedExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolClaimedExit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolClaimedExit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolClaimedExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolClaimedExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolClaimedExit represents a ClaimedExit event raised by the Pool contract.
type PoolClaimedExit struct {
	Exitee    common.Address
	Recipient common.Address
	Round     uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimedExit is a free log retrieval operation binding the contract event 0x4febfd878de8dfeadee5baa2dcf0e8e5ae4103684573cb3aa22ffb0fd02d5721.
//
// Solidity: event ClaimedExit(address indexed exitee, address indexed recipient, uint64 indexed round)
func (_Pool *PoolFilterer) FilterClaimedExit(opts *bind.FilterOpts, exitee []common.Address, recipient []common.Address, round []uint64) (*PoolClaimedExitIterator, error) {

	var exiteeRule []interface{}
	for _, exiteeItem := range exitee {
		exiteeRule = append(exiteeRule, exiteeItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "ClaimedExit", exiteeRule, recipientRule, roundRule)
	if err != nil {
		return nil, err
	}
	return &PoolClaimedExitIterator{contract: _Pool.contract, event: "ClaimedExit", logs: logs, sub: sub}, nil
}

// WatchClaimedExit is a free log subscription operation binding the contract event 0x4febfd878de8dfeadee5baa2dcf0e8e5ae4103684573cb3aa22ffb0fd02d5721.
//
// Solidity: event ClaimedExit(address indexed exitee, address indexed recipient, uint64 indexed round)
func (_Pool *PoolFilterer) WatchClaimedExit(opts *bind.WatchOpts, sink chan<- *PoolClaimedExit, exitee []common.Address, recipient []common.Address, round []uint64) (event.Subscription, error) {

	var exiteeRule []interface{}
	for _, exiteeItem := range exitee {
		exiteeRule = append(exiteeRule, exiteeItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "ClaimedExit", exiteeRule, recipientRule, roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolClaimedExit)
				if err := _Pool.contract.UnpackLog(event, "ClaimedExit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimedExit is a log parse operation binding the contract event 0x4febfd878de8dfeadee5baa2dcf0e8e5ae4103684573cb3aa22ffb0fd02d5721.
//
// Solidity: event ClaimedExit(address indexed exitee, address indexed recipient, uint64 indexed round)
func (_Pool *PoolFilterer) ParseClaimedExit(log types.Log) (*PoolClaimedExit, error) {
	event := new(PoolClaimedExit)
	if err := _Pool.contract.UnpackLog(event, "ClaimedExit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolClaimedMintIterator is returned from FilterClaimedMint and is used to iterate over the raw logs and unpacked data for ClaimedMint events raised by the Pool contract.
type PoolClaimedMintIterator struct {
	Event *PoolClaimedMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolClaimedMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolClaimedMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolClaimedMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolClaimedMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolClaimedMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolClaimedMint represents a ClaimedMint event raised by the Pool contract.
type PoolClaimedMint struct {
	Minter       common.Address
	PositionTick *big.Int
	Round        uint64
	Recipient    common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClaimedMint is a free log retrieval operation binding the contract event 0x1a5c9dafba043d990d534e332f02d83d1a2172cc6522c341a7f08297d45db5a8.
//
// Solidity: event ClaimedMint(address indexed minter, int24 indexed positionTick, uint64 indexed round, address recipient)
func (_Pool *PoolFilterer) FilterClaimedMint(opts *bind.FilterOpts, minter []common.Address, positionTick []*big.Int, round []uint64) (*PoolClaimedMintIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var positionTickRule []interface{}
	for _, positionTickItem := range positionTick {
		positionTickRule = append(positionTickRule, positionTickItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "ClaimedMint", minterRule, positionTickRule, roundRule)
	if err != nil {
		return nil, err
	}
	return &PoolClaimedMintIterator{contract: _Pool.contract, event: "ClaimedMint", logs: logs, sub: sub}, nil
}

// WatchClaimedMint is a free log subscription operation binding the contract event 0x1a5c9dafba043d990d534e332f02d83d1a2172cc6522c341a7f08297d45db5a8.
//
// Solidity: event ClaimedMint(address indexed minter, int24 indexed positionTick, uint64 indexed round, address recipient)
func (_Pool *PoolFilterer) WatchClaimedMint(opts *bind.WatchOpts, sink chan<- *PoolClaimedMint, minter []common.Address, positionTick []*big.Int, round []uint64) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var positionTickRule []interface{}
	for _, positionTickItem := range positionTick {
		positionTickRule = append(positionTickRule, positionTickItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "ClaimedMint", minterRule, positionTickRule, roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolClaimedMint)
				if err := _Pool.contract.UnpackLog(event, "ClaimedMint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimedMint is a log parse operation binding the contract event 0x1a5c9dafba043d990d534e332f02d83d1a2172cc6522c341a7f08297d45db5a8.
//
// Solidity: event ClaimedMint(address indexed minter, int24 indexed positionTick, uint64 indexed round, address recipient)
func (_Pool *PoolFilterer) ParseClaimedMint(log types.Log) (*PoolClaimedMint, error) {
	event := new(PoolClaimedMint)
	if err := _Pool.contract.UnpackLog(event, "ClaimedMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolCollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the Pool contract.
type PoolCollectIterator struct {
	Event *PoolCollect // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolCollect)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolCollect)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolCollect represents a Collect event raised by the Pool contract.
type PoolCollect struct {
	Owner        common.Address
	PositionTick *big.Int
	Recipient    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCollect is a free log retrieval operation binding the contract event 0x8f2cd17f522199653b5efd04284ecf83ff604a050a7a8c124e11f79924306e5b.
//
// Solidity: event Collect(address indexed owner, int24 indexed positionTick, address recipient, uint96 amount)
func (_Pool *PoolFilterer) FilterCollect(opts *bind.FilterOpts, owner []common.Address, positionTick []*big.Int) (*PoolCollectIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var positionTickRule []interface{}
	for _, positionTickItem := range positionTick {
		positionTickRule = append(positionTickRule, positionTickItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Collect", ownerRule, positionTickRule)
	if err != nil {
		return nil, err
	}
	return &PoolCollectIterator{contract: _Pool.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x8f2cd17f522199653b5efd04284ecf83ff604a050a7a8c124e11f79924306e5b.
//
// Solidity: event Collect(address indexed owner, int24 indexed positionTick, address recipient, uint96 amount)
func (_Pool *PoolFilterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *PoolCollect, owner []common.Address, positionTick []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var positionTickRule []interface{}
	for _, positionTickItem := range positionTick {
		positionTickRule = append(positionTickRule, positionTickItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Collect", ownerRule, positionTickRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolCollect)
				if err := _Pool.contract.UnpackLog(event, "Collect", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollect is a log parse operation binding the contract event 0x8f2cd17f522199653b5efd04284ecf83ff604a050a7a8c124e11f79924306e5b.
//
// Solidity: event Collect(address indexed owner, int24 indexed positionTick, address recipient, uint96 amount)
func (_Pool *PoolFilterer) ParseCollect(log types.Log) (*PoolCollect, error) {
	event := new(PoolCollect)
	if err := _Pool.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolEnteredIterator is returned from FilterEntered and is used to iterate over the raw logs and unpacked data for Entered events raised by the Pool contract.
type PoolEnteredIterator struct {
	Event *PoolEntered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolEntered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolEntered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolEntered represents a Entered event raised by the Pool contract.
type PoolEntered struct {
	Sender     common.Address
	Round      uint64
	Claimer    common.Address
	AmountSent *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEntered is a free log retrieval operation binding the contract event 0xcc2db95feaea18b716e394e12503afc7e9b4914f0fce01087f189875edaceb43.
//
// Solidity: event Entered(address indexed sender, uint64 indexed round, address indexed claimer, uint96 amountSent)
func (_Pool *PoolFilterer) FilterEntered(opts *bind.FilterOpts, sender []common.Address, round []uint64, claimer []common.Address) (*PoolEnteredIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Entered", senderRule, roundRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &PoolEnteredIterator{contract: _Pool.contract, event: "Entered", logs: logs, sub: sub}, nil
}

// WatchEntered is a free log subscription operation binding the contract event 0xcc2db95feaea18b716e394e12503afc7e9b4914f0fce01087f189875edaceb43.
//
// Solidity: event Entered(address indexed sender, uint64 indexed round, address indexed claimer, uint96 amountSent)
func (_Pool *PoolFilterer) WatchEntered(opts *bind.WatchOpts, sink chan<- *PoolEntered, sender []common.Address, round []uint64, claimer []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Entered", senderRule, roundRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolEntered)
				if err := _Pool.contract.UnpackLog(event, "Entered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEntered is a log parse operation binding the contract event 0xcc2db95feaea18b716e394e12503afc7e9b4914f0fce01087f189875edaceb43.
//
// Solidity: event Entered(address indexed sender, uint64 indexed round, address indexed claimer, uint96 amountSent)
func (_Pool *PoolFilterer) ParseEntered(log types.Log) (*PoolEntered, error) {
	event := new(PoolEntered)
	if err := _Pool.contract.UnpackLog(event, "Entered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolExitedIterator is returned from FilterExited and is used to iterate over the raw logs and unpacked data for Exited events raised by the Pool contract.
type PoolExitedIterator struct {
	Event *PoolExited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolExited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolExited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolExited represents a Exited event raised by the Pool contract.
type PoolExited struct {
	Exitee       common.Address
	Round        uint64
	Claimer      common.Address
	SharesLocked *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExited is a free log retrieval operation binding the contract event 0x4b57c930249a429dcb198d48b8c5cdfb16ddc71cfb1535c8391baf12501b7041.
//
// Solidity: event Exited(address indexed exitee, uint64 indexed round, address claimer, uint128 sharesLocked)
func (_Pool *PoolFilterer) FilterExited(opts *bind.FilterOpts, exitee []common.Address, round []uint64) (*PoolExitedIterator, error) {

	var exiteeRule []interface{}
	for _, exiteeItem := range exitee {
		exiteeRule = append(exiteeRule, exiteeItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Exited", exiteeRule, roundRule)
	if err != nil {
		return nil, err
	}
	return &PoolExitedIterator{contract: _Pool.contract, event: "Exited", logs: logs, sub: sub}, nil
}

// WatchExited is a free log subscription operation binding the contract event 0x4b57c930249a429dcb198d48b8c5cdfb16ddc71cfb1535c8391baf12501b7041.
//
// Solidity: event Exited(address indexed exitee, uint64 indexed round, address claimer, uint128 sharesLocked)
func (_Pool *PoolFilterer) WatchExited(opts *bind.WatchOpts, sink chan<- *PoolExited, exitee []common.Address, round []uint64) (event.Subscription, error) {

	var exiteeRule []interface{}
	for _, exiteeItem := range exitee {
		exiteeRule = append(exiteeRule, exiteeItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Exited", exiteeRule, roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolExited)
				if err := _Pool.contract.UnpackLog(event, "Exited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExited is a log parse operation binding the contract event 0x4b57c930249a429dcb198d48b8c5cdfb16ddc71cfb1535c8391baf12501b7041.
//
// Solidity: event Exited(address indexed exitee, uint64 indexed round, address claimer, uint128 sharesLocked)
func (_Pool *PoolFilterer) ParseExited(log types.Log) (*PoolExited, error) {
	event := new(PoolExited)
	if err := _Pool.contract.UnpackLog(event, "Exited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the Pool contract.
type PoolInitializeIterator struct {
	Event *PoolInitialize // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolInitialize)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolInitialize)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolInitialize represents a Initialize event raised by the Pool contract.
type PoolInitialize struct {
	Fr   *big.Int
	Tick *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x963b19b795d0d7f389b1248834a4da01059a1bca8509a324f0df4b9d9208a075.
//
// Solidity: event Initialize(int80 fr, int24 tick)
func (_Pool *PoolFilterer) FilterInitialize(opts *bind.FilterOpts) (*PoolInitializeIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return &PoolInitializeIterator{contract: _Pool.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x963b19b795d0d7f389b1248834a4da01059a1bca8509a324f0df4b9d9208a075.
//
// Solidity: event Initialize(int80 fr, int24 tick)
func (_Pool *PoolFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *PoolInitialize) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolInitialize)
				if err := _Pool.contract.UnpackLog(event, "Initialize", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialize is a log parse operation binding the contract event 0x963b19b795d0d7f389b1248834a4da01059a1bca8509a324f0df4b9d9208a075.
//
// Solidity: event Initialize(int80 fr, int24 tick)
func (_Pool *PoolFilterer) ParseInitialize(log types.Log) (*PoolInitialize, error) {
	event := new(PoolInitialize)
	if err := _Pool.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Pool contract.
type PoolMintIterator struct {
	Event *PoolMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolMint represents a Mint event raised by the Pool contract.
type PoolMint struct {
	Sender       common.Address
	PositionTick *big.Int
	Round        uint64
	Claimer      common.Address
	AmountSent   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x51c0f7607cd62fa0b901c1d125a2b45f70b72849ba5b0e2e552366282e75033d.
//
// Solidity: event Mint(address indexed sender, int24 indexed positionTick, uint64 indexed round, address claimer, uint96 amountSent)
func (_Pool *PoolFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address, positionTick []*big.Int, round []uint64) (*PoolMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var positionTickRule []interface{}
	for _, positionTickItem := range positionTick {
		positionTickRule = append(positionTickRule, positionTickItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Mint", senderRule, positionTickRule, roundRule)
	if err != nil {
		return nil, err
	}
	return &PoolMintIterator{contract: _Pool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x51c0f7607cd62fa0b901c1d125a2b45f70b72849ba5b0e2e552366282e75033d.
//
// Solidity: event Mint(address indexed sender, int24 indexed positionTick, uint64 indexed round, address claimer, uint96 amountSent)
func (_Pool *PoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PoolMint, sender []common.Address, positionTick []*big.Int, round []uint64) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var positionTickRule []interface{}
	for _, positionTickItem := range positionTick {
		positionTickRule = append(positionTickRule, positionTickItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Mint", senderRule, positionTickRule, roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolMint)
				if err := _Pool.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0x51c0f7607cd62fa0b901c1d125a2b45f70b72849ba5b0e2e552366282e75033d.
//
// Solidity: event Mint(address indexed sender, int24 indexed positionTick, uint64 indexed round, address claimer, uint96 amountSent)
func (_Pool *PoolFilterer) ParseMint(log types.Log) (*PoolMint, error) {
	event := new(PoolMint)
	if err := _Pool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Pool contract.
type PoolSwapIterator struct {
	Event *PoolSwap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolSwap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolSwap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolSwap represents a Swap event raised by the Pool contract.
type PoolSwap struct {
	LiquidityMoved *big.Int
	Tick           *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd91abeef74e03de9570ddb10febb3e28fc4f9b3ed1bf7d0d803db21c4dca4faa.
//
// Solidity: event Swap(uint96 liquidityMoved, int24 tick)
func (_Pool *PoolFilterer) FilterSwap(opts *bind.FilterOpts) (*PoolSwapIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &PoolSwapIterator{contract: _Pool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd91abeef74e03de9570ddb10febb3e28fc4f9b3ed1bf7d0d803db21c4dca4faa.
//
// Solidity: event Swap(uint96 liquidityMoved, int24 tick)
func (_Pool *PoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *PoolSwap) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolSwap)
				if err := _Pool.contract.UnpackLog(event, "Swap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwap is a log parse operation binding the contract event 0xd91abeef74e03de9570ddb10febb3e28fc4f9b3ed1bf7d0d803db21c4dca4faa.
//
// Solidity: event Swap(uint96 liquidityMoved, int24 tick)
func (_Pool *PoolFilterer) ParseSwap(log types.Log) (*PoolSwap, error) {
	event := new(PoolSwap)
	if err := _Pool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Pool contract.
type PoolTransferIterator struct {
	Event *PoolTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PoolTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTransfer represents a Transfer event raised by the Pool contract.
type PoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pool *PoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolTransferIterator{contract: _Pool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pool *PoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTransfer)
				if err := _Pool.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pool *PoolFilterer) ParseTransfer(log types.Log) (*PoolTransfer, error) {
	event := new(PoolTransfer)
	if err := _Pool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
