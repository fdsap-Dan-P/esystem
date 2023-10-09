
CREATE TABLE IF NOT EXISTS esystemdump.BranchList(
    BrCode varChar(2) NOT NULL,
    EbSysDate Date NOT NULL,
    RunState SmallInt NOT NULL,
    OrgAddress VarChar(200) NOT NULL,
    TaxInfo VarChar(300) NOT NULL,
    DefCity VarChar(50) NOT NULL,
    DefProvince VarChar(50) NOT NULL,
    DefCountry VarChar(50) NOT NULL,
    DefZip VarChar(50) NOT NULL,
    WaivableInt bool NOT NULL,
    DBVersion VarChar(10) NOT NULL,
    ESystemVer bytea NOT NULL,
    NewBrCode SmallInt NOT NULL,
    LastConnection Date NOT NULL,
    CONSTRAINT pkBranchList PRIMARY KEY (BrCode)
); 

CREATE TABLE IF NOT EXISTS esystemdump.ModifiedTable(
    BrCode varChar(2) NOT NULL,
    DumpTable varChar(100) NOT NULL,
    LastModCtr BigInt NOT NULL,
    CONSTRAINT pk_ModifiedTable PRIMARY KEY (BrCode, DumpTable)
); 

---------------------------------------------------------------------------
CREATE or REPLACE FUNCTION esystemdump.trgModified() returns trigger AS 
---------------------------------------------------------------------------
$$ 
DECLARE 
   tbl name ;--VarChar(255);
BEGIN
   tbl = TG_TABLE_NAME; --::regclass::text;
 --select TG_TABLE_NAME::varChar(255);
 INSERT INTO esystemdump.ModifiedTable(BrCode, DumpTable, LastModCtr)
   SELECT NEW.BrCode, tbl, NEW.ModCtr
     ON CONFLICT(BrCode, DumpTable) 
   DO UPDATE SET
    LastModCtr = EXCLUDED.LastModCtr;
 --  END IF;
RETURN NEW; 
END $$ Language plpgsql;
 

CREATE TABLE IF NOT EXISTS esystemdump.Area(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    AreaCode BigInt NOT NULL,
    Area VarChar(30) NULL,
    CONSTRAINT pk_Area PRIMARY KEY (BrCode, AreaCode, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.Unit(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    UnitCode BigInt NOT NULL,
    Unit VarChar(100) NULL,
    AreaCode BigInt NULL,
    FName VarChar(100) NULL,
    LName VarChar(100) NULL,
    MName VarChar(100) NULL,
    VatReg VarChar(100) NULL,
    UnitAddress VarChar(200) NULL,
    CONSTRAINT pk_Unit PRIMARY KEY (BrCode, UnitCode, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.Center(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    CenterCode VarChar(7) NOT NULL,
    CenterName VarChar(100) NULL,
    CenterAddress VarChar(200) NULL,
    MeetingDay BigInt NULL,
    Unit BigInt NULL,
    DateEstablished Date NULL,
    AOID BigInt NULL,
    
    CONSTRAINT pk_Center PRIMARY KEY (BrCode, CenterCode, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.Customer(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    CID BigInt NOT NULL,
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
    Religion BigInt NULL,
    CONSTRAINT pk_Customer PRIMARY KEY (BrCode, CID, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.Addresses(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    CID BigInt NOT NULL,
    SeqNum BigInt NOT NULL,
    AddressDetails VarChar(200) NULL,
    Barangay VarChar(200) NULL,
    City VarChar(200) NULL,
    Province VarChar(200) NULL,
    Phone1 VarChar(200) NULL,
    Phone2 VarChar(200) NULL,
    Phone3 VarChar(200) NULL,
    Phone4 VarChar(200) NULL,
    CONSTRAINT pk_Addresses PRIMARY KEY (BrCode, SeqNum, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.LnMaster(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    CID BigInt NOT NULL,
    Acc VarChar(22) NOT NULL,
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
    DisbBy VarChar(200) NULL,
    CONSTRAINT pk_LnMaster PRIMARY KEY (BrCode, Acc, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.SaMaster(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    Acc VarChar(22) NOT NULL,
    CID BigInt NOT NULL,
    Type BigInt NOT NULL,
    Balance Numeric(18,2) NULL,
    DoLastTrn Date NULL,
    DoStatus Date NULL,
    Dopen Date NULL,
    DoMaturity Date NULL,
    Status VarChar(200) NULL,
    CONSTRAINT pk_SaMaster PRIMARY KEY (BrCode, Acc, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.TrnMaster(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    Acc VarChar(22) NOT NULL,
    TrnDate Date NOT NULL,
    Trn BigInt NOT NULL,
    TrnType BigInt NULL,
    OrNo BigInt NULL,
    Prin Numeric(18,2) NOT NULL,
    IntR Numeric(18,2) NOT NULL,
    WaivedInt Numeric(18,2) NOT NULL,
    RefNo VarChar(200) NULL,
    UserName VarChar(30) NULL,
    Particular VarChar(200) NULL,
    CONSTRAINT pk_TrnMaster PRIMARY KEY (BrCode, TrnDate, Trn, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.SaTrnMaster(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    Acc VarChar(22) NOT NULL,
    TrnDate Date NOT NULL,
    Trn BigInt NOT NULL,
    TrnType BigInt NULL,
    OrNo BigInt NULL,
    TrnAmt Numeric(18,2) NULL,
    RefNo VarChar(200) NULL,
    Particular VarChar(200) NOT NULL,
    TermId VarChar(200) NOT NULL,
    UserName VarChar(30) NOT NULL,
    PendApprove VarChar(1) NOT NULL,
    CONSTRAINT pk_SaTrnMaster PRIMARY KEY (BrCode, TrnDate, Trn, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.LoanInst(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    Acc VarChar(22) NOT NULL,
    Dnum BigInt NOT NULL,
    DueDate Date NOT NULL,
    InstFlag BigInt NOT NULL,
    DuePrin Numeric(18,2) NOT NULL,
    DueInt Numeric(18,2) NOT NULL,
    UpInt Numeric(18,2) NOT NULL,
    CONSTRAINT pk_LoanInst PRIMARY KEY (BrCode, Acc, Dnum, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.LnChrgData(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    Acc VarChar(22) NOT NULL,
    ChrgCode BigInt NOT NULL,
    RefAcc VarChar(22) NULL,
    ChrAmnt Numeric(18,2) NOT NULL,
    CONSTRAINT pk_LnChrgData PRIMARY KEY (BrCode, Acc, ChrgCode, RefAcc, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.CustAddInfoList(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    InfoCode BigInt NOT NULL,
    InfoOrder VarChar(200) NOT NULL,
    Title VarChar(200) NOT NULL,
    InfoType VarChar(200) NOT NULL,
    InfoLen BigInt NOT NULL,
    InfoFormat VarChar(200) NOT NULL,
    InputType BigInt NOT NULL,
    InfoSource VarChar(200) NOT NULL,
    CONSTRAINT pk_CustAddInfoList PRIMARY KEY (BrCode, InfoCode, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.CustAddInfoGroup(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    InfoGroup BigInt NOT NULL,
    GroupTitle VarChar(200) NULL,
    Remarks VarChar(200) NULL,
    ReqOnEntry Bool NULL,
    ReqOnExit Bool NULL,
    Link2Loan BigInt NULL,
    Link2Save BigInt NULL,
    CONSTRAINT pk_CustAddInfoGroup PRIMARY KEY (BrCode, InfoGroup, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.CustAddInfoGroupNeed(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    InfoGroup BigInt NOT NULL,
    InfoCode BigInt NOT NULL,
    InfoProcess VarChar(200) NULL,
    CONSTRAINT pk_CustAddInfoGroupNeed PRIMARY KEY (BrCode, InfoGroup, InfoCode, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.CustAddInfo(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    CID BigInt NOT NULL,
    InfoDate Date NOT NULL,
    InfoCode BigInt NOT NULL,
    Info VarChar(200) NOT NULL,
    InfoValue BigInt NOT NULL,
    CONSTRAINT pk_CustAddInfo PRIMARY KEY (BrCode, CID, InfoCode, InfoDate, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.MutualFund(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    CID BigInt NOT NULL,
    OrNo BigInt NULL,
    TrnDate Date NOT NULL,
    TrnType VarChar(200) NULL,
    TrnAmt Numeric(18,2) NOT NULL,
    UserName VarChar(30) NULL,
    CONSTRAINT pk_MutualFund PRIMARY KEY (BrCode, CID, OrNo, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.ReferencesDetails(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    ID BigInt NOT NULL,
    RefID BigInt NOT NULL,
    PurposeDescription VarChar(200) NULL,
    ParentID BigInt NULL,
    CodeID BigInt NULL,
    Stat BigInt NULL,
    CONSTRAINT pk_ReferencesDetails PRIMARY KEY (BrCode, ID, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.CenterWorker(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    AOID BigInt NOT NULL,
    Lname VarChar(100) NULL,
    FName VarChar(100) NULL,
    Mname VarChar(100) NULL,
    PhoneNumber VarChar(200) NULL,
    CONSTRAINT pk_CenterWorker PRIMARY KEY (BrCode, AOID, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.Writeoff(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    Acc VarChar(22) NOT NULL,
    DisbDate Date NOT NULL,
    Principal Numeric(18,2) NOT NULL,
    Interest Numeric(18,2) NOT NULL,
    BalPrin Numeric(18,2) NOT NULL,
    BalInt Numeric(18,2) NOT NULL,
    TrnDate Date NOT NULL,
    AcctType VarChar(22) NOT NULL,
    Print VarChar(1) NULL,
    PostedBy VarChar(200) NULL,
    VerifiedBy VarChar(200) NULL,
    CONSTRAINT pk_Writeoff PRIMARY KEY (BrCode, Acc, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.Accounts(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    Acc VarChar(22) NOT NULL,
    Title VarChar(200) NOT NULL,
    Category BigInt NOT NULL,
    Type VarChar(200) NOT NULL,
    MainCD VarChar(200) NULL,
    Parent VarChar(200) NULL,
    CONSTRAINT pk_Accounts PRIMARY KEY (BrCode, Acc, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.JnlHeaders(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    Trn VarChar(17) NOT NULL,
    TrnDate Date NOT NULL,
    Particulars VarChar(200) NOT NULL,
    UserName VarChar(30) NULL,
    Code BigInt NOT NULL,
    CONSTRAINT pk_JnlHeaders PRIMARY KEY (BrCode, Trn, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.JnlDetails(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    Acc VarChar(22) NOT NULL,
    Trn VarChar(17) NOT NULL,
    Series BigInt NULL,
    Debit Numeric(18,2) NULL,
    Credit Numeric(18,2) NULL,
    CONSTRAINT pk_JnlDetails PRIMARY KEY (BrCode, Trn, Acc, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.LedgerDetails(
    ModCtr BigInt NOT NULL,
    BrCode VarChar(2) NOT NULL,
    ModAction VarChar(1) NOT NULL,
    TrnDate Date NOT NULL,
    Acc VarChar(22) NOT NULL,
    Balance Numeric(18,2) NOT NULL,
    CONSTRAINT pk_LedgerDetails PRIMARY KEY (BrCode, TrnDate, Acc, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.UsersList(
    ModCtr BigInt NOT NULL,
    BrCode  varchar(2) NOT NULL,
    UserID  varchar(50) NOT NULL,
    ModAction VarChar(1) NOT NULL,
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
    Passwd  bytea,
    Attempt smallint,
    DateLocked  date,
    Remarks varchar(1000),
    Picture bytea,
    isLoggedIn  bool,
    AccountExpirationDt date,
    ComputerName  varchar(50),
    ValidateEdataBackup int,
CONSTRAINT pkUsersList PRIMARY KEY (BrCode, UserID, ModAction)
); 


  CREATE TABLE IF NOT EXISTS esystemdump.MultiplePaymentReceipt(
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
CONSTRAINT pkMultiplePaymentReceipt PRIMARY KEY (BrCode, OrNo, ModAction)
); 

CREATE TABLE IF NOT EXISTS esystemdump.InActiveCID(
  ModCtr BigInt NOT NULL,
  BrCode VarChar(2) NOT NULL,
  ModAction  varchar(1) NOT NULL,
  CID numeric NOT NULL,
  InActive bool NOT NULL,
  DateStart Date NOT NULL,
  DateEnd Date NOT NULL,
  UserId varchar(50) NOT NULL,
  DeactivatedBy varchar(50) NULL,
CONSTRAINT pkInActiveCID PRIMARY KEY (BrCode, CID, DateStart, ModAction)
);   

CREATE TABLE IF NOT EXISTS esystemdump.ReactivateWriteoff(
  ModCtr BigInt NOT NULL,
  BrCode VarChar(2) NOT NULL,
  ModAction  varchar(1) NOT NULL,
  ID numeric NOT NULL,
  CID numeric NOT NULL,
  DeactivateBy varchar(50) NULL,
  ReactivateBy varchar(50) NULL,
  Status SmallInt NULL,  
  StatusDate Date NULL,
CONSTRAINT pkReactivateWriteoff PRIMARY KEY (BrCode, ID, ModAction)
);   

CREATE TABLE IF NOT EXISTS esystemdump.LnBeneficiary(
  ModCtr BigInt NOT NULL,
  BrCode VarChar(2) NOT NULL,
  ModAction  varchar(1) NOT NULL,
  Acc varchar(22) NOT NULL,
  bDay date NOT NULL,
  Educ_Lvl varchar(5) NULL,
  Gender bool NULL,
  Last_Name varchar(100) NULL,
  First_Name varchar(100) NULL,
  Middle_Name varchar(100) NULL,
  Remarks varchar(200) NULL,
CONSTRAINT pkLnBeneficiary PRIMARY KEY (BrCode, Acc, ModAction)
);   

CREATE TABLE esystemdump.ColSht(
  BrCode VarChar(2) NOT NULL,
  APPTYPE Int,
  Code int,
  Status Int,
  Acc varchar(22),
  CID int,
  UM varchar(50),
  ClientName varchar(55),
  CenterCode varchar(7),
  CenterName varchar(100),
  ManCode Int,
  Unit varchar(100),
  AreaCode Int,
  Area varchar(30),
  StaffName varchar(50),
  AcctType Int,
  AcctDesc varchar(100),
  DisbDate Date,
  DateStart Date,
  Maturity Date,
  Principal Numeric(18,2),
  Interest Numeric(18,2),
  Gives Int,
  BalPrin Numeric(18,2),
  BalInt Numeric(18,2),
  Amort Numeric(18,2),
  DuePrin Numeric(18,2),
  DueInt Numeric(18,2),
  LoanBal Numeric(18,2),
  SaveBal Numeric(18,2),
  WaivedInt Numeric(18,2),
  UnPaidCtr int,
  WrittenOff int,
  OrgName varchar(100),
  OrgAddress varchar(301),
  MeetingDate Date,
  MeetingDay smallint,
  SharesOfStock Numeric(18,2),
  DateEstablished Date,
  Classification int,
  WriteOff int,
CONSTRAINT pkColSht PRIMARY KEY (Acc)
);
CREATE INDEX IF NOT EXISTS idxColSht ON esystemdump.ColSht(BrCode);

---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgArea_Ins on esystemdump.Area;
CREATE TRIGGER trgArea_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.Area
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgUnit_Ins on esystemdump.Unit;
CREATE TRIGGER trgUnit_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.Unit
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgCenter_Ins on esystemdump.Center;
CREATE TRIGGER trgCenter_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.Center
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgCustomer_Ins on esystemdump.Customer;
CREATE TRIGGER trgCustomer_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.Customer
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgAddresses_Ins on esystemdump.Addresses;
CREATE TRIGGER trgAddresses_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.Addresses
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trglnMaster_Ins on esystemdump.lnMaster;
CREATE TRIGGER trglnMaster_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.lnMaster
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgsaMaster_Ins on esystemdump.saMaster;
CREATE TRIGGER trgsaMaster_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.saMaster
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgtrnMaster_Ins on esystemdump.trnMaster;
CREATE TRIGGER trgtrnMaster_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.trnMaster
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgsaTrnMaster_Ins on esystemdump.saTrnMaster;
CREATE TRIGGER trgsaTrnMaster_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.saTrnMaster
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgLoanInst_Ins on esystemdump.LoanInst;
CREATE TRIGGER trgLoanInst_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.LoanInst
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trglnChrgData_Ins on esystemdump.lnChrgData;
CREATE TRIGGER trglnChrgData_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.lnChrgData
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgCustAddInfoList_Ins on esystemdump.CustAddInfoList;
CREATE TRIGGER trgCustAddInfoList_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.CustAddInfoList
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgCustAddInfoGroup_Ins on esystemdump.CustAddInfoGroup;
CREATE TRIGGER trgCustAddInfoGroup_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.CustAddInfoGroup
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgCustAddInfoGroupNeed_Ins on esystemdump.CustAddInfoGroupNeed;
CREATE TRIGGER trgCustAddInfoGroupNeed_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.CustAddInfoGroupNeed
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgCustAddInfo_Ins on esystemdump.CustAddInfo;
CREATE TRIGGER trgCustAddInfo_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.CustAddInfo
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgMutualFund_Ins on esystemdump.MutualFund;
CREATE TRIGGER trgMutualFund_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.MutualFund
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgReferencesDetails_Ins on esystemdump.ReferencesDetails;
CREATE TRIGGER trgReferencesDetails_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.ReferencesDetails
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgCenterWorker_Ins on esystemdump.CenterWorker;
CREATE TRIGGER trgCenterWorker_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.CenterWorker
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgWriteoff_Ins on esystemdump.Writeoff;
CREATE TRIGGER trgWriteoff_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.Writeoff
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgAccounts_Ins on esystemdump.Accounts;
CREATE TRIGGER trgAccounts_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.Accounts
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgjnlHeaders_Ins on esystemdump.jnlHeaders;
CREATE TRIGGER trgjnlHeaders_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.jnlHeaders
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgjnlDetails_Ins on esystemdump.jnlDetails;
CREATE TRIGGER trgjnlDetails_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.jnlDetails
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgLedgerDetails_Ins on esystemdump.LedgerDetails;
CREATE TRIGGER trgLedgerDetails_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.LedgerDetails
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgUsersList_Ins on esystemdump.UsersList;
CREATE TRIGGER trgUsersList_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.UsersList
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgMultiplePaymentReceipt_Ins on esystemdump.MultiplePaymentReceipt;
CREATE TRIGGER trgMultiplePaymentReceipt_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.MultiplePaymentReceipt
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgInActiveCID_Ins on esystemdump.InActiveCID;
CREATE TRIGGER trgInActiveCID_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.InActiveCID
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgReactivateWriteoff_Ins on esystemdump.ReactivateWriteoff;
CREATE TRIGGER trgReactivateWriteoff_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.ReactivateWriteoff
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgLnBeneficiary_Ins on esystemdump.LnBeneficiary;
CREATE TRIGGER trgLnBeneficiary_Ins
---------------------------------------------------------------------------
    AFTER INSERT OR UPDATE ON esystemdump.LnBeneficiary
    FOR EACH ROW
    EXECUTE PROCEDURE esystemdump.trgModified();
    




-- CREATE UNIQUE INDEX IF NOT EXISTS idxMultiplePaymentReceipt ON staging.MultiplePaymentReceipt(BrCode, OrNo);
