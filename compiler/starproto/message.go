package starproto

import (
	"fmt"
	"hash/fnv"
	"sort"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/pkg/errors"
	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
	"google.golang.org/protobuf/types/dynamicpb"
)

func NewStarProtoMessage(msg *dynamic.Message) *starProtoMessage {
	wrapper := &starProtoMessage{
		msg:       msg,
		desc:      msg.GetMessageDescriptor(),
		attrCache: make(map[string]starlark.Value),
	}

	return wrapper
}

func ToProtoMessage(val starlark.Value) (*dynamic.Message, bool) {
	if msg, ok := val.(*starProtoMessage); ok {
		return msg.msg, true
	}
	return nil, false
}

func ToDynamicPb(msg *dynamic.Message) *dynamicpb.Message {
	mdp := msg.GetMessageDescriptor().UnwrapMessage()
	dynamicMessage := dynamicpb.NewMessage(mdp)
	msg.ConvertTo(dynamicMessage)
	return dynamicMessage
}

func (msg *starProtoMessage) AsDynamicPb() *dynamicpb.Message {
	return ToDynamicPb(msg.msg)
}

// A Starlark built-in type representing a Protobuf message. Provides attributes
// for accessing message fields using their original protobuf names.
type starProtoMessage struct {
	starlark.Mapping
	msg    *dynamic.Message
	desc   *desc.MessageDescriptor
	frozen bool

	// lets the message wrapper keep track of per-field wrappers, for freezing.
	attrCache map[string]starlark.Value
}

func (msg *starProtoMessage) String() string {
	return fmt.Sprintf("<%s %s>", msg.Type(), msg.msg)
}
func (msg *starProtoMessage) Type() string         { return msg.msg.XXX_MessageName() }
func (msg *starProtoMessage) Truth() starlark.Bool { return starlark.True }

func (msg *starProtoMessage) Freeze() {
	if !msg.frozen {
		msg.frozen = true
		for _, attr := range msg.attrCache {
			attr.Freeze()
		}
	}
}

func (msg *starProtoMessage) CompareSameType(op syntax.Token, y starlark.Value, depth int) (bool, error) {
	other, ok := y.(*starProtoMessage)
	if !ok {
		return false, nil
	}

	switch op {
	case syntax.EQL:
		eql := dynamic.Equal(msg.msg, other.msg)
		return eql, nil
	case syntax.NEQ:
		eql := dynamic.Equal(msg.msg, other.msg)
		return !eql, nil
	default:
		return false, fmt.Errorf("only == and != operations are supported on protobufs, found %s", op.String())
	}
}

func (msg *starProtoMessage) Hash() (uint32, error) {
	h := fnv.New32()
	b, err := msg.msg.Marshal()
	if err != nil {
		return 0, errors.Wrap(err, "failed to hash")
	}
	_, err = h.Write(b)
	if err != nil {
		return 0, errors.Wrap(err, "failed to hash")
	}

	return h.Sum32(), nil
}

func (msg *starProtoMessage) checkMutable(verb string) error {
	if msg.frozen {
		return fmt.Errorf("cannot %s frozen message", verb)
	}
	return nil
}

func (msg *starProtoMessage) Attr(name string) (starlark.Value, error) {
	if name == "__proto__" {
		return msg, nil
	}
	if attr, ok := msg.attrCache[name]; ok {
		return attr, nil
	}
	field := msg.desc.FindFieldByName(name)
	if field == nil {
		return nil, starlark.NoSuchAttrError(fmt.Sprintf("field %s not found in message %s", name, msg.desc.GetFullyQualifiedName()))
	}

	val := &fieldValue{
		desc: field,
		msg:  msg.msg,
	}
	out := valueToStarlark(val)
	if msg.frozen {
		out.Freeze()
	}
	msg.attrCache[name] = out
	return out, nil
}

func (msg *starProtoMessage) AttrNames() []string {
	var names []string
	for _, field := range msg.desc.GetFields() {
		names = append(names, field.GetName())
	}
	sort.Strings(names)
	return names
}

func (msg *starProtoMessage) SetKey(name starlark.Value, star starlark.Value) error {
	if key, ok := name.(starlark.String); ok {
		return msg.SetField(string(key), star)
	}
	return fmt.Errorf("key type must by `string`, got %T", name)
}

func (msg *starProtoMessage) SetField(name string, star starlark.Value) error {
	field := msg.desc.FindFieldByName(name)
	if field == nil {
		return starlark.NoSuchAttrError(fmt.Sprintf("cannot set field %s not found in message %s", name, msg.desc.GetFullyQualifiedName()))
	}

	val, err := valueFromStarlark(field, star)
	if err != nil {
		return err
	}

	if err := msg.checkMutable("set field of"); err != nil {
		return err
	}

	if oneof := field.GetOneOf(); oneof != nil {
		for _, choice := range oneof.GetChoices() {
			delete(msg.attrCache, choice.GetName())
		}
	} else {
		delete(msg.attrCache, name)
	}

	return msg.msg.TrySetField(field, val)
}

var (
	_ starlark.HasAttrs    = (*starProtoMessage)(nil)
	_ starlark.HasSetKey   = (*starProtoMessage)(nil)
	_ starlark.HasSetField = (*starProtoMessage)(nil)
	_ starlark.Comparable  = (*starProtoMessage)(nil)
)
