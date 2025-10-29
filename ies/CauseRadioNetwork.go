package ies

import "github.com/lvdund/ngap/aper"

const (
	CauseRadioNetworkUnspecified                            aper.Enumerated = 0
	CauseRadioNetworkRlfailurerlc                           aper.Enumerated = 1
	CauseRadioNetworkUnknownoralreadyallocatedgnbcuuef1Apid aper.Enumerated = 2
	CauseRadioNetworkUnknownoralreadyallocatedgnbduuef1Apid aper.Enumerated = 3
	CauseRadioNetworkUnknownorinconsistentpairofuef1Apid    aper.Enumerated = 4
	CauseRadioNetworkInteractionwithotherprocedure          aper.Enumerated = 5
	CauseRadioNetworkNotsupportedqcivalue                   aper.Enumerated = 6
	CauseRadioNetworkActiondesirableforradioreasons         aper.Enumerated = 7
	CauseRadioNetworkNoradioresourcesavailable              aper.Enumerated = 8
	CauseRadioNetworkProcedurecancelled                     aper.Enumerated = 9
	CauseRadioNetworkNormalrelease                          aper.Enumerated = 10
	CauseRadioNetworkCellnotavailable                       aper.Enumerated = 11
	CauseRadioNetworkRlfailureothers                        aper.Enumerated = 12
	CauseRadioNetworkUerejection                            aper.Enumerated = 13
	CauseRadioNetworkResourcesnotavailablefortheslice       aper.Enumerated = 14
	CauseRadioNetworkAmfinitiatedabnormalrelease            aper.Enumerated = 15
	CauseRadioNetworkReleaseduetopreemption                 aper.Enumerated = 16
	CauseRadioNetworkPlmnnotservedbythegnbcu                aper.Enumerated = 17
	CauseRadioNetworkMultipledrbidinstances                 aper.Enumerated = 18
	CauseRadioNetworkUnknowndrbid                           aper.Enumerated = 19
	CauseRadioNetworkMultiplebhrlcchidinstances             aper.Enumerated = 20
	CauseRadioNetworkUnknownbhrlcchid                       aper.Enumerated = 21
	CauseRadioNetworkChocpcresourcestobechanged             aper.Enumerated = 22
	CauseRadioNetworkNpnnotsupported                        aper.Enumerated = 23
	CauseRadioNetworkNpnaccessdenied                        aper.Enumerated = 24
	CauseRadioNetworkGnbcucellcapacityexceeded              aper.Enumerated = 25
	CauseRadioNetworkReportcharacteristicsempty             aper.Enumerated = 26
	CauseRadioNetworkExistingmeasurementid                  aper.Enumerated = 27
	CauseRadioNetworkMeasurementtemporarilynotavailable     aper.Enumerated = 28
	CauseRadioNetworkMeasurementnotsupportedfortheobject    aper.Enumerated = 29
)

type CauseRadioNetwork struct {
	Value aper.Enumerated
}

func (ie *CauseRadioNetwork) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 29}, true)
	return
}

func (ie *CauseRadioNetwork) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 29}, true)
	ie.Value = aper.Enumerated(v)
	return
}
