package queue

type EventType int32 

const (
	MD_ONDISCONNECTED EventType = 3000 + iota

	MD_ONERROR // 3001

	MD_ONSUBMARKETDATA // 3002

	MD_ONUNSUBMARKETDATA // 3003

	MD_ONDEPTHMARKETDATA // 3004

	MD_ONSUBORDERBOOK // 3005

	MD_ONUNSUBORDERBOOK // 3006

	MD_ONORDERBOOK // 3007

	MD_ONSUBTICKBYTICK // 3008

	MD_ONUNSUBTICKBYTICK // 3009

	MD_ONTICKBYTICK // 3010

	MD_ONSUBSCRIBEALLMARKETDATA // 3011

	MD_ONUNSUBSCRIBEALLMARKETDATA // 3012

	MD_ONSUBSCRIBEALLORDERBOOK // 3013

	MD_ONUNSUBSCRIBEALLORDERBOOK // 3014

	MD_ONSUBSCRIBEALLTICKBYTICK // 3015

	MD_ONUNSUBSCRIBEALLTICKBYTICK // 3016

	MD_ONQUERYALLTICKERS // 3017

	MD_ONQUERYTICKERSPRICEINFO // 3018

	MD_ONSUBSCRIBEALLOPTIONMARKETDATA // 3019

	MD_ONUNSUBSCRIBEALLOPTIONMARKETDATA // 3020

	MD_ONSUBSCRIBEALLOPTIONORDERBOOK // 3021

	MD_ONUNSUBSCRIBEALLOPTIONORDERBOOK // 3022

	MD_ONSUBSCRIBEALLOPTIONTICKBYTICK // 3023

	MD_ONUNSUBSCRIBEALLOPTIONTICKBYTICK // 3024

	TRADER_ONDISCONNECTED // 3025

	TRADER_ONERROR // 3026

	TRADER_ONORDEREVENT // 3027

	TRADER_ONTRADEEVENT // 3028

	TRADER_ONCANCELORDERERROR // 3029

	TRADER_ONQUERYORDER // 3030

	TRADER_ONQUERYORDERBYPAGE // 3031

	TRADER_ONQUERYTRADE // 3032

	TRADER_ONQUERYTRADEBYPAGE // 3033

	TRADER_ONQUERYPOSITION // 3034

	TRADER_ONQUERYASSET // 3035

	TRADER_ONQUERYSTRUCTUREDFUND // 3036

	TRADER_ONQUERYFUNDTRANSFER // 3037

	TRADER_ONFUNDTRANSFER // 3038

	TRADER_ONQUERYETF // 3039

	TRADER_ONQUERYETFBASKET // 3040

	TRADER_ONQUERYIPOINFOLIST // 3041

	TRADER_ONQUERYIPOQUOTAINFO // 3042

	TRADER_ONQUERYOPTIONAUCTIONINFO // 3043

)

const (
	MD_RELEASE EventType = 4000 + iota

	MD_SETUDPBUFFERSIZE // 4001

	MD_REGISTERSPI // 4002

	MD_SETHEARTBEATINTERVAL // 4003

	MD_SUBSCRIBEMARKETDATA // 4004

	MD_UNSUBSCRIBEMARKETDATA // 4005

	MD_SUBSCRIBEORDERBOOK // 4006

	MD_UNSUBSCRIBEORDERBOOK // 4007

	MD_SUBSCRIBETICKBYTICK // 4008

	MD_UNSUBSCRIBETICKBYTICK // 4009

	MD_SUBSCRIBEALLMARKETDATA // 4010

	MD_UNSUBSCRIBEALLMARKETDATA // 4011

	MD_SUBSCRIBEALLORDERBOOK // 4012

	MD_UNSUBSCRIBEALLORDERBOOK // 4013

	MD_SUBSCRIBEALLTICKBYTICK // 4014

	MD_UNSUBSCRIBEALLTICKBYTICK // 4015

	MD_LOGIN // 4016

	MD_LOGOUT // 4017

	MD_QUERYALLTICKERS // 4018

	MD_QUERYTICKERSPRICEINFO // 4019

	MD_QUERYALLTICKERSPRICEINFO // 4020

	MD_SUBSCRIBEALLOPTIONMARKETDATA // 4021

	MD_UNSUBSCRIBEALLOPTIONMARKETDATA // 4022

	MD_SUBSCRIBEALLOPTIONORDERBOOK // 4023

	MD_UNSUBSCRIBEALLOPTIONORDERBOOK // 4024

	MD_SUBSCRIBEALLOPTIONTICKBYTICK // 4025

	MD_UNSUBSCRIBEALLOPTIONTICKBYTICK // 4026

	TRADER_RELEASE // 4027

	TRADER_REGISTERSPI // 4028

	TRADER_GETCLIENTIDBYXTPID // 4029

	TRADER_SUBSCRIBEPUBLICTOPIC // 4030

	TRADER_SETSOFTWAREVERSION // 4031

	TRADER_SETSOFTWAREKEY // 4032

	TRADER_SETHEARTBEATINTERVAL // 4033

	TRADER_LOGIN // 4034

	TRADER_LOGOUT // 4035

	TRADER_ISSERVERRESTART // 4036

	TRADER_INSERTORDER // 4037

	TRADER_CANCELORDER // 4038

	TRADER_QUERYORDERBYXTPID // 4039

	TRADER_QUERYORDERS // 4040

	TRADER_QUERYORDERSBYPAGE // 4041

	TRADER_QUERYTRADESBYXTPID // 4042

	TRADER_QUERYTRADES // 4043

	TRADER_QUERYTRADESBYPAGE // 4044

	TRADER_QUERYPOSITION // 4045

	TRADER_QUERYASSET // 4046

	TRADER_QUERYSTRUCTUREDFUND // 4047

	TRADER_FUNDTRANSFER // 4048

	TRADER_QUERYFUNDTRANSFER // 4049

	TRADER_QUERYETF // 4050

	TRADER_QUERYETFTICKERBASKET // 4051

	TRADER_QUERYIPOINFOLIST // 4052

	TRADER_QUERYIPOQUOTAINFO // 4053

	TRADER_QUERYOPTIONAUCTIONINFO // 4054

)