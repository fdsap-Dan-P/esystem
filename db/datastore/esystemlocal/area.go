package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const createArea = `-- name: CreateArea: one
IF EXISTS (SELECT AreaCode FROM Area WHERE AreaCode = $1)
BEGIN
  UPDATE Area SET 
    Area = $2 
  WHERE AreaCode = $1
END ELSE BEGIN
  INSERT INTO Area (
    AreaCode, Area
  ) 
  VALUES ($1, $2) 
END
`

type AreaRequest struct {
	AreaCode int64          `json:"areaCode"`
	Area     sql.NullString `json:"area"`
}

func (q *QueriesLocal) CreateArea(ctx context.Context, arg AreaRequest) error {
	_, err := q.db.ExecContext(ctx, createArea,
		arg.AreaCode,
		arg.Area,
	)
	return err
}

const deleteArea = `-- name: DeleteArea :exec
DELETE FROM Area WHERE AreaCode = $1
`

func (q *QueriesLocal) DeleteArea(ctx context.Context, areaCode int64) error {
	_, err := q.db.ExecContext(ctx, deleteArea, areaCode)
	return err
}

type AreaInfo struct {
	ModCtr    int64          `json:"modCtr"`
	BrCode    string         `json:"brCode"`
	ModAction string         `json:"modAction"`
	AreaCode  int64          `json:"areaCode"`
	Area      sql.NullString `json:"area"`
}

// -- name: GetArea :one
const getArea = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, AreaCode, Area
FROM OrgParms, Area d
INNER JOIN Modified m on m.UniqueKeyInt1 = d.AreaCode
`

func scanRowArea(row *sql.Row) (AreaInfo, error) {
	var i AreaInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.AreaCode,
		&i.Area,
	)
	return i, err
}

func scanRowsArea(rows *sql.Rows) ([]AreaInfo, error) {
	items := []AreaInfo{}
	for rows.Next() {
		var i AreaInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.AreaCode,
			&i.Area,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (q *QueriesLocal) GetArea(ctx context.Context, areaCode int64) (AreaInfo, error) {
	sql := fmt.Sprintf("%s WHERE m.TableName = 'Area' AND Uploaded = 0 AND AreaCode = $1", getArea)
	row := q.db.QueryRowContext(ctx, sql, areaCode)
	return scanRowArea(row)
}

type ListAreaParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) AreaCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, AreaCode, Area
FROM OrgParms, Area d
`, filenamePath)
}

func (q *QueriesLocal) ListArea(ctx context.Context) ([]AreaInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE m.TableName = 'Area' AND Uploaded = 0`,
		getArea)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsArea(rows)
}

// -- name: UpdateArea :one
const updateArea = `
UPDATE Area SET 
  Area = $2 
WHERE AreaCode = $1`

func (q *QueriesLocal) UpdateArea(ctx context.Context, arg AreaRequest) error {
	_, err := q.db.ExecContext(ctx, updateArea,
		arg.AreaCode,
		arg.Area,
	)
	return err
}

func (q *QueriesLocal) UpdateModCtr(ctx context.Context, arg AreaRequest) error {
	_, err := q.db.ExecContext(ctx, updateArea,
		arg.AreaCode,
		arg.Area,
	)
	return err
}
