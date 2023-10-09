package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createEducational = `-- name: CreateEducational: one
INSERT INTO Educational (
IIId, Series, Level_Id, Course_Type_Id, Course_Id, 
Course, School, Address_Detail, Address_URL, Geography_Id, Start_Date, 
End_Date, Period_Date, Completed, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, 
$13, $14, $15
) 
ON CONFLICT(IIID, Series) DO UPDATE SET 
	level_id = excluded.level_id,
	course_type_id = excluded.course_type_id,
	course_id = excluded.course_id,
	course = excluded.course,
	school = excluded.school,
	address_detail = excluded.address_detail,
	address_url = excluded.address_url,
	geography_id = excluded.geography_id,
	start_date = excluded.start_date,
	end_date = excluded.end_date,
	period_date = excluded.period_date,
	completed = excluded.completed,
	other_info = excluded.other_info
	
RETURNING UUId, IIId, Series, Level_Id, Course_Type_Id, Course_Id, 
Course, School, Address_Detail, Address_URL, Geography_Id, Start_Date, 
End_Date, Period_Date, Completed, Other_Info

`

type EducationalRequest struct {
	Uuid          uuid.UUID      `json:"uuid"`
	Iiid          int64          `json:"iiid"`
	Series        int16          `json:"series"`
	LevelId       sql.NullInt64  `json:"levelId"`
	CourseTypeId  sql.NullInt64  `json:"courseTypeId"`
	CourseId      sql.NullInt64  `json:"courseId"`
	Course        string         `json:"course"`
	School        string         `json:"school"`
	AddressDetail sql.NullString `json:"addressDetail"`
	AddressUrl    sql.NullString `json:"addressUrl"`
	GeographyId   sql.NullInt64  `json:"geographyId"`
	StartDate     sql.NullTime   `json:"startDate"`
	EndDate       sql.NullTime   `json:"endDate"`
	PeriodDate    sql.NullString `json:"periodDate"`
	Completed     sql.NullBool   `json:"completed"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}

func (q *QueriesIdentity) CreateEducational(ctx context.Context, arg EducationalRequest) (model.Educational, error) {
	row := q.db.QueryRowContext(ctx, createEducational,
		arg.Iiid,
		arg.Series,
		arg.LevelId,
		arg.CourseTypeId,
		arg.CourseId,
		arg.Course,
		arg.School,
		arg.AddressDetail,
		arg.AddressUrl,
		arg.GeographyId,
		arg.StartDate,
		arg.EndDate,
		arg.PeriodDate,
		arg.Completed,
		arg.OtherInfo,
	)
	var i model.Educational
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.LevelId,
		&i.CourseTypeId,
		&i.CourseId,
		&i.Course,
		&i.School,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.StartDate,
		&i.EndDate,
		&i.PeriodDate,
		&i.Completed,
		&i.OtherInfo,
	)
	return i, err
}

const deleteEducational = `-- name: DeleteEducational :exec
DELETE FROM Educational
WHERE uuid = $1
`

func (q *QueriesIdentity) DeleteEducational(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteEducational, uuid)
	return err
}

type EducationalInfo struct {
	Uuid          uuid.UUID      `json:"uuid"`
	Iiid          int64          `json:"iiid"`
	Series        int16          `json:"series"`
	LevelId       sql.NullInt64  `json:"levelId"`
	CourseTypeId  sql.NullInt64  `json:"courseTypeId"`
	CourseId      sql.NullInt64  `json:"courseId"`
	Course        string         `json:"course"`
	School        string         `json:"school"`
	AddressDetail sql.NullString `json:"addressDetail"`
	AddressUrl    sql.NullString `json:"addressUrl"`
	GeographyId   sql.NullInt64  `json:"geographyId"`
	StartDate     sql.NullTime   `json:"startDate"`
	EndDate       sql.NullTime   `json:"endDate"`
	PeriodDate    sql.NullString `json:"periodDate"`
	Completed     sql.NullBool   `json:"completed"`
	OtherInfo     sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getEducational = `-- name: GetEducational :one
SELECT 
mr.UUId, IIId, Series, Level_Id, Course_Type_Id, 
Course_Id, Course, School, Address_Detail, Address_URL, Geography_Id, 
Start_Date, End_Date, Period_Date, Completed, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Educational d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesIdentity) GetEducational(ctx context.Context, uuid uuid.UUID) (EducationalInfo, error) {
	row := q.db.QueryRowContext(ctx, getEducational, uuid)
	var i EducationalInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.LevelId,
		&i.CourseTypeId,
		&i.CourseId,
		&i.Course,
		&i.School,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.StartDate,
		&i.EndDate,
		&i.PeriodDate,
		&i.Completed,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getEducationalbyUuId = `-- name: GetEducationalbyUuId :one
SELECT 
mr.UUId, IIId, Series, Level_Id, Course_Type_Id, 
Course_Id, Course, School, Address_Detail, Address_URL, Geography_Id, 
Start_Date, End_Date, Period_Date, Completed, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Educational d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesIdentity) GetEducationalbyUuId(ctx context.Context, uuid uuid.UUID) (EducationalInfo, error) {
	row := q.db.QueryRowContext(ctx, getEducationalbyUuId, uuid)
	var i EducationalInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.LevelId,
		&i.CourseTypeId,
		&i.CourseId,
		&i.Course,
		&i.School,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.StartDate,
		&i.EndDate,
		&i.PeriodDate,
		&i.Completed,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listEducational = `-- name: ListEducational:many
SELECT 
mr.UUId, IIId, Series, Level_Id, Course_Type_Id, 
Course_Id, Course, School, Address_Detail, Address_URL, Geography_Id, 
Start_Date, End_Date, Period_Date, Completed, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Educational d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Iiid = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListEducationalParams struct {
	Iiid   int64 `json:"iiid"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesIdentity) ListEducational(ctx context.Context, arg ListEducationalParams) ([]EducationalInfo, error) {
	rows, err := q.db.QueryContext(ctx, listEducational, arg.Iiid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []EducationalInfo{}
	for rows.Next() {
		var i EducationalInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.Iiid,
			&i.Series,
			&i.LevelId,
			&i.CourseTypeId,
			&i.CourseId,
			&i.Course,
			&i.School,
			&i.AddressDetail,
			&i.AddressUrl,
			&i.GeographyId,
			&i.StartDate,
			&i.EndDate,
			&i.PeriodDate,
			&i.Completed,
			&i.OtherInfo,

			&i.ModCtr,
			&i.Created,
			&i.Updated,
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

const updateEducational = `-- name: UpdateEducational :one
UPDATE Educational SET 
IIId = $2,
Series = $3,
Level_Id = $4,
Course_Type_Id = $5,
Course_Id = $6,
Course = $7,
School = $8,
Address_Detail = $9,
Address_URL = $10,
Geography_Id = $11,
Start_Date = $12,
End_Date = $13,
Period_Date = $14,
Completed = $15,
Other_Info = $16
WHERE uuid = $1
RETURNING UUId, IIId, Series, Level_Id, Course_Type_Id, Course_Id, 
Course, School, Address_Detail, Address_URL, Geography_Id, Start_Date, 
End_Date, Period_Date, Completed, Other_Info
`

func (q *QueriesIdentity) UpdateEducational(ctx context.Context, arg EducationalRequest) (model.Educational, error) {
	row := q.db.QueryRowContext(ctx, updateEducational,

		arg.Uuid,
		arg.Iiid,
		arg.Series,
		arg.LevelId,
		arg.CourseTypeId,
		arg.CourseId,
		arg.Course,
		arg.School,
		arg.AddressDetail,
		arg.AddressUrl,
		arg.GeographyId,
		arg.StartDate,
		arg.EndDate,
		arg.PeriodDate,
		arg.Completed,
		arg.OtherInfo,
	)
	var i model.Educational
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.LevelId,
		&i.CourseTypeId,
		&i.CourseId,
		&i.Course,
		&i.School,
		&i.AddressDetail,
		&i.AddressUrl,
		&i.GeographyId,
		&i.StartDate,
		&i.EndDate,
		&i.PeriodDate,
		&i.Completed,
		&i.OtherInfo,
	)
	return i, err
}
