package ies

import "github.com/lvdund/ngap/aper"

const (
	EUTRASpecialSubframePatternsSsp0  aper.Enumerated = 0
	EUTRASpecialSubframePatternsSsp1  aper.Enumerated = 1
	EUTRASpecialSubframePatternsSsp2  aper.Enumerated = 2
	EUTRASpecialSubframePatternsSsp3  aper.Enumerated = 3
	EUTRASpecialSubframePatternsSsp4  aper.Enumerated = 4
	EUTRASpecialSubframePatternsSsp5  aper.Enumerated = 5
	EUTRASpecialSubframePatternsSsp6  aper.Enumerated = 6
	EUTRASpecialSubframePatternsSsp7  aper.Enumerated = 7
	EUTRASpecialSubframePatternsSsp8  aper.Enumerated = 8
	EUTRASpecialSubframePatternsSsp9  aper.Enumerated = 9
	EUTRASpecialSubframePatternsSsp10 aper.Enumerated = 10
)

type EUTRASpecialSubframePatterns struct {
	Value aper.Enumerated
}

func (ie *EUTRASpecialSubframePatterns) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 10}, true)
	return
}

func (ie *EUTRASpecialSubframePatterns) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 10}, true)
	ie.Value = aper.Enumerated(v)
	return
}
