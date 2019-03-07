package epitomecl

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on wire codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetName{}, "epitomecl/SetName", nil)
	cdc.RegisterConcrete(MsgBuyName{}, "epitomecl/BuyName", nil)
}
