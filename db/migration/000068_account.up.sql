----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account (
----------------------------------------------------------------------------------------
ID BIGSERIAL NOT NULL,
UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
Customer_ID bigint NOT NULL,
Acc varchar(21) NOT NULL,
Alternate_Acc varchar(30) NULL,
Account_Name varchar(100) NOT NULL,
Balance numeric(16,6) NOT NULL DEFAULT 0,
Non_Current numeric(16,6) NOT NULL DEFAULT 0,
Contract_Date Date NULL,
Credit numeric(16,6) NOT NULL DEFAULT 0,
Debit numeric(16,6) NOT NULL DEFAULT 0,
isOpen bool NULL,
isBudget bool NULL,
Last_Activity_Date Date NULL,
Open_Date Date NOT NULL DEFAULT CURRENT_Date,
Passbook_Line int2 NOT NULL DEFAULT 0,
Pending_Trn_Amt numeric(16,6) NOT NULL DEFAULT 0,
Principal numeric(16,6) NOT NULL DEFAULT 0,
Class_ID bigint NOT NULL,
Account_Type_ID bigint NOT NULL,
Budget_Account_ID bigint NULL,
Category_ID bigint NULL,
Currency varchar(3) NOT NULL,
Office_ID bigint NOT NULL,
Referredby_ID bigint NULL,
Status_Code integer NOT NULL,
Closed Boolean NOT NULL DEFAULT FALSE,
Remarks varchar(200) NULL,
Other_Info jsonb NULL,
  
  CONSTRAINT Account_pkey PRIMARY KEY (ID),
  CONSTRAINT idxAccount_Acc UNIQUE (Acc),
  CONSTRAINT idxAccount_altAcc UNIQUE (Alternate_Acc),
  
  CONSTRAINT fk_Account_Customer FOREIGN KEY (Customer_ID) REFERENCES Customer(ID),
  CONSTRAINT fk_Account_Account_Type FOREIGN KEY (Account_Type_ID) REFERENCES Account_Type(ID),
  CONSTRAINT fk_Account_Class FOREIGN KEY (Category_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Account_Category FOREIGN KEY (Class_ID) REFERENCES Account_Class(ID),
  CONSTRAINT fk_Account_Budget_Account FOREIGN KEY (Budget_Account_ID) REFERENCES Account(ID),
  CONSTRAINT fk_Account_Office FOREIGN KEY (Office_ID ) REFERENCES Office(ID),
  CONSTRAINT fk_Account_referedby FOREIGN KEY (Referredby_ID) REFERENCES Identity_Info(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_UUID ON public.Account(UUID);
CREATE INDEX IF NOT EXISTS idxAccount_Status ON public.Account(Status_Code);


DROP TRIGGER IF EXISTS trgAccount_Ins on Account;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();

  
DROP TRIGGER IF EXISTS trgAccount_upd on Account;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccount_del on Account;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


DROP TRIGGER IF EXISTS trgAccount_del on Account;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwAccount
----------------------------------------------------------------------------------------
AS SELECT 
    dt.ID, mr.UUID,
    cust.Central_Office_ID, cust.CID,     
    cust.ID Customer_ID, cust.Customer_Alt_ID, 
    
    ii.Birth_Place_ID, ii.Last_Name, ii.First_Name, ii.Middle_Name,
    
    dt.Acc, dt.Alternate_Acc, dt.Account_Name,
    
    o.ID Office_ID , o.UUID OfficeUUID, o.Office_Name,
    
    atyp.ID Account_Type_ID, atyp.UUID Account_TypeUUID, atyp.Account_Type,   
    
    ac.Product_ID, ac.Product_Name,
    
    ac.Group_ID, ac.GroupUUID, ac.AccountGroup,

    ac.cur_id, ac.CurUUID, ac.CurrentAcc, ac.CurrentTitle,
    ac.Noncur_id, ac.NonCurUUID, ac.Non_CurrentAcc, ac.Non_CurrentTitle,

    ac.BS_Acc_ID, ac.BS_AccUUID, ac.BS_Acc, ac.BSTitle, 
    ac.IS_Acc_ID, ac.IS_AccUUID, ac.IS_Acc, ac.ISTitle, 
    
    ac.Class_ID, ac.Account_Class,    
    
    dt.Currency, 

    stat.ID Status_ID , stat.UUID StatusUUID, stat.Title Status,
       
    refby.ID Referredby_ID, refby.UUID ReferredbyUUID, 
    refby.Last_Name ReferredbyLast_Name, refby.First_Name ReferredbyFirst_Name,     
    
    budAcc.ID Budget_Account_ID, budAcc.UUID Budget_AccountUUID, budAcc.Acc Budget_AccountAcc,    
    
    
    dt.isBudget,
    
    dt.Principal,
    dt.Debit, 
    dt.Credit,
    dt.Balance,
    dt.Pending_Trn_Amt,
    dt.Passbook_Line,
    dt.Contract_Date,
    dt.Open_Date,
    dt.Last_Activity_Date,
    dt.Remarks,
    
    mr.Mod_Ctr,
    dt.Other_Info,
    mr.Created,
    mr.Updated 
    
   FROM Account dt
   JOIN Main_Record mr ON dt.UUID = mr.UUID
   JOIN Customer cust ON cust.ID = dt.Customer_ID
   JOIN Identity_Info ii ON ii.ID = cust.IIID
   JOIN Account_Type atyp ON atyp.ID = dt.Account_Type_ID
   JOIN Office o on o.ID = dt.Office_ID 

   LEFT JOIN vwAccount_Class ac on ac.ID = dt.Class_ID

   LEFT JOIN Reference stat on stat.Code = dt.Status_Code and lower(stat.Ref_Type) = 'accountstatus'
   LEFT JOIN Identity_Info refby on refby.ID = dt.Referredby_ID
   LEFT JOIN Account budAcc on budAcc.ID = dt.Budget_Account_ID
   ; 
 
 /*
 ---------------------------------------------------------------------------------------- 
  INSERT INTO Account(
---------------------------------------------------------------------------------------- 
    Customer_ID, Acc, Alternate_Acc, Account_Name, 
    Balance, Non_Current, Contract_Date, Credit, Debit, isBudget, 
    Last_Activity_Date, Open_Date, Passbook_Line, Pending_Trn_Amt, 
    Principal, Class_ID, Account_Type_ID, Budget_Account_ID, 
    Currency, Office_ID , Referredby_ID, Status_ID 
    ) 
    
SELECT
    cust.ID Customer_ID, Acc.Acc, Acc.Alternate_Acc, Acc.Account_Name, 
    Acc.Balance, Acc.Non_Current, cast(Acc.Contract_Date as Date), Acc.Credit, Acc.Debit, Acc.isBudget, 
    cast(Acc.Last_Activity_Date as Date), cast(Acc.Open_Date as Date), Acc.Passbook_Line, Acc.Pending_Trn_Amt, 
    Acc.Principal, ac.ID Class_ID, y.ID Account_Type_ID, bud.ID Budget_Account_ID, 
    acc.Currency, o.ID Office_ID, refby.ID Referredby_ID, stat.ID Status_ID 
  FROM (Values
      ('10001','1001-0001-0000001','1001-0001-0000001','United Church of Christ in the Philippines',0,0,'01/01/2020',0,0,FALSE,'01/01/2020','01/01/2020',1,0,0,'Current','Sikap 1',null,'PHP','10019','1050','Current'),
      ('10001','1001-0001-0000002','1001-0001-0000002','United Church of Christ in the Philippines',100,0,'01/01/2020',0,100,FALSE,'01/01/2020','01/01/2020',1,0,0,'Current','Sikap 1',null,'PHP','10019','1050','Current')
      )
  Acc(
    Customer_Alt_ID, Acc, Alternate_Acc, Account_Name, Balance, Non_Current, Contract_Date, Credit, Debit,
    isBudget, Last_Activity_Date, Open_Date, Passbook_Line, Pending_Trn_Amt, Principal,
    Account_Class, Account_Type, Budget_AccountAltID, Currency, OfficeAltID, ReferredbyAltID, Status 
    )

  LEFT JOIN Customer cust on cust.Customer_Alt_ID = Acc.Customer_Alt_ID
  LEFT JOIN Account_Type y on y.Account_Type = Acc.Account_Type 
  LEFT JOIN vwReference acr on acr.Title = Acc.Account_Class and acr.Ref_Type  = 'LoanClass'
  LEFT JOIN Account_Class ac 
    on  ac.Class_ID = acr.ID 
    and ac.Product_ID = y.Product_ID 
    and ac.Group_ID = y.Group_ID
  LEFT JOIN Account bud on bud.Alternate_Acc = Acc.Budget_AccountAltID
  LEFT JOIN Office o on o.Alternate_ID = Acc.OfficeAltid
  LEFT JOIN Identity_Info refby on refby.Alternate_ID = ReferredbyAltID
  LEFT JOIN vwReference stat on stat.Title = Acc.Status and stat.Ref_Type  = 'LoanStatus'
  
  ON CONFLICT(Alternate_Acc) DO UPDATE SET 
    Customer_ID = excluded.Customer_ID,
    Acc = excluded.Acc,
    Alternate_Acc = excluded.Alternate_Acc,
    Account_Name = excluded.Account_Name,
    Balance = excluded.Balance,
    Non_Current = excluded.Non_Current,
    Contract_Date = excluded.Contract_Date,
    Credit = excluded.Credit,
    Debit = excluded.Debit,
    isBudget = excluded.isBudget,
    Last_Activity_Date = excluded.Last_Activity_Date,
    Open_Date = excluded.Open_Date,
    Passbook_Line = excluded.Passbook_Line,
    Pending_Trn_Amt = excluded.Pending_Trn_Amt,
    Principal = excluded.Principal,
    Class_ID = excluded.Class_ID,
    Account_Type_ID = excluded.Account_Type_ID,
    Budget_Account_ID = excluded.Budget_Account_ID,
    Currency = excluded.Currency,
    Office_ID  = excluded.Office_ID ,
    Referredby_ID = excluded.Referredby_ID,
    Status_ID  = excluded.Status_ID   
  ;
*/

----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS "entries" (
----------------------------------------------------------------------------------------
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" numeric(16,6) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
 CONSTRAINT fk_entries_acc FOREIGN KEY (account_id) REFERENCES account(ID)
);

----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS "transfers" (
----------------------------------------------------------------------------------------
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" numeric(16,6) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
 CONSTRAINT fk_transfers_facc FOREIGN KEY (from_account_id) REFERENCES account(ID),
 CONSTRAINT fk_transfers_tacc FOREIGN KEY (to_account_id) REFERENCES account(ID)
);

CREATE INDEX IF NOT EXISTS idxentries_acc ON public.entries(Account_ID);

CREATE INDEX IF NOT EXISTS idxtransfers_facc ON public.transfers(from_account_id);
CREATE INDEX IF NOT EXISTS idxtransfers_tacc ON public.transfers(to_account_id);
CREATE INDEX IF NOT EXISTS idxtransfers_ftacc ON public.transfers(from_account_id, to_account_id);

-- ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

-- ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "account" ("id");

-- ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "account" ("id");

--CREATE INDEX IF NOT EXISTS ON "account" (Customer_ID);

--CREATE INDEX IF NOT EXISTS ON "entries" (Account_ID);

--CREATE INDEX IF NOT EXISTS ON "transfers" ("from_account_id");

--CREATE INDEX IF NOT EXISTS ON "transfers" ("to_account_id");

--CREATE INDEX IF NOT EXISTS ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';
