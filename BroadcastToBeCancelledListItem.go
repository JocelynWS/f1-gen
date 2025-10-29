package f1ap

import (
	"github.com/lvdund/ngap/aper"
)

type BroadcastToBeCancelledListItem struct {
	BroadcastToBeCancelledItem BroadcastToBeCancelledItem `mandatory`
}

func (ie *BroadcastToBeCancelledListItem) Encode(w *aper.AperWriter) (err error) {
	return ie.BroadcastToBeCancelledItem.Encode(w)
}

func (ie *BroadcastToBeCancelledListItem) Decode(r *aper.AperReader) (err error) {
	return ie.BroadcastToBeCancelledItem.Decode(r)
}
