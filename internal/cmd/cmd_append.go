package cmd

import (
	"github.com/dicedb/dice/internal/object"
	dstore "github.com/dicedb/dice/internal/store"
	"github.com/dicedb/dicedb-go/wire"
)

var cAppend = &DiceDBCommand{
	Name:      "APPEND",
	HelpShort: "returns the length of the value stored at a specified key",
	Eval:      evalAppend,
}

func init() {
	commandRegistry.AddCommand(cAppend)
}

func evalAppend(c *Cmd, s *dstore.Store) (*CmdRes, error) {
	if len(c.C.Args) != 2 {
		return cmdResNil, errWrongArgumentCount("APPEND")
	}

	key := c.C.Args[0]
	val := c.C.Args[1]
	obj := s.Get(key)

	if obj == nil {
		obj = s.NewObj(val, -1, object.ObjTypeString)
		s.Put(key, obj)

		return &CmdRes{R: &wire.Response{
			Value: &wire.Response_VInt{
				VInt: int64(len(val)),
			},
		}}, nil
	}

	objVal, ok := obj.Value.(string)

	if !ok {
		return cmdResNil, errWrongTypeOperation("APPEND")
	}

	newVal := objVal + val
	obj.Value = newVal
	s.Put(key, obj)

	return &CmdRes{R: &wire.Response{
		Value: &wire.Response_VInt{
			VInt: int64(len(newVal)),
		},
	}}, nil
}
