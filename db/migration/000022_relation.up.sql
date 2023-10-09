----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Relation (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  IIID bigint NOT NULL,
  Series int2 NOT NULL,  
  Relation_IIID bigint NOT NULL,
  Type_ID bigint NOT NULL,
  Relation_Date Date NULL,
  Other_Info jsonb NULL,
  CONSTRAINT Relation_pkey PRIMARY KEY (UUID),
  CONSTRAINT idxRelationID  UNIQUE (IIID, Relation_IIID, Type_ID),
  CONSTRAINT fkRelationIdentity_Info FOREIGN KEY (IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fkRelationRelation FOREIGN KEY (Relation_IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fkRelation_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID)
);

DROP TRIGGER IF EXISTS trgRelationIns on Relation;
---------------------------------------------------------------------------
CREATE TRIGGER trgRelationIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Relation
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgRelationupd on Relation;
---------------------------------------------------------------------------
CREATE TRIGGER trgRelationupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Relation
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgRelation_del on Relation;
---------------------------------------------------------------------------
CREATE TRIGGER trgRelation_del
---------------------------------------------------------------------------
    AFTER DELETE ON Relation
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwRelation
----------------------------------------------------------------------------------------
AS SELECT
    mr.UUID,
    p.IIID IIID, b.Series,
    p.Title,    
    p.Last_Name, p.First_Name, p.Middle_Name, p.Mother_Maiden_Name,
    p.Birthday,  p.Sex,
    
    p.Gender_ID, p.Gender,    
    
    p.Civil_Status_ID , p.Civil_Status,
        
    p.Birth_Place_ID, p.Birth_Place, p.Full_Birth_Place, p.Birth_PlaceURL, 
    
    p.Contact_ID, p.Contact_Last_Name, p.Contact_First_Name, p.Contact_Middle_Name,
    
    p.Alternate_ID, p.Identity_Map_ID, p.Simple_Name, 

    p.iiMod_Ctr,
    p.iiOther_Info,
    p.iiCreated,
    p.iiUpdated,

    p.Current_Address_ID,
    p.Current_Detail,
    p.Current_URL,    
    p.Location,
    p.Full_Location,
 
    p.Marriage_Date,
    p.Known_Language,
    p.isAdopted,
  
    p.Source_Income_ID, p.Source_Income,
    p.Disability_ID, p.Disability,
    p.Occupation_ID, p.Occupation,
    p.Sector_ID, p.Sector,
    p.Industry_ID, p.industry,
    p.Religion_ID, p.Religion,
    p.Nationality_ID, p.Nationality,     
   
    relii.ID RelationID, 
    relii.Last_Name RelationLast_Name, relii.First_Name RelationFirst_Name,
    
    typ.ID Type_ID, typ.UUID Relation_TypeUUID, typ.Title Relation_Type,
    
    mr.Mod_Ctr,
    b.Other_Info,
    mr.Created,
    mr.Updated 
   FROM Relation b
   INNER JOIN Main_Record mr on mr.UUID = b.UUID
   LEFT JOIN vwperson p ON p.IIID = b.Relation_IIID
   JOIN Identity_Info relii ON b.IIID = relii.ID
   LEFT JOIN Reference typ ON typ.ID = b.Type_ID;

/*  
   INSERT into Relation(
      IIID, Series, Relation_IIID, Type_ID,   Relation_Date) 
   SELECT 
      ii.ID IIID, a.Series, ir.ID RelationID, typ.ID Type_ID, 
      cast(a.Relation_Date as Date) Relation_Date
   FROM (Values
      ('101','100','Spouse',1,'1997/01/01',NULL)
       )   
    a(Alternate_ID, AlternateRelationid, Relation_Type, Series, Relation_Date, Other_Info)  
      
  LEFT JOIN Identity_Info ii   on ii.Alternate_ID = a.Alternate_ID
  LEFT JOIN Identity_Info ir   on ir.Alternate_ID = a.AlternateRelationid
  LEFT JOIN Relation rel       on rel.IIID = ii.ID
  LEFT JOIN vwReference typ    on lower(typ.Title) = lower(a.Relation_Type) 
     and typ.Ref_Type = 'RelationshipType' 
     
  ON CONFLICT(IIID, Relation_IIID, Type_ID) 
  DO UPDATE SET
    IIID = excluded.IIID,  
    Relation_IIID = excluded.Relation_IIID,   
    Type_ID = excluded.Type_ID,   
    Relation_Date = excluded.Relation_Date;  
*/
 

