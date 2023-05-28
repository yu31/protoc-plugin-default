package godefault

import (
	"reflect"
	"strings"

	"github.com/yu31/protoc-kit-go/helper/pkerror"

	"github.com/yu31/protoc-plugin-default/xgo/pb/pbdefault"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func (p *Plugin) getTagName(ot interface{}) string {
	name := reflect.TypeOf(ot).Elem().Name()
	name = strings.TrimPrefix(name, "FieldType_")
	name = strings.TrimPrefix(name, "FieldOptions_")
	name = strings.TrimSuffix(name, "_")
	name = strings.ToLower(name)
	return name
}

func (p *Plugin) loadOneOfOptions(oneOf *protogen.Oneof) *pbdefault.OneOfOptions {
	i := proto.GetExtension(oneOf.Desc.Options(), pbdefault.E_Oneof)
	options := i.(*pbdefault.OneOfOptions)
	if options == nil {
		options = &pbdefault.OneOfOptions{}
	}
	return options
}

func (p *Plugin) loadFieldOptions(field *protogen.Field) *pbdefault.FieldOptions {
	i := proto.GetExtension(field.Desc.Options(), pbdefault.E_Field)
	options := i.(*pbdefault.FieldOptions)
	if options == nil {
		options = &pbdefault.FieldOptions{}
	}
	return options
}

func (p *Plugin) loadMapOptions(field *protogen.Field) *pbdefault.MapOptions {
	fieldOptions := p.loadFieldOptions(field)
	if fieldOptions == nil || fieldOptions.Kind == nil {
		return nil
	}

	switch ot := fieldOptions.Kind.(type) {
	case *pbdefault.FieldOptions_Map:
		return ot.Map
	default:
		err := pkerror.New(
			"type map only supports option kind <map> and you provided: <%s>", p.getTagName(ot),
		)
		panic(err)
	}
}

func (p *Plugin) loadRepeatedOptions(field *protogen.Field) *pbdefault.RepeatedOptions {
	fieldOptions := p.loadFieldOptions(field)
	if fieldOptions == nil || fieldOptions.Kind == nil {
		return nil
	}

	switch ot := fieldOptions.Kind.(type) {
	case *pbdefault.FieldOptions_Repeated:
		return ot.Repeated
	default:
		err := pkerror.New(
			"type repeated only supports option kind <repeated> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}

func (p *Plugin) loadPlainOptions(field *protogen.Field) *pbdefault.PlainOptions {
	fieldOptions := p.loadFieldOptions(field)
	if fieldOptions == nil || fieldOptions.Kind == nil {
		return nil
	}

	switch ot := fieldOptions.Kind.(type) {
	case *pbdefault.FieldOptions_Plain:
		return ot.Plain
	default:
		err := pkerror.New(
			"type plain only supports option kind <plain> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}

// load field type.

func (p *Plugin) loadFieldTypeInt32(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Int32 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Int32:
		return ot
	default:
		err := pkerror.New(
			"type int32 only supports option kind <int32> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeInt64(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Int64 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Int64:
		return ot
	default:
		err := pkerror.New(
			"type int64 only supports option kind <int64> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeSInt32(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Sint32 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Sint32:
		return ot
	default:
		err := pkerror.New(
			"type sint32 only supports option kind <sint32> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeSInt64(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Sint64 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Sint64:
		return ot
	default:
		err := pkerror.New(
			"type sint64 only supports option kind <sint64> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeSFixed32(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Sfixed32 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Sfixed32:
		return ot
	default:
		err := pkerror.New(
			"type sfixed32 only supports option kind <sfixed32> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeSFixed64(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Sfixed64 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Sfixed64:
		return ot
	default:
		err := pkerror.New(
			"type sfixed64 only supports option kind <sfixed64> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}

func (p *Plugin) loadFieldTypeUInt32(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Uint32 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Uint32:
		return ot
	default:
		err := pkerror.New(
			"type uin32 only supports option kind <uin32> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeUInt64(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Uint64 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Uint64:
		return ot
	default:
		err := pkerror.New(
			"type uin64 only supports option kind <uin64> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeFixed32(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Fixed32 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Fixed32:
		return ot
	default:
		err := pkerror.New(
			"type fixed32 only supports option kind <fixed32> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeFixed64(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Fixed64 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Fixed64:
		return ot
	default:
		err := pkerror.New(
			"type fixed64 only supports option kind <fixed64> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}

func (p *Plugin) loadFieldTypeFloat(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Float {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Float:
		return ot
	default:
		err := pkerror.New(
			"type float only supports option kind <float> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeDouble(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Double {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Double:
		return ot
	default:
		err := pkerror.New(
			"type double only supports option kind <double> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}

func (p *Plugin) loadFieldTypeBool(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Bool {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Bool:
		return ot
	default:
		err := pkerror.New(
			"type bool only supports option kind <bool> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeEnum(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Enum {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Enum:
		return ot
	default:
		err := pkerror.New(
			"type enum only supports option kind <enum> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeString(fieldType *pbdefault.FieldType) *pbdefault.FieldType_String_ {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_String_:
		return ot
	default:
		err := pkerror.New(
			"type string only supports option kind <string> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeBytes(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Bytes {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Bytes:
		return ot
	default:
		err := pkerror.New(
			"type bytes only supports option kind <bytes> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeMessage(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Message {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Message:
		return ot
	default:
		err := pkerror.New(
			"type message only supports option kind <message> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeAny(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Any {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Any:
		return ot
	default:
		err := pkerror.New(
			"type any only supports option kind <any> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeDuration(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Duration {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Duration:
		return ot
	default:
		err := pkerror.New(
			"type duration only supports option kind <duration> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeTimestamp(fieldType *pbdefault.FieldType) *pbdefault.FieldType_Timestamp {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefault.FieldType_Timestamp:
		return ot
	default:
		err := pkerror.New(
			"type timestamp only supports option kind <timestamp> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
