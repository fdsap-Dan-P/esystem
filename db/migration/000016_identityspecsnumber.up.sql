----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Identity_Specs_Number (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  IIID bigint NOT NULL,
  Specs_Code varchar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Value numeric(16,6) NOT NULL,
  Value2 numeric(16,6) NOT NULL,
  Measure_ID bigint NULL,

  CONSTRAINT Identity_Specs_Number_pkey PRIMARY KEY (UUID),
  CONSTRAINT Identity_Specs_Number_ID FOREIGN KEY (IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fk_Identity_Specs_Unitmeasure FOREIGN KEY (Measure_ID) REFERENCES Reference(ID),
  CONSTRAINT Identity_Specs_Number_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxIdentity_Specs_Number_Unique ON public.Identity_Specs_Number(IIID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxIdentity_Specs_Number_Code ON public.Identity_Specs_Date(IIID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgIdentity_Specs_Number_Ins on Identity_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgIdentity_Specs_Number_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Identity_Specs_Number
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgIdentity_Specs_Number_upd on Identity_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgIdentity_Specs_Number_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Identity_Specs_Number
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgIdentity_Specs_Number_del on Identity_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgIdentity_Specs_Number_del
---------------------------------------------------------------------------
    AFTER DELETE ON Identity_Specs_Number
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 
