package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SRSSpatialRelation struct {
	SpatialRelationforResourceID SpatialRelationforResourceIDItem `mandatory`
	// IEExtensions *ProtocolExtensionContainerSRSSpatialRelationExtIEs `optional`
}

func (ie *SRSSpatialRelation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.SpatialRelationforResourceID.Encode(w); err != nil {
		err = utils.WrapError("Encode SpatialRelationforResourceID", err)
		return
	}
	return
}

func (ie *SRSSpatialRelation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.SpatialRelationforResourceID.Decode(r); err != nil {
		err = utils.WrapError("Read SpatialRelationforResourceID", err)
		return
	}
	return
}
