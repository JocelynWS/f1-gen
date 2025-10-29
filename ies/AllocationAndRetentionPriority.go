package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AllocationAndRetentionPriority struct {
	PriorityLevel           int64                   `lb:0,ub:15,mandatory`
	PreEmptionCapability    PreEmptionCapability    `mandatory`
	PreEmptionVulnerability PreEmptionVulnerability `mandatory`
	// IEExtensions * `optional`
}

func (ie *AllocationAndRetentionPriority) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PriorityLevel := NewINTEGER(ie.PriorityLevel, aper.Constraint{Lb: 0, Ub: 15}, false)
	if err = tmp_PriorityLevel.Encode(w); err != nil {
		err = utils.WrapError("Encode PriorityLevel", err)
		return
	}
	if err = ie.PreEmptionCapability.Encode(w); err != nil {
		err = utils.WrapError("Encode PreEmptionCapability", err)
		return
	}
	if err = ie.PreEmptionVulnerability.Encode(w); err != nil {
		err = utils.WrapError("Encode PreEmptionVulnerability", err)
		return
	}
	return
}
func (ie *AllocationAndRetentionPriority) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PriorityLevel := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 15},
		ext: false,
	}
	if err = tmp_PriorityLevel.Decode(r); err != nil {
		err = utils.WrapError("Read PriorityLevel", err)
		return
	}
	ie.PriorityLevel = int64(tmp_PriorityLevel.Value)
	if err = ie.PreEmptionCapability.Decode(r); err != nil {
		err = utils.WrapError("Read PreEmptionCapability", err)
		return
	}
	if err = ie.PreEmptionVulnerability.Decode(r); err != nil {
		err = utils.WrapError("Read PreEmptionVulnerability", err)
		return
	}
	return
}
