package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type MDTConfiguration struct {
	MdtActivation          MDTActivation    `mandatory`
	MeasurementsToActivate aper.BitString   `lb:8,ub:8,mandatory,valueExt`
	M2Configuration        *M2Configuration `optional`
	M5Configuration        *M5Configuration `optional`
	M6Configuration        *M6Configuration `optional`
	M7Configuration        *M7Configuration `optional`
	// IEExtensions *optional
}

func (ie *MDTConfiguration) Encode(w *aper.AperWriter) (err error) {
	optionals := []byte{0x0} // no IEExtensions
	w.WriteBits(optionals, 1)

	if err = ie.MdtActivation.Encode(w); err != nil {
		return utils.WrapError("Encode MdtActivation", err)
	}

	tmp_MeasurementsToActivate := NewBITSTRING(ie.MeasurementsToActivate, aper.Constraint{Lb: 8, Ub: 8}, true)
	if err = tmp_MeasurementsToActivate.Encode(w); err != nil {
		return utils.WrapError("Encode MeasurementsToActivate", err)
	}

	if ie.M2Configuration != nil {
		if err = ie.M2Configuration.Encode(w); err != nil {
			return utils.WrapError("Encode M2Configuration", err)
		}
	}
	if ie.M5Configuration != nil {
		if err = ie.M5Configuration.Encode(w); err != nil {
			return utils.WrapError("Encode M5Configuration", err)
		}
	}
	if ie.M6Configuration != nil {
		if err = ie.M6Configuration.Encode(w); err != nil {
			return utils.WrapError("Encode M6Configuration", err)
		}
	}
	if ie.M7Configuration != nil {
		if err = ie.M7Configuration.Encode(w); err != nil {
			return utils.WrapError("Encode M7Configuration", err)
		}
	}
	return
}

func (ie *MDTConfiguration) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBits(1); err != nil {
		return
	}

	if err = ie.MdtActivation.Decode(r); err != nil {
		return utils.WrapError("Read MdtActivation", err)
	}

	tmp_MeasurementsToActivate := BITSTRING{c: aper.Constraint{Lb: 8, Ub: 8}, ext: true}
	if err = tmp_MeasurementsToActivate.Decode(r); err != nil {
		return utils.WrapError("Read MeasurementsToActivate", err)
	}
	ie.MeasurementsToActivate = aper.BitString{
		Bytes:   tmp_MeasurementsToActivate.Value.Bytes,
		NumBits: tmp_MeasurementsToActivate.Value.NumBits,
	}

	if aper.IsBitSet(ie.MeasurementsToActivate.Bytes, 1) {
		tmp := new(M2Configuration)
		if err = tmp.Decode(r); err != nil {
			return utils.WrapError("Read M2Configuration", err)
		}
		ie.M2Configuration = tmp
	}
	if aper.IsBitSet(ie.MeasurementsToActivate.Bytes, 4) {
		tmp := new(M5Configuration)
		if err = tmp.Decode(r); err != nil {
			return utils.WrapError("Read M5Configuration", err)
		}
		ie.M5Configuration = tmp
	}
	if aper.IsBitSet(ie.MeasurementsToActivate.Bytes, 6) {
		tmp := new(M6Configuration)
		if err = tmp.Decode(r); err != nil {
			return utils.WrapError("Read M6Configuration", err)
		}
		ie.M6Configuration = tmp
	}
	if aper.IsBitSet(ie.MeasurementsToActivate.Bytes, 7) {
		tmp := new(M7Configuration)
		if err = tmp.Decode(r); err != nil {
			return utils.WrapError("Read M7Configuration", err)
		}
		ie.M7Configuration = tmp
	}

	return
}
