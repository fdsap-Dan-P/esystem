----------------------------------------------------------------------------------------
 CREATE TABLE IF NOT EXISTS public.Personal_Info (
----------------------------------------------------------------------------------------
  ID bigint NOT NULL,  
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Maiden_LName varchar(50) NULL,
  Maiden_FName varchar(50) NULL,
  Maiden_MName varchar(50) NULL,
  Spouse_FName VarChar(50) NULL,
  Spouse_LName VarChar(50) NULL,
  Spouse_MName VarChar(50) NULL,
  Mother_Maiden_LName varchar(50) NULL,
  Mother_Maiden_FName varchar(50) NULL,
  Mother_Maiden_MName varchar(50) NULL,
  Marriage_Date Date NULL,
  isAdopted bool NULL,
  Known_Language varchar(1000) NULL,
  Industry_ID bigint NULL,
  Nationality_ID bigint NULL,
  Occupation_ID bigint NULL,
  Education_ID bigint NULL,
  Religion_ID bigint NULL,
  Sector_ID bigint NULL,
  Source_Income_ID bigint NULL,
  Disability_ID bigint NULL,
  Business_Name VarChar(100) NULL,
  Business_Address VarChar(200) NULL,
  Business_Address_Id bigint NULL,
  Business_Position VarChar(50) NULL,
  Other_Info jsonb NULL,

  CONSTRAINT Personal_Info_pkey PRIMARY KEY (ID),
  CONSTRAINT fkpersonalindustry FOREIGN KEY (Industry_ID) REFERENCES Reference(ID),
  CONSTRAINT fkpersonalNationality FOREIGN KEY (Nationality_ID) REFERENCES Reference(ID),
  CONSTRAINT fkpersonalOccupation FOREIGN KEY (Occupation_ID) REFERENCES Reference(ID),
  CONSTRAINT fkpersonalEducation FOREIGN KEY (Education_ID) REFERENCES Reference(ID),
  CONSTRAINT fkpersonalReligion FOREIGN KEY (Religion_ID) REFERENCES Reference(ID),
  CONSTRAINT fkpersonalSector FOREIGN KEY (Sector_ID) REFERENCES Reference(ID),
  CONSTRAINT fkpersonalSourceofincome FOREIGN KEY (Source_Income_ID) REFERENCES Reference(ID),
  CONSTRAINT fkpersonalDisability FOREIGN KEY (Disability_ID) REFERENCES Reference(ID),
  CONSTRAINT fkpersonalBusAddress FOREIGN KEY (Business_Address_Id) REFERENCES Geography(ID),
  CONSTRAINT fkpersonalIdentity FOREIGN KEY (ID) REFERENCES Identity_Info(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxPersonal_Info_UUID ON public.Personal_Info(UUID);

DROP TRIGGER IF EXISTS trgPersonal_InfoIns on Personal_Info;
---------------------------------------------------------------------------
CREATE TRIGGER trgPersonal_InfoIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Personal_Info
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();

DROP TRIGGER IF EXISTS trgPersonal_Infoupd on Personal_Info;
---------------------------------------------------------------------------
CREATE TRIGGER trgPersonal_Infoupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Personal_Info
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgPersonal_Info_del on Personal_Info;
---------------------------------------------------------------------------
CREATE TRIGGER trgPersonal_Info_del
---------------------------------------------------------------------------
    AFTER DELETE ON Personal_Info
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwperson
----------------------------------------------------------------------------------------
AS SELECT 
    ii.ID IIID,  mr.UUID,
    ii.Title,    
    ii.Last_Name, ii.First_Name, ii.Middle_Name, ii.Mother_Maiden_Name,
    ii.Birthday, ii.Sex,
    
    ii.Gender_ID, ii.Gender,    
    
    ii.Civil_Status_ID , ii.Civil_Status,
        
    ii.Birth_Place_ID, ii.Birth_Place, ii.Full_Birth_Place, ii.Birth_PlaceURL, 
    
    ii.Contact_ID, ii.Contact_Last_Name, ii.Contact_First_Name, ii.Contact_Middle_Name,
    
    ii.Alternate_ID, ii.Phone, ii.Email, ii.Identity_Map_ID, ii.Simple_Name, 
    
    ii.Mod_Ctr iiMod_Ctr,
    ii.Other_Info iiOther_Info,
    ii.Created iiCreated,
    ii.Updated iiUpdated,
        
    adrs.UUID Address_UUID,
    adrs.Geography_ID Current_Address_ID,
    adrs.Detail Current_Detail,
    adrs.URL Current_URL,    
    vgeo.Location,
    vgeo.Full_Location,
 
    p.Marriage_Date,
    p.Known_Language,
    p.isAdopted,
  
    inc.ID Source_Income_ID, inc.UUID Source_IncomeUUID, inc.Title Source_Income,
    dis.ID Disability_ID, dis.UUID DisabilityUUID, dis.Title Disability,
    occ.ID Occupation_ID, occ.UUID OccupationUUID, occ.Title Occupation,
    sec.ID Sector_ID, sec.UUID SectorUUID, sec.Title Sector,
    ind.ID Industry_ID, ind.UUID IndustryUUID, ind.Title industry,
    rel.ID Religion_ID, rel.UUID ReligionUUID, rel.Title Religion,
    nat.ID Nationality_ID, nat.UUID NationalityUUID, nat.Title Nationality,     
   
    mr.Mod_Ctr,
    p.Other_Info,
    mr.Created,
    mr.Updated 
  FROM vwIdentity_Info ii
  LEFT JOIN Main_Record mr on mr.UUID = ii.UUID
  LEFT JOIN Personal_Info p ON p.ID = ii.ID  
  LEFT JOIN Address_List adrs ON adrs.IIID = ii.ID AND adrs.Series = 1
  LEFT JOIN vwGeography vgeo ON vgeo.ID = adrs.Geography_ID
  LEFT JOIN Reference inc ON inc.ID = p.Source_Income_ID
  LEFT JOIN Reference dis ON dis.ID = p.Disability_ID
  LEFT JOIN Reference occ ON occ.ID = p.Occupation_ID
  LEFT JOIN Reference sec ON sec.ID = p.Sector_ID
  LEFT JOIN Reference ind ON ind.ID = p.Industry_ID
  LEFT JOIN Reference rel ON rel.ID = p.Religion_ID
  LEFT JOIN Reference nat ON nat.ID = p.Nationality_ID;

--------------------------------------------------------------------------------------------
DO $$
--------------------------------------------------------------------------------------------
DECLARE 
  Titles TEXT DEFAULT '';
  rec   RECORD;
  cur CURSOR FOR 

  SELECT 
    ii.IIID,
    a.Last_Name, a.First_Name, a.Middle_Name, a.Mother_Maiden_Name, 
    cast(a.Birthday as Date) Birthday, 
    Title.Title, Civil.ID Civil_Status_ID , a.Birth_Place, 
    a.Sex, Gender.ID Gender_ID, 
    a.Current_Address, a.Current_Address_Detail, a.Current_Address_URL,
    Contact.ID Contact_ID, income.ID Source_Income_ID, 
    Disability.ID Disability_ID, a.isAdopted, Occupation.ID Occupation_ID, 
    a.Sector, industry.ID Industry_ID, Religion.ID Religion_ID, 
    Nationality.ID Nationality_ID, cast(a.Marriage_Date as Date) Marriage_Date, a.Known_Language, 
    a.Alternate_ID, cast(a.data as jsonb) data1
    
 FROM (Values
      ('Application','Roderick','','','2023-04-14',null,'Single','San Pablo City',TRUE,'Male','Soledad San Pablo City, Laguna','Purok 3',null,null,'Apps','',null,FALSE,'',null,'','','Filipino',null,null,null),
      ('Mercado','Roderick','de Guzman','Florcita Dia de Guzman','1979/04/14',null,'Married','San Pablo City',TRUE,'Male','Soledad San Pablo City, Laguna','Purok 3',null,null,'100','Salary',null,FALSE,'Computer Programmer ',null,'Financial services; professional services','United Church of Christ in the Philippines','Filipino','1997/09/21','Tagalog; English',null),
      ('Mercado','Olive','Maluping','Amparo Lopez','1976/06/09',null,'Married','San Pablo City',FALSE,'Female','Soledad San Pablo City, Laguna','Purok 3',null,null,'101','Salary',null,FALSE,'Elementary School Teachers, Except Special Education',null,'Education ','United Church of Christ in the Philippines','Filipino','1997/09/21','Tagalog; English',null)
      )   
    a(Last_Name, First_Name, Middle_Name, Mother_Maiden_Name, Birthday, 
      Title, Civil_Status, Birth_Place, Sex, Gender, 
      Current_Address, Current_Address_Detail, Current_Address_URL, ContactaltID, 
      Alternate_ID, Source_Income, Disability, isAdopted, Occupation, Sector, 
      industry, Religion, Nationality, Marriage_Date, Known_Language, data)  
      
  LEFT JOIN vwperson ii             on ii.Alternate_ID = a.Alternate_ID
  LEFT JOIN vwReference Title       on lower(Title.Title) = lower(a.Title) and Title.Ref_Type = 'Title'  
  LEFT JOIN vwReference Civil       on lower(Civil.Title) = lower(a.Civil_Status)          and Civil.Ref_Type = 'Civil_Status'  
  LEFT JOIN vwReference Gender      on lower(Gender.Title)        = lower(a.Gender)        and Gender.Ref_Type = 'Gender'
  LEFT JOIN vwReference Disability  on lower(Disability.Title)    = lower(a.Disability)    and Disability.Ref_Type = 'Disabilities'
  LEFT JOIN vwReference income      on lower(income.Title)        = lower(a.Source_Income) and income.Ref_Type = 'SourceofIncome'
  LEFT JOIN vwReference Occupation  on lower(Occupation.Title)    = lower(a.Occupation)    and Occupation.Ref_Type = 'Occupation'
  LEFT JOIN vwReference industry    on lower(industry.Title)      = lower(a.industry)      and industry.Ref_Type = 'Industry'
  LEFT JOIN vwReference Religion    on lower(Religion.Title)      = lower(a.Religion)      and Religion.Ref_Type = 'Religion'
  LEFT JOIN vwReference Nationality on lower(Nationality.Title)   = lower(a.Nationality)   and Nationality.Ref_Type = 'Nationality'
  LEFT JOIN Identity_Info Contact   on Contact.Alternate_ID      = a.ContactaltID
  ;  

  i integer;
  Current_Address_ID bigint;
  Birth_Place_ID bigint;
  AddressType_ID bigint;
BEGIN
   -- Open the cursor
   Open cur;
   
   LOOP

      FETCH cur INTO rec;
    -- exit when no more row to fetch
      EXIT WHEN NOT FOUND;
 
      SELECT a.ID into Current_Address_ID 
      FROM searchLocation(rec.Current_Address,1) a;
    
      SELECT a.ID into Birth_Place_ID 
      FROM searchLocation(rec.Birth_Place,1) a;
    
      SELECT r.ID into AddressType_ID 
      FROM Reference r
      inner join Reference_Type y on r.Type_ID = y.ID
      WHERE y.Title = 'AddressType' and r.Title = 'Current Address';

     INSERT INTO Identity_Info (
       isPerson, Title, Last_Name, First_Name, 
       Middle_Name, Mother_Maiden_Name, Birthday, 
       Sex, Gender_ID, Civil_Status_ID , Birth_Place_ID, 
       Contact_ID, Alternate_ID, Identity_Map_ID
       )
     SELECT 
       true isPerson, rec.Title, rec.Last_Name, rec.First_Name, 
       rec.Middle_Name, rec.Mother_Maiden_Name, rec.Birthday, 
       rec.Sex, rec.Gender_ID, rec.Civil_Status_ID , Birth_Place_ID Birth_Place_ID, 
       rec.Contact_ID, rec.Alternate_ID, cast(null as bigint) Identity_Map_ID
     ON CONFLICT(Alternate_ID)
     DO UPDATE SET
       isPerson = excluded.isPerson, 
       Title = excluded.Title, 
       Last_Name = excluded.Last_Name, 
       First_Name = excluded.First_Name, 
       Middle_Name = excluded.Middle_Name, 
       Mother_Maiden_Name = excluded.Mother_Maiden_Name, 
       Birthday = excluded.Birthday, 
       Sex = excluded.Sex, 
       Gender_ID = excluded.Gender_ID, 
       Civil_Status_ID  = excluded.Civil_Status_ID , 
       Birth_Place_ID = excluded.Birth_Place_ID, 
       Contact_ID = excluded.Contact_ID,
       Identity_Map_ID = excluded.Identity_Map_ID     
     ;

   
     -- Address_List
     INSERT INTO 
       Address_List(IIID, Detail, Series, 
       Geography_ID, Type_ID, URL)   
     SELECT
      ii.ID IIID, rec.Current_Address_Detail Address_Detail, 1 Series, 
       Current_Address_ID Geography_ID, AddressType_ID AddressType_ID, 
       rec.Current_Address_URL Address_URL
     FROM Identity_Info ii 
     WHERE ii.Alternate_ID = rec.Alternate_ID                   
     ON CONFLICT (IIID, Series)
     DO UPDATE SET 
        Detail      = EXCLUDED.Detail,
        Geography_ID = EXCLUDED.Geography_ID,
        Type_ID      = EXCLUDED.Type_ID, 
        URL         = EXCLUDED.URL;
   
     INSERT INTO Personal_Info(ID, 
       isAdopted, Source_Income_ID, Disability_ID, Occupation_ID, Sector_ID, Industry_ID,   
       Religion_ID, Nationality_ID, Marriage_Date, Known_Language) 
     SELECT ii.ID, 
       rec.isAdopted, rec.Source_Income_ID, rec.Disability_ID, rec.Occupation_ID, null, rec.Industry_ID, 
       rec.Religion_ID, rec.Nationality_ID, rec.Marriage_Date, rec.Known_Language
     FROM Identity_Info ii 
     WHERE ii.Alternate_ID = rec.Alternate_ID
       and not        
       (rec.isAdopted is null and 
        rec.Source_Income_ID is null and
        rec.Disability_ID is null and
        rec.Occupation_ID is null and
        rec.Industry_ID is null and
        rec.Religion_ID is null and
        rec.Nationality_ID is null and
        rec.Marriage_Date is null and
        rec.Known_Language is null)
       
     ON CONFLICT(ID)
     DO UPDATE SET
       isAdopted = excluded.isAdopted, Source_Income_ID = excluded.Source_Income_ID, Disability_ID = excluded.Disability_ID, 
       Occupation_ID = excluded.Occupation_ID, Sector_ID = excluded.Sector_ID, Industry_ID = excluded.Industry_ID,   
       Religion_ID = excluded.Religion_ID, Nationality_ID = excluded.Religion_ID, Marriage_Date = excluded.Marriage_Date, 
       Known_Language = excluded.Marriage_Date;          


     --EXIT WHEN TRUE;

   END LOOP;
  
   -- Close the cursor
   CLOSE cur;

END; $$ ;

/*
--------------------------------------------------------------------------------------------
DO $$
--------------------------------------------------------------------------------------------
DECLARE 
  Titles TEXT DEFAULT '';
  rec   RECORD;
  cur CURSOR FOR 

  SELECT 
    ii.IIID,
    a.Last_Name, a.First_Name, a.Middle_Name, a.Mother_Maiden_Name, 
    cast(a.Birthday as Date) Birthday, 
    Title, Civil.ID Civil_Status_ID , a.Birth_Place, 
    a.Sex, Gender.ID Gender_ID, 
    a.Current_Address, a.Current_Address_Detail, a.Current_Address_URL,
    Contact.ID Contact_ID, income.ID Source_Income_ID, 
    Disability.ID Disability_ID, a.isAdopted, Occupation.ID Occupation_ID, 
    a.Sector, industry.ID Industry_ID, Religion.ID Religion_ID, 
    Nationality.ID Nationality_ID, cast(a.Marriage_Date as Date) Marriage_Date, a.Known_Language, 
    a.Alternate_ID, cast(a.data as jsonb) data1
    
 FROM (Values
      ('Mercado','Roderick','de Guzman','Florcita Dia de Guzman','1979/04/14',null,'Married','San Pablo City',TRUE,'Male','Soledad San Pablo City, Laguna','Purok 3',null,null,'100','Salary',null,FALSE,'Computer Programmer ',null,'Financial services; professional services','United Church of Christ in the Philippines','Filipino','1997/09/21','Tagalog; English',null),
      ('Mercado','Olive','Maluping','Amparo Lopez','1976/06/09',null,'Married','San Pablo City',FALSE,'Female','Soledad San Pablo City, Laguna','Purok 3',null,null,'101','Salary',null,FALSE,'Elementary School Teachers, Except Special Education',null,'Education ','United Church of Christ in the Philippines','Filipino','1997/09/21','Tagalog; English',null),
      ('Anicete','Marlyn','B',null,null,null,null,null,FALSE,'Female',null,null,null,null,'1001',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Anilor',null,null,null,null,'Ptr.',null,null,FALSE,'Female',null,null,null,null,'1002',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Apole','Ivie',null,null,null,'Rev.',null,null,FALSE,'Female',null,null,null,null,'1003',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Baldonado','Eric','P',null,null,'DM',null,null,TRUE,'Male',null,null,null,null,'1004',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Bayot','Joel',null,null,null,'CM',null,null,TRUE,'Male',null,null,null,null,'1005',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Billena','Ruth',null,null,null,'CM',null,null,FALSE,'Female',null,null,null,null,'1006',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Cabico','Butch',null,null,null,'Ptr.',null,null,TRUE,'Male',null,null,null,null,'1007',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Cabiso','Elmer',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1008',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Del Valle','Elena',null,null,null,'CM',null,null,FALSE,'Female',null,null,null,null,'1009',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Destor','Danny',null,null,null,null,null,null,TRUE,'Male',null,null,null,null,'1010',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Destor','Gener','S',null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1011',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Divino','Samson',null,null,null,'CM',null,null,TRUE,'Male',null,null,null,null,'1012',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Egar','Juliet',null,null,null,'Ptr.',null,null,FALSE,'Female',null,null,null,null,'1013',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Ercia','Annie',null,null,null,null,null,null,FALSE,'Female',null,null,null,null,'1014',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Esperanza','Ariel',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1015',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Espineli','Ceasar',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1016',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Fabreag','Leovino',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1017',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Faller','Martha',null,null,null,'Rev.',null,null,FALSE,'Female',null,null,null,null,'1018',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Fortes','Billy',null,null,null,'Atty.',null,null,TRUE,'Male',null,null,null,null,'1019',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Francisco','Danny',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1020',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Galo','Deogracio',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1021',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Guerrero','Angelo',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1022',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Guerrero','John Angelo',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1023',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Igdanes','Lemuel',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1024',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Laudencia','Edgar',null,null,null,null,null,null,TRUE,'Male',null,null,null,null,'1025',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Laundencia','Edgar',null,null,null,'Ptr.',null,null,TRUE,'Male',null,null,null,null,'1026',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Magculang','Ricky',null,null,null,null,null,null,TRUE,'Male',null,null,null,null,'1027',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Marinas','Benny',null,null,null,'Ptr.',null,null,TRUE,'Male',null,null,null,null,'1028',null,null,NULL,null,null,null,null,null,null,null,null),
      ('McDivith','Helen',null,null,null,null,null,null,FALSE,'Female',null,null,null,null,'1029',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Medel','Dandy',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1030',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Millares','Nehemia',null,null,null,'Rev.',null,null,FALSE,'Female',null,null,null,null,'1031',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Miriam',null,null,null,null,null,null,null,FALSE,'Female',null,null,null,null,'1032',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Obando',null,null,null,null,'Ptr.',null,null,FALSE,'Female',null,null,null,null,'1033',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Onate','Biboy',null,null,null,'Ptr.',null,null,TRUE,'Male',null,null,null,null,'1034',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Paller','Victor','L',null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1035',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Pedrina','Yolly',null,null,null,null,null,null,FALSE,'Female',null,null,null,null,'1036',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Plotado','Elizabeth',null,null,null,'Ptr.',null,null,FALSE,'Female',null,null,null,null,'1037',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Pontillas','Randy',null,null,null,'Ptr.',null,null,TRUE,'Male',null,null,null,null,'1038',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Pring','Luis',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1039',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Robles','Gilbert',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1040',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Romaquin','Nathaniel',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1041',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Roxas','Noel',null,null,null,'Ptr.',null,null,TRUE,'Male',null,null,null,null,'1042',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Saliendra','Luisito','H',null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1043',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Saliendra','Ronnie',null,null,null,null,null,null,TRUE,'Male',null,null,null,null,'1044',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Samonte','Enong',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1045',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Sediarin',null,null,null,null,'Rev.',null,null,FALSE,'Female',null,null,null,null,'1046',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Sito',null,null,null,null,'Ptr.',null,null,TRUE,'Male',null,null,null,null,'1047',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Solis','Juliet',null,null,null,'Rev.',null,null,FALSE,'Female',null,null,null,null,'1048',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Sopranes','Amy',null,null,null,'Ptr.',null,null,FALSE,'Female',null,null,null,null,'1049',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Suarez','Gilbert',null,null,null,'Ptr.',null,null,TRUE,'Male',null,null,null,null,'1050',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Suyom','June',null,null,null,null,null,null,TRUE,'Male',null,null,null,null,'1051',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Tadiosa','Gailry',null,null,null,'Ptr.',null,null,TRUE,'Male',null,null,null,null,'1052',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Tendero','Bernabeth',null,null,null,null,null,null,FALSE,'Female',null,null,null,null,'1053',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Tendero','Joel','E',null,null,'Bp.',null,null,TRUE,'Male',null,null,null,null,'1054',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Turgo','Billy',null,null,null,'Ptr.',null,null,TRUE,'Male',null,null,null,null,'1055',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Skagersten','Anna',null,null,null,null,null,null,FALSE,'Female',null,null,null,null,'1056',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Espineli','Ceazar',null,null,null,'CM',null,null,TRUE,'Male',null,null,null,null,'1057',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Saliendra','Celso',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1058',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Dela Cruz','Eleanor',null,null,null,'CM',null,null,TRUE,'Male',null,null,null,null,'1059',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Tequis','Hamuel',null,null,null,'Bp.',null,null,TRUE,'Male',null,null,null,null,'1060',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Delos Santos','Ian Lloid',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1061',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Calimutan','Joram',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1062',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Bueno','Junwel','M',null,null,'CM',null,null,TRUE,'Male',null,null,null,null,'1063',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Ledo','Kreng',null,null,null,null,null,null,TRUE,'Male',null,null,null,null,'1064',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Peralta','Laarni',null,null,null,'Ptr.',null,null,FALSE,'Female',null,null,null,null,'1065',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Dolo','Leiza',null,null,null,'Rev.',null,null,FALSE,'Female',null,null,null,null,'1066',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Labuntog','Melzar',null,null,null,'Bp.',null,null,TRUE,'Male',null,null,null,null,'1067',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Cantor','Ofelia ',null,null,null,null,null,null,FALSE,'Female',null,null,null,null,'1068',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Arce','Pastive',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1069',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Rizal','Pumphrey',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1070',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Abad','Reynaldo',null,null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1071',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Bahia','Rizaldo','Y',null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1072',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Eslabon','Robert','T',null,null,'Rev.',null,null,TRUE,'Male',null,null,null,null,'1073',null,null,NULL,null,null,null,null,null,null,null,null),
      ('Tripulca','Wilmer',null,null,null,'Ptr.',null,null,TRUE,'Male',null,null,null,null,'1074',null,null,NULL,null,null,null,null,null,null,null,null)
      )   
    a(Last_Name, First_Name, Middle_Name, Mother_Maiden_Name, Birthday, 
      Title, Civil_Status, Birth_Place, Sex, Gender, 
      Current_Address, Current_Address_Detail, Current_Address_URL, ContactaltID, 
      Alternate_ID, Source_Income, Disability, isAdopted, Occupation, Sector, 
      industry, Religion, Nationality, Marriage_Date, Known_Language, data)  
      
  LEFT JOIN vwperson ii             on ii.Alternate_ID = a.Alternate_ID
  LEFT JOIN vwReference Title       on lower(Title.Title) = lower(a.Title) and Title.Ref_Type = 'Title'  
  LEFT JOIN vwReference Civil       on lower(Civil.Title) = lower(a.Civil_Status)          and Civil.Ref_Type = 'Civil_Status'  
  LEFT JOIN vwReference Gender      on lower(Gender.Title)        = lower(a.Gender)        and Gender.Ref_Type = 'Gender'
  LEFT JOIN vwReference Disability  on lower(Disability.Title)    = lower(a.Disability)    and Disability.Ref_Type = 'Disabilities'
  LEFT JOIN vwReference income      on lower(income.Title)        = lower(a.Source_Income) and income.Ref_Type = 'SourceofIncome'
  LEFT JOIN vwReference Occupation  on lower(Occupation.Title)    = lower(a.Occupation)    and Occupation.Ref_Type = 'Occupation'
  LEFT JOIN vwReference industry    on lower(industry.Title)      = lower(a.industry)      and industry.Ref_Type = 'Industry'
  LEFT JOIN vwReference Religion    on lower(Religion.Title)      = lower(a.Religion)      and Religion.Ref_Type = 'Religion'
  LEFT JOIN vwReference Nationality on lower(Nationality.Title)   = lower(a.Nationality)   and Nationality.Ref_Type = 'Nationality'
  LEFT JOIN Identity_Info Contact   on Contact.Alternate_ID      = a.ContactaltID
  ;  

  i integer;
  Current_Address_ID bigint;
  Birth_Place_ID bigint;
  AddressType_ID bigint;
BEGIN
   -- Open the cursor
   Open cur;
   
   LOOP

      FETCH cur INTO rec;
    -- exit when no more row to fetch
      EXIT WHEN NOT FOUND;
 
      SELECT a.ID into Current_Address_ID 
      FROM searchLocation(rec.Current_Address,1) a;
    
      SELECT a.ID into Birth_Place_ID 
      FROM searchLocation(rec.Birth_Place,1) a;
    
      SELECT r.ID into AddressType_ID 
      FROM Reference r
      inner join Reference_Type y on r.Type_ID = y.ID
      WHERE y.Title = 'AddressType' and r.Title = 'Current Address';

     INSERT INTO Identity_Info (
       isPerson, Title_ID, Last_Name, First_Name, 
       Middle_Name, Mother_Maiden_Name, Birthday, 
       Sex, Gender_ID, Civil_Status_ID , Birth_Place_ID, 
       Contact_ID, Alternate_ID, Identity_Map_ID
       )
     SELECT 
       true isPerson, rec.Title_ID, rec.Last_Name, rec.First_Name, 
       rec.Middle_Name, rec.Mother_Maiden_Name, rec.Birthday, 
       rec.Sex, rec.Gender_ID, rec.Civil_Status_ID , Birth_Place_ID Birth_Place_ID, 
       rec.Contact_ID, rec.Alternate_ID, cast(null as bigint) Identity_Map_ID
     ON CONFLICT(Alternate_ID)
     DO UPDATE SET
       isPerson = excluded.isPerson, 
       Title_ID = excluded.Title_ID, 
       Last_Name = excluded.Last_Name, 
       First_Name = excluded.First_Name, 
       Middle_Name = excluded.Middle_Name, 
       Mother_Maiden_Name = excluded.Mother_Maiden_Name, 
       Birthday = excluded.Birthday, 
       Sex = excluded.Sex, 
       Gender_ID = excluded.Gender_ID, 
       Civil_Status_ID  = excluded.Civil_Status_ID , 
       Birth_Place_ID = excluded.Birth_Place_ID, 
       Contact_ID = excluded.Contact_ID,
       Identity_Map_ID = excluded.Identity_Map_ID     
     ;

   
     -- Address_List
     INSERT INTO 
       Address_List(IIID, Detail, Series, 
       Geography_ID, Type_ID, URL)   
     SELECT
      ii.ID IIID, rec.Current_Address_Detail Address_Detail, 1 Series, 
       Current_Address_ID Geography_ID, AddressType_ID AddressType_ID, 
       rec.Current_Address_URL Address_URL
     FROM Identity_Info ii 
     WHERE ii.Alternate_ID = rec.Alternate_ID                   
     ON CONFLICT (IIID, Series)
     DO UPDATE SET 
        Detail      = EXCLUDED.Detail,
        Geography_ID = EXCLUDED.Geography_ID,
        Type_ID      = EXCLUDED.Type_ID, 
        URL         = EXCLUDED.URL;
   
     INSERT INTO Personal_Info(ID, 
       isAdopted, Source_Income_ID, Disability_ID, Occupation_ID, Sector_ID, Industry_ID,   
       Religion_ID, Nationality_ID, Marriage_Date, Known_Language) 
     SELECT ii.ID, 
       rec.isAdopted, rec.Source_Income_ID, rec.Disability_ID, rec.Occupation_ID, null, rec.Industry_ID, 
       rec.Religion_ID, rec.Nationality_ID, rec.Marriage_Date, rec.Known_Language
     FROM Identity_Info ii 
     WHERE ii.Alternate_ID = rec.Alternate_ID
       and not        
       (rec.isAdopted is null and 
        rec.Source_Income_ID is null and
        rec.Disability_ID is null and
        rec.Occupation_ID is null and
        rec.Industry_ID is null and
        rec.Religion_ID is null and
        rec.Nationality_ID is null and
        rec.Marriage_Date is null and
        rec.Known_Language is null)
       
     ON CONFLICT(ID)
     DO UPDATE SET
       isAdopted = excluded.isAdopted, Source_Income_ID = excluded.Source_Income_ID, Disability_ID = excluded.Disability_ID, 
       Occupation_ID = excluded.Occupation_ID, Sector_ID = excluded.Sector_ID, Industry_ID = excluded.Industry_ID,   
       Religion_ID = excluded.Religion_ID, Nationality_ID = excluded.Religion_ID, Marriage_Date = excluded.Marriage_Date, 
       Known_Language = excluded.Marriage_Date;          


     --EXIT WHEN TRUE;

   END LOOP;
  
   -- Close the cursor
   CLOSE cur;

END; $$ ;

*/