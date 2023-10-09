--- drop table Account_Beneficiary cascade
----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Beneficiary (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Account_ID bigint NOT NULL,
  IIID bigint NOT NULL,
  Series int2 NOT NULL,
  Beneficiary_Type_ID bigint NOT NULL,
  Relationship_Type_ID bigint NOT NULL,
  Other_Info jsonb NULL,

  CONSTRAINT Account_Beneficiary_pkey PRIMARY KEY (Account_ID, IIID),
  CONSTRAINT fk_Account_Beneficiary_Acc FOREIGN KEY (Account_ID) REFERENCES Account(ID),
  CONSTRAINT fk_Account_Beneficiary_IIID FOREIGN KEY (IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fk_Account_Beneficiary_Type FOREIGN KEY (Beneficiary_Type_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Account_Beneficiary_Rel FOREIGN KEY (Relationship_Type_ID) REFERENCES Reference(ID)
);

CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Beneficiary_UUID ON public.Account_Beneficiary(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS fk_Account_Beneficiary_Account 
  ON public.Account_Beneficiary (Account_ID, IIID);

DROP TRIGGER IF EXISTS trgAccount_Beneficiary_Ins on Account_Beneficiary;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Beneficiary_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Beneficiary
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();

DROP TRIGGER IF EXISTS trgAccount_Beneficiary_upd on Account_Beneficiary;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Beneficiary_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Beneficiary
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccount_Beneficiary_del on Account_Beneficiary;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Beneficiary_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Beneficiary
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwAccount_Beneficiary
----------------------------------------------------------------------------------------
AS SELECT 

    mr.UUID,
    p.IIID, 
    p.Title,    
    p.Last_Name, p.First_Name, p.Middle_Name, p.Mother_Maiden_Name,
    p.Birthday, p.Sex,
    
    p.Gender_ID, p.Gender,    
    
    p.Civil_Status_ID , p.Civil_Status,
        
    p.Birth_Place_ID, 
    p.Birth_Place, p.Full_Birth_Place,
    p.Birth_PlaceURL, 
    
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
    p.isadopted,
  
    p.Source_Income_ID, p.Source_Income,
    p.Disability_ID, p.Disability,
    p.Occupation_ID, p.Occupation,
    p.Sector_ID, p.Sector,
    p.Industry_ID, p.industry,
    p.Religion_ID, p.Religion,
    p.Nationality_ID, p.Nationality,         
        
    c.ID Customer_ID, c.UUID CustomerUUID, 
    a.ID AccoountID, a.UUID AccoountUUID,
    
    bentyp.ID BeneficiaryType_ID, bentyp.UUID BeneficiaryTypeUUID, bentyp.Title BeneficiaryTypeTitle,
    
    reltyp.ID Relation_Type_ID, reltyp.UUID Relation_TypeUUID, reltyp.Title Relation_TypeTitle,
 
    mr.Mod_Ctr,
    b.Other_Info,
    mr.Created,
    mr.Updated 
    
   FROM Account_Beneficiary b
   INNER JOIN Main_Record mr on mr.UUID = b.UUID  
   JOIN vwperson p ON b.IIID = p.IIID
   LEFT JOIN Account a ON a.ID = b.Account_ID
   LEFT JOIN Customer c ON c.ID = a.Customer_ID
   LEFT JOIN Reference bentyp ON bentyp.ID = b.Beneficiary_Type_ID
   LEFT JOIN Reference reltyp ON reltyp.ID = b.Relationship_Type_ID;

/*
  INSERT INTO Account_Beneficiary(
    Account_ID, Series, Beneficiary_Type_ID, 
    IIID, Relationship_Type_ID)
  SELECT 
    Acc.ID Account_ID, Series, benetyp.ID Type_ID, 
    ii.ID IIID, relType.ID Relation_Type_ID
    
   FROM (Values
      ('1001-0001-0000002', 1, 'Irrevocable','101', 'Spouse')
      )   
    a(Alternate_Acc, Series, Account_BeneficiaryType, Alternate_ID, Relation_Type)  

  LEFT JOIN Account  Acc        on Acc.Alternate_Acc = a.Alternate_Acc
  LEFT JOIN vwReference benetyp on lower(benetyp.Title)      = lower(a.Account_BeneficiaryType) 
    and lower(benetyp.Ref_Type) = 'beneficiarytype'
  LEFT JOIN vwReference relType on lower(relType.Title)      = lower(a.Relation_Type)    
    and lower(relType.Ref_Type) = 'relationshiptype'
  LEFT JOIN Identity_Info ii    on ii.Alternate_ID = a.Alternate_ID
 
  ON CONFLICT(Account_ID, IIID)
  DO UPDATE SET
    Account_ID = excluded.Account_ID,
    IIID = excluded.IIID,
    Beneficiary_Type_ID = excluded.Beneficiary_Type_ID,
    Relationship_Type_ID = excluded.Relationship_Type_ID
  ;  


-- select * from vwReference v2  where Ref_Type = 'RelationshipType'
select * from vwAccount_Beneficiary;
--select * from vwperson v order by first_Name
*/