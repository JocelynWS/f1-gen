package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type BAPlayerBHRLCchannelMappingInfo struct {
	BAPlayerBHRLCchannelMappingInfoToAdd    []BAPlayerBHRLCchannelMappingInfoItem `lb:1,ub:maxnoofMappingEntries,optional`
	BAPlayerBHRLCchannelMappingInfoToRemove []MappingInformationIndex             `lb:1,ub:maxnoofMappingEntries,optional`
	// IEExtensions * `optional`
}

func (ie *BAPlayerBHRLCchannelMappingInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.BAPlayerBHRLCchannelMappingInfoToAdd != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.BAPlayerBHRLCchannelMappingInfoToRemove != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if len(ie.BAPlayerBHRLCchannelMappingInfoToAdd) > 0 {
		tmp := Sequence[*BAPlayerBHRLCchannelMappingInfoItem]{
			Value: []*BAPlayerBHRLCchannelMappingInfoItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofMappingEntries},
			ext:   false,
		}
		for _, i := range ie.BAPlayerBHRLCchannelMappingInfoToAdd {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode BAPlayerBHRLCchannelMappingInfoToAdd", err)
			return
		}
	}
	if len(ie.BAPlayerBHRLCchannelMappingInfoToRemove) > 0 {
		tmp := Sequence[*MappingInformationIndex]{
			Value: []*MappingInformationIndex{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofMappingEntries},
			ext:   false,
		}
		for _, i := range ie.BAPlayerBHRLCchannelMappingInfoToRemove {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode BAPlayerBHRLCchannelMappingInfoToRemove", err)
			return
		}
	}
	return
}
func (ie *BAPlayerBHRLCchannelMappingInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_BAPlayerBHRLCchannelMappingInfoToAdd := Sequence[*BAPlayerBHRLCchannelMappingInfoItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMappingEntries},
			ext: false,
		}
		fn := func() *BAPlayerBHRLCchannelMappingInfoItem { return new(BAPlayerBHRLCchannelMappingInfoItem) }
		if err = tmp_BAPlayerBHRLCchannelMappingInfoToAdd.Decode(r, fn); err != nil {
			err = utils.WrapError("Read BAPlayerBHRLCchannelMappingInfoToAdd", err)
			return
		}
		ie.BAPlayerBHRLCchannelMappingInfoToAdd = []BAPlayerBHRLCchannelMappingInfoItem{}
		for _, i := range tmp_BAPlayerBHRLCchannelMappingInfoToAdd.Value {
			ie.BAPlayerBHRLCchannelMappingInfoToAdd = append(ie.BAPlayerBHRLCchannelMappingInfoToAdd, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_BAPlayerBHRLCchannelMappingInfoToRemove := Sequence[*MappingInformationIndex]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMappingEntries},
			ext: false,
		}
		fn := func() *MappingInformationIndex { return new(MappingInformationIndex) }
		if err = tmp_BAPlayerBHRLCchannelMappingInfoToRemove.Decode(r, fn); err != nil {
			err = utils.WrapError("Read BAPlayerBHRLCchannelMappingInfoToRemove", err)
			return
		}
		ie.BAPlayerBHRLCchannelMappingInfoToRemove = []MappingInformationIndex{}
		for _, i := range tmp_BAPlayerBHRLCchannelMappingInfoToRemove.Value {
			ie.BAPlayerBHRLCchannelMappingInfoToRemove = append(ie.BAPlayerBHRLCchannelMappingInfoToRemove, *i)
		}
	}
	return
}
