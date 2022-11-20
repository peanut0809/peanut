package markting

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CardReq struct {
	g.Meta             `path:"/card/add" tags:"card" method:"post"`
	OrgId              int    `json:"org_id" v:"bail|required|integer|length:4,15#|主办id不能为空|主办标识必须是数值|主办id长度在{min}到{max}之间"`
	EcardNmae          string `json:"card_name" v:"bain|required|length:1,50#次卡名称不能为空|次卡名称不能为空长度在{min}到{max}之间"`
	SaleStartDate      string `json:"sale_start_date" v:"date-format:Y-m-d H:i:s#售卖开始时间数据异常"`
	SaleStopDate       string `json:"sale_stop_date" v:"date-format:Y-m-d H:i:s#售卖开始时间数据异常"`
	Colour             string `json:"colour"`
	ImgId              int    `json:"img_id"`
	IssueNum           int    `json:"issue_num" v:"bail|required|integer|length:1,10000#发行量不能为空|发行量必须是数值|发行量长度在{min}到{max}之间"`
	Price              string `json:"price" v:"float"`
	ValidStatus        int    `json:"valid_status" v:"in:1,2,3#有效期状态异常"`
	ValidUntil         string `json:"valid_until"`
	Usage              int    `json:"usage" v:"in:1,2#使用方式异常"`
	BeforehandNum      int    `json:"beforehand_num" v:"bail|required|integer|length:1,10000#预约次数必填|预约次数必须整数|预约次数异常"`
	UserInformation    string `json:"user_information"`
	CardExplain        string `json:"card_explain"`
	BuyNum             int    `json:"buy_num"`
	ActivationCodeNum  int    `json:"activation_code_num"`
	ActivationUsageNum int    `json:"activation_usage_num"`
	Stock              int    `json:"stock" v:"bail|required|integer#库存必填|库存必须整数"`
	SaleStatus         int    `json:"sale_status" v:"in:1,2#使用方式异常"`
	CheckNum           int    `json:"check_num"`
	TextColour         string `json:"text_colour" v:"length:1,100#年卡文字颜色长度异常"`
	Cancellation       int    `json:"cancellation" v:"in:0,1#是否可取消预约状态异常"`
	CancellationTime   string `json:"cancellation_time"`
	GiveMember         int    `json:"give_member" v:"in:0,1#赠送会员等级状态异常"`
	DeliveryAddress    string `json:"delivery_address"  v:"length:1,300#商品取货地址长度异常"`
	ContainGoods       int    `json:"contain_goods" v:"in:0,1#是否包含商品异常"`
	//EcardExtend        []EcardExtendItems `json:"ecard_extend" v:"required"`
}

type EcardExtendItems struct {
	EventeId     int `json:"event_id" v:"bail|required|integer|length:4,15#活动id不能为空|活动标识必须是数值|活动id长度在{min}到{max}之间"`
	ScreeningsId int `json:"screenings_id" v:"bail|required|integer|length:4,15#场次id不能为空|场次标识必须是数值|场次id长度在{min}到{max}之间"`

	PriceId         string `json:"price_id" v:"required#票品id集合不能为空"`
	ActivityNum     int    `json:"activity_num"`
	EveryNum        int    `json:"every_num"`
	EveryConsumenum int    `json:"every_consume_num"`
}

type CardCreateRes struct {
	Items bool `json:"items"`
}
