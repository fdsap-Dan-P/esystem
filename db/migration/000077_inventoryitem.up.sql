--drop table Inventory_Item cascade
-- update schema_migrations  set dirty = false
----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Inventory_Item (
----------------------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Item_Code varchar(50) NULL,
  Bar_Code varchar(48) NULL,
  Item_Name varchar(100) NOT NULL,
  Unique_Variation varchar(50) NOT NULL,
  Parent_ID bigint NULL,
  Generic_Name_ID bigint NULL,
  Brand_Name_ID bigint NULL,
  Measure_ID bigint NOT NULL, 
  Image_Id bigint NULL,  
  Remarks varchar(1000) NOT NULL,
  Vec_Simple_Name tsvector,
  Other_Info jsonb NULL,  
  
  CONSTRAINT Inventory_Item_pkey PRIMARY KEY (ID),
  CONSTRAINT Inventory_Item_img FOREIGN KEY (Image_ID) REFERENCES Documents(ID),
  CONSTRAINT Inventory_Item_pkg FOREIGN KEY (Parent_ID) REFERENCES Inventory_Item(ID),
  CONSTRAINT fk_Inventory_Item_gen FOREIGN KEY (Generic_Name_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Inventory_Item_bn FOREIGN KEY (Brand_Name_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Inventory_Item_UnitMeasure FOREIGN KEY (Measure_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxInventory_Item_UUID ON public.Inventory_Item(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxInventory_Item_Code ON public.Inventory_Item(Item_Code);
CREATE INDEX IF NOT EXISTS idxInventory_Item_Variation ON public.Inventory_Item(Unique_Variation);

--------------------------------------------------------------------------------------------------
CREATE INDEX IF NOT EXISTS idx_InventoryItem_lexeme 
--------------------------------------------------------------------------------------------------
  ON public.Inventory_Item 
  USING gin (Vec_Simple_Name);
--------------------------------------------------------------------------------------------------

---------------------------------------------------------------------------
-- INSERT InventoryItem Trigger
CREATE or REPLACE FUNCTION InventoryItem_INSERT() returns trigger AS 
---------------------------------------------------------------------------
$$ 
DECLARE 
   ctr int8 := Nextval('Mod_Ctr_seq');
   upd timestamptz := CURRENT_TIMESTAMP;

   SimpleName VarChar(1000);
   VecSimpleName tsvector;
   payload JSON;
BEGIN
   payload = json_build_object( 
      'action', TG_OP,
      'table',  TG_TABLE_NAME::regclass::text,
      'data',  json_build_object( 
          'UUID', NEW.UUID,
          'Mod_Ctr', ctr
      ) 
   );
     
  INSERT INTO Main_Record(UUID, TableName, Mod_Ctr, Created)
  SELECT NEW.UUID, TG_TABLE_NAME::regclass::text, ctr, upd
  ON CONFLICT DO NOTHING;
   
  INSERT INTO Modified(Mod_Ctr, UUID, Updated)  
  SELECT ctr, NEW.UUID, upd;
  
  SELECT 
     setweight(to_tsvector('Simple'::regConfig, unAccent(COALESCE(', '::text || p1.Item_Name::text, ''::text))), 'A'::"char") ||
     setweight(to_tsvector('Simple'::regConfig, unAccent(COALESCE(', '::text || p1.Parent::text, ''::text))), 'B'::"char") ||
     setweight(to_tsvector('Simple'::regConfig, unAccent(COALESCE(', '::text || p1.Brand_Name::text, ''::text))), 'C'::"char") || 
     setweight(to_tsvector('Simple'::regConfig, unAccent(COALESCE(', '::text || p1.Generic_Name::text, ''::text))), 'D'::"char") VecSimpleName
     
  INTO VecSimpleName
  FROM 
   (SELECT 
      r.ID, r.Item_Name,
      COALESCE(p.Item_Name,'') Parent, COALESCE(b.Title,'') Brand_Name, COALESCE(g.Title,'') Generic_Name
      --Simpleword(r.Title::text ||  ', '::text || r.Short_Name || ', '::text ) Simple_Name,
      --Simpleword(COALESCE(', '::text || p.Simple_Name,'')) Parent, Simpleword(y.Title) Ref_Type
    FROM 
      (SELECT new.Id, new.Parent_Id, new.Item_Name, new.Generic_Name_ID, new.Brand_Name_ID) r
    LEFT JOIN Reference g on g.Id = r.Generic_Name_ID
    LEFT JOIN Reference b on b.Id = r.Brand_Name_ID
    LEFT JOIN Inventory_Item p on r.Parent_Id = p.Id
    ) p1;    
   
   
  NEW.Vec_Simple_Name = VecSimpleName;
  PERFORM pg_notify('mychan', payload::text);   
RETURN NEW;
END $$ Language plpgsql;
 

--------------------------------------------------------------------------------------------------
CREATE or REPLACE FUNCTION InventoryItem_UPDATE() RETURNS TRIGGER Language plpgsql AS 
$function$ 
--------------------------------------------------------------------------------------------------
DECLARE 
  ctr int8 := Nextval('Mod_Ctr_seq');
  upd timestamptz := CURRENT_TIMESTAMP;
  SimpleName VarChar(1000);
  VecSimpleName tsvector;
  payload JSON;
BEGIN
  payload = json_build_object( 
    'action', TG_OP,
    'table',  TG_TABLE_NAME::regclass::text,
    'data',  json_build_object( 
    'UUID', NEW.UUID,
    'Mod_Ctr', ctr) 
   );
     
  IF NOT (COALESCE(OLD.Item_Name,'')  =  COALESCE(NEW.Item_Name,'') 
     AND  COALESCE(OLD.Parent_ID,-1) =  COALESCE(NEW.Parent_ID,-1) 
     AND  COALESCE(OLD.Brand_Name_Id,-1) =  COALESCE(NEW.Brand_Name_id,-1) 
     AND  COALESCE(OLD.Generic_Name_Id,-1) =  COALESCE(NEW.Generic_Name_Id,-1) 
     ) THEN 
    -- NEW.Simple_Location := Simpleword(new.Location);
    INSERT INTO Modified(Mod_Ctr, UUID, Updated)
    ValueS(ctr, NEW.UUID, upd);
     
    UPDATE Main_Record 
    SET Updated = Upd
    WHERE UUID = NEW.UUID;
  
    SELECT 
       setweight(to_tsvector('Simple'::regConfig, unAccent(COALESCE(', '::text || p1.Item_Name::text, ''::text))), 'A'::"char") ||
       setweight(to_tsvector('Simple'::regConfig, unAccent(COALESCE(', '::text || p1.Parent::text, ''::text))), 'B'::"char") ||
       setweight(to_tsvector('Simple'::regConfig, unAccent(COALESCE(', '::text || p1.Brand_Name::text, ''::text))), 'C'::"char") || 
       setweight(to_tsvector('Simple'::regConfig, unAccent(COALESCE(', '::text || p1.Generic_Name::text, ''::text))), 'D'::"char") VecSimpleName
       
    INTO VecSimpleName
    FROM 
     (SELECT 
        r.ID, r.Item_Name,
        COALESCE(p.Item_Name,'') Parent, COALESCE(b.Title,'') Brand_Name, COALESCE(g.Title,'') Generic_Name
        --Simpleword(r.Title::text ||  ', '::text || r.Short_Name || ', '::text ) Simple_Name,
        --Simpleword(COALESCE(', '::text || p.Simple_Name,'')) Parent, Simpleword(y.Title) Ref_Type
      FROM 
        (SELECT new.Id, new.Parent_Id, new.Item_Name, new.Generic_Name_ID, new.Brand_Name_ID) r
      LEFT JOIN Reference g on g.Id = r.Generic_Name_ID
      LEFT JOIN Reference b on b.Id = r.Brand_Name_ID
      LEFT JOIN Inventory_Item p on r.Parent_Id = p.Id
      ) p1;
   
    NEW.Vec_Simple_Name = VecSimpleName;

    PERFORM pg_notify('mychan', payload::text);   
  END IF;
 
RETURN NEW;
END; 
$function$;
 
DROP TRIGGER IF EXISTS trgInventoryItem_Ins on Inventory_Item;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventoryItem_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Inventory_Item
    FOR EACH ROW
    EXECUTE PROCEDURE InventoryItem_INSERT();

DROP TRIGGER IF EXISTS trgInventoryItem_upd on Inventory_Item;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventoryItem_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Inventory_Item
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE InventoryItem_UPDATE();

DROP TRIGGER IF EXISTS trgInventory_Item_del on Inventory_Item;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Item_del
---------------------------------------------------------------------------
    AFTER DELETE ON Inventory_Item
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
  
  INSERT into Inventory_Item(
    UUID, Bar_Code, Item_Name, Unique_Variation, Parent_ID, Generic_Name_ID,
    Brand_Name_ID, Measure_ID, Remarks)
  SELECT 
      cast(Acc.UUID as UUID) UUID, Bar_Code, Acc.Item_Name, Unique_Variation,  null Parent_ID, gen.ID Generic_Name_ID, 
      bn.ID Brand_Name_ID, uc.ID Measure_ID, Acc.Remarks
  FROM (Values
      ('090db518-587c-41a3-9baa-9dc70dae58f8', '0705632441947', 'Colgate Ultra', '30 ML', 'Tooth Paste', 'Colgate', 'ML', 'Remarks'),
      ('0df94671-3193-4440-bf0d-ec7f171b294e', '0705632441948', 'Colgate White', '30 ML', 'Tooth Paste', 'Colgate', 'ML', 'Remarks')
      )   
    Acc(
      UUID, Bar_Code, Item_Name, Unique_Variation, Generic_Name, Brand_Name, measure, Remarks
      )
  LEFT JOIN vwReference gen on gen.Title = Acc.Generic_Name  and gen.Ref_Type = 'GenericName'
  LEFT JOIN vwReference bn on bn.Title = Acc.Brand_Name  and bn.Ref_Type = 'BrandName'
  LEFT JOIN vwReference uc on uc.Short_Name = Acc.measure and uc.Ref_Type = 'UnitMeasure'
  
  --select * from vwReference v2 where Title = 'Colgate' and Ref_Type = 'Brand_Name '
  ON CONFLICT(UUID) DO UPDATE SET
    Bar_Code = excluded.Bar_Code,
    Item_Name = excluded.Item_Name,
    Unique_Variation =  excluded.Unique_Variation,
    Parent_ID = excluded.Parent_ID,
    Generic_Name_ID = excluded.Generic_Name_ID,
    Brand_Name_ID = excluded.Brand_Name_ID,
    Measure_ID = excluded.Measure_ID,
    Remarks = excluded.Remarks;
  
