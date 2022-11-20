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

// 次卡列表

type CardGetListInput struct {
	CardName string // 次卡名称
	Page     int    // 分页号码
	Size     int    // 分页数量，最大50
}

// ContentGetListOutput 查询列表结果
type CardGetListOutput struct {
	List  []CardGetListOutputItem `json:"list" description:"列表"`
	Page  int                     `json:"page" description:"分页码"`
	Size  int                     `json:"size" description:"分页数量"`
	Total int                     `json:"total" description:"数据总数"`
}

type CardGetListOutputItem struct {
	CardItems *CardItems `json:"card_items"`
}

type CardItems struct {
	OrgId         int    `json:"org_id"`
	CardName      string `json:"card_name"`
	SaleStartDate string `json:"sale_start_date"`
	SaleStopDate  string `json:"sale_stop_date"`
	Price         string `json:"price"`
	BeforehandNum int    `json:"beforehand_num"`
	BuyNum        int    `json:"buy_num"`
}

//次卡列表
