package parser

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"

	"github.com/blackstork-io/fabric/parser/blocks/internal/tree"
	"github.com/blackstork-io/fabric/parser/definitions"
	"github.com/blackstork-io/fabric/pkg/diagnostics"
)

// Collection of defined blocks

type DefinedBlocks struct {
	GlobalConfig *definitions.GlobalConfig
	Config       nodeMap[pluginKindT, nodeMap[pluginNameT, nodeMap[blockNameT, *definitions.Config]]]
	Documents    nodeMap[documentNameT, *definitions.Document]
	Sections     nodeMap[sectionNameT, *definitions.Section]
	Plugins      nodeMap[pluginKindT, nodeMap[pluginNameT, nodeMap[blockNameT, definitions.PluginIface]]]
}

func mapGetOrInit[K1, K2 comparable, V any](m map[K1]map[K2]V, key K1) (innerMap map[K2]V) {
	innerMap, found := m[key]
	if !found {
		innerMap = map[K2]V{}
		m[key] = innerMap
	}
	return
}

func mapToCty(m map[string]map[string]cty.Value) (res map[string]cty.Value) {
	res = make(map[string]cty.Value, len(m))
	for k, v := range m {
		if len(v) == 0 {
			continue
		}
		res[k] = cty.MapVal(v)
	}
	return
}

func PluginMapToCty[V definitions.FabricBlock](plugins map[definitions.Key]V) (content, data cty.Value) {
	// [plugin_kind][plugin_name][block_name]*definitions.Plugin

	pluginMap := [2]map[string]map[string]cty.Value{
		{},
		{},
	}
	for k, v := range plugins {
		var idx int
		switch k.PluginKind {
		case definitions.BlockKindContent:
			idx = 0
		case definitions.BlockKindData:
			idx = 1
		default:
			panic("must be exhaustive")
		}
		blockNameToVal := mapGetOrInit(pluginMap[idx], k.PluginName)
		blockNameToVal[k.BlockName] = definitions.ToCtyValue(v)
	}
	pluginKindToVal := [2]cty.Value{}

	for idx, pl := range pluginMap {
		if len(pl) == 0 {
			continue
		}
		pluginKindToVal[idx] = cty.MapVal(mapToCty(pl))
	}
	return pluginKindToVal[0], pluginKindToVal[1]
}

func (db *DefinedBlocks) AsValueMap() map[string]cty.Value {
	content, _ := db.Plugins.Map[definitions.BlockKindContent].AsCtyValue()
	data, _ := db.Plugins.Map[definitions.BlockKindData].AsCtyValue()
	config, _ := db.Config.AsCtyValue()
	sections, _ := db.Sections.AsCtyValue()
	return map[string]cty.Value{
		definitions.BlockKindContent: content,
		definitions.BlockKindData:    data,
		definitions.BlockKindSection: sections,
		definitions.BlockKindConfig:  config,
	}
}

func (db *DefinedBlocks) DefaultConfigFor(plugin *definitions.Plugin) (config *definitions.Config) {
	return db.DefaultConfig(plugin.Kind(), plugin.Name())
}

func (db *DefinedBlocks) DefaultConfig(pluginKind, pluginName string) (config *definitions.Config) {
	return db.Config.Get(pluginKind).Get(pluginName).Get("")
}

func (db *DefinedBlocks) Merge(other *DefinedBlocks) (diags diagnostics.Diag) {
	if other.GlobalConfig != nil {
		if db.GlobalConfig != nil {
			diags.Add("Global config declared multiple times", "")
		}
		db.GlobalConfig = other.GlobalConfig
	}
	diags.Extend(MergeNestedMap(db.Config, other.Config))
	diags.Extend(MergeMap(db.Documents, other.Documents))
	diags.Extend(MergeMap(db.Sections, other.Sections))
	diags.Extend(MergeNestedMap(db.Plugins, other.Plugins))
	return
}

func AddIfMissing[M ~map[K]V, K comparable, V definitions.FabricBlock](m M, key K, newBlock V) *hcl.Diagnostic {
	if origBlock, found := m[key]; found {
		kind := origBlock.GetHCLBlock().Type
		origDefRange := origBlock.GetHCLBlock().DefRange
		return &hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  fmt.Sprintf("Duplicate '%s' declaration", kind),
			Detail:   fmt.Sprintf("'%s' with the same name originally defined at %s:%d", kind, origDefRange.Filename, origDefRange.Start.Line),
			Subject:  newBlock.GetHCLBlock().DefRange.Ptr(),
		}
	}
	m[key] = newBlock
	return nil
}

func NewDefinedBlocks() *DefinedBlocks {
	return &DefinedBlocks{
		Config:    NewMap[pluginKindT, nodeMap[pluginNameT, nodeMap[blockNameT, *definitions.Config]]](),
		Documents: NewMap[documentNameT, *definitions.Document](),
		Sections:  NewMap[sectionNameT, *definitions.Section](),
		Plugins:   NewMap[pluginKindT, nodeMap[pluginNameT, nodeMap[blockNameT, definitions.PluginIface]]](), // map[definitions.Key]*definitions.Plugin{},
	}
}

type DocumentsMap struct {
	tree.NodeSigil
	Map map[string]*definitions.Document
}

func NewMap[K key, V val]() nodeMap[K, V] {
	return nodeMap[K, V]{
		Map: map[string]V{},
	}
}

type (
	documentNameT string
	pluginKindT   string
	pluginNameT   string
	blockNameT    string
	sectionNameT  string
)

func (documentNameT) FriendlyName() string {
	return "document name"
}

func (pluginKindT) FriendlyName() string {
	return "plugin kind"
}

func (pluginNameT) FriendlyName() string {
	return "plugin name"
}

func (blockNameT) FriendlyName() string {
	return "block name"
}

func (sectionNameT) FriendlyName() string {
	return "section name"
}

type key interface {
	~string
	tree.Namer
}

type val interface {
	tree.Node
	tree.CtyAble
}

type nodeMap[K key, V val] struct {
	tree.NodeSigil
	Map map[string]V
}

var _ interface {
	tree.Node
	tree.CtyAble
	tree.StrIndexable
} = nodeMap[documentNameT, *definitions.Document]{}

func (m nodeMap[K, V]) CtyType() cty.Type {
	var v V
	return cty.Map(v.CtyType())
}

func (m nodeMap[K, V]) AsCtyValue() cty.Value {
	if len(m.Map) == 0 {
		var v V
		return cty.MapValEmpty(v.CtyType())
	}
	mp := make(map[string]cty.Value, len(m.Map))
	for k, v := range m.Map {
		mp[k] = v.AsCtyValue()
	}
	return cty.MapVal(mp)
}

func (m nodeMap[K, V]) FriendlyName() string {
	var k K
	var v V
	return fmt.Sprintf(
		"map of string (%s) to %s", k.FriendlyName(), v.FriendlyName(),
	)
}

func (m nodeMap[K, V]) IndexStr(idx string) tree.Node {
	if v, found := m.Map[idx]; found {
		return v
	}
	return nil
}

func (m nodeMap[K, V]) Get(idx string) V {
	if m.Map != nil {
		if v, found := m.Map[idx]; found {
			return v
		}
	}
	var v V
	return v
}

func IterMap[V val](m nodeMap[pluginKindT, nodeMap[pluginNameT, nodeMap[blockNameT, V]]], fn func(pluginKind, pluginName, blockName string, val V)) {
	for k1, v1 := range m.Map {
		for k2, v2 := range v1.Map {
			for k3, v3 := range v2.Map {
				fn(k1, k2, k3, v3)
			}
		}
	}
}

func SetIfMissing[V interface {
	definitions.FabricBlock
	val
}](m nodeMap[pluginKindT, nodeMap[pluginNameT, nodeMap[blockNameT, V]]], pluginKind, pluginName, blockName string, val V) *hcl.Diagnostic {
	m1, found := m.Map[pluginKind]
	if !found {
		m1 = NewMap[pluginNameT, nodeMap[blockNameT, V]]()
		m2 := NewMap[blockNameT, V]()
		m2.Map[blockName] = val
		m1.Map[pluginName] = m2
		m.Map[pluginKind] = m1
		return nil
	}
	m2, found := m1.Map[pluginName]
	if !found {
		m2 = NewMap[blockNameT, V]()
		m2.Map[blockName] = val
		m1.Map[pluginName] = m2
		return nil
	}
	return AddIfMissing(m2.Map, blockName, val)
}

func MergeNestedMap[V interface {
	definitions.FabricBlock
	val
}](dst, src nodeMap[pluginKindT, nodeMap[pluginNameT, nodeMap[blockNameT, V]]]) (diags diagnostics.Diag) {
	for k1, src1 := range src.Map {
		dst1, found := dst.Map[k1]
		if !found {
			dst.Map[k1] = src1
			continue
		}
		for k2, src2 := range src1.Map {
			dst2, found := dst1.Map[k2]
			if !found {
				dst1.Map[k2] = src2
				continue
			}
			diags.Extend(MergeMap(dst2, src2))
		}
	}
	return nil
}

func MergeMap[K key, V interface {
	definitions.FabricBlock
	val
}](dst, src nodeMap[K, V]) (diags diagnostics.Diag) {
	for k, v := range src.Map {
		diags.Append(AddIfMissing(dst.Map, k, v))
	}
	return
}
