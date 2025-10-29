package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SLDRBInformation struct {
	SLDRBQoS               PC5QoSParameters         `mandatory`
	FlowsMappedToSLDRBList []FlowsMappedToSLDRBItem `lb:1,ub:maxnoofQoSFlows,mandatory,valExt`
}

func (ie *SLDRBInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 0)
	if err = ie.SLDRBQoS.Encode(w); err != nil {
		err = utils.WrapError("Encode SLDRBQoS", err)
		return
	}
	if len(ie.FlowsMappedToSLDRBList) > 0 {
		tmp := Sequence[*FlowsMappedToSLDRBItem]{
			Value: []*FlowsMappedToSLDRBItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofQoSFlows},
			ext:   true,
		}
		for _, i := range ie.FlowsMappedToSLDRBList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode FlowsMappedToSLDRBList", err)
			return
		}
	} else {
		err = utils.WrapError("FlowsMappedToSLDRBList is nil", err)
		return
	}
	return
}
func (ie *SLDRBInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if err = ie.SLDRBQoS.Decode(r); err != nil {
		err = utils.WrapError("Read SLDRBQoS", err)
		return
	}
	tmp_FlowsMappedToSLDRBList := Sequence[*FlowsMappedToSLDRBItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofQoSFlows},
		ext: true,
	}
	fn := func() *FlowsMappedToSLDRBItem { return new(FlowsMappedToSLDRBItem) }
	if err = tmp_FlowsMappedToSLDRBList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read FlowsMappedToSLDRBList", err)
		return
	}
	ie.FlowsMappedToSLDRBList = []FlowsMappedToSLDRBItem{}
	for _, i := range tmp_FlowsMappedToSLDRBList.Value {
		ie.FlowsMappedToSLDRBList = append(ie.FlowsMappedToSLDRBList, *i)
	}
	return
}
