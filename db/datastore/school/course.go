package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createCourse = `-- name: CreateCourse: one
INSERT INTO Course (
    UUID, School_ID, Course_Title, Course_Ref_ID, Status_ID, Remarks, Other_Info
 ) VALUES ($1, $2, $3, $4, $5, $6, $7) 
 RETURNING 
   Id, Uuid, School_ID, Course_Title, Course_Ref_ID, Status_ID, Remarks, Other_Info
`

type CourseRequest struct {
	Id          int64          `json:"id"`
	Uuid        uuid.UUID      `json:"uuid"`
	SchoolId    int64          `json:"schoolID"`
	CourseTitle string         `json:"courseTitle"`
	CourseRefId int64          `json:"courseRefID"`
	StatusId    int64          `json:"statusID"`
	Remarks     string         `json:"remarks"`
	OtherInfo   sql.NullString `json:"otherInfo"`
}

func (q *QueriesSchool) CreateCourse(ctx context.Context, arg CourseRequest) (model.Course, error) {
	if arg.Uuid == uuid.Nil {
		arg.Uuid, _ = uuid.NewUUID()
	}
	row := q.db.QueryRowContext(ctx, createCourse,
		arg.Uuid,
		arg.SchoolId,
		arg.CourseTitle,
		arg.CourseRefId,
		arg.StatusId,
		arg.Remarks,
		arg.OtherInfo,
	)

	var i model.Course
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SchoolId,
		&i.CourseTitle,
		&i.CourseRefId,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteCourse = `-- name: DeleteCourse :exec
DELETE FROM Course
WHERE id = $1
`

func (q *QueriesSchool) DeleteCourse(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCourse, id)
	return err
}

type CourseInfo struct {
	Id          int64          `json:"id"`
	Uuid        uuid.UUID      `json:"uuid"`
	SchoolId    int64          `json:"schoolID"`
	CourseTitle string         `json:"courseTitle"`
	CourseRefId int64          `json:"courseRefId"`
	StatusId    int64          `json:"statusId"`
	Remarks     string         `json:"remarks"`
	OtherInfo   sql.NullString `json:"otherInfo"`
	ModCtr      int64          `json:"modCtr"`
	Created     sql.NullTime   `json:"created"`
	Updated     sql.NullTime   `json:"updated"`
}

func populateCourse(q *QueriesSchool, ctx context.Context, sql string) (CourseInfo, error) {
	var i CourseInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SchoolId,
		&i.CourseTitle,
		&i.CourseRefId,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateCourses(q *QueriesSchool, ctx context.Context, sql string) ([]CourseInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []CourseInfo{}
	for rows.Next() {
		var i CourseInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.SchoolId,
			&i.CourseTitle,
			&i.CourseRefId,
			&i.StatusId,
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

const courseSQL = `-- name: courseSQL
SELECT 
  d.Id, mr.Uuid, d.School_ID, d.Course_Title, d.Course_Ref_ID, d.Status_ID, 
  d.Remarks, d.Other_Info,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Course d 
INNER JOIN Main_Record mr on mr.Uuid = d.Uuid
`

func (q *QueriesSchool) GetCourse(ctx context.Context, id int64) (CourseInfo, error) {
	return populateCourse(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", courseSQL, id))
}

func (q *QueriesSchool) GetCoursebyUuid(ctx context.Context, uuid uuid.UUID) (CourseInfo, error) {
	return populateCourse(q, ctx, fmt.Sprintf("%s WHERE mr.Uuid = '%v'", courseSQL, uuid))
}

type ListCourseParams struct {
	SchoolId int64 `json:"schoolId"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *QueriesSchool) ListCourse(ctx context.Context, arg ListCourseParams) ([]CourseInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.School_Id = %v LIMIT %d OFFSET %d",
			courseSQL, arg.SchoolId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.School_Id = %v ", courseSQL, arg.SchoolId)
	}
	return populateCourses(q, ctx, sql)
}

const updateCourse = `-- name: UpdateCourse :one
UPDATE Course SET 
  School_ID = $2,
  Course_Title = $3,
  Course_Ref_ID = $4,
  Status_ID = $5,
  Remarks = $6,
  Other_Info = $7
WHERE id = $1
RETURNING 
  Id, Uuid, School_ID, Course_Title, Course_Ref_ID, Status_ID, Remarks, Other_Info
`

func (q *QueriesSchool) UpdateCourse(ctx context.Context, arg CourseRequest) (model.Course, error) {
	row := q.db.QueryRowContext(ctx, updateCourse,
		arg.Id,
		arg.SchoolId,
		arg.CourseTitle,
		arg.CourseRefId,
		arg.StatusId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.Course
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SchoolId,
		&i.CourseTitle,
		&i.CourseRefId,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
