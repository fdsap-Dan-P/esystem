----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Customer_Specs_Ref (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Customer_ID bigint NOT NULL,
  Specs_Code varchar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Ref_ID bigint NOT NULL,
  
  CONSTRAINT Customer_Specs_Ref_pkey PRIMARY KEY (UUID),
  CONSTRAINT Customer_Specs_Ref_ID FOREIGN KEY (Customer_ID) REFERENCES Customer(ID),
  CONSTRAINT Customer_Specs_Ref_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID),
  CONSTRAINT Customer_Specs_Ref_Ref FOREIGN KEY (Ref_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxCustomer_Specs_Ref_Unique ON public.Customer_Specs_Ref(Customer_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxCustomer_Specs_Ref_Code ON public.Customer_Specs_Ref(Customer_ID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgCustomer_Specs_Ref_Ins on Customer_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_Specs_Ref_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Customer_Specs_Ref
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgCustomer_Specs_Ref_upd on Customer_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_Specs_Ref_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Customer_Specs_Ref
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgCustomer_Specs_Ref_del on Customer_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_Specs_Ref_del
---------------------------------------------------------------------------
    AFTER DELETE ON Customer_Specs_Ref
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 