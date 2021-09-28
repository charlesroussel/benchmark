package models

type BidBodyRequest struct {
	ID     string               `json:"id"`
	App    BidBodyRequestApp    `json:"app"`
	At     int                  `json:"at"`
	TMax   int                  `json:"tmax"`
	Imp    []BidBodyRequestImp  `json:"imp"`
	Regs   BidBodyRequestRegs   `json:"regs"`
	Device BidBodyRequestDevice `json:"device"`
	Site   BidBodyRequestSite   `json:"site"`
	User   BidBodyRequestUser   `json:"user"`
}

type BidBodyRequestImp struct {
	ID       string                  `json:"id"`
	BidFloor float32                 `json:"bidfloor"`
	TagID    string                  `json:"tagid"`
	Banner   BidBodyRequestImpBanner `json:"banner"`
}

type BidBodyRequestImpBanner struct {
	Pos    int                             `json:"pos"`
	Format []BidBodyRequestImpBannerFormat `json:"format"`
}

type BidBodyRequestImpBannerFormat struct {
	W int `json:"w"`
	H int `json:"h"`
}

type BidBodyRequestRegs struct {
	Ext BidBodyRequestRegsExt `json:"ext"`
}

type BidBodyRequestRegsExt struct {
	Gdpr int `json:"gdpr"`
}

type BidBodyRequestSite struct {
	ID        string                      `json:"id"`
	Domain    string                      `json:"domain"`
	Publisher BidBodyRequestSitePublisher `json:"publisher"`
}

type BidBodyRequestSitePublisher struct {
	ID string `json:"id"`
}

type BidBodyRequestUser struct {
	ID         string                `json:"id"`
	Ext        BidBodyRequestUserExt `json:"ext"`
	CustomData string                `json:"customdata"`
}

type BidBodyRequestUserExt struct {
	Consent string `json:"consent"`
}

type BidBodyRequestApp struct {
	ID     string `json:"id"`
	Bundle string `json:"bundle"`
}

type BidBodyRequestDevice struct {
	UA         string `json:"ua"`
	DeviceType string `json:"devicetype"`
	Ip         string `json:"ip"`
	Make       string `json:"make"`
	Model      string `json:"model"`
	OS         string `json:"os"`
	OSVersion  string `json:"osv"`
}

type BidResponse struct {
	ID      string               `json:"id"`
	SeatBid []BidResponseSeatBid `json:"seatbid"`
}

type BidResponseSeatBid struct {
	Bid []BidResponseSeatBidBid `json:"bid"`
}

type BidResponseSeatBidBid struct {
	ID      string                   `json:"id"`
	ImpID   string                   `json:"impid"`
	Price   float64                  `json:"price"`
	NUrl    string                   `json:"nurl"`
	Adm     string                   `json:"adm"`
	ADomain []string                 `json:"adomain"`
	W       int                      `json:"w"`
	H       int                      `json:"h"`
	Ext     BidResponseSeatBidBidExt `json:"ext"`
}

type BidResponseSeatBidBidExt struct {
	UserID          string `json:"userid"`
	UserSync        bool   `json:"usersync"`
	AdvertiserID    string `json:"advertiserid"`
	IsOmidCompliant bool   `json:"isomidcompliant"`
	AdContent       string `json:"adcontent"`
	AdvertId        string `json:"advertid"`
	CampaignId      string `json:"campaignid"`
	MediaType       string `json:"mediatype"`
	LandingPageURL  string `json:"landing_page_url"`
}
