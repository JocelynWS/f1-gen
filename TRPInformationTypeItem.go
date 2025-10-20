package ies

import "github.com/lvdund/ngap/aper"

const (
	TRPInformationTypeItemNrpci                          aper.Enumerated = 0
	TRPInformationTypeItemNgrancgi                       aper.Enumerated = 1
	TRPInformationTypeItemArfcn                          aper.Enumerated = 2
	TRPInformationTypeItemPrsconfig                      aper.Enumerated = 3
	TRPInformationTypeItemSsbconfig                      aper.Enumerated = 4
	TRPInformationTypeItemSfninittime                    aper.Enumerated = 5
	TRPInformationTypeItemSpatialdirectinfo              aper.Enumerated = 6
	TRPInformationTypeItemGeocoord                       aper.Enumerated = 7
	TRPInformationTypeItemTrpinformationtyperesponselist aper.Enumerated = 8
	TRPInformationTypeItemTrpinformationtyperesponseitem aper.Enumerated = 9
	TRPInformationTypeItemPcinr                          aper.Enumerated = 10
	TRPInformationTypeItemNgrancgi2                      aper.Enumerated = 11
	TRPInformationTypeItemNrarfcn                        aper.Enumerated = 12
	TRPInformationTypeItemPrsconfiguration               aper.Enumerated = 13
	TRPInformationTypeItemSsbinformation                 aper.Enumerated = 14
	TRPInformationTypeItemSfninitialisationtime          aper.Enumerated = 15
	TRPInformationTypeItemSpatialdirectioninformation    aper.Enumerated = 16
	TRPInformationTypeItemGeographicalcoordinates        aper.Enumerated = 17
	TRPInformationTypeItemChoiceextension                aper.Enumerated = 18
)

type TRPInformationTypeItem struct {
	Value aper.Enumerated
}

func (ie *TRPInformationTypeItem) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 18}, true)
	return
}

func (ie *TRPInformationTypeItem) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 18}, true)
	ie.Value = aper.Enumerated(v)
	return
}
