----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Inventory_Specs_String (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Inventory_Item_ID bigint NOT NULL,
  Specs_Code VarChar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Value text NOT NULL,
  
  CONSTRAINT Inventory_Specs_String_pkey PRIMARY KEY (UUID),
  CONSTRAINT Inventory_Specs_String_ID FOREIGN KEY (Inventory_Item_ID) REFERENCES Inventory_Item(ID),
  CONSTRAINT Inventory_Specs_String_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxInventory_Specs_String_Unique ON public.Inventory_Specs_String(Inventory_Item_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxInventory_Specs_String_Code ON public.Inventory_Specs_String(Inventory_Item_ID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgInventory_Specs_String_Ins on Inventory_Specs_String;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Specs_String_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Inventory_Specs_String
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgInventory_Specs_String_upd on Inventory_Specs_String;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Specs_String_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Inventory_Specs_String
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgInventory_Specs_String_del on Inventory_Specs_String;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Specs_String_del
---------------------------------------------------------------------------
    AFTER DELETE ON Inventory_Specs_String
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 