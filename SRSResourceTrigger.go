package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SRSResourceTrigger struct {
	AperiodicSRSResourceTriggerList []AperiodicSRSResourceTrigger `lb:1,ub:maxnoofSRSTriggerStates,mandatory`
	// IEExtensions * `optional`
}

func (ie *SRSResourceTrigger) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if len(ie.AperiodicSRSResourceTriggerList) > 0 {
		tmp := Sequence[*AperiodicSRSResourceTrigger]{
			Value: []*AperiodicSRSResourceTrigger{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofSRSTriggerStates},
			ext:   false,
		}
		for _, i := range ie.AperiodicSRSResourceTriggerList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode AperiodicSRSResourceTriggerList", err)
			return
		}
	} else {
		err = utils.WrapError("AperiodicSRSResourceTriggerList is nil", err)
		return
	}
	return
}
func (ie *SRSResourceTrigger) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_AperiodicSRSResourceTriggerList := Sequence[*AperiodicSRSResourceTrigger]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofSRSTriggerStates},
		ext: false,
	}
	fn := func() *AperiodicSRSResourceTrigger { return new(AperiodicSRSResourceTrigger) }
	if err = tmp_AperiodicSRSResourceTriggerList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read AperiodicSRSResourceTriggerList", err)
		return
	}
	ie.AperiodicSRSResourceTriggerList = []AperiodicSRSResourceTrigger{}
	for _, i := range tmp_AperiodicSRSResourceTriggerList.Value {
		ie.AperiodicSRSResourceTriggerList = append(ie.AperiodicSRSResourceTriggerList, *i)
	}
	return
}
