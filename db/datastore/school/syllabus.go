package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"simplebank/model"
)

const createSyllabus = `-- name: CreateSyllabus: one
INSERT INTO Syllabus (
  Uuid, Course_Id, Version, Course_Year, Semister_Id, Status_Id, Date_Implement, Remarks, Other_Info
 ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
 ON CONFLICT(UUID)
 DO UPDATE SET
	course_id =  EXCLUDED.course_id,
	version =  EXCLUDED.version,
	course_year =  EXCLUDED.course_year,
	semister_id =  EXCLUDED.semister_id,
	status_id =  EXCLUDED.status_id,
	date_implement =  EXCLUDED.date_implement,
	remarks =  EXCLUDED.remarks,
	other_info =  EXCLUDED.other_info
RETURNING 
	Id, Uuid, Course_Id, Version, Course_Year, Semister_Id, Status_Id, Date_Implement, Remarks, Other_Info
`

type SyllabusRequest struct {
	Id            int64          `json:"id"`
	Uuid          uuid.UUID      `json:"uuid"`
	CourseId      int64          `json:"courseId"`
	Version       string         `json:"version"`
	CourseYear    int32          `json:"courseYear"`
	SemisterId    int64          `json:"semisterId"`
	StatusId      int64          `json:"statusId"`
	DateImplement time.Time      `json:"dateImplement"`
	Remarks       string         `json:"remarks"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}

func (q *QueriesSchool) CreateSyllabus(ctx context.Context, arg SyllabusRequest) (model.Syllabus, error) {
	if arg.Uuid == uuid.Nil {
		arg.Uuid, _ = uuid.NewUUID()
	}
	row := q.db.QueryRowContext(ctx, createSyllabus,
		arg.Uuid,
		arg.CourseId,
		arg.Version,
		arg.CourseYear,
		arg.SemisterId,
		arg.StatusId,
		arg.DateImplement,
		arg.Remarks,
		arg.OtherInfo,
	)

	var i model.Syllabus
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CourseId,
		&i.Version,
		&i.CourseYear,
		&i.SemisterId,
		&i.StatusId,
		&i.DateImplement,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteSyllabus = `-- name: DeleteSyllabus :exec
DELETE FROM Syllabus
WHERE id = $1
`

func (q *QueriesSchool) DeleteSyllabus(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteSyllabus, id)
	return err
}

type SyllabusInfo struct {
	Id             int64          `json:"id"`
	Uuid           uuid.UUID      `json:"uuid"`
	CourseId       int64          `json:"courseId"`
	CourseTitle    string         `json:"courseTitle"`
	CourseRefId    int64          `json:"courseRefId"`
	CourseRefTitle string         `json:"courseRefTitle"`
	SchoolId       int64          `json:"schoolId"`
	Version        string         `json:"version"`
	CourseYear     int32          `json:"courseYear"`
	SemisterId     int64          `json:"semisterId"`
	Semister       string         `json:"semister"`
	StatusId       int64          `json:"statusId"`
	Status         sql.NullString `json:"status"`
	DateImplement  time.Time      `json:"dateImplement"`
	Remarks        string         `json:"remarks"`
	OtherInfo      sql.NullString `json:"otherInfo"`
	ModCtr         int64          `json:"modCtr"`
	Created        sql.NullTime   `json:"created"`
	Updated        sql.NullTime   `json:"updated"`
}

func populateSyllabus(q *QueriesSchool, ctx context.Context, sql string) (SyllabusInfo, error) {
	var i SyllabusInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CourseId,
		&i.CourseTitle,
		&i.CourseRefId,
		&i.CourseRefTitle,
		&i.SchoolId,
		&i.Version,
		&i.CourseYear,
		&i.SemisterId,
		&i.Semister,
		&i.StatusId,
		&i.Status,
		&i.DateImplement,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateSyllabuss(q *QueriesSchool, ctx context.Context, sql string) ([]SyllabusInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []SyllabusInfo{}
	for rows.Next() {
		var i SyllabusInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.CourseId,
			&i.CourseTitle,
			&i.CourseRefId,
			&i.CourseRefTitle,
			&i.SchoolId,
			&i.Version,
			&i.CourseYear,
			&i.SemisterId,
			&i.Semister,
			&i.StatusId,
			&i.Status,
			&i.DateImplement,
			&i.Remarks,
			&i.OtherInfo,
			&i.ModCtr,
			&i.Created,
			&i.Updated,
		)
		if err != nil {
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

const syllabusSQL = `-- name: syllabusSQL
SELECT 
  d.Id, mr.Uuid, d.Course_Id, corsref.Title Course_Ref_Title, 
  corsref.Id Course_Ref_Id, cors.Course_Title, cors.School_Id, d.Version, d.Course_Year, d.Semister_Id, sem.Title Semister, 
  d.Status_Id, stat.Title Status, d.Date_Implement, d.Remarks, d.Other_Info,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Syllabus d 
INNER JOIN Main_Record mr on mr.Uuid = d.Uuid
INNER JOIN Course cors on cors.Id = d.Course_ID
INNER JOIN Reference sem on sem.Id = d.Semister_Id
INNER JOIN Reference corsref on corsref.Id = cors.Course_Ref_Id
LEFT JOIN Reference stat on stat.Id = d.Status_Id
`

func (q *QueriesSchool) GetSyllabus(ctx context.Context, id int64) (SyllabusInfo, error) {
	return populateSyllabus(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", syllabusSQL, id))
}

func (q *QueriesSchool) GetSyllabusbyUuid(ctx context.Context, uuid uuid.UUID) (SyllabusInfo, error) {
	return populateSyllabus(q, ctx, fmt.Sprintf("%s WHERE mr.Uuid = '%v'", syllabusSQL, uuid))
}

type ListSyllabusParams struct {
	CourseId int64 `json:"courseId"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *QueriesSchool) ListSyllabus(ctx context.Context, arg ListSyllabusParams) ([]SyllabusInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Course_Id = %v LIMIT %d OFFSET %d",
			syllabusSQL, arg.CourseId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Course_Id = %v ", syllabusSQL, arg.CourseId)
	}
	return populateSyllabuss(q, ctx, sql)
}

const updateSyllabus = `-- name: UpdateSyllabus :one
UPDATE Syllabus SET 
  Course_Id = $2,
  Version = $3,
  Course_Year = $4,
  Semister_Id = $5,
  Status_Id = $6,
  Date_Implement = $7,
  Remarks = $8,
  Other_Info = $9
WHERE id = $1
RETURNING 
  Id, Uuid, Course_Id, Version, Course_Year, Semister_Id, Status_Id, Date_Implement, Remarks, Other_Info
`

func (q *QueriesSchool) UpdateSyllabus(ctx context.Context, arg SyllabusRequest) (model.Syllabus, error) {
	row := q.db.QueryRowContext(ctx, updateSyllabus,
		arg.Id,
		arg.CourseId,
		arg.Version,
		arg.CourseYear,
		arg.SemisterId,
		arg.StatusId,
		arg.DateImplement,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Syllabus
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CourseId,
		&i.Version,
		&i.CourseYear,
		&i.SemisterId,
		&i.StatusId,
		&i.DateImplement,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
