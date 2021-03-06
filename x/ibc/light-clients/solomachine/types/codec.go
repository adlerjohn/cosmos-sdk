package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	clientexported "github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
)

// RegisterInterfaces register the ibc channel submodule interfaces to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*clientexported.ClientState)(nil),
		&ClientState{},
	)
	registry.RegisterImplementations(
		(*clientexported.ConsensusState)(nil),
		&ConsensusState{},
	)
	registry.RegisterImplementations(
		(*clientexported.Header)(nil),
		&Header{},
	)
	registry.RegisterImplementations(
		(*clientexported.Misbehaviour)(nil),
		&Misbehaviour{},
	)
}

var (
	// SubModuleCdc references the global x/ibc/light-clients/solomachine module codec. Note, the codec
	// should ONLY be used in certain instances of tests and for JSON encoding..
	//
	// The actual codec used for serialization should be provided to x/ibc/light-clients/solomachine and
	// defined at the application level.
	SubModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
)
