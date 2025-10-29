package f1ap

import "github.com/lvdund/ngap/aper"

type CellsToBeBroadcastListItem struct {
	CellsToBeBroadcastItem CellsToBeBroadcastItem `mandatory`
}

func (ie *CellsToBeBroadcastListItem) Encode(w *aper.AperWriter) (err error) {
	if err = ie.CellsToBeBroadcastItem.Encode(w); err != nil {
		return
	}
	return
}

func (ie *CellsToBeBroadcastListItem) Decode(r *aper.AperReader) (err error) {
	if err = ie.CellsToBeBroadcastItem.Decode(r); err != nil {
		return
	}
	return
}
