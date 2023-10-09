----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Customer_Specs_Date (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Customer_ID bigint NOT NULL,
  Specs_Code varchar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Value Date NOT NULL,
  Value2 Date NOT NULL,
  
  CONSTRAINT Customer_Specs_Date_pkey PRIMARY KEY (UUID),
  CONSTRAINT Customer_Specs_Date_ID FOREIGN KEY (Customer_ID) REFERENCES Customer(ID),
  CONSTRAINT Customer_Specs_Date_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxCustomer_Specs_Date_Unique ON public.Customer_Specs_Date(Customer_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxCustomer_Specs_Date_Code ON public.Customer_Specs_Date(Customer_ID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgCustomer_Specs_Date_Ins on Customer_Specs_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_Specs_Date_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Customer_Specs_Date
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgCustomer_Specs_Date_upd on Customer_Specs_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_Specs_Date_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Customer_Specs_Date
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgCustomer_Specs_Date_del on Customer_Specs_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_Specs_Date_del
---------------------------------------------------------------------------
    AFTER DELETE ON Customer_Specs_Date
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 
  