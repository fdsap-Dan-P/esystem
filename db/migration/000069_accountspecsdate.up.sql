----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Specs_Date (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Account_ID bigint NOT NULL,
  Specs_Code varchar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Value Date NOT NULL,
  Value2 Date NOT NULL,
  
  CONSTRAINT Account_Specs_Date_pkey PRIMARY KEY (UUID),
  CONSTRAINT Account_Specs_Date_ID FOREIGN KEY (Account_ID) REFERENCES Account(ID),
  CONSTRAINT Account_Specs_Date_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Specs_Date_Unique ON public.Account_Specs_Date(Account_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Specs_Date_Code ON public.Account_Specs_Date(Account_ID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgAccount_Specs_Date_Ins on Account_Specs_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Specs_Date_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Specs_Date
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgAccount_Specs_Date_upd on Account_Specs_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Specs_Date_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Specs_Date
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgAccount_Specs_Date_del on Account_Specs_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Specs_Date_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Specs_Date
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 
  