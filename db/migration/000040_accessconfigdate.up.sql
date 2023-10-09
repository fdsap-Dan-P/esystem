----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Access_Config_Date (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Role_ID bigint NOT NULL,
  Config_Code varchar(50) NOT NULL,
  Config_ID bigint NOT NULL,
  Value Date NOT NULL,
  Value2 Date NOT NULL,
  
  CONSTRAINT Access_Config_Date_pkey PRIMARY KEY (UUID),
  CONSTRAINT fk_Access_Config_Role   FOREIGN KEY (Role_ID)   REFERENCES Access_Role(ID),
  CONSTRAINT fk_Access_Config_Config FOREIGN KEY (Config_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccess_Config_Date_Unique ON public.Access_Config_Date(Role_ID, Config_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccess_Config_Date_Code ON public.Access_Config_Date(Role_ID, lower(Config_Code));

DROP TRIGGER IF EXISTS trgAccess_Config_Date_Ins on Access_Config_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Config_Date_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Access_Config_Date
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericConfigInsert();
 
DROP TRIGGER IF EXISTS trgAccess_Config_Date_upd on Access_Config_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Config_Date_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Access_Config_Date
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericConfigUpdate();

DROP TRIGGER IF EXISTS trgAccess_Config_Date_del on Access_Config_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Config_Date_del
---------------------------------------------------------------------------
    AFTER DELETE ON Access_Config_Date
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 
/*

  INSERT INTO Access_Config(Role_ID, Config_ID, Value_Int, Value_Decimal, Value_Date, Value_String)
  SELECT 
    r.ID Role_ID, c.ID Config_ID, 
    a.Value Value_Int, null Value_Decimal, null Value_Date, null Value_String
  FROM (Values
      ('Admin', 'PasswordComplexity', 4),
      ('Bookkeeper', 'PasswordComplexity', 4),
      ('Cashier', 'PasswordComplexity', 4),
      ('Pastor', 'PasswordComplexity', 4),
      ('Member', 'PasswordComplexity', 4)
      )   
    a(Access_Name, Access_Config, Value), vwReference c, Access_Role r
   WHERE c.Ref_Type = 'Access_Config' and c.Title = a.Access_Config
     and r.Access_Name = a.Access_Name
   ORDER BY 1
  ON CONFLICT(Role_ID, Config_ID)
  DO UPDATE SET 
    Value_Int = EXCLUDED.Value_Int, 
    Value_Decimal = EXCLUDED.Value_Decimal, 
    Value_Date = EXCLUDED.Value_Date, 
    Value_String = EXCLUDED.Value_String;

*/