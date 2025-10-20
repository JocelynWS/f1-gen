package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DUCURIMInformation struct {
	VictimgNBSetID       aper.BitString          `lb:22,ub:22,mandatory`
	RIMRSDetectionStatus RIMRSDetectionStatus    `mandatory`
	AggressorCellList    []AggressorCellListItem `lb:1,ub:maxCellingNBDU,mandatory`
	// IEExtensions * `optional`
}

func (ie *DUCURIMInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_VictimgNBSetID := NewBITSTRING(ie.VictimgNBSetID, aper.Constraint{Lb: 22, Ub: 22}, false)
	if err = tmp_VictimgNBSetID.Encode(w); err != nil {
		err = utils.WrapError("Encode VictimgNBSetID", err)
		return
	}
	if err = ie.RIMRSDetectionStatus.Encode(w); err != nil {
		err = utils.WrapError("Encode RIMRSDetectionStatus", err)
		return
	}
	if len(ie.AggressorCellList) > 0 {
		tmp := Sequence[*AggressorCellListItem]{
			Value: []*AggressorCellListItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
			ext:   false,
		}
		for _, i := range ie.AggressorCellList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode AggressorCellList", err)
			return
		}
	} else {
		err = utils.WrapError("AggressorCellList is nil", err)
		return
	}
	return
}
func (ie *DUCURIMInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_VictimgNBSetID := BITSTRING{
		c:   aper.Constraint{Lb: 22, Ub: 22},
		ext: false,
	}
	if err = tmp_VictimgNBSetID.Decode(r); err != nil {
		err = utils.WrapError("Read VictimgNBSetID", err)
		return
	}
	ie.VictimgNBSetID = aper.BitString{Bytes: tmp_VictimgNBSetID.Value.Bytes, NumBits: tmp_VictimgNBSetID.Value.NumBits}
	if err = ie.RIMRSDetectionStatus.Decode(r); err != nil {
		err = utils.WrapError("Read RIMRSDetectionStatus", err)
		return
	}
	tmp_AggressorCellList := Sequence[*AggressorCellListItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxCellingNBDU},
		ext: false,
	}
	fn := func() *AggressorCellListItem { return new(AggressorCellListItem) }
	if err = tmp_AggressorCellList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read AggressorCellList", err)
		return
	}
	ie.AggressorCellList = []AggressorCellListItem{}
	for _, i := range tmp_AggressorCellList.Value {
		ie.AggressorCellList = append(ie.AggressorCellList, *i)
	}
	return
}
