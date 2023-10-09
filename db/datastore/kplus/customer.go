package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

const kPLUSCustomerSQL = `-- name: kPLUSCustomerSQL :one
SELECT * 
FROM
(SELECT 
    id_number INAIIID, c.ID Customer_Id, ii.ID IIID, c.cid, 
    ii.last_name lastName, COALESCE(ii.first_name,'') firstName, COALESCE(ii.middle_name,'') middleName, 
    COALESCE(per.maiden_fname,'') maidenFName, COALESCE(per.maiden_Lname,'') maidenLName, COALESCE(maiden_mname,'') maidenMName, 
    COALESCE(ii.birthday,'1900-01-01') doBirth, COALESCE(geo.Location,'') birthPlace, ii.sex sex,  
    COALESCE(civil.title,'') civilStatus, COALESCE(ii.title,'') title, COALESCE(stat.short_name,'15') status, 
    COALESCE(stat.title,'') statusDesc, COALESCE(cls.code,0) classification,  
    COALESCE(cls.title,'') classificationDesc, COALESCE(subcls.code,0) subClassification,  COALESCE(subcls.title,'') subClassificationDesc,
    COALESCE(per.business_name,'') business, COALESCE(c.date_entry,'1900-01-01') doEntry,  
    COALESCE(c.date_recognized,'1900-01-01')  doRecognized, COALESCE(c.date_resigned,'1900-01-01')  doResigned, 
    br.code brCode, br.office_name branchName, o.code unitCode,  
    o.office_name unitName, cg.code centerCode, cg.group_name centerName, c.dosri dosri, 
    FullNameTFMLS(refer.Title, refer.Last_Name, refer.First_Name, refer.Middle_Name, refer.Suffix_Name) reffered,  
    COALESCE(c.remarks,''), COALESCE(c.Primary_Acc,'') accountNumber, '' searchName, 
    COALESCE(per.maiden_fname,'') memberMaidenFName, COALESCE(per.maiden_lname,'') memberMaidenLName, COALESCE(per.maiden_mname,'') memberMaidenMName
FROM
  identity_info ii 
INNER JOIN Customer c on ii.id = c.iiid 
LEFT JOIN personal_info per on ii.id = per.id 
INNER JOIN reference stat on c.status_code = stat.code and lower(stat.ref_type) = 'customerstatus'
INNER JOIN reference civil on ii.civil_status_id = civil.id
INNER JOIN reference cls on c.classification_id = cls.id
INNER JOIN reference subcls on c.sub_classification_id = subcls.id
INNER JOIN customer_group cg on c.customer_group_id = cg.id
INNER JOIN office o on cg.office_id = o.id 
INNER JOIN office br on o.parent_id = br.id 
LEFT JOIN identity_info refer on c.refferedby_id = refer.id 
LEFT JOIN Geography geo on ii.birth_place_id = geo.ID
INNER JOIN IDs on ids.iiid  = ii.ID
INNER JOIN Reference idType on idType.id = ids.Type_id and lower(idType.title) = 'inai-iiid') d
`
const kPLUSCustomerInfo = `-- name: kPLUSCustomerInfo :one
SELECT 
  o.Code BrCode, COALESCE(Inaiiid, '') Inaiiid, cus.id Customer_ID, ii.Id IIID, cus.cid, ii.alternate_id, cus.customer_alt_id, 
  ii.title, ii.last_name, ii.first_name, ii.middle_name, ii.birthday,
  o.Id Office_Id, o.Short_Name Office_Short_Name, o.Office_Name,
  COALESCE(Phone, '') Phone, COALESCE(eMail, '') eMail
FROM Customer cus
INNER JOIN Office o on o.Id = cus.office_id 
INNER JOIN Identity_Info ii on cus.iiid = ii.Id
LEFT JOIN 
  (SELECT iiid, MAX(CASE WHEN lower(idType.title) = 'inai-iiid' THEN IDs.id_number ELSE '' END) Inaiiid
   FROM IDs
   INNER JOIN Reference idType on idType.id = ids.Type_id 
   GROUP BY iiid
  ) ids on ids.iiid  = ii.ID
`

type CustomerInfo struct {
	BrCode          string       `json:"brCode"`
	INAIIID         string       `json:"iNAIIID"`
	CustomerId      int64        `json:"customerId"`
	Iiid            int64        `json:"iIID"`
	Cid             int64        `json:"cid"`
	AlternateId     string       `json:"alternateId"`
	CustomerAltId   string       `json:"customerAltId"`
	Title           string       `json:"title"`
	LastName        string       `json:"lastName"`
	FirstName       string       `json:"firstName"`
	MiddleName      string       `json:"middleName"`
	Birthday        sql.NullTime `json:"birthday"`
	OfficeId        int64        `json:"officeId"`
	OfficeShortName string       `json:"officeShortName"`
	OfficeName      string       `json:"officeName"`
	Phone           string       `json:"phone"`
	Email           string       `json:"email"`
}

func populateKPLUSCustomersInfo(q *QueriesKPlus, ctx context.Context, sql string) ([]CustomerInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []CustomerInfo{}
	for rows.Next() {
		var i CustomerInfo
		err := rows.Scan(
			&i.BrCode,
			&i.INAIIID,
			&i.CustomerId,
			&i.Iiid,
			&i.Cid,
			&i.AlternateId,
			&i.CustomerAltId,
			&i.Title,
			&i.LastName,
			&i.FirstName,
			&i.MiddleName,
			&i.Birthday,
			&i.OfficeId,
			&i.OfficeShortName,
			&i.OfficeName,
			&i.Phone,
			&i.Email,
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

type CustomersInfoParam struct {
	BrCode        string `json:"brCode"`
	INAIIID       string `json:"iNAIIID"`
	CustomerId    int64  `json:"customerId"`
	Iiid          int64  `json:"iIID"`
	Cid           int64  `json:"cid"`
	AlternateId   string `json:"alternateId"`
	CustomerAltId string `json:"customerAltId"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	SearchName    string `json:"searchName"`
}

func (q *QueriesKPlus) GetCustomersInfo(ctx context.Context, arr CustomersInfoParam) ([]CustomerInfo, error) {
	sql := kPLUSCustomerInfo
	if arr.SearchName != "" {
		sql = fmt.Sprintf(`SELECT d.* FROM (%v) d INNER JOIN SearchName('%v', 10) s on s.Id = d.IIID WHERE true`, kPLUSCustomerInfo, arr.SearchName)
	} else {
		sql = fmt.Sprintf(`SELECT * FROM (%v) d WHERE true`, kPLUSCustomerInfo)
	}

	if arr.BrCode != "" {
		sql = fmt.Sprintf(`%v and BrCode = '%v'`, sql, arr.BrCode)
	}
	if arr.INAIIID != "" {
		sql = fmt.Sprintf(`%v and INAIIID = '%v'`, sql, arr.INAIIID)
	}
	if arr.CustomerId != 0 {
		sql = fmt.Sprintf(`%v and CustomerId = %v`, sql, arr.CustomerId)
	}
	if arr.Iiid != 0 {
		sql = fmt.Sprintf(`%v and Iiid = %v`, sql, arr.Iiid)
	}
	if arr.Cid != 0 {
		sql = fmt.Sprintf(`%v and Cid = %v`, sql, arr.Cid)
	}
	if arr.AlternateId != "" {
		sql = fmt.Sprintf(`%v and AlternateId = '%v'`, sql, arr.AlternateId)
	}
	if arr.CustomerAltId != "" {
		sql = fmt.Sprintf(`%v and CustomerAltId = '%v'`, sql, arr.CustomerAltId)
	}
	if arr.Phone != "" {
		sql = fmt.Sprintf(`%v and Phone = '%v'`, sql, arr.Phone)
	}
	if arr.Email != "" {
		sql = fmt.Sprintf(`%v and Email = '%v'`, sql, arr.Email)
	}
	log.Println(sql)
	return populateKPLUSCustomersInfo(q, ctx, sql)
}

func populateKPLUSCustomer(q *QueriesKPlus, ctx context.Context, sql string) (KPLUSCustomer, error) {
	var i KPLUSCustomer
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.INAIIID,
		&i.CustomerId,
		&i.IIID,
		&i.Cid,
		&i.LastName,
		&i.FirstName,
		&i.MiddleName,
		&i.MaidenFName,
		&i.MaidenLName,
		&i.MaidenMName,
		&i.DoBirth,
		&i.BirthPlace,
		&i.Sex,
		&i.CivilStatus,
		&i.Title,
		&i.Status,
		&i.StatusDesc,
		&i.Classification,
		&i.ClassificationDesc,
		&i.SubClassification,
		&i.SubClassificationDesc,
		&i.Business,
		&i.DoEntry,
		&i.DoRecognized,
		&i.DoResigned,
		&i.BrCode,
		&i.BranchName,
		&i.UnitCode,
		&i.UnitName,
		&i.CenterCode,
		&i.CenterName,
		&i.Dosri,
		&i.Reffered,
		&i.Remarks,
		&i.AccountNumber,
		&i.SearchName,
		&i.MemberMaidenFName,
		&i.MemberMaidenLName,
		&i.MemberMaidenMName,
	)
	return i, err
}

func populateKPLUSCustomers(q *QueriesKPlus, ctx context.Context, sql string) ([]KPLUSCustomer, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []KPLUSCustomer{}
	for rows.Next() {
		var i KPLUSCustomer
		err := rows.Scan(
			&i.INAIIID,
			&i.CustomerId,
			&i.IIID,
			&i.Cid,
			&i.LastName,
			&i.FirstName,
			&i.MiddleName,
			&i.MaidenFName,
			&i.MaidenLName,
			&i.MaidenMName,
			&i.DoBirth,
			&i.BirthPlace,
			&i.Sex,
			&i.CivilStatus,
			&i.Title,
			&i.Status,
			&i.StatusDesc,
			&i.Classification,
			&i.ClassificationDesc,
			&i.SubClassification,
			&i.SubClassificationDesc,
			&i.Business,
			&i.DoEntry,
			&i.DoRecognized,
			&i.DoResigned,
			&i.BrCode,
			&i.BranchName,
			&i.UnitCode,
			&i.UnitName,
			&i.CenterCode,
			&i.CenterName,
			&i.Dosri,
			&i.Reffered,
			&i.Remarks,
			&i.AccountNumber,
			&i.SearchName,
			&i.MemberMaidenFName,
			&i.MemberMaidenLName,
			&i.MemberMaidenMName,
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

func (q *QueriesKPlus) SearchCustomerCID(ctx context.Context, cid int64) (KPLUSCustomer, error) {
	return populateKPLUSCustomer(q, ctx, fmt.Sprintf("%s WHERE lower(trim(INAIIID)) = '%v' ", kPLUSCustomerSQL, strconv.FormatInt(cid, 10)))
}
