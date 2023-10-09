---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Unit_Conversion (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Type_ID bigint NOT NULL,
  From_ID bigint NOT NULL,
  To_ID bigint NOT NULL,
  Value numeric(16,6) NOT NULL DEFAULT 0,
  Other_Info jsonb NULL,
  
  CONSTRAINT Unit_Conversion_pkey PRIMARY KEY (ID),
  CONSTRAINT idxConversionID  UNIQUE (From_ID, To_ID),
  CONSTRAINT fk_Unit_Conversion_Reference FOREIGN KEY (Type_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Unit_Conversion_from FOREIGN KEY (From_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Unit_Conversion_to FOREIGN KEY (To_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxUnit_Conversion_UUID ON public.Unit_Conversion(UUID);

DROP TRIGGER IF EXISTS trgUnit_Conversion_Ins on Unit_Conversion;
---------------------------------------------------------------------------
CREATE TRIGGER trgUnit_Conversion_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Unit_Conversion
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
  
DROP TRIGGER IF EXISTS trgUnit_Conversion_upd on Unit_Conversion;
---------------------------------------------------------------------------
CREATE TRIGGER trgUnit_Conversion_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Unit_Conversion
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();
    
DROP TRIGGER IF EXISTS trgUnit_Conversion_del on Unit_Conversion;
---------------------------------------------------------------------------
CREATE TRIGGER trgUnit_Conversion_del
---------------------------------------------------------------------------
    AFTER DELETE ON Unit_Conversion
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwUnit_Conversion
----------------------------------------------------------------------------------------
AS SELECT 
    c.ID, mr.UUID,
    typ.ID AS Type_ID,
    typ.Title AS Conversion_Type,
    typ.Short_Name AS Conversion_Type_Short_Name,
    
    base.ID AS From_ID, base.UUID AS ConversionFromUUID,
    base.Title AS Conversion_from,
    base.Short_Name AS Conversion_from_Short_Name,

    
    cur.ID AS To_ID, cur.UUID AS ConversionToUUD,
    cur.Title AS Conversion_to,
    cur.Short_Name AS Conversion_to_Short_Name,
    
    c.Value,
    
    mr.Mod_Ctr,
    c.Other_Info,
    mr.Created,
    mr.Updated
   FROM Unit_Conversion c
     INNER JOIN Main_Record mr on mr.UUID = c.UUID
     JOIN Reference typ ON typ.ID = c.Type_ID
     JOIN Reference base ON base.ID = c.From_ID
     JOIN Reference cur ON cur.ID = c.To_ID;

   INSERT INTO 
     Unit_Conversion(Type_ID, From_ID, To_ID, Value)   
   SELECT y.ID Type_ID, f.ID From_ID, t.ID To_ID, Value
   FROM 
    (Values
      ('Distance', 'Meter', 'Centimeter', 100)
    ) a(Conversion_Type, Conversion_from, Conversion_to, Value), vwReference y, vwReference f, vwReference t
   WHERE 
       y.Ref_Type = 'ConversionType' and y.Title = a.Conversion_Type
   and f.Ref_Type = 'UnitMeasure' and f.Title = a.Conversion_from
   and t.Ref_Type = 'UnitMeasure' and t.Title = a.Conversion_to    
   
   ON CONFLICT(From_ID, To_ID)
   DO UPDATE SET Type_ID = EXCLUDED.Type_ID, Value = EXCLUDED.Value;
   

