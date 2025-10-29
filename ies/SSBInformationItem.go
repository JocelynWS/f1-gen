package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SSBInformationItem struct {
	SSBConfiguration SSBTFConfiguration `mandatory`
	PCINR            NRPCI              `mandatory`
	// IEExtensions * `optional`
}

func (ie *SSBInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.SSBConfiguration.Encode(w); err != nil {
		err = utils.WrapError("Encode SSBConfiguration", err)
		return
	}
	if err = ie.PCINR.Encode(w); err != nil {
		err = utils.WrapError("Encode PCINR", err)
		return
	}
	return
}
func (ie *SSBInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.SSBConfiguration.Decode(r); err != nil {
		err = utils.WrapError("Read SSBConfiguration", err)
		return
	}
	if err = ie.PCINR.Decode(r); err != nil {
		err = utils.WrapError("Read PCINR", err)
		return
	}
	return
}
