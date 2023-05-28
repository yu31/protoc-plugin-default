package godefaults

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/yu31/protoc-kit-go/helper/pkerror"

	"github.com/yu31/protoc-plugin-defaults/xgo/pb/pbdefaults"
)

func (p *Plugin) getValueForFieldAny(typeAny *pbdefaults.TypeAny) string {
	msgDef := p.lookupMessageByTypeUrl(*typeAny.TypeUrl)
	msgType := p.g.QualifiedGoIdent(msgDef.GoIdent)
	anyValue := fmt.Sprintf("&%s{}", msgType)
	return anyValue
}

func (p *Plugin) lookupMessageByTypeUrl(typeUrl string) *protogen.Message {
	var importPath string
	var msgName string

	i := strings.Index(typeUrl, ".")
	if i == -1 {
		err := pkerror.New("invalid format of type_url [%s]", typeUrl)
		panic(err)
	}
	// The first charset is "." represents the message is location in the current proto file.
	if i == 0 {
		importPath = "."
	} else {
		importPath = strings.TrimSpace(typeUrl[:i])
	}
	msgName = strings.TrimSpace(typeUrl[i+1:])

	if importPath == "" {
		err := pkerror.New("not found import path in type_url [%s]", typeUrl)
		panic(err)
	}
	if msgName == "" {
		err := pkerror.New("not found message name in type_url [%s]", typeUrl)
		panic(err)
	}
	if importPath != "." {
		importPath = importPath + ".proto"
	}
	return p.lookupMessageFromInputFiles(importPath, msgName)
}

func (p *Plugin) lookupMessageFromInputFiles(importPath string, msgName string) *protogen.Message {
	var file *protogen.File

	if importPath == "." {
		file = p.file
	} else {
		for _, f := range p.pp.Files {
			if f.Desc.Path() == importPath {
				file = f
				break
			}
		}
	}
	if file == nil {
		err := pkerror.New("cannot find proto file [%s] in import path", importPath)
		panic(err)
	}

	lookupMsg := func(messages []*protogen.Message, name string) *protogen.Message {
		for _, m := range messages {
			if string(m.Desc.Name()) == name {
				return m
			}
		}
		err := pkerror.New("cannot find proto message [%s] in file [%s]", msgName, importPath)
		panic(err)
	}

	parts := strings.Split(msgName, ".")
	message := lookupMsg(file.Messages, parts[0])

	for _, sub := range parts[1:] {
		messages := message.Messages
		message = lookupMsg(messages, sub)
	}
	return message
}
