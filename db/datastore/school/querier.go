package db

import (
	"context"
	"database/sql"
	"fmt"
	dsRef "simplebank/db/datastore/reference"
	dsTrn "simplebank/db/datastore/transaction"
	"simplebank/model"

	"github.com/google/uuid"
)

type QuerierSchool interface {
	CreateCourse(ctx context.Context, arg CourseRequest) (model.Course, error)
	GetCourse(ctx context.Context, id int64) (CourseInfo, error)
	GetCoursebyUuid(ctx context.Context, uuid uuid.UUID) (CourseInfo, error)
	ListCourse(ctx context.Context, arg ListCourseParams) ([]CourseInfo, error)
	UpdateCourse(ctx context.Context, arg CourseRequest) (model.Course, error)
	DeleteCourse(ctx context.Context, id int64) error

	CreateSubject(ctx context.Context, arg SubjectRequest) (model.Subject, error)
	GetSubject(ctx context.Context, id int64) (SubjectInfo, error)
	GetSubjectbyUuid(ctx context.Context, uuid uuid.UUID) (SubjectInfo, error)
	ListSubject(ctx context.Context, arg ListSubjectParams) ([]SubjectInfo, error)
	UpdateSubject(ctx context.Context, arg SubjectRequest) (model.Subject, error)
	DeleteSubject(ctx context.Context, id int64) error

	CreateSyllabus(ctx context.Context, arg SyllabusRequest) (model.Syllabus, error)
	GetSyllabus(ctx context.Context, id int64) (SyllabusInfo, error)
	GetSyllabusbyUuid(ctx context.Context, uuid uuid.UUID) (SyllabusInfo, error)
	ListSyllabus(ctx context.Context, arg ListSyllabusParams) ([]SyllabusInfo, error)
	UpdateSyllabus(ctx context.Context, arg SyllabusRequest) (model.Syllabus, error)
	DeleteSyllabus(ctx context.Context, id int64) error

	CreateSchoolSection(ctx context.Context, arg SchoolSectionRequest) (model.SchoolSection, error)
	GetSchoolSection(ctx context.Context, id int64) (SchoolSectionInfo, error)
	GetSchoolSectionbyUuid(ctx context.Context, uuid uuid.UUID) (SchoolSectionInfo, error)
	ListSchoolSection(ctx context.Context, arg ListSchoolSectionParams) ([]SchoolSectionInfo, error)
	UpdateSchoolSection(ctx context.Context, arg SchoolSectionRequest) (model.SchoolSection, error)
	DeleteSchoolSection(ctx context.Context, id int64) error

	CreateSectionSubject(ctx context.Context, arg SectionSubjectRequest) (model.SectionSubject, error)
	GetSectionSubject(ctx context.Context, id int64) (SectionSubjectInfo, error)
	GetSectionSubjectbyUuid(ctx context.Context, uuid uuid.UUID) (SectionSubjectInfo, error)
	ListSectionSubject(ctx context.Context, arg ListSectionSubjectParams) ([]SectionSubjectInfo, error)
	UpdateSectionSubject(ctx context.Context, arg SectionSubjectRequest) (model.SectionSubject, error)
	DeleteSectionSubject(ctx context.Context, uuid uuid.UUID) error

	CreateAccountSubject(ctx context.Context, arg AccountSubjectRequest) (model.AccountSubject, error)
	GetAccountSubject(ctx context.Context, id int64) (AccountSubjectInfo, error)
	GetAccountSubjectbyUuid(ctx context.Context, uuid uuid.UUID) (AccountSubjectInfo, error)
	ListAccountSubject(ctx context.Context, arg ListAccountSubjectParams) ([]AccountSubjectInfo, error)
	UpdateAccountSubject(ctx context.Context, arg AccountSubjectRequest) (model.AccountSubject, error)
	DeleteAccountSubject(ctx context.Context, uuid uuid.UUID) error

	CreateSubjectEvent(ctx context.Context, arg SubjectEventRequest) (model.SubjectEvent, error)
	GetSubjectEvent(ctx context.Context, id int64) (SubjectEventInfo, error)
	GetSubjectEventbyUuid(ctx context.Context, uuid uuid.UUID) (SubjectEventInfo, error)
	ListSubjectEvent(ctx context.Context, arg ListSubjectEventParams) ([]SubjectEventInfo, error)
	UpdateSubjectEvent(ctx context.Context, arg SubjectEventRequest) (model.SubjectEvent, error)
	DeleteSubjectEvent(ctx context.Context, uuid uuid.UUID) error

	CreateSyllabusSubject(ctx context.Context, arg SyllabusSubjectRequest) (model.SyllabusSubject, error)
	GetSyllabusSubject(ctx context.Context, id int64) (SyllabusSubjectInfo, error)
	GetSyllabusSubjectbyUuid(ctx context.Context, uuid uuid.UUID) (SyllabusSubjectInfo, error)
	ListSyllabusSubject(ctx context.Context, arg ListSyllabusSubjectParams) ([]SyllabusSubjectInfo, error)
	UpdateSyllabusSubject(ctx context.Context, arg SyllabusSubjectRequest) (model.SyllabusSubject, error)
	DeleteSyllabusSubject(ctx context.Context, uuid uuid.UUID) error

	CreateAnswer(ctx context.Context, arg AnswerRequest) (model.Answer, error)
	GetAnswer(ctx context.Context, id int64) (AnswerInfo, error)
	GetAnswerbyUuid(ctx context.Context, uuid uuid.UUID) (AnswerInfo, error)
	ListAnswer(ctx context.Context, arg ListAnswerParams) ([]AnswerInfo, error)
	UpdateAnswer(ctx context.Context, arg AnswerRequest) (model.Answer, error)
	DeleteAnswer(ctx context.Context, uuid uuid.UUID) error

	CreateQuestionaire(ctx context.Context, arg QuestionaireRequest) (model.Questionaire, error)
	GetQuestionaire(ctx context.Context, id int64) (QuestionaireInfo, error)
	GetQuestionairebyUuid(ctx context.Context, uuid uuid.UUID) (QuestionaireInfo, error)
	ListQuestionaire(ctx context.Context, arg ListQuestionaireParams) ([]QuestionaireInfo, error)
	UpdateQuestionaire(ctx context.Context, arg QuestionaireRequest) (model.Questionaire, error)
	DeleteQuestionaire(ctx context.Context, uuid uuid.UUID) error

	CreateQuestion(ctx context.Context, arg QuestionRequest) (model.Question, error)
	GetQuestion(ctx context.Context, id int64) (QuestionInfo, error)
	GetQuestionbyUuid(ctx context.Context, uuid uuid.UUID) (QuestionInfo, error)
	ListQuestion(ctx context.Context, arg ListQuestionParams) ([]QuestionInfo, error)
	UpdateQuestion(ctx context.Context, arg QuestionRequest) (model.Question, error)
	DeleteQuestion(ctx context.Context, uuid uuid.UUID) error
}

var _ QuerierSchool = (*QueriesSchool)(nil)

var _ StoreSchool = (*SQLStoreSchool)(nil)

func (q *QueriesSchool) WithTx(tx *sql.Tx) *QueriesSchool {
	return &QueriesSchool{
		db: tx,
	}
}

// SQLStore provides all functions to execute SQL queriesUser and users
type SQLStoreSchool struct {
	db *sql.DB
	*QueriesSchool
	*dsRef.QueriesReference
	*dsTrn.QueriesTransaction
}

// Store defines all functions to execute db queriesUser and users
type StoreSchool interface {
	QuerierSchool
	dsRef.QuerierReference
	dsTrn.QuerierTransaction
}

// NewStore creates a new store
func NewStoreSchool(db *sql.DB) StoreSchool {
	return &SQLStoreSchool{
		QueriesSchool:      New(db),
		QueriesReference:   dsRef.New(db),
		QueriesTransaction: dsTrn.New(db),
		db:                 db,
	}
}

// ExecTx executes a function within a database transaction
func (store *SQLStoreSchool) ExecTx(ctx context.Context,
	fn func(*QueriesSchool) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
