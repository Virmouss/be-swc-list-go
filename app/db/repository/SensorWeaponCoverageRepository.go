package repository

import (
	"be-swc-list/app/db"
	"be-swc-list/app/model"
	"log"

	"gorm.io/gorm"
)

type SensorWeaponCoverageRepo struct {
	db *gorm.DB
}

func NewSWCRepo() *SensorWeaponCoverageRepo {

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	return &SensorWeaponCoverageRepo{db: db}
}

func FindAllByGroup(db *gorm.DB, page, pageSize int32) []model.SWCParameterResult {
	var results []model.SWCParameterResult

	//Calculate offset for Pagination
	offset := (page - 1) * pageSize

	err := db.Where("is_deleted = ?", false).
		Order("group_swc").
		Limit(int(pageSize)).
		Offset(int(offset)).
		Find(&results).Error

	if err != nil {
		return nil
	}

	return results
}

func FindAllByParams(db *gorm.DB, page, pageSize int32) []model.SWCParameterResult {
	var results []model.SWCParameterResult

	offset := (page - 1) * pageSize

	err := db.Where("is_deleted = ?", false).
		Order("group_swc").
		Limit(int(pageSize)).
		Offset(int(offset)).
		Find(&results).Error

	if err != nil {
		return nil
	}

	return results
}

func FindGroup(db *gorm.DB) []model.SWCParameterResult {
	var results []model.SWCParameterResult

	err := db.Raw("SELECT group_swc FROM sensor_weapon_coverage WHERE is_deleted = false GROUP BY group_swc, environment").Scan(&results).Error

	if err != nil {
		return nil
	}

	return results
}

func GetSWCParameterResults(db *gorm.DB, page, pageSize int32) []model.SWCParameterResult {
	var results []model.SWCParameterResult

	offset := (page - 1) * pageSize

	query := `
		SELECT  
			group_swc,  
			environment,  
			sp.type, 
			STRING_AGG(CAST(id AS TEXT), ',') AS id,  
			STRING_AGG(CAST(value AS TEXT), ',') AS value,  
			STRING_AGG(CAST(default_value AS TEXT), ',') AS default_value, 
			STRING_AGG(CAST(item AS TEXT), ',') AS item, 
			STRING_AGG(CAST(created_at AS TEXT), ',') AS created_at, 
			STRING_AGG(CAST(updated_at AS TEXT), ',') AS updated_at, 
			STRING_AGG(CAST(is_deleted AS TEXT), ',') AS is_deleted, 
			STRING_AGG(CAST(unit AS TEXT), ',') AS unit 
		FROM  
			swc_parameter sp  
		WHERE  
			sp.is_deleted = false  
				
			AND ( 
				sp.item ILIKE '%Start Bearing%' OR 
				sp.item ILIKE '%End Bearing%' OR 
				sp.item ILIKE '%Max. Range%' OR 
				sp.item ILIKE '%Min. Range%' 
			) 
		GROUP BY  
			group_swc,  
			environment,  
			sp.type 
		HAVING  
			STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%Start Bearing%'  
			AND STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%End Bearing%' 
			AND STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%Max. Range%' 
			AND STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%Min. Range%'
		LIMIT ? OFFSET ?
	`
	result := db.Raw(query, pageSize, offset).Scan(&results)
	if result.Error != nil {
		return nil
	}

	return results

}

func GetSWCParameterResultsParams(db *gorm.DB, pattern string, page, pageSize int32) []model.SWCParameterResult {
	var results []model.SWCParameterResult

	// Calculate the offset for pagination
	offset := (page - 1) * pageSize

	query := `
        SELECT group_swc, environment, sp.type,
               STRING_AGG(CAST(id AS TEXT), ',') AS id,
               STRING_AGG(CAST(value AS TEXT), ',') AS value,
               STRING_AGG(CAST(default_value AS TEXT), ',') AS default_value,
               STRING_AGG(CAST(item AS TEXT), ',') AS item,
               STRING_AGG(CAST(created_at AS TEXT), ',') AS created_at,
               STRING_AGG(CAST(updated_at AS TEXT), ',') AS updated_at,
               STRING_AGG(CAST(is_deleted AS TEXT), ',') AS is_deleted,
               STRING_AGG(CAST(unit AS TEXT), ',') AS unit
        FROM swc_parameter sp
        WHERE sp.is_deleted = false
          AND LOWER(group_swc) LIKE LOWER(CONCAT('%', CAST(? AS TEXT) , '%'))
          AND (
              sp.item ILIKE '%Start Bearing%' OR
              sp.item ILIKE '%End Bearing%' OR
              sp.item ILIKE '%Max. Range%' OR
              sp.item ILIKE '%Min. Range%'
          )
        GROUP BY group_swc, environment, sp.type
        HAVING
          STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%Start Bearing%'
          AND STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%End Bearing%'
          AND STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%Max. Range%'
          AND STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%Min. Range%'
        LIMIT ? OFFSET ?
    `
	result := db.Raw(query, pattern, pageSize, offset).Scan(&results)
	if result.Error != nil {
		return nil
	}

	return results
}

func GetSWCParameterEnvAirResultsParams(db *gorm.DB, pattern string, page, pageSize int32) []model.SWCParameterResult {
	var results []model.SWCParameterResult

	offset := (page - 1) * pageSize

	query := `
		SELECT group_swc , environment, sp.type, 
               STRING_AGG(CAST(id AS TEXT), ',') AS id,  
                   STRING_AGG(CAST(value AS TEXT), ',') AS value,  
                   STRING_AGG(CAST(default_value AS TEXT), ',') AS default_value, 
                   STRING_AGG(CAST(item AS TEXT), ',') AS item, 
                   STRING_AGG(CAST(created_at AS TEXT), ',') AS created_at, 
                   STRING_AGG(CAST(updated_at AS TEXT), ',') AS updated_at, 
                   STRING_AGG(CAST(is_deleted AS TEXT), ',') AS is_deleted, 
                   STRING_AGG(CAST(unit AS TEXT), ',') AS unit 
            FROM swc_parameter sp WHERE sp.is_deleted = false AND sp.environment = 'air' and lower(group_swc) LIKE lower (concat('%', CAST(? AS TEXT), '%')) 
                AND ( 
                    sp.item ILIKE '%Start Bearing%' OR 
                    sp.item ILIKE '%End Bearing%' OR 
                    sp.item ILIKE '%Max. Range%' OR 
                    sp.item ILIKE '%Min. Range%' 
                ) 
            GROUP BY group_swc, environment, sp.type
            HAVING  
                STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%Start Bearing%'  
                AND STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%End Bearing%' 
                AND STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%Max. Range%' 
                AND STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%Min. Range%'
			LIMIT ? OFFSET ?
	`
	result := db.Raw(query, pattern, pageSize, offset).Scan(&results)
	if result.Error != nil {
		return nil
	}

	return results

}

func GetSWCParameterEnvAirResults(db *gorm.DB, page, pageSize int32) []model.SWCParameterResult {
	var results []model.SWCParameterResult

	offset := (page - 1) * pageSize

	query := `
		SELECT group_swc , environment, sp.type, 
               STRING_AGG(CAST(id AS TEXT), ',') AS id,  
                   STRING_AGG(CAST(value AS TEXT), ',') AS value,  
                   STRING_AGG(CAST(default_value AS TEXT), ',') AS default_value, 
                   STRING_AGG(CAST(item AS TEXT), ',') AS item, 
                   STRING_AGG(CAST(created_at AS TEXT), ',') AS created_at, 
                   STRING_AGG(CAST(updated_at AS TEXT), ',') AS updated_at, 
                   STRING_AGG(CAST(is_deleted AS TEXT), ',') AS is_deleted, 
                   STRING_AGG(CAST(unit AS TEXT), ',') AS unit 
            FROM swc_parameter sp WHERE sp.is_deleted = false AND sp.environment = 'air'  
                AND ( 
                    sp.item ILIKE '%Start Bearing%' OR 
                    sp.item ILIKE '%End Bearing%' OR 
                    sp.item ILIKE '%Max. Range%' OR 
                    sp.item ILIKE '%Min. Range%' 
                ) 
            GROUP BY group_swc, environment, sp.type
            HAVING  
                STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%Start Bearing%'  
                AND STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%End Bearing%' 
                AND STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%Max. Range%' 
                AND STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%Min. Range%'
			LIMIT ? OFFSET ?
		`

	result := db.Raw(query, pageSize, offset).Scan(&results)
	if result.Error != nil {
		return nil
	}

	return results

}

func GetSWCParameterEnvNonAirResultsParams(db *gorm.DB, pattern string, page, pageSize int32) []model.SWCParameterResult {
	var results []model.SWCParameterResult

	// Calculate the offset for pagination
	offset := (page - 1) * pageSize

	query := `
        SELECT group_swc, environment, sp.type,
			STRING_AGG(CAST(id AS TEXT), ',') AS id,
			STRING_AGG(CAST(value AS TEXT), ',') AS value,
			STRING_AGG(CAST(default_value AS TEXT), ',') AS default_value,
			STRING_AGG(CAST(item AS TEXT), ',') AS item,
			STRING_AGG(CAST(created_at AS TEXT), ',') AS created_at,
			STRING_AGG(CAST(updated_at AS TEXT), ',') AS updated_at,
			STRING_AGG(CAST(is_deleted AS TEXT), ',') AS is_deleted,
			STRING_AGG(CAST(unit AS TEXT), ',') AS unit
		FROM swc_parameter sp 
		WHERE sp.is_deleted = false 
		AND sp.environment = 'non-air'
		AND lower(group_swc) LIKE lower(concat('%', CAST(? AS TEXT), '%'))
		AND (
			sp.item ILIKE '%Start Bearing%' OR
			sp.item ILIKE '%End Bearing%' OR
			sp.item ILIKE '%Max. Range%' OR
			sp.item ILIKE '%Min. Range%'
		)
		GROUP BY group_swc, environment, sp.type
		LIMIT ? OFFSET ?
    `
	result := db.Raw(query, pattern, pageSize, offset).Scan(&results)
	if result.Error != nil {
		return nil
	}
	return results
}

func GetSWCParameterEnvNonAirResults(db *gorm.DB, page, pageSize int32) []model.SWCParameterResult {
	var results []model.SWCParameterResult

	// Calculate the offset for pagination
	offset := (page - 1) * pageSize

	query := `
		SELECT group_swc , environment, sp.type, 
               STRING_AGG(CAST(id AS TEXT), ',') AS id,  
                   STRING_AGG(CAST(value AS TEXT), ',') AS value,  
                   STRING_AGG(CAST(default_value AS TEXT), ',') AS default_value, 
                   STRING_AGG(CAST(item AS TEXT), ',') AS item, 
                   STRING_AGG(CAST(created_at AS TEXT), ',') AS created_at, 
                   STRING_AGG(CAST(updated_at AS TEXT), ',') AS updated_at, 
                   STRING_AGG(CAST(is_deleted AS TEXT), ',') AS is_deleted, 
                   STRING_AGG(CAST(unit AS TEXT), ',') AS unit 
            FROM swc_parameter sp WHERE sp.is_deleted = false AND sp.environment = 'air'  
                AND ( 
                    sp.item ILIKE '%Start Bearing%' OR 
                    sp.item ILIKE '%End Bearing%' OR 
                    sp.item ILIKE '%Max. Range%' OR 
                    sp.item ILIKE '%Min. Range%' 
                ) 
            GROUP BY group_swc, environment, sp.type
            HAVING  
                STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%Start Bearing%'  
                AND STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%End Bearing%' 
                AND STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%Max. Range%' 
                AND STRING_AGG(CAST(item AS TEXT), ',') ILIKE '%Min. Range%'
			LIMIT ? OFFSET ?
		`

	result := db.Raw(query, pageSize, offset).Scan(&results)
	if result.Error != nil {
		return nil
	}

	return results

}

func GetSWCCount(db *gorm.DB) []model.SWCParameterResult {
	var results []model.SWCParameterResult

	query := `
		SELECT group_swc , environment, sp.type, 
               STRING_AGG(CAST(id AS TEXT), ',') AS id,  
                   STRING_AGG(CAST(value AS TEXT), ',') AS value,  
                   STRING_AGG(CAST(default_value AS TEXT), ',') AS default_value, 
                   STRING_AGG(CAST(item AS TEXT), ',') AS item, 
                   STRING_AGG(CAST(created_at AS TEXT), ',') AS created_at, 
                   STRING_AGG(CAST(updated_at AS TEXT), ',') AS updated_at, 
                   STRING_AGG(CAST(is_deleted AS TEXT), ',') AS is_deleted, 
                   STRING_AGG(CAST(unit AS TEXT), ',') AS unit 
            FROM swc_parameter sp WHERE sp.is_deleted = false  
            GROUP BY group_swc, environment, sp.type
	`
	result := db.Raw(query).Scan(&results)
	if result.Error != nil {
		return nil
	}

	return results
}

func GetSWCCountEnvi(db *gorm.DB, pattern string) []model.SWCParameterResult {
	var results []model.SWCParameterResult

	query := `
	SELECT group_swc , environment, sp.type, 
               STRING_AGG(CAST(id AS TEXT), ',') AS id,  
                   STRING_AGG(CAST(value AS TEXT), ',') AS value,  
                   STRING_AGG(CAST(default_value AS TEXT), ',') AS default_value, 
                   STRING_AGG(CAST(item AS TEXT), ',') AS item, 
                   STRING_AGG(CAST(created_at AS TEXT), ',') AS created_at, 
                   STRING_AGG(CAST(updated_at AS TEXT), ',') AS updated_at, 
                   STRING_AGG(CAST(is_deleted AS TEXT), ',') AS is_deleted, 
                   STRING_AGG(CAST(unit AS TEXT), ',') AS unit 
            FROM swc_parameter sp WHERE sp.is_deleted = false AND sp.environment = CAST(? AS TEXT)  
            GROUP BY group_swc, environment, sp.type
	`
	result := db.Raw(query, pattern).Scan(&results)
	if result.Error != nil {
		return nil
	}

	return results
}
