package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const createCenter = `-- name: CreateCenter: one
IF EXISTS (SELECT Center_Code FROM CentWorker_Det WHERE Center_Code = $1) 
BEGIN 
  UPDATE Center SET 
    Center_Name = $2,
    Center_Address = $3,
    Center_Meet_Day = $4,
    Unit = $5,
    DateEstablished = $6
  WHERE Center_Code = $1;
END ELSE BEGIN
  INSERT INTO Center (
	Center_Code, Center_Name, Center_Address, Center_Meet_Day, Unit, DateEstablished
  ) 
  VALUES ($1, $2, $3, $4, $5, $6);
END;
IF EXISTS (SELECT Center_Code FROM CentWorker_Det WHERE Center_Code = $1) 
BEGIN 
  UPDATE CentWorker_Det 
    SET CenterW_Id = 16 
    WHERE Center_Code = $1;
END ELSE BEGIN
  INSERT INTO CentWorker_Det(Center_Code, CenterW_Id) 
    VALUES ($1, $7);
END;
`

type CenterRequest struct {
	CenterCode      string         `json:"centerCode"`
	CenterName      sql.NullString `json:"centerName"`
	CenterAddress   sql.NullString `json:"centerAddress"`
	MeetingDay      sql.NullInt64  `json:"meetingDay"`
	Unit            sql.NullInt64  `json:"unit"`
	DateEstablished sql.NullTime   `json:"dateEstablished"`
	AOID            sql.NullInt64  `json:"aoId"`
}

func (q *QueriesLocal) CreateCenter(ctx context.Context, arg CenterRequest) error {
	_, err := q.db.ExecContext(ctx, createCenter,
		arg.CenterCode,
		arg.CenterName,
		arg.CenterAddress,
		arg.MeetingDay,
		arg.Unit,
		arg.DateEstablished,
		arg.AOID,
	)
	return err
}

const deleteCenter = `-- name: DeleteCenter :exec
DELETE FROM CentWorker_Det WHERE Center_Code = $1;
DELETE FROM Center WHERE Center_Code = $1;
`

func (q *QueriesLocal) DeleteCenter(ctx context.Context, centerCode string) error {
	_, err := q.db.ExecContext(ctx, deleteCenter, centerCode)
	return err
}

type CenterInfo struct {
	ModCtr          int64          `json:"modCtr"`
	BrCode          string         `json:"brCode"`
	ModAction       string         `json:"modAction"`
	CenterCode      string         `json:"centerCode"`
	CenterName      sql.NullString `json:"centerName"`
	CenterAddress   sql.NullString `json:"centerAddress"`
	MeetingDay      sql.NullInt64  `json:"meetingDay"`
	Unit            sql.NullInt64  `json:"unit"`
	DateEstablished sql.NullTime   `json:"dateEstablished"`
	AOID            sql.NullInt64  `json:"aoID"`
}

// -- name: GetCenter :one
const getCenter = `
SELECT 
  m.ModCtr, OrgParms.DefBranch_Code BrCode, m.ModAction, d.Center_Code, Center_Name, Center_Address, Center_Meet_Day, Unit, DateEstablished,
  isNull(w.CenterW_Id,-1) CenterW_Id
FROM OrgParms, Center d
LEFT JOIN CentWorker_Det w on d.Center_Code = w.Center_Code
INNER JOIN Modified m on m.UniqueKeyString1 = d.Center_Code
`

func scanRowCenter(row *sql.Row) (CenterInfo, error) {
	var i CenterInfo
	err := row.Scan(
		&i.ModCtr, &i.BrCode, &i.ModAction,
		&i.CenterCode,
		&i.CenterName,
		&i.CenterAddress,
		&i.MeetingDay,
		&i.Unit,
		&i.DateEstablished,
		&i.AOID,
	)
	return i, err
}

func scanRowsCenter(rows *sql.Rows) ([]CenterInfo, error) {
	items := []CenterInfo{}
	for rows.Next() {
		var i CenterInfo
		if err := rows.Scan(
			&i.ModCtr,
			&i.BrCode,
			&i.ModAction,
			&i.CenterCode,
			&i.CenterName,
			&i.CenterAddress,
			&i.MeetingDay,
			&i.Unit,
			&i.DateEstablished,
			&i.AOID,
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

func (q *QueriesLocal) GetCenter(ctx context.Context, centerCode string) (CenterInfo, error) {
	sql := fmt.Sprintf("%s WHERE  m.TableName = 'Center' AND Uploaded = 0 and d.Center_Code = $1", getCenter)
	row := q.db.QueryRowContext(ctx, sql, centerCode)
	return scanRowCenter(row)
}

type ListCenterParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesLocal) CenterCSV(ctx context.Context, filenamePath string) error {
	return q.Sql2Csv(ctx, `
SELECT 
  0 ModCtr, OrgParms.DefBranch_Code BrCode, d.Center_Code, Center_Name, 
  Center_Address, Center_Meet_Day, Unit, DateEstablished,
  isNull(w.CenterW_Id,-1) CenterW_Id
FROM OrgParms, Center d
LEFT JOIN CentWorker_Det w on d.Center_Code = w.Center_Code
	`, filenamePath)
}

func (q *QueriesLocal) ListCenter(ctx context.Context) ([]CenterInfo, error) {
	sql := fmt.Sprintf(
		`%v WHERE  m.TableName = 'Center' AND Uploaded = 0`,
		getCenter)
	log.Println(sql)
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsCenter(rows)
}

// -- name: UpdateCenter :one
const updateCenter = `
UPDATE Center SET 
  Center_Name = $2,
  Center_Address = $3,
  Center_Meet_Day = $4,
  Unit = $5,
  DateEstablished = $6  
WHERE Center_Code = $1;

UPDATE CentWorker_Det 
  SET CenterW_Id = 16 
  WHERE Center_Code = $1;
`

func (q *QueriesLocal) UpdateCenter(ctx context.Context, arg CenterRequest) error {
	_, err := q.db.ExecContext(ctx, updateCenter,
		arg.CenterCode,
		arg.CenterName,
		arg.CenterAddress,
		arg.MeetingDay,
		arg.Unit,
		arg.DateEstablished,
	)
	return err
}
