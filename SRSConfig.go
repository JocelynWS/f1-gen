package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SRSConfig struct {
	SRSResourceList       []SRSResource           `lb:1,ub:maxnoSRSResources,optional,valExt`
	PosSRSResourceList    []PosSRSResourceItem    `lb:1,ub:maxnoSRSPosResources,optional,valExt`
	SRSResourceSetList    []SRSResourceSetItem    `lb:1,ub:maxnoSRSResources,optional,valExt`
	PosSRSResourceSetList []PosSRSResourceSetItem `lb:1,ub:maxnoSRSPosResourceSetss,optional,valExt`
	// IEExtensions * `optional`
}

func (ie *SRSConfig) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.SRSResourceList != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.PosSRSResourceList != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.SRSResourceSetList != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.PosSRSResourceSetList != nil {
		aper.SetBit(optionals, 4)
	}
	w.WriteBits(optionals, 5)
	if len(ie.SRSResourceList) > 0 {
		tmp := Sequence[*SRSResource]{
			Value: []*SRSResource{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoSRSResources},
			ext:   true,
		}
		for _, i := range ie.SRSResourceList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode SRSResourceList", err)
			return
		}
	}
	if len(ie.PosSRSResourceList) > 0 {
		tmp := Sequence[*PosSRSResourceItem]{
			Value: []*PosSRSResourceItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoSRSPosResources},
			ext:   true,
		}
		for _, i := range ie.PosSRSResourceList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSRSResourceList", err)
			return
		}
	}
	if len(ie.SRSResourceSetList) > 0 {
		tmp := Sequence[*SRSResourceSetItem]{
			Value: []*SRSResourceSetItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoSRSResources},
			ext:   true,
		}
		for _, i := range ie.SRSResourceSetList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode SRSResourceSetList", err)
			return
		}
	}
	if len(ie.PosSRSResourceSetList) > 0 {
		tmp := Sequence[*PosSRSResourceSetItem]{
			Value: []*PosSRSResourceSetItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoSRSPosResourceSets},
			ext:   true,
		}
		for _, i := range ie.PosSRSResourceSetList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode PosSRSResourceSetList", err)
			return
		}
	}
	return
}
func (ie *SRSConfig) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(5); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_SRSResourceList := Sequence[*SRSResource]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoSRSResources},
			ext: true,
		}
		fn := func() *SRSResource { return new(SRSResource) }
		if err = tmp_SRSResourceList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read SRSResourceList", err)
			return
		}
		ie.SRSResourceList = []SRSResource{}
		for _, i := range tmp_SRSResourceList.Value {
			ie.SRSResourceList = append(ie.SRSResourceList, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_PosSRSResourceList := Sequence[*PosSRSResourceItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoSRSPosResources},
			ext: true,
		}
		fn := func() *PosSRSResourceItem { return new(PosSRSResourceItem) }
		if err = tmp_PosSRSResourceList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read PosSRSResourceList", err)
			return
		}
		ie.PosSRSResourceList = []PosSRSResourceItem{}
		for _, i := range tmp_PosSRSResourceList.Value {
			ie.PosSRSResourceList = append(ie.PosSRSResourceList, *i)
		}
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_SRSResourceSetList := Sequence[*SRSResourceSetItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoSRSResourceSets},
			ext: true,
		}
		fn := func() *SRSResourceSetItem { return new(SRSResourceSetItem) }
		if err = tmp_SRSResourceSetList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read SRSResourceSetList", err)
			return
		}
		ie.SRSResourceSetList = []SRSResourceSetItem{}
		for _, i := range tmp_SRSResourceSetList.Value {
			ie.SRSResourceSetList = append(ie.SRSResourceSetList, *i)
		}
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_PosSRSResourceSetList := Sequence[*PosSRSResourceSetItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoSRSPosResourceSets},
			ext: true,
		}
		fn := func() *PosSRSResourceSetItem { return new(PosSRSResourceSetItem) }
		if err = tmp_PosSRSResourceSetList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read PosSRSResourceSetList", err)
			return
		}
		ie.PosSRSResourceSetList = []PosSRSResourceSetItem{}
		for _, i := range tmp_PosSRSResourceSetList.Value {
			ie.PosSRSResourceSetList = append(ie.PosSRSResourceSetList, *i)
		}
	}
	return
}
