package f1ap

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
	for i := range ie.ListofDLPRSResourceARP {
		if err = ie.ListofDLPRSResourceARP[i].Encode(w); err != nil {
			err = utils.WrapError("Encode ListofDLPRSResourceARP", err)
			return
		}
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
	for i := range ie.ListofDLPRSResourceARP {
		if err = ie.ListofDLPRSResourceARP[i].Decode(r); err != nil {
			err = utils.WrapError("Read ListofDLPRSResourceARP", err)
			return
		}
	}
	return
}
