// automatically generated, do not modify

package data

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Monster struct {
	_tab flatbuffers.Table
}

func (rcv *Monster) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Monster) Pos(obj *Vec2) *Vec2 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Vec2)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Monster) Hp() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 100
}

func (rcv *Monster) Name() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.String(o)
	}
	return ""
}

func (rcv *Monster) Color() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 3
}

func (rcv *Monster) TestType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) Test(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func MonsterStart(builder *flatbuffers.Builder) { builder.StartObject(6) }
func MonsterAddPos(builder *flatbuffers.Builder, pos flatbuffers.UOffsetT) { builder.PrependStructSlot(0, flatbuffers.UOffsetT(pos), 0) }
func MonsterAddHp(builder *flatbuffers.Builder, hp int16) { builder.PrependInt16Slot(1, hp, 100) }
func MonsterAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(name), 0) }
func MonsterAddColor(builder *flatbuffers.Builder, color int8) { builder.PrependInt8Slot(3, color, 3) }
func MonsterAddTestType(builder *flatbuffers.Builder, testType byte) { builder.PrependByteSlot(4, testType, 0) }
func MonsterAddTest(builder *flatbuffers.Builder, test flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(test), 0) }
func MonsterEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
