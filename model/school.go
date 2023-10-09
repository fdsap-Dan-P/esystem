package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Course struct {
	Id          int64          `json:"id"`
	Uuid        uuid.UUID      `json:"uuid"`
	SchoolId    int64          `json:"schoolId"`
	CourseTitle string         `json:"courseTitle"`
	CourseRefId int64          `json:"courseRefId"`
	StatusId    int64          `json:"statusId"`
	Remarks     string         `json:"remarks"`
	OtherInfo   sql.NullString `json:"otherInfo"`
}

type Subject struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	SubjectTitle string         `json:"subjectTitle"`
	SubjectRefId int64          `json:"subjectRefId"`
	TypeId       sql.NullInt64  `json:"typeId"`
	Remarks      string         `json:"remarks"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}

type Syllabus struct {
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

type SchoolSection struct {
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

type SectionSubject struct {
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

type AccountSubject struct {
	Id               int64           `json:"id"`
	Uuid             uuid.UUID       `json:"uuid"`
	AccountId        int64           `json:"accountId"`
	SectionSubjectId int64           `json:"sectionSubjectId"`
	SubjectId        int64           `json:"subjectId"`
	Ratings1stQtr    decimal.Decimal `json:"ratings1stQtr"`
	Ratings2ndQtr    decimal.Decimal `json:"ratings2ndQtr"`
	Ratings3rdQtr    decimal.Decimal `json:"ratings3rdQtr"`
	Ratings4thQtr    decimal.Decimal `json:"ratings4thQtr"`
	RatingsFinal     decimal.Decimal `json:"ratingsFinal"`
	AttendanceCtr    int64           `json:"attendanceCtr"`
	AbsentCtr        int64           `json:"absentCtr"`
	LateCtr          int64           `json:"lateCtr"`
	StatusId         int64           `json:"statusId"`
	Remarks          string          `json:"remarks"`
	OtherInfo        sql.NullString  `json:"otherInfo"`
}

type SubjectEvent struct {
	Id               int64           `json:"id"`
	Uuid             uuid.UUID       `json:"uuid"`
	TypeId           int64           `json:"typeId"`
	TicketItemId     int64           `json:"ticketItemId"`
	Iiid             int64           `json:"iiid"`
	SectionSubjectId int64           `json:"sectionSubjectId"`
	EventDate        time.Time       `json:"eventDate"`
	GradingPeriod    int16           `json:"gradingPeriod"`
	ItemCount        decimal.Decimal `json:"itemCount"`
	StatusId         int64           `json:"statusId"`
	Remarks          sql.NullString  `json:"remarks"`
	OtherInfo        sql.NullString  `json:"otherInfo"`
}

type SyllabusSubject struct {
	Id         int64          `json:"id"`
	Uuid       uuid.UUID      `json:"uuid"`
	SyllabusId int64          `json:"syllabusId"`
	SubjectId  int64          `json:"subjectId"`
	Units      int64          `json:"units"`
	TypeId     int64          `json:"typeId"`
	StatusId   int64          `json:"statusId"`
	Remarks    string         `json:"remarks"`
	OtherInfo  sql.NullString `json:"otherInfo"`
}

type Questionnaire struct {
	Id              int64          `json:"id"`
	Uuid            uuid.UUID      `json:"uuid"`
	Code            string         `json:"code"`
	Version         int64          `json:"version"`
	Title           string         `json:"title"`
	TypeId          int64          `json:"typeId"`
	SubjectId       sql.NullInt64  `json:"subjectId"`
	DateRevised     time.Time      `json:"dateRevised"`
	OfficeId        sql.NullInt64  `json:"officeId"`
	AuthorId        sql.NullInt64  `json:"authorId"`
	StatusId        int64          `json:"statusId"`
	PointEquivalent sql.NullString `json:"pointEquivalent"`
	Remarks         sql.NullString `json:"remarks"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}

type Question struct {
	Id             int64          `json:"id"`
	Uuid           uuid.UUID      `json:"uuid"`
	QuestionaireId int64          `json:"questionaireId"`
	Series         int16          `json:"series"`
	TypeId         int64          `json:"typeId"`
	QuestionItem   string         `json:"questionItem"`
	Choices        sql.NullString `json:"choices"`
	AnswerType     sql.NullString `json:"answerType"`
	ParentId       sql.NullInt64  `json:"parentId"`
	StatusId       int64          `json:"statusId"`
	Remarks        sql.NullString `json:"remarks"`
	OtherInfo      sql.NullString `json:"otherInfo"`
}

type Answer struct {
	Id             int64           `json:"id"`
	Uuid           uuid.UUID       `json:"uuid"`
	SubjectEventId int64           `json:"subjectEventId"`
	QuestionId     int64           `json:"questionId"`
	Answers        sql.NullString  `json:"answers"`
	Points         decimal.Decimal `json:"points"`
	Remarks        sql.NullString  `json:"remarks"`
	OtherInfo      sql.NullString  `json:"otherInfo"`
}
