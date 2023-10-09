----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Specs_Ref (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Account_ID bigint NOT NULL,
  Specs_Code varchar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Ref_Id bigint NOT NULL,
  
  CONSTRAINT Account_Specs_Ref_pkey PRIMARY KEY (UUID),
  CONSTRAINT Account_Specs_Ref_ID FOREIGN KEY (Account_ID) REFERENCES Account(ID),
  CONSTRAINT Account_Specs_Ref_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID),
  CONSTRAINT Account_Specs_Ref_Ref FOREIGN KEY (Ref_Id) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Specs_Ref_Unique ON public.Account_Specs_Ref(Account_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Specs_Ref_Code ON public.Account_Specs_Ref(Account_ID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgAccount_Specs_Ref_Ins on Account_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Specs_Ref_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Specs_Ref
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgAccount_Specs_Ref_upd on Account_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Specs_Ref_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Specs_Ref
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgAccount_Specs_Ref_del on Account_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Specs_Ref_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Specs_Ref
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 