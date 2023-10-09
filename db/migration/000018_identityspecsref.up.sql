----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Identity_Specs_Ref (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  IIID bigint NOT NULL,
  Specs_Code varchar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Ref_ID bigint NOT NULL,
  
  CONSTRAINT Identity_Specs_Ref_pkey PRIMARY KEY (UUID),
  CONSTRAINT Identity_Specs_Ref_ID FOREIGN KEY (IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT Identity_Specs_Ref_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID),
  CONSTRAINT Identity_Specs_Ref_Ref FOREIGN KEY (Ref_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxIdentity_Specs_Ref_Unique ON public.Identity_Specs_Ref(IIID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxIdentity_Specs_Ref_Code ON public.Identity_Specs_Date(IIID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgIdentity_Specs_Ref_Ins on Identity_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgIdentity_Specs_Ref_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Identity_Specs_Ref
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgIdentity_Specs_Ref_upd on Identity_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgIdentity_Specs_Ref_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Identity_Specs_Ref
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgIdentity_Specs_Ref_del on Identity_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgIdentity_Specs_Ref_del
---------------------------------------------------------------------------
    AFTER DELETE ON Identity_Specs_Ref
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 