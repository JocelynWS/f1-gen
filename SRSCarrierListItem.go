package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SRSCarrierListItem struct {
	PointA                    int64                `lb:0,ub:3279165,mandatory`
	UplinkChannelBWPerSCSList []SCSSpecificCarrier `lb:0,ub:maxnoSCSs,mandatory`
	ActiveULBWP               ActiveULBWP          `mandatory`
	Pci                       *NRPCI               `optional`
	// IEExtensions * `optional`
}

func (ie *SRSCarrierListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.Pci != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_PointA := NewINTEGER(ie.PointA, aper.Constraint{Lb: 0, Ub: 3279165}, false)
	if err = tmp_PointA.Encode(w); err != nil {
		err = utils.WrapError("Encode PointA", err)
		return
	}
	if len(ie.UplinkChannelBWPerSCSList) > 0 {
		tmp := Sequence[*SCSSpecificCarrier]{
			Value: []*SCSSpecificCarrier{},
			c:     aper.Constraint{Lb: 0, Ub: maxnoSCSs},
			ext:   false,
		}
		for _, i := range ie.UplinkChannelBWPerSCSList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode UplinkChannelBWPerSCSList", err)
			return
		}
	} else {
		err = utils.WrapError("UplinkChannelBWPerSCSList is nil", err)
		return
	}
	if err = ie.ActiveULBWP.Encode(w); err != nil {
		err = utils.WrapError("Encode ActiveULBWP", err)
		return
	}
	if ie.Pci != nil {
		if err = ie.Pci.Encode(w); err != nil {
			err = utils.WrapError("Encode Pci", err)
			return
		}
	}
	return
}
func (ie *SRSCarrierListItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_PointA := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 3279165},
		ext: false,
	}
	if err = tmp_PointA.Decode(r); err != nil {
		err = utils.WrapError("Read PointA", err)
		return
	}
	ie.PointA = int64(tmp_PointA.Value)
	tmp_UplinkChannelBWPerSCSList := Sequence[*SCSSpecificCarrier]{
		c:   aper.Constraint{Lb: 0, Ub: maxnoSCSs},
		ext: false,
	}
	fn := func() *SCSSpecificCarrier { return new(SCSSpecificCarrier) }
	if err = tmp_UplinkChannelBWPerSCSList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read UplinkChannelBWPerSCSList", err)
		return
	}
	ie.UplinkChannelBWPerSCSList = []SCSSpecificCarrier{}
	for _, i := range tmp_UplinkChannelBWPerSCSList.Value {
		ie.UplinkChannelBWPerSCSList = append(ie.UplinkChannelBWPerSCSList, *i)
	}
	if err = ie.ActiveULBWP.Decode(r); err != nil {
		err = utils.WrapError("Read ActiveULBWP", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(NRPCI)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read Pci", err)
			return
		}
		ie.Pci = tmp
	}
	return
}
