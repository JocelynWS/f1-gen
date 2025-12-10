package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DRBInformation struct {
	DRBQoS               QoSFlowLevelQoSParameters `mandatory`
	SNSSAI               SNSSAI                    `mandatory`
	NotificationControl  *NotificationControl      `optional`
	FlowsMappedToDRBList []FlowsMappedToDRBItem    `lb:1,ub:maxnoofQoSFlows,mandatory,valueExt`
	// IEExtensions * `optional`
}

func (ie *DRBInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.NotificationControl != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.DRBQoS.Encode(w); err != nil {
		err = utils.WrapError("Encode DRBQoS", err)
		return
	}
	if err = ie.SNSSAI.Encode(w); err != nil {
		err = utils.WrapError("Encode SNSSAI", err)
		return
	}
	if ie.NotificationControl != nil {
		if err = ie.NotificationControl.Encode(w); err != nil {
			err = utils.WrapError("Encode NotificationControl", err)
			return
		}
	}
	if len(ie.FlowsMappedToDRBList) > 0 {
		tmp := Sequence[*FlowsMappedToDRBItem]{
			Value: []*FlowsMappedToDRBItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofQoSFlows},
			ext:   true,
		}
		for _, i := range ie.FlowsMappedToDRBList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode FlowsMappedToDRBList", err)
			return
		}
	} else {
		err = utils.WrapError("FlowsMappedToDRBList is nil", err)
		return
	}
	return
}
func (ie *DRBInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.DRBQoS.Decode(r); err != nil {
		err = utils.WrapError("Read DRBQoS", err)
		return
	}
	if err = ie.SNSSAI.Decode(r); err != nil {
		err = utils.WrapError("Read SNSSAI", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(NotificationControl)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read NotificationControl", err)
			return
		}
		ie.NotificationControl = tmp
	}
	tmp_FlowsMappedToDRBList := Sequence[*FlowsMappedToDRBItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofQoSFlows},
		ext: true,
	}
	fn := func() *FlowsMappedToDRBItem { return new(FlowsMappedToDRBItem) }
	if err = tmp_FlowsMappedToDRBList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read FlowsMappedToDRBList", err)
		return
	}
	ie.FlowsMappedToDRBList = []FlowsMappedToDRBItem{}
	for _, i := range tmp_FlowsMappedToDRBList.Value {
		ie.FlowsMappedToDRBList = append(ie.FlowsMappedToDRBList, *i)
	}
	return
}
