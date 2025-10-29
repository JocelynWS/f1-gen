package f1ap

const (
	PrivateIEIDPresentNothing uint64 = iota
	PrivateIEIDPresentLocal
	PrivateIEIDPresentGlobal
)

// type PrivateIEID struct {
// 	Choice uint64
// 	Local  *int64
// 	Global *OBJECT
// }

// func (ie *PrivateIEID) Encode(w *aper.AperWriter) (err error) {
// 	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
// 		return
// 	}
// 	switch ie.Choice {
// 	case PrivateIEIDPresentLocal:
// 		tmp := NewINTEGER(*ie.Local, aper.Constraint{Lb: 0, Ub: 65535}, false)
// 		err = tmp.Encode(w)
// 	case PrivateIEIDPresentGlobal:
// 		err = ie.Global.Encode(w)
// 	}
// 	return
// }

// func (ie *PrivateIEID) Decode(r *aper.AperReader) (err error) {
// 	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
// 		return
// 	}
// 	switch ie.Choice {
// 	case PrivateIEIDPresentLocal:
// 		tmp := NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 65535}, false)
// 		if err = tmp.Decode(r); err != nil {
// 			return
// 		}
// 		ie.Local = (*int64)(&tmp.Value)
// 	case PrivateIEIDPresentGlobal:
// 		var tmp OBJECT
// 		if err = tmp.Decode(r); err != nil {
// 			return
// 		}
// 		ie.Global = &tmp
// 	}
// 	return
// }
