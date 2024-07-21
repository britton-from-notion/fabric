package eval

import (
	"github.com/blackstork-io/fabric/parser/definitions"
	"github.com/blackstork-io/fabric/plugin/dataspec"
)

type PluginAction struct {
	PluginName string
	BlockName  string
	Meta       *definitions.MetaBlock
	Config     *dataspec.Block
	Args       *dataspec.Block
}
