package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SliceToReportItem struct {
	PLMNIdentity []byte       `lb:3,ub:3,mandatory`
	SNSSAIlist   []SNSSAIItem `lb:1,ub:maxnoofSliceItems,mandatory,valExt`
	// IEExtensions * `optional`
}

func (ie *SliceToReportItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		err = utils.WrapError("Encode PLMNIdentity", err)
		return
	}
	if len(ie.SNSSAIlist) > 0 {
		tmp := Sequence[*SNSSAIItem]{
			Value: []*SNSSAIItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
			ext:   true,
		}
		for _, i := range ie.SNSSAIlist {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode SNSSAIlist", err)
			return
		}
	} else {
		err = utils.WrapError("SNSSAIlist is nil", err)
		return
	}
	return
}
func (ie *SliceToReportItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PLMNIdentity := OCTETSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = tmp_PLMNIdentity.Decode(r); err != nil {
		err = utils.WrapError("Read PLMNIdentity", err)
		return
	}
	ie.PLMNIdentity = tmp_PLMNIdentity.Value
	tmp_SNSSAIlist := Sequence[*SNSSAIItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
		ext: true,
	}
	fn := func() *SNSSAIItem { return new(SNSSAIItem) }
	if err = tmp_SNSSAIlist.Decode(r, fn); err != nil {
		err = utils.WrapError("Read SNSSAIlist", err)
		return
	}
	ie.SNSSAIlist = []SNSSAIItem{}
	for _, i := range tmp_SNSSAIlist.Value {
		ie.SNSSAIlist = append(ie.SNSSAIlist, *i)
	}
	return
}
