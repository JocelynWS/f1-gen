package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type FlowsMappedToDRBItem struct {
	QoSFlowIdentifier         int64                     `lb:0,ub:63,mandatory`
	QoSFlowLevelQoSParameters QoSFlowLevelQoSParameters `mandatory`
	// IEExtensions * `optional`
}

func (ie *FlowsMappedToDRBItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_QoSFlowIdentifier := NewINTEGER(ie.QoSFlowIdentifier, aper.Constraint{Lb: 0, Ub: 63}, false)
	if err = tmp_QoSFlowIdentifier.Encode(w); err != nil {
		err = utils.WrapError("Encode QoSFlowIdentifier", err)
		return
	}
	if err = ie.QoSFlowLevelQoSParameters.Encode(w); err != nil {
		err = utils.WrapError("Encode QoSFlowLevelQoSParameters", err)
		return
	}
	return
}
func (ie *FlowsMappedToDRBItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_QoSFlowIdentifier := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 63},
		ext: false,
	}
	if err = tmp_QoSFlowIdentifier.Decode(r); err != nil {
		err = utils.WrapError("Read QoSFlowIdentifier", err)
		return
	}
	ie.QoSFlowIdentifier = int64(tmp_QoSFlowIdentifier.Value)
	if err = ie.QoSFlowLevelQoSParameters.Decode(r); err != nil {
		err = utils.WrapError("Read QoSFlowLevelQoSParameters", err)
		return
	}
	return
}
