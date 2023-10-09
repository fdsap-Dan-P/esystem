CREATE TABLE IF NOT EXISTS staging.Office(
  UUID uuid NULL,	
  Code VarChar(50) NULL,
  ShortName VarChar(200) NULL,
  OfficeName VarChar(200) NULL,
  DateStablished Date Null,
  OfficeType VarChar(200) NULL,
  ParentUUID
  AlternateID VarChar(200) NULL,
  AddressDetail VarChar(200) NULL,
  Barangay VarChar(200) NULL,
  City VarChar(200) NULL,
  Province VarChar(200) NULL,
  Mobile VarChar(200) NULL,
  eMail VarChar(200) NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS idxOffice ON staging.Office(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxOffice ON staging.Office(Code);

CREATE TABLE IF NOT EXISTS staging.Employee(
    UUID uuid NULL,
    EmpNo VarChar(50) NULL,
    FName VarChar(100) NULL,
    LName VarChar(100) NULL,
    MName VarChar(100) NULL,
    Suffix VarChar(15) NULL,
    Position VarChar(100) NULL,
    OfficeId BigInt NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS idxEmployee ON staging.Employee(EmpNo);
CREATE UNIQUE INDEX IF NOT EXISTS idxEmployeeuuid ON staging.Employee(uuid);

CREATE TABLE IF NOT EXISTS staging.AOMap(
    EmpNo VarChar(50) NULL,
    BrCode VarChar(2) NULL,
    AOID BigInt NULL
); 
CREATE INDEX IF NOT EXISTS idxAOMap ON staging.AOMap(EmpNo);
CREATE UNIQUE INDEX IF NOT EXISTS idxAOMap2 ON staging.AOMap(BrCode,AOID);




INSERT INTO staging.AOMap(EmpNo, BrCode, AOID)
select * from 
(select a.empno, a.BrCode, a.AOID from 
(values 
('200601-00565','E3',39),
('200712-01365','E3',52),
('200804-01644','E3',66),
('201309-05147','E3',10),
('201407-06085','E3',23),
('201408-06263','E3',28),
('201408-06267','E3',63),
('201410-06461','E3',35),
('201506-07540','E3',44),
('201510-08476','E3',20),
('201511-08577','E3',7),
('201511-08579','E3',30),
('201511-08638','E3',18),
('201608-10959','E3',1),
('201612-12469','E3',36),
('201706-13191','E3',15),
('201707-14095','E3',22),
('201707-14531','E3',42),
('201710-16800','E3',17),
('201801-20708','E3',24),
('201802-17763','E3',29),
('201804-18338','E3',50),
('201805-20101','E3',51),
('201807-20704','E3',55),
('201807-21831','E3',67),
('201808-21881','E3',26),
('201811-24500','E3',16),
('201903-26328','E3',57),
('201905-28117','E3',14),
('201905-28120','E3',4),
('201906-29039','E3',11),
('202102-35883','E3',2),
('202208-40126','E3',31),
('202208-40353','E3',60),
('202210-41718','E3',68),
('201801-17287','E3',65),
('202203-38617','E3',9),
('202208-40736','E3',54),
('202208-40798','E3',45),
('202209-41252','E3',46),
('201807-20705','E3',3),
('201707-14532','E3',5),
('201702-13011','E3',6),
('201410-06461','E3',8),
('201811-24569','E3',21),
('ed5bb1cf-79ac-48d0-8f29-49aab3d3c94a','E3',12),
('097e7677-4de9-44b9-be74-cb81127d5734','E3',13),
('201602-09155','E3',19),
('05494285-4614-40cd-a38d-e7c50c0266c6','E3',25),
('a4392c1e-e334-43b5-99ea-e6b02854dd34','E3',27),
('c7e43b5f-3c53-4bba-9112-7872e7e83206','E3',32),
('f690dcca-2360-4295-b117-5b084671f623','E3',33),
('6ba6091b-71a8-4b66-8875-e9ef1849bf33','E3',34),
('4c326a3a-20da-406e-a6f4-9a58dcc7bc06','E3',37),
('21762780-aa25-4d1c-b829-323d08f8f148','E3',38),
('c910b629-2163-4e3a-b936-6d19b13c7206','E3',40),
('11012eb8-4986-4dd9-bf07-0dede61aa8d6','E3',41),
('b2cae61c-a26d-4f75-9b36-fdd67c859c1d','E3',43),
('09d5665f-a698-4a38-acfd-6fd2c34016a5','E3',47),
('3a2ee21f-ca21-44cd-ad25-cb1c9134501a','E3',48),
('4ef0409f-f8bd-4177-8519-feebfadeb0e6','E3',49),
('1acc3f22-21c7-4cbc-9bb0-73e6b87e0ceb','E3',53),
('46bc30d9-f0d0-4be8-9e40-61f3a06fed65','E3',56),
('53b666a9-3dae-422d-af8e-224dfec46c82','E3',58),
('7a9e1b43-4679-4fc3-b1c7-ada7b90e815d','E3',59),
('7f94970e-40b5-4f7d-806d-2bbaa8a43eeb','E3',61),
('977ed959-9280-4ac0-8fdf-2a2bda6b02ec','E3',62),
('c3bf163a-2a6e-4717-8231-244400f5c3ca','E3',64))
a(EmpNo, BrCode, AOID)
left join staging.aomap m on a.empno = m.empno 
where m.empno is null)a

on conflict(BrCode, AOID) do update 
set 
  EmpNo = excluded.EmpNo
  
CREATE TABLE IF NOT EXISTS staging.Area(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    AreaCode BigInt NULL,
    Area VarChar(30) NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxArea ON staging.Area(BrCode, AreaCode);

CREATE TABLE IF NOT EXISTS staging.Unit(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    UnitCode BigInt NULL,
    Unit VarChar(100) NULL,
    AreaCode BigInt NULL,
    FName VarChar(100) NULL,
    LName VarChar(100) NULL,
    MName VarChar(100) NULL,
    VatReg VarChar(100) NULL,
    UnitAddress VarChar(200) NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxUnit ON staging.Unit(BrCode, UnitCode);

CREATE TABLE IF NOT EXISTS staging.Center(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    CenterCode VarChar(7) NULL,
    CenterName VarChar(100) NULL,
    CenterAddress VarChar(200) NULL,
    MeetingDay BigInt NULL,
    Unit BigInt NULL,
    DateEstablished Date NULL
    AOID BigInt NULL
); 

CREATE UNIQUE INDEX IF NOT EXISTS idxCenter ON staging.Center(BrCode, CenterCode);

CREATE TABLE IF NOT EXISTS staging.Customer(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    CID BigInt NULL,
    CenterCode VarChar(200) NULL,
    Title BigInt NULL,
    LName VarChar(100) NULL,
    FName VarChar(100) NULL,
    MName VarChar(100) NULL,
    MaidenFName VarChar(100) NULL,
    MaidenLName VarChar(100) NULL,
    MaidenMName VarChar(100) NULL,
    Sex VarChar(1) NULL,
    BirthDate Date NULL,
    BirthPlace VarChar(100) NULL,
    CivilStatus BigInt NULL,
    CustType BigInt NULL,
    Remarks VarChar(200) NULL,
    Status BigInt NULL,
    Classification BigInt NULL,
    DepoType VarChar(200) NULL,
    SubClassification BigInt NULL,
    PledgeAmount Numeric(18,2) NULL,
    MutualAmount Numeric(18,2) NULL,
    PangarapAmount Numeric(18,2) NULL,
    KatuparanAmount Numeric(18,2) NULL,
    InsuranceAmount Numeric(18,2) NULL,
    AccPledge Numeric(18,2) NULL,
    AccMutual Numeric(18,2) NULL,
    AccPang Numeric(18,2) NULL,
    AccKatuparan Numeric(18,2) NULL,
    AccInsurance Numeric(18,2) NULL,
    LoanLimit Numeric(18,2) NULL,
    CreditLimit Numeric(18,2) NULL,
    DateRecognized Date NULL,
    DateResigned Date NULL,
    DateEntry Date NULL,
    GoldenLifeDate Date NULL,
    Restricted VarChar(200) NULL,
    Borrower VarChar(200) NULL,
    CoMaker VarChar(200) NULL,
    Guarantor VarChar(200) NULL,
    DOSRI BigInt NULL,
    IDCode1 BigInt NULL,
    IDNum1 VarChar(200) NULL,
    IDCode2 BigInt NULL,
    IDNum2 VarChar(200) NULL,
    Contact1 VarChar(200) NULL,
    Contact2 VarChar(200) NULL,
    Phone1 VarChar(200) NULL,
    Reffered1 VarChar(200) NULL,
    Reffered2 VarChar(200) NULL,
    Reffered3 VarChar(200) NULL,
    Education BigInt NULL,
    Validity1 Date NULL,
    Validity2 Date NULL,
    BusinessType BigInt NULL,
    AccountNumber VarChar(30) NULL,
    IIID BigInt NULL,
    Religion BigInt NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxCustomer ON staging.Customer(BrCode, CID);

CREATE TABLE IF NOT EXISTS staging.Addresses(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    CID BigInt NULL,
    SeqNum BigInt NULL,
    AddressDetails VarChar(200) NULL,
    Barangay VarChar(200) NULL,
    City VarChar(200) NULL,
    Province VarChar(200) NULL,
    Phone1 VarChar(200) NULL,
    Phone2 VarChar(200) NULL,
    Phone3 VarChar(200) NULL,
    Phone4 VarChar(200) NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxAddresses ON staging.Addresses(BrCode, SeqNum);

CREATE TABLE IF NOT EXISTS staging.LnMaster(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    CID BigInt NULL,
    Acc VarChar(22) NULL,
    AcctType BigInt NULL,
    DisbDate Date NULL,
    Principal Numeric(18,2) NULL,
    Interest Numeric(18,2) NULL,
    NetProceed Numeric(18,2) NULL,
    Gives BigInt NULL,
    Frequency BigInt NULL,
    AnnumDiv BigInt NULL,
    Prin Numeric(18,2) NULL,
    IntR Numeric(18,2) NULL,
    WaivedInt Numeric(18,2) NULL,
    WeeksPaid BigInt NULL,
    DoMaturity Date NULL,
    ConIntRate Numeric(18,2) NULL,
    Status VarChar(200) NULL,
    Cycle BigInt NULL,
    LNGrpCode BigInt NULL,
    Proff BigInt NULL,
    FundSource VarChar(200) NULL,
    DOSRI Bool NULL,
    LnCategory BigInt NULL,
    OpenDate Date NULL,
    LastTrnDate Date NULL,
    DisbBy VarChar(200) NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxLnMaster ON staging.LnMaster(BrCode, Acc);

CREATE TABLE IF NOT EXISTS staging.SaMaster(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    Acc VarChar(22) NULL,
    CID BigInt NULL,
    Type BigInt NULL,
    Balance Numeric(18,2) NULL,
    DoLastTrn Date NULL,
    DoStatus Date NULL,
    Dopen Date NULL,
    DoMaturity Date NULL,
    Status VarChar(200) NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxSaMaster ON staging.SaMaster(BrCode, Acc);

CREATE TABLE IF NOT EXISTS staging.TrnMaster(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    Acc VarChar(22) NULL,
    TrnDate Date NULL,
    Trn BigInt NULL,
    TrnType BigInt NULL,
    OrNo BigInt NULL,
    Prin Numeric(18,2) NULL,
    IntR Numeric(18,2) NULL,
    WaivedInt Numeric(18,2) NULL,
    RefNo VarChar(200) NULL,
    UserName VarChar(30) NULL,
    Particular VarChar(200) NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxTrnMaster ON staging.TrnMaster(BrCode, TrnDate, Trn);

CREATE TABLE IF NOT EXISTS staging.SaTrnMaster(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    Acc VarChar(22) NULL,
    TrnDate Date NULL,
    Trn BigInt NULL,
    TrnType BigInt NULL,
    OrNo BigInt NULL,
    TrnAmt Numeric(18,2) NULL,
    RefNo VarChar(200) NULL,
    Particular VarChar(200) NULL,
    TermId VarChar(200) NULL,
    UserName VarChar(30) NULL,
    PendApprove VarChar(1) NOT NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxSaTrnMaster ON staging.SaTrnMaster(BrCode, TrnDate, Trn);

CREATE TABLE IF NOT EXISTS staging.LoanInst(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    Acc VarChar(22) NULL,
    Dnum BigInt NULL,
    DueDate Date NULL,
    InstFlag BigInt NULL,
    DuePrin Numeric(18,2) NULL,
    DueInt Numeric(18,2) NULL,
    UpInt Numeric(18,2) NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxLoanInst ON staging.LoanInst(BrCode, Acc, Dnum);

CREATE TABLE IF NOT EXISTS staging.LnChrgData(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    Acc VarChar(22) NULL,
    ChrgCode BigInt NULL,
    RefAcc VarChar(22) NULL,
    ChrAmnt Numeric(18,2) NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxLnChrgData ON staging.LnChrgData(BrCode, Acc, ChrgCode, ChrgCode);

CREATE TABLE IF NOT EXISTS staging.CustAddInfoList(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    InfoCode BigInt NULL,
    InfoOrder VarChar(200) NULL,
    Title VarChar(200) NULL,
    InfoType VarChar(200) NULL,
    InfoLen BigInt NULL,
    InfoFormat VarChar(200) NULL,
    InputType BigInt NULL,
    InfoSource VarChar(200) NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxCustAddInfoList ON staging.CustAddInfoList(BrCode, InfoCode);

CREATE TABLE IF NOT EXISTS staging.CustAddInfoGroup(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    InfoGroup BigInt NULL,
    GroupTitle VarChar(200) NULL,
    Remarks VarChar(200) NULL,
    ReqOnEntry Bool NULL,
    ReqOnExit Bool NULL,
    Link2Loan BigInt NULL,
    Link2Save BigInt NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxCustAddInfoGroup ON staging.CustAddInfoGroup(BrCode, InfoGroup);

CREATE TABLE IF NOT EXISTS staging.CustAddInfoGroupNeed(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    InfoGroup BigInt NULL,
    InfoCode BigInt NULL,
    InfoProcess VarChar(200) NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxCustAddInfoGroupNeed ON staging.CustAddInfoGroupNeed(BrCode, InfoGroup, InfoCode);

CREATE TABLE IF NOT EXISTS staging.CustAddInfo(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    CID BigInt NULL,
    InfoDate Date NULL,
    InfoCode BigInt NULL,
    Info VarChar(200) NULL,
    InfoValue BigInt NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxCustAddInfo ON staging.CustAddInfo(BrCode, CID, InfoCode, InfoDate);

CREATE TABLE IF NOT EXISTS staging.MutualFund(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    CID BigInt NULL,
    OrNo BigInt NULL,
    TrnDate Date NULL,
    TrnType VarChar(200) NULL,
    TrnAmt Numeric(18,2) NULL,
    UserName VarChar(30) NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxMutualFund ON staging.MutualFund(BrCode, CID, OrNo);

CREATE TABLE IF NOT EXISTS staging.ReferencesDetails(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    ID BigInt NULL,
    RefID BigInt NULL,
    PurposeDescription VarChar(200) NULL,
    ParentID BigInt NULL,
    CodeID BigInt NULL,
    Stat BigInt NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxReferencesDetails ON staging.ReferencesDetails(BrCode, ID);

CREATE TABLE IF NOT EXISTS staging.CenterWorker(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    AOID BigInt NULL,
    Lname VarChar(100) NULL,
    FName VarChar(100) NULL,
    Mname VarChar(100) NULL,
    PhoneNumber VarChar(200) NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxCenterWorker ON staging.CenterWorker(BrCode, AOID);

CREATE TABLE IF NOT EXISTS staging.Writeoff(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    Acc VarChar(22) NULL,
    DisbDate Date NULL,
    Principal Numeric(18,2) NULL,
    Interest Numeric(18,2) NULL,
    BalPrin Numeric(18,2) NULL,
    BalInt Numeric(18,2) NULL,
    TrnDate Date NULL,
    AcctType VarChar(22) NULL,
    Print VarChar(1) NULL,
    PostedBy VarChar(200) NULL,
    VerifiedBy VarChar(200) NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxWriteoff ON staging.Writeoff(BrCode, Acc);

CREATE TABLE IF NOT EXISTS staging.Accounts(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    Acc VarChar(22) NULL,
    Title VarChar(200) NULL,
    Category BigInt NULL,
    Type VarChar(200) NULL,
    MainCD VarChar(200) NULL,
    Parent VarChar(200) NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxAccounts ON staging.Accounts(BrCode, Acc);

CREATE TABLE IF NOT EXISTS staging.JnlHeaders(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    Trn VarChar(17) NULL,
    TrnDate Date NULL,
    Particulars VarChar(300) NULL,
    UserName VarChar(30) NULL,
    Code BigInt NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxJnlHeaders ON staging.JnlHeaders(BrCode, Trn);

CREATE TABLE IF NOT EXISTS staging.JnlDetails(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    Acc VarChar(22) NULL,
    Trn VarChar(17) NULL,
    Series BigInt NULL,
    Debit Numeric(18,2) NULL,
    Credit Numeric(18,2) NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxJnlDetails ON staging.JnlDetails(BrCode, Trn, Acc);

CREATE TABLE IF NOT EXISTS staging.LedgerDetails(
    ModCtr BigInt NULL,
    BrCode VarChar(2) NULL,
    TrnDate Date NULL,
    Acc VarChar(22) NULL,
    Balance Numeric(18,2) NULL
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxLedgerDetails ON staging.LedgerDetails(BrCode, TrnDate, Acc);

CREATE TABLE IF NOT EXISTS staging.UsersList(
    ModCtr BigInt NULL,
    BrCode  varchar(2) NOT NULL,
    UserID  varchar(50) NOT NULL,
    AccessCode  numeric,
    LName varchar(50),
    Fname  varchar(50),
    MName varchar(1),
    DateHired date,
    BirthDay  date,
    DateGiven date,
    DateExpired date,
    Address varchar(200),
    Position  varchar(100),
    AreaCode  numeric,
    ManCode numeric,
    AddInfo varchar(1000),
    -- Passwd  bytea,
    Attempt smallint,
    DateLocked  date,
    Remarks varchar(1000),
    -- Picture bytea,
    isLoggedIn  bool,
    AccountExpirationDt date
    -- ComputerName  varchar(50),
    -- ValidateEdataBackup int,
CONSTRAINT pkUsersList PRIMARY KEY (BrCode, UserID)
); 
--CREATE UNIQUE INDEX IF NOT EXISTS idxUsersList ON staging.UsersList(BrCode, UserID);

  CREATE TABLE IF NOT EXISTS staging.MultiplePaymentReceipt(
    ModCtr BigInt NULL,
    BrCode  varchar(2) NOT NULL,
    ModAction  varchar(1) NOT NULL,
    TrnDate  Date NOT NULL,
    OrNo  numeric NOT NULL,
    CID numeric NOT NULL,
    PrNo numeric NOT NULL,
    UserName varchar(50) NOT NULL,
    TermId varchar(20) NOT NULL,
    AmtPaid numeric NOT NULL,
CONSTRAINT pkMultiplePaymentReceipt PRIMARY KEY (BrCode, OrNo)
); 

CREATE TABLE staging.InActiveCID(
  ModCtr BigInt NOT NULL,
  BrCode VarChar(2) NOT NULL,
  CID numeric NOT NULL,
  InActive bool NOT NULL,
  DateStart Date NOT NULL,
  DateEnd Date NOT NULL,
  UserId varchar(50) NOT NULL,
  DeactivatedBy varchar(50) NULL,
CONSTRAINT pkInActiveCID PRIMARY KEY (BrCode, CID, DateStart)
);   

CREATE TABLE IF NOT EXISTS staging.ReactivateWriteoff(
  ModCtr BigInt NOT NULL,
  BrCode VarChar(2) NOT NULL,
  ID numeric NOT NULL,
  CID numeric NOT NULL,
  DeactivateBy varchar(50) NULL,
  ReactivateBy varchar(50) NULL,
  Status SmallInt NULL,  
  StatusDate Date NULL,
CONSTRAINT pkReactivateWriteoff PRIMARY KEY (BrCode, ID)
);   

CREATE TABLE IF NOT EXISTS staging.LnBeneficiary(
  ModCtr BigInt NOT NULL,
  BrCode VarChar(2) NOT NULL,
  Acc varchar(22) NOT NULL,
  bDay date NOT NULL,
  Educ_Lvl varchar(5) NULL,
  Gender bool NULL,
  Last_Name varchar(100) NULL,
  First_Name varchar(100) NULL,
  Middle_Name varchar(100) NULL,
  Remarks varchar(200) NULL,
CONSTRAINT pkLnBeneficiary PRIMARY KEY (BrCode, Acc)
);   

-- CREATE UNIQUE INDEX IF NOT EXISTS idxMultiplePaymentReceipt ON staging.MultiplePaymentReceipt(BrCode, OrNo);