package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ULBHNonUPTrafficMappingItem struct {
	NonUPTrafficType NonUPTrafficType `mandatory`
	BHInfo           BHInfo           `mandatory`
	// IEExtensions * `optional`
}

func (ie *ULBHNonUPTrafficMappingItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.NonUPTrafficType.Encode(w); err != nil {
		err = utils.WrapError("Encode NonUPTrafficType", err)
		return
	}
	if err = ie.BHInfo.Encode(w); err != nil {
		err = utils.WrapError("Encode BHInfo", err)
		return
	}
	return
}
func (ie *ULBHNonUPTrafficMappingItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.NonUPTrafficType.Decode(r); err != nil {
		err = utils.WrapError("Read NonUPTrafficType", err)
		return
	}
	if err = ie.BHInfo.Decode(r); err != nil {
		err = utils.WrapError("Read BHInfo", err)
		return
	}
	return
}
