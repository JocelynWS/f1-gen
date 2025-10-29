package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GeographicalCoordinates struct {
	TRPPositionDefinitionType TRPPositionDefinitionType `mandatory`
	DLPRSResourceCoordinates  *DLPRSResourceCoordinates `optional`
	// IEExtensions * `optional`
}

func (ie *GeographicalCoordinates) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DLPRSResourceCoordinates != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.TRPPositionDefinitionType.Encode(w); err != nil {
		err = utils.WrapError("Encode TRPPositionDefinitionType", err)
		return
	}
	if ie.DLPRSResourceCoordinates != nil {
		if err = ie.DLPRSResourceCoordinates.Encode(w); err != nil {
			err = utils.WrapError("Encode DLPRSResourceCoordinates", err)
			return
		}
	}
	return
}
func (ie *GeographicalCoordinates) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.TRPPositionDefinitionType.Decode(r); err != nil {
		err = utils.WrapError("Read TRPPositionDefinitionType", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(DLPRSResourceCoordinates)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read DLPRSResourceCoordinates", err)
			return
		}
		ie.DLPRSResourceCoordinates = tmp
	}
	return
}
