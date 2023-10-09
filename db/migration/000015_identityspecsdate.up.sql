----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Identity_Specs_Date (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  IIID bigint NOT NULL,
  Specs_Code varchar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Value Date NOT NULL,
  Value2 Date NOT NULL,
  
  CONSTRAINT Identity_Specs_Date_pkey PRIMARY KEY (UUID),
  CONSTRAINT Identity_Specs_Date_ID FOREIGN KEY (IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT Identity_Specs_Date_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxIdentity_Specs_Date_Unique ON public.Identity_Specs_Date(IIID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxIdentity_Specs_Date_Code ON public.Identity_Specs_Date(IIID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgIdentity_Specs_Date_Ins on Identity_Specs_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgIdentity_Specs_Date_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Identity_Specs_Date
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgIdentity_Specs_Date_upd on Identity_Specs_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgIdentity_Specs_Date_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Identity_Specs_Date
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate(); 

DROP TRIGGER IF EXISTS trgIdentity_Specs_Date_del on Identity_Specs_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgIdentity_Specs_Date_del
---------------------------------------------------------------------------
    AFTER DELETE ON Identity_Specs_Date
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 
  