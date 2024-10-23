package services

import (
	"be-swc-list/app/db"
	"be-swc-list/app/db/repository"
	"be-swc-list/app/generated/swc"
	"be-swc-list/app/model"
	"context"
	"log"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type SensorWeaponCoverageService struct {
	db *gorm.DB
}

func NewSensorWeaponCoverageService() *SensorWeaponCoverageService {

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	return &SensorWeaponCoverageService{db: db}
}

func (s *SensorWeaponCoverageService) GetSwcList(ctx context.Context, swcRequest *swc.GetSWCListReq) (*swc.GetSWCListRes, error) {
	var usersPage []model.SWCParameterResult

	pageNumber := swcRequest.PageNumber
	pageSize := swcRequest.PageSize
	environments := swcRequest.Environment

	if environments == "air" {
		if swcRequest.Group != "" {
			usersPage = repository.GetSWCParameterEnvAirResultsParams(s.db, swcRequest.Group, pageNumber, pageSize)
		} else {
			usersPage = repository.GetSWCParameterEnvAirResults(s.db, pageNumber, pageSize)
		}
	} else if environments == "non-air" {
		if swcRequest.Group != "" {
			usersPage = repository.GetSWCParameterEnvNonAirResultsParams(s.db, swcRequest.Group, pageNumber, pageSize)
		} else {
			usersPage = repository.GetSWCParameterEnvNonAirResults(s.db, pageNumber, pageSize)
		}
	} else if environments == "" {
		if swcRequest.Group != "" {
			usersPage = repository.GetSWCParameterResultsParams(s.db, swcRequest.Group, pageNumber, pageSize)
		} else {
			log.Println("EMPTY")
			usersPage = repository.GetSWCParameterResults(s.db, pageNumber, pageSize)
		}

	}

	minRange := &swc.MinRange{
		Default:     0,
		Value:       0,
		IsAvailable: false,
	}
	maxRange := &swc.MaxRange{
		Default:     0,
		Value:       0,
		IsAvailable: false,
	}

	startBearing := &swc.StartBearing{
		Default:     0,
		Value:       0,
		IsAvailable: false,
	}
	endBearing := &swc.EndBearing{
		Default:     0,
		Value:       0,
		IsAvailable: false,
	}

	legnthData := 0
	var swcListData []*swc.GetSWCListRes_SWCListData
	for _, result := range usersPage {
		groupSWC := result.GroupSWC
		environment := result.Environment
		typeValue := result.Type
		valueStr := result.Value
		defaultValueStr := result.DefaultValue
		item := result.Item
		items := strings.Split(item, ",")

		valueStrArr := strings.Split(valueStr, ",")
		defaultValueStrArr := strings.Split(defaultValueStr, ",")
		itemsArr := strings.Split(item, ",")

		var values []float64
		var defaultValues []float64

		// Convert value strings to float64
		for _, v := range valueStrArr {
			if val, err := strconv.ParseFloat(v, 64); err == nil {
				values = append(values, val)
			}
		}

		for _, dv := range defaultValueStrArr {
			if defVal, err := strconv.ParseFloat(dv, 64); err == nil {
				defaultValues = append(defaultValues, defVal)
			}
		}

		for i := range itemsArr {
			switch items[i] {
			case "Min. Range":
				minRange = &swc.MinRange{
					Default:     defaultValues[i],
					Value:       values[i],
					IsAvailable: true,
				}
			case "Max. Range":
				maxRange = &swc.MaxRange{
					Default:     defaultValues[i],
					Value:       values[i],
					IsAvailable: true,
				}
			case "Start Bearing":
				startBearing = &swc.StartBearing{
					Default:     defaultValues[i],
					Value:       values[i],
					IsAvailable: true,
				}
			case "End Bearing":
				endBearing = &swc.EndBearing{
					Default:     defaultValues[i],
					Value:       values[i],
					IsAvailable: true,
				}
			}
		}
		if minRange.IsAvailable && maxRange.IsAvailable && startBearing.IsAvailable && endBearing.IsAvailable {
			itemList := &swc.ItemList{
				MinRange:     minRange,
				MaxRange:     maxRange,
				StartBearing: startBearing,
				EndBearing:   endBearing,
			}
			swcListData = append(swcListData, &swc.GetSWCListRes_SWCListData{
				Type:        typeValue,
				GroupSWC:    groupSWC,
				ItemList:    itemList,
				Environment: environment,
			})
			legnthData++
		}
		//Reset Value
		minRange = &swc.MinRange{}
		maxRange = &swc.MaxRange{}
		startBearing = &swc.StartBearing{}
		endBearing = &swc.EndBearing{}
	}
	totalDataAll := s.GetCountAllData()

	swcResponse := &swc.GetSWCListRes{
		TotalData:    int32(legnthData),
		TotalDataAll: int32(totalDataAll),
		SwcListData:  swcListData,
	}
	return swcResponse, nil
}

func (s *SensorWeaponCoverageService) GetSwcDatabyId(ctx context.Context, swcReq *swc.GetSWCParameterByIdReq) (*model.SensorWeaponCoverage, error) {
	var swcRecord model.SensorWeaponCoverage

	if err := s.db.First(&swcRecord, swcReq.Id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("SWC with ID %d not found", swcReq.Id)
			return nil, err
		}
		log.Printf("Failed to retrieve SWC: %v", err)
		return nil, err
	}

	return &swcRecord, nil

}

func (s *SensorWeaponCoverageService) GetCountAllData() int {
	var totalData int

	datas := repository.GetSWCCount(s.db)

	for _, result := range datas {
		item := result.Item
		items := strings.Split(item, ",")

		if len(items) == 4 {
			totalData++
		}
	}
	return totalData

}
