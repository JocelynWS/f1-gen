package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type FlowsMappedToSLDRBItem struct {
	Pc5QoSFlowIdentifier int64 `lb:1,ub:2048,mandatory`
	// IEExtensions * `optional`
}

func (ie *FlowsMappedToSLDRBItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_Pc5QoSFlowIdentifier := NewINTEGER(ie.Pc5QoSFlowIdentifier, aper.Constraint{Lb: 1, Ub: 2048}, false)
	if err = tmp_Pc5QoSFlowIdentifier.Encode(w); err != nil {
		err = utils.WrapError("Encode Pc5QoSFlowIdentifier", err)
		return
	}
	return
}
func (ie *FlowsMappedToSLDRBItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_Pc5QoSFlowIdentifier := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 2048},
		ext: false,
	}
	if err = tmp_Pc5QoSFlowIdentifier.Decode(r); err != nil {
		err = utils.WrapError("Read Pc5QoSFlowIdentifier", err)
		return
	}
	ie.Pc5QoSFlowIdentifier = int64(tmp_Pc5QoSFlowIdentifier.Value)
	return
}
