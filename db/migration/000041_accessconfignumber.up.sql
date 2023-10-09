----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Access_Config_Number (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Role_ID bigint NOT NULL,
  Config_Code varchar(50) NOT NULL,
  Config_ID bigint NOT NULL,
  Value numeric(16,6) NOT NULL,
  Value2 numeric(16,6) NOT NULL,
  Measure_ID bigint NULL,

  CONSTRAINT Access_Config_Number_pkey PRIMARY KEY (UUID),
  CONSTRAINT fk_Access_Config_role   FOREIGN KEY (Role_ID)   REFERENCES Access_Role(ID),
  CONSTRAINT fk_Access_Config_Config FOREIGN KEY (Config_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccess_Config_Number_Unique ON public.Access_Config_Number(Role_ID, Config_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccess_Config_Number_Code ON public.Access_Config_Number(Role_ID, lower(Config_Code));

DROP TRIGGER IF EXISTS trgAccess_Config_Number_Ins on Access_Config_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Config_Number_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Access_Config_Number
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericConfigInsert();
 
DROP TRIGGER IF EXISTS trgAccess_Config_Number_upd on Access_Config_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Config_Number_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Access_Config_Number
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericConfigUpdate();

DROP TRIGGER IF EXISTS trgAccess_Config_Number_del on Access_Config_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Config_Number_del
---------------------------------------------------------------------------
    AFTER DELETE ON Access_Config_Number
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 
 