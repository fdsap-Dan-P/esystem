----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Interest (
----------------------------------------------------------------------------------------
  Account_ID bigint NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Interest numeric(16,6) NOT NULL DEFAULT 0,
  Effective_Rate float8 NOT NULL DEFAULT 0,
  Interest_Rate float8 NOT NULL DEFAULT 0,
  Credit numeric(16,6) NOT NULL DEFAULT 0,
  Debit numeric(16,6) NOT NULL DEFAULT 0,
  Accruals numeric(16,6) NOT NULL DEFAULT 0,
  Waived_Int numeric(16,6) NOT NULL DEFAULT 0,
  Last_Accrued_Date Date NULL,

  CONSTRAINT Account_Interest_pkey PRIMARY KEY (Account_ID),
  CONSTRAINT fk_Account_InterestID FOREIGN KEY (Account_ID) REFERENCES Account(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Interest_UUID ON public.Account_Interest(UUID);

DROP TRIGGER IF EXISTS trgAccount_Interest_Ins on Account_Interest;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Interest_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Interest
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAccount_Interest_upd on Account_Interest;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Interest_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Interest
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccount_Interest_del on Account_Interest;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Interest_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Interest
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

/*
INSERT into Account_Interest(
      Account_ID, Accruals, Credit, Debit, 
      Interest, Effective_Rate, Interest_Rate, 
      Last_Accrued_Date, Waived_Int) 
  
  SELECT 
      a.ID Account_ID, Acc.Accruals, Acc.Credit, Acc.Debit, 
      Acc.Interest, Acc.Effective_Rate, Acc.Interest_Rate, 
      cast(Acc.Last_Accrued_Date as Date), Acc.Waived_Int
  FROM (Values
      ('1001-0001-0000001',0, 0, 0, 1000, 56, 24, '01-01-2020', 0)
      )   
    Acc(
      Alternate_Acc, Accruals, Credit, Debit, 
      Effective_Rate, Interest, Interest_Rate, 
      Last_Accrued_Date, Waived_Int
      )
  LEFT JOIN Account a on a.Alternate_Acc = Acc.Alternate_Acc 
  ON CONFLICT(Account_ID) DO UPDATE SET
    Accruals = excluded.Accruals,
    Credit = excluded.Credit,
    Debit = excluded.Debit,
    Effective_Rate = excluded.Effective_Rate,
    Interest = excluded.Interest,
    Interest_Rate = excluded.Interest_Rate,
    Last_Accrued_Date = excluded.Last_Accrued_Date,
    Waived_Int = excluded.Waived_Int
;
*/