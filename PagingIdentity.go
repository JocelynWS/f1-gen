package f1ap

import "github.com/lvdund/ngap/aper"

const (
	PagingIdentityPresentNothing uint64 = iota
	PagingIdentityPresentRANUEPagingIdentity
	PagingIdentityPresentCNUEPagingIdentity
)

type PagingIdentity struct {
	Choice              uint64
	RANUEPagingIdentity *RANUEPagingIdentity
	CNUEPagingIdentity  *CNUEPagingIdentity
	// ChoiceExtension
}

func (ie *PagingIdentity) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case PagingIdentityPresentRANUEPagingIdentity:
		err = ie.RANUEPagingIdentity.Encode(w)
	case PagingIdentityPresentCNUEPagingIdentity:
		err = ie.CNUEPagingIdentity.Encode(w)
	}
	return
}

func (ie *PagingIdentity) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case PagingIdentityPresentRANUEPagingIdentity:
		var tmp RANUEPagingIdentity
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.RANUEPagingIdentity = &tmp
	case PagingIdentityPresentCNUEPagingIdentity:
		var tmp CNUEPagingIdentity
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CNUEPagingIdentity = &tmp
	}
	return
}
