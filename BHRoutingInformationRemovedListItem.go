package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type BHRoutingInformationRemovedListItem struct {
	BAPRoutingID BAPRoutingID `mandatory`
	// IEExtensions * `optional`
}

func (ie *BHRoutingInformationRemovedListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.BAPRoutingID.Encode(w); err != nil {
		err = utils.WrapError("Encode BAPRoutingID", err)
		return
	}
	return
}
func (ie *BHRoutingInformationRemovedListItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.BAPRoutingID.Decode(r); err != nil {
		err = utils.WrapError("Read BAPRoutingID", err)
		return
	}
	return
}
