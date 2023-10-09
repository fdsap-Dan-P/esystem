package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/sqltocsv"
	"github.com/shopspring/decimal"
)

func (q *QueriesLocal) Sql2Csv(ctx context.Context, sql string, filenamePath string) error {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		fmt.Println("Sql2Csv-> Error:", err)
		return err
	}
	defer rows.Close()
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Current directory:", currentDir)

	csvConverter := sqltocsv.New(rows)
	csvConverter.TimeFormat = time.RFC3339
	csvConverter.Delimiter = '|'
	csvConverter.WriteHeaders = false
	return csvConverter.WriteFile(filenamePath)
}

type Area struct {
	ModCtr    int64          `json:"modCtr"`
	BrCode    string         `json:"brCode"`
	ModAction string         `json:"modAction"`
	AreaCode  int64          `json:"areaCode"`
	Area      sql.NullString `json:"area"`
}

type Unit struct {
	ModCtr      int64          `json:"modCtr"`
	BrCode      string         `json:"brCode"`
	ModAction   string         `json:"modAction"`
	UnitCode    int64          `json:"unitCode"`
	Unit        sql.NullString `json:"unit"`
	AreaCode    sql.NullInt64  `json:"areaCode"`
	FName       sql.NullString `json:"fName"`
	LName       sql.NullString `json:"lName"`
	MName       sql.NullString `json:"mName"`
	VatReg      sql.NullString `json:"vatReg"`
	UnitAddress sql.NullString `json:"unitAddress"`
}

type Center struct {
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

type Customer struct {
	ModCtr            int64               `json:"modCtr"`
	BrCode            string              `json:"brCode"`
	ModAction         string              `json:"modAction"`
	CID               int64               `json:"CID"`
	CenterCode        sql.NullString      `json:"centerCode"`
	Title             sql.NullInt64       `json:"title"`
	LName             sql.NullString      `json:"lName"`
	FName             sql.NullString      `json:"fName"`
	MName             sql.NullString      `json:"mName"`
	MaidenFName       sql.NullString      `json:"maidenFName"`
	MaidenLName       sql.NullString      `json:"maidenLName"`
	MaidenMName       sql.NullString      `json:"maidenMName"`
	Sex               sql.NullString      `json:"sex"`
	BirthDate         sql.NullTime        `json:"birthDate"`
	BirthPlace        sql.NullString      `json:"birthPlace"`
	CivilStatus       sql.NullInt64       `json:"civilStatus"`
	CustType          sql.NullInt64       `json:"custType"`
	Remarks           sql.NullString      `json:"remarks"`
	Status            sql.NullInt64       `json:"status"`
	Classification    sql.NullInt64       `json:"classification"`
	DepoType          sql.NullString      `json:"depoType"`
	SubClassification sql.NullInt64       `json:"subClassification"`
	PledgeAmount      decimal.NullDecimal `json:"pledgeAmount"`
	MutualAmount      decimal.NullDecimal `json:"mutualAmount"`
	PangarapAmount    decimal.NullDecimal `json:"pangarapAmount"`
	KatuparanAmount   decimal.NullDecimal `json:"katuparanAmount"`
	InsuranceAmount   decimal.NullDecimal `json:"insuranceAmount"`
	AccPledge         decimal.NullDecimal `json:"accPledge"`
	AccMutual         decimal.NullDecimal `json:"accMutual"`
	AccPang           decimal.NullDecimal `json:"accPang"`
	AccKatuparan      decimal.NullDecimal `json:"accKatuparan"`
	AccInsurance      decimal.NullDecimal `json:"accInsurance"`
	LoanLimit         decimal.NullDecimal `json:"loanLimit"`
	CreditLimit       decimal.NullDecimal `json:"creditLimit"`
	DateRecognized    sql.NullTime        `json:"dateRecognized"`
	DateResigned      sql.NullTime        `json:"dateResigned"`
	DateEntry         sql.NullTime        `json:"dateEntry"`
	GoldenLifeDate    sql.NullTime        `json:"goldenLifeDate"`
	Restricted        sql.NullString      `json:"restricted"`
	Borrower          sql.NullString      `json:"borrower"`
	CoMaker           sql.NullString      `json:"coMaker"`
	Guarantor         sql.NullString      `json:"guarantor"`
	DOSRI             sql.NullInt64       `json:"dOSRI"`
	IDCode1           sql.NullInt64       `json:"iDCode1"`
	IDNum1            sql.NullString      `json:"iDNum1"`
	IDCode2           sql.NullInt64       `json:"iDCode2"`
	IDNum2            sql.NullString      `json:"iDNum2"`
	Contact1          sql.NullString      `json:"contact1"`
	Contact2          sql.NullString      `json:"contact2"`
	Phone1            sql.NullString      `json:"phone1"`
	Reffered1         sql.NullString      `json:"reffered1"`
	Reffered2         sql.NullString      `json:"reffered2"`
	Reffered3         sql.NullString      `json:"reffered3"`
	Education         sql.NullInt64       `json:"education"`
	Validity1         sql.NullTime        `json:"validity1"`
	Validity2         sql.NullTime        `json:"validity2"`
	BusinessType      sql.NullInt64       `json:"businessType"`
	AccountNumber     sql.NullString      `json:"accountNumber"`
	IIID              sql.NullInt64       `json:"iIID"`
	Religion          sql.NullInt64       `json:"religion"`
}

type Addresses struct {
	ModCtr         int64          `json:"modCtr"`
	BrCode         string         `json:"brCode"`
	ModAction      string         `json:"modAction"`
	CID            int64          `json:"CID"`
	SeqNum         int64          `json:"seqNum"`
	AddressDetails sql.NullString `json:"addressDetails"`
	Barangay       sql.NullString `json:"barangay"`
	City           sql.NullString `json:"city"`
	Province       sql.NullString `json:"province"`
	Phone1         sql.NullString `json:"phone1"`
	Phone2         sql.NullString `json:"phone2"`
	Phone3         sql.NullString `json:"phone3"`
	Phone4         sql.NullString `json:"phone4"`
}

type LnMaster struct {
	ModCtr      int64               `json:"modCtr"`
	BrCode      string              `json:"brCode"`
	ModAction   string              `json:"modAction"`
	CID         int64               `json:"CID"`
	Acc         string              `json:"acc"`
	AcctType    sql.NullInt64       `json:"acctType"`
	DisbDate    sql.NullTime        `json:"disbDate"`
	Principal   decimal.NullDecimal `json:"principal"`
	Interest    decimal.NullDecimal `json:"interest"`
	NetProceed  decimal.NullDecimal `json:"netProceed"`
	Gives       sql.NullInt64       `json:"gives"`
	Frequency   sql.NullInt64       `json:"frequency"`
	AnnumDiv    sql.NullInt64       `json:"annumDiv"`
	Prin        decimal.NullDecimal `json:"prin"`
	IntR        decimal.NullDecimal `json:"intR"`
	WaivedInt   decimal.NullDecimal `json:"waivedInt"`
	WeeksPaid   sql.NullInt64       `json:"weeksPaid"`
	DoMaturity  sql.NullTime        `json:"doMaturity"`
	ConIntRate  decimal.NullDecimal `json:"conIntRate"`
	Status      sql.NullString      `json:"status"`
	Cycle       sql.NullInt64       `json:"cycle"`
	LNGrpCode   sql.NullInt64       `json:"lNGrpCode"`
	Proff       sql.NullInt64       `json:"proff"`
	FundSource  sql.NullString      `json:"fundSource"`
	DOSRI       sql.NullBool        `json:"DOSRI"`
	LnCategory  sql.NullInt64       `json:"lnCategory"`
	OpenDate    sql.NullTime        `json:"openDate"`
	LastTrnDate sql.NullTime        `json:"lastTrnDate"`
	DisbBy      sql.NullString      `json:"disbBy"`
}

type SaMaster struct {
	ModCtr     int64               `json:"modCtr"`
	BrCode     string              `json:"brCode"`
	ModAction  string              `json:"modAction"`
	Acc        string              `json:"acc"`
	CID        int64               `json:"cID"`
	Type       int64               `json:"type"`
	Balance    decimal.NullDecimal `json:"balance"`
	DoLastTrn  sql.NullTime        `json:"doLastTrn"`
	DoStatus   sql.NullTime        `json:"doStatus"`
	Dopen      sql.NullTime        `json:"dopen"`
	DoMaturity sql.NullTime        `json:"doMaturity"`
	Status     sql.NullString      `json:"status"`
}

type TrnMaster struct {
	ModCtr     int64           `json:"modCtr"`
	BrCode     string          `json:"brCode"`
	ModAction  string          `json:"modAction"`
	Acc        string          `json:"acc"`
	TrnDate    time.Time       `json:"trnDate"`
	Trn        int64           `json:"trn"`
	TrnType    sql.NullInt64   `json:"trnType"`
	OrNo       sql.NullInt64   `json:"orNo"`
	Prin       decimal.Decimal `json:"prin"`
	IntR       decimal.Decimal `json:"intR"`
	WaivedInt  decimal.Decimal `json:"waivedInt"`
	RefNo      sql.NullString  `json:"refNo"`
	UserName   sql.NullString  `json:"userName"`
	Particular sql.NullString  `json:"particular"`
}

type SaTrnMaster struct {
	ModCtr      int64               `json:"modCtr"`
	BrCode      string              `json:"brCode"`
	ModAction   string              `json:"modAction"`
	Acc         string              `json:"acc"`
	TrnDate     time.Time           `json:"trnDate"`
	Trn         int64               `json:"trn"`
	TrnType     sql.NullInt64       `json:"trnType"`
	OrNo        sql.NullInt64       `json:"orNo"`
	TrnAmt      decimal.NullDecimal `json:"trnAmt"`
	RefNo       sql.NullString      `json:"refNo"`
	Particular  string              `json:"particular"`
	TermId      string              `json:"termId"`
	UserName    string              `json:"userName"`
	PendApprove string              `json:"pendApprove"`
}

type LoanInst struct {
	ModCtr    int64           `json:"modCtr"`
	BrCode    string          `json:"brCode"`
	ModAction string          `json:"modAction"`
	Acc       string          `json:"acc"`
	Dnum      int64           `json:"dnum"`
	DueDate   time.Time       `json:"dueDate"`
	InstFlag  int64           `json:"instFlag"`
	DuePrin   decimal.Decimal `json:"duePrin"`
	DueInt    decimal.Decimal `json:"dueInt"`
	UpInt     decimal.Decimal `json:"upInt"`
}

type LnChrgData struct {
	ModCtr    int64           `json:"modCtr"`
	BrCode    string          `json:"brCode"`
	ModAction string          `json:"modAction"`
	Acc       string          `json:"acc"`
	ChrgCode  int64           `json:"chrgCode"`
	RefAcc    sql.NullString  `json:"refAcc"`
	ChrAmnt   decimal.Decimal `json:"chrAmnt"`
}

type CustAddInfoList struct {
	ModCtr     int64  `json:"modCtr"`
	BrCode     string `json:"brCode"`
	ModAction  string `json:"modAction"`
	InfoCode   int64  `json:"infoCode"`
	InfoOrder  string `json:"infoOrder"`
	Title      string `json:"title"`
	InfoType   string `json:"infoType"`
	InfoLen    int64  `json:"infoLen"`
	InfoFormat string `json:"infoFormat"`
	InputType  int64  `json:"inputType"`
	InfoSource string `json:"infoSource"`
}

type CustAddInfoGroup struct {
	ModCtr     int64          `json:"modCtr"`
	BrCode     string         `json:"brCode"`
	ModAction  string         `json:"modAction"`
	InfoGroup  int64          `json:"infoGroup"`
	GroupTitle sql.NullString `json:"groupTitle"`
	Remarks    sql.NullString `json:"remarks"`
	ReqOnEntry sql.NullBool   `json:"reqOnEntry"`
	ReqOnExit  sql.NullBool   `json:"reqOnExit"`
	Link2Loan  sql.NullInt64  `json:"link2Loan"`
	Link2Save  sql.NullInt64  `json:"link2Save"`
}

type CustAddInfoGroupNeed struct {
	ModCtr      int64          `json:"modCtr"`
	BrCode      string         `json:"brCode"`
	ModAction   string         `json:"modAction"`
	InfoGroup   int64          `json:"infoGroup"`
	InfoCode    int64          `json:"infoCode"`
	InfoProcess sql.NullString `json:"infoProcess"`
}

type CustAddInfo struct {
	ModCtr    int64     `json:"modCtr"`
	BrCode    string    `json:"brCode"`
	ModAction string    `json:"modAction"`
	CID       int64     `json:"cID"`
	InfoDate  time.Time `json:"infoDate"`
	InfoCode  int64     `json:"infoCode"`
	Info      string    `json:"info"`
	InfoValue int64     `json:"infoValue"`
}

type MutualFund struct {
	ModCtr    int64           `json:"modCtr"`
	BrCode    string          `json:"brCode"`
	ModAction string          `json:"modAction"`
	CID       int64           `json:"CID"`
	OrNo      sql.NullInt64   `json:"orNo"`
	TrnDate   time.Time       `json:"trnDate"`
	TrnType   sql.NullString  `json:"trnType"`
	TrnAmt    decimal.Decimal `json:"trnAmt"`
	UserName  sql.NullString  `json:"userName"`
}

type ReferencesDetails struct {
	ModCtr             int64          `json:"modCtr"`
	BrCode             string         `json:"brCode"`
	ModAction          string         `json:"modAction"`
	ID                 int64          `json:"id"`
	RefID              int64          `json:"refID"`
	PurposeDescription sql.NullString `json:"purposeDescription"`
	ParentID           sql.NullInt64  `json:"parentID"`
	CodeID             sql.NullInt64  `json:"codeID"`
	Stat               sql.NullInt64  `json:"stat"`
}

type CenterWorker struct {
	ModCtr      int64          `json:"modCtr"`
	BrCode      string         `json:"brCode"`
	ModAction   string         `json:"modAction"`
	AOID        int64          `json:"aoID"`
	Lname       sql.NullString `json:"lname"`
	FName       sql.NullString `json:"fName"`
	Mname       sql.NullString `json:"mname"`
	PhoneNumber sql.NullString `json:"phoneNumber"`
	EmpNo       sql.NullString `json:"empNo"`
}

type Writeoff struct {
	ModCtr     int64           `json:"modCtr"`
	BrCode     string          `json:"brCode"`
	ModAction  string          `json:"modAction"`
	Acc        string          `json:"acc"`
	DisbDate   time.Time       `json:"disbDate"`
	Principal  decimal.Decimal `json:"principal"`
	Interest   decimal.Decimal `json:"interest"`
	BalPrin    decimal.Decimal `json:"balPrin"`
	BalInt     decimal.Decimal `json:"balInt"`
	TrnDate    time.Time       `json:"trnDate"`
	AcctType   string          `json:"acctType"`
	Print      sql.NullString  `json:"print"`
	PostedBy   sql.NullString  `json:"postedBy"`
	VerifiedBy sql.NullString  `json:"verifiedBy"`
}

type Accounts struct {
	ModCtr    int64          `json:"modCtr"`
	BrCode    string         `json:"brCode"`
	ModAction string         `json:"modAction"`
	Acc       string         `json:"acc"`
	Title     string         `json:"title"`
	Category  int64          `json:"category"`
	Type      string         `json:"type"`
	MainCD    sql.NullString `json:"mainCD"`
	Parent    sql.NullString `json:"parent"`
}

type JnlHeaders struct {
	ModCtr      int64          `json:"modCtr"`
	BrCode      string         `json:"brCode"`
	ModAction   string         `json:"modAction"`
	Trn         string         `json:"trn"`
	TrnDate     time.Time      `json:"trnDate"`
	Particulars string         `json:"particulars"`
	UserName    sql.NullString `json:"userName"`
	Code        int64          `json:"code"`
}

type JnlDetails struct {
	ModCtr    int64               `json:"modCtr"`
	BrCode    string              `json:"brCode"`
	ModAction string              `json:"modAction"`
	Acc       string              `json:"acc"`
	Trn       string              `json:"trn"`
	Series    sql.NullInt64       `json:"series"`
	Debit     decimal.NullDecimal `json:"debit"`
	Credit    decimal.NullDecimal `json:"credit"`
}

type LedgerDetails struct {
	ModCtr    int64           `json:"modCtr"`
	BrCode    string          `json:"brCode"`
	ModAction string          `json:"modAction"`
	TrnDate   time.Time       `json:"trnDate"`
	Acc       string          `json:"acc"`
	Balance   decimal.Decimal `json:"balance"`
}

type UsersList struct {
	ModCtr              int64          `json:"modCtr"`
	BrCode              string         `json:"brCode"`
	ModAction           string         `json:"modAction"`
	UserId              string         `json:"userId"`
	AccessCode          sql.NullInt64  `json:"accessCode"`
	LName               string         `json:"lName"`
	FName               string         `json:"fName"`
	MName               string         `json:"mName"`
	DateHired           sql.NullTime   `json:"dateHired"`
	BirthDay            sql.NullTime   `json:"birthDay"`
	DateGiven           sql.NullTime   `json:"dateGiven"`
	DateExpired         sql.NullTime   `json:"dateExpired"`
	Address             sql.NullString `json:"address"`
	Position            sql.NullString `json:"position"`
	AreaCode            sql.NullInt64  `json:"areaCode"`
	ManCode             sql.NullInt64  `json:"manCode"`
	AddInfo             sql.NullString `json:"addInfo"`
	Passwd              []byte         `json:"passwd"`
	Attempt             sql.NullInt64  `json:"attempt"`
	DateLocked          sql.NullTime   `json:"dateLocked"`
	Remarks             sql.NullString `json:"remarks"`
	Picture             []byte         `json:"picture"`
	IsLoggedIn          bool           `json:"isLoggedIn"`
	AccountExpirationDt sql.NullTime   `json:"accountExpirationDt"`
}

type MultiplePaymentReceipt struct {
	ModCtr    int64           `json:"modCtr"`
	BrCode    string          `json:"brCode"`
	ModAction string          `json:"modAction"`
	TrnDate   time.Time       `json:"trnDate"`
	OrNo      int64           `json:"orNo"`
	CID       int64           `json:"cID"`
	PrNo      int64           `json:"prNo"`
	UserName  string          `json:"userName"`
	TermId    string          `json:"ermId"`
	AmtPaid   decimal.Decimal `json:"amtPaid"`
}

type InActiveCID struct {
	ModCtr        int64          `json:"modCtr"`
	BrCode        string         `json:"brCode"`
	ModAction     string         `json:"modAction"`
	CID           int64          `json:"cid"`
	InActive      bool           `json:"inActive"`
	DateStart     time.Time      `json:"dateStart"`
	DateEnd       sql.NullTime   `json:"dateEnd"`
	UserId        string         `json:"userId"`
	DeactivatedBy sql.NullString `json:"deactivatedBy"`
}

type ReactivateWriteoff struct {
	ModCtr       int64          `json:"modCtr"`
	BrCode       string         `json:"brCode"`
	ModAction    string         `json:"modAction"`
	ID           int64          `json:"id"`
	CID          int64          `json:"cid"`
	DeactivateBy sql.NullString `json:"deactivatedBy"`
	ReactivateBy sql.NullString `json:"reactivatedBy"`
	Status       int64          `json:"status"`
	StatusDate   time.Time      `json:"statusDate"`
}

type LnBeneficiary struct {
	ModCtr     int64          `json:"modCtr"`
	BrCode     string         `json:"brCode"`
	ModAction  string         `json:"modAction"`
	Acc        string         `json:"acc"`
	BDay       time.Time      `json:"bDay"`
	EducLvl    string         `json:"educLvl"`
	Gender     bool           `json:"gender"`
	LastName   sql.NullString `json:"lastName"`
	FirstName  sql.NullString `json:"firstName"`
	MiddleName sql.NullString `json:"middleName"`
	Remarks    sql.NullString `json:"remarks"`
}

type ColSht struct {
	BrCode          string          `json:"brCode"`
	AppType         int64           `json:"appType"`
	Code            int64           `json:"code"`
	Status          int64           `json:"status"`
	Acc             string          `json:"acc"`
	CID             int64           `json:"cID"`
	UM              string          `json:"uM"`
	ClientName      string          `json:"clientName"`
	CenterCode      string          `json:"centerCode"`
	CenterName      string          `json:"centerName"`
	ManCode         int64           `json:"manCode"`
	Unit            string          `json:"unit"`
	AreaCode        int64           `json:"areaCode"`
	Area            string          `json:"area"`
	StaffName       string          `json:"staffName"`
	AcctType        int64           `json:"acctType"`
	AcctDesc        string          `json:"acctDesc"`
	DisbDate        time.Time       `json:"disbDate"`
	DateStart       time.Time       `json:"dateStart"`
	Maturity        time.Time       `json:"maturity"`
	Principal       decimal.Decimal `json:"principal"`
	Interest        decimal.Decimal `json:"interest"`
	Gives           int64           `json:"gives"`
	BalPrin         decimal.Decimal `json:"balPrin"`
	BalInt          decimal.Decimal `json:"balInt"`
	Amort           decimal.Decimal `json:"amort"`
	DuePrin         decimal.Decimal `json:"duePrin"`
	DueInt          decimal.Decimal `json:"dueInt"`
	LoanBal         decimal.Decimal `json:"loanBal"`
	SaveBal         decimal.Decimal `json:"saveBal"`
	WaivedInt       decimal.Decimal `json:"waivedInt"`
	UnPaidCtr       int64           `json:"unPaidCtr"`
	WrittenOff      int64           `json:"WrittenOff"`
	OrgName         string          `json:"orgName"`
	OrgAddress      string          `json:"orgAddress"`
	MeetingDate     time.Time       `json:"meetingDate"`
	MeetingDay      int64           `json:"meetingDay"`
	SharesOfStock   decimal.Decimal `json:"sharesOfStock"`
	DateEstablished time.Time       `json:"dateEstablished"`
	Classification  int64           `json:"classification"`
	WriteOff        int64           `json:"writeOff"`
}
