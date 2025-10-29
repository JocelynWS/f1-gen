package f1ap

import (
	"github.com/lvdund/ngap/aper"
)

type EndpointIPAddressAndPort struct {
	EndpointIPAddress TransportLayerAddress `mandatory,ignore`
	//IEExtensions      *ProtocolExtensionContainer `optional,ignore`
}

func (e *EndpointIPAddressAndPort) Encode(w *aper.AperWriter) error {
	return e.EndpointIPAddress.Encode(w)
}

func (e *EndpointIPAddressAndPort) Decode(r *aper.AperReader) error {
	return e.EndpointIPAddress.Decode(r)
}
