package f1ap

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	SSBPositionsInBurstPresentShortBitmap  aper.Enumerated = 0
	SSBPositionsInBurstPresentMediumBitmap aper.Enumerated = 1
	SSBPositionsInBurstPresentLongBitmap   aper.Enumerated = 2
	SSBPositionsInBurstPresentNothing      aper.Enumerated = 3
)

type SSBPositionsInBurst struct {
	Present      aper.Enumerated
	ShortBitmap  *aper.BitString
	MediumBitmap *aper.BitString
	LongBitmap   *aper.BitString
}

func (ie *SSBPositionsInBurst) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	if err = w.WriteInteger(int64(ie.Present), &aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode Present Index", err)
	}

	switch ie.Present {
	case SSBPositionsInBurstPresentShortBitmap:
		if err = w.WriteBitString(ie.ShortBitmap.Bytes, uint(ie.ShortBitmap.NumBits), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Encode ShortBitmap", err)
		}
	case SSBPositionsInBurstPresentMediumBitmap:
		if err = w.WriteBitString(ie.MediumBitmap.Bytes, uint(ie.MediumBitmap.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode MediumBitmap", err)
		}
	case SSBPositionsInBurstPresentLongBitmap:
		if err = w.WriteBitString(ie.LongBitmap.Bytes, uint(ie.LongBitmap.NumBits), &aper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Encode LongBitmap", err)
		}
	default:
		return fmt.Errorf("Invalid CHOICE index for SSBPositionsInBurst: %d", ie.Present)
	}

	return
}

func (ie *SSBPositionsInBurst) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	var present int64
	present, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return utils.WrapError("Decode Present Index", err)
	}
	ie.Present = aper.Enumerated(present)

	switch ie.Present {
	case SSBPositionsInBurstPresentShortBitmap:
		bits, numBits, err := r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false)
		if err != nil {
			return utils.WrapError("Decode ShortBitmap", err)
		}
		ie.ShortBitmap = &aper.BitString{
			Bytes:   bits,
			NumBits: uint64(numBits),
		}
	case SSBPositionsInBurstPresentMediumBitmap:
		bits, numBits, err := r.ReadBitString(&aper.Constraint{Lb: 8, Ub: 8}, false)
		if err != nil {
			return utils.WrapError("Decode MediumBitmap", err)
		}
		ie.MediumBitmap = &aper.BitString{
			Bytes:   bits,
			NumBits: uint64(numBits),
		}
	case SSBPositionsInBurstPresentLongBitmap:
		bits, numBits, err := r.ReadBitString(&aper.Constraint{Lb: 64, Ub: 64}, false)
		if err != nil {
			return utils.WrapError("Decode LongBitmap", err)
		}
		ie.LongBitmap = &aper.BitString{
			Bytes:   bits,
			NumBits: uint64(numBits),
		}
	default:
		return fmt.Errorf("Invalid CHOICE index for SSBPositionsInBurst: %d", ie.Present)
	}

	return
}
