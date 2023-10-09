----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Access_Config_String (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Role_ID bigint NOT NULL,
  Config_Code varchar(50) NOT NULL,
  Config_ID bigint NOT NULL,
  Value text NOT NULL,
  
  CONSTRAINT Access_Config_String_pkey PRIMARY KEY (UUID),
  CONSTRAINT fk_Access_Config_role   FOREIGN KEY (Role_ID)   REFERENCES Access_Role(ID),
  CONSTRAINT fk_Access_Config_Config FOREIGN KEY (Config_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccess_Config_String_Unique ON public.Access_Config_String(Role_ID, Config_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccess_Config_String_Code ON public.Access_Config_String(Role_ID, lower(Config_Code));

DROP TRIGGER IF EXISTS trgAccess_Config_String_Ins on Access_Config_String;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Config_String_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Access_Config_String
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericConfigInsert();
 
DROP TRIGGER IF EXISTS trgAccess_Config_String_upd on Access_Config_String;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Config_String_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Access_Config_String
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericConfigUpdate();

DROP TRIGGER IF EXISTS trgAccess_Config_String_del on Access_Config_String;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Config_String_del
---------------------------------------------------------------------------
    AFTER DELETE ON Access_Config_String
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 