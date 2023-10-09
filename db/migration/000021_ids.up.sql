-- drop table IDs cascade
----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.IDs (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  IIID bigint NOT NULL,
  Series int2 NOT NULL,
  ID_Number varchar(30) NOT NULL,
  Registration_Date Date NULL,
  Validity_Date Date NULL,
  Type_ID bigint NOT NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT IDs_pkey PRIMARY KEY (IIID, Series),
  CONSTRAINT fk_IDs_Identity_Info FOREIGN KEY (IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fk_IDs_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxIDs_UUID ON public.IDs(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxIDs_IDs ON public.IDs(IIID, Type_ID, lower(trim(ID_Number)));

DROP TRIGGER IF EXISTS trgIDs_Ins on IDs;
---------------------------------------------------------------------------
CREATE TRIGGER trgIDs_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON IDs
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgIDs_upd on IDs;
---------------------------------------------------------------------------
CREATE TRIGGER trgIDs_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON IDs
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgIDs_del on IDs;
---------------------------------------------------------------------------
CREATE TRIGGER trgIDs_del
---------------------------------------------------------------------------
    AFTER DELETE ON IDs
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwIDs
----------------------------------------------------------------------------------------
AS SELECT 
     mr.UUID,

     ii.ID IIID, ii.Alternate_ID,
     ii.Last_Name, ii.First_Name, ii.Middle_Name,

     c.Series,
    
     typ.ID Type_ID, typ.UUID IDTypeUUID, typ.Title AS IDType,
    
     c.ID_Number,
     c.Registration_Date,
     c.Validity_Date,
    
     mr.Mod_Ctr,
     c.Other_Info,
     mr.Created,
     mr.Updated 
   FROM IDs c
   JOIN Main_Record mr ON mr.UUID = c.UUID
   JOIN Identity_Info ii ON c.IIID = ii.ID
   LEFT JOIN Reference typ ON typ.ID = c.Type_ID;

/*
    INSERT into IDs( 
      IIID,   Series,   Type_ID,   ID_Number,   Registration_Date,   Validity_Date) 
    SELECT 
      ii.ID, a.Series, typ.ID Type_ID, 
      a.ID_Number, cast(a.Registration_Date as Date), cast(a.Validity_Date as Date)
    FROM (Values
      ('100',1,'PRC','12313 123','2019/01/01','2022/01/01',NULL),
      ('101',2,'PRC','432213 c3','2017/01/01','2020/01/01',NULL)
         )   
      a(Alternate_ID, Series, id_Type, ID_Number, Registration_Date, Validity_Date, Other_Info)  
      
    LEFT JOIN Identity_Info ii   on ii.Alternate_ID = a.Alternate_ID
    LEFT JOIN IDs c              on c.IIID = ii.ID and c.Series = a.Series
    LEFT JOIN vwReference typ    on lower(typ.Title) = lower(a.id_Type) and typ.Ref_Type = 'IDType'  
    ON CONFLICT (IIID, Series) DO UPDATE SET 
      IIID = excluded.IIID,   
      Series = excluded.Series,   
      Type_ID = excluded.Type_ID,   
      ID_Number = excluded.ID_Number,   
      Registration_Date = excluded.Registration_Date,   
      Validity_Date = excluded.Validity_Date
  ; 
*/