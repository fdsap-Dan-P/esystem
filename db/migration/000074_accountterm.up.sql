----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Term (
----------------------------------------------------------------------------------------
  Account_ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Frequency int2 NOT NULL,
  N int2 NOT NULL,
  Paid_N int2 NOT NULL,
  Fixed_Due numeric(16,6) NOT NULL,
  Cummulative_Due numeric(16,6) NOT NULL,
  Date_Start Date NOT NULL,
  Maturity Date NOT NULL,
  
  CONSTRAINT Account_Term_pkey PRIMARY KEY (Account_ID),
  CONSTRAINT fkAccount_Term_Acc FOREIGN KEY (Account_ID) REFERENCES Account(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Term_UUID ON public.Account_Term(UUID);

DROP TRIGGER IF EXISTS trgAccount_Term_Ins on Account_Term;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Term_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Term
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAccount_Term_upd on Account_Term;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Term_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Term
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccount_Term_del on Account_Term;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Term_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Term
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

/*
INSERT into Account_Term(
      Account_ID, frequency, n, Paid_N, 
      Fixed_Due, Cummulative_Due, Date_Start, Maturity) 
  
  SELECT 
      a.ID Account_ID, frequency, n, Paid_N, 
      Fixed_Due, Cummulative_Due, CAST(Date_Start as Date), CAST(Maturity as Date)
  FROM (Values
      ('1001-0001-0000001', 50, 25, 0, 0, 0, '01-01-2020', '01-01-2021')
      )   
    Acc(
      Alternate_Acc, Frequency, n, Paid_N, 
      Fixed_Due, Cummulative_Due, Date_Start, Maturity
      )
  LEFT JOIN Account a on a.Alternate_Acc = Acc.Alternate_Acc 
  ON CONFLICT(Account_ID) DO UPDATE SET
    frequency = excluded.frequency,
    n = excluded.n,
    Paid_N = excluded.Paid_N,
    Fixed_Due = excluded.Fixed_Due,
    Cummulative_Due = excluded.Cummulative_Due,
    Date_Start = excluded.Date_Start,
    Maturity = excluded.Maturity
    ;
  */