package Card

import (
	"context"
	"fmt"

	"golang8090/internal/dao"
	"golang8090/internal/model/entity"
	"golang8090/internal/service"

	"golang8090/internal/model"

	"github.com/gogf/gf/v2/database/gdb"
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

	g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		CardId, err := dao.EventeECard.Ctx(ctx).InsertAndGetId(g.Map{
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
		if err != nil { // 插入 card主表
			return err
		}
		CardExtendItems := in.EcardExtend
		for _, items := range CardExtendItems {
			result, err := dao.EventeECardExtend.Ctx(ctx).Insert(g.Map{
				"card_id":           CardId,
				"org_id":            in.OrgId,
				"event_id":          items.EventeId,
				"screenings_id":     items.ScreeningsId,
				"price_id":          items.PriceId,
				"activity_num":      items.ActivityNum,
				"every_num":         items.EveryNum,
				"every_consume_num": items.EveryConsumenum,
			}) //插入 card_extend子表
			result.RowsAffected()
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// GetList 查询内容列表
func (s *sCard) GetList(ctx context.Context, in model.CardGetListInput) (out *model.CardGetListOutput, err error) {
	var (
		m = dao.EventeECard.Ctx(ctx)
	)
	out = &model.CardGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	fmt.Println(listModel)
	// 执行查询
	var list []*entity.EventeECard
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Card
	if err := listModel.ScanList(&out.List, "Card"); err != nil {
		return out, err
	}
	return
}
