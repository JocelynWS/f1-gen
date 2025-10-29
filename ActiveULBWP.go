package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ActiveULBWP struct {
	LocationAndBandwidth    int64             `lb:0,ub:37949,madatory,valExt`
	SubcarrierSpacing       SubcarrierSpacing `madatory,valExt`
	CyclicPrefix            CyclicPrefix      `madatory`
	TxDirectCurrentLocation int64             `lb:0,ub:3301,madatory,valExt`
	Shift7dot5kHz           *Shift7dot5kHz    `optional`
	SRSConfig               SRSConfig         `madatory`
	// IEExtensions *ActiveULBWPExtIEs `optional`
}

func (ie *ActiveULBWP) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.Shift7dot5kHz != nil {
		aper.SetBit(optionals, 0)
	}
	if err = w.WriteBits(optionals, 2); err != nil {
		return
	}
	tmp_LocationAndBandwidth := NewINTEGER(ie.LocationAndBandwidth, aper.Constraint{Lb: 0, Ub: 37949}, true)
	if err = tmp_LocationAndBandwidth.Encode(w); err != nil {
		err = utils.WrapError("Encode LocationAndBandwidth", err)
		return
	}
	if err = ie.SubcarrierSpacing.Encode(w); err != nil {
		err = utils.WrapError("Encode SubcarrierSpacing", err)
		return
	}
	if err = ie.CyclicPrefix.Encode(w); err != nil {
		err = utils.WrapError("Encode CyclicPrefix", err)
		return
	}
	tmp_TxDirectCurrentLocation := NewINTEGER(ie.TxDirectCurrentLocation, aper.Constraint{Lb: 0, Ub: 3301}, true)
	if err = tmp_TxDirectCurrentLocation.Encode(w); err != nil {
		err = utils.WrapError("Encode TxDirectCurrentLocation", err)
		return
	}
	if ie.Shift7dot5kHz != nil {
		if err = ie.Shift7dot5kHz.Encode(w); err != nil {
			err = utils.WrapError("Encode Shift7dot5kHz", err)
			return
		}
	}
	if err = ie.SRSConfig.Encode(w); err != nil {
		err = utils.WrapError("Encode SRSConfig", err)
		return
	}
	return
}

func (ie *ActiveULBWP) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_LocationAndBandwidth := INTEGER{c: aper.Constraint{Lb: 0, Ub: 37949}, ext: true}
	if err = tmp_LocationAndBandwidth.Decode(r); err != nil {
		err = utils.WrapError("Read LocationAndBandwidth", err)
		return
	}
	ie.LocationAndBandwidth = int64(tmp_LocationAndBandwidth.Value)
	if err = ie.SubcarrierSpacing.Decode(r); err != nil {
		err = utils.WrapError("Read SubcarrierSpacing", err)
		return
	}
	if err = ie.CyclicPrefix.Decode(r); err != nil {
		err = utils.WrapError("Read CyclicPrefix", err)
		return
	}
	tmp_TxDirectCurrentLocation := INTEGER{c: aper.Constraint{Lb: 0, Ub: 3301}, ext: true}
	if err = tmp_TxDirectCurrentLocation.Decode(r); err != nil {
		err = utils.WrapError("Read TxDirectCurrentLocation", err)
		return
	}
	ie.TxDirectCurrentLocation = int64(tmp_TxDirectCurrentLocation.Value)
	if aper.IsBitSet(optionals, 0) {
		tmp := new(Shift7dot5kHz)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read Shift7dot5kHz", err)
			return
		}
		ie.Shift7dot5kHz = tmp
	}
	if err = ie.SRSConfig.Decode(r); err != nil {
		err = utils.WrapError("Read SRSConfig", err)
		return
	}
	return
}
