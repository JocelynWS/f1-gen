package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DLPRSResourceSetARP struct {
	DlPRSResourceSetID          PRSResourceSetID            `mandatory`
	DLPRSResourceSetARPLocation DLPRSResourceSetARPLocation `mandatory`
	ListofDLPRSResourceARP      []DLPRSResourceARP          `mandatory,lb:1,ub:maxnoofPRSResourcesPerSet`
	// IEExtensions * `optional,ignore`
}

func (ie *DLPRSResourceSetARP) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.DlPRSResourceSetID.Encode(w); err != nil {
		err = utils.WrapError("Encode DlPRSResourceSetID", err)
		return
	}
	if err = ie.DLPRSResourceSetARPLocation.Encode(w); err != nil {
		err = utils.WrapError("Encode DLPRSResourceSetARPLocation", err)
		return
	}
	tmp := Sequence[*DLPRSResourceARP]{Value: []*DLPRSResourceARP{}, c: aper.Constraint{Lb: 1, Ub: maxnoofPRSResourcesPerSet}, ext: true}
	for i := range ie.ListofDLPRSResourceARP {
		tmp.Value = append(tmp.Value, &ie.ListofDLPRSResourceARP[i])
	}
	if err = tmp.Encode(w); err != nil {
		err = utils.WrapError("Encode ListofDLPRSResourceARP", err)
		return
	}
	return
}

func (ie *DLPRSResourceSetARP) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.DlPRSResourceSetID.Decode(r); err != nil {
		err = utils.WrapError("Read DlPRSResourceSetID", err)
		return
	}
	if err = ie.DLPRSResourceSetARPLocation.Decode(r); err != nil {
		err = utils.WrapError("Read DLPRSResourceSetARPLocation", err)
		return
	}
	tmp := Sequence[*DLPRSResourceARP]{c: aper.Constraint{Lb: 1, Ub: maxnoofPRSResourcesPerSet}, ext: true}
	fn := func() *DLPRSResourceARP { return new(DLPRSResourceARP) }
	if err = tmp.Decode(r, fn); err != nil {
		err = utils.WrapError("Read ListofDLPRSResourceARP", err)
		return
	}
	ie.ListofDLPRSResourceARP = []DLPRSResourceARP{}
	for _, v := range tmp.Value {
		ie.ListofDLPRSResourceARP = append(ie.ListofDLPRSResourceARP, *v)
	}
	return
}
