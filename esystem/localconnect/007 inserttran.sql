CREATE OR REPLACE FUNCTION UpdateAllBalance(lim integer)
RETURNS VOID AS $$
----------------------------------------------------------------------------------------

DECLARE
  rowVar RECORD;
  balPrin Numeric = 0;
  balInt Numeric = 0;
  AccId bigint = 0;
  curTrn CURSOR FOR 
    SELECT t.uuid, t.Account_Id, t.trn_prin, t.trn_int 
    FROM account_Tran  t
    inner join 
      (Select id from account where isbudget = false limit lim) a on a.id = t.account_id 
    WHERE trn_type_code not in (3100,3400)  -- account_id = 249518 
    ORDER BY account_id, value_date, series;
--  rowVar account_Tran%ROWTYPE;
--  rowVar RECORD;
BEGIN
    OPEN curTrn;
    LOOP
        FETCH curTrn INTO rowVar;
        EXIT WHEN NOT FOUND;

        IF rowVar.account_id <> AccId THEN
            SELECT Principal, Interest INTO balPrin, balInt
            FROM (
                SELECT a.ID, Principal, COALESCE(ai.Interest, 0) AS Interest
                FROM Account a
                LEFT JOIN Account_Interest ai ON a.id = ai.account_id 
                WHERE a.ID = rowVar.Account_Id
            ) AS acc;
            AccId := rowVar.account_id ;
        END IF;

        -- Update the column
        balPrin := balPrin + rowVar.trn_Prin;
        balInt := balInt + rowVar.trn_Int;

        -- Perform additional modifications if needed
        -- Update the row in the table
        UPDATE account_Tran SET bal_Prin = balPrin, bal_Int = balInt WHERE uuid = rowVar.uuid;
        UPDATE account set isbudget = true where id = rowVar.Account_Id;
    END LOOP;
    --COMMIT;
    CLOSE curTrn;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS staging.TrnMap(
  UUID uuid NOT NULL,  
  BrCode VarChar(2) NOT NULL, 
  TrnDate Date NULL,
  UserName VarChar(200) NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS idxTrnMapUnq ON staging.TrnMap(BrCode, TrnDate, lower(UserName));
CREATE UNIQUE INDEX IF NOT EXISTS idxTrnMapUUID ON staging.TrnMap(UUID);


-- Mapping for trnMaster
INSERT INTO staging.TrnMap(UUID, BrCode, TrnDate, UserName)
SELECT uuid_generate_v4() uuid, t.BrCode, t.TrnDate, t.UserName
FROM
 (SELECT BrCode, TrnDate, trim(lower(COALESCE(UserName,'SA'))) UserName FROM staging.TrnMaster GROUP BY BrCode, TrnDate, trim(lower(COALESCE(UserName,'SA'))) )t
ON CONFLICT(BrCode, TrnDate, lower(UserName)) DO NOTHING;


-- Mapping for satrnMaster
INSERT INTO staging.TrnMap(UUID, BrCode, TrnDate, UserName)
SELECT uuid_generate_v4() uuid, t.BrCode, t.TrnDate, t.UserName
FROM  
 (SELECT BrCode, TrnDate, trim(lower(COALESCE(UserName,'SA'))) UserName FROM staging.saTrnMaster GROUP BY BrCode, TrnDate,  trim(lower(COALESCE(UserName,'SA'))) ) t
ON CONFLICT(BrCode, TrnDate, lower(UserName)) DO NOTHING;

-- Mapping for Mutual_Fund
INSERT INTO staging.TrnMap(UUID, BrCode, TrnDate, UserName)
SELECT uuid_generate_v4() uuid, t.BrCode, t.TrnDate, t.UserName
FROM
 (SELECT BrCode, TrnDate, trim(lower(COALESCE(UserName,'SA'))) UserName FROM staging.MutualFund GROUP BY BrCode, TrnDate, trim(lower(COALESCE(UserName,'SA'))) ) t
ON CONFLICT(BrCode, TrnDate, lower(UserName)) DO NOTHING;

-- Insert to "Ticket" from TrnMap
INSERT INTO Ticket (UUID, central_office_id, ticket_type_id, ticket_date, postedby_id, status_id, remarks)
  SELECT m.UUID, co.ID central_office_id, typ.ID ticket_type_id, m.TrnDate ticket_date, u.Id postedby_id, stat.Id status_id, '' remarks
  FROM staging.TrnMap m 
  LEFT JOIN Office co on co.Code = 'CI'
  LEFT JOIN Reference typ on lower(typ.Ref_Type) = 'tickettype' and lower(typ.Title) = lower('Over the Counter')
  LEFT JOIN Reference stat on lower(stat.Ref_Type) = 'ticketstatus' and lower(stat.Title) = lower('Completed')
  LEFT JOIN Users u on lower(u.login_name) = lower(m.BrCode || '-' || trim(m.UserName))
ON CONFLICT(UUID) DO UPDATE SET
  central_office_id =  EXCLUDED.central_office_id,
  ticket_type_id =  EXCLUDED.ticket_type_id,
  ticket_date =  EXCLUDED.ticket_date,
  postedby_id =  EXCLUDED.postedby_id,
  status_id =  EXCLUDED.status_id,
  remarks =  EXCLUDED.remarks;
  

--   SELECT m.*, m.UUID, co.ID central_office_id, typ.ID ticket_type_id, m.TrnDate ticket_date, u.Id postedby_id, stat.Id status_id, '' remarks
--   FROM staging.TrnMap m 
--   LEFT JOIN Office co on co.Code = 'CI'
--   LEFT JOIN Reference typ on lower(typ.Ref_Type) = 'tickettype' and lower(typ.Title) = lower('Over the Counter')
--   LEFT JOIN Reference stat on lower(stat.Ref_Type) = 'ticketstatus' and lower(stat.Title) = lower('Completed')
--   LEFT JOIN Users u on lower(u.login_name) = lower(m.BrCode || '-' || trim(m.UserName))
-- where u.id is null


-- Insert to TrnHead from trnMaster
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
  (SELECT m.BrCode, m.BrCode || '-' || to_char(t.OrNo, 'fm00000000000000') || '-' || c.Id OrNo2, c.Id ticket_id, c.Ticket_Date trn_Date, c.postedby_id User_Id,
          Max(COALESCE(mp.CID,lm.CID)) CID, PrNo
   FROM staging.trnMap m
   INNER JOIN staging.trnmaster t on t.BrCode = m.BrCode and lower(trim(COALESCE(t.username,'SA'))) = lower(m.username) and t.trndate = m.trnDate
   INNER JOIN staging.lnMaster lm on t.BrCode = lm.BrCode and t.acc = lm.acc    
   INNER JOIN ticket c on c.Uuid = m.Uuid
   LEFT JOIN staging.multiplepaymentreceipt mp on mp.BrCode = t.BrCode and mp.OrNo = t.OrNo
   GROUP BY  m.BrCode, t.OrNo, c.postedby_id, PrNo, c.Id, c.Ticket_Date
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
   INNER JOIN staging.satrnmaster t on t.BrCode = m.BrCode and lower(trim(COALESCE(t.username,'SA'))) = lower(m.username) and t.trndate = m.trnDate
   INNER JOIN staging.saMaster lm on t.BrCode = lm.BrCode and t.acc = lm.acc    
   INNER JOIN ticket c on c.Uuid = m.Uuid
   LEFT JOIN staging.multiplepaymentreceipt mp on mp.BrCode = t.BrCode and mp.OrNo = t.OrNo
   GROUP BY  m.BrCode, COALESCE(t.OrNo,0), c.postedby_id, PrNo, c.Id, c.Ticket_Date
   ) t
LEFT JOIN vwoffice br on lower(br.Code) = lower(t.BrCode) and br.officetype = 'Area'
LEFT JOIN customer cus on cus.customer_alt_id = t.BrCode || '-' || t.CID
LEFT JOIN Reference typ on lower(typ.Ref_Type) = 'trnheadtype' and lower(typ.Title) = lower('Multiple Payment')
ON CONFLICT (Trn_Serial) DO nothing UPDATE SET 
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
   INNER JOIN staging.mutualfund t on t.BrCode = m.BrCode and lower(trim(COALESCE(t.username,'SA'))) = lower(m.username) and t.trndate = m.trnDate
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
  d.trn_prin, d.trn_int, d.waived_Int, d.Particular
FROM
  (SELECT
    th.id trn_head_id, Row_Number() OVER(Partition by t.OrNo Order by t.TrnDate) series,
    t.BrCode||to_char(t.trndate, 'YYYYMMDD')||'L'||t.Trn Alternate_Key, 
    t.TrnDate value_date, acc.Id account_id, t.TrnType Trn_Type_Code, 'PHP' currency, null item_id, false passbook_posted, 
    t.Prin trn_prin, t.Intr trn_int, t.waivedint waived_Int, t.Particular
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
  trn_prin, trn_int, waived_Int, Particular) ','/var/lib/postgresql/AccTran.csv');



COPY (
SELECT uuid_generate_v4() uuid, 
  d.trn_head_id, d.series, d.Alternate_Key,
  d.value_date, d.account_id, d.Trn_Type_Code, d.currency, d.item_id, d.passbook_posted, 
  d.trn_prin, d.trn_int, d.waived_Int, d.Particular
FROM
(SELECT
  th.id trn_head_id, Row_Number() OVER(Partition by t.OrNo Order by t.TrnDate) + 100000000 series, 
    t.BrCode||to_char(t.trndate, 'YYYYMMDD')||'S'||t.Trn Alternate_Key, 
  t.TrnDate value_date, acc.Id account_id, t.TrnType Trn_Type_Code, 'PHP' currency, null item_id, false passbook_posted, 
  CASE WHEN t.TrnType % 2 = 1 THEN - t.TrnAmt ELSE t.TrnAmt END trn_prin, 
  0 trn_int, 0 waived_Int, t.Particular
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
  trn_prin, trn_int, waived_Int, Particular) ','/var/lib/postgresql/AccsaTran.csv');
   

COPY (
SELECT uuid_generate_v4() uuid, 
  d.trn_head_id, d.series, d.Alternate_Key,
  d.value_date, d.account_id, d.Trn_Type_Code, d.currency, d.item_id, d.passbook_posted, 
  d.trn_prin, d.trn_int, d.waived_Int, d.Particular
FROM
(SELECT
  th.id trn_head_id, Row_Number() OVER(Partition by t.OrNo Order by t.TrnDate) + 200000000 series, 
    t.BrCode||to_char(t.trndate, 'YYYYMMDD')||'M'|| Row_Number() OVER(Partition by t.TrnDate Order by t.orno, t.CID) Alternate_Key, 
  t.TrnDate value_date, acc.Id account_id, t.TrnType Trn_Type_Code, 'PHP' currency, null item_id, false passbook_posted, 
  t.TrnAmt trn_prin, 0 trn_int, 0 waived_Int, '' Particular
FROM staging.trnMap m
INNER JOIN staging.mutualfund t on t.BrCode = m.BrCode and trim(lower(COALESCE(t.username,'SA'))) = trim(lower(m.username)) and t.trndate = m.trnDate
INNER JOIN ticket c on c.Uuid = m.Uuid
INNER JOIN trn_head th on th.Trn_Serial = m.BrCode || '-' || to_char(COALESCE(t.OrNo,0), 'fm00000000000000') || '-' || c.Id 
LEFT JOIN Account acc on acc.alternate_acc = t.BrCode ||'MBA'|| t.CID 
) d
LEFT JOIN Account_Tran accTrn on accTrn.Alternate_Key = d.Alternate_Key
where accTrn.uuid is null
)
TO '/var/lib/postgresql/AccMBATran.csv'  WITH DELIMITER '|' CSV HEADER;


select loaddata ('Account_Tran (uuid, trn_head_id, series, Alternate_Key,
  value_date, account_id, Trn_Type_Code, currency, item_id, passbook_posted, 
  trn_prin, trn_int, waived_Int, Particular) ','/var/lib/postgresql/AccMBATran.csv');
   

SELECT t.username, m.BrCode, m.BrCode || '-' || to_char(t.OrNo, 'fm00000000000000') OrNo2, c.Id ticket_id, max(c.Ticket_Date) trn_Date, c.postedby_id User_Id,
          Max(COALESCE(mp.CID,lm.CID)) CID, PrNo
   FROM staging.trnMap m
   INNER JOIN staging.trnmaster t on t.BrCode = m.BrCode and t.username = m.username and t.trndate = m.trnDate
   INNER JOIN staging.lnMaster lm on t.BrCode = lm.BrCode and t.acc = lm.acc    
   INNER JOIN ticket c on c.Uuid = m.Uuid
   LEFT JOIN staging.multiplepaymentreceipt mp on mp.BrCode = t.BrCode and mp.OrNo = t.OrNo
   where t.orno = 48086
   GROUP BY  m.BrCode, t.OrNo, c.postedby_id, PrNo, t.username, c.Id
   
  

  
COPY (
  SELECT 
    trn.account_id, trn.uuid, trn.trn_head_id, trn.series, trn.value_date,
    acc.principal-sum(COALESCE(p.trn_prin,0)) Bal_prin, acc.interest-sum(COALESCE(p.trn_int,0)) Bal_int
  FROM account_Tran trn
  INNER JOIN 
    (SELECT a.ID, Principal, COALESCE(ai.Interest,0) Interest
     FROM Account a
     LEFT JOIN Account_Interest ai on a.id = ai.account_id 
     ) acc on acc.ID = trn.account_id
  LEFT JOIN account_Tran p on p.account_id = trn.account_id and (p.value_date < trn.value_date or (p.value_date = trn.value_date and p.series <= trn.series)) and p.trn_type_code not in (3100,3400)
  GROUP BY trn.account_id, trn.uuid, trn.trn_head_id, trn.series, acc.Principal, acc.Interest, trn.value_date)
TO '/var/lib/postgresql/accBal.csv'  WITH DELIMITER '|' CSV HEADER;


select loaddata ('Account_Tran (uuid, trn_head_id, series, Alternate_Key,
  value_date, account_id, Trn_Type_Code, currency, item_id, passbook_posted, 
  trn_prin, trn_int, waived_Int, Particular) ','/var/lib/postgresql/AccMBATran.csv');
   