package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GNBCUSystemInformation struct {
	Sibtypetobeupdatedlist []SibtypetobeupdatedListItem `mandatory,lb:1,ub:maxnoofSIBTypes`
	// IEExtensions * `optional,ignore`
}

func (ie *GNBCUSystemInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp := Sequence[*SibtypetobeupdatedListItem]{Value: []*SibtypetobeupdatedListItem{}, c: aper.Constraint{Lb: 1, Ub: maxnoofSIBTypes}, ext: true}
	for i := range ie.Sibtypetobeupdatedlist {
		tmp.Value = append(tmp.Value, &ie.Sibtypetobeupdatedlist[i])
	}
	if err = tmp.Encode(w); err != nil {
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
	tmp := Sequence[*SibtypetobeupdatedListItem]{c: aper.Constraint{Lb: 1, Ub: maxnoofSIBTypes}, ext: true}
	fn := func() *SibtypetobeupdatedListItem { return new(SibtypetobeupdatedListItem) }
	if err = tmp.Decode(r, fn); err != nil {
		err = utils.WrapError("Read Sibtypetobeupdatedlist", err)
		return
	}
	ie.Sibtypetobeupdatedlist = []SibtypetobeupdatedListItem{}
	for _, v := range tmp.Value {
		ie.Sibtypetobeupdatedlist = append(ie.Sibtypetobeupdatedlist, *v)
	}
	return
}
