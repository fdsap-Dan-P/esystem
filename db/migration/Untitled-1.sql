----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Inventory_Detail (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Account_Inventory_Id numeric(16,6) NOT NULL,
  Inventory_Item_Id numeric(16,6) NOT NULL,
  ItemCount numeric(16,6) NOT NULL,
  Unit_Price numeric(16,6) NOT NULL, 
  Book_Value numeric(16,6) NOT NULL,
  Unit numeric(16,6) NOT NULL,
  Measure_ID bigint NOT NULL,
  Tax_Rate numeric(10,6) NOT NULL,
  Batch_Number varchar(30) NULL,
  Date_Manufactured  Date NULL,
  Date_Expired Date NULL,
  Remarks varchar(1000) NOT NULL,
  Other_Info jsonb NULL,  
  
  CONSTRAINT Inventory_Detail_pkey PRIMARY KEY (ID),
  CONSTRAINT Inventory_Detail_AccInv FOREIGN KEY (Account_Inventory_Id) REFERENCES Account_Inventory(ID),
  CONSTRAINT fk_Inventory_Detail_Unitmeasure FOREIGN KEY (Measure_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxInventory_Detail_UUID ON public.Inventory_Detail(UUID);

DROP TRIGGER IF EXISTS trgInventory_Detail_Ins on Inventory_Detail;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Detail_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Inventory_Detail
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgInventory_Detail_upd on Inventory_Detail;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Detail_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Inventory_Detail
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();
