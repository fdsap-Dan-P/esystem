----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Inventory (
----------------------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Account_ID bigint NOT NULL,
  Repository_ID bigint NULL,
  Bar_Code varchar(48) NULL, 
  Code varchar(50) NOT NULL,
  Quantity numeric(16,6) NOT NULL,
  Unit_Price numeric(16,6) NOT NULL, 
  Book_Value numeric(16,6) NOT NULL,
  Discount numeric(10,6) NOT NULL, 
  Tax_Rate numeric(10,6) NOT NULL,
  Remarks varchar(1000) NOT NULL,
  Other_Info jsonb NULL,  
  
  CONSTRAINT Account_Inventory_pkey PRIMARY KEY (ID),
  CONSTRAINT Account_Inventory_ID FOREIGN KEY (Account_ID) REFERENCES Account(ID),
  CONSTRAINT Account_Inventory_Repository FOREIGN KEY (Repository_ID) REFERENCES Inventory_Repository(ID)
  -- CONSTRAINT Account_Inventory_pkg FOREIGN KEY (Package_ID) REFERENCES Account_Inventory(ID),
  -- CONSTRAINT fk_Account_Inventory_Unitmeasure FOREIGN KEY (Measure_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Inventory_UUID ON public.Account_Inventory(UUID);

DROP TRIGGER IF EXISTS trgAccount_Inventory_Ins on Account_Inventory;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Inventory_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Inventory
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAccount_Inventory_upd on Account_Inventory;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Inventory_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Inventory
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccount_Inventory_del on Account_Inventory;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Inventory_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Inventory
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
/* 
  INSERT into Account_Inventory(
      UUID, Account_ID, Bar_Code, Code, Quantity, Unit_Price, Book_Value, Discount, Tax_Rate, Remarks)
  SELECT 
      cast(Acc.UUID as UUID), a.ID Account_ID, Bar_Code, Code,
      Acc.Quantity, Acc.Unit_Price, Book_Value, Acc.Discount, Tax_Rate, Acc.Remarks
  FROM (Values
      ('c3476afe-bd50-49e6-8de3-074555a8e1bd', '1001-0001-0000001', null, 'Code', 25, 110, 2750, 0, .12, 'Test Product'),
      ('b35e39e8-885b-41a6-a070-a249c2a099e5', '1001-0001-0000001', null, 'Code', 25, 110, 2750, 0, .12, 'Test Product')
    )   
    Acc(
      UUID, Alternate_Acc, Bar_Code, Code, Quantity, Unit_Price, Book_Value, Discount, Tax_Rate, Remarks
      )
  LEFT JOIN Account a on a.Alternate_Acc = Acc.Alternate_Acc 
  
  --select * from vwReference v2 where Title = 'Colgate' and Ref_Type = 'Brand_Name '
  ON CONFLICT(UUID) DO UPDATE SET
    Account_ID = excluded.Account_ID,
    Bar_Code = excluded.Bar_Code,
    Code = excluded.Code,
    Quantity = excluded.Quantity,
    Unit_Price  = excluded.Unit_Price ,
    Book_Value = excluded.Book_Value,
    Discount = excluded.Discount,
    Tax_Rate = excluded.Tax_Rate,
    Remarks = excluded.Remarks
    ;
  
 */