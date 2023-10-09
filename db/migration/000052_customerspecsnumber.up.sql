----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Customer_Specs_Number (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Customer_ID bigint NOT NULL,
  Specs_Code varchar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Value numeric(16,6) NOT NULL,
  Value2 numeric(16,6) NOT NULL,
  Measure_ID bigint NULL,

  CONSTRAINT Customer_Specs_Number_pkey PRIMARY KEY (UUID),
  CONSTRAINT Customer_Specs_Number_ID FOREIGN KEY (Customer_ID) REFERENCES Customer(ID),
  CONSTRAINT fk_Customer_Specs_Unitmeasure FOREIGN KEY (Measure_ID) REFERENCES Reference(ID),
  CONSTRAINT Customer_Specs_Number_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxCustomer_Specs_Number_Unique ON public.Customer_Specs_Number(Customer_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxCustomer_Specs_Number_Code ON public.Customer_Specs_Number(Customer_ID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgCustomer_Specs_Number_Ins on Customer_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_Specs_Number_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Customer_Specs_Number
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgCustomer_Specs_Number_upd on Customer_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_Specs_Number_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Customer_Specs_Number
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgCustomer_Specs_Number_del on Customer_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_Specs_Number_del
---------------------------------------------------------------------------
    AFTER DELETE ON Customer_Specs_Number
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 
