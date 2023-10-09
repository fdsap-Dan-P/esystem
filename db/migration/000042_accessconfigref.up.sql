----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Access_Config_Ref (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Role_ID bigint NOT NULL,
  Config_Code varchar(50) NOT NULL,
  Config_ID bigint NOT NULL,
  Ref_ID bigint NOT NULL,
  
  CONSTRAINT Access_Config_Ref_pkey PRIMARY KEY (UUID),
  CONSTRAINT fk_Access_Config_role   FOREIGN KEY (Role_ID)   REFERENCES Access_Role(ID),
  CONSTRAINT fk_Access_Config_Config FOREIGN KEY (Config_ID) REFERENCES Reference(ID),
  CONSTRAINT Access_Config_Ref_Ref FOREIGN KEY (Ref_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccess_Config_Ref_Unique ON public.Access_Config_Ref(Role_ID, Config_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccess_Config_Ref_Code ON public.Access_Config_Ref(Role_ID, lower(Config_Code));

DROP TRIGGER IF EXISTS trgAccess_Config_Ref_Ins on Access_Config_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Config_Ref_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Access_Config_Ref
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericConfigInsert();
 
DROP TRIGGER IF EXISTS trgAccess_Config_Ref_upd on Access_Config_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Config_Ref_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Access_Config_Ref
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericConfigUpdate();

DROP TRIGGER IF EXISTS trgAccess_Config_Ref_del on Access_Config_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Config_Ref_del
---------------------------------------------------------------------------
    AFTER DELETE ON Access_Config_Ref
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 