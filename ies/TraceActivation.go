package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TraceActivation struct {
	TraceID                        []byte         `lb:8,ub:8,mandatory`
	InterfacesToTrace              aper.BitString `lb:8,ub:8,mandatory`
	TraceDepth                     TraceDepth     `mandatory`
	TraceCollectionEntityIPAddress aper.BitString `lb:1,ub:160,mandatory,valueExt`
	// IEExtensions * `optional`
}

func (ie *TraceActivation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_TraceID := NewOCTETSTRING(ie.TraceID, aper.Constraint{Lb: 8, Ub: 8}, false)
	if err = tmp_TraceID.Encode(w); err != nil {
		err = utils.WrapError("Encode TraceID", err)
		return
	}
	tmp_InterfacesToTrace := NewBITSTRING(ie.InterfacesToTrace, aper.Constraint{Lb: 8, Ub: 8}, false)
	if err = tmp_InterfacesToTrace.Encode(w); err != nil {
		err = utils.WrapError("Encode InterfacesToTrace", err)
		return
	}
	if err = ie.TraceDepth.Encode(w); err != nil {
		err = utils.WrapError("Encode TraceDepth", err)
		return
	}
	tmp_TraceCollectionEntityIPAddress := NewBITSTRING(ie.TraceCollectionEntityIPAddress, aper.Constraint{Lb: 1, Ub: 160}, true)
	if err = tmp_TraceCollectionEntityIPAddress.Encode(w); err != nil {
		err = utils.WrapError("Encode TraceCollectionEntityIPAddress", err)
		return
	}
	return
}
func (ie *TraceActivation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_TraceID := OCTETSTRING{
		c:   aper.Constraint{Lb: 8, Ub: 8},
		ext: false,
	}
	if err = tmp_TraceID.Decode(r); err != nil {
		err = utils.WrapError("Read TraceID", err)
		return
	}
	ie.TraceID = tmp_TraceID.Value
	tmp_InterfacesToTrace := BITSTRING{
		c:   aper.Constraint{Lb: 8, Ub: 8},
		ext: false,
	}
	if err = tmp_InterfacesToTrace.Decode(r); err != nil {
		err = utils.WrapError("Read InterfacesToTrace", err)
		return
	}
	ie.InterfacesToTrace = aper.BitString{Bytes: tmp_InterfacesToTrace.Value.Bytes, NumBits: tmp_InterfacesToTrace.Value.NumBits}
	if err = ie.TraceDepth.Decode(r); err != nil {
		err = utils.WrapError("Read TraceDepth", err)
		return
	}
	tmp_TraceCollectionEntityIPAddress := BITSTRING{
		c:   aper.Constraint{Lb: 1, Ub: 160},
		ext: true,
	}
	if err = tmp_TraceCollectionEntityIPAddress.Decode(r); err != nil {
		err = utils.WrapError("Read TraceCollectionEntityIPAddress", err)
		return
	}
	ie.TraceCollectionEntityIPAddress = aper.BitString{Bytes: tmp_TraceCollectionEntityIPAddress.Value.Bytes, NumBits: tmp_TraceCollectionEntityIPAddress.Value.NumBits}
	return
}
