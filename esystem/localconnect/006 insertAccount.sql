-- INSERT Account from lnMaster
INSERT INTO Account(
  customer_id, acc, alternate_acc, account_name, balance, 
  non_current, contract_date, credit, debit, isbudget, 
  last_activity_date, open_date, passbook_line, pending_trn_amt, principal, 
  class_id, account_type_id, budget_account_id, currency, office_id, 
  referredby_id, status_code, remarks, other_info
  )
SELECT
--  typ.Id, cls.id cls, grp.id grp, 
  c.id customer_id, a.BrCode || a.Acc acc, a.BrCode || a.Acc alternate_acc, 
  FullNameTFMLS(ii.Title, ii.Last_Name, ii.First_Name, ii.Middle_Name, ii.Suffix_Name) account_name, 
  a.Principal - a.Prin balance, 
  0 non_current, a.disbdate contract_date, a.Prin credit, 0 debit, false isbudget, 
  a.lasttrndate last_activity_date, a.opendate open_date, 0 passbook_line,  0 pending_trn_amt, a.Principal principal, 
  accCls.Id class_id, typ.Id account_type_id, null budget_account_id, 'PHP' currency, o.Id office_id, 
  null referredby_id, 
  CASE WHEN wo.acc is null and a.status not in (98, 99) THEN a.status::integer ELSE 97 END status_code, '' remarks, null other_info
FROM staging.lnMaster a
INNER JOIN Customer c on c.customer_alt_id = a.BrCode || '-' || a.CID
INNER JOIN Identity_Info ii on ii.ID = c.IIID
INNER JOIN Account_Type typ on typ.Code = a.accttype 
LEFT JOIN vwoffice o  on o.officetype = 'Area' and o.Code = a.BrCode
LEFT JOIN Reference cls on lower(cls.ref_type) = 'loanclass' and lower(cls.title) = 'current'
LEFT JOIN Reference grp on lower(grp.ref_type) = 'accounttypegroup' and lower(grp.title) = 'microfinance'
LEFT JOIN Account_Class accCls on accCls.product_id = typ.product_id and accCls.group_id = grp.Id and cls.id = accCls.class_id 
LEFT JOIN staging.writeoff wo on a.brcode = wo.brcode and a.Acc = wo.Acc
ON CONFLICT(alternate_acc)
DO UPDATE SET
  customer_id =  EXCLUDED.customer_id,
  acc =  EXCLUDED.acc,
  account_name =  EXCLUDED.account_name,
  balance =  EXCLUDED.balance,
  non_current =  EXCLUDED.non_current,
  contract_date =  EXCLUDED.contract_date,
  credit =  EXCLUDED.credit,
  debit =  EXCLUDED.debit,
  isbudget =  EXCLUDED.isbudget,
  last_activity_date =  EXCLUDED.last_activity_date,
  open_date =  EXCLUDED.open_date,
  passbook_line =  EXCLUDED.passbook_line,
  pending_trn_amt =  EXCLUDED.pending_trn_amt,
  principal =  EXCLUDED.principal,
  class_id =  EXCLUDED.class_id,
  account_type_id =  EXCLUDED.account_type_id,
  budget_account_id =  EXCLUDED.budget_account_id,
  currency =  EXCLUDED.currency,
  office_id =  EXCLUDED.office_id,
  referredby_id =  EXCLUDED.referredby_id,
  status_code =  EXCLUDED.status_code,
  remarks =  EXCLUDED.remarks,
  other_info =  EXCLUDED.other_info;  
  
-- INSERT Account for saMaster
ININSERT INTO Account(
  customer_id, acc, alternate_acc, account_name, balance, 
  non_current, contract_date, credit, debit, isbudget, 
  last_activity_date, open_date, passbook_line, pending_trn_amt, principal, 
  class_id, account_type_id, budget_account_id, currency, office_id, 
  referredby_id, status_code, remarks, other_info
  )
SELECT
--  a."type",typ.Id, cls.id cls, grp.id grp, 
  c.id customer_id, a.BrCode || a.Acc acc, a.BrCode || a.Acc alternate_acc, 
  FullNameTFMLS(ii.Title, ii.Last_Name, ii.First_Name, ii.Middle_Name, ii.Suffix_Name) account_name, 
  a.balance, 
  0 non_current, c.date_entry contract_date, 0 credit, 0 debit, false isbudget, 
  a.dolasttrn last_activity_date, a.dopen open_date, 0 passbook_line,  0 pending_trn_amt, 0 principal, 
  accCls.Id class_id, typ.Id account_type_id, null budget_account_id, 'PHP' currency, o.Id office_id, 
  null referredby_id, a.status::integer status_code, '' remarks, null other_info
FROM staging.saMaster a
INNER JOIN Customer c on c.customer_alt_id = a.BrCode || '-' || a.CID
INNER JOIN Identity_Info ii on ii.ID = c.IIID
INNER JOIN Account_Type typ on typ.Code = a."type" 
INNER JOIN product prod on prod.ID = typ.product_id and prod.product_name = 'Savings' 
LEFT JOIN vwoffice o  on o.officetype = 'Area' and o.Code = a.BrCode
LEFT JOIN Reference cls on lower(cls.ref_type) = 'savingsclass' and lower(cls.title) = 'active'
LEFT JOIN Reference grp on lower(grp.ref_type) = 'accounttypegroup' and lower(grp.title) = 'microfinance'
LEFT JOIN Account_Class accCls on accCls.product_id = typ.product_id and accCls.group_id = grp.Id and cls.id = accCls.class_id 
ON CONFLICT(alternate_acc)
DO UPDATE SET
  customer_id =  EXCLUDED.customer_id,
  acc =  EXCLUDED.acc,
  account_name =  EXCLUDED.account_name,
  balance =  EXCLUDED.balance,
  non_current =  EXCLUDED.non_current,
  contract_date =  EXCLUDED.contract_date,
  credit =  EXCLUDED.credit,
  debit =  EXCLUDED.debit,
  isbudget =  EXCLUDED.isbudget,
  last_activity_date =  EXCLUDED.last_activity_date,
  open_date =  EXCLUDED.open_date,
  passbook_line =  EXCLUDED.passbook_line,
  pending_trn_amt =  EXCLUDED.pending_trn_amt,
  principal =  EXCLUDED.principal,
  class_id =  EXCLUDED.class_id,
  account_type_id =  EXCLUDED.account_type_id,
  budget_account_id =  EXCLUDED.budget_account_id,
  currency =  EXCLUDED.currency,
  office_id =  EXCLUDED.office_id,
  referredby_id =  EXCLUDED.referredby_id,
  status_code =  EXCLUDED.status_code,
  remarks =  EXCLUDED.remarks,
  other_info =  EXCLUDED.other_info;  
  
  -- INSERT Account for MBA from Customer
INSERT INTO Account(
  customer_id, acc, alternate_acc, account_name, balance, 
  non_current, contract_date, credit, debit, isbudget, 
  last_activity_date, open_date, passbook_line, pending_trn_amt, principal, 
  class_id, account_type_id, budget_account_id, currency, office_id, 
  referredby_id, status_code, remarks, other_info
  )
SELECT
  c.id customer_id, a.BrCode ||'MBA'|| a.CID acc, a.BrCode ||'MBA'|| a.CID alternate_acc, 
  FullNameTFMLS(ii.Title, ii.Last_Name, ii.First_Name, ii.Middle_Name, ii.Suffix_Name) account_name, 
  0 balance, 
  0 non_current, c.date_entry contract_date, 0 credit, 0 debit, false isbudget, 
  null last_activity_date, c.date_entry open_date, 0 passbook_line,  0 pending_trn_amt, 0 principal, 
  accCls.Id class_id, typ.Id account_type_id, null budget_account_id, 'PHP' currency, o.Id office_id, 
  null referredby_id, CASE WHEN a.status::integer=2 THEN 99 ELSE 10 END status_code, '' remarks, null other_info
FROM staging.Customer a
INNER JOIN (VALUES(100,'Golden Life 100'),
      (50,'Golden Life 50'),
      (20,'Mutual Fund'),
      (0,'Mutual Fund')
      ) ins(code, AccType) on ins.Code = a.mutualamount 
INNER JOIN Customer c on c.customer_alt_id = a.BrCode || '-' || a.CID
INNER JOIN Identity_Info ii on ii.ID = c.IIID
INNER JOIN Account_Type typ on typ.account_type = ins.accType 
LEFT JOIN vwoffice o  on o.officetype = 'Area' and o.Code = a.BrCode
LEFT JOIN Reference cls on lower(cls.ref_type) = 'accountclass' and lower(cls.title) = 'current'
LEFT JOIN Reference grp on lower(grp.ref_type) = 'accounttypegroup' and lower(grp.title) = 'microinsurance'
LEFT JOIN Account_Class accCls on accCls.product_id = typ.product_id and accCls.group_id = grp.Id and cls.id = accCls.class_id 
INNER JOIN 
 (SELECT distinct s.BrCode, s.cid
  FROM staging.saMaster s
  LEFT JOIN 
   (SELECT cid FROM staging.mutualfund m group by cid) m on m.cid = s.CID
    WHERE type = 60 and (m.cid is not null or status <> '99')
  ) mf on a.brcode = mf.BrCode and a.cid = mf.cid

ON CONFLICT(alternate_acc)
DO UPDATE SET
  customer_id =  EXCLUDED.customer_id,
  acc =  EXCLUDED.acc,
  account_name =  EXCLUDED.account_name,
  balance =  EXCLUDED.balance,
  non_current =  EXCLUDED.non_current,
  contract_date =  EXCLUDED.contract_date,
  credit =  EXCLUDED.credit,
  debit =  EXCLUDED.debit,
  isbudget =  EXCLUDED.isbudget,
  last_activity_date =  EXCLUDED.last_activity_date,
  open_date =  EXCLUDED.open_date,
  passbook_line =  EXCLUDED.passbook_line,
  pending_trn_amt =  EXCLUDED.pending_trn_amt,
  principal =  EXCLUDED.principal,
  class_id =  EXCLUDED.class_id,
  account_type_id =  EXCLUDED.account_type_id,
  budget_account_id =  EXCLUDED.budget_account_id,
  currency =  EXCLUDED.currency,
  office_id =  EXCLUDED.office_id,
  referredby_id =  EXCLUDED.referredby_id,
  status_code =  EXCLUDED.status_code,
  remarks =  EXCLUDED.remarks,
  other_info =  EXCLUDED.other_info;  

INSERT INTO Account_Interest(
  account_id, interest, effective_rate, interest_rate, credit, 
  debit, accruals, waived_int, last_accrued_date
  )
SELECT 
  acc.Id account_id, a.interest, ConIntRate effective_rate, a.interest/a.principal interest_rate, 
  a.intr credit, 0 debit, 0 accruals, a.waivedint waived_int, a.DoMaturity last_accrued_date
FROM staging.lnMaster a
LEFT JOIN Account acc on acc.alternate_acc = a.BrCode || a.Acc
ON CONFLICT(account_id)
DO UPDATE SET
  interest =  EXCLUDED.interest,
  effective_rate =  EXCLUDED.effective_rate,
  interest_rate =  EXCLUDED.interest_rate,
  credit =  EXCLUDED.credit,
  debit =  EXCLUDED.debit,
  accruals =  EXCLUDED.accruals,
  waived_int =  EXCLUDED.waived_int,
  last_accrued_date =  EXCLUDED.last_accrued_date;

-- Insert Account_Term from lnMaster
INSERT INTO Account_Term(
  account_id, frequency, n, paid_n, fixed_due, cummulative_due, 
  date_start, maturity
  )  
SELECT
  acc.Id account_id, CASE frequency WHEN 0 THEN 50 WHEN 1 THEN 12 WHEN 2 THEN 24 ELSE frequency END, gives n, 
  weekspaid paid_n, 0 fixed_due, 0 cummulative_due, i.duedate date_start, a.domaturity maturity
FROM staging.lnMaster a
LEFT JOIN Account acc on acc.alternate_acc = a.BrCode || a.Acc
INNER JOIN staging.loaninst i on i.brcode = a.brcode and i.acc = a.acc and i.dnum = 1
ON CONFLICT(account_id)
DO UPDATE SET
  frequency =  EXCLUDED.frequency,
  n =  EXCLUDED.n,
  paid_n =  EXCLUDED.paid_n,
  fixed_due =  EXCLUDED.fixed_due,
  cummulative_due =  EXCLUDED.cummulative_due,
  date_start =  EXCLUDED.date_start,
  maturity =  EXCLUDED.maturity;


-- Insert Account_Term from Customer for CBU
INSERT INTO Account_Term(
  account_id, frequency, n, paid_n, fixed_due, cummulative_due, 
  date_start, maturity
  )  
SELECT
  acc.Id account_id, 50 frequency, 0 n, 
  0 paid_n, c.pledgeamount fixed_due, c.accpledge cummulative_due, c.daterecognized date_start, '1900-01-01'::date maturity
FROM staging.customer c 
INNER JOIN staging.saMaster a on c.BrCode = a.BrCode and c.CID = a.CID and a.Status::integer = 10
INNER JOIN Account acc on acc.alternate_acc = a.BrCode || a.Acc
WHERE c.status in (0, 1, 4) and c.pledgeamount > 0 
ON CONFLICT(account_id)
DO UPDATE SET
  frequency =  EXCLUDED.frequency,
  n =  EXCLUDED.n,
  paid_n =  EXCLUDED.paid_n,
  fixed_due =  EXCLUDED.fixed_due,
  cummulative_due =  EXCLUDED.cummulative_due,
  date_start =  EXCLUDED.date_start,
  maturity =  EXCLUDED.maturity;

-- Insert Account_Term from Customer for MBA
INSERT INTO Account_Term(
  account_id, frequency, n, paid_n, fixed_due, cummulative_due, 
  date_start, maturity
  )  
SELECT
  acc.Id account_id, 50 frequency, 0 n, 
  0 paid_n, c.mutualamount fixed_due, c.accmutual cummulative_due, c.daterecognized date_start, '1900-01-01'::date maturity
FROM staging.customer c 
INNER JOIN Account acc on acc.alternate_acc = c.BrCode ||'MBA'|| c.CID
WHERE c.status in (0, 1, 4) and c.mutualamount > 0 
ON CONFLICT(account_id)
DO UPDATE SET
  frequency =  EXCLUDED.frequency,
  n =  EXCLUDED.n,
  paid_n =  EXCLUDED.paid_n,
  fixed_due =  EXCLUDED.fixed_due,
  cummulative_due =  EXCLUDED.cummulative_due,
  date_start =  EXCLUDED.date_start,
  maturity =  EXCLUDED.maturity;
  

INSERT INTO Schedule(
  account_id, series, due_date, due_prin, due_int, 
  end_prin, end_int, carrying_value, realizable
  )
SELECT 
  acc.Id account_id, l.dNum series, l.duedate due_date, l.dueprin due_prin, l.dueint due_int, 
  m.Principal - p.DuePrin end_prin, m.Interest - p.DueInt end_int, m.Principal -p.DuePrin - p.DueInt + p.UpInt carrying_value, l.UpInt realizable
FROM staging.loaninst l
INNER JOIN staging.lnMaster m on m.BrCode = l.BrCode and m.Acc = l.Acc
LEFT JOIN Account acc on acc.alternate_acc = l.BrCode || l.Acc
INNER JOIN 
  (SELECT l.BrCode, l.Acc, l.dNum, Sum(p.DuePrin) DuePrin, Sum(p.DueInt) DueInt, Sum(p.UpInt) UpInt
   FROM staging.loaninst l
   INNER JOIN staging.loaninst p on l.BrCode = p.BrCode and l.Acc = p.Acc and p.dNum <= l.dNum
   GROUP BY l.BrCode, l.Acc, l.dNum   
  ) p on p.BrCode = l.BrCode and p.Acc = l.Acc and p.dNum = l.dNum;
ON CONFLICT(account_id, series)
DO UPDATE SET
  due_date =  EXCLUDED.due_date,
  due_prin =  EXCLUDED.due_prin,
  due_int =  EXCLUDED.due_int,
  end_prin =  EXCLUDED.end_prin,
  end_int =  EXCLUDED.end_int,
  carrying_value =  EXCLUDED.carrying_value,
  realizable =  EXCLUDED.realizable,
  other_info =  EXCLUDED.other_info;

