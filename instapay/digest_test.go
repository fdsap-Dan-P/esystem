package instapay

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDigest(t *testing.T) {
	xml := []byte("<Message xmlns='urn:instapay2019' xmlns:head='urn:iso:std:iso:20022:tech:xsd:head.001.001.01' xmlns:er='urn:iso:std:iso:20022:tech:xsd:admn.005.001.01'><AppHdr><head:Fr><head:FIId><head:FinInstnId><head:BICFI>CAMZPHM2XXX</head:BICFI></head:FinInstnId></head:FIId></head:Fr><head:To><head:FIId><head:FinInstnId><head:BICFI>BNNNPH22XXX</head:BICFI></head:FinInstnId></head:FIId></head:To><head:BizMsgIdr>B20230320CAMZPHM2XXXB51110787460656</head:BizMsgIdr><head:MsgDefIdr>admn.005.001.01</head:MsgDefIdr><head:CreDt>2023-03-20T01:20:werZ</head:CreDt><head:Sgntr></head:Sgntr></AppHdr><EchoRequest><er:AdmnEchoReq><er:GrpHdr><er:MsgId>M20230321CAMZPHM2XXXB00000000120</er:MsgId><er:CreDtTm>2023-03-21T01:22:57.541</er:CreDtTm></er:GrpHdr><er:EchoTxInf><er:FnctnCd>731</er:FnctnCd><er:InstrId>20230321CAMZPHM2XXXB0000000001</er:InstrId><er:InstgAgt><er:FinInstnId><er:BIC>CAMZPHM2XXX</er:BIC></er:FinInstnId></er:InstgAgt><er:InstdAgt><er:FinInstnId><er:BIC>BNNNPH22XXX</er:BIC></er:FinInstnId></er:InstdAgt></er:EchoTxInf></er:AdmnEchoReq></EchoRequest></Message>")
	val := Digest(xml)
	fmt.Printf("xml: %x", val)
	require.Equal(t, val, "ppp")

}
