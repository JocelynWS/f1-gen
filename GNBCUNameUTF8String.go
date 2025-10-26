package ies

import (
	"github.com/lvdund/ngap/aper"
)

type GNBCUNameUTF8String struct {
	Value string `aper:"sizeExt,sizeLB:1,sizeUB:150"`
}

func (ie *GNBCUNameUTF8String) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString([]byte(ie.Value), &aper.Constraint{Lb: 1, Ub: 150}, true)
	return
}

func (ie *GNBCUNameUTF8String) Decode(r *aper.AperReader) (err error) {
	var v []byte
	if v, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 150}, true); err != nil {
		return
	}
	ie.Value = string(v)
	return
}
