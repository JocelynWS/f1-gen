package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ChildNodesListItem struct {
	GNBCUUEF1APID      int64               `lb:0,ub:4294967295,mandatory,reject`
	GNBDUUEF1APID      int64               `lb:0,ub:4294967295,mandatory,reject`
	ChildNodeCellsItem *ChildNodeCellsItem `lb:1,ub:maxnoofChildIABNodes,optional`
	// IEExtensions           *ProtocolExtensionContainer `optional,ignore`
}

func (ie *ChildNodesListItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}

	optionals := []byte{0x0}
	if ie.ChildNodeCellsItem != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)

	if err = w.WriteInteger(ie.GNBCUUEF1APID, &aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return utils.WrapError("Encode GNBCUUEF1APID", err)
	}
	if err = w.WriteInteger(ie.GNBDUUEF1APID, &aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return utils.WrapError("Encode GNBDUUEF1APID", err)
	}

	if ie.ChildNodeCellsItem != nil && len(ie.ChildNodeCellsItem.Value) > 0 {
		tmp := Sequence[*ChildNodesListItem]{
			Value: []*ChildNodesListItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofChildIABNodes},
			ext:   false,
		}
		for _, item := range ie.ChildNodeCellsItem.Value {
			tmp.Value = append(tmp.Value, item)
		}
		if err = tmp.Encode(w); err != nil {
			return utils.WrapError("Encode ChildNodeCellsItem", err)
		}
	}

	return
}

func (ie *ChildNodesListItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}

	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}

	if ie.GNBCUUEF1APID, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return utils.WrapError("Decode GNBCUUEF1APID", err)
	}
	if ie.GNBDUUEF1APID, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return utils.WrapError("Decode GNBDUUEF1APID", err)
	}

	if aper.IsBitSet(optionals, 1) {
		tmp_ChildNodeCellsItem := Sequence[*ChildNodesListItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofChildIABNodes},
			ext: false,
		}
		fn := func() *ChildNodesListItem { return new(ChildNodesListItem) }
		if err = tmp_ChildNodeCellsItem.Decode(r, fn); err != nil {
			return utils.WrapError("Decode ChildNodeCellsItem", err)
		}
		ie.ChildNodeCellsItem = &ChildNodeCellsItem{}
		for _, item := range tmp_ChildNodeCellsItem.Value {
			ie.ChildNodeCellsItem.Value = append(ie.ChildNodeCellsItem.Value, item)
		}
	}

	return
}

type ChildNodeCellsItem struct {
	Value []*ChildNodesListItem
}
