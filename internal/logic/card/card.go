package Card

import (
	"context"

	"golang8090/internal/dao"
	"golang8090/internal/service"

	"golang8090/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type sCard struct{}

func init() {
	service.RegisterCard(New())
}

func New() *sCard {
	return &sCard{}
}

// Create 创建内容
func (s *sCard) Create(ctx context.Context, in model.CardCreateInput) (out bool, err error) {

	id, err := dao.EventeECard.Ctx(ctx).InsertAndGetId(g.Map{
		"org_id":               in.OrgId,
		"card_name":            in.EcardNmae,
		"sale_start_date":      in.SaleStartDate,
		"sale_stop_date":       in.SaleStopDate,
		"colour":               in.Colour,
		"img_id":               in.ImgId,
		"issue_num":            in.IssueNum,
		"price":                in.Price,
		"valid_status":         in.ValidStatus,
		"valid_until":          in.ValidUntil,
		"usage":                in.Usage,
		"beforehand_num":       in.BeforehandNum,
		"user_information":     in.UserInformation,
		"card_explain":         in.CardExplain,
		"buy_num":              in.BuyNum,
		"activation_code_num":  in.ActivationCodeNum,
		"activation_usage_num": in.ActivationUsageNum,
		"stock":                in.Stock,
		"sale_status":          in.SaleStatus,
		"check_num":            in.CheckNum,
		"text_colour":          in.TextColour,
		"cancellation":         in.Cancellation,
		"give_member":          in.GiveMember,
		"delivery_address":     in.DeliveryAddress,
		"contain_goods":        in.ContainGoods,
	})
	if id > 0 {
		return true, nil
	} else {
		return false, err
	}
}
