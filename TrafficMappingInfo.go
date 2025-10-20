package ies

import "github.com/lvdund/ngap/aper"

const (
	TrafficMappingInfoPresentNothing uint64 = iota
	TrafficMappingInfoPresentIPtolayer2TrafficMappingInfo
	TrafficMappingInfoPresentBAPlayerBHRLCchannelMappingInfo
)

type TrafficMappingInfo struct {
	Choice                          uint64
	IPtolayer2TrafficMappingInfo    *IPtolayer2TrafficMappingInfo
	BAPlayerBHRLCchannelMappingInfo *BAPlayerBHRLCchannelMappingInfo
	// ChoiceExtension // ChoiceExtensions
}

func (ie *TrafficMappingInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case TrafficMappingInfoPresentIPtolayer2TrafficMappingInfo:
		err = ie.IPtolayer2TrafficMappingInfo.Encode(w)
	case TrafficMappingInfoPresentBAPlayerBHRLCchannelMappingInfo:
		err = ie.BAPlayerBHRLCchannelMappingInfo.Encode(w)
	}
	return
}

func (ie *TrafficMappingInfo) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case TrafficMappingInfoPresentIPtolayer2TrafficMappingInfo:
		var tmp IPtolayer2TrafficMappingInfo
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.IPtolayer2TrafficMappingInfo = &tmp
	case TrafficMappingInfoPresentBAPlayerBHRLCchannelMappingInfo:
		var tmp BAPlayerBHRLCchannelMappingInfo
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.BAPlayerBHRLCchannelMappingInfo = &tmp
	}
	return
}
