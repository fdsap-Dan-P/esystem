----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Type_Filter (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Account_Type_ID bigint NOT NULL,
  Central_Office_ID bigint NOT NULL,
  Allow bool NOT NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT Account_Type_Filter_pkey PRIMARY KEY (Account_Type_ID, Central_Office_ID),
  CONSTRAINT fk_Account_Type_Filter_Type FOREIGN KEY (Account_Type_ID) REFERENCES Account_Type(ID),
  CONSTRAINT fk_Account_Type_Filter_Office FOREIGN KEY (Central_Office_ID) REFERENCES Office(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Type_Filter_UUID ON public.Account_Type_Filter(UUID);

DROP TRIGGER IF EXISTS trgAccount_Type_Filter_Ins on Account_Type_Filter;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Type_Filter_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Type_Filter
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAccount_Type_Filter_upd on Account_Type_Filter;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Type_Filter_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Type_Filter
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccount_Type_Filter_del on Account_Type_Filter;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Type_Filter_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Type_Filter
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
