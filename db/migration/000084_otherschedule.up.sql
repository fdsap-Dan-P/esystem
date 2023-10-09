---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Other_Schedule (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Account_ID bigint NOT NULL,
  Charge_ID bigint NOT NULL,
  Series int2 NOT NULL,
  Due_Date  Date NOT NULL,
  Due_Amt  numeric(16,6) NOT NULL DEFAULT 0,
  Realizable numeric(16,6) NOT NULL DEFAULT 0,
  End_Bal numeric(16,6) NOT NULL DEFAULT 0,
  Other_Info jsonb NULL,
  
  CONSTRAINT Other_Schedule_pkey PRIMARY KEY (Account_ID, Charge_ID, Series),
  CONSTRAINT fk_Other_Schedule_Account FOREIGN KEY (Account_ID) REFERENCES Account(ID),
  CONSTRAINT fk_Other_Schedule_chrg FOREIGN KEY (Charge_ID) REFERENCES Charge_Type(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxOther_Schedule_UUID ON public.Other_Schedule(UUID);

DROP TRIGGER IF EXISTS trgOther_Schedule_Ins on Other_Schedule;
---------------------------------------------------------------------------
CREATE TRIGGER trgOther_Schedule_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Other_Schedule
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgOther_Schedule_upd on Other_Schedule;
---------------------------------------------------------------------------
CREATE TRIGGER trgOther_Schedule_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Other_Schedule
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgOther_Schedule_del on Other_Schedule;
---------------------------------------------------------------------------
CREATE TRIGGER trgOther_Schedule_del
---------------------------------------------------------------------------
    AFTER DELETE ON Other_Schedule
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


/*
INSERT into Other_Schedule(Account_ID, Charge_ID, Series, Due_Date , Due_Amt , Realizable, End_Bal)    
  SELECT
    a.ID Account_ID, chrg.ID Charge_ID, Series, cast(Due_Date  as Date), Due_Amt , Realizable, End_Bal
  FROM (Values
      ('1001-0001-0000001','Service Fee',1,'01/08/2020',20,20,80),
      ('1001-0001-0000001','Service Fee',2,'01/15/2020',20,20,60),
      ('1001-0001-0000001','Service Fee',3,'01/22/2020',20,20,40),
      ('1001-0001-0000001','Service Fee',4,'01/29/2020',20,20,20),
      ('1001-0001-0000001','Service Fee',5,'02/05/2020',20,20,0)
      )   
  Acc(
    Alternate_Acc, Charge, Series, Due_Date , Due_Amt , Realizable, End_Bal)
    
  LEFT JOIN Account a on a.Alternate_Acc = Acc.Alternate_Acc
  LEFT JOIN Charge_Type chrg on chrg.Charge_Type = Acc.Charge

  ON CONFLICT(Account_ID, Charge_ID, Series) DO UPDATE SET
    Account_ID = excluded.Account_ID,
    Charge_ID = excluded.Charge_ID,
    Series = excluded.Series,
    Due_Date  = excluded.Due_Date ,
    Due_Amt  = excluded.Due_Amt ,
    Realizable = excluded. Realizable,
    End_Bal = excluded.End_Bal
  ;  
*/