---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Tran (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Trn_Head_ID bigint NOT NULL,
  Series bigint NOT NULL,
  Alternate_Key varchar(20) NULL,
  Value_Date Date NOT NULL,
  Account_ID bigint NOT NULL,
  Trn_Type_Code SmallInt NOT NULL,
  Currency varchar(3) NOT NULL,
  Item_ID bigint NULL,
  Passbook_Posted bool NOT NULL,
  Trn_Prin numeric(16,6) NOT NULL DEFAULT 0,
  Trn_Int numeric(16,6) NOT NULL DEFAULT 0,
  Waived_Int numeric(16,6) NOT NULL DEFAULT 0,
  Bal_Prin numeric(16,6) NOT NULL DEFAULT 0,
  Bal_Int numeric(16,6) NOT NULL DEFAULT 0,
  Particular varchar(300) NOT NULL DEFAULT '',
  Cancelled bool NOT NULL DEFAULT false,
  Other_Info jsonb NULL,
  
  CONSTRAINT Account_Tran_pkey PRIMARY KEY (Trn_Head_ID, Series),
  CONSTRAINT fkAccount_TranTrn_Head FOREIGN KEY (Trn_Head_ID) REFERENCES Trn_Head(ID),
  CONSTRAINT fkAccount_TranAccount FOREIGN KEY (Account_ID) REFERENCES Account(ID),
  CONSTRAINT fkAccount_TranItem FOREIGN KEY (Item_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Tran_UUID ON public.Account_Tran(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Tran_Alt ON public.Account_Tran(Alternate_Key);
CREATE INDEX IF NOT EXISTS idxAccount_Tran_Acc ON public.Account_Tran(Account_ID);
CREATE INDEX IF NOT EXISTS idxAccount_Tran_Code ON public.Account_Tran(Account_ID, Trn_Type_Code);

DROP TRIGGER IF EXISTS trgAccount_TranIns on Account_Tran;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_TranIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Tran
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();

DROP TRIGGER IF EXISTS trgAccount_Tranupd on Account_Tran;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Tranupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Tran
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccount_Tran_del on Account_Tran;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Tran_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Tran
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

/*

  INSERT INTO Account_Tran(
    Trn_Head_ID, Series, Value_Date, Account_ID, Currency, Item_ID, Passbook_Posted, Trn_Prin, Trn_Int)
  SELECT 
    h.ID Trn_Head_ID, a.Series, Value_Date, Acc.ID Account_ID, a.Currency, i.ID Item_ID, Passbook_Posted, Trn_Prin, Trn_Int
    
   FROM (Values
      ('2af90d74-3bee-48c5-8935-443edafb8f5a'::UUID, 1, '01-01-2020'::Date, '1001-0001-0000001',  'PHP', 'General Assembly Meeting', TRUE, 0, 0)
      )   
    a(Trn_HeadUUID, Series, Value_Date, Alternate_Acc, Currency, Item, Passbook_Posted, Trn_Prin, Trn_Int)  

  LEFT JOIN Trn_Head h on h.UUID = a.Trn_HeadUUID
  LEFT JOIN vwReference i    on lower(i.Title) = lower(a.Item) and lower(i.Ref_Type) = 'churchfunditem'
  LEFT JOIN Account Acc      on Acc.Alternate_Acc = a.Alternate_Acc

  ON CONFLICT(Trn_Head_ID, Series)
  DO UPDATE SET
    Value_Date = excluded.Value_Date,
    Trn_Head_ID = excluded.Trn_Head_ID,
    Account_ID = excluded.Account_ID,
    Currency = excluded.Currency,
    Passbook_Posted = excluded.Passbook_Posted,
    Trn_Prin = excluded.Trn_Prin,
    Trn_Int = excluded.Trn_Int
  ;   
*/


---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.TempAccBal (
---------------------------------------------------------------------------
  Account_ID bigint NOT NULL,
  UUID uuid NOT NULL,
  Trn_Head_ID bigint NOT NULL,
  Series bigint NOT NULL,
  Value_Date Date NOT NULL,
  Bal_Prin numeric(16,6) NOT NULL DEFAULT 0,
  Bal_Int numeric(16,6) NOT NULL DEFAULT 0,
  
  CONSTRAINT TempAccBal_pkey PRIMARY KEY (UUID)
);

/*


DECLARE
  curTrn CURSOR FOR 
    SELECT uuid, trn_prin, trn_int FROM account_Tran 
    WHERE account_id = 249518 and trn_type_code not in (3100,3400)  
    ORDER BY account_id, value_date, series
    FOR UPDATE;
  
    rowVar RECORD;
    balPrin Numeric = 0;
    balInt Numeric = 0;
    AccId bigint = 0;
BEGIN
    OPEN curTrn;
    LOOP
        FETCH curTrn INTO rowVar;
        EXIT WHEN NOT FOUND;

        IF rowVar.account_id <> AccId
        BEGIN
          SELECT Principal, Interest into balPrin, BalInt
          FROM
            (SELECT a.ID, Principal, COALESCE(ai.Interest,0) Interest
             FROM Account a
             LEFT JOIN Account_Interest ai on a.id = ai.account_id 
             WHERE a.ID = rowVar.Account_Id
             ) acc
        END;
        -- Update the column
        balPrin := balPrin + rowVar.trn_Prin ;
        balInt := balInt + rowVar.trn_Int ;

        -- Perform additional modifications if needed
        -- Update the row in the table
        UPDATE account_Tran SET bal_Prin = balPrin, bal_Int = balInt WHERE uuid = rowVar.uuid;
    END LOOP;

    CLOSE curTrn;
END;


  SELECT 
    trn.account_id, trn.uuid, trn.trn_head_id, trn.series, trn.value_date,
    acc.principal-sum(COALESCE(p.trn_prin,0)) Bal_prin, acc.interest-sum(COALESCE(p.trn_int,0)) Bal_int
  FROM account_Tran trn
  */
