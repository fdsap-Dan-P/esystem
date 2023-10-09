----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Inventory_Specs_Number (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Inventory_Item_ID bigint NOT NULL,
  Specs_Code VarChar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Value numeric(16,6) NOT NULL,
  Value2 numeric(16,6) NOT NULL,
  Measure_ID bigint NULL,

  CONSTRAINT Inventory_Specs_Number_pkey PRIMARY KEY (UUID),
  CONSTRAINT Inventory_Specs_Number_ID FOREIGN KEY (Inventory_Item_ID) REFERENCES Inventory_Item(ID),
  CONSTRAINT fk_Inventory_Specs_Unitmeasure FOREIGN KEY (Measure_ID) REFERENCES Reference(ID),
  CONSTRAINT Inventory_Specs_Number_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxInventory_Specs_Number_Unique ON public.Inventory_Specs_Number(Inventory_Item_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxInventory_Specs_Number_Code ON public.Inventory_Specs_Number(Inventory_Item_ID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgInventory_Specs_Number_Ins on Inventory_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Specs_Number_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Inventory_Specs_Number
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgInventory_Specs_Number_upd on Inventory_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Specs_Number_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Inventory_Specs_Number
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgInventory_Specs_Number_del on Inventory_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Specs_Number_del
---------------------------------------------------------------------------
    AFTER DELETE ON Inventory_Specs_Number
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 
