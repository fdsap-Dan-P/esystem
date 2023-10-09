---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Office_Account (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Office_ID     bigint NOT NULL,
  Type_ID      bigint NOT NULL,
  Currency  varchar(3) NOT NULL,
  Partition_ID bigint NULL,
  Balance       numeric(16,6) NOT NULL DEFAULT 0,
  Pending_Trn_Amt numeric(16,6) NOT NULL DEFAULT 0,
  Budget        numeric(16,6) NOT NULL DEFAULT 0,
  Last_Activity_Date  Date NULL,
  Status_Code    integer NOT NULL,
  Remarks varchar(200) NULL,
  Other_Info jsonb NULL,
 
  CONSTRAINT Office_Account_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Office_Account_Office FOREIGN KEY (Office_ID ) REFERENCES Office(ID),
  CONSTRAINT fk_Office_Account_Office_Account_Type FOREIGN KEY (Type_ID) REFERENCES Office_Account_Type(ID),
  CONSTRAINT fk_Office_Account_Partition FOREIGN KEY (Partition_ID ) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxOffice_Account_UUID ON public.Office_Account(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxOffice_Account_Status ON public.Office_Account(Status_Code);
CREATE UNIQUE INDEX IF NOT EXISTS idxOffice_Account_Unq 
   on Office_Account(Office_ID , Type_ID, lower(Currency), COALESCE(Partition_ID ,0 ));

DROP TRIGGER IF EXISTS trgOffice_Account_Ins on Office_Account;
---------------------------------------------------------------------------
CREATE TRIGGER trgOffice_Account_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Office_Account
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgOffice_Account_upd on Office_Account;
---------------------------------------------------------------------------
CREATE TRIGGER trgOffice_Account_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Office_Account
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgOffice_Account_del on Office_Account;
---------------------------------------------------------------------------
CREATE TRIGGER trgOffice_Account_del
---------------------------------------------------------------------------
    AFTER DELETE ON Office_Account
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

  
   INSERT into Office_Account(
    Office_ID , Type_ID, Currency, Partition_ID , 
    Balance, Pending_Trn_Amt, Budget, Last_Activity_Date, 
    Status_Code , Remarks
    ) 
 
  SELECT
    o.ID Office_ID , y.ID Type_ID, acc.Currency, par.ID Partition_ID , 
    Balance, Pending_Trn_Amt, Budget, cast(Last_Activity_Date as Date),
    stat.Code Status_Code , Remarks
  FROM (Values
      ('10019', 'Cash', 'PHP', 'FundSource', 'GSB', 100, 0, 200, '01/01/2020', 30, 'Remarks')
      )   
  Acc(
    Office_altid, Office_Account_Type, Currency, Partition_Type, Partition_Title, 
    Balance, Pending_Trn_Amt, Budget, Last_Activity_Date, 
    StatusCode, Remarks
    )
    
  LEFT JOIN Office o on o.Alternate_ID = Acc.Office_altid
  LEFT JOIN Office_Account_Type y on y.Office_Account_Type = Acc.Office_Account_Type

  LEFT JOIN Reference stat on stat.code = acc.StatusCode and lower(stat.Ref_Type) = 'accountstatus'
  LEFT JOIN vwReference par on par.Ref_Type = Acc.Partition_Type and par.Title = Acc.Partition_Title
  
  ON CONFLICT(Office_ID , Type_ID, lower(Currency), COALESCE(Partition_ID ,0 )) DO UPDATE SET
    Office_ID  = excluded.Office_ID ,
    Type_ID = excluded.Type_ID,
    Currency = excluded.Currency,
    Balance = excluded.Balance,
    Pending_Trn_Amt = excluded.Pending_Trn_Amt,
    Budget = excluded.Budget,
    Last_Activity_Date = excluded.Last_Activity_Date,
    Partition_ID  = excluded.Partition_ID,
    Status_Code  = excluded.Status_Code,
    Remarks = excluded.Remarks
  ;