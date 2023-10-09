package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
)

const createSectionSubject = `-- name: CreateSectionSubject: one
INSERT INTO Section_Subject(
   uuid, school_section_id, subject_id, teacher_id, 
   type_id, status_id, schedule_code, schedule_json, 
   remarks, other_info )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT (uuid)
DO UPDATE SET
	school_section_id = excluded.school_section_id,
	subject_id = excluded.subject_id,
	teacher_id = excluded.teacher_id,
	type_id = excluded.type_id,
	status_id = excluded.status_id,
	schedule_code = excluded.schedule_code,
	schedule_json = excluded.schedule_json,
	remarks = excluded.remarks,
	other_info = excluded.other_info
RETURNING id, uuid, school_section_id, subject_id, teacher_id, 
type_id, status_id, schedule_code, schedule_json, 
remarks, other_info`

type SectionSubjectRequest struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	SchoolSectionId int64          `json:"schoolSectionId"`
	SubjectId       int64          `json:"subjectId"`
	TeacherId       sql.NullInt64  `json:"teacherId"`
	TypeId          int64          `json:"typeId"`
	StatusId        int64          `json:"statusId"`
	ScheduleCode    sql.NullString `json:"scheduleCode"`
	ScheduleJson    sql.NullString `json:"scheduleJson"`
	Remarks         string         `json:"remarks"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}

func (q *QueriesSchool) CreateSectionSubject(ctx context.Context, arg SectionSubjectRequest) (model.SectionSubject, error) {
	row := q.db.QueryRowContext(ctx, createSectionSubject,
		arg.Uuid,
		arg.SchoolSectionId,
		arg.SubjectId,
		arg.TeacherId,
		arg.TypeId,
		arg.StatusId,
		arg.ScheduleCode,
		arg.ScheduleJson,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.SectionSubject
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SchoolSectionId,
		&i.SubjectId,
		&i.TeacherId,
		&i.TypeId,
		&i.StatusId,
		&i.ScheduleCode,
		&i.ScheduleJson,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteSectionSubject = `-- name: DeleteSectionSubject :exec
DELETE FROM Section_Subject
WHERE uuid = $1
`

func (q *QueriesSchool) DeleteSectionSubject(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteSectionSubject, uuid)
	return err
}

type SectionSubjectInfo struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	SchoolSectionId int64          `json:"schoolSectionId"`
	SubjectId       int64          `json:"subjectId"`
	TeacherId       sql.NullInt64  `json:"teacherId"`
	TypeId          int64          `json:"typeId"`
	StatusId        int64          `json:"statusId"`
	ScheduleCode    sql.NullString `json:"scheduleCode"`
	ScheduleJson    sql.NullString `json:"scheduleJson"`
	Remarks         string         `json:"remarks"`
	OtherInfo       sql.NullString `json:"otherInfo"`
	ModCtr          int64          `json:"modCtr"`
	Created         sql.NullTime   `json:"created"`
	Updated         sql.NullTime   `json:"updated"`
}

const sectionSubjectSQL = `-- name: SectionSubjectSQL :one
SELECT
d.Id, mr.UUID, school_section_id, subject_id, teacher_id, type_id, status_id, schedule_code, schedule_json, remarks, other_info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Section_Subject d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateSectionSubject(q *QueriesSchool, ctx context.Context, sql string) (SectionSubjectInfo, error) {
	var i SectionSubjectInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SchoolSectionId,
		&i.SubjectId,
		&i.TeacherId,
		&i.TypeId,
		&i.StatusId,
		&i.ScheduleCode,
		&i.ScheduleJson,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateSectionSubjects(q *QueriesSchool, ctx context.Context, sql string) ([]SectionSubjectInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SectionSubjectInfo{}
	for rows.Next() {
		var i SectionSubjectInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.SchoolSectionId,
			&i.SubjectId,
			&i.TeacherId,
			&i.TypeId,
			&i.StatusId,
			&i.ScheduleCode,
			&i.ScheduleJson,
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

func (q *QueriesSchool) GetSectionSubject(ctx context.Context, id int64) (SectionSubjectInfo, error) {
	return populateSectionSubject(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", sectionSubjectSQL, id))
}

func (q *QueriesSchool) GetSectionSubjectbyUuid(ctx context.Context, uuid uuid.UUID) (SectionSubjectInfo, error) {
	return populateSectionSubject(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", sectionSubjectSQL, uuid))
}

type ListSectionSubjectParams struct {
	SchoolSectionId int64 `json:"schoolSectionId"`
	Limit           int32 `json:"limit"`
	Offset          int32 `json:"offset"`
}

func (q *QueriesSchool) ListSectionSubject(ctx context.Context, arg ListSectionSubjectParams) ([]SectionSubjectInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			sectionSubjectSQL, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf(sectionSubjectSQL)
	}
	return populateSectionSubjects(q, ctx, sql)
}

const updateSectionSubject = `-- name: UpdateSectionSubject :one
UPDATE Section_Subject SET 
uuid = $2,
school_section_id = $3,
subject_id = $4,
teacher_id = $5,
type_id = $6,
status_id = $7,
schedule_code = $8,
schedule_json = $9,
remarks = $10,
other_info = $11
WHERE id = $1
RETURNING id, uuid, school_section_id, subject_id, teacher_id, type_id, status_id, schedule_code, schedule_json, remarks, other_info
`

func (q *QueriesSchool) UpdateSectionSubject(ctx context.Context, arg SectionSubjectRequest) (model.SectionSubject, error) {
	row := q.db.QueryRowContext(ctx, updateSectionSubject,
		arg.Id,
		arg.Uuid,
		arg.SchoolSectionId,
		arg.SubjectId,
		arg.TeacherId,
		arg.TypeId,
		arg.StatusId,
		arg.ScheduleCode,
		arg.ScheduleJson,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.SectionSubject
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SchoolSectionId,
		&i.SubjectId,
		&i.TeacherId,
		&i.TypeId,
		&i.StatusId,
		&i.ScheduleCode,
		&i.ScheduleJson,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
