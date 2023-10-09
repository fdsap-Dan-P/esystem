drop view if exists ColSht;
CREATE OR REPLACE VIEW ColSht as
WITH grp AS (
  SELECT 
    g.code, g.title, g.pattern
  FROM ( 
  VALUES 
    (0.1,'Sipag Loans'::text,'%Sipag%'::text), 
    (0.2,'Agri Loan Program'::text,'%Agri%'::text), 
    (0.3,'Small Business Loans'::text,'%Small%'::text), 
    (0.4,'UNLAD Loans'::text,'%Unlad%'::text), 
    (0.5,'w/ Loans'::text,''::text)) 
   g(code, title, pattern)
  ), 
acc AS (
  SELECT
    acc.id Account_Id, cus.iiid, cus.id customerid, cus.central_office_id centralofficeid, cus.cid, acc.acc, acc.alternate_acc alternateacc,
    FullName(ii.last_name, ii.first_name, ii.middle_name) clientname,
    grp.id centercode, grp.group_name CenterName, grp.officer_id aoID,
    sc.system_date - (date_part('dow'::text, sc.system_date) - grp.meeting_day::double precision - 1)::integer AS meetingdate,
    grp.date_stablished datestablished, 
    grp.meeting_day meetingday, unit.id mancode,
    unit.office_name unit, unit.Officer_iiid UnitManager_iiid,
    ar.code BrCode, ar.id areacode, ar.Officer_iiid AreaManager_iiid, 
    ar.office_name Area, ar.address_detail AreaAddress, --acc.status_code, 
    acctype.account_type,
    prod.code AppType,
    CASE prod.product_name 
      WHEN 'Savings' THEN 1
      WHEN 'Loan' THEN 2
      WHEN 'Collecting Facility' THEN 3 ELSE NULL::integer END AS code,
    accType.code accttype, acctype.account_type acctdesc,
    acc.Contract_Date disbdate, term.date_start datestart, term.maturity maturity,
    acc.principal,  COALESCE(ai.interest,0) interest,
    COALESCE(term.n,0) gives,
    acc.debit - acc.credit prin,
    COALESCE(ai.debit,0) - COALESCE(ai.credit,0) intr,
    acc.principal - acc.credit + acc.debit balprin,
    COALESCE(ai.interest,0) - COALESCE(ai.credit,0) + COALESCE(ai.debit,0) balint,
--            a.balance,
    COALESCE(ai.waived_int,0) waivedint, 
    COALESCE(term.fixed_due, 0) fixed_due, COALESCE(term.cummulative_due ,0) cummulative_due,
    sc.system_date - (date_part('dow'::text, sc.system_date) - 6::double precision)::integer AS branchsysdate,
    CASE WHEN acc.status_code = 97 THEN 0::numeric ELSE acc.principal - (acc.credit - acc.debit) END AS wbalprin,
  --  CASE WHEN a.apptype = 3 THEN acc.principal - acc.prin ELSE 0::numeric END AS lbalprin,
   --         COALESCE(grp_1.code, 0.5) AS grpcode,
    cls.Code classification, cls.Title classdesc,
    subcls.Code subclassification, subcls.Title subclassdesc,
    cus.status_code, stat.Title StatusDesc, central.office_name OrgName, central.address_detail OrgAddress, accType.normal_balance,
    csn.value SharesOfStock
  FROM Customer cus
  INNER JOIN Office central on central.id = cus.central_office_id 
  INNER JOIN identity_info ii on cus.iiid = ii.id 
  INNER JOIN Customer_Group grp on cus.customer_group_id = grp.id
  INNER JOIN Office unit on grp.office_id = unit.id
  INNER JOIN Office ar on unit.parent_id = ar.id
  INNER JOIN system_config sc on ar.id = sc.office_id 
  INNER JOIN Account acc on cus.id = acc.customer_id 
  INNER JOIN account_type accType on acc.account_type_id  = accType.id 
  INNER JOIN product prod on accType.product_id = prod.id 
  LEFT JOIN account_interest ai on acc.id = ai.account_id 
  LEFT JOIN account_term term on acc.id = term.account_id 
  LEFT JOIN Reference stat on cus.status_code = stat.code and lower(stat.ref_type) = 'customerstatus'
  LEFT JOIN Reference cls on cus.classification_id = cls.id 
  LEFT JOIN Reference subcls on cus.sub_classification_id = subcls.id 
  LEFT JOIN customer_specs_number csn on csn.customer_id  = cus.id and lower(csn.Specs_Code) = 'shareofstocks'
--  LEFT JOIN Write
-- select * from customer cus where cus.customer_alt_id = 'E3-6686118'
  WHERE acc.isopen  --and cus.id = 4798
  ),
  
  
sched as (
  SELECT 
    acc.account_Id, acc.acc, acc.branchsysdate,
    max(CASE WHEN s.due_date < (acc.branchsysdate - 7) AND (acc.BalPrin + acc.BalInt) > (s.end_prin + s.end_int) 
        THEN s.due_date ELSE acc.datestart END
        ) AS startarrears,
    sum(CASE WHEN s.due_date < (acc.branchsysdate - 7) AND (acc.BalPrin + acc.BalInt) > (s.end_prin + s.end_int) 
        THEN 1 ELSE 0 END
        ) AS defctr,
    max(COALESCE(s.due_prin + s.due_int, acc.Fixed_Due)) AS amort,
    sum(CASE WHEN s.due_date <= (acc.branchsysdate - 7) and s.end_prin < acc.BalPrin 
        THEN 1 ELSE 0 END) UnpaidCtr,
    sum(CASE WHEN s.due_date <= (acc.branchsysdate - 7) 
        THEN s.due_prin ELSE 0::numeric END
        ) + acc.prin AS defprin,
    sum(CASE WHEN s.due_date <= (acc.branchsysdate - 7) 
        THEN s.due_int ELSE 0::numeric END
        ) + acc.intr AS defint, 
    sum(CASE  WHEN s.due_date <= acc.branchsysdate 
        THEN s.due_prin ELSE 0::numeric END
        ) + acc.prin AS curprin,
    sum(CASE WHEN s.due_date <= acc.branchsysdate 
        THEN s.due_int ELSE 0::numeric END
        ) + acc.intr AS curint,
    sum(CASE WHEN s.due_date <= (acc.branchsysdate + 7) 
        THEN s.due_prin ELSE 0::numeric END
        ) + acc.prin AS nextprin,
    sum(CASE WHEN s.due_date <= (acc.branchsysdate + 7) 
        THEN s.due_int ELSE 0::numeric END
        ) + acc.intr AS nextint
  --  acc.wbalprin,
  --            acc_1.lbalprin,
  --            acc_1.grpcode
  FROM Acc
  LEFT JOIN Schedule s on s.account_id  = acc.account_id
--  where acc = 'E30101-4041-0528625'  
  GROUP BY acc.account_Id, acc.acc, acc.branchsysdate, acc.prin, acc.intr
  ),  
sht as (
SELECT 
  acc.BrCode, acc.AppType, acc.Code, 
  CASE 
    WHEN COALESCE(acc.wbalprin, 0::numeric) > 0::numeric THEN 4
    WHEN COALESCE(sched.defctr, 0::numeric) > 0::numeric THEN 3
    WHEN acc.BalPrin > 0 and acc.Code = 2 THEN 2 ELSE 0 END,  
  -- StatusDesc, 
  acc.Acc, acc.Iiid, acc.customerid, acc.centralofficeid, acc.cid,
  acc.ClientName, acc.CenterCode, acc.CenterName, acc.ManCode, acc.Unit, 
  acc.UnitManager_iiid, fullname(um.last_name , um.first_name , um.middle_name) UM, 
  acc.AreaCode, acc.Area, ao.StaffName, 
  acc.Accttype, acc.AcctDesc, acc.Disbdate, 
  acc.DateStart, acc.Maturity, acc.Principal, acc.Interest, 
  acc.Gives, acc.BalPrin, acc.BalInt, 
  COALESCE(sched.Amort,acc.fixed_Due) Amort, 
  CASE WHEN acc.code = 2 THEN COALESCE(sched.CurPrin,0) ELSE cummulative_due END DuePrin, 
  CASE WHEN acc.code = 2 THEN COALESCE(sched.CurInt,0) ELSE 0 END DueInt, 

  CASE WHEN acc.code = 2 THEN acc.balprin+ acc.balint ELSE 0 END LoanBal,
  CASE WHEN acc.code = 1 THEN acc.balprin+ acc.balint ELSE 0 END SaveBal,
  acc.Waivedint, COALESCE(sched.UnpaidCtr,0) UnpaidCtr, 
  CASE WHEN acc.wbalprin > 0 THEN 1 ELSE 0 END Writtenoff, 
  acc.Classification, Acc.Classification ClassDesc, 
  CASE WHEN acc.Status_Code = 24 THEN 1 ELSE 0 END Writeoff,
  acc.status_code, acc.StatusDesc, acc.OrgName, acc.OrgAddress, acc.meetingDay, acc.meetingDate,
  acc.datestablished, acc.normal_balance, alternateacc,
  acc.SharesOfStock
FROM Acc 
LEFT JOIN Sched on acc.Account_Id = Sched.Account_Id
LEFT JOIN identity_info um on acc.UnitManager_iiid = um.id
LEFT JOIN identity_info am on acc.AreaManager_iiid = am.id
LEFT JOIN 
  (SELECT e.ID, FullName(ii.last_name, ii.first_name , ii.middle_name) StaffName
   FROM employee e 
   INNER JOIN identity_info ii on e.iiid = ii.id
  ) ao on ao.ID = acc.AOID
WHERE not (acc.Code = 3 and acc.fixed_due = 0)
)
SELECT 
  BrCode, AppType, code, status_code Status, StatusDesc, Acc, 
  sht.Iiid, sht.customerid, sht.centralofficeid, sht.cid, 
  UM, ClientName, CenterCode, CenterName, ManCode, Unit, 
  AreaCode, Area, StaffName, AcctType, AcctDesc, DisbDate, 
  COALESCE(DateStart,'01-01-1900'::date) DateStart, COALESCE(Maturity,'01-01-1900'::date) Maturity, Principal, Interest, Gives, 
  BalPrin, BalInt, Amort, DuePrin, DueInt, 
  LoanBal, SaveBal, WaivedInt, UnPaidCtr, WrittenOff WritenOff, 
  OrgName, OrgAddress, MeetingDate, MeetingDay, 
  SharesOfStock, datestablished DateEstablished, Classification, ClassDesc, WriteOff
FROM sht;
