/*
UPDATE  account.accounts set prin = t.prin, intR = t.Intr, WaivedInt = t.WaivedInt
FROM 
 (SELECT t.Acc,
      Sum(CASE WHEN acc.AppType = 0 and TrnType % 2 = 0 THEN -t.Prin ELSE t.Prin END) Prin, 
      sum(CASE WHEN acc.AppType = 0 and TrnType % 2 = 0 THEN -t.Intr ELSE t.Intr END) Intr,
      Sum(CASE WHEN acc.AppType = 0 and TrnType % 2 = 0 THEN -t.TrnAmt ELSE t.TrnAmt END) TrnAmt, 
      Sum(t.WaivedInt) WaivedInt 
 FROM DailyTransaction t
 INNER JOIN account.Accounts acc on acc.Acc = t.Acc
 INNER JOIN Accs on accs.acc = acc.acc
 WHERE TrnType not in (3100,3400) 
 GROUP BY t.Acc
  ) t
 where accounts.acc = t.Acc
 

select balance,* from account.accounts a  where cid = 9361293
 
select prin,* from account.accounts a where acc = '4001-0001-34183185'
order by trndate,trn

select * from account.mapaccount m where acc = 022F-4001-009234945

select * from dailytransaction d where acc = '4001-0001-34183185' order by trndate,trn

SELECT *
FROM
(SELECT 
  ma.BrCode, ma.Acc, ma.NewAcc,
  CASE WHEN trnDate Between '2021-08-04'::date and '2021-08-07' THEN Prin+IntR ELSE 0 END Wk1,
  CASE WHEN trnDate Between '2021-08-08'::date and '2021-08-13' THEN Prin+IntR ELSE 0 END Wk2
FROM dailytransaction dt 
INNER JOIN account.MapAccount ma on ma.newacc = dt.acc
where ma.BrCode in ('04','05','07')
) core

FULL JOIN 
(SELECT 
  dt.BrCode, dt.Acc, ma.NewAcc, dt.TrnDate, dt.Trn,
  CASE WHEN trnDate Between '2021-08-04'::date and '2021-08-07' THEN Prin+IntR ELSE 0 END Wk1,
  CASE WHEN trnDate Between '2021-08-08'::date and '2021-08-13' THEN Prin+IntR ELSE 0 END Wk2
FROM account.dat_trnmaster dt 
INNER JOIN account.MapAccount ma on ma.BrCode = dt.BrCode and ma.Acc = dt.Acc
where dt.BrCode in ('04','05','07') 
UNION ALL
SELECT 
  dt.BrCode, dt.CID::VarChar(10) CID, ma.NewAcc, dt.TrnDate, dt.OrNo Trn,
  CASE WHEN trnDate Between '2021-08-04'::date and '2021-08-07' THEN TrnAmt ELSE 0 END Wk1,
  CASE WHEN trnDate Between '2021-08-08'::date and '2021-08-13' THEN TrnAmt ELSE 0 END Wk2
FROM account.dat_mutual_fund dt 
INNER JOIN account.MapAccount ma on ma.BrCode = dt.BrCode and ma.Acc = dt.CID::VarChar(10)
where dt.BrCode in ('04','05','07') 
) es on core.NewAcc = es.NewAcc

WHERE abs(COALESCE(es.Wk1,0)-COALESCE(core.Wk1,0)) > 0 or abs(COALESCE(es.Wk2,0)-COALESCE(core.Wk2,0)) > 0 


select * from account.dat_mutual_fund

select   CASE WHEN trnDate Between '2021-08-04'::date and '2021-08-07' THEN Prin+IntR ELSE 0 END Wk1,
  CASE WHEN trnDate Between '2021-08-08'::date and '2021-08-13' THEN Prin+IntR ELSE 0 END Wk2,
  * from dailytransaction d 
where acc = '1012-0000-29021288' 
and trnDate Between '2021-08-04'::date and '2021-08-13'
order by trndate,trn


SELECT 
  dt.BrCode, dt.Acc, ma.NewAcc, dt.TrnDate, dt.Trn,
  CASE WHEN trnDate Between '2021-08-04'::date and '2021-08-07' THEN Prin+IntR ELSE 0 END Wk1,
  CASE WHEN trnDate Between '2021-08-08'::date and '2021-08-13' THEN Prin+IntR ELSE 0 END Wk2
FROM account.dat_trnmaster dt 
INNER JOIN account.MapAccount ma on ma.BrCode = dt.BrCode and ma.Acc = dt.Acc
where dt.BrCode in ('04','05','07') 

*/
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
INNER JOIN Account.BrCodes on brCodes.BrCode = t.BrCode
INNER JOIN Account.MapAccount ma on ma.BrCode = t.BrCode and ma.Acc = t.Acc
INNER JOIN Account.Accounts m on ma.NewAcc = m.Acc
--INNER JOIN staging.stg_lnMaster l on l.BrCode = t.BrCode and l.Acc = t.Acc
--LEFT JOIN staging.INAI_CustomerMap iiid on iiid.BrCode = t.BrCode and iiid.CID = t.CID
INNER JOIN account.Customer c on m.CID = c.CID 
WHERE ma.Archived = false --and t.BrCode in ('2F','M6','N1','W5','P7')
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
INNER JOIN Account.BrCodes on brCodes.BrCode = t.BrCode
INNER JOIN Account.MapAccount ma on ma.BrCode = t.BrCode and ma.Acc = t.Acc
INNER JOIN Account.Accounts m on ma.NewAcc = m.Acc
--INNER JOIN staging.stg_lnMaster l on l.BrCode = t.BrCode and l.Acc = t.Acc
--LEFT JOIN staging.INAI_CustomerMap iiid on iiid.BrCode = t.BrCode and iiid.CID = t.CID
INNER JOIN account.Customer c on m.CID = c.CID 
WHERE ma.Archived = TRUE --and t.BrCode in ('2F','M6','N1','W5','P7')
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
INNER JOIN Account.BrCodes on brCodes.BrCode = t.BrCode
INNER JOIN Account.MapAccount ma on ma.BrCode = t.BrCode and ma.Acc = t.CID::varchar(10)
INNER JOIN Account.Accounts acc on acc.Acc = ma.NewAcc
INNER JOIN Account.Customer cus on cus.CID = Acc.CID
WHERE ma.Archived = FALSE --and t.BrCode in ('2F','M6','N1','W5','P7')
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
INNER JOIN Account.BrCodes on brCodes.BrCode = t.BrCode
INNER JOIN Account.MapAccount ma on ma.BrCode = t.BrCode and ma.Acc = t.CID::varchar(10)
INNER JOIN Account.Accounts acc on acc.Acc = ma.NewAcc
INNER JOIN Account.Customer cus on cus.CID = Acc.CID
WHERE ma.Archived = TRUE --and t.BrCode in ('2F','M6','N1','W5','P7')
ON CONFLICT(Acc, TrnDate, Trn) DO NOTHING;


/*
update dailytransaction set particulars = 'delete'
from dailytransaction d 
left join accounts a on d.acc = a.acc
where a.acc is null and dailytransaction.kafkaid = d.kafkaid 
select * from dailytransaction where  particulars = 'delete'

*/