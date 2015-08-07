// automatically generated, do not modify

package data

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Vec2 struct {
	_tab flatbuffers.Struct
}

func (rcv *Vec2) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Vec2) X() float32 { return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(0)) }
func (rcv *Vec2) Y() float32 { return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(4)) }

func CreateVec2(builder *flatbuffers.Builder, x float32, y float32) flatbuffers.UOffsetT {
    builder.Prep(4, 8)
    builder.PrependFloat32(y)
    builder.PrependFloat32(x)
    return builder.Offset()
}
