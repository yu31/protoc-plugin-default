package godefaults

import (
	"reflect"
	"strings"

	"github.com/yu31/protoc-kit-go/helper/pkerror"

	"github.com/yu31/protoc-plugin-defaults/xgo/pb/pbdefaults"

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

func (p *Plugin) loadOneOfOptions(oneOf *protogen.Oneof) *pbdefaults.OneOfOptions {
	i := proto.GetExtension(oneOf.Desc.Options(), pbdefaults.E_Oneof)
	options := i.(*pbdefaults.OneOfOptions)
	if options == nil {
		options = &pbdefaults.OneOfOptions{}
	}
	return options
}

func (p *Plugin) loadFieldOptions(field *protogen.Field) *pbdefaults.FieldOptions {
	i := proto.GetExtension(field.Desc.Options(), pbdefaults.E_Field)
	options := i.(*pbdefaults.FieldOptions)
	if options == nil {
		options = &pbdefaults.FieldOptions{}
	}
	return options
}

func (p *Plugin) loadMapOptions(field *protogen.Field) *pbdefaults.MapOptions {
	fieldOptions := p.loadFieldOptions(field)
	if fieldOptions == nil || fieldOptions.Kind == nil {
		return nil
	}

	switch ot := fieldOptions.Kind.(type) {
	case *pbdefaults.FieldOptions_Map:
		return ot.Map
	default:
		err := pkerror.New(
			"type map only supports option kind <map> and you provided: <%s>", p.getTagName(ot),
		)
		panic(err)
	}
}

func (p *Plugin) loadRepeatedOptions(field *protogen.Field) *pbdefaults.RepeatedOptions {
	fieldOptions := p.loadFieldOptions(field)
	if fieldOptions == nil || fieldOptions.Kind == nil {
		return nil
	}

	switch ot := fieldOptions.Kind.(type) {
	case *pbdefaults.FieldOptions_Repeated:
		return ot.Repeated
	default:
		err := pkerror.New(
			"type repeated only supports option kind <repeated> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}

func (p *Plugin) loadPlainOptions(field *protogen.Field) *pbdefaults.PlainOptions {
	fieldOptions := p.loadFieldOptions(field)
	if fieldOptions == nil || fieldOptions.Kind == nil {
		return nil
	}

	switch ot := fieldOptions.Kind.(type) {
	case *pbdefaults.FieldOptions_Plain:
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

func (p *Plugin) loadFieldTypeInt32(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Int32 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Int32:
		return ot
	default:
		err := pkerror.New(
			"type int32 only supports option kind <int32> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeInt64(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Int64 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Int64:
		return ot
	default:
		err := pkerror.New(
			"type int64 only supports option kind <int64> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeSInt32(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Sint32 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Sint32:
		return ot
	default:
		err := pkerror.New(
			"type sint32 only supports option kind <sint32> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeSInt64(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Sint64 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Sint64:
		return ot
	default:
		err := pkerror.New(
			"type sint64 only supports option kind <sint64> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeSFixed32(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Sfixed32 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Sfixed32:
		return ot
	default:
		err := pkerror.New(
			"type sfixed32 only supports option kind <sfixed32> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeSFixed64(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Sfixed64 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Sfixed64:
		return ot
	default:
		err := pkerror.New(
			"type sfixed64 only supports option kind <sfixed64> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}

func (p *Plugin) loadFieldTypeUInt32(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Uint32 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Uint32:
		return ot
	default:
		err := pkerror.New(
			"type uin32 only supports option kind <uin32> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeUInt64(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Uint64 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Uint64:
		return ot
	default:
		err := pkerror.New(
			"type uin64 only supports option kind <uin64> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeFixed32(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Fixed32 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Fixed32:
		return ot
	default:
		err := pkerror.New(
			"type fixed32 only supports option kind <fixed32> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeFixed64(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Fixed64 {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Fixed64:
		return ot
	default:
		err := pkerror.New(
			"type fixed64 only supports option kind <fixed64> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}

func (p *Plugin) loadFieldTypeFloat(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Float {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Float:
		return ot
	default:
		err := pkerror.New(
			"type float only supports option kind <float> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeDouble(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Double {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Double:
		return ot
	default:
		err := pkerror.New(
			"type double only supports option kind <double> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}

func (p *Plugin) loadFieldTypeBool(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Bool {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Bool:
		return ot
	default:
		err := pkerror.New(
			"type bool only supports option kind <bool> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeEnum(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Enum {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Enum:
		return ot
	default:
		err := pkerror.New(
			"type enum only supports option kind <enum> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeString(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_String_ {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_String_:
		return ot
	default:
		err := pkerror.New(
			"type string only supports option kind <string> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeBytes(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Bytes {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Bytes:
		return ot
	default:
		err := pkerror.New(
			"type bytes only supports option kind <bytes> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeMessage(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Message {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Message:
		return ot
	default:
		err := pkerror.New(
			"type message only supports option kind <message> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeAny(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Any {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Any:
		return ot
	default:
		err := pkerror.New(
			"type any only supports option kind <any> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeDuration(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Duration {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Duration:
		return ot
	default:
		err := pkerror.New(
			"type duration only supports option kind <duration> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
func (p *Plugin) loadFieldTypeTimestamp(fieldType *pbdefaults.FieldType) *pbdefaults.FieldType_Timestamp {
	if fieldType == nil || fieldType.Kind == nil {
		return nil
	}

	switch ot := fieldType.Kind.(type) {
	case *pbdefaults.FieldType_Timestamp:
		return ot
	default:
		err := pkerror.New(
			"type timestamp only supports option kind <timestamp> and you provided: <%s>",
			p.getTagName(ot),
		)
		panic(err)
	}
}
