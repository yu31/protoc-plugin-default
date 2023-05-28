package godefault

import (
	"fmt"
	"os"
	"sync"

	"github.com/yu31/protoc-kit-go/helper/pkbuffer"
	"github.com/yu31/protoc-kit-go/helper/pkwkt"
	"github.com/yu31/protoc-kit-go/utils/pkfield"
	"github.com/yu31/protoc-kit-go/utils/pkmessage"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/yu31/protoc-kit-go/helper/pkerror"
	"github.com/yu31/protoc-kit-go/utils/pkident"

	"github.com/yu31/protoc-plugin-default/xgo/pb/pbdefault"

	"github.com/yu31/protoc-plugin-default/cmd/internal/godefault/pkg/importpkg"
)

func (p *Plugin) generateCodeForMessage(msg *protogen.Message) {
	p.once = &sync.Once{}
	p.message = msg
	p.extend = pkbuffer.New()

	p.g.P("// Set default value for message ", msg.Desc.Name(), " that in file ", p.file.Desc.Path())
	p.g.P("func (x *", msg.GoIdent.GoName, ") ", p.getNameDefaultMethod(), "() {")
	p.g.P("if x == nil {")
	p.g.P("    return")
	p.g.P("}")

	// Generate validate field method for field.
	for _, field := range pkmessage.LoadFieldLists(msg) {
		catch := pkerror.Recover("godefault", p.file, p.message, field, func() {
			p.generateCodeForField(field)
		})
		if catch {
			os.Exit(1)
		}
	}
	p.g.P("}")

	// write extend codes.
	p.g.P(p.extend.String())
}

func (p *Plugin) generateCodeForField(field *protogen.Field) {
	var skipEval bool
	switch {
	case field.Desc.IsMap():
		skipEval = p.generateCodeForMap(field)
	case field.Desc.IsList():
		skipEval = p.generateCodeForRepeated(field)
	case pkfield.FieldIsOneOf(field):
		p.generateCodeForOneOf(field)
		p.generateCodeForInOneOf(field)
	default:
		skipEval = p.generateCodeForPlain(field, false)
	}

	if skipEval {
		return
	}
	p.generateCodeForEval(field, false)
}

func (p *Plugin) generateCodeForOneOf(field *protogen.Field) {
	opts := p.loadOneOfOptions(field.Oneof)
	if opts.Field == nil || *opts.Field == "" {
		return
	}

	typeSet := *opts.Field

	var defaultField *protogen.Field
	for _, f := range field.Oneof.Fields {
		if string(f.Desc.Name()) == typeSet {
			defaultField = f
			break
		}
	}
	if defaultField == nil {
		err := pkerror.New("cannot found the default field [%s] in oneof parts.", typeSet)
		panic(err)
	}

	oneOfName := field.Oneof.GoName
	oneOfType := p.g.QualifiedGoIdent(defaultField.GoIdent)
	p.g.P("if x.", oneOfName, " == nil", " {")
	p.g.P("    x.", oneOfName, " = ", "new(", oneOfType, ")")
	p.g.P("}")
}

func (p *Plugin) generateCodeForInOneOf(field *protogen.Field) {
	oneOfName := field.Oneof.GoName

	p.g.P("switch o := x.", oneOfName, ".(type) {")
	for _, oneOfField := range field.Oneof.Fields {
		p.g.P("case *", p.g.QualifiedGoIdent(oneOfField.GoIdent), ":")
		skipEval := p.generateCodeForPlain(oneOfField, true)
		if !skipEval {
			p.generateCodeForEval(oneOfField, true)
		}
	}
	p.g.P("default:")
	p.g.P("    _ = o // to avoid unused panic")
	p.g.P("}")
}

func (p *Plugin) generateCodeForPlain(field *protogen.Field, inOneOf bool) (skipEval bool) {
	opts := p.loadPlainOptions(field)
	if opts == nil {
		return
	}
	if opts.SkipEval != nil && *opts.SkipEval {
		skipEval = true
	}
	if opts.Value == nil {
		return
	}

	var varName string
	if inOneOf {
		varName = "o" + "." + field.GoName
	} else {
		varName = "x" + "." + field.GoName
	}

	isOptional := pkfield.FieldIsOptional(field)

	_condApply, valueDefault := p.genCodeByFileType(field, opts.Value, varName, skipEval)
	if _condApply == "" {
		return
	}

	var ignoreEmtpy bool
	if opts.IgnoreEmpty != nil && *opts.IgnoreEmpty {
		ignoreEmtpy = true
	}

	var condApply string
	switch {
	case ignoreEmtpy && field.Desc.Kind() == protoreflect.BytesKind:
		condApply = varName + " == nil"
	case ignoreEmtpy && field.Desc.Kind() == protoreflect.MessageKind:
		condApply = varName + " == nil"
	case ignoreEmtpy && isOptional:
		condApply = varName + " == nil"
	case !ignoreEmtpy && isOptional:
		condApply = varName + " == nil || *" + _condApply
	default:
		condApply = _condApply
	}

	p.g.P("if ", condApply, " {")
	if isOptional {
		goType := pkfield.FieldGoType(p.g, field)
		p.g.P("v := ", goType, "(", valueDefault, ")")
		p.g.P(varName, " = &v")
	} else {
		p.g.P(varName, " = ", valueDefault)
	}
	p.g.P("}")
	return
}
func (p *Plugin) generateCodeForRepeated(field *protogen.Field) (skipEval bool) {
	opts := p.loadRepeatedOptions(field)
	if opts == nil {
		return
	}

	if opts.SkipEval != nil && *opts.SkipEval {
		skipEval = true
	}

	if len(opts.Items) == 0 {
		return
	}

	var buf pkbuffer.Buffer

	buf.WriteString(pkfield.FieldGoType(p.g, field))
	buf.WriteString("{")
	buf.WriteStringLn("")

	for _, v := range opts.Items {
		_, item := p.genCodeByFileType(field, v, "", skipEval)
		buf.WriteString(item)
		buf.WriteStringLn(",")
	}
	buf.WriteString("}")

	var condApply string

	if opts.IgnoreEmpty != nil && *opts.IgnoreEmpty {
		condApply = "x." + field.GoName + " == nil"
	} else {
		condApply = "len(x." + field.GoName + ") == 0"
	}

	p.g.P("if ", condApply, "{")
	p.g.P("    x.", field.GoName, " = ", buf.String())
	p.g.P("}")
	return
}
func (p *Plugin) generateCodeForMap(field *protogen.Field) (skipEval bool) {
	opts := p.loadMapOptions(field)
	if opts == nil {
		return
	}
	if opts.SkipEval != nil && *opts.SkipEval {
		skipEval = true
	}
	if len(opts.Kvs) == 0 {
		return
	}

	duplicate := make(map[string]struct{}, len(opts.Kvs))

	var buf pkbuffer.Buffer

	buf.WriteString(pkfield.FieldGoType(p.g, field))
	buf.WriteString("{")
	buf.WriteStringLn("")

	for _, kv := range opts.Kvs {
		_, key := p.genCodeByFileType(field.Message.Fields[0], kv.Key, "", skipEval)
		_, value := p.genCodeByFileType(field.Message.Fields[1], kv.Value, "", skipEval)

		if _, ok := duplicate[key]; ok {
			err := pkerror.New("found duplicate map key: <%s>", key)
			panic(err)
		}
		duplicate[key] = struct{}{}

		buf.WriteString(key)
		buf.WriteString(": ")
		buf.WriteString(value)
		buf.WriteStringLn(", ")
	}
	buf.WriteString("}")

	var condApply string

	if opts.IgnoreEmpty != nil && *opts.IgnoreEmpty {
		condApply = "x." + field.GoName + " == nil"
	} else {
		condApply = "len(x." + field.GoName + ") == 0"
	}

	p.g.P("if ", condApply, "{")
	p.g.P("    x.", field.GoName, " = ", buf.String())
	p.g.P("}")
	return
}

func (p *Plugin) genCodeByFileType(field *protogen.Field, filedType *pbdefault.FieldType,
	varName string, skipEval bool) (cond string, value string) {

	switch field.Desc.Kind() {
	case protoreflect.Int32Kind:
		value = "0"
		typeSet := p.loadFieldTypeInt32(filedType)
		if typeSet != nil {
			cond = varName + " == 0"
			value = pkident.Int32(typeSet.Int32)
		}
	case protoreflect.Sint32Kind:
		value = "0"
		typeSet := p.loadFieldTypeSInt32(filedType)
		if typeSet != nil {
			cond = varName + " == 0"
			value = pkident.Int32(typeSet.Sint32)
		}
	case protoreflect.Sfixed32Kind:
		value = "0"
		typeSet := p.loadFieldTypeSFixed32(filedType)
		if typeSet != nil {
			cond = varName + " == 0"
			value = pkident.Int32(typeSet.Sfixed32)
		}
	case protoreflect.Int64Kind:
		value = "0"
		typeSet := p.loadFieldTypeInt64(filedType)
		if typeSet != nil {
			cond = varName + " == 0"
			value = pkident.Int64(typeSet.Int64)
		}
	case protoreflect.Sint64Kind:
		value = "0"
		typeSet := p.loadFieldTypeSInt64(filedType)
		if typeSet != nil {
			cond = varName + " == 0"
			value = pkident.Int64(typeSet.Sint64)
		}
	case protoreflect.Sfixed64Kind:
		value = "0"
		typeSet := p.loadFieldTypeSFixed64(filedType)
		if typeSet != nil {
			cond = varName + " == 0"
			value = pkident.Int64(typeSet.Sfixed64)
		}
	case protoreflect.Uint32Kind:
		value = "0"
		typeSet := p.loadFieldTypeUInt32(filedType)
		if typeSet != nil {
			cond = varName + " == 0"
			value = pkident.UInt32(typeSet.Uint32)
		}
	case protoreflect.Fixed32Kind:
		value = "0"
		typeSet := p.loadFieldTypeFixed32(filedType)
		if typeSet != nil {
			cond = varName + " == 0"
			value = pkident.UInt32(typeSet.Fixed32)
		}
	case protoreflect.Uint64Kind:
		value = "0"
		typeSet := p.loadFieldTypeUInt64(filedType)
		if typeSet != nil {
			cond = varName + " == 0"
			value = pkident.UInt64(typeSet.Uint64)
		}
	case protoreflect.Fixed64Kind:
		value = "0"
		typeSet := p.loadFieldTypeFixed64(filedType)
		if typeSet != nil {
			cond = varName + " == 0"
			value = pkident.UInt64(typeSet.Fixed64)
		}
	case protoreflect.FloatKind:
		value = "0"
		typeSet := p.loadFieldTypeFloat(filedType)
		if typeSet != nil {
			cond = varName + " == 0"
			value = pkident.Float32(typeSet.Float)
		}
	case protoreflect.DoubleKind:
		value = "0"
		typeSet := p.loadFieldTypeDouble(filedType)
		if typeSet != nil {
			cond = varName + " == 0"
			value = pkident.Float64(typeSet.Double)
		}
	case protoreflect.BoolKind:
		value = "false"
		typeSet := p.loadFieldTypeBool(filedType)
		if typeSet != nil {
			cond = varName + "== false"
			value = pkident.Bool(typeSet.Bool)
		}
	case protoreflect.EnumKind:
		value = "0"
		typeSet := p.loadFieldTypeEnum(filedType)
		if typeSet != nil {
			cond = varName + " == 0"
			value = pkident.Int32(typeSet.Enum)
		}
	case protoreflect.StringKind:
		value = `""`
		typeSet := p.loadFieldTypeString(filedType)
		if typeSet != nil {
			cond = varName + ` == ""`
			value = pkident.String(typeSet.String_)
		}
	case protoreflect.BytesKind:
		value = "[]byte{}"
		typeSet := p.loadFieldTypeBytes(filedType)
		if typeSet != nil {
			cond = "len(" + varName + ")" + " == 0"
			value = pkident.Bytes(typeSet.Bytes)
		}
	case protoreflect.MessageKind:
		value = "nil"
		switch pkwkt.Lookup(string(field.Message.Desc.FullName())) {
		case pkwkt.Any:
			typeSet := p.loadFieldTypeAny(filedType)
			// The `Import` and `Message` both are empty represents user not set, We can ignore it.
			if typeSet != nil && typeSet.Any != nil && typeSet.Any.TypeUrl != nil {
				p.generateCodeForNewAny()
				//cond = varName + " == nil"
				cond = varName + `.GetTypeUrl() == ""`
				anyValue := p.getValueForFieldAny(typeSet.Any)
				value = fmt.Sprintf("x.%s(%s, %v)", methodNewAny, anyValue, skipEval)
			}
		case pkwkt.Duration:
			typeSet := p.loadFieldTypeDuration(filedType)
			if typeSet != nil && typeSet.Duration != nil {
				cond = varName + ".AsDuration() == 0"
				goType := p.g.QualifiedGoIdent(importpkg.DurationPB.Ident("Duration"))
				value = fmt.Sprintf("&%s{Seconds: %d, Nanos: %d}", goType, typeSet.Duration.Seconds, typeSet.Duration.Nanos)
			}
		case pkwkt.Timestamp:
			typeSet := p.loadFieldTypeTimestamp(filedType)
			if typeSet != nil && typeSet.Timestamp != nil {
				cond = varName + ".AsTime().UnixNano() == 0"
				goType := p.g.QualifiedGoIdent(importpkg.TimestampPB.Ident("Timestamp"))
				value = fmt.Sprintf("&%s{Seconds: %d, Nanos: %d}", goType, typeSet.Timestamp.Seconds, typeSet.Timestamp.Nanos)
			}
		default:
			typeSet := p.loadFieldTypeMessage(filedType)
			if typeSet != nil && typeSet.Message != nil && typeSet.Message.Init != nil && *typeSet.Message.Init {
				cond = varName + " == nil"
				value = "&" + p.g.QualifiedGoIdent(field.Message.GoIdent) + "{}"
			}
		}
	default:
		panic("invalid cases.")
	}
	return
}

func (p *Plugin) generateCodeForEval(field *protogen.Field, inOneOf bool) {
	if pkfield.FieldIsOneOf(field) && !inOneOf {
		return
	}
	if !pkfield.FieldContainMessage(field) || pkwkt.Valid(field) {
		return
	}

	methodName := p.getNameDefaultMethod()

	evalMethod := func(varName string) {
		p.g.P("if ", varName, " != nil {")
		p.g.P(varName, ".", methodName, "()")
		p.g.P("}")
	}

	switch {
	case field.Desc.IsList():
		p.g.P("for _, ri := range x.", field.GoName, "{")
		evalMethod("ri")
		p.g.P("}")
	case field.Desc.IsMap():
		p.g.P("for _, mv := range x.", field.GoName, "{")
		evalMethod("mv")
		p.g.P("}")
	default:
		if inOneOf {
			evalMethod("o." + field.GoName)
		} else {
			evalMethod("x." + field.GoName)
		}
	}
}

func (p *Plugin) generateCodeForNewAny() {
	p.once.Do(func() {
		msgGoName := p.message.GoIdent.GoName
		methodName := p.getNameDefaultMethod()
		newMethod := p.g.QualifiedGoIdent(importpkg.AnyPB.Ident("New"))
		anyType := p.g.QualifiedGoIdent(importpkg.AnyPB.Ident("Any"))
		msgType := p.g.QualifiedGoIdent(importpkg.Proto.Ident("Message"))

		buf := p.extend
		buf.WriteStringLn("// ", methodNewAny, " used to create instance of type any")
		buf.WriteStringLn("func (*", msgGoName, ") ", methodNewAny, "(src ", msgType, ", skipEval bool) *", anyType, " {")
		buf.WriteStringLn("    if !skipEval {")
		buf.WriteStringLn("        if dt, ok := interface{}(src).(interface {", methodName, "() }); ok {")
		buf.WriteStringLn("            dt.", methodName, "()")
		buf.WriteStringLn("        }")
		buf.WriteStringLn("    }")
		buf.WriteStringLn("    v, err := ", newMethod, "(src)")
		buf.WriteStringLn("    if err != nil {") // Ignore error ?
		buf.WriteStringLn("        panic(err)")
		buf.WriteStringLn("    }")
		buf.WriteStringLn("    return v")
		buf.WriteStringLn("}")
	})
}
