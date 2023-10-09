package db

type KPLUSResponse struct {
	RetCode   int64  `json:"retCode"`
	Message   string `json:"message"`
	Reference string `json:"reference"`
}

type KPLUSCustomer struct {
	INAIIID               string `json:"iNAIIID"`
	CustomerId            int64  `json:"customerId"`
	IIID                  int64  `json:"iIID"`
	Cid                   int64  `json:"cid"`
	LastName              string `json:"lastName"`
	FirstName             string `json:"firstName"`
	MiddleName            string `json:"middleName"`
	MaidenFName           string `json:"maidenFName"`
	MaidenLName           string `json:"maidenLName"`
	MaidenMName           string `json:"maidenMName"`
	DoBirth               string `json:"doBirth"`
	BirthPlace            string `json:"birthPlace"`
	Sex                   string `json:"sex"`
	CivilStatus           string `json:"civilStatus"`
	Title                 string `json:"title"`
	Status                int64  `json:"status"`
	StatusDesc            string `json:"statusDesc"`
	Classification        int64  `json:"classification"`
	ClassificationDesc    string `json:"classificationDesc"`
	SubClassification     int64  `json:"subClassification"`
	SubClassificationDesc string `json:"subClassificationDesc"`
	Business              string `json:"business"`
	DoEntry               string `json:"doEntry"`
	DoRecognized          string `json:"doRecognized"`
	DoResigned            string `json:"doResigned"`
	BrCode                string `json:"brCode"`
	BranchName            string `json:"branchName"`
	UnitCode              string `json:"unitCode"`
	UnitName              string `json:"unitName"`
	CenterCode            string `json:"centerCode"`
	CenterName            string `json:"centerName"`
	Dosri                 bool   `json:"dosri"`
	Reffered              string `json:"reffered"`
	Remarks               string `json:"remarks"`
	AccountNumber         string `json:"accountNumber"`
	SearchName            string `json:"searchName"`
	MemberMaidenFName     string `json:"memberMaidenFName"`
	MemberMaidenLName     string `json:"memberMaidenLName"`
	MemberMaidenMName     string `json:"memberMaidenMName"`
}

type Savings struct {
	BrCode     string  `json:"brCode"`
	INAIIID    string  `json:"iNAIIID"`
	CustomerId int64   `json:"customerId"`
	IIID       int64   `json:"iIID"`
	Cid        int64   `json:"cid"`
	Acc        string  `json:"acc"`
	AcctType   int64   `json:"acctType"`
	AccDesc    string  `json:"accDesc"`
	Dopen      string  `json:"dopen"`
	StatusDesc string  `json:"statusDesc"`
	Balance    float64 `json:"balance"`
	Status     int64   `json:"status"`
}

type ByTrnDateTrn []TransactionHistory

func (s ByTrnDateTrn) Len() int {
	return len(s)
}

func (s ByTrnDateTrn) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByTrnDateTrn) Less(i, j int) bool {
	if s[i].TrnDate == s[j].TrnDate {
		return s[i].Trn > s[j].Trn
	}
	return s[i].TrnDate > s[j].TrnDate
}

type TransactionHistory struct {
	Acc           string  `json:"acc"`
	AccountId     string  `json:"accountId"`
	TrnDate       string  `json:"trnDate"`
	ValueDate     string  `json:"valueDate"`
	TrnHeadId     int64   `json:"trnHeadId"`
	Trn           int64   `json:"trn"`
	AlternateKey  string  `json:"alternateKey"`
	Prin          float64 `json:"prin"`
	Intr          float64 `json:"intr"`
	TrnAmount     float64 `json:"trnAmount"`
	BalPrin       float64 `json:"balPrin"`
	BalInt        float64 `json:"BalInt"`
	Balance       float64 `json:"balance"`
	PaidPrin      float64 `json:"paidPrin"`
	PaidInt       float64 `json:"paidInt"`
	Particulars   string  `json:"particulars"`
	TrnType       string  `json:"trnType"`
	Username      string  `json:"username"`
	IsFinancial   bool    `json:"isFinancial"`
	NormalBalance bool    `json:"normalBalance"`
}

type ColShtperAcc struct {
	INAIIID         string  `json:"iNAIIID"`
	BrCode          string  `json:"brCode"`
	AppType         int64   `json:"appType"`
	Code            int64   `json:"code"`
	Status          int64   `json:"status"`
	StatusDesc      string  `json:"statusDesc"`
	Acc             string  `json:"acc"`
	Iiid            int64   `json:"iiid"`
	CustomerId      int64   `json:"customerId"`
	CentralOfficeId int64   `json:"centralOfficeId"`
	CID             int64   `json:"cID"`
	UM              string  `json:"uM"`
	ClientName      string  `json:"clientName"`
	CenterCode      int64   `json:"centerCode"`
	CenterName      string  `json:"centerName"`
	ManCode         int64   `json:"manCode"`
	Unit            string  `json:"unit"`
	AreaCode        int64   `json:"areaCode"`
	Area            string  `json:"area"`
	StaffName       string  `json:"staffName"`
	AcctType        int64   `json:"acctType"`
	AcctDesc        string  `json:"acctDesc"`
	DisbDate        string  `json:"disbDate"`
	DateStart       string  `json:"dateStart"`
	Maturity        string  `json:"maturity"`
	Principal       float64 `json:"principal"`
	Interest        float64 `json:"interest"`
	Gives           int64   `json:"gives"`
	IbalPrin        float64 `json:"ibalPrin"`
	IbalInt         float64 `json:"ibalInt"`
	BalPrin         float64 `json:"balPrin"`
	BalInt          float64 `json:"balInt"`
	Amort           float64 `json:"amort"`
	DuePrin         float64 `json:"duePrin"`
	DueInt          float64 `json:"dueInt"`
	LoanBal         float64 `json:"loanBal"`
	SaveBal         float64 `json:"saveBal"`
	WaivedInt       float64 `json:"waivedInt"`
	UnPaidCtr       int64   `json:"unPaidCtr"`
	WritenOff       int64   `json:"writenOff"`
	Classification  int64   `json:"classification"`
	ClassDesc       int64   `json:"classDesc"`
	WriteOff        int64   `json:"writeOff"`
	Pay             float64 `json:"pay"`
	Withdraw        float64 `json:"withdraw"`
	Type            int64   `json:"type"`
	OrgName         string  `json:"orgName"`
	OrgAddress      string  `json:"orgAddress"`
	MeetingDate     string  `json:"meetingDate"`
	MeetingDay      int64   `json:"meetingDay"`
	SharesOfStock   float64 `json:"sharesOfStock"`
	DateEstablished string  `json:"dateEstablished"`
	Uuid            string  `json:"uuid"`
}

type References struct {
	Status             string `json:"status"`
	RefType            string `json:"refType"`
	Code               string `json:"code"`
	ShortName          string `json:"shortName"`
	OwnerRID           string `json:"ownerRID"`
	StatusID           string `json:"statusID"`
	Title              string `json:"title"`
	RefTypeID          string `json:"refTypeID"`
	RefTypeRID         string `json:"refTypeRID"`
	RefTypeTitle       string `json:"refTypeTitle"`
	RefID              string `json:"refID"`
	RefRID             string `json:"refRID"`
	ParentID           string `json:"parentID"`
	RefTypeParentID    string `json:"refTypeParentID"`
	RefTypeParent      string `json:"refTypeParent"`
	RefTypeTitleParent string `json:"refTypeTitleParent"`
	Xml                string `json:"xml"`
	Parent             string `json:"parent"`
}

type Loan struct {
	Acc         string `json:"acc"`
	Status      string `json:"status"`
	DateRelease string `json:"dateRelease"`
	AcctType    string `json:"acctType"`
	Principal   string `json:"principal"`
	Interest    string `json:"interest"`
	Oth         string `json:"oth"`
	Balance     string `json:"balance"`
	Term        string `json:"term"`
	PaidTerm    string `json:"paidTerm"`
}

type LoanInfo struct {
	Cid                 string `json:"cid"`
	Acc                 string `json:"acc"`
	AppType             string `json:"appType"`
	AcctType            string `json:"acctType"`
	Accdesc             string `json:"accdesc"`
	Dopen               string `json:"dopen"`
	Domaturity          string `json:"domaturity"`
	Term                string `json:"term"`
	Weekspaid           string `json:"weekspaid"`
	Status              string `json:"status"`
	Principal           string `json:"principal"`
	Interest            string `json:"interest"`
	Others              string `json:"others"`
	Discounted          string `json:"discounted"`
	Netproceed          string `json:"netproceed"`
	Balance             string `json:"balance"`
	Prin                string `json:"prin"`
	Intr                string `json:"intr"`
	Oth                 string `json:"oth"`
	Penalty             string `json:"penalty"`
	Waivedint           string `json:"waivedint"`
	Disbby              string `json:"disbby"`
	Approvby            string `json:"approvby"`
	Cycle               string `json:"cycle"`
	Frequency           string `json:"frequency"`
	Annumdiv            string `json:"annumdiv"`
	Lngrpcode           string `json:"lngrpcode"`
	Proff               string `json:"proff"`
	Fundsource          string `json:"fundsource"`
	Conintrate          string `json:"conintrate"`
	Amortcond           string `json:"amortcond"`
	Amortcondvalue      string `json:"amortcondvalue"`
	Classification_code string `json:"classification_code"`
	Classification_type string `json:"classification_type"`
	Remarks             string `json:"remarks"`
	Amort               string `json:"amort"`
	IsLumpsum           string `json:"isLumpsum"`
	LoanID              string `json:"loanID"`
	Charges             string `json:"charges"`
}

type Charges struct {
	Acc       string `json:"acc"`
	Charges   string `json:"charges"`
	AmortList string `json:"amortList"`
}

type LoanInfoAmort struct {
	Dnum        string `json:"dnum"`
	Acc         string `json:"acc"`
	DueDate     string `json:"dueDate"`
	InstFlag    string `json:"instFlag"`
	Prin        string `json:"prin"`
	Intr        string `json:"intr"`
	Oth         string `json:"oth"`
	Penalty     string `json:"penalty"`
	EndBal      string `json:"endBal"`
	EndInt      string `json:"endInt"`
	EndOth      string `json:"endOth"`
	InstPd      string `json:"instPd"`
	PenPd       string `json:"penPd"`
	CarVal      string `json:"carVal"`
	UpInt       string `json:"upInt"`
	ServFee     string `json:"servFee"`
	PledgeAmort string `json:"pledgeAmort"`
}

type SavingForSuperApp struct {
	Cid              string `json:"cid"`
	FullName         string `json:"fullName"`
	Acc              string `json:"acc"`
	Balance          string `json:"balance"`
	Withdrawable     string `json:"withdrawable"`
	CenterCode       string `json:"centerCode"`
	UnitCode         string `json:"unitCode"`
	CenterName       string `json:"centerName"`
	UnitName         string `json:"unitName"`
	WithdrawalAmount string `json:"withdrawalAmount"`
}

type FundTransfer struct {
	SourceAccount       string `json:"sourceAccount"`
	TargetAccount       string `json:"targetAccount"`
	Amount              string `json:"amount"`
	Username            string `json:"username"`
	TrnReference        string `json:"trnReference"`
	Particulars         string `json:"particulars"`
	TransFee            string `json:"transFee"`
	TransFeeParticulars string `json:"transFeeParticulars"`
}
