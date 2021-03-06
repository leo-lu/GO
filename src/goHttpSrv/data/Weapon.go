// automatically generated, do not modify

package data

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Weapon struct {
	_tab flatbuffers.Table
}

func GetRootAsWeapon(buf []byte, offset flatbuffers.UOffsetT) *Weapon {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Weapon{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *Weapon) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Weapon) Pos(obj *Vec2) *Vec2 {
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

func (rcv *Weapon) Hp() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 150
}

func WeaponStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func WeaponAddPos(builder *flatbuffers.Builder, pos flatbuffers.UOffsetT) { builder.PrependStructSlot(0, flatbuffers.UOffsetT(pos), 0) }
func WeaponAddHp(builder *flatbuffers.Builder, hp int16) { builder.PrependInt16Slot(1, hp, 150) }
func WeaponEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
