----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Trn_Head_Specs_Ref (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Trn_Head_ID bigint NOT NULL,
  Specs_Code VarChar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Ref_ID bigint NOT NULL,
  
  CONSTRAINT Trn_Head_Specs_Ref_pkey PRIMARY KEY (UUID),
  CONSTRAINT Trn_Head_Specs_Ref_ID FOREIGN KEY (Trn_Head_ID) REFERENCES Trn_Head(ID),
  CONSTRAINT Trn_Head_Specs_Ref_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID),
  CONSTRAINT Trn_Head_Specs_Ref_Ref FOREIGN KEY (Ref_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxTrn_Head_Specs_Ref_Unique ON public.Trn_Head_Specs_Ref(Trn_Head_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxTrn_Head_Specs_Ref_Code ON public.Trn_Head_Specs_Ref(Trn_Head_ID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgTrn_Head_Specs_Ref_Ins on Trn_Head_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgTrn_Head_Specs_Ref_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Trn_Head_Specs_Ref
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgTrn_Head_Specs_Ref_upd on Trn_Head_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgTrn_Head_Specs_Ref_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Trn_Head_Specs_Ref
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgTrn_Head_Specs_Ref_del on Trn_Head_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgTrn_Head_Specs_Ref_del
---------------------------------------------------------------------------
    AFTER DELETE ON Trn_Head_Specs_Ref
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 