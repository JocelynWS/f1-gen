package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TRPMeasurementRequestItem struct {
	TRPID                   int64                    `lb:0,ub:4095,mandatory,valExt`
	SearchWindowInformation *SearchWindowInformation `optional`
	//IEExtensions            *ProtocolExtensionContainer `optional`
}

func (ie *TRPMeasurementRequestItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.SearchWindowInformation != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 1)

	tmp_TRPID := NewINTEGER(ie.TRPID, aper.Constraint{Lb: 0, Ub: 4095}, true)
	if err = tmp_TRPID.Encode(w); err != nil {
		err = utils.WrapError("Encode TRPID", err)
		return
	}

	if ie.SearchWindowInformation != nil {
		if err = ie.SearchWindowInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode SearchWindowInformation", err)
			return
		}
	}

	return
}

func (ie *TRPMeasurementRequestItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(1); err != nil {
		return
	}

	tmp_TRPID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4095},
		ext: true,
	}
	if err = tmp_TRPID.Decode(r); err != nil {
		err = utils.WrapError("Read TRPID", err)
		return
	}
	ie.TRPID = int64(tmp_TRPID.Value)

	if aper.IsBitSet(optionals, 1) {
		tmp := new(SearchWindowInformation)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read SearchWindowInformation", err)
			return
		}
		ie.SearchWindowInformation = tmp
	}

	return
}
