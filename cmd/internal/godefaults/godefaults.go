package godefaults

import (
	"flag"
	"sync"

	"github.com/yu31/protoc-kit-go/helper/pkbuffer"
	"github.com/yu31/protoc-kit-go/utils/pkmessage"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/yu31/protoc-kit-go/pkgenerator"
)

const version = "0.0.1"

// Plugin for implements pgkgenerator.Plugin
type Plugin struct {
	flags    flag.FlagSet
	funcName string // The generated method func name for defaults.
	g        *protogen.GeneratedFile

	pp   *protogen.Plugin
	file *protogen.File

	messages []*protogen.Message

	// The message of currently being processed.
	once    *sync.Once
	message *protogen.Message
	extend  *pkbuffer.Buffer
}

func New() pkgenerator.Plugin {
	p := &Plugin{}
	// Usage: --godefaults_opt="fn=SetDefaults"
	p.flags.StringVar(&p.funcName, "fn", "", "The method func name")
	return p
}

// Name identifies the Plugin.
func (p *Plugin) Name() string {
	return "defaults"
}

// Version identifies the Plugin version.
func (p *Plugin) Version() string {
	return version
}

// ParamFunc is used for accept parameters from the command line.
func (p *Plugin) ParamFunc() func(name, value string) error {
	return p.flags.Set
}

// Init is called once before code generated.
// The `file` will be ignored if return false.
func (p *Plugin) Init(pp *protogen.Plugin, file *protogen.File) bool {
	if len(file.Messages) == 0 {
		return false
	}
	p.pp = pp
	p.file = file
	p.messages = pkmessage.LoadValidMessages(file.Messages)
	return true
}

// Generate to generate codes for specified file.
// except for the imports, by calling the generator's methods P, In, and Out.
func (p *Plugin) Generate(g *protogen.GeneratedFile) {
	p.g = g
	for _, msg := range p.messages {
		p.generateCodeForMessage(msg)
	}
}

func (p *Plugin) getNameDefaultsMethod() string {
	if p.funcName != "" {
		return p.funcName
	}
	return "SetDefaults"
}
