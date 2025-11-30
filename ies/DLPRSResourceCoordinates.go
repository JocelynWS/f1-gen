package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DLPRSResourceCoordinates struct {
	ListofDLPRSResourceSetARP []DLPRSResourceSetARP `mandatory,lb:1,ub:maxnoofPRSResourceSets`
	// IEExtensions * `optional,ignore`
}

func (ie *DLPRSResourceCoordinates) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp := Sequence[*DLPRSResourceSetARP]{Value: []*DLPRSResourceSetARP{}, c: aper.Constraint{Lb: 1, Ub: maxnoofPRSResourceSets}, ext: true}
	for i := range ie.ListofDLPRSResourceSetARP {
		tmp.Value = append(tmp.Value, &ie.ListofDLPRSResourceSetARP[i])
	}
	if err = tmp.Encode(w); err != nil {
		err = utils.WrapError("Encode ListofDLPRSResourceSetARP", err)
		return
	}
	return
}

func (ie *DLPRSResourceCoordinates) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp := Sequence[*DLPRSResourceSetARP]{c: aper.Constraint{Lb: 1, Ub: maxnoofPRSResourceSets}, ext: true}
	fn := func() *DLPRSResourceSetARP { return new(DLPRSResourceSetARP) }
	if err = tmp.Decode(r, fn); err != nil {
		err = utils.WrapError("Read ListofDLPRSResourceSetARP", err)
		return
	}
	ie.ListofDLPRSResourceSetARP = []DLPRSResourceSetARP{}
	for _, v := range tmp.Value {
		ie.ListofDLPRSResourceSetARP = append(ie.ListofDLPRSResourceSetARP, *v)
	}
	return
}
