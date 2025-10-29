package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type NRFreqInfo struct {
	NRARFCN        int64            `lb:0,ub:maxNRARFCN,mandatory`
	SulInformation *SULInformation  `optional`
	FreqBandListNr []FreqBandNrItem `lb:1,ub:maxnoofNrCellBands,mandatory,valExt`
	// IEExtensions *optional
}

func (ie *NRFreqInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.SulInformation != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)

	tmp_NRARFCN := NewINTEGER(ie.NRARFCN, aper.Constraint{Lb: 0, Ub: maxNRARFCN}, false)
	if err = tmp_NRARFCN.Encode(w); err != nil {
		return utils.WrapError("Encode NRARFCN", err)
	}

	if ie.SulInformation != nil {
		if err = ie.SulInformation.Encode(w); err != nil {
			return utils.WrapError("Encode SulInformation", err)
		}
	}

	tmp := Sequence[*FreqBandNrItem]{Value: []*FreqBandNrItem{}, c: aper.Constraint{Lb: 1, Ub: maxnoofNrCellBands}, ext: true}
	for i := range ie.FreqBandListNr {
		tmp.Value = append(tmp.Value, &ie.FreqBandListNr[i])
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode FreqBandListNr", err)
	}

	return
}

func (ie *NRFreqInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	optionals, err := r.ReadBits(2)
	if err != nil {
		return
	}

	tmp_NRARFCN := INTEGER{c: aper.Constraint{Lb: 0, Ub: maxNRARFCN}, ext: false}
	if err = tmp_NRARFCN.Decode(r); err != nil {
		return utils.WrapError("Read NRARFCN", err)
	}
	ie.NRARFCN = int64(tmp_NRARFCN.Value)

	if aper.IsBitSet(optionals, 0) {
		tmp := new(SULInformation)
		if err = tmp.Decode(r); err != nil {
			return utils.WrapError("Read SulInformation", err)
		}
		ie.SulInformation = tmp
	}

	tmp_Seq := Sequence[*FreqBandNrItem]{c: aper.Constraint{Lb: 1, Ub: maxnoofNrCellBands}, ext: true}
	fn := func() *FreqBandNrItem { return new(FreqBandNrItem) }
	if err = tmp_Seq.Decode(r, fn); err != nil {
		return utils.WrapError("Read FreqBandListNr", err)
	}
	ie.FreqBandListNr = []FreqBandNrItem{}
	for _, v := range tmp_Seq.Value {
		ie.FreqBandListNr = append(ie.FreqBandListNr, *v)
	}

	return
}
