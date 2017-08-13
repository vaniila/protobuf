package bitflags

import (
	"github.com/vaniila/protobuf/gogoproto"
	"github.com/vaniila/protobuf/plugin/testgen"
	"github.com/vaniila/protobuf/protoc-gen-gogo/generator"
)

type test struct {
	*generator.Generator
}

func init() {
	testgen.RegisterTestPlugin(NewTest)
}

func NewTest(g *generator.Generator) testgen.TestPlugin {
	return &test{g}
}

func (p *test) Generate(imports generator.PluginImports, file *generator.FileDescriptor) bool {
	for _, message := range file.Messages() {
		if gogoproto.HasBitflags(file.FileDescriptorProto, message.DescriptorProto) {
			return true
		}
	}
	return false
}
