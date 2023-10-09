  -- INSERT system_config
  INSERT INTO system_config(
    office_id, gl_date, last_accruals, last_month_end, 
    next_month_end, system_date, run_state
  )
  SELECT 
    ofc.id office_id, ebsysdate gl_date, ebsysdate last_accruals, ebsysdate last_month_end, 
    ebsysdate next_month_end, ebsysdate system_date, 0 run_state 
  FROM 
    vwoffice ofc 
    ,(SELECT '2022-10-28'::date ebsysdate ) config
  WHERE officetype  = 'Area' and code = 'E3'
  ON CONFLICT(office_id) DO UPDATE SET
    gl_date = excluded.gl_date, 
    last_accruals = excluded.last_accruals, 
    last_month_end = excluded.last_month_end, 
    next_month_end = excluded.next_month_end, 
    system_date = excluded.system_date, 
    run_state = excluded.run_state;
  
-- Mapping for jnlHeader
INSERT INTO staging.TrnMap(UUID, BrCode, TrnDate, UserName)
SELECT uuid_generate_v4() uuid, t.BrCode, t.TrnDate, t.UserName
FROM
 (SELECT 
    BrCode, TrnDate, trim(lower(COALESCE(UserName,'SA'))) UserName 
  FROM staging.jnlheaders 
  GROUP BY BrCode, TrnDate, trim(lower(COALESCE(UserName,'SA'))) 
 ) t
ON CONFLICT(BrCode, TrnDate, lower(UserName)) DO NOTHING;

-- Insert to "Ticket" from TrnMap
INSERT INTO Ticket (UUID, central_office_id, ticket_type_id, ticket_date, postedby_id, status_id, remarks)
  SELECT m.UUID, co.ID central_office_id, typ.ID ticket_type_id, m.TrnDate ticket_date, u.Id postedby_id, stat.Id status_id, '' remarks
  FROM staging.TrnMap m 
  LEFT JOIN Office co on co.Code = 'CI'
  LEFT JOIN Reference typ on lower(typ.Ref_Type) = 'tickettype' and lower(typ.Title) = lower('Over the Counter')
  LEFT JOIN Reference stat on lower(stat.Ref_Type) = 'ticketstatus' and lower(stat.Title) = lower('Completed')
  LEFT JOIN Users u on lower(u.login_name) = lower(m.BrCode || '-' || trim(m.UserName))
ON CONFLICT(UUID) DO NOTHING;

-- Insert to TrnHead from jnlHeader
INSERT INTO Trn_Head(
  Trn_Serial, ticket_id, trn_date, type_id, Particular, office_id, 
  user_id, transacting_iiid, orno, isfinal, ismanual, 
  alternate_trn, reference
)  
SELECT 
  t.OrNo2 Trn_Serial, t.ticket_id, t.trn_date, typ.Id type_id, t.Particular, br.Id office_id, 
  t.User_Id, cus.iiid transacting_iiid, t.OrNo2 orno, true isfinal, true ismanual, 
  null alternate_trn, COALESCE (t.PrNo::varchar(50),'') reference
FROM   
  (SELECT 
    m.BrCode, m.BrCode || '-' || t.trn OrNo2, c.Id ticket_id, 
    c.Ticket_Date trn_Date, c.postedby_id User_Id,
    NULL CID, '' PrNo, t.particulars Particular,
    CASE t.code WHEN 0 THEN 'Manual Entry' ELSE 'Automated Entry' END trnType
   FROM staging.trnMap m
   INNER JOIN staging.jnlheaders t on t.BrCode = m.BrCode and lower(trim(COALESCE(t.username,'SA'))) = lower(m.username) and t.trndate = m.trnDate
   INNER JOIN ticket c on c.Uuid = m.Uuid
   ) t
LEFT JOIN vwoffice br on lower(br.Code) = lower(t.BrCode) and br.officetype = 'Area'
LEFT JOIN customer cus on cus.customer_alt_id = t.BrCode || '-' || t.CID
LEFT JOIN Reference typ on lower(typ.Ref_Type) = 'trnheadtype' and lower(typ.Title) = lower(trnType)
ON CONFLICT (Trn_Serial) DO UPDATE SET 
  ticket_id =  EXCLUDED.ticket_id,
  trn_date =  EXCLUDED.trn_date,
  type_id =  EXCLUDED.type_id,
  Particular = excluded.Particular,
  office_id =  EXCLUDED.office_id,
  user_id =  EXCLUDED.user_id,
  transacting_iiid =  EXCLUDED.transacting_iiid,
  orno =  EXCLUDED.orno,
  isfinal =  EXCLUDED.isfinal,
  ismanual =  EXCLUDED.ismanual,
  alternate_trn =  EXCLUDED.alternate_trn,
  reference =  EXCLUDED.reference;

-- Insert to TrnHead from satrnMaster
INSERT INTO Trn_Head(
  Trn_Serial, ticket_id, trn_date, type_id, Particular, office_id, 
  user_id, transacting_iiid, orno, isfinal, ismanual, 
  alternate_trn, reference
)  
SELECT 
  t.OrNo2 Trn_Serial, t.ticket_id, t.trn_date, typ.Id type_id, '' Particular, br.Id office_id, 
  t.User_Id, cus.iiid transacting_iiid, t.OrNo2 orno, true isfinal, true ismanual, 
  null alternate_trn, COALESCE (t.PrNo::varchar(50),'') reference
FROM   
  (SELECT m.BrCode, m.BrCode || '-' || to_char(COALESCE(t.OrNo,0), 'fm00000000000000') || '-' || c.Id OrNo2, c.Id ticket_id, c.Ticket_Date trn_Date, c.postedby_id User_Id,
          Max(COALESCE(mp.CID,lm.CID)) CID, PrNo
   FROM staging.trnMap m
   INNER JOIN staging.satrnmaster t on t.BrCode = m.BrCode and t.username = m.username and t.trndate = m.trnDate
   INNER JOIN staging.saMaster lm on t.BrCode = lm.BrCode and t.acc = lm.acc    
   INNER JOIN ticket c on c.Uuid = m.Uuid
   LEFT JOIN staging.multiplepaymentreceipt mp on mp.BrCode = t.BrCode and mp.OrNo = t.OrNo
   GROUP BY  m.BrCode, COALESCE(t.OrNo,0), c.postedby_id, PrNo, c.Id, c.Ticket_Date
   ) t
LEFT JOIN vwoffice br on lower(br.Code) = lower(t.BrCode) and br.officetype = 'Area'
LEFT JOIN customer cus on cus.customer_alt_id = t.BrCode || '-' || t.CID
LEFT JOIN Reference typ on lower(typ.Ref_Type) = 'trnheadtype' and lower(typ.Title) = lower('Multiple Payment')
ON CONFLICT (Trn_Serial) DO UPDATE SET 
  ticket_id =  EXCLUDED.ticket_id,
  trn_date =  EXCLUDED.trn_date,
  type_id =  EXCLUDED.type_id,
  Particular = excluded.Particular,
  office_id =  EXCLUDED.office_id,
  user_id =  EXCLUDED.user_id,
  transacting_iiid =  EXCLUDED.transacting_iiid,
  orno =  EXCLUDED.orno,
  isfinal =  EXCLUDED.isfinal,
  ismanual =  EXCLUDED.ismanual,
  alternate_trn =  EXCLUDED.alternate_trn,
  reference =  EXCLUDED.reference;

-- Insert to TrnHead from MutualFund
INSERT INTO Trn_Head(
  Trn_Serial, ticket_id, trn_date, type_id, Particular, office_id, 
  user_id, transacting_iiid, orno, isfinal, ismanual, 
  alternate_trn, reference
)  
SELECT 
  t.OrNo2 Trn_Serial, t.ticket_id, t.trn_date, typ.Id type_id, '' Particular, br.Id office_id, 
  t.User_Id, cus.iiid transacting_iiid, t.OrNo2 orno, true isfinal, true ismanual, 
  null alternate_trn, COALESCE (t.PrNo::varchar(50),'') reference
FROM   
  (SELECT 
    m.BrCode, m.BrCode || '-' || to_char(COALESCE(t.OrNo,0), 'fm00000000000000') || '-' || c.Id OrNo2, c.Id ticket_id, 
    c.Ticket_Date trn_Date, c.postedby_id User_Id,
    Max(COALESCE(mp.CID,t.CID)) CID, PrNo
   FROM staging.trnMap m
   INNER JOIN staging.mutualfund t on t.BrCode = m.BrCode and t.username = m.username and t.trndate = m.trnDate
   INNER JOIN ticket c on c.Uuid = m.Uuid
   LEFT JOIN staging.multiplepaymentreceipt mp on mp.BrCode = t.BrCode and mp.OrNo = t.OrNo
   GROUP BY  m.BrCode, COALESCE(t.OrNo,0), c.postedby_id, PrNo, c.Id, c.Ticket_Date
   ) t
LEFT JOIN vwoffice br on lower(br.Code) = lower(t.BrCode) and br.officetype = 'Area'
LEFT JOIN customer cus on cus.customer_alt_id = t.BrCode || '-' || t.CID
LEFT JOIN Reference typ on lower(typ.Ref_Type) = 'trnheadtype' and lower(typ.Title) = lower('Multiple Payment')
ON CONFLICT (Trn_Serial) DO UPDATE SET 
  ticket_id =  EXCLUDED.ticket_id,
  trn_date =  EXCLUDED.trn_date,
  type_id =  EXCLUDED.type_id,
  Particular = excluded.Particular,
  office_id =  EXCLUDED.office_id,
  user_id =  EXCLUDED.user_id,
  transacting_iiid =  EXCLUDED.transacting_iiid,
  orno =  EXCLUDED.orno,
  isfinal =  EXCLUDED.isfinal,
  ismanual =  EXCLUDED.ismanual,
  alternate_trn =  EXCLUDED.alternate_trn,
  reference =  EXCLUDED.reference;

COPY
(SELECT uuid_generate_v4() uuid, 
  d.trn_head_id, d.series, d.Alternate_Key,
  d.value_date, d.account_id, d.Trn_Type_Code, d.currency, d.item_id, d.passbook_posted, 
  d.trn_prin, d.trn_int, d.waived_Int
FROM
  (SELECT
    th.id trn_head_id, Row_Number() OVER(Partition by t.OrNo Order by t.TrnDate) series,
    t.BrCode||to_char(t.trndate, 'YYYYMMDD')||'L'||t.Trn Alternate_Key, 
    t.TrnDate value_date, acc.Id account_id, t.TrnType Trn_Type_Code, 'PHP' currency, null item_id, false passbook_posted, 
    t.Prin trn_prin, t.Intr trn_int, t.waivedint waived_Int
  FROM staging.trnMap m
  INNER JOIN staging.trnmaster t on t.BrCode = m.BrCode and   trim(lower(COALESCE(t.username,'SA'))) = m.username and t.trndate = m.trnDate
  INNER JOIN ticket c on c.Uuid = m.Uuid
  INNER JOIN trn_head th on th.Trn_Serial = m.BrCode || '-' || to_char(t.OrNo, 'fm00000000000000') || '-' || c.Id 
  LEFT JOIN Account acc on acc.alternate_acc = t.BrCode || trim(t.Acc)
  ) d
LEFT JOIN Account_Tran accTrn on accTrn.Alternate_Key = d.Alternate_Key
WHERE accTrn.uuid is null
)
TO '/var/lib/postgresql/AccTran.csv'  WITH DELIMITER '|' CSV HEADER;

select loaddata ('Account_Tran (uuid, trn_head_id, series, Alternate_Key,
  value_date, account_id, Trn_Type_Code, currency, item_id, passbook_posted, 
  trn_prin, trn_int, waived_Int) ','/var/lib/postgresql/AccTran.csv');



COPY (
SELECT uuid_generate_v4() uuid, 
  d.trn_head_id, d.series, d.Alternate_Key,
  d.value_date, d.account_id, d.Trn_Type_Code, d.currency, d.item_id, d.passbook_posted, 
  d.trn_prin, d.trn_int, d.waived_Int
FROM
(SELECT
  th.id trn_head_id, Row_Number() OVER(Partition by t.OrNo Order by t.TrnDate) + 100000000 series, 
    t.BrCode||to_char(t.trndate, 'YYYYMMDD')||'S'||t.Trn Alternate_Key, 
  t.TrnDate value_date, acc.Id account_id, t.TrnType Trn_Type_Code, 'PHP' currency, null item_id, false passbook_posted, 
  t.TrnAmt trn_prin, 0 trn_int, 0 waived_Int
FROM staging.trnMap m
INNER JOIN staging.satrnmaster t on t.BrCode = m.BrCode and trim(lower(COALESCE(t.username,'SA'))) = trim(lower(m.username)) and t.trndate = m.trnDate
INNER JOIN ticket c on c.Uuid = m.Uuid
INNER JOIN trn_head th on th.Trn_Serial = m.BrCode || '-' || to_char(COALESCE(t.OrNo,0), 'fm00000000000000') || '-' || c.Id 
LEFT JOIN Account acc on acc.alternate_acc = t.BrCode || trim(t.Acc)
) d
LEFT JOIN Account_Tran accTrn on accTrn.Alternate_Key = d.Alternate_Key
where accTrn.uuid is null
)
TO '/var/lib/postgresql/AccsaTran.csv'  WITH DELIMITER '|' CSV HEADER;


select loaddata ('Account_Tran (uuid, trn_head_id, series, Alternate_Key,
  value_date, account_id, Trn_Type_Code, currency, item_id, passbook_posted, 
  trn_prin, trn_int, waived_Int) ','/var/lib/postgresql/AccsaTran.csv');
   

SELECT t.username, m.BrCode, m.BrCode || '-' || to_char(t.OrNo, 'fm00000000000000') OrNo2, c.Id ticket_id, max(c.Ticket_Date) trn_Date, c.postedby_id User_Id,
          Max(COALESCE(mp.CID,lm.CID)) CID, PrNo
   FROM staging.trnMap m
   INNER JOIN staging.trnmaster t on t.BrCode = m.BrCode and t.username = m.username and t.trndate = m.trnDate
   INNER JOIN staging.lnMaster lm on t.BrCode = lm.BrCode and t.acc = lm.acc    
   INNER JOIN ticket c on c.Uuid = m.Uuid
   LEFT JOIN staging.multiplepaymentreceipt mp on mp.BrCode = t.BrCode and mp.OrNo = t.OrNo
   where t.orno = 48086
   GROUP BY  m.BrCode, t.OrNo, c.postedby_id, PrNo, t.username, c.Id
   
   
   select * from staging.trnmap where 
   
   selec
   

SQL Error [23505]: ERROR: duplicate key value violates unique constraint "trn_headserial"
  Detail: Key (trn_serial)=(E3-00000000048086) already exists.
  ERROR: duplicate key value violates unique constraint "trn_headserial"
  Detail: Key (trn_serial)=(E3-00000000048086) already exists.
  ERROR: duplicate key value violates unique constraint "trn_headserial"
  Detail: Key (trn_serial)=(E3-00000000048086) already exists.

  
  select * from mu
  
select * from Trn_Head

   select currval('wee')
   
   select *multi
  
select * from Reference stat where lower(typ.Ref_Type) = 'ticketstatus' and lower(typ.Title) = lower('Completed')




