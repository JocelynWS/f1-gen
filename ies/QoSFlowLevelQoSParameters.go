package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type QoSFlowLevelQoSParameters struct {
	QoSCharacteristics               QoSCharacteristics                  `madatory`
	NGRANAllocationRetentionPriority NGRANAllocationAndRetentionPriority `madatory`
	GBRQoSFlowInformation            *GBRQoSFlowInformation              `optional`
	ReflectiveQoSAttribute           *ReflectiveQoSAttribute             `optional,valExt`
	// IEExtensions *QoSFlowLevelQoSParametersExtIEs `optional`
}

func (ie *QoSFlowLevelQoSParameters) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.GBRQoSFlowInformation != nil {
		aper.SetBit(optionals, 0)
	}
	if ie.ReflectiveQoSAttribute != nil {
		aper.SetBit(optionals, 1)
	}
	if err = w.WriteBits(optionals, 3); err != nil {
		return
	}

	if err = ie.QoSCharacteristics.Encode(w); err != nil {
		err = utils.WrapError("Encode QoSCharacteristics", err)
		return
	}
	if err = ie.NGRANAllocationRetentionPriority.Encode(w); err != nil {
		err = utils.WrapError("Encode NGRANAllocationRetentionPriority", err)
		return
	}
	if ie.GBRQoSFlowInformation != nil {
		if err = ie.GBRQoSFlowInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode GBRQoSFlowInformation", err)
			return
		}
	}
	if ie.ReflectiveQoSAttribute != nil {
		if err = ie.ReflectiveQoSAttribute.Encode(w); err != nil {
			err = utils.WrapError("Encode ReflectiveQoSAttribute", err)
			return
		}
	}

	return
}

func (ie *QoSFlowLevelQoSParameters) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}

	if err = ie.QoSCharacteristics.Decode(r); err != nil {
		err = utils.WrapError("Read QoSCharacteristics", err)
		return
	}
	if err = ie.NGRANAllocationRetentionPriority.Decode(r); err != nil {
		err = utils.WrapError("Read NGRANAllocationRetentionPriority", err)
		return
	}

	if aper.IsBitSet(optionals, 0) {
		tmp := new(GBRQoSFlowInformation)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read GBRQoSFlowInformation", err)
			return
		}
		ie.GBRQoSFlowInformation = tmp
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(ReflectiveQoSAttribute)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read ReflectiveQoSAttribute", err)
			return
		}
		ie.ReflectiveQoSAttribute = tmp
	}

	return
}
