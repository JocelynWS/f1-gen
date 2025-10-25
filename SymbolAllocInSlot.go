package ies

import "github.com/lvdund/ngap/aper"

const (
	SymbolAllocInSlotPresentNothing uint64 = iota
	SymbolAllocInSlotPresentAllDL
	SymbolAllocInSlotPresentAllUL
	SymbolAllocInSlotPresentBothDLAndUL
)

type SymbolAllocInSlot struct {
	Choice      uint64
	AllDL       *NULL
	AllUL       *NULL
	BothDLAndUL *NumDLULSymbols
	// ChoiceExtension
}

func (ie *SymbolAllocInSlot) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case SymbolAllocInSlotPresentAllDL:
		err = ie.AllDL.Encode(w)
	case SymbolAllocInSlotPresentAllUL:
		err = ie.AllUL.Encode(w)
	case SymbolAllocInSlotPresentBothDLAndUL:
		err = ie.BothDLAndUL.Encode(w)
	}
	return
}

func (ie *SymbolAllocInSlot) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case SymbolAllocInSlotPresentAllDL:
		var tmp NULL
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.AllDL = &tmp
	case SymbolAllocInSlotPresentAllUL:
		var tmp NULL
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.AllUL = &tmp
	case SymbolAllocInSlotPresentBothDLAndUL:
		var tmp NumDLULSymbols
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.BothDLAndUL = &tmp
	}
	return
}
