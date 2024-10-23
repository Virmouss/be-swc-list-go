package router

import (
	"context"

	"be-swc-list/app/generated/swc"
	"be-swc-list/app/model"
)

type SwcService interface {
	GetSwcList(context.Context, *swc.GetSWCListReq) (*swc.GetSWCListRes, error)
	GetSwcDatabyId(context.Context, *swc.GetSWCParameterByIdReq) (*model.SensorWeaponCoverage, error)
}
