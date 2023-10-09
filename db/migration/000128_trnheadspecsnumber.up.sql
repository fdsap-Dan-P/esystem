----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Trn_Head_Specs_Number (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Trn_Head_ID bigint NOT NULL,
  Specs_Code VarChar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Value numeric(16,6) NOT NULL,
  Value2 numeric(16,6) NOT NULL,
  Measure_ID bigint NULL,

  CONSTRAINT Trn_Head_Specs_Number_pkey PRIMARY KEY (UUID),
  CONSTRAINT Trn_Head_Specs_Number_ID FOREIGN KEY (Trn_Head_ID) REFERENCES Trn_Head(ID),
  CONSTRAINT fk_Trn_Head_Specs_Unitmeasure FOREIGN KEY (Measure_ID) REFERENCES Reference(ID),
  CONSTRAINT Trn_Head_Specs_Number_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxTrn_Head_Specs_Number_Unique ON public.Trn_Head_Specs_Number(Trn_Head_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxTrn_Head_Specs_Number_Code ON public.Trn_Head_Specs_Number(Trn_Head_ID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgTrn_Head_Specs_Number_Ins on Trn_Head_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgTrn_Head_Specs_Number_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Trn_Head_Specs_Number
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgTrn_Head_Specs_Number_upd on Trn_Head_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgTrn_Head_Specs_Number_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Trn_Head_Specs_Number
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgTrn_Head_Specs_Number_del on Trn_Head_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgTrn_Head_Specs_Number_del
---------------------------------------------------------------------------
    AFTER DELETE ON Trn_Head_Specs_Number
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 
