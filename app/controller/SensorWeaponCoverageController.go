package controller

import (
	"be-swc-list/app/generated/swc"
	"be-swc-list/app/router"
	"context"

	"google.golang.org/grpc"
)

type SensorWeaponCoverageController struct {
	swcService router.SwcService
	swc.UnimplementedSWCListServer
}

func NewGrpcSWCService(grpc *grpc.Server, swcService router.SwcService) {

	gRPCHandler := &SensorWeaponCoverageController{
		swcService: swcService,
	}

	swc.RegisterSWCListServer(grpc, gRPCHandler)
}

func (h *SensorWeaponCoverageController) GetSWCDatabyID(ctx context.Context, req *swc.GetSWCParameterByIdReq) (*swc.GetSWCParameterByIdRes, error) {
	swcId := &swc.GetSWCParameterByIdReq{Id: int32(req.Id)}

	swcData, err := h.swcService.GetSwcDatabyId(ctx, swcId)
	if err != nil {
		return nil, err
	}

	swcResponses := &swc.GetSWCParameterByIdRes{
		Id:          swcData.ID,
		Type:        swcData.Type,
		GroupSWC:    swcData.GroupSWC,
		Item:        swcData.Item,
		Environment: swcData.Environment,
		Value:       swcData.Value,
		Default:     swcData.Value,
		Unit:        swcData.Unit,
		UpdatedAt:   swcData.UpdatedAt,
		CreatedAt:   swcData.CreatedAt,
	}

	return swcResponses, nil

}

func (h *SensorWeaponCoverageController) GetSWCList(ctx context.Context, req *swc.GetSWCListReq) (*swc.GetSWCListRes, error) {
	swcListRequest := swc.GetSWCListReq{
		PageNumber:  req.PageNumber,
		PageSize:    req.PageSize,
		Environment: req.Environment,
		Group:       req.Group,
	}

	swcList, err := h.swcService.GetSwcList(ctx, &swcListRequest)
	if err != nil {
		return nil, err
	}

	return swcList, err

}
