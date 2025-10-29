package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ConditionalInterDUMobilityInformation struct {
	CHOTrigger          CHOTriggerInterDU `mandatory`
	TargetGNBDUUEF1APID *int64            `lb:0,ub:4294967295,optional`
	// IEExtensions *ProtocolExtensionContainerConditionalInterDUMobilityInformationExtIEs `optional`
}

func (ie *ConditionalInterDUMobilityInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.TargetGNBDUUEF1APID != nil {
		aper.SetBit(optionals, 0)
	}
	w.WriteBits(optionals, 1)

	if err = ie.CHOTrigger.Encode(w); err != nil {
		err = utils.WrapError("Encode CHOTrigger", err)
		return
	}

	if ie.TargetGNBDUUEF1APID != nil {
		tmp := INTEGER{
			Value: aper.Integer(*ie.TargetGNBDUUEF1APID),
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode TargetGNBDUUEF1APID", err)
			return
		}
	}

	return
}

func (ie *ConditionalInterDUMobilityInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	var optionals []byte
	if optionals, err = r.ReadBits(1); err != nil {
		return
	}

	if err = ie.CHOTrigger.Decode(r); err != nil {
		err = utils.WrapError("Read CHOTrigger", err)
		return
	}

	if aper.IsBitSet(optionals, 0) {
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read TargetGNBDUUEF1APID", err)
			return
		}
		ie.TargetGNBDUUEF1APID = (*int64)(&tmp.Value)
	}

	return
}
