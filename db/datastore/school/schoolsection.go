package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createSchoolSection = `-- name: CreateSchoolSection: one
INSERT INTO School_Section (
	Uuid, Syllabus_Id, School_Id, Course_Id, Start_Date, End_Date, 
	Adviser_Id, Status_Id, Section_Name, Remarks, Other_Info
 ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
ON CONFLICT(UUID)
DO UPDATE SET
	syllabus_id =  EXCLUDED.syllabus_id,
	school_id =  EXCLUDED.school_id,
	course_id =  EXCLUDED.course_id,
	start_date =  EXCLUDED.start_date,
	end_date =  EXCLUDED.end_date,
	adviser_id =  EXCLUDED.adviser_id,
	status_id =  EXCLUDED.status_id,
	section_name =  EXCLUDED.section_name,
	remarks =  EXCLUDED.remarks,
	other_info =  EXCLUDED.other_info
 RETURNING 
   Id, Uuid, Syllabus_Id, School_Id, Course_Id, Start_Date, 
   End_Date, Adviser_Id, Status_Id, Section_Name, Remarks, Other_Info
`

type SchoolSectionRequest struct {
	Id          int64          `json:"id"`
	Uuid        uuid.UUID      `json:"uuid"`
	SyllabusId  int64          `json:"syllabusId"`
	SchoolId    int64          `json:"schoolId"`
	CourseId    int64          `json:"courseId"`
	StartDate   sql.NullTime   `json:"startDate"`
	EndDate     sql.NullTime   `json:"endDate"`
	AdviserId   sql.NullInt64  `json:"adviserId"`
	StatusId    int64          `json:"statusId"`
	SectionName sql.NullString `json:"sectionName"`
	Remarks     string         `json:"remarks"`
	OtherInfo   sql.NullString `json:"otherInfo"`
}

func (q *QueriesSchool) CreateSchoolSection(ctx context.Context, arg SchoolSectionRequest) (model.SchoolSection, error) {
	if arg.Uuid == uuid.Nil {
		arg.Uuid, _ = uuid.NewUUID()
	}
	row := q.db.QueryRowContext(ctx, createSchoolSection,
		arg.Uuid,
		arg.SyllabusId,
		arg.SchoolId,
		arg.CourseId,
		arg.StartDate,
		arg.EndDate,
		arg.AdviserId,
		arg.StatusId,
		arg.SectionName,
		arg.Remarks,
		arg.OtherInfo,
	)

	var i model.SchoolSection
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SyllabusId,
		&i.SchoolId,
		&i.CourseId,
		&i.StartDate,
		&i.EndDate,
		&i.AdviserId,
		&i.StatusId,
		&i.SectionName,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteSchoolSection = `-- name: DeleteSchoolSection :exec
DELETE FROM School_Section
WHERE id = $1
`

func (q *QueriesSchool) DeleteSchoolSection(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteSchoolSection, id)
	return err
}

type SchoolSectionInfo struct {
	Id             int64         `json:"id"`
	Uuid           uuid.UUID     `json:"uuid"`
	SyllabusId     int64         `json:"syllabusId"`
	SchoolId       int64         `json:"schoolId"`
	CourseId       int64         `json:"courseId"`
	CourseTitle    string        `json:"courseTitle"`
	CourseRefId    int64         `json:"courseRefId"`
	CourseRefTitle string        `json:"courseRefTitle"`
	StartDate      sql.NullTime  `json:"startDate"`
	EndDate        sql.NullTime  `json:"endDate"`
	AdviserId      sql.NullInt64 `json:"adviserId"`

	AdviserTitle      sql.NullString `json:"adviserTitle"`
	AdviserLastName   sql.NullString `json:"adviserLastName"`
	AdviserFirstName  sql.NullString `json:"adviserFirstName"`
	AdviserMiddleName sql.NullString `json:"adviserMiddleName"`

	StatusId    int64          `json:"statusId"`
	Status      sql.NullString `json:"status"`
	SectionName sql.NullString `json:"sectionName"`
	Remarks     string         `json:"remarks"`
	OtherInfo   sql.NullString `json:"otherInfo"`
	ModCtr      int64          `json:"modCtr"`
	Created     sql.NullTime   `json:"created"`
	Updated     sql.NullTime   `json:"updated"`
}

func populateSchoolSection(q *QueriesSchool, ctx context.Context, sql string) (SchoolSectionInfo, error) {
	var i SchoolSectionInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SyllabusId,
		&i.SchoolId,
		&i.CourseId,
		&i.CourseTitle,
		&i.CourseRefId,
		&i.CourseRefTitle,
		&i.StartDate,
		&i.EndDate,
		&i.AdviserId,
		&i.AdviserTitle,
		&i.AdviserLastName,
		&i.AdviserFirstName,
		&i.AdviserMiddleName,
		&i.StatusId,
		&i.Status,
		&i.SectionName,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateSchoolSections(q *QueriesSchool, ctx context.Context, sql string) ([]SchoolSectionInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []SchoolSectionInfo{}
	for rows.Next() {
		var i SchoolSectionInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.SyllabusId,
			&i.SchoolId,
			&i.CourseId,
			&i.CourseTitle,
			&i.CourseRefId,
			&i.CourseRefTitle,
			&i.StartDate,
			&i.EndDate,
			&i.AdviserId,
			&i.AdviserTitle,
			&i.AdviserLastName,
			&i.AdviserFirstName,
			&i.AdviserMiddleName,
			&i.StatusId,
			&i.Status,
			&i.SectionName,
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

const schoolSectionSQL = `-- name: schoolSectionSQL
SELECT 
  d.Id, mr.Uuid, Syllabus_Id, d.School_Id, Course_Id, cors.Course_Title,
  corsref.Id Course_Ref_Id, corsref.Title Course_Ref_Title, Start_Date, End_Date, 
  Adviser_Id, title.short_name, adviser.Last_Name, adviser.first_name, adviser.middle_name, 
  d.Status_Id, stat.Title Status, Section_Name, d.Remarks, d.Other_Info,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM School_Section d 
INNER JOIN Main_Record mr on mr.Uuid = d.Uuid
INNER JOIN Course cors on cors.Id = d.Course_ID
INNER JOIN Reference corsref on corsref.Id = d.Course_Id
LEFT JOIN Reference stat on stat.Id = d.Status_Id
LEFT JOIN identity_info adviser on adviser.Id = d.Adviser_Id
LEFT JOIN Reference title on title.Id = adviser.title_id
`

func (q *QueriesSchool) GetSchoolSection(ctx context.Context, id int64) (SchoolSectionInfo, error) {
	return populateSchoolSection(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", schoolSectionSQL, id))
}

func (q *QueriesSchool) GetSchoolSectionbyUuid(ctx context.Context, uuid uuid.UUID) (SchoolSectionInfo, error) {
	return populateSchoolSection(q, ctx, fmt.Sprintf("%s WHERE mr.Uuid = '%v'", schoolSectionSQL, uuid))
}

type ListSchoolSectionParams struct {
	SyllabusId int64 `json:"syllabusId"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *QueriesSchool) ListSchoolSection(ctx context.Context, arg ListSchoolSectionParams) ([]SchoolSectionInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Syllabus_Id = %v LIMIT %d OFFSET %d",
			schoolSectionSQL, arg.SyllabusId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Syllabus_Id = %v ", schoolSectionSQL, arg.SyllabusId)
	}
	return populateSchoolSections(q, ctx, sql)
}

const updateSchoolSection = `-- name: UpdateSchoolSection :one
UPDATE School_Section SET 
  Syllabus_Id = $2,
  School_Id = $3,
  Course_Id = $4,
  Start_Date = $5,
  End_Date = $6,
  Adviser_Id = $7,
  Status_Id = $8,
  Section_Name = $9,
  Remarks = $10,
  Other_Info = $11
WHERE id = $1
RETURNING 
Id, Uuid, Syllabus_Id, School_Id, Course_Id, Start_Date, End_Date, Adviser_Id, Status_Id, Section_Name, Remarks, Other_Info
`

func (q *QueriesSchool) UpdateSchoolSection(ctx context.Context, arg SchoolSectionRequest) (model.SchoolSection, error) {
	row := q.db.QueryRowContext(ctx, updateSchoolSection,
		arg.Id,
		arg.SyllabusId,
		arg.SchoolId,
		arg.CourseId,
		arg.StartDate,
		arg.EndDate,
		arg.AdviserId,
		arg.StatusId,
		arg.SectionName,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.SchoolSection
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.SyllabusId,
		&i.SchoolId,
		&i.CourseId,
		&i.StartDate,
		&i.EndDate,
		&i.AdviserId,
		&i.StatusId,
		&i.SectionName,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
