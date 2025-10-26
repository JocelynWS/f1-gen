package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ULAoA struct {
	AzimuthAoA            int64                  `lb:0,ub:3599,mandatory`
	ZenithAoA             *int64                 `lb:0,ub:1799,optional`
	AngleCoordinateSystem *AngleCoordinateSystem `optional`
	//IEExtensions          *ProtocolExtensionContainer `optional`
}

func (ie *ULAoA) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.ZenithAoA != nil {
		aper.SetBit(optionals, 0)
	}
	if ie.AngleCoordinateSystem != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)

	tmpAzimuth := NewINTEGER(ie.AzimuthAoA, aper.Constraint{Lb: 0, Ub: 3599}, false)
	if err = tmpAzimuth.Encode(w); err != nil {
		err = utils.WrapError("Encode AzimuthAoA", err)
		return
	}

	if ie.ZenithAoA != nil {
		tmpZenith := NewINTEGER(*ie.ZenithAoA, aper.Constraint{Lb: 0, Ub: 1799}, false)
		if err = tmpZenith.Encode(w); err != nil {
			err = utils.WrapError("Encode ZenithAoA", err)
			return
		}
	}

	if ie.AngleCoordinateSystem != nil {
		if err = ie.AngleCoordinateSystem.Encode(w); err != nil {
			err = utils.WrapError("Encode AngleCoordinateSystem", err)
			return
		}
	}

	return
}

func (ie *ULAoA) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}

	tmpAzimuth := INTEGER{c: aper.Constraint{Lb: 0, Ub: 3599}, ext: false}
	if err = tmpAzimuth.Decode(r); err != nil {
		err = utils.WrapError("Decode AzimuthAoA", err)
		return
	}
	ie.AzimuthAoA = int64(tmpAzimuth.Value)

	if aper.IsBitSet(optionals, 0) {
		tmpZenith := INTEGER{c: aper.Constraint{Lb: 0, Ub: 1799}, ext: false}
		if err = tmpZenith.Decode(r); err != nil {
			err = utils.WrapError("Decode ZenithAoA", err)
			return
		}
		ie.ZenithAoA = (*int64)(&tmpZenith.Value)
	}

	if aper.IsBitSet(optionals, 1) {
		tmp := new(AngleCoordinateSystem)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Decode AngleCoordinateSystem", err)
			return
		}
		ie.AngleCoordinateSystem = tmp
	}

	return
}
