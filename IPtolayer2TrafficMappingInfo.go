package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type IPtolayer2TrafficMappingInfo struct {
	IPtolayer2TrafficMappingInfoToAdd    []IPtolayer2TrafficMappingInfoItem `lb:1,ub:maxnoofMappingEntries,optional`
	IPtolayer2TrafficMappingInfoToRemove []MappingInformationIndex          `lb:1,ub:maxnoofMappingEntries,optional`
	// IEExtensions * `optional`
}

func (ie *IPtolayer2TrafficMappingInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.IPtolayer2TrafficMappingInfoToAdd != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.IPtolayer2TrafficMappingInfoToRemove != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if len(ie.IPtolayer2TrafficMappingInfoToAdd) > 0 {
		tmp := Sequence[*IPtolayer2TrafficMappingInfoItem]{
			Value: []*IPtolayer2TrafficMappingInfoItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofMappingEntries},
			ext:   false,
		}
		for _, i := range ie.IPtolayer2TrafficMappingInfoToAdd {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode IPtolayer2TrafficMappingInfoToAdd", err)
			return
		}
	}
	if len(ie.IPtolayer2TrafficMappingInfoToRemove) > 0 {
		tmp := Sequence[*MappingInformationIndex]{
			Value: []*MappingInformationIndex{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofMappingEntries},
			ext:   false,
		}
		for _, i := range ie.IPtolayer2TrafficMappingInfoToRemove {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode IPtolayer2TrafficMappingInfoToRemove", err)
			return
		}
	}
	return
}
func (ie *IPtolayer2TrafficMappingInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_IPtolayer2TrafficMappingInfoToAdd := Sequence[*IPtolayer2TrafficMappingInfoItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMappingEntries},
			ext: false,
		}
		fn := func() *IPtolayer2TrafficMappingInfoItem { return new(IPtolayer2TrafficMappingInfoItem) }
		if err = tmp_IPtolayer2TrafficMappingInfoToAdd.Decode(r, fn); err != nil {
			err = utils.WrapError("Read IPtolayer2TrafficMappingInfoToAdd", err)
			return
		}
		ie.IPtolayer2TrafficMappingInfoToAdd = []IPtolayer2TrafficMappingInfoItem{}
		for _, i := range tmp_IPtolayer2TrafficMappingInfoToAdd.Value {
			ie.IPtolayer2TrafficMappingInfoToAdd = append(ie.IPtolayer2TrafficMappingInfoToAdd, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_IPtolayer2TrafficMappingInfoToRemove := Sequence[*MappingInformationIndex]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMappingEntries},
			ext: false,
		}
		fn := func() *MappingInformationIndex { return new(MappingInformationIndex) }
		if err = tmp_IPtolayer2TrafficMappingInfoToRemove.Decode(r, fn); err != nil {
			err = utils.WrapError("Read IPtolayer2TrafficMappingInfoToRemove", err)
			return
		}
		ie.IPtolayer2TrafficMappingInfoToRemove = []MappingInformationIndex{}
		for _, i := range tmp_IPtolayer2TrafficMappingInfoToRemove.Value {
			ie.IPtolayer2TrafficMappingInfoToRemove = append(ie.IPtolayer2TrafficMappingInfoToRemove, *i)
		}
	}
	return
}
