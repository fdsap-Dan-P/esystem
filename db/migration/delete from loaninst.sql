delete from loaninst 
where acc in
(select a.Acc 
 from account.mapcustomer2 c
 inner join account.accounts a on c.CID = a.CID)


ArchivedTransaction

----------------------
-- Loans/Savings Transaction
----------------------
INSERT INTO DailyTransaction (
     cid, acc, trndate, trntype, trn, orno, prin, intr, oth, 
     waivedint, trnamt, balance, pendapprove, trnmnem_cd, username, 
     termid, refno, trndesc, particulars, apptype, "time", unitcode,
     AlternateTrn
  )
SELECT 
  m.CID cid, m.acc, t.trndate, t.trntype, t.trn, t.orno, 
  CASE WHEN trnType = 3400 THEN 0 ELSE t.prin END Prin, 
  CASE WHEN trnType = 3400 THEN 0 ELSE t.Intr END Intr, 0 oth, 
  t.waivedint, t.Prin+t.Intr TrnAmt, 0 Balance, 1 pendapprove, 0 trnmnem_cd, '' username, 
  '' termid, CASE WHEN length(t.refno) < 8 and t.refno != '' and t.refno ~ '^[0-9]*$'  THEN t.refno::Numeric ELSE 0 END refno, 
  t.trn trndesc, Left(COALESCE(t.particular,''),100) particular, CASE WHEN m.AcctType = 60 THEN 0 ELSE 1 END apptype, trndate "time", c.unit_code unitcode,
  CASE WHEN m.apptype = 3 THEN 'L' ELSE 'S' END || t.BrCode || to_char(t.TrnDate, 'YYYYMMDD') || to_char(trn,'0000000000') AlternateTrn
FROM staging.stg_trnMaster t
INNER JOIN Account.MapAccount ma on ma.BrCode = t.BrCode and ma.Acc = t.Acc
INNER JOIN Account.Accounts m on ma.NewAcc = m.Acc
--INNER JOIN staging.stg_lnMaster l on l.BrCode = t.BrCode and l.Acc = t.Acc
--LEFT JOIN staging.INAI_CustomerMap iiid on iiid.BrCode = t.BrCode and iiid.CID = t.CID
INNER JOIN account.Customer c on m.CID = c.CID 
WHERE ma.Archived = false and t.BrCode in ('2F','M6','N1','W5','P7')
ON CONFLICT(Acc, TrnDate, Trn) DO UPDATE SET
  Prin = EXCLUDED.Prin,
  Intr = EXCLUDED.Intr
;

INSERT INTO ArchivedTransaction (
     cid, acc, trndate, trntype, trn, orno, prin, intr, oth, 
     waivedint, trnamt, balance, pendapprove, trnmnem_cd, username, 
     termid, refno, trndesc, particulars, apptype, "time", unitcode,
     AlternateTrn
  )
SELECT 
  m.CID cid, m.acc, t.trndate, t.trntype, t.trn, t.orno, 
  CASE WHEN trnType = 3400 THEN 0 ELSE t.prin END Prin, 
  CASE WHEN trnType = 3400 THEN 0 ELSE t.Intr END Intr, 0 oth, 
  t.waivedint, t.Prin+t.Intr TrnAmt, 0 Balance, 1 pendapprove, 0 trnmnem_cd, '' username, 
  '' termid, CASE WHEN length(t.refno) < 8 and t.refno != '' and t.refno ~ '^[0-9]*$'  THEN t.refno::Numeric ELSE 0 END refno, 
  t.trn trndesc, Left(COALESCE(t.particular,''),100) particular, CASE WHEN m.AcctType = 60 THEN 0 ELSE 1 END apptype, trndate "time", c.unit_code unitcode,
  CASE WHEN m.apptype = 3 THEN 'L' ELSE 'S' END || t.BrCode || to_char(t.TrnDate, 'YYYYMMDD') || to_char(trn,'0000000000') AlternateTrn
FROM staging.stg_trnMaster t
INNER JOIN Account.MapAccount ma on ma.BrCode = t.BrCode and ma.Acc = t.Acc
INNER JOIN Account.Accounts m on ma.NewAcc = m.Acc
--INNER JOIN staging.stg_lnMaster l on l.BrCode = t.BrCode and l.Acc = t.Acc
--LEFT JOIN staging.INAI_CustomerMap iiid on iiid.BrCode = t.BrCode and iiid.CID = t.CID
INNER JOIN account.Customer c on m.CID = c.CID 
WHERE ma.Archived = TRUE and t.BrCode in ('2F','M6','N1','W5','P7')
ON CONFLICT(Acc, TrnDate, Trn) DO NOTHING ;

----------------------
-- MBA
----------------------
INSERT INTO DailyTransaction (
     cid, acc, trndate, trntype, trn, orno, prin, intr, oth, 
     waivedint, trnamt, balance, pendapprove, trnmnem_cd, username, 
     termid, refno, trndesc, particulars, apptype, "time", unitcode,
     AlternateTrn
  )
SELECT 
  acc.CID, acc.acc, t.trndate, 0 trntype, t.orno trn, t.orno, 
  t.TrnAmt Prin, 0 Intr, 0 oth, 
  0 waivedint, t.TrnAmt, 0 Balance, 1 pendapprove, 0 trnmnem_cd, '' username, 
  '' termid, 0 refno, 
  '' trndesc, '' particular, 2 apptype, trndate "time", cus.unit_code unitcode,
  'M' || t.BrCode || to_char(t.TrnDate, 'YYYYMMDD') || to_char(ORNO,'0000000000') AlternateTrn
FROM staging.stg_Mutual_Fund t
INNER JOIN Account.MapAccount ma on ma.BrCode = t.BrCode and ma.Acc = t.CID::varchar(10)
INNER JOIN Account.Accounts acc on acc.Acc = ma.NewAcc
INNER JOIN Account.Customer cus on cus.CID = Acc.CID
WHERE ma.Archived = FALSE t.BrCode in ('2F','M6','N1','W5','P7')
ON CONFLICT(Acc, TrnDate, Trn) DO NOTHING;


INSERT INTO ArchivedTransaction (
     cid, acc, trndate, trntype, trn, orno, prin, intr, oth, 
     waivedint, trnamt, balance, pendapprove, trnmnem_cd, username, 
     termid, refno, trndesc, particulars, apptype, "time", unitcode,
     AlternateTrn
  )
SELECT 
  acc.CID, acc.acc, t.trndate, 0 trntype, t.orno trn, t.orno, 
  t.TrnAmt Prin, 0 Intr, 0 oth, 
  0 waivedint, t.TrnAmt, 0 Balance, 1 pendapprove, 0 trnmnem_cd, '' username, 
  '' termid, 0 refno, 
  '' trndesc, '' particular, 2 apptype, trndate "time", cus.unit_code unitcode,
  'M' || t.BrCode || to_char(t.TrnDate, 'YYYYMMDD') || to_char(ORNO,'0000000000') AlternateTrn
FROM staging.stg_Mutual_Fund t
INNER JOIN Account.MapAccount ma on ma.BrCode = t.BrCode and ma.Acc = t.CID::varchar(10)
INNER JOIN Account.Accounts acc on acc.Acc = ma.NewAcc
INNER JOIN Account.Customer cus on cus.CID = Acc.CID
WHERE ma.Archived = TRUE t.BrCode in ('2F','M6','N1','W5','P7')
ON CONFLICT(Acc, TrnDate, Trn) DO NOTHING;

update dailytransaction set particulars = 'delete'
from dailytransaction d 
left join accounts a on d.acc = a.acc
where a.acc is null and dailytransaction.kafkaid = d.kafkaid 



select * from dailytransaction where  particulars = 'delete'



