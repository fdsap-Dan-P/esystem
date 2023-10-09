----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Customer_Beneficiary (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Customer_ID bigint NOT NULL,
  Series int2 NOT NULL,  
  IIID bigint NOT NULL,
  Type_ID bigint NOT NULL,
  Relation_Type_ID bigint NOT NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT Customer_Beneficiary_pkey PRIMARY KEY (UUID),
  CONSTRAINT idxCustomer_Beneficiary_ID  UNIQUE (Customer_ID, IIID),
  CONSTRAINT fkCustomer_Beneficiary_IIID FOREIGN KEY (IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fkCustomer_Beneficiary_Cust FOREIGN KEY (Customer_ID) REFERENCES Customer(ID),
  CONSTRAINT fkCustomer_Beneficiary_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID),
  CONSTRAINT fkCustomer_Beneficiary_Rel FOREIGN KEY (Relation_Type_ID) REFERENCES Reference(ID)
);

DROP TRIGGER IF EXISTS trgCustomer_BeneficiaryIns on Customer_Beneficiary;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_BeneficiaryIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Customer_Beneficiary
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgCustomer_Beneficiaryupd on Customer_Beneficiary;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_Beneficiaryupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Customer_Beneficiary
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgCustomer_Beneficiary_del on Customer_Beneficiary;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_Beneficiary_del
---------------------------------------------------------------------------
    AFTER DELETE ON Customer_Beneficiary
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwCustomer_Beneficiary
----------------------------------------------------------------------------------------
AS SELECT

    mr.UUID,
    b.Customer_ID,   
    b.Series,
    typ.ID Type_ID, typ.UUID TypeUUID, typ.Title BeneficiaryType,
    reltyp.ID Relation_Type_ID, reltyp.UUID Relation_TypeUUID, reltyp.Title Relation_Type,

    p.Title,    
    p.Last_Name, p.First_Name, p.Middle_Name, p.Mother_Maiden_Name,
    p.Birthday,  p.Sex,
    
    p.Gender_ID, p.Gender,    
    
    p.Civil_Status_ID , p.Civil_Status,
        
    p.Birth_Place_ID, p.Birth_Place, p.Full_Birth_Place, p.Birth_PlaceURL, 
    
    p.Contact_ID, p.Contact_Last_Name, p.Contact_First_Name, p.Contact_Middle_Name,
    
    p.Alternate_ID, p.Phone, p.Email, p.Identity_Map_ID, p.Simple_Name, 

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
      
    mr.Mod_Ctr,
    b.Other_Info,
    mr.Created,
    mr.Updated 
   FROM Customer_Beneficiary b
   INNER JOIN Main_Record mr on mr.UUID = b.UUID
   INNER JOIN Customer cust on cust.ID = b.Customer_ID
   LEFT JOIN vwperson p ON p.IIID = b.IIID
   LEFT JOIN Reference typ ON typ.ID = b.Type_ID
   LEFT JOIN Reference Reltyp ON Reltyp.ID = b.Relation_Type_ID;
  
 ----------------------------------------------------------------------------------------
  INSERT into Customer_Beneficiary(
      Customer_ID, Series, IIID, Type_ID, Relation_Type_ID, Other_Info) 
   SELECT 
      c.ID Customer_ID, a.Series, ii.ID IIID, typ.ID Type_ID, relTyp.ID RelType_ID, NULL Other_Info
   FROM (Values
      ('10001', 1, '101', 'Irrevocable', 'Contact Person', NULL)
       )   
    a(Customer_Alt_ID, Series, Alternate_ID, Type_ID, Relation_Type, Other_Info)  

  INNER JOIN Customer c on c.Customer_Alt_ID = a.Customer_Alt_ID
  LEFT JOIN Identity_Info ii   on ii.Alternate_ID = a.Alternate_ID
  LEFT JOIN Customer_Beneficiary rel on rel.IIID = ii.ID
  LEFT JOIN vwReference typ    on lower(typ.Title) = lower(a.Type_ID) 
     and lower(typ.Ref_Type) = 'beneficiarytype' 
  LEFT JOIN vwReference reltyp    on lower(reltyp.Title) = lower(a.Relation_Type) 
     and lower(reltyp.Ref_Type) = 'relationshiptype' 
     
  ON CONFLICT(Customer_ID, IIID) DO UPDATE SET
    Customer_ID = excluded.Customer_ID,  
    Series = excluded.Series,   
    IIID = excluded.IIID,   
    Type_ID = excluded.Type_ID,
    Other_Info = excluded.Other_Info ;  
