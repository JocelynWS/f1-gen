package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type RRCVersion struct {
	LatestRRCVersion aper.BitString `lb:3,ub:3,mandatory`
	// IEExtensions     *ProtocolExtensionContainer  `optional`
}

func (ie *RRCVersion) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)

	tmp_LatestRRCVersion := NewBITSTRING(ie.LatestRRCVersion, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_LatestRRCVersion.Encode(w); err != nil {
		err = utils.WrapError("Encode LatestRRCVersion", err)
		return
	}

	return
}

func (ie *RRCVersion) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}

	tmp_LatestRRCVersion := BITSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = tmp_LatestRRCVersion.Decode(r); err != nil {
		err = utils.WrapError("Read LatestRRCVersion", err)
		return
	}
	ie.LatestRRCVersion = aper.BitString{
		Bytes:   tmp_LatestRRCVersion.Value.Bytes,
		NumBits: tmp_LatestRRCVersion.Value.NumBits,
	}

	return
}
