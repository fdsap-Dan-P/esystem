package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"simplebank/model"
)

const createInventoryItem = `-- name: CreateInventoryItem: one
INSERT INTO Inventory_Item (
    Uuid, Item_Code, Bar_Code, Item_Name, Unique_Variation, Parent_ID, Generic_Name_ID,
    Brand_Name_ID, Measure_ID, Image_ID, Remarks, Other_Info
 ) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
 ON CONFLICT(UUID)
 DO UPDATE SET
    Item_Code = EXCLUDED.Item_Code,
	bar_code =  EXCLUDED.bar_code,
	item_name =  EXCLUDED.item_name,
	unique_variation =  EXCLUDED.unique_variation,
	parent_id =  EXCLUDED.parent_id,
	generic_name_id =  EXCLUDED.generic_name_id,
	brand_name_id =  EXCLUDED.brand_name_id,
	measure_id =  EXCLUDED.measure_id,
	image_id =  EXCLUDED.image_id,
	remarks =  EXCLUDED.remarks,
	vec_simple_name =  EXCLUDED.vec_simple_name,
	other_info =  EXCLUDED.other_info
RETURNING Id, UUID, Item_Code, Bar_Code, Item_Name, Unique_Variation, Parent_ID, Generic_Name_ID,
 Brand_Name_ID, Measure_ID, Image_ID, Remarks, Other_Info
`

type InventoryItemRequest struct {
	Id                int64                         `json:"Id"`
	Uuid              uuid.UUID                     `json:"uuid"`
	ItemCode          string                        `json:"itemCode"`
	BarCode           sql.NullString                `json:"barCode"`
	ItemName          string                        `json:"itemName"`
	UniqueVariation   string                        `json:"uniqueVariation"`
	ParentId          sql.NullInt64                 `json:"parentId"`
	GenericNameId     sql.NullInt64                 `json:"genericNameId"`
	BrandNameId       sql.NullInt64                 `json:"brandNameId"`
	MeasureId         int64                         `json:"measureId"`
	ImageId           sql.NullInt64                 `json:"imageId"`
	Remarks           string                        `json:"remarks"`
	OtherInfo         sql.NullString                `json:"otherInfo"`
	InventoryItemList []InventoryItemRequest        `json:"inventoryItemList"`
	SpecsStringList   []InventorySpecsStringRequest `json:"specsStringList"`
	SpecsNumberList   []InventorySpecsNumberRequest `json:"specsNumberList"`
	SpecsDateList     []InventorySpecsDateRequest   `json:"specsDateList"`
}

func (q *QueriesAccount) CreateInventoryItem(ctx context.Context, arg InventoryItemRequest) (model.InventoryItem, error) {
	row := q.db.QueryRowContext(ctx, createInventoryItem,
		arg.Uuid,
		arg.ItemCode,
		arg.BarCode,
		arg.ItemName,
		arg.UniqueVariation,
		arg.ParentId,
		arg.GenericNameId,
		arg.BrandNameId,
		arg.MeasureId,
		arg.ImageId,
		arg.Remarks,
		arg.OtherInfo,
	)

	var i model.InventoryItem
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ItemCode,
		&i.BarCode,
		&i.ItemName,
		&i.UniqueVariation,
		&i.ParentId,
		&i.GenericNameId,
		&i.BrandNameId,
		&i.MeasureId,
		&i.ImageId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteInventoryItem = `-- name: DeleteInventoryItem :exec
DELETE FROM Inventory_Item
WHERE id = $1
`

func (q *QueriesAccount) DeleteInventoryItem(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteInventoryItem, id)
	return err
}

type InventoryItemInfo struct {
	Id                 int64          `json:"Id"`
	Uuid               uuid.UUID      `json:"uuid"`
	ItemCode           string         `json:"itemCode"`
	BarCode            sql.NullString `json:"barCode"`
	ItemName           string         `json:"itemName"`
	UniqueVariation    string         `json:"uniqueVariation"`
	ParentId           sql.NullInt64  `json:"parentId"`
	ParentCode         sql.NullInt64  `json:"parentCode"`
	ParentItemName     sql.NullString `json:"parentItemName"`
	GenericNameId      sql.NullInt64  `json:"genericNameId"`
	GenericCode        sql.NullInt64  `json:"genericCode"`
	GenericShortName   sql.NullString `json:"genericShortName"`
	GenericTitle       sql.NullString `json:"genericTitle"`
	BrandNameId        sql.NullInt64  `json:"brandNameId"`
	BrandNameCode      sql.NullInt64  `json:"brandNameCode"`
	BrandNameShortName sql.NullString `json:"brandNameShortName"`
	BrandNameTitle     sql.NullString `json:"brandNameTitle"`
	MeasureId          int64          `json:"measureId"`
	MeasureCode        sql.NullInt64  `json:"measureCode"`
	MeasureShortName   sql.NullString `json:"measureShortName"`
	MeasureTitle       sql.NullString `json:"measureTitle"`
	ImageId            sql.NullInt64  `json:"imageId"`
	ImageCode          sql.NullString `json:"imageCode"`
	ImageFilePath      sql.NullString `json:"imageFilePath"`
	ImageThumbnail     []byte         `json:"imageThumbnail"`
	Remarks            string         `json:"remarks"`
	OtherInfo          sql.NullString `json:"otherInfo"`
	ModCtr             int64          `json:"modCtr"`
	Created            sql.NullTime   `json:"created"`
	Updated            sql.NullTime   `json:"updated"`
	Weight             float64        `json:"weight"`
	VecSimpleName      string         `json:"vecSimpleName"`

	// Child                []model.InventoryItem        `json:"child"`
	// InventorySpecsNumber []model.InventorySpecsNumber `json:"inventorySpecsNumber"`
	// InventorySpecsDate   []model.InventorySpecsDate   `json:"inventorySpecsDate"`
	// InventorySpecsString []model.InventorySpecsString `json:"inventorySpecsString"`
}

func populateInventoryItem(q *QueriesAccount, ctx context.Context, sql string, param ...interface{}) ([]InventoryItemInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql, param...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []InventoryItemInfo{}
	for rows.Next() {
		var i InventoryItemInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.ItemCode,
			&i.BarCode,
			&i.ItemName,
			&i.UniqueVariation,
			&i.ParentId,
			&i.ParentCode,
			&i.ParentItemName,
			&i.GenericNameId,
			&i.GenericCode,
			&i.GenericShortName,
			&i.GenericTitle,
			&i.BrandNameId,
			&i.BrandNameCode,
			&i.BrandNameShortName,
			&i.BrandNameTitle,
			&i.MeasureId,
			&i.MeasureCode,
			&i.MeasureShortName,
			&i.MeasureTitle,
			&i.ImageId,
			&i.ImageCode,
			&i.ImageFilePath,
			&i.ImageThumbnail,
			&i.Remarks,
			&i.OtherInfo,
			&i.ModCtr,
			&i.Created,
			&i.Updated,
			&i.Weight,
			&i.VecSimpleName,
		)
		if err != nil {
			return items, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return items, err
	}
	if err := rows.Err(); err != nil {
		return items, err
	}
	return items, nil
}

const inventoryItemSQL = `-- name: inventoryItemSQL
SELECT 
  d.Id, mr.Uuid, d.Item_Code, d.Bar_Code, d.Item_Name, d.Unique_Variation, 
  d.Parent_ID, p.Item_Code Parent_Code, p.Item_Name Parent_Item_Name, 
  gn.Id Generic_Name_ID, gn.Code Generic_Code, 
  gn.short_name Generic_Short_Name, gn.Title Generic_Title,
  br.Id Brand_Name_ID, br.Code Brand_Name_Code, 
  br.short_name Brand_Name_Short_Name, br.Title Brand_Name_Title,
  ms.Id Measure_ID, ms.Code Measure_Code, 
  ms.short_name Measure_Short_Name, ms.Title Measure_Title,
  dc.Id Image_ID, dc.Code Image_Code, dc.File_Path Image_File_Path, 
  dc.ThumbNail Image_ThumbNail,
  d.Remarks, d.Other_Info, 
  mr.Mod_Ctr, mr.Created, mr.Updated, 0 Weight, d.Vec_Simple_Name
FROM Inventory_Item d 
LEFT JOIN Inventory_Item p on p.Id = d.Parent_ID 
LEFT JOIN Reference br  on br.Id = d.Brand_Name_ID
LEFT JOIN Reference gn on gn.Id = d.Generic_Name_ID
LEFT JOIN Reference ms on ms.Id = d.Measure_ID
LEFT JOIN Documents dc on dc.Id = d.Image_Id
INNER JOIN Main_Record mr on mr.Uuid = d.Uuid 
`

func (q *QueriesAccount) GetInventoryItem(
	ctx context.Context, id int64) (InventoryItemInfo, error) {
	script := fmt.Sprintf(`%v WHERE d.id = $1`, inventoryItemSQL)
	log.Printf("stript %v: %v", id, script)
	items, err := populateInventoryItem(q, ctx, script, id)

	log.Printf("GetInventoryItem %v", items)
	if len(items) > 0 {
		return items[0], err
	} else {
		return InventoryItemInfo{}, fmt.Errorf("inventoryItem ID:%v not found", id)
	}
}

func (q *QueriesAccount) GetInventoryItembyCode(
	ctx context.Context, itemCode string) (InventoryItemInfo, error) {
	script := fmt.Sprintf(`%v WHERE d.Item_Code = $1`, inventoryItemSQL)
	log.Printf("stript %v: %v", itemCode, script)
	items, err := populateInventoryItem(q, ctx, script, itemCode)

	log.Printf("GetInventoryItem %v", items)
	if len(items) > 0 {
		return items[0], err
	} else {
		return InventoryItemInfo{}, fmt.Errorf("inventoryItem itemCode:%v not found", itemCode)
	}
}

func (q *QueriesAccount) GetInventoryItembyUuid(ctx context.Context, uuid uuid.UUID) (InventoryItemInfo, error) {
	script := fmt.Sprintf(`%v WHERE mr.uuid = $1`, inventoryItemSQL)
	items, err := populateInventoryItem(q, ctx, script, uuid)
	if len(items) > 0 {
		return items[0], err
	} else {
		return InventoryItemInfo{}, fmt.Errorf("inventoryItem UUID:%v not found", uuid)
	}
}

func (q *QueriesAccount) GetInventoryItembyName(ctx context.Context, name string) (InventoryItemInfo, error) {
	script := fmt.Sprintf(`%v WHERE lower(d.Item_Name) = $1`, inventoryItemSQL)
	items, err := populateInventoryItem(q, ctx, script, name)
	if len(items) > 0 {
		return items[0], err
	} else {
		return InventoryItemInfo{}, fmt.Errorf("inventoryItem Name:%v not found", name)
	}
}

type ListInventoryItembyGenericParams struct {
	GenericId int64 `json:"genericId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesAccount) ListInventoryItembyGeneric(ctx context.Context, arg ListInventoryItembyGenericParams) ([]InventoryItemInfo, error) {
	sql := fmt.Sprintf(`%v WHERE d.Generic_Name_ID = $1 LIMIT $2 OFFSET $3 `, inventoryItemSQL)
	return populateInventoryItem(q, ctx, sql, arg.GenericId, arg.Limit, arg.Offset)
}

type ListInventoryItembyBrandParams struct {
	BrandId int64 `json:"brandId"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *QueriesAccount) ListInventoryItembyBrand(ctx context.Context, arg ListInventoryItembyBrandParams) ([]InventoryItemInfo, error) {
	sql := fmt.Sprintf(`%v WHERE d.Brand_Name_ID = $1 LIMIT $2 OFFSET $3 `, inventoryItemSQL)
	return populateInventoryItem(q, ctx, sql, arg.BrandId, arg.Limit, arg.Offset)
}

type InventoryItemFilterParams struct {
	Filter string `json:"filter"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *QueriesAccount) InventoryItemFilter(ctx context.Context, arg InventoryItemFilterParams) ([]InventoryItemInfo, error) {
	sql := fmt.Sprintf("%s FROM ( %s ) d, %s) s WHERE Vec_Simple_Name @@ loc_or ORDER BY %s",
		"SELECT d.* ",
		inventoryItemSQL,
		fmt.Sprintf("(SELECT plainto_tsquery_or('%s') loc_or", arg.Filter),
		"ts_rank_cd(Vec_Simple_Name, loc_or, 32) desc")
	log.Println(sql)
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d", sql, arg.Limit, arg.Offset)
	}
	return populateInventoryItem(q, ctx, sql)
}

type SpecsCondition string
type ConditionValueUsed int8

const (
	Value1 ConditionValueUsed = 1
	Value2 ConditionValueUsed = 2
	Both   ConditionValueUsed = 3
)

const (
	Equal              SpecsCondition = "Equal"
	NotEqual           SpecsCondition = "NotEqual"
	BeginsWith         SpecsCondition = "BeginsWith"
	Contains           SpecsCondition = "Contains"
	EndsWith           SpecsCondition = "EndsWith"
	GreaterThan        SpecsCondition = "GreaterThan"
	LessThan           SpecsCondition = "LessThan"
	GreaterThanorEqual SpecsCondition = "GreaterThanorEqual"
	LessThanorEqual    SpecsCondition = "LessThanorEqual"
	Between            SpecsCondition = "Between"
)

type SearchSpecsString struct {
	ItemCode   int64          `json:"ItemCode"`
	ItemId     int64          `json:"itemid"`
	Condition  SpecsCondition `json:"Condition"`
	Value      string         `json:"value"`
	Weight     float64        `json:"weight"`
	UsedWeight float64        `json:"usedWeight"`
}

type SearchSpecsDate struct {
	ItemId     int64              `json:"itemid"`
	ItemCode   int64              `json:"ItemCode"`
	Condition  SpecsCondition     `json:"Condition"`
	Value1     time.Time          `json:"value1"`
	Value2     time.Time          `json:"value2"`
	ValueUsed  ConditionValueUsed `json:"valueUsed"`
	Weight     float64            `json:"weight"`
	UsedWeight float64            `json:"usedWeight"`
}

type SearchSpecsNumber struct {
	ItemId     int64              `json:"itemid"`
	ItemCode   int64              `json:"ItemCode"`
	Condition  SpecsCondition     `json:"Condition"`
	Value1     decimal.Decimal    `json:"value1"`
	Value2     decimal.Decimal    `json:"value2"`
	ValueUsed  ConditionValueUsed `json:"valueUsed"`
	Weight     float64            `json:"weight"`
	UsedWeight float64            `json:"usedWeight"`
}

type InventoryItemSearchParams struct {
	SearchSpecsString []SearchSpecsString `json:"searchSpecsString"`
	SearchSpecsDate   []SearchSpecsDate   `json:"searchSpecsDate"`
	SearchSpecsNumber []SearchSpecsNumber `json:"searchSpecsNumber"`
	Limit             int32               `json:"limit"`
	Offset            int32               `json:"offset"`
}

func (inv InventoryItemSearchParams) searchSql() string {
	var filterCtr int = 0
	var wctr int = 0
	var w float64 = 0
	var usedw float64 = 0
	for _, a := range inv.SearchSpecsString {
		filterCtr++
		if a.Weight > 0 {
			w += a.Weight
			wctr++
		}
	}
	for _, a := range inv.SearchSpecsDate {
		filterCtr++
		if a.Weight > 0 {
			w += a.Weight
			wctr++
		}
	}
	for _, a := range inv.SearchSpecsNumber {
		filterCtr++
		if a.Weight > 0 {
			w += a.Weight
			wctr++
		}
	}

	usedw = 1
	if wctr > 0 {
		usedw = (w / float64(wctr))
	}

	w += float64(filterCtr-wctr) * usedw

	log.Printf("sdfsdf usedw: %v; filterCtr:%v; wctr:%v", usedw, filterCtr, wctr)
	log.Printf("sdfsdf w: %v", w)

	var sql string = ""
	var s string = ""
	var weight float64 = 0
	for i, a := range inv.SearchSpecsString {
		if a.Weight > 0 {
			weight = a.Weight / w
		} else {
			weight = usedw / w
		}
		inv.SearchSpecsString[i].UsedWeight = w
		switch a.Condition {
		case Equal:
			s = fmt.Sprintf("(specs_id = %v and lower(value) = '%v')", a.ItemId, strings.ToLower(a.Value))
		case NotEqual:
			s = fmt.Sprintf("(specs_id = %v and lower(value) <> '%v')", a.ItemId, strings.ToLower(a.Value))
		case BeginsWith:
			s = fmt.Sprintf("(specs_id = %v and lower(value) like '%%%v')", a.ItemId, strings.ToLower(a.Value))
		case Contains:
			s = fmt.Sprintf("(specs_id = %v and lower(value) like '%%%v%%')", a.ItemId, strings.ToLower(a.Value))
		case EndsWith:
			s = fmt.Sprintf("(specs_id = %v and lower(value) like '%v%%')", a.ItemId, strings.ToLower(a.Value))
		default:
			s = fmt.Sprintf("(specs_id = %v and lower(value) = '%v')", a.ItemId, strings.ToLower(a.Value))
		}

		// Equal              SpecsCondition = "Equal"
		// NotEqual           SpecsCondition = "NotEqual"
		// BeginsWith         SpecsCondition = "BeginsWith"
		// Contains           SpecsCondition = "Contains"
		// EndsWith           SpecsCondition = "EndsWith"
		// GreaterThan        SpecsCondition = "GreaterThan"
		// LessThan           SpecsCondition = "LessThan"
		// GreaterThanorEqual SpecsCondition = "GreaterThanorEqual"
		// LessThanorEqual    SpecsCondition = "LessThanorEqual"
		// Between
		if sql != "" {
			sql += "\n UNION ALL \n"
		}
		sql += fmt.Sprintf(
			"SELECT inventory_item_id, %v weight FROM inventory_specs_string WHERE %v",
			weight, s)
	}
	var valueUsed string = "value"

	for i, a := range inv.SearchSpecsDate {
		if a.Weight > 0 {
			weight = a.Weight / w
		} else {
			weight = usedw / w
		}
		inv.SearchSpecsDate[i].UsedWeight = w

		switch a.ValueUsed {
		case Value2:
			valueUsed = "value2"
		default:
			valueUsed = "value"
		}

		switch a.Condition {
		case Equal:
			if a.ValueUsed == Both {
				s = fmt.Sprintf("(specs_id = %v and (value = '%v' or value2 = '%v'))", a.ItemId, a.Value1.Format("2006-01-02"), a.Value1.Format("2006-01-02"))
			} else {
				s = fmt.Sprintf("(specs_id = %v and %s = '%v')", a.ItemId, valueUsed, a.Value1.Format("2006-01-02"))
			}
		case NotEqual:
			if a.ValueUsed == Both {
				s = fmt.Sprintf("(specs_id = %v and (value <> '%v' or value2 <> '%v'))", a.ItemId, a.Value1.Format("2006-01-02"), a.Value1.Format("2006-01-02"))
			} else {
				s = fmt.Sprintf("(specs_id = %v and %s <> '%v')", a.ItemId, valueUsed, a.Value1.Format("2006-01-02"))
			}
		case GreaterThan:
			if a.ValueUsed == Both {
				s = fmt.Sprintf("(specs_id = %v and (value > '%v' or value2 > '%v'))", a.ItemId, a.Value1.Format("2006-01-02"), a.Value1.Format("2006-01-02"))
			} else {
				s = fmt.Sprintf("(specs_id = %v and %s > '%v')", a.ItemId, valueUsed, a.Value1.Format("2006-01-02"))
			}
		case LessThan:
			if a.ValueUsed == Both {
				s = fmt.Sprintf("(specs_id = %v and (value < '%v' or value2 < '%v'))", a.ItemId, a.Value1.Format("2006-01-02"), a.Value1.Format("2006-01-02"))
			} else {
				s = fmt.Sprintf("(specs_id = %v and %s < '%v')", a.ItemId, valueUsed, a.Value1.Format("2006-01-02"))
			}
		case GreaterThanorEqual:
			if a.ValueUsed == Both {
				s = fmt.Sprintf("(specs_id = %v and (value >= '%v' or value2 >= '%v'))", a.ItemId, a.Value1.Format("2006-01-02"), a.Value1.Format("2006-01-02"))
			} else {
				s = fmt.Sprintf("(specs_id = %v and %s >= '%v')", a.ItemId, valueUsed, a.Value1.Format("2006-01-02"))
			}
		case LessThanorEqual:
			if a.ValueUsed == Both {
				s = fmt.Sprintf("(specs_id = %v and (value <= '%v' or value2 <= '%v'))", a.ItemId, a.Value1.Format("2006-01-02"), a.Value1.Format("2006-01-02"))
			} else {
				s = fmt.Sprintf("(specs_id = %v and %s <= '%v')", a.ItemId, valueUsed, a.Value1.Format("2006-01-02"))
			}
		case Between:
			if a.ValueUsed == Both {
				s = fmt.Sprintf("(specs_id = %v and (value between '%v' and '%v' or value2 between '%v' and '%v'))", a.ItemId, a.Value1.Format("2006-01-02"), a.Value1.Format("2006-01-02"), a.Value1.Format("2006-01-02"), a.Value1.Format("2006-01-02"))
			} else {
				s = fmt.Sprintf("(specs_id = %v and %s between '%v' and '%v')", a.ItemId, valueUsed, a.Value1.Format("2006-01-02"), a.Value2.Format("2006-01-02"))
			}
		default:
			if a.ValueUsed == Both {
				s = fmt.Sprintf("(specs_id = %v and (value = '%v' or value2 = '%v'))", a.ItemId, a.Value1.Format("2006-01-02"), a.Value1.Format("2006-01-02"))
			} else {
				s = fmt.Sprintf("(specs_id = %v and %s = '%v')", a.ItemId, valueUsed, a.Value1.Format("2006-01-02"))
			}
		}

		if sql != "" {
			sql += "\n UNION ALL \n"
		}
		sql += fmt.Sprintf(
			"SELECT inventory_item_id, %v weight FROM inventory_specs_date WHERE %v",
			weight, s)
	}

	for i, a := range inv.SearchSpecsNumber {
		if a.Weight > 0 {
			weight = a.Weight / w
		} else {
			weight = usedw / w
		}
		inv.SearchSpecsNumber[i].UsedWeight = w

		switch a.ValueUsed {
		case Value2:
			valueUsed = "value2"
		default:
			valueUsed = "value"
		}

		switch a.Condition {
		case Equal:
			if a.ValueUsed == Both {
				s = fmt.Sprintf("(specs_id = %v and (value = %v or value2 = %v))", a.ItemId, a.Value1, a.Value1)
			} else {
				s = fmt.Sprintf("(specs_id = %v and %s = %v)", a.ItemId, valueUsed, a.Value1)
			}
		case NotEqual:
			if a.ValueUsed == Both {
				s = fmt.Sprintf("(specs_id = %v and (value <> '%v' or value2 <> '%v'))", a.ItemId, a.Value1, a.Value1)
			} else {
				s = fmt.Sprintf("(specs_id = %v and %s <> '%v')", a.ItemId, valueUsed, a.Value1)
			}
		case GreaterThan:
			if a.ValueUsed == Both {
				s = fmt.Sprintf("(specs_id = %v and (value > '%v' or value2 > '%v'))", a.ItemId, a.Value1, a.Value1)
			} else {
				s = fmt.Sprintf("(specs_id = %v and %s > '%v')", a.ItemId, valueUsed, a.Value1)
			}
		case LessThan:
			if a.ValueUsed == Both {
				s = fmt.Sprintf("(specs_id = %v and (value < '%v' or value2 < '%v'))", a.ItemId, a.Value1, a.Value1)
			} else {
				s = fmt.Sprintf("(specs_id = %v and %s < '%v')", a.ItemId, valueUsed, a.Value1)
			}
		case GreaterThanorEqual:
			if a.ValueUsed == Both {
				s = fmt.Sprintf("(specs_id = %v and (value >= %v or value2 >= %v))", a.ItemId, a.Value1, a.Value1)
			} else {
				s = fmt.Sprintf("(specs_id = %v and %s >= %v)", a.ItemId, valueUsed, a.Value1)
			}
		case LessThanorEqual:
			if a.ValueUsed == Both {
				s = fmt.Sprintf("(specs_id = %v and (value <= %v or value2 <= %v))", a.ItemId, a.Value1, a.Value1)
			} else {
				s = fmt.Sprintf("(specs_id = %v and %s <= %v)", a.ItemId, valueUsed, a.Value1)
			}
		case Between:
			if a.ValueUsed == Both {
				s = fmt.Sprintf("(specs_id = %v and (value between %v and %v or value2 between %v and %v))", a.ItemId, a.Value1, a.Value1, a.Value1, a.Value1)
			} else {
				s = fmt.Sprintf("(specs_id = %v and %s between %v and %v)", a.ItemId, valueUsed, a.Value1, a.Value2)
			}
		default:
			if a.ValueUsed == Both {
				s = fmt.Sprintf("(specs_id = %v and (value = %v or value2 = %v))", a.ItemId, a.Value1, a.Value1)
			} else {
				s = fmt.Sprintf("(specs_id = %v and %s = %v)", a.ItemId, valueUsed, a.Value1)
			}
		}

		if sql != "" {
			sql += "\nUNION ALL\n"
		}
		sql += fmt.Sprintf(
			"SELECT inventory_item_id, %v weight FROM inventory_specs_number WHERE %v",
			weight, s)
	}

	log.Printf(fmt.Sprintf("SELECT inventory_item_id, sum(weight) weight FROM (%s) s GROUP BY inventory_item_id", sql))
	return fmt.Sprintf("SELECT inventory_item_id, sum(weight) weight FROM (%s) s GROUP BY inventory_item_id", sql)
}

func (q *QueriesAccount) InventoryItemSearch(ctx context.Context, arg InventoryItemSearchParams) ([]InventoryItemInfo, error) {
	s := arg.searchSql()
	log.Println("--------------- TEST: InventoryItemSearch -------------")
	log.Println(s)
	sql := fmt.Sprintf(`
SELECT 
  d.id, d.uuid, d.item_code, d.bar_code, d.item_name, 
  d.unique_variation, d.parent_id, d.parent_code, d.parent_item_name, d.generic_name_id, 
  d.generic_code, d.generic_short_name, d.generic_title, d.brand_name_id, d.brand_name_code, 
  d.brand_name_short_name, d.brand_name_title, d.measure_id, d.measure_code, 
  d.measure_short_name, d.measure_title, d.image_id, d.image_code, 
  d.image_file_path, d.image_thumbnail, d.remarks, d.other_info, 
  d.mod_ctr, d.created, d.updated, d.weight, d.vec_simple_name
FROM (%s) d, (%s) s WHERE s.inventory_item_id = d.id `, inventoryItemSQL, s)

	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d", sql, arg.Limit, arg.Offset)
	}
	log.Println(sql)
	return populateInventoryItem(q, ctx, sql)
}

const updateInventoryItem = `-- name: UpdateInventoryItem :one
UPDATE Inventory_Item SET 
  Item_Code = $2,
  Bar_Code = $3,
  Item_Name = $4,
  Unique_Variation = $5,
  Parent_Id = $6,
  Generic_Name_Id = $7,
  Brand_Name_Id = $8,
  Measure_Id = $9,
  Image_Id = $10,
  Remarks = $11,
  Other_Info = $12
WHERE id = $1
RETURNING  Id, UUID, Item_Code, Bar_Code, Item_Name, Unique_Variation, Parent_ID, Generic_Name_ID,
Brand_Name_ID, Measure_ID, Image_ID, Remarks, Other_Info
`

func (q *QueriesAccount) UpdateInventoryItem(
	ctx context.Context, arg InventoryItemRequest) (model.InventoryItem, error) {
	row := q.db.QueryRowContext(ctx, updateInventoryItem,
		arg.Id,
		arg.ItemCode,
		arg.BarCode,
		arg.ItemName,
		arg.UniqueVariation,
		arg.ParentId,
		arg.GenericNameId,
		arg.BrandNameId,
		arg.MeasureId,
		arg.ImageId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.InventoryItem
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ItemCode,
		&i.BarCode,
		&i.ItemName,
		&i.UniqueVariation,
		&i.ParentId,
		&i.GenericNameId,
		&i.BrandNameId,
		&i.MeasureId,
		&i.ImageId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

// TransferTx performs a money transfer from one account to the other.
// It creates the transfer, add account entries, and update accounts' balance within a database transaction
func (store *SQLStoreAccount) CreateInventoryItemFull(
	ctx context.Context, arg InventoryItemRequest) (model.InventoryItem, error) {

	var result model.InventoryItem

	err := store.ExecTx(ctx, func(q *QueriesAccount) error {
		var err error

		result, err = store.CreateInventoryItem(context.Background(), arg)
		if err != nil {
			return err
		}

		fmt.Printf("CreateInventoryItemFull:Get by d1%+v\n", arg.SpecsNumberList)
		for _, specs := range arg.SpecsDateList {
			specs.InventoryItemId = result.Id
			_, r := store.CreateInventorySpecsDate(context.Background(), specs)
			if r != nil {
				return err
			}
		}

		for _, specs := range arg.SpecsStringList {
			specs.InventoryItemId = result.Id
			_, r := store.CreateInventorySpecsString(context.Background(), specs)
			if r != nil {
				return err
			}
		}

		for _, specs := range arg.SpecsNumberList {
			specs.InventoryItemId = result.Id
			_, r := store.CreateInventorySpecsNumber(context.Background(), specs)
			if r != nil {
				return err
			}
		}

		for _, inv := range arg.InventoryItemList {
			inv.ParentId.Int64 = result.Id
			inv.ParentId.Valid = true

			_, r := store.CreateInventoryItemFull(context.Background(), inv)
			if r != nil {
				return err
			}
		}

		return err
	})

	return result, err
}
