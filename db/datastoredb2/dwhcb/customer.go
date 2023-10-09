package dwhcb

import (
	"context"
	"database/sql"
	"fmt"
	// "log"
)

type CustomerInfo struct {
	CUSTOMER_ID        sql.NullInt64  `json:"CUSTOMER_ID"`
	CO_CODE            sql.NullString `json:"CO_CODE"`
	COMPANY_BOOK       sql.NullString `json:"COMPANY_BOOK"`
	CUSTOMER_NAME      sql.NullString `json:"CUSTOMER_NAME"`
	L_DATE_RECOG       sql.NullTime   `json:"L_DATE_RECOG"`
	DATE_OF_BIRTH      sql.NullTime   `json:"DATE_OF_BIRTH"`
	ADDRESS            sql.NullString `json:"ADDRESS"`
	TOWN_COUNTRY       sql.NullString `json:"TOWN_COUNTRY"`
	MEMBER_FLAG        sql.NullString `json:"MEMBER_FLAG"`
	SECTOR_ID          sql.NullInt64  `json:"SECTOR_ID"`
	INDUSTRY_ID        sql.NullInt64  `json:"INDUSTRY_ID"`
	CUSTOMER_STATUS_ID sql.NullInt64  `json:"CUSTOMER_STATUS_ID"`
	NATIONALITY        sql.NullString `json:"NATIONALITY"`
	LEGAL_ID           sql.NullString `json:"LEGAL_ID"`
	LEGAL_DOC_NAME     sql.NullString `json:"LEGAL_DOC_NAME"`
	MARITAL_STATUS     sql.NullString `json:"MARITAL_STATUS"`
	RISK_CLASS_ID      sql.NullString `json:"RISK_CLASS_ID"`
	L_BARANGAY_CODE    sql.NullString `json:"L_BARANGAY_CODE"`
	L_PROVINCE         sql.NullString `json:"L_PROVINCE"`
	COMPANY_ID         sql.NullString `json:"COMPANY_ID"`
	CENTER_ID          sql.NullString `json:"CENTER_ID"`
	AGRARIAN_FLAG      sql.NullString `json:"AGRARIAN_FLAG"`
	RURAL_FLAG         sql.NullString `json:"RURAL_FLAG"`
	SMS_1              sql.NullString `json:"SMS_1"`
	L_SHARE_AMOUNT     sql.NullString `json:"L_SHARE_AMOUNT"`
	L_PPI_SCORE        sql.NullInt64  `json:"L_PPI_SCORE"`
	LAST_PAYMENT_DATE  sql.NullTime   `json:"LAST_PAYMENT_DATE"`
	L_MEMBER_CLASS     sql.NullString `json:"L_MEMBER_CLASS"`
	L_GLIP_FLAG        sql.NullString `json:"L_GLIP_FLAG"`
	ALT_CUST_ID        sql.NullString `json:"ALT_CUST_ID"`
	GIVEN_NAMES        sql.NullString `json:"GIVEN_NAMES"`
	L_MIDDLE_NAME      sql.NullString `json:"L_MIDDLE_NAME"`
	FAMILY_NAME        sql.NullString `json:"FAMILY_NAME"`
	L_BARANGAY_NAME    sql.NullString `json:"L_BARANGAY_NAME"`
	GENDER             sql.NullString `json:"GENDER"`
	ZIPCODE            sql.NullString `json:"ZIPCODE"`
	L_TIN_ID           sql.NullString `json:"L_TIN_ID"`
	L_SSS_ID           sql.NullString `json:"L_SSS_ID"`
	L_PHILHEALTH_ID    sql.NullString `json:"L_PHILHEALTH_ID"`
	TITLE              sql.NullString `json:"TITLE"`
	L_REL_DOSRI        sql.NullString `json:"L_REL_DOSRI"`
	L_OCCUPATION       sql.NullString `json:"L_OCCUPATION"`
	L_EDUCATION        sql.NullString `json:"L_EDUCATION"`
	L_DISABL_FLAG      sql.NullString `json:"L_DISABL_FLAG"`
	SPOUSE_NAME        sql.NullString `json:"SPOUSE_NAME"`
	MBA_WEEKLY_AMOUNT  sql.NullString `json:"MBA_WEEKLY_AMOUNT"`
	RESIGNED_REMARK    sql.NullString `json:"RESIGNED_REMARK"`
	L_GLIP_MBA_AMT     sql.NullString `json:"L_GLIP_MBA_AMT"`
	L_RELIGION         sql.NullString `json:"L_RELIGION"`
}

func populateData(q *QueriesCustomer, ctx context.Context, sql string) (CustomerInfo, error) {
	var i CustomerInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.CUSTOMER_ID,
		&i.CO_CODE,
		&i.COMPANY_BOOK,
		&i.CUSTOMER_NAME,
		&i.L_DATE_RECOG,
		&i.DATE_OF_BIRTH,
		&i.ADDRESS,
		&i.TOWN_COUNTRY,
		&i.MEMBER_FLAG,
		&i.SECTOR_ID,
		&i.INDUSTRY_ID,
		&i.CUSTOMER_STATUS_ID,
		&i.NATIONALITY,
		&i.LEGAL_ID,
		&i.LEGAL_DOC_NAME,
		&i.MARITAL_STATUS,
		&i.RISK_CLASS_ID,
		&i.L_BARANGAY_CODE,
		&i.L_PROVINCE,
		&i.COMPANY_ID,
		&i.CENTER_ID,
		&i.AGRARIAN_FLAG,
		&i.RURAL_FLAG,
		&i.SMS_1,
		&i.L_SHARE_AMOUNT,
		&i.L_PPI_SCORE,
		&i.LAST_PAYMENT_DATE,
		&i.L_MEMBER_CLASS,
		&i.L_GLIP_FLAG,
		&i.ALT_CUST_ID,
		&i.GIVEN_NAMES,
		&i.L_MIDDLE_NAME,
		&i.FAMILY_NAME,
		&i.L_BARANGAY_NAME,
		&i.GENDER,
		&i.ZIPCODE,
		&i.L_TIN_ID,
		&i.L_SSS_ID,
		&i.L_PHILHEALTH_ID,
		&i.TITLE,
		&i.L_REL_DOSRI,
		&i.L_OCCUPATION,
		&i.L_EDUCATION,
		&i.L_DISABL_FLAG,
		&i.SPOUSE_NAME,
		&i.MBA_WEEKLY_AMOUNT,
		&i.RESIGNED_REMARK,
		&i.L_GLIP_MBA_AMT,
		&i.L_RELIGION,
	)
	return i, err
}

func populateDatas(q *QueriesCustomer, ctx context.Context, sql string) ([]CustomerInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []CustomerInfo{}
	for rows.Next() {
		var i CustomerInfo
		err := rows.Scan(
			&i.CUSTOMER_ID,
			&i.CO_CODE,
			&i.COMPANY_BOOK,
			&i.CUSTOMER_NAME,
			&i.L_DATE_RECOG,
			&i.DATE_OF_BIRTH,
			&i.ADDRESS,
			&i.TOWN_COUNTRY,
			&i.MEMBER_FLAG,
			&i.SECTOR_ID,
			&i.INDUSTRY_ID,
			&i.CUSTOMER_STATUS_ID,
			&i.NATIONALITY,
			&i.LEGAL_ID,
			&i.LEGAL_DOC_NAME,
			&i.MARITAL_STATUS,
			&i.RISK_CLASS_ID,
			&i.L_BARANGAY_CODE,
			&i.L_PROVINCE,
			&i.COMPANY_ID,
			&i.CENTER_ID,
			&i.AGRARIAN_FLAG,
			&i.RURAL_FLAG,
			&i.SMS_1,
			&i.L_SHARE_AMOUNT,
			&i.L_PPI_SCORE,
			&i.LAST_PAYMENT_DATE,
			&i.L_MEMBER_CLASS,
			&i.L_GLIP_FLAG,
			&i.ALT_CUST_ID,
			&i.GIVEN_NAMES,
			&i.L_MIDDLE_NAME,
			&i.FAMILY_NAME,
			&i.L_BARANGAY_NAME,
			&i.GENDER,
			&i.ZIPCODE,
			&i.L_TIN_ID,
			&i.L_SSS_ID,
			&i.L_PHILHEALTH_ID,
			&i.TITLE,
			&i.L_REL_DOSRI,
			&i.L_OCCUPATION,
			&i.L_EDUCATION,
			&i.L_DISABL_FLAG,
			&i.SPOUSE_NAME,
			&i.MBA_WEEKLY_AMOUNT,
			&i.RESIGNED_REMARK,
			&i.L_GLIP_MBA_AMT,
			&i.L_RELIGION,
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

const selectSQL = `-- name: selectSQL
SELECT 
  c.CUSTOMER_ID, CO_CODE, COMPANY_BOOK, CUSTOMER_NAME, L_DATE_RECOG, DATE_OF_BIRTH, 
  ADDRESS, TOWN_COUNTRY, MEMBER_FLAG, SECTOR_ID, INDUSTRY_ID, CUSTOMER_STATUS_ID, 
  NATIONALITY, LEGAL_ID, LEGAL_DOC_NAME, MARITAL_STATUS, RISK_CLASS_ID, L_BARANGAY_CODE, 
  L_PROVINCE, COMPANY_ID, CENTER_ID, AGRARIAN_FLAG, RURAL_FLAG, SMS_1, L_SHARE_AMOUNT, 
  L_PPI_SCORE, LAST_PAYMENT_DATE, L_MEMBER_CLASS, L_GLIP_FLAG, ALT_CUST_ID, GIVEN_NAMES, 
  L_MIDDLE_NAME, FAMILY_NAME, L_BARANGAY_NAME, GENDER, ZIPCODE, L_TIN_ID, L_SSS_ID, 
  L_PHILHEALTH_ID, TITLE, L_REL_DOSRI, L_OCCUPATION, L_EDUCATION, L_DISABL_FLAG, 
  SPOUSE_NAME, MBA_WEEKLY_AMOUNT, RESIGNED_REMARK, L_GLIP_MBA_AMT, L_RELIGION
FROM 
  dwh_ph.customer c
`

func (q *QueriesCustomer) GetCustomerInfo(ctx context.Context, id int64) (CustomerInfo, error) {
	return populateData(q, ctx,
		fmt.Sprintf("%s WHERE c.CUSTOMER_ID = %v", selectSQL, id))
}

type ListCustomerParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesCustomer) ListCustomer(ctx context.Context, arg ListCustomerParams) ([]CustomerInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d",
			selectSQL, arg.Limit, arg.Offset)
	} else {
		sql = selectSQL
	}
	return populateDatas(q, ctx, sql)
}
