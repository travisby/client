// Auto-generated by avdl-compiler v1.3.19 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/gregor_ui.avdl

package keybase1

import (
	gregor1 "github.com/keybase/client/go/protocol/gregor1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type PushReason int

const (
	PushReason_NONE        PushReason = 0
	PushReason_RECONNECTED PushReason = 1
	PushReason_NEW_DATA    PushReason = 2
)

func (o PushReason) DeepCopy() PushReason { return o }

var PushReasonMap = map[string]PushReason{
	"NONE":        0,
	"RECONNECTED": 1,
	"NEW_DATA":    2,
}

var PushReasonRevMap = map[PushReason]string{
	0: "NONE",
	1: "RECONNECTED",
	2: "NEW_DATA",
}

func (e PushReason) String() string {
	if v, ok := PushReasonRevMap[e]; ok {
		return v
	}
	return ""
}

type PushStateArg struct {
	State  gregor1.State `codec:"state" json:"state"`
	Reason PushReason    `codec:"reason" json:"reason"`
}

func (o PushStateArg) DeepCopy() PushStateArg {
	return PushStateArg{
		State:  o.State.DeepCopy(),
		Reason: o.Reason.DeepCopy(),
	}
}

type PushOutOfBandMessagesArg struct {
	Oobm []gregor1.OutOfBandMessage `codec:"oobm" json:"oobm"`
}

func (o PushOutOfBandMessagesArg) DeepCopy() PushOutOfBandMessagesArg {
	return PushOutOfBandMessagesArg{
		Oobm: (func(x []gregor1.OutOfBandMessage) []gregor1.OutOfBandMessage {
			var ret []gregor1.OutOfBandMessage
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Oobm),
	}
}

type GregorUIInterface interface {
	PushState(context.Context, PushStateArg) error
	PushOutOfBandMessages(context.Context, []gregor1.OutOfBandMessage) error
}

func GregorUIProtocol(i GregorUIInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.gregorUI",
		Methods: map[string]rpc.ServeHandlerDescription{
			"pushState": {
				MakeArg: func() interface{} {
					ret := make([]PushStateArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PushStateArg)
					if !ok {
						err = rpc.NewTypeError((*[]PushStateArg)(nil), args)
						return
					}
					err = i.PushState(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"pushOutOfBandMessages": {
				MakeArg: func() interface{} {
					ret := make([]PushOutOfBandMessagesArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PushOutOfBandMessagesArg)
					if !ok {
						err = rpc.NewTypeError((*[]PushOutOfBandMessagesArg)(nil), args)
						return
					}
					err = i.PushOutOfBandMessages(ctx, (*typedArgs)[0].Oobm)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type GregorUIClient struct {
	Cli rpc.GenericClient
}

func (c GregorUIClient) PushState(ctx context.Context, __arg PushStateArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.gregorUI.pushState", []interface{}{__arg}, nil)
	return
}

func (c GregorUIClient) PushOutOfBandMessages(ctx context.Context, oobm []gregor1.OutOfBandMessage) (err error) {
	__arg := PushOutOfBandMessagesArg{Oobm: oobm}
	err = c.Cli.Call(ctx, "keybase.1.gregorUI.pushOutOfBandMessages", []interface{}{__arg}, nil)
	return
}
