package controller

import (
	"context"

	v1 "golang8090/api/markting"
	"golang8090/internal/model"
	"golang8090/internal/service"
)

// Card 内容管理
var Card = cCard{}

type cCard struct{}

func (a *cCard) Create(ctx context.Context, req *v1.CardReq) (res *v1.CardCreateRes, err error) {
	out, err := service.Card().Create(ctx, model.CardCreateInput{
		CardCreateUpdateBase: model.CardCreateUpdateBase{
			OrgId:              req.OrgId,
			EcardNmae:          req.EcardNmae,
			SaleStartDate:      req.SaleStartDate,
			SaleStopDate:       req.SaleStopDate,
			Colour:             req.Colour,
			ImgId:              req.ImgId,
			IssueNum:           req.IssueNum,
			Price:              req.Price,
			ValidStatus:        req.ValidStatus,
			ValidUntil:         req.ValidUntil,
			Usage:              req.Usage,
			BeforehandNum:      req.BeforehandNum,
			UserInformation:    req.UserInformation,
			CardExplain:        req.CardExplain,
			BuyNum:             req.BuyNum,
			ActivationCodeNum:  req.ActivationCodeNum,
			ActivationUsageNum: req.ActivationUsageNum,
			Stock:              req.Stock,
			SaleStatus:         req.SaleStatus,
			CheckNum:           req.CheckNum,
			TextColour:         req.TextColour,
			Cancellation:       req.Cancellation,
			CancellationTime:   req.CancellationTime,
			GiveMember:         req.GiveMember,
			DeliveryAddress:    req.DeliveryAddress,
			ContainGoods:       req.ContainGoods,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.CardCreateRes{out}, nil
}

func (a *cCard) Items(ctx context.Context, req *v1.CardGetListCommonReq) (res *v1.CardGetListCommonRes, err error) {
	out, err := service.Card().GetList(ctx, model.CardGetListInput{
		CardName: req.CardName,
		Page:     req.Page,
		Size:     req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CardGetListCommonRes{
		Page:  out.Page,
		Size:  out.Size,
		Total: out.Total,
		List:  "",
	}, nil
}
