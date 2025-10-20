package ies

import (
	"github.com/lvdund/ngap/aper"
)

type ENUMERATED struct {
	Value aper.Enumerated
	c     aper.Constraint
	ext   bool
}

func NewENUMERATED(v int64, c aper.Constraint, ext bool) ENUMERATED {
	return ENUMERATED{
		Value: aper.Enumerated(v),
		c:     c,
		ext:   ext,
	}
}
func (t *ENUMERATED) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteEnumerate(uint64(t.Value), t.c, t.ext)
	return
}
func (t *ENUMERATED) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadEnumerate(t.c, t.ext)
	t.Value = aper.Enumerated(v)
	return
}

type BITSTRING struct {
	Value aper.BitString
	c     aper.Constraint
	ext   bool
}

func NewBITSTRING(v aper.BitString, c aper.Constraint, ext bool) BITSTRING {
	return BITSTRING{
		Value: aper.BitString{
			Bytes:   v.Bytes,
			NumBits: v.NumBits,
		},
		c:   c,
		ext: ext,
	}
}
func (t *BITSTRING) Encode(w *aper.AperWriter) (err error) {
	if t.c.Lb == t.c.Ub {
		t.Value.NumBits = uint64(t.c.Lb)
	} else if len(t.Value.Bytes)*8 < int(t.c.Lb) {
		t.Value.NumBits = uint64(t.c.Lb)
	}
	err = w.WriteBitString(t.Value.Bytes, uint(t.Value.NumBits), &t.c, t.ext)
	return
}
func (t *BITSTRING) Decode(r *aper.AperReader) (err error) {
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&t.c, t.ext); err != nil {
		return
	}
	t.Value.Bytes = v
	t.Value.NumBits = uint64(n)
	return
}

type OCTETSTRING struct {
	Value aper.OctetString
	c     aper.Constraint
	ext   bool
}

func NewOCTETSTRING(v []byte, c aper.Constraint, ext bool) OCTETSTRING {
	return OCTETSTRING{
		Value: v,
		c:     c,
		ext:   ext,
	}
}
func (t *OCTETSTRING) Encode(w *aper.AperWriter) (err error) {
	if t.c.Lb == t.c.Ub && t.c.Lb == 0 {
		err = w.WriteOctetString(t.Value, nil, t.ext)
	} else {
		err = w.WriteOctetString(t.Value, &t.c, t.ext)
	}
	return
}
func (t *OCTETSTRING) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if t.c.Lb == t.c.Ub && t.c.Lb == 0 {
		if v, err = r.ReadOctetString(nil, t.ext); err != nil {
			return
		}
	} else {
		if v, err = r.ReadOctetString(&t.c, t.ext); err != nil {
			return
		}
	}

	t.Value = v
	return
}

type INTEGER struct {
	Value aper.Integer
	c     aper.Constraint
	ext   bool
}

func NewINTEGER(v int64, c aper.Constraint, ext bool) INTEGER {
	return INTEGER{
		Value: aper.Integer(v),
		c:     c,
		ext:   ext,
	}
}
func (t *INTEGER) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteInteger(int64(t.Value), &t.c, t.ext)
	return
}
func (t *INTEGER) Decode(r *aper.AperReader) (err error) {
	v, err := r.ReadInteger(&t.c, t.ext)
	t.Value = aper.Integer(v)
	return
}

type Sequence[T aper.IE] struct {
	Value []T
	c     aper.Constraint
	ext   bool
}

func NewSequence[T aper.IE](items []T, c aper.Constraint, ext bool) Sequence[T] {
	return Sequence[T]{
		Value: items,
		c:     c,
		ext:   ext,
	}
}

func (s *Sequence[T]) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[T](s.Value, w, &s.c, s.ext); err != nil {
		return
	}
	return
}
func (s *Sequence[T]) Decode(r *aper.AperReader, fn func() T) (err error) {
	var newItems []T
	newItems, err = aper.ReadSequenceOfEx(fn, r, &s.c, s.ext)
	if err != nil {
		return
	}
	s.Value = []T{}
	s.Value = append(s.Value, newItems...)
	return
}

type NULL struct{}

func NewNULL() NULL {
	return NULL{}
}

func (n *NULL) Encode(w *aper.AperWriter) (err error) {
	return
}

func (n *NULL) Decode(r *aper.AperReader) (err error) {
	return
}


// temparory
type TAC struct {
	Value []byte
}

func (ie *TAC) Encode(w *aper.AperWriter) (err error) {
	return
}
func (ie *TAC) Decode(r *aper.AperReader) (err error) {
	return
}

type PLMNIdentity struct {
	Value []byte
}

func (ie *PLMNIdentity) Encode(w *aper.AperWriter) (err error) {
	return
}
func (ie *PLMNIdentity) Decode(r *aper.AperReader) (err error) {
	return
}

type EmergencyAreaID struct {
	Value []byte
}

func (ie *EmergencyAreaID) Encode(w *aper.AperWriter) (err error) {
	return
}
func (ie *EmergencyAreaID) Decode(r *aper.AperReader) (err error) {
	return
}

type TransportLayerAddress struct {
	Value []byte
}

func (ie *TransportLayerAddress) Encode(w *aper.AperWriter) (err error) {
	return
}
func (ie *TransportLayerAddress) Decode(r *aper.AperReader) (err error) {
	return
}

const (
	maxNRARFCN                            int64 = 3279165
	maxnoofErrors                         int64 = 256
	maxnoofIndividualF1ConnectionsToReset int64 = 65536
	maxCellingNBDU                        int64 = 512
	maxnoofSCells                         int64 = 32
	maxnoofSRBs                           int64 = 8
	maxnoofDRBs                           int64 = 64
	maxnoofULUPTNLInformation             int64 = 2
	maxnoofDLUPTNLInformation             int64 = 2
	maxnoofBPLMNs                         int64 = 6
	maxnoofCandidateSpCells               int64 = 64
	maxnoofPotentialSpCells               int64 = 64
	maxnoofNrCellBands                    int64 = 32
	maxnoofSIBTypes                       int64 = 32
	maxnoofSITypes                        int64 = 32
	maxnoofPagingCells                    int64 = 512
	maxnoofTNLAssociations                int64 = 32
	maxnoofQoSFlows                       int64 = 64
	maxnoofSliceItems                     int64 = 1024
	maxCellineNB                          int64 = 256
	maxnoofExtendedBPLMNs                 int64 = 6
	maxnoofUEIDs                          int64 = 65536
	maxnoofBPLMNsNR                       int64 = 12
	maxnoofUACPLMNs                       int64 = 12
	maxnoofUACperPLMN                     int64 = 64
	maxnoofAdditionalSIBs                 int64 = 63
	maxnoofslots                          int64 = 5120
	maxnoofTLAs                           int64 = 16
	maxnoofGTPTLAs                        int64 = 16
	maxnoofBHRLCChannels                  int64 = 65536
	maxnoofRoutingEntries                 int64 = 1024
	maxnoofIABSTCInfo                     int64 = 45
	maxnoofSymbols                        int64 = 14
	maxnoofServingCells                   int64 = 32
	maxnoofDUFSlots                       int64 = 320
	maxnoofHSNASlots                      int64 = 5120
	maxnoofServedCellsIAB                 int64 = 512
	maxnoofChildIABNodes                  int64 = 1024
	maxnoofNonUPTrafficMappings           int64 = 32
	maxnoofTLAsIAB                        int64 = 1024
	maxnoofMappingEntries                 int64 = 67108864
	maxnoofDSInfo                         int64 = 64
	maxnoofEgressLinks                    int64 = 2
	maxnoofULUPTNLInformationforIAB       int64 = 32678
	maxnoofUPTNLAddresses                 int64 = 8
	maxnoofSLDRBs                         int64 = 512
	maxnoofQoSParaSets                    int64 = 8
	maxnoofPC5QoSFlows                    int64 = 2048
	maxnoofSSBAreas                       int64 = 64
	maxnoofPhysicalResourceBlocks         int64 = 275
	maxnoofPhysicalResourceBlocksMinusOne int64 = 274
	maxnoofPRACHconfigs                   int64 = 16
	maxnoofRACHReports                    int64 = 64
	maxnoofRLFReports                     int64 = 64
	maxnoofAdditionalPDCPDuplicationTNL   int64 = 2
	maxnoofRLCDuplicationState            int64 = 3
	maxnoofCHOcells                       int64 = 8
	maxnoofMDTPLMNs                       int64 = 16
	maxnoofCAGsupported                   int64 = 12
	maxnoofNIDsupported                   int64 = 12
	maxnoofNRSCSs                         int64 = 5
	maxnoofExtSliceItems                  int64 = 65535
	maxnoofPosMeas                        int64 = 16384
	maxnoofTRPInfoTypes                   int64 = 64
	maxnoofTRPs                           int64 = 65535
	maxnoofSRSTriggerStates               int64 = 3
	maxnoofSpatialRelations               int64 = 64
	maxnoBcastCell                        int64 = 16384
	maxnoofAngleInfo                      int64 = 65535
	maxnoofLCSGCSTranslation              int64 = 3
	maxnoofPath                           int64 = 2
	maxnoofMeasECID                       int64 = 64
	maxnoofSSBs                           int64 = 255
	maxnoSRSResourceSets                  int64 = 16
	maxnoSRSResourcePerSet                int64 = 16
	maxnoSRSCarriers                      int64 = 32
	maxnoSCSs                             int64 = 5
	maxnoSRSResources                     int64 = 64
	maxnoSRSPosResources                  int64 = 64
	maxnoSRSPosResourceSets               int64 = 16
	maxnoSRSPosResourcePerSet             int64 = 16
	maxnoofPRSResourceSets                int64 = 2
	maxnoofPRSResourcesPerSet             int64 = 64
	maxNoOfMeasTRPs                       int64 = 64
	maxnoofPRSresourceSets                int64 = 8
	maxnoofPRSresources                   int64 = 64
)
