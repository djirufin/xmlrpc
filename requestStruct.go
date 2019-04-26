package xmlrpc

import "time"

//ReqGetBalance : Request to getBalance struct
type ReqGetBalance struct {
	OriginNodeType      string    `xml:"originNodeType"`
	OriginHostName      string    `xml:"originHostName"`
	OriginTransactionID string    `xml:"originTransactionID"`
	OriginTimeStamp     time.Time `xml:"originTimeStamp"`
	SubscriberNumber    string    `xml:"subscriberNumber"`
	SubscriberNumberNAI int       `xml:"subscriberNumberNAI"`
}

//ReqGetOffers : Request to getBalance struct
type ReqGetOffers struct {
	OriginNodeType         string    `xml:"originNodeType"`
	OriginHostName         string    `xml:"originHostName"`
	OriginTransactionID    string    `xml:"originTransactionID"`
	OriginTimeStamp        time.Time `xml:"originTimeStamp"`
	SubscriberNumber       string    `xml:"subscriberNumber"`
	SubscriberNumberNAI    int       `xml:"subscriberNumberNAI"`
	OfferRequestedTypeFlag string    `xml:"offerRequestedTypeFlag"`
}

//RepGetBalance : Response from getBalance struct
type RepGetBalance struct {
	AccountFlagsAfter           AccountFlagsAfterStruct
	AccountFlagsBefore          AccountFlagsAfterStruct
	AccountValue1               string
	CreditClearanceDate         time.Time
	Currency1                   string
	DedicatedAccountInformation []DaInfo
	LanguageIDCurrent           int
	OfferInformationList        OfferInfo
	OriginTransactionID         string
	ResponseCode                int
	ServiceClassCurrent         int
	ServiceFeeExpiryDate        time.Time
	ServiceRemovalDate          time.Time
	SupervisionExpiryDate       time.Time
}

//AccountFlagsAfterStruct : details
type AccountFlagsAfterStruct struct {
	ActivationStatusFlag               bool
	NegativeBarringStatusFlag          bool
	ServiceFeePeriodExpiryFlag         bool
	ServiceFeePeriodWarningActiveFlag  bool
	SupervisionPeriodExpiryFlag        bool
	SupervisionPeriodWarningActiveFlag bool
}

//DaInfo :DedicatedAccountInformation details struct
type DaInfo struct {
	DedicatedAccountActiveValue1 string
	DedicatedAccountID           int
	DedicatedAccountUnitType     int
	DedicatedAccountValue1       string
	ExpiryDate                   time.Time
	StartDate                    time.Time
}

//OfferInfo : Offer details struct
type OfferInfo struct {
	OfferID    int
	OfferType  int
	ExpiryDate time.Time
	StartDate  time.Time
}
