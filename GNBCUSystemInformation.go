package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GNBCUSystemInformation struct {
	Sibtypetobeupdatedlist SEQUENCE `mandatory`
	// IEExtensions * `optional`
}

func (ie *GNBCUSystemInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.Sibtypetobeupdatedlist.Encode(w); err != nil {
		err = utils.WrapError("Encode Sibtypetobeupdatedlist", err)
		return
	}
	return
}
func (ie *GNBCUSystemInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.Sibtypetobeupdatedlist.Decode(r); err != nil {
		err = utils.WrapError("Read Sibtypetobeupdatedlist", err)
		return
	}
	return
}
