----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Type_Group (
----------------------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Product_ID bigint NOT NULL,
  Group_ID bigint NOT NULL,
  Account_Type_Group varchar(255) NOT NULL,
--   IIID bigint NULL,
  Normal_Balance bool NOT NULL,
  Isgl bool NOT NULL,
  Active bool NOT NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT Account_Type_Group_pkey PRIMARY KEY (ID),
  CONSTRAINT idxAccount_Type_Group_Code UNIQUE (Product_ID, Group_ID),
  CONSTRAINT idxAccount_Type_Group_Name UNIQUE (Account_Type_Group),
  CONSTRAINT fk_Account_Type_Group_Product FOREIGN KEY (Product_ID) REFERENCES Product(ID),
  CONSTRAINT fk_Account_Type_Group_Group FOREIGN KEY (Group_ID) REFERENCES Reference(ID)
--   CONSTRAINT fk_Account_Type_Group_IIID FOREIGN KEY (IIID) REFERENCES Identity_Info(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Type_Group_UUID ON public.Account_Type_Group(UUID);

DROP TRIGGER IF EXISTS trgAccount_Type_Group_Ins on Account_Type_Group;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Type_Group_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Type_Group
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAccount_Type_Group_upd on Account_Type_Group;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Type_Group_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Type_Group
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccount_Type_Group_del on Account_Type_Group;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Type_Group_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Type_Group
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

