
----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Inventory_Detail (
----------------------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Account_Inventory_Id bigint NOT NULL,
  Inventory_Item_Id bigint NOT NULL,
  Repository_Id bigint NULL,
  Supplier_Id bigint NULL,
  Unit_Price numeric(16,6) NOT NULL,
  Book_Value numeric(16,6) NOT NULL,
  Unit numeric(16,6) NOT NULL,
  Measure_ID bigint NOT NULL,
  Batch_Number varchar(30) NULL,
  Date_Manufactured  Date NULL,
  Date_Expired Date NULL,
  Remarks varchar(1000) NOT NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT Inventory_Detail_pkey PRIMARY KEY (ID),
  CONSTRAINT Inventory_Detail_AccInv FOREIGN KEY (Account_Inventory_Id) REFERENCES Account_Inventory(ID),
  CONSTRAINT Inventory_Detail_InvItem FOREIGN KEY (Inventory_Item_Id) REFERENCES Inventory_Item(ID),
  CONSTRAINT Inventory_Detail_Supplier FOREIGN KEY (Supplier_Id) REFERENCES Identity_Info(ID),
  CONSTRAINT Inventory_Detail_Repository FOREIGN KEY (Repository_Id) REFERENCES Inventory_Repository(ID),
  CONSTRAINT fk_Inventory_Detail_Unitmeasure FOREIGN KEY (Measure_ID) REFERENCES Reference(ID)
);

CREATE UNIQUE INDEX IF NOT EXISTS idxInventory_Detail_UUID ON public.Inventory_Detail(UUID);
CREATE INDEX IF NOT EXISTS idxInventory_Detail_InvID ON public.Inventory_Detail(Account_Inventory_Id, Inventory_Item_Id);

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

DROP TRIGGER IF EXISTS trgInventory_Detail_del on Inventory_Detail;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Detail_del
---------------------------------------------------------------------------
    AFTER DELETE ON Inventory_Detail
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

/*
  INSERT into Inventory_Detail(
    UUID, Account_Inventory_Id, Inventory_Item_Id, Unit_Price, 
    Book_Value, Unit, Measure_ID, Batch_Number,
    Date_Manufactured, Date_Expired, Remarks)
  SELECT 
    cast(Acc.UUID as UUID) UUID, a.ID Account_ID, ii.id Inventory_Item_Id,
    Acc.Unit_Price, acc.Book_Value , Acc.Unit, uc.ID Measure_ID, Batch_Number,
    to_date(Date_Manufactured,'yyyy-mm-dd') Date_Manufactured, to_date(Date_Expired,'yyyy-mm-dd') Date_Expired, Acc.Remarks
  FROM (Values
    ('a7fa02d3-5c91-4683-8429-a88d96030c8a', 'c3476afe-bd50-49e6-8de3-074555a8e1bd', '090db518-587c-41a3-9baa-9dc70dae58f8', 25, 110, 2750, 2, 'Box30', 'Batch001', '2021-01-01', '2021-03-31', 'Remarks'),
    ('07ac4c5c-2527-43f7-9739-33399ee7e7a3', 'b35e39e8-885b-41a6-a070-a249c2a099e5', '0df94671-3193-4440-bf0d-ec7f171b294e', 25, 110, 2750, 5, 'Box30', 'Batch002', '2021-01-01', '2021-03-31', 'Remarks')
    )   
    Acc(UUID, Acc_UUID, Item_UUID, Inventory, Unit_Price, Book_Value, Unit, Measure, Batch_Number, Date_Manufactured, Date_Expired, Remarks)
  LEFT JOIN Account_Inventory a on a.UUID = cast(Acc.Acc_UUID  as UUID)
  LEFT JOIN Inventory_Item ii on ii.UUID = cast(Acc.Item_UUID  as UUID)
  LEFT JOIN vwReference uc on uc.Short_Name = Acc.measure and uc.Ref_Type = 'UnitMeasure'
  
  --select * from vwReference v2 where Title = 'Colgate' and Ref_Type = 'Brand_Name '
  ON CONFLICT(UUID) DO UPDATE SET
    Account_Inventory_Id = excluded.Account_Inventory_Id,
    Inventory_Item_Id = excluded.Inventory_Item_Id,
    Unit_Price = excluded.Unit_Price,
    Book_Value = excluded.Book_Value,
    Unit = excluded.Unit, 
    Measure_ID = excluded.Measure_ID, 
    Batch_Number = excluded.Batch_Number, 
    Date_Manufactured = excluded.Date_Manufactured,  
    Date_Expired = excluded.Date_Expired, 
    Remarks = excluded.Remarks;
*/  
  
  