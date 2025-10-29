package f1ap

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type RLCDuplicationStateItem struct {
	DuplicationState DuplicationState `mandatory`
	// IEExtensions *ProtocolExtensionContainerRLCDuplicationStateItemExtIEs `optional`
}

func (ie *RLCDuplicationStateItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if err = w.WriteBits(optionals, 1); err != nil {
		return
	}

	if err = ie.DuplicationState.Encode(w); err != nil {
		err = utils.WrapError("Encode DuplicationState", err)
		return
	}

	return
}

func (ie *RLCDuplicationStateItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	if _, err = r.ReadBits(1); err != nil {
		return
	}

	if err = ie.DuplicationState.Decode(r); err != nil {
		err = utils.WrapError("Read DuplicationState", err)
		return
	}

	return
}
