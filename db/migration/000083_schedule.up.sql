---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Schedule (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Account_ID bigint NOT NULL,
  Series int2 NOT NULL,
  Due_Date  Date NOT NULL,
  Due_Prin  numeric(16,6) NOT NULL,
  Due_Int  numeric(16,6) NOT NULL,
  End_Prin  numeric(16,6) NOT NULL,
  End_Int  numeric(16,6) NOT NULL,
  Carrying_Value  numeric(16,6) NOT NULL,
  Realizable numeric(16,6) NOT NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT Schedule_pkey PRIMARY KEY (Account_ID, Series),
  CONSTRAINT fk_Schedule_Acc FOREIGN KEY (Account_ID) REFERENCES Account(ID)
 );
CREATE UNIQUE INDEX IF NOT EXISTS idxSchedule_UUID ON public.Schedule(UUID);
CREATE INDEX IF NOT EXISTS idxSchedule_Acc ON public.Schedule USING btree (Account_ID);

DROP TRIGGER IF EXISTS trgSchedule_Ins on Schedule;
---------------------------------------------------------------------------
CREATE TRIGGER trgSchedule_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Schedule
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgSchedule_upd on Schedule;
---------------------------------------------------------------------------
CREATE TRIGGER trgSchedule_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Schedule
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgSchedule_del on Schedule;
---------------------------------------------------------------------------
CREATE TRIGGER trgSchedule_del
---------------------------------------------------------------------------
    AFTER DELETE ON Schedule
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

/*
  INSERT into Schedule(
    Account_ID, Series, Due_Date , 
    Due_Prin , Due_Int , End_Prin , End_Int , 
    Carrying_Value , Realizable
    )    
  SELECT
    a.ID Account_ID, Series, cast(Due_Date  as Date), 
    Due_Prin , Due_Int , End_Prin , End_Int , 
    Carrying_Value , Realizable
  FROM (Values
      ('1001-0001-0000001',1,'01/08/2020',100,20,2400,480,2400,20),
      ('1001-0001-0000001',2,'01/15/2020',100,20,2300,460,2300,20),
      ('1001-0001-0000001',3,'01/22/2020',100,20,2200,440,2200,20),
      ('1001-0001-0000001',4,'01/29/2020',100,20,2100,420,2100,20),
      ('1001-0001-0000001',5,'02/05/2020',100,20,2000,400,2000,20),
      ('1001-0001-0000001',6,'02/12/2020',100,20,1900,380,1900,20),
      ('1001-0001-0000001',7,'02/19/2020',100,20,1800,360,1800,20),
      ('1001-0001-0000001',8,'02/26/2020',100,20,1700,340,1700,20),
      ('1001-0001-0000001',9,'03/04/2020',100,20,1600,320,1600,20),
      ('1001-0001-0000001',10,'03/11/2020',100,20,1500,300,1500,20),
      ('1001-0001-0000001',11,'03/18/2020',100,20,1400,280,1400,20),
      ('1001-0001-0000001',12,'03/25/2020',100,20,1300,260,1300,20),
      ('1001-0001-0000001',13,'04/01/2020',100,20,1200,240,1200,20),
      ('1001-0001-0000001',14,'04/08/2020',100,20,1100,220,1100,20),
      ('1001-0001-0000001',15,'04/15/2020',100,20,1000,200,1000,20),
      ('1001-0001-0000001',16,'04/22/2020',100,20,900,180,900,20),
      ('1001-0001-0000001',17,'04/29/2020',100,20,800,160,800,20),
      ('1001-0001-0000001',18,'05/06/2020',100,20,700,140,700,20),
      ('1001-0001-0000001',19,'05/13/2020',100,20,600,120,600,20),
      ('1001-0001-0000001',20,'05/20/2020',100,20,500,100,500,20),
      ('1001-0001-0000001',21,'05/27/2020',100,20,400,80,400,20),
      ('1001-0001-0000001',22,'06/03/2020',100,20,300,60,300,20),
      ('1001-0001-0000001',23,'06/10/2020',100,20,200,40,200,20),
      ('1001-0001-0000001',24,'06/17/2020',100,20,100,20,100,20),
      ('1001-0001-0000001',25,'06/24/2020',100,20,0,0,0,20)
      )   
  Acc(
    Alternate_Acc, Series, Due_Date , 
    Due_Prin , Due_Int , End_Prin , End_Int , 
    Carrying_Value , Realizable)
    
  LEFT JOIN Account a on a.Alternate_Acc = Acc.Alternate_Acc

  ON CONFLICT(Account_ID, Series) DO UPDATE SET
    Account_ID = excluded.Account_ID,
    Series = excluded.Series,
    Due_Date  = excluded.Due_Date ,
    Due_Prin  = excluded.Due_Prin ,
    Due_Int  = excluded.Due_Int ,
    End_Prin  = excluded.End_Prin ,
    End_Int  = excluded.End_Int ,
    Carrying_Value  = excluded.Carrying_Value ,
    Realizable = excluded.Realizable,
    Other_Info = excluded.Other_Info
  ;  
*/