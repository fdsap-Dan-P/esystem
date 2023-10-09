
/*
CREATE TABLE IF NOT EXISTS MapBranch(
  IIID Int4, 
  BrCode VarChar(4), 
  CONSTRAINT MapBranch_pkey PRIMARY KEY (IIID)
  );
CREATE UNIQUE INDEX IF NOT EXISTS idxMapBranch_Code ON public.MapBranch(BrCode);



-- // MapCenter2
CREATE TABLE IF NOT EXISTS MapCenter2(
  IIID Int4, 
  BrCode VarChar(4), 
  UnitIIID Int4,
  CenterCode VarChar(10),
  CONSTRAINT MapCenter2_pkey PRIMARY KEY (IIID)
  );
CREATE UNIQUE INDEX IF NOT EXISTS idxMapCenter2_Code ON public.MapCenter2(BrCode, CenterCode);

-- // MapCustomer2
CREATE TABLE IF NOT EXISTS MapCustomer2(
  IIID Int4, 
  BrCode VarChar(4), 
  CID Numeric,
  CONSTRAINT MapCustomer2_pkey PRIMARY KEY (BrCode, CID)
  );

CREATE UNIQUE INDEX IF NOT EXISTS idxMapCustomer2_Code ON public.MapCustomer2(IIID);

-- // MapUnit2
CREATE TABLE IF NOT EXISTS MapUnit2(
  IIID Int4, 
  BrCode VarChar(4), 
  UnitCode Numeric,
  CONSTRAINT MapUnit2_pkey PRIMARY KEY (IIID)
  );
CREATE UNIQUE INDEX IF NOT EXISTS idxMapUnit2_Code ON public.MapUnit2(BrCode, UnitCode);

-- // Fillup MapAccount2
CREATE TABLE IF NOT EXISTS MapAccount2(
  NewAcc VarChar(25),
  BrCode VarChar(4), 
  Acc VarChar(25),
  Archived Bool,
  CONSTRAINT MapAccount2_pkey PRIMARY KEY (BrCode, Acc));
  CREATE UNIQUE INDEX IF NOT EXISTS idxMapAccount2_Code ON public.MapAccount2(NewAcc);


-- // MapCenter
CREATE TABLE IF NOT EXISTS MapCenter(
  IIID Int4, 
  BrCode VarChar(4), 
  UnitIIID Int4,
  CenterCode VarChar(10),
  CONSTRAINT MapCenter_pkey PRIMARY KEY (IIID)
  );
CREATE UNIQUE INDEX IF NOT EXISTS idxMapCenter_Code ON public.MapCenter(BrCode, CenterCode);

-- // MapUnit
CREATE TABLE IF NOT EXISTS MapCustomer(
  IIID Int4, 
  BrCode VarChar(4), 
  CID Numeric,
  CONSTRAINT MapCustomer_pkey PRIMARY KEY (BrCode, CID)
  );

CREATE UNIQUE INDEX IF NOT EXISTS idxMapCustomer_Code ON public.MapCustomer(IIID);

-- // MapUnit
CREATE TABLE IF NOT EXISTS MapUnit(
  IIID Int4, 
  BrCode VarChar(4), 
  UnitCode Numeric,
  CONSTRAINT MapUnit_pkey PRIMARY KEY (IIID)
  );
CREATE UNIQUE INDEX IF NOT EXISTS idxMapUnit_Code ON public.MapUnit(BrCode, UnitCode);

-- // MapBranch
CREATE TABLE IF NOT EXISTS MapBranch(
  IIID Int4, 
  BrCode VarChar(4), 
  NewBrCode VarChar(4),
  CONSTRAINT MapBranch_pkey PRIMARY KEY (IIID)
  );
CREATE UNIQUE INDEX IF NOT EXISTS idxMapBranch_Code ON public.MapBranch(BrCode);

-- // Fillup MapAccount
CREATE TABLE IF NOT EXISTS MapAccount(
  NewAcc VarChar(25),
  BrCode VarChar(4), 
  Acc VarChar(25),
  Archived Bool,
  CONSTRAINT MapAccount_pkey PRIMARY KEY (BrCode, Acc));
  CREATE UNIQUE INDEX IF NOT EXISTS idxMapAccount_Code ON public.MapAccount(NewAcc);
  
  CREATE TABLE BrCodes (BrCode VarChar(4))
  create index unique BrCodes
*/
    
DELETE FROM BrCodes;

INSERT INTO BrCodes
SELECT * FROM (VALUES
  ('04'),
  ('I3'),
  ('L7'),
  ('M6'),
  ('N1'),
  ('O1'),
  ('P7'),
  ('W5')
) b(BrCode)
ON CONFLICT DO NOTHING;

   INSERT INTO BrCodes
SELECT distinct brcode FROM dat_customer 

ON CONFLICT DO NOTHING;


--// Get unmap Customer
  SELECT 
    c.OfficeIIID CenterIIID, c.OfficeName CenterName,
    u.OfficeIIID UnitIIID, u.OfficeName UnitName, 
    dc.*
  FROM dat_customer dc
  --inner join dat_samaster dl on dl.brcode = dc.brcode and dl.cid = dc.cid
  LEFT JOIN mapcustomer mc on dc.BrCode = mc.brCode and mc.cid = dc.cid
  LEFT JOIN mapCenter cen on cen.brcode = dc.brcode and cen.centercode = dc.center_code 
  LEFT JOIN Offices c on c.officeiiid = cen.IIID 
  LEFT JOIN Offices u on u.officeiiid = c.parentiiid
  WHERE mc.IIID is null

--/ Get not Map Center 
  SELECT distinct sc.*, u.IIID UnitIIID, dat_unit.unit 
  FROM dat_center sc 
  INNER JOIN dat_customer cus on cus.brcode = sc.brcode and cus.center_Code = sc.centercode
  LEFT JOIN dat_unit on dat_unit.brcode = sc.brcode and dat_unit.unitcode = sc.unitcode 
  LEFT JOIN mapcenter mc on mc.brcode = sc.brcode and mc.centercode = sc.centercode 
  LEFT JOIN mapunit u on u.BrCode = sc.brcode and u.unitcode = sc.unitcode 
  WHERE mc.IIID is null 
  
--// Test Duplicate Clients
  SELECT BrCode, CID
  FROM staging.stg_Customer sc
  GROUP BY BrCode, CID
  HAVING Count(CID) > 1

  --/ Get not Map Center
  SELECT 
    c.OfficeIIID CenterIIID, c.OfficeName CenterName,
    u.OfficeIIID UnitIIID, u.OfficeName UnitName, 
    dc.*
  FROM staging.stg_customer dc
  --inner join dat_samaster dl on dl.brcode = dc.brcode and dl.cid = dc.cid
  LEFT JOIN mapcustomer mc on dc.BrCode = mc.brCode and mc.cid = dc.cid
  LEFT JOIN mapCenter cen on cen.brcode = dc.brcode and cen.centercode = dc.center_code 
  LEFT JOIN Offices c on c.officeiiid = cen.IIID 
  LEFT JOIN Offices u on u.officeiiid = c.parentiiid
  WHERE mc.IIID is null
  
  
SELECT BrCode, CID, IIID
FROM
 (SELECT 
    COALESCE(sc.BrCode, mc.BrCode) BrCode, 
    COALESCE(sc.CID, mc.CID) CID, mc.IIID
  FROM 
   (SELECT sc.BrCode, CID 
    FROM staging.stg_Customer sc
    INNER JOIN BrCodes on BrCodes.BrCode = sc.BrCode 
    ) sc
  FULL JOIN 
   (SELECT sc.BrCode, CID, IIID
    FROM MapCustomer sc
    INNER JOIN BrCodes on BrCodes.BrCode = sc.BrCode 
   ) mc on sc.BrCode = mc.BrCode and sc.CID = mc.CID
  WHERE mc.IIID is NULL or sc.CID is NULL
 ) c



-- // Fillup MapBranch
INSERT INTO MapBranch(IIID, BrCode, NewBrCode)
  SELECT OfficeIIID, BrCode, NewBrCode
  FROM staging.stg_brcode sb
  WHERE officeiiid  is not null
  ON CONFLICT (BrCode) DO 
    UPDATE SET 
      IIID = EXCLUDED.IIID,
       NewBrCode = EXCLUDED.NewBrCode;

 -- Update Branch Staging
 INSERT INTO MapBranch(IIID,BrCode, NewBrCode)
 -- Copy to MapBranch
 SELECT OfficeIIID, BrCode, NewBrCode FROM Staging.stg_brcode WHERE OfficeIIID is NOT NULL 
 ON CONFLICT(BrCode) DO
 UPDATE SET IIID = EXCLUDED.IIID;
 
-- // Fillup MapUnit
INSERT INTO MapUnit(IIID, BrCode, UnitCode)
  SELECT DISTINCT o.officeiiid, m.BrCode, right(o.code,3)::Numeric UnitCode --, o.officename 
  FROM Offices o 
  INNER JOIN MapBranch m on m.IIID = o.ParentIIID
  INNER JOIN staging.stg_unit u on u.BrCode = m.BrCode and u.unitcode::VarChar(3) = right(o.code,3)
  WHERE lower(o.officetype) ='unit' and m.brcode = 'C1' -- order by unitcode
  ON CONFLICT (BrCode, UnitCode) 
  DO UPDATE SET 
    IIID = EXCLUDED.IIID;
    
 -- // Fillup MapCenter 
INSERT INTO MapCenter(IIID, BrCode, UnitIIID, CenterCode)
  SELECT 
    DISTINCT c.OfficeIIID CenterIIID,  sb.brcode,  u.OfficeIIID, sc.centercode --, c.Code
  FROM Offices c
  INNER JOIN Offices u on u.officeiiid = c.parentiiid 
  INNER JOIN Offices b on b.officeiiid = u.parentiiid 
  INNER JOIN staging.stg_brcode sb on sb.officeiiid = b.officeiiid 
  INNER JOIN staging.stg_center sc on 
    substring(c.code FROM (length(c.code) - position('-' in reverse(c.code))+2) FOR 10) = sc.centercode 
    and sc.brcode = sb.brcode 
  WHERE lower(c.OfficeType) = 'center' and sb.brcode in ('2F','M6','N1','W5','P7') -- and c.OfficeIIID = 9277565
  and c.Officeiiid not in (select iiid from mapcenter)  
  ON CONFLICT (BrCode, CenterCode) 
  DO UPDATE SET 
    IIID = EXCLUDED.IIID;
   
  -- // Update customer CenterCode -> only  when there is a change in center parent
  UPDATE customer set branch_code = cus.Branch_Code, unit_code = cus.unit_code 
  FROM
   (SELECT cus.cid, b.officeiiid Branch_Code, b.officename, u.officeiiid unit_code 
    FROM Customer cus
    INNER JOIN Offices c on c.officeiiid = cus.center_code 
    INNER JOIN Offices u on u.officeiiid = c.parentiiid 
    INNER JOIN Offices b on b.officeiiid = u.parentiiid 
    WHERE b.officeiiid <> cus.branch_code
    ) cus
  WHERE cus.cid = customer.cid 
  
    UPDATE customer set branch_code = cus.Branch_Code, unit_code = cus.unit_code 
  FROM
   (SELECT cus.cid, b.officeiiid Branch_Code, b.officename, u.officeiiid unit_code 
    FROM Customer cus
    INNER JOIN Offices c on c.officeiiid = cus.center_code 
    INNER JOIN Offices u on u.officeiiid = c.parentiiid 
    INNER JOIN Offices b on b.officeiiid = u.parentiiid 
    WHERE u.officeiiid <> cus.unit_code 
    ) cus
  WHERE cus.cid = customer.cid 
  
  --WHERE mc.IIID = Customer.CID and customer.center_code = -1
  
  --// Get Customer with no centers
  SELECT distinct cen.*
  FROM Customer 
  INNER JOIN mapCustomer mc on mc.IIID = Customer.CID 
  INNER JOIN staging.stg_customer c on c.BrCode = mc.BrCode and c.CID = mc.CID
  INNER JOIN staging.stg_Center cen on cen.brcode = c.brcode and cen.centercode = c.center_code 
  WHERE mc.IIID = Customer.CID and customer.center_code = -1
  
----------------------
-- SystemReferences
----------------------  
INSERT INTO SystemReferences (
  refid, reftypeid, reftyperid, reftype, reftypetitle, code, 
  title, shortname, parentid, parent, ownerrid, statusid, status, reftypeparentid, 
  reftypeparent, reftypetitleparent, "xml"
)
SELECT 
  (select max(refid)+1 from SystemReferences), r.reftypeid, r.reftyperid, r.reftype, r.reftypetitle, 0 code, 
  'N/A' title, r.shortname, r.parentid, r.parent, r.ownerrid, r.statusid, r.status, r.reftypeparentid, 
  r.reftypeparent, r.reftypetitleparent, r."xml"  
FROM SystemReferences r
LEFT JOIN SystemReferences p on lower(p.reftype) = 'typeofbusiness' and p.title = 'N/A'
WHERE lower(r.reftype) = 'typeofbusiness' and lower(r.title) = 'retail trading'
  and p.title is null
   
 /* 
  select sl.*
from accounts a 
inner join mapaccount ma on ma.newacc = a.acc
inner join stg_loaninst sl on sl.brcode = ma.brcode and sl.acc  = ma.acc
left join archivedloaninst h on a.acc = h.acc 
where apptype = 3 and h.acc is null and archived = true
  
  select sl.*
from accounts a 
inner join mapaccount ma on ma.newacc = a.acc
inner join stg_loaninst sl on sl.brcode = ma.brcode and sl.acc  = ma.acc
left join loaninst h on a.acc = h.acc 
where apptype = 3 and h.acc is null and archived = false

  select sl.*
from accounts a 
inner join mapaccount ma on ma.newacc = a.acc
inner join stg_trnmaster sl on sl.brcode = ma.brcode and sl.acc  = ma.acc
left join dailytransaction h on a.acc = h.acc 
where a.apptype = 3 and h.acc is null and archived = false
  
SELECT mc.brcode, mc.centercode, mc.iiid CenterIIID, 
  c.OfficeName CenterName, u.UnitCode, c.Parentiiid UnitIIID, c.officeparent UnitName
from mapcenter mc 
inner join offices c on c.officeiiid = mc.iiid
inner join mapunit u on u.iiid = c.parentIIID 
where mc.brcode in ('2F','M6','N1','W5','P7')

  --/ Insert customer mapping
INSERT INTO MapCustomer
SELECT IIID, BrCode, CID 
FROM mapcustomer ic 
WHERE BrCode = '2F'
ON CONFLICT (BrCode, CID) 
DO UPDATE SET 
  IIID = EXCLUDED.IIID;

  --/ Other way of insert
INSERT INTO MapCustomer
SELECT IIID, BrCode, CID 
FROM mapcustomer ic 
WHERE BrCode = '2F'
ON CONFLICT (IIID) 
DO UPDATE SET 
  BrCode = EXCLUDED.BrCode,
  CID = EXCLUDED.CID;
*/
----------------------
-- Customer
----------------------
INSERT INTO Customer (
  cid, cname, fname, mname, maidenfname, maidenlname, maidenmname, 
  dobirth, birthplace, sex, civilstatus, title, status, classification, 
  subclassification, doentry, dorecognized, doresigned, center_code, branch_code, unit_code,
  dosri, reffered1, remarks, Business, accountnumber)
SELECT
  CID, lName, FName, Mname, MaidenFName, MaidenLName, MaidenMName, 
  BirthDate, left(BirthPlace,30) BirthPlace, sex.RefID Sex, civil.refid CivilStatus,  title.refID Title, stat.refID Status, 
  COALESCE (min(cls.refID),0)  Classification, min(COALESCE(subcls.refID,-1)) subclassification, 
  EntryDate, RecognizeDate, ResignDate, 
  COALESCE(c.CenterIIID,-1) Center_Code, COALESCE(c.BranchIIID,-1) branch_Code, COALESCE(c.unitIIID,-1) Unit_Code,
  DOSRI, Reffered1, Remarks, min(bus.refID) Business, 
  AddCheckDigit(NewBrCode || RIGHT('00000000' || OldCID::VarChar(8),8))  AccountNumber
FROM
 (SELECT 
    b.NewBrCode, iiid.iiid CID, lName, FName, Mname, MaidenFName, MaidenLName, MaidenMName, 
    BirthDate, BirthPlace, CASE WHEN Sex = 'M' THEN 'Male' ELSE 'Female' END Sex, CivilStatus, 
    Title, CASE Status WHEN 0 THEN 'Active' WHEN 1 THEN 'InActive' WHEN 2 THEN 'Resigned' WHEN 4 THEN 'Balik CARD' ELSE '' END Status, 
    Classification, subclassification, EntryDate, RecognizeDate, ResignDate, Center_Code, 
    CASE WHEN DOSRI THEN 1 ELSE 0 END DOSRI, Reffered1, Remarks, 'N/A' BusinessType,
    b.IIID BranchIIID, cen.UnitIIID, cen.IIID CenterIIID, iiid.CID OldCID
  FROM staging.stg_Customer c
  INNER JOIN MapBranch b on c.BrCode = b.BrCode
  INNER JOIN MapCenter cen on cen.BrCode = c.BrCode and cen.CenterCode = c.center_code 
  -- INNER JOIN mapcustomer iiid on iiid.BrCode = c.BrCode and iiid.CID = c.CID
  INNER JOIN  MapCustomer iiid  on iiid.BrCode = c.BrCode and iiid.CID = c.CID and COALESCE(iiid.ctr,1)=1
  INNER JOIN BrCodes on brCodes.BrCode = c.BrCode
  --INNER JOIN CIDS on CIDS.cid = iiid.iiid
  /*
   (SELECT m.IIID, c.BrCode, c.CID 
    FROM MapCustomer m
    INNER JOIN mapcustomer c 
      on m.BrCode = c.BrCode and m.CID = c.CID
    --WHERE m.IIID <> c.IIID
    ) iiid  on iiid.BrCode = c.BrCode and iiid.CID = c.CID
    */
  --WHERE b.BrCode in ('2F','M6','N1','W5','P7')
    --where c.brcode = 'L7' and c.cid = 5566431
 ) c
LEFT JOIN SystemReferences civil on lower(civil.Title) = lower(c.CivilStatus) and lower(civil.RefType) = 'civilstatus'
LEFT JOIN SystemReferences sex on lower(sex.Title) = lower(c.Sex) and lower(sex.RefType) = 'gender'
LEFT JOIN SystemReferences title on lower(title.Title) = lower(c.title) and lower(title.RefType) = 'title'
LEFT JOIN SystemReferences stat on lower(stat.Title) = lower(c.status) and lower(stat.RefType) = 'customerstatus'
LEFT JOIN SystemReferences cls on lower(cls.Title) = lower(c.classification) and lower(cls.RefType) = 'classification'
LEFT JOIN SystemReferences subcls on lower(subcls.Title) = lower(c.subclassification) and lower(subcls.RefType) = 'sub-classification'
LEFT JOIN SystemReferences bus on lower(bus.Title) = lower(c.BusinessType) and lower(bus.RefType) = lower('typeofbusiness')

GROUP BY 
  CID, lName, FName, Mname, MaidenFName, MaidenLName, MaidenMName, 
  BirthDate, BirthPlace, sex.RefID, civil.refid,  title.refID, stat.refID, 
  EntryDate, RecognizeDate, ResignDate, c.centerIIID, c.unitIIID, c.branchIIID,
  DOSRI, Reffered1, Remarks, NewBrCode, OldCID
ON CONFLICT (CID) DO 
  UPDATE SET 
    cname = EXCLUDED.cname,
    fname = EXCLUDED.fname,
    mname = EXCLUDED.mname,
    maidenfname = EXCLUDED.maidenfname,
    maidenlname = EXCLUDED.maidenlname,
    maidenmname = EXCLUDED.maidenmname,
    dobirth = EXCLUDED.dobirth,
    birthplace = EXCLUDED.birthplace,
    sex = EXCLUDED.sex,
    civilstatus = EXCLUDED.civilstatus,
    title = EXCLUDED.title,
    status = EXCLUDED.status,
    classification = EXCLUDED.classification,
    subclassification = EXCLUDED.subclassification,
    doentry = EXCLUDED.doentry,
    dorecognized = EXCLUDED.dorecognized,
    doresigned = EXCLUDED.doresigned,
    center_code = EXCLUDED.center_code,
    branch_code = EXCLUDED.branch_code,
    unit_code = EXCLUDED.unit_code,
    dosri = EXCLUDED.dosri,
    reffered1 = EXCLUDED.reffered1,
    remarks = EXCLUDED.remarks,
    business = EXCLUDED.business,
    accountnumber = EXCLUDED.accountnumber
  ;

-- get not inserted Customer
SELECT cen.*,mc.*,sc.* 
FROM staging.stg_Customer sc
LEFT JOIN mapCustomer mc on sc.BrCode = mc.BrCode and sc.cid = mc.cid
LEFT JOIN customer c on c.cid = mc.iiid
LEFT JOIN mapcenter cen  on cen.brcode  = sc.brcode and cen.centercode =sc.center_code 
INNER JOIN BrCodes on brCodes.BrCode = sc.BrCode
WHERE c.cid is null 

select *
FROM staging.stg_lnmaster ss 
INNER JOIN BrCodes on brCodes.BrCode = ss.BrCode
left join mapaccount ma on ss.brcode = ma.brcode and ss.acc = ma.acc 
where ma.newacc is null 

/*
insert into mapaccount 
SELECT a.newacc, a.brcode, oldacc, true
from staging.stg_lnmaster ss
INNER JOIN 
(select acc newAcc, substring(Alternateacc,3,17) oldAcc, left(AlternateAcc,2) BrCode  
 FROM accounts
 WHERE left(AlternateAcc,2) in ('2F','M6','N1','W5','P7')
) a on a.OldAcc = ss.Acc and a.BrCode = ss.BrCode
where a.newAcc is null and ss.brcode in ('2F','M6','N1','W5','P7')
*/

-- // INSERT Account Mapping FROM LOANS
INSERT INTO MapAccount(BrCode, Acc, Archived)
SELECT c.BrCode, Acc, abs(Principal-Prin+Interest-IntR) = 0
FROM staging.stg_lnmaster c
INNER JOIN MapBranch b on c.BrCode = b.BrCode
--WHERE BrCode in ('2F','M6','N1','W5','P7')
ON CONFLICT (BrCode, Acc) 
  DO UPDATE 
    SET Archived = EXCLUDED.Archived;

-- // INSERT Account Mapping FROM Savings
INSERT INTO MapAccount(BrCode, Acc, Archived)
SELECT c.BrCode, Acc, abs(Balance) = 0 and dateOpen < '2021/01/01'::date
FROM staging.stg_saMaster c
INNER JOIN BrCodes on brCodes.BrCode = c.BrCode
INNER JOIN MapBranch b on c.BrCode = b.BrCode
-- WHERE BrCode in ('2F','M6','N1','W5','P7')
ON CONFLICT (BrCode, Acc) 
  DO UPDATE 
    SET Archived = EXCLUDED.Archived;

-- // INSERT Account Mapping FROM Mutual_Fund
INSERT INTO MapAccount(BrCode, Acc, Archived)
SELECT BrCode, CID::VarChar(10), TRUE
FROM
 (SELECT cus.BrCode, cus.CID
  FROM staging.stg_mutual_fund cus 
  INNER JOIN BrCodes on brCodes.BrCode = cus.BrCode
  INNER JOIN MapBranch b on cus.BrCode = b.BrCode
  GROUP BY cus.BrCode, cid) a  
ON CONFLICT (BrCode, Acc) 
  DO UPDATE 
    SET Archived = EXCLUDED.Archived;
   

-- //  TAG All Current Account to Archive=FALSE UPDATE Account Archived 
UPDATE MapAccount
SET Archived = FALSE 
WHERE Archived AND 
EXISTS 
 (SELECT Acc FROM staging.stg_trnmaster t
  WHERE MapAccount.Acc = t.Acc and MapAccount.BrCode = t.BrCode
   and trnDate > '2020-12-31'::date);

 --- next
UPDATE MapAccount
SET Archived = FALSE 
WHERE Archived AND 
EXISTS 
 (SELECT CID::VarChar(10) 
  FROM staging.stg_mutual_fund t
  WHERE MapAccount.Acc = CID::VarChar(10) and MapAccount.BrCode = t.BrCode
   and trnDate > '2020-12-31'::date);
 
   -- // UPDATE Loan Account new Acc
UPDATE MapAccount
  SET newAcc = NewAcc(a.AcctType::varchar(10), lPad(y.module_code, 2, '0') || lPad(y.code, 2, '0') || '-')
  FROM
    staging.stg_lnMaster a 
  INNER JOIN AcctParms y on y.AcctType = a.AcctType
  WHERE MapAccount.newAcc is Null and a.BrCode = MapAccount.BrCode  and a.Acc = MapAccount.Acc;
  
  -- // UPDATE Savings Account new Acc
  UPDATE MapAccount
  SET newAcc = NewAcc(60::varchar(10), lPad(y.module_code, 2, '0') || lPad(y.code, 2, '0') || '-')
  FROM
    staging.stg_saMaster a 
  INNER JOIN AcctParms y on y.AcctType = 60
  WHERE MapAccount.newAcc is Null and a.BrCode = MapAccount.BrCode  and a.Acc = MapAccount.Acc;

  -- // UPDATE MBA Account new Acc
   UPDATE MapAccount
  SET newAcc = NewAcc(y.AcctType::varchar(10), lPad(y.module_code, 2, '0') || lPad(y.code, 2, '0') || '-')
  FROM
    mapcustomer a 
  INNER JOIN AcctParms y on y.AcctType = 201
  WHERE MapAccount.newAcc is Null and a.BrCode = MapAccount.BrCode  and a.CID::VarChar(10) = MapAccount.Acc;
  
  --// Check Account without new Acc
  SELECT * from mapaccount where newacc is null
  
 -- //Insert Loan Accounts
INSERT INTO Accounts(
  CID, Acc, 
  AlternateAcc, AppType, AcctType, AccDesc, dOpen, doMaturity, Term, WeeksPaid, 
  Status, Principal, Interest, "others", Discounted, NetProceed, Balance, 
  Prin, Intr, Oth, Penalty, Waivedint, DisbBy, ApprovBy, "cycle", Frequency, 
  AnnumDiv, lnGrpCode, Proff, Fundsource, ConIntRate, AmortCond, AmortCondValue, 
  Classification_Code, Classification_Type, Remarks, LumpSum, LoanID, LastTransactionDt
 )
SELECT 
  IIID.IIID CID, 
  ma.NewAcc Acc,
  a.BrCode || a.Acc AlternateAcc, y.AppType, a.AcctType, y.AcctDesc AccDesc, a.DisbDate dOpen, a.doMaturity, a.Gives Term, a.WeeksPaid, 
  a.Status, a.Principal, a.Interest, 0 "others", 0 Discounted, 0 NetProceed, 
  a.Principal + a.Interest - a.Prin - a.IntR - a.WaivedInt Balance, 
  a.Prin, a.Intr, 0 Oth, 0 Penalty, a.WaivedInt, '' DisbBy, '' ApprovBy, a."cycle", a.Frequency, 
  a.AnnumDiv, a.lnGrpCode, a.Proff, 0 Fundsource, a.ConIntRate, 0 AmortCond, 0 AmortCondValue, 
  0 Classification_Code, 0 Classification_Type, '' Remarks, 0 LumpSum, 0 LoanID, NULL LastTransactionDt
FROM staging.stg_lnMaster a
INNER JOIN BrCodes on brCodes.BrCode = a.BrCode
INNER JOIN MapAccount ma on ma.BrCode = a.BrCode and ma.Acc = a.Acc 
INNER JOIN AcctParms y on y.AcctType = a.AcctType
INNER JOIN mapcustomer iiid on iiid.BrCode = a.BrCode and iiid.CID = a.CID
INNER JOIN MapBranch b on a.BrCode = b.BrCode
-- WHERE a.BrCode in ('2F','M6','N1','W5','P7')
ON CONFLICT(Acc) DO NOTHING;

--// Insert accounts from Savings
INSERT INTO Accounts(
  CID, Acc, AlternateAcc, AppType, AcctType, AccDesc, dOpen, doMaturity, Term, WeeksPaid, 
  Status, Principal, Interest, "others", Discounted, NetProceed, Balance, 
  Prin, Intr, Oth, Penalty, Waivedint, DisbBy, ApprovBy, "cycle", Frequency, 
  AnnumDiv, lnGrpCode, Proff, Fundsource, ConIntRate, AmortCond, AmortCondValue, 
  Classification_Code, Classification_Type, Remarks, LumpSum, LoanID, LastTransactionDt
 )
SELECT 
  iiid.iiid CID, 
  ma.NewAcc Acc,
  a.BrCode || a.Acc AlternateAcc,  0 AppType, 60 AcctType, 'Pledge Account' AccDesc, dateopen dOpen, dateopen::date doMaturity, 0 Term, 0 WeeksPaid, 
  CASE WHEN a.Balance=0 THEN 99 ELSE 10 END Status, cus.Pledge_Amount Principal, 0 Interest, 0 "others", 0 Discounted, 0 NetProceed, 
  a.Balance, 
  cus.accpledge Prin, 0 Intr, 0 Oth, 0 Penalty, 0 WaivedInt, '' DisbBy, '' ApprovBy, 0 "cycle", 0 Frequency, 
  0 AnnumDiv, 0 lnGrpCode, 0 Proff, 0 Fundsource, 0 ConIntRate, 0 AmortCond, 0 AmortCondValue, 
  0 Classification_Code, 0 Classification_Type, '' Remarks, 0 LumpSum, 0 LoanID, lasttrndate LastTransactionDt
FROM staging.stg_saMaster a
INNER JOIN BrCodes on brCodes.BrCode = a.BrCode
INNER JOIN MapAccount ma on ma.BrCode = a.BrCode and ma.Acc = a.Acc 
INNER JOIN mapcustomer iiid on iiid.BrCode = a.BrCode and iiid.CID = a.CID
INNER JOIN staging.stg_Customer cus on cus.BrCode = iiid.BrCode and cus.CID = iiid.CID
--where a.BrCode in ('2F','M6','N1','W5','P7')
ON CONFLICT(Acc) DO UPDATE SET
  Acc = EXCLUDED.Acc,
  Intr = EXCLUDED.Intr,
  CID =  EXCLUDED.CID,
  Principal = EXCLUDED.Principal,
  Prin = EXCLUDED.Prin;

-- INSERT Accounts for MBA
INSERT INTO Accounts(
  CID, Acc, AlternateAcc, AppType, AcctType, AccDesc, dOpen, doMaturity, Term, WeeksPaid, 
  Status, Principal, Interest, "others", Discounted, NetProceed, Balance, 
  Prin, Intr, Oth, Penalty, Waivedint, DisbBy, ApprovBy, "cycle", Frequency, 
  AnnumDiv, lnGrpCode, Proff, Fundsource, ConIntRate, AmortCond, AmortCondValue, 
  Classification_Code, Classification_Type, Remarks, LumpSum, LoanID, LastTransactionDt
 )
 SELECT 
  iiid.iiid CID, 
  ma.NewAcc Acc,
  cus.BrCode  || '-' || cus.CID::VarChar(10) AlternateAcc, y.AppType,  y.AcctType, y.AcctDesc AccDesc, cus.RecognizeDate dOpen, 
  cus.RecognizeDate doMaturity, 0 Term, 0 WeeksPaid, 
  10 Status, cus.mutual_amount Principal, 0 Interest, 0 "others", 0 Discounted, 0 NetProceed, 
  0 Balance, 
  cus.accmutual Prin, 0 Intr, 0 Oth, 0 Penalty, 0 WaivedInt, '' DisbBy, '' ApprovBy, 0 "cycle", 0 Frequency, 
  0 AnnumDiv, 0 lnGrpCode, 0 Proff, 0 Fundsource, 0 ConIntRate, 0 AmortCond, 0 AmortCondValue, 
  0 Classification_Code, 0 Classification_Type, '' Remarks, 0 LumpSum, 0 LoanID, cus.resignDate LastTransactionDt  
FROM staging.stg_Customer cus 
INNER JOIN BrCodes on brCodes.BrCode = cus.BrCode
INNER JOIN MapAccount ma on ma.BrCode = cus.BrCode and ma.Acc = cus.CID::VarChar(10) 
INNER JOIN MapCustomer iiid on iiid.BrCode = cus.BrCode and iiid.CID = cus.CID
INNER JOIN AcctParms y on y.AcctType = 201
-- WHERE  cus.BrCode in ('2F','M6','N1','W5','P7')
--WHERE cus.mutual_amount > 0 and cus.Status in (0,1,4)
 -- and cus.CID in (SELECT cid from staging.stg_samaster sa where sa.balance > 0)
   --and cus.cid = 1536891
ON CONFLICT(Acc) DO UPDATE SET
  --Acc = EXCLUDED.Acc,
  Intr = EXCLUDED.Intr,
  CID  = EXCLUDED.CID,
  Principal = EXCLUDED.Principal,
  Prin = EXCLUDED.Prin;

----------------------
-- lnChrgData for previous balance
----------------------  
INSERT INTO lnChrgData(
  Acc, Chd, ChrgCode, ChrDesc, ChrAmnt, ChrBal, refacc) 
SELECT m.Acc Acc, Chd, ChrgCode, ChrDesc, Sum(ChrAmnt) ChrAmnt, Sum(ChrBal) ChrBal, m2.Acc refacc
FROM
  (SELECT
    c.BrCode, Acc, 0 Chd, ChrgCode, 
    'Previous Loan' ChrDesc, ChrAmnt, ChrAmnt ChrBal, refacc
  FROM staging.stg_lnChrgData c
  INNER JOIN BrCodes on brCodes.BrCode = c.BrCode
  WHERE chrgcode = 18 --and BrCode in ('2F','M6','N1','W5','P7')
  ) a
INNER JOIN MapAccount m on m.BrCode = a.BrCode and m.Acc = a.Acc
INNER JOIN MapAccount m2 on m2.BrCode = a.BrCode and m2.Acc = a.refAcc
--LEFT JOIN Accounts m on m.AlternateAcc = a.BrCode || a.Acc
--LEFT JOIN Accounts m2 on m2.AlternateAcc = a.BrCode || a.refAcc
GROUP BY m.Acc, Chd, ChrgCode, ChrDesc, m2.Acc 
ON CONFLICT DO NOTHING;

----------------------
-- lnChrgData not included in previous balance
----------------------  
INSERT INTO lnChrgData(
  Acc, Chd, ChrgCode, ChrDesc, ChrAmnt, ChrBal, refacc)
SELECT m.Acc Acc, Chd, ChrgCode, ChrDesc, Sum(ChrAmnt) ChrAmnt, Sum(ChrBal) ChrBal, NULL refacc
FROM
  (SELECT
    c.BrCode, Acc, 0 Chd, ChrgCode, 
    CASE ChrgCode 
    WHEN 11 THEN 'Retention for Pledge Savings'
    WHEN 14 THEN 'Service Fee'
    WHEN 18 THEN 'Previous Loan'
    WHEN 16 THEN 'LRF' ELSE '' END ChrDesc, ChrAmnt, ChrAmnt ChrBal, refacc
  FROM staging.stg_lnChrgData c
  INNER JOIN BrCodes on brCodes.BrCode = c.BrCode
  WHERE chrgcode <> 18 -- and brCode in ('2F','M6','N1','W5','P7')
  ) a
INNER JOIN MapAccount m on m.BrCode = a.BrCode and m.Acc = a.Acc  
GROUP BY m.Acc, Chd, ChrgCode, ChrDesc
ON CONFLICT DO NOTHING;
 
 
--// Insert writeoff
INSERT INTO Writeoff(
  Acc, disbdate, principal, interest, balprin, balint, trndate, accttype, print, postedby, verifiedby, termid) 
  SELECT m.Acc, a.disbdate, a.Principal, a.Interest, a.Balprin, a.Balint, a.trndate, a.accttype, a.print, 
     a.postedby, a.verifiedby, a.termid
  FROM staging.stg_Writeoff a
  INNER JOIN BrCodes on brCodes.BrCode = a.BrCode
  INNER JOIN MapAccount m on m.BrCode = a.BrCode and m.Acc = a.Acc
  --LEFT JOIN Accounts m on m.AlternateAcc = a.BrCode || a.Acc
  WHERE NOT EXISTS
    (SELECT w.Acc FROM Writeoff w WHERE w.Acc = m.Acc)
  ON CONFLICT DO NOTHING;

INSERT INTO ArchivedLoanInst(dNum, Acc, DueDate, Prin, IntR, UpInt)
  SELECT dNum, m.NewAcc, DueDate, Prin, IntR, UpInt
  FROM staging.stg_loaninst i
  INNER JOIN BrCodes on brCodes.BrCode = i.BrCode
  INNER JOIN MapAccount m on i.BrCode = m.BrCode and i.Acc = m.Acc
  WHERE m.Archived = True 
  ON CONFLICT (Acc, dNum) DO NOTHING;   
 
/*
select distinct ma.brcode
FROM mapAccount ma 
inner join accounts a on a.acc = ma.newacc 
left join loaninst i on i.acc = ma.newacc
where i.acc is null and ma.archived = false and a.apptype = 3

insert into accs
select ma.newacc
FROM mapAccount ma 
inner join accounts a on a.acc = ma.newacc 
left join loaninst i on i.acc = ma.newacc
where i.acc is null and ma.archived = false and a.apptype = 3

*/

INSERT INTO LoanInst( dnum, acc, duedate, instflag, prin, intr, oth, penalty, endbal, endint, endoth, instpd, penpd, carval, upint, servfee, pledgeamort)
  SELECT dnum, m.NewAcc, duedate, 0 instflag, prin, intr, 0 oth, 0 penalty, 0 endbal, 0 endint, 0 endoth, 0 instpd, 0 penpd, 0 carval, upint, 0 servfee, 0 pledgeamort 
  FROM staging.stg_loaninst i
  INNER JOIN BrCodes on brCodes.BrCode = i.BrCode
  INNER JOIN MapAccount m on i.BrCode = m.BrCode and i.Acc = m.Acc
  --INNER JOIN accs on m.newacc = accs.acc
 -- WHERE m.Archived = false and i.BrCode in ('2F','M6','N1','W5','P7')
  ON CONFLICT (Acc, dNum) DO NOTHING; 
    
-- // Update LoanInst Balances
UPDATE LoanInst SET 
  EndBal = a.EndBal,
  InstPD = a.InstPD,
  CarVal = a.CarVal,
  endint = a.endint,
  instflag = CASE WHEN a.InstPD = a.Prin + a.IntR THEN 9 ELSE 0 END 
FROM
  (SELECT
    dnum, i.Acc, duedate, instflag,
    prin,intr, endbal,endint, 
    CASE WHEN EndBal+EndInt+Prin+Intr-Bal > Prin+Intr THEN Prin+Intr
         WHEN EndBal+EndInt+Prin+Intr-Bal < 0 THEN 0 
         ELSE EndBal+EndInt+Prin+Intr-Bal END instpd, 
    EndBal-(Interest-EndInt)+SupInt carval
    FROM 
     (SELECT 
        l.Acc, i.dNum, i.duedate, i.instflag, i.prin, i.intr,
        l.Principal-Sum(s.prin) endbal,l.Interest-Sum(s.Intr) endint, 
        Sum(s.upint) supint, l.Interest,
        l.Principal + l.Interest - l.Prin - l.Intr Bal
      FROM LoanInst i
      INNER JOIN Accounts l on i.Acc = l.Acc
      INNER JOIN LoanInst s on s.Acc = l.Acc and s.dNum <= i.dNum
      INNER JOIN MapAccount a on a.NewAcc = i.Acc
      INNER JOIN BrCodes on brCodes.BrCode = a.BrCode
      --WHERE a.BrCode in ('2F','M6','N1','W5','P7')
      GROUP BY 
        l.Acc, i.dnum, i.acc, i.duedate, i.instflag, i.prin, i.intr, l.Prin, l.Intr,
        l.Principal, l.Interest, i.upint
      ) i 
   ) a  
WHERE a.Acc = LoanInst.Acc and a.dNum = LoanInst.dNum 

 