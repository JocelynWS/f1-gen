package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NID struct {
	Value []byte
}

func NewNID(bits []byte) NID {
	if len(bits) < 6 {
		panic("NID must be 44 bits (6 bytes including padding)")
	}
	return NID{Value: bits[:6]}
}

func (i *NID) Encode(w *aper.AperWriter) (err error) {
	bits := aper.BitString{
		Bytes:   i.Value,
		NumBits: 44,
	}

	tmp_NID := NewBITSTRING(bits, aper.Constraint{Lb: 44, Ub: 44}, false)
	if err = tmp_NID.Encode(w); err != nil {
		err = utils.WrapError("Encode NID", err)
		return
	}
	return
}

func (i *NID) Decode(r *aper.AperReader) (err error) {
	tmp_NID := BITSTRING{
		c:   aper.Constraint{Lb: 44, Ub: 44},
		ext: false,
	}
	if err = tmp_NID.Decode(r); err != nil {
		err = utils.WrapError("Read NID", err)
		return
	}
	i.Value = tmp_NID.Value.Bytes
	return
}
