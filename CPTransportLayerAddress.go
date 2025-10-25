package ies

import "github.com/lvdund/ngap/aper"

const (
	CPTransportLayerAddressPresentNothing uint64 = iota
	CPTransportLayerAddressPresentEndpointIPAddress
	CPTransportLayerAddressPresentEndpointIPAddressAndPort
)

type CPTransportLayerAddress struct {
	Choice                   uint64
	EndpointIPAddress        *TransportLayerAddress
	EndpointIPAddressAndPort *EndpointIPAddressAndPort
	// ChoiceExtension
}

func (ie *CPTransportLayerAddress) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case CPTransportLayerAddressPresentEndpointIPAddress:
		err = ie.EndpointIPAddress.Encode(w)
	case CPTransportLayerAddressPresentEndpointIPAddressAndPort:
		err = ie.EndpointIPAddressAndPort.Encode(w)
	}
	return
}

func (ie *CPTransportLayerAddress) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case CPTransportLayerAddressPresentEndpointIPAddress:
		var tmp TransportLayerAddress
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EndpointIPAddress = &tmp
	case CPTransportLayerAddressPresentEndpointIPAddressAndPort:
		var tmp EndpointIPAddressAndPort
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EndpointIPAddressAndPort = &tmp
	}
	return
}
