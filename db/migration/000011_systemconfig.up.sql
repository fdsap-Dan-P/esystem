---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.System_Config (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Office_ID bigint NOT NULL,
  GL_Date Date NOT NULL,
  Last_Accruals Date NOT NULL,
  Last_Month_End Date NOT NULL,
  Next_Month_End Date NOT NULL,
  System_Date Date NOT NULL,
  Run_State int2 NOT NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT System_Config_pkey PRIMARY KEY (Office_ID ),
  CONSTRAINT fkSystem_ConfigOffice FOREIGN KEY (Office_ID ) REFERENCES Office(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxSystem_Config_UUID ON public.System_Config(UUID);

DROP TRIGGER IF EXISTS trgSystem_ConfigIns on System_Config;
---------------------------------------------------------------------------
CREATE TRIGGER trgSystem_ConfigIns
---------------------------------------------------------------------------
    BEFORE INSERT ON System_Config
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgSystem_Configupd on System_Config;
---------------------------------------------------------------------------
CREATE TRIGGER trgSystem_Configupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON System_Config
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgSystem_Config_del on System_Config;
---------------------------------------------------------------------------
CREATE TRIGGER trgSystem_Config_del
---------------------------------------------------------------------------
    AFTER DELETE ON System_Config
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

  INSERT into System_Config(
    Office_ID , GL_Date, Last_Accruals, Last_Month_End, Next_Month_End, System_Date, Run_State
    )    
  SELECT
    o.ID Office_ID , cast('01/01/2020' as Date) GL_Date, cast('01/01/2020' as Date) Last_Accruals, 
    cast('01/01/2020' as Date) Last_Month_End, cast('01/01/2020' as Date) Next_Month_End, 
    cast('01/01/2020' as Date) System_Date, 0 Run_State
  FROM Office o WHERE o.Alternate_ID = '10019'

  ON CONFLICT(Office_ID ) DO UPDATE SET
    GL_Date = excluded.GL_Date,
    Last_Accruals = excluded.Last_Accruals,
    Last_Month_End = excluded.Last_Month_End,
    Next_Month_End = excluded.Next_Month_End,
    Run_State = excluded.Run_State,
    System_Date = excluded.System_Date
  ;  
