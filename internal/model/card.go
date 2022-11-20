package model

// ContentCreateUpdateBase 创建/修改内容基类
type CardCreateUpdateBase struct {
	OrgId              int
	EcardNmae          string
	SaleStartDate      string
	SaleStopDate       string
	Colour             string
	ImgId              int
	IssueNum           int
	Price              string
	ValidStatus        int
	ValidUntil         string
	Usage              int
	BeforehandNum      int
	UserInformation    string
	CardExplain        string
	BuyNum             int
	ActivationCodeNum  int
	ActivationUsageNum int
	Stock              int
	SaleStatus         int
	CheckNum           int
	TextColour         string
	Cancellation       int
	CancellationTime   string
	GiveMember         int
	DeliveryAddress    string
	ContainGoods       int
	EcardExtend        []EcardExtendItems
}

type EcardExtendItems struct {
	EventeId        int
	ScreeningsId    int
	PriceId         string
	ActivityNum     int
	EveryNum        int
	EveryConsumenum int
}

// ContentCreateInput 创建内容
type CardCreateInput struct {
	CardCreateUpdateBase
}

// CardCreateOutput 创建内容返回结果
type CardCreateOutput struct {
	Items bool `json:"items"`
}
