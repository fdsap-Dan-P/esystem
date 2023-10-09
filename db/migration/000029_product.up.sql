---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Product (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Code SmallInt NOT NULL,
  Product_Name varchar(50) NOT NULL,
  Description varchar(500) NULL,
  Normal_Balance bool NOT NULL,
  IsGl bool NOT NULL,
  Other_Info jsonb NULL,
  CONSTRAINT Product_pkey PRIMARY KEY (ID)
  );
  
CREATE UNIQUE INDEX IF NOT EXISTS idxProduct_UUID ON public.Product(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxProduct_Code ON public.Product(Code);
CREATE UNIQUE INDEX IF NOT EXISTS idxProduct_Title ON public.Product(lower(Product_Name));

DROP TRIGGER IF EXISTS trgProduct_Ins on Product;
---------------------------------------------------------------------------
CREATE TRIGGER trgProduct_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Product
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgProduct_upd on Product;
---------------------------------------------------------------------------
CREATE TRIGGER trgProduct_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Product
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgProduct_del on Product;
---------------------------------------------------------------------------
CREATE TRIGGER trgProduct_del
---------------------------------------------------------------------------
    AFTER DELETE ON Product
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

  INSERT INTO Product(Code, Product_Name, Normal_Balance, isgl)
  SELECT 
    Code, Product_Name, Normal_Balance, isgl   
    FROM (Values
      (0,'Accounting',TRUE, TRUE),
      (100,'Cash',TRUE, FALSE),
      (115,'Bank Deposit',TRUE, FALSE),
      (110,'Investment',TRUE, FALSE),
      (120,'Inventory',TRUE, FALSE),
      (130,'Loan',TRUE, FALSE),
      (140,'Receivable',TRUE, FALSE),
      (150,'Fixed Asset',TRUE, FALSE),
      (160,'Property Investment',FALSE, FALSE),
      (210,'Payable',FALSE, FALSE),
      (211,'Collecting Facility',FALSE, FALSE),
      (212,'Remittance',FALSE, FALSE),
      (230,'Savings',FALSE, FALSE),
      (250,'Insurance',FALSE, FALSE),
      (260,'Loans Payable',FALSE, FALSE),
      (310,'Capital Stocks',FALSE, FALSE),
      (510,'Service',FALSE, FALSE),
      (520,'Rental',FALSE, FALSE),
      (550,'Sales',FALSE, FALSE),
      (530,'Interest Income',FALSE, FALSE),
      (580,'Donation',FALSE, FALSE),
      (610,'Interest Expenses',TRUE, FALSE),
      (620,'Reimbursement',TRUE, FALSE),
      (630,'Payroll',TRUE, FALSE),
      (710,'School',TRUE, FALSE)
      
      )   
     a(Code, Product_Name, Normal_Balance, isgl)
   ON CONFLICT(Code)
   DO UPDATE SET Product_Name = EXCLUDED.Product_Name, Normal_Balance=EXCLUDED.Normal_Balance, isgl = EXCLUDED.isgl;
  ;

