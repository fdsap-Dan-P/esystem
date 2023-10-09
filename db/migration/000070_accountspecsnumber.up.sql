----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Specs_Number (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Account_ID bigint NOT NULL,
  Specs_Code varchar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Value numeric(16,6) NOT NULL,
  Value2 numeric(16,6) NOT NULL,
  Measure_ID bigint NULL,

  CONSTRAINT Account_Specs_Number_pkey PRIMARY KEY (UUID),
  CONSTRAINT Account_Specs_Number_ID FOREIGN KEY (Account_ID) REFERENCES Account(ID),
  CONSTRAINT fk_Account_Specs_Unitmeasure FOREIGN KEY (Measure_ID) REFERENCES Reference(ID),
  CONSTRAINT Account_Specs_Number_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Specs_Number_Unique ON public.Account_Specs_Number(Account_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Specs_Number_Code ON public.Account_Specs_Number(Account_ID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgAccount_Specs_Number_Ins on Account_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Specs_Number_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Specs_Number
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgAccount_Specs_Number_upd on Account_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Specs_Number_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Specs_Number
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgAccount_Specs_Number_del on Account_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Specs_Number_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Specs_Number
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 
