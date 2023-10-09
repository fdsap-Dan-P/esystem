---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Identity_Info (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Identity_Map_ID bigint NULL,
  isPerson bool NOT NULL,
  Alternate_ID VarChar(50) NULL,
  Title varchar(15) NULL,
  Last_Name varchar(50) NOT NULL,
  First_Name varchar(50) NULL,
  Middle_Name varchar(50) NULL,
  Suffix_Name varchar(10) NULL,
  Professional_Suffixes varchar(50) NULL,
  Mother_Maiden_Name varchar(150) NULL,
  Birthday Date NULL,
  Sex bool NULL,
  Gender_id bigint NULL,
  Civil_Status_ID bigint NULL,
  Birth_Place_ID bigint NULL,
  Contact_ID bigint NULL,
  Phone VarChar(20) NULL,
  Email VarChar(100) NULL,
  Simple_Name varchar(100) NULL,
  Vec_Simple_Name tsvector,
  Vec_Full_Simple_Name tsvector,
  Other_Info jsonb NULL,

  CONSTRAINT Identity_Info_pkey PRIMARY KEY (ID),
  CONSTRAINT fkIdentity_InfoMap FOREIGN KEY (Identity_Map_ID) REFERENCES Identity_Info(ID),
  CONSTRAINT idxIdentity_Infoalt UNIQUE (Alternate_ID),
  CONSTRAINT fkIdentity_InfoGender FOREIGN KEY (Gender_id) REFERENCES Reference(ID),
  CONSTRAINT fkIdentity_InfoCivil_Status FOREIGN KEY (Civil_Status_ID ) REFERENCES Reference(ID),
  CONSTRAINT fkIdentity_InfoBirth_Place FOREIGN KEY (Birth_Place_ID) REFERENCES Geography(ID),
  CONSTRAINT fkIdentity_Infocontanct FOREIGN KEY (Contact_ID) REFERENCES Identity_Info(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxIdentity_Info_UUID ON public.Identity_Info(UUID);

CREATE INDEX IF NOT EXISTS idxIdentity_Info_Email ON public.Identity_Info(Phone);
CREATE INDEX IF NOT EXISTS idxIdentity_Info_Phone ON public.Identity_Info(Email);


DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fkofficeofficer') THEN    
      ALTER TABLE Office ADD CONSTRAINT fkOfficeOfficer FOREIGN KEY (Officer_IIID) REFERENCES Identity_info(ID);

    END IF;
END;
$$;


----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwIdentity_Info
----------------------------------------------------------------------------------------
AS SELECT 
    ii.ID, mr.UUID,
    ii.Title,
    ii.Last_Name, ii.First_Name, ii.Middle_Name, ii.Mother_Maiden_Name,
    ii.Birthday, ii.Sex,
    
    Gender.ID Gender_id, Gender.UUID Gender_UUID, Gender.Title Gender,    
    
    Civil.ID Civil_Status_ID , Civil.UUID Civil_Status_UUID, Civil.Title Civil_Status,
        
    bPlace.ID Birth_Place_ID, bPlace.UUID Birth_Place_UUID,
    bPlace.Location Birth_Place, bPlace.Full_Location Full_Birth_Place,
    bPlace.Address_URL Birth_PlaceURL, 
    
    Contact.ID Contact_ID, Contact.UUID ContactUUID,
    Contact.Last_Name Contact_Last_Name, Contact.First_Name Contact_First_Name, 
    Contact.Middle_Name Contact_Middle_Name,
    
    ii.Alternate_ID, 
    ii.Phone, ii.Email, ii.Identity_Map_ID, ii.Simple_Name, 
    
    mr.Mod_Ctr,
    ii.Other_Info,
    mr.Created,
    mr.Updated 
  FROM Identity_Info ii
  LEFT JOIN Main_Record mr on mr.UUID = ii.UUID
  LEFT JOIN Reference Gender ON Gender.ID = ii.Gender_id
  LEFT JOIN Reference Civil ON Civil.ID = ii.civil_status_id 
  LEFT JOIN vwGeography bPlace ON bPlace.ID = Birth_Place_ID

  LEFT JOIN Identity_Info Contact ON Contact.ID = ii.Contact_ID;

---------------------------------------------------------------------------
-- INSERT Identity_Info Trigger
CREATE or REPLACE FUNCTION Identity_InfoINSERT() returns trigger AS 
---------------------------------------------------------------------------
$$ 
DECLARE 
   ctr int8 := Nextval('Mod_Ctr_seq');
   upd timestamptz := CURRENT_TIMESTAMP;
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
 
    
  NEW.Simple_Name := Simpleword(NEW.First_Name || ', '::text || NEW.Last_Name || ', '::text || NEW.Middle_Name);   
  NEW.Vec_Simple_Name := to_tsvector('Simple'::regConfig, unAccent(NEW.Simple_Name::text));
  NEW.Vec_Full_Simple_Name := setweight( to_tsvector('Simple'::regConfig, unAccent(NEW.Simple_Name::text)), 'A'::"char" );
  
  PERFORM pg_notify('mychan', payload::text);
  RETURN NEW;
END $$ Language plpgsql;

/*
: invalid input syntax for type boolean: "(Olive,t,Mercado)"
  Where: PL/pgSQL function Identity_Infoupdate() line 3 at IF
SQL statement "INSERT INTO Identity_Info (
       isPerson, Title_ID, Last_Name, First_Name, 
       Middle_Name, Mother_Maiden_Name, Birthday, 

*/

----------------------------------------------------------------------------------------------
-- UPDATE Identity_Info Trigger
CREATE OR REPLACE FUNCTION Identity_InfoUPDATE() RETURNS trigger Language plpgsql
----------------------------------------------------------------------------------------------
AS $function$ 
DECLARE 
   ctr int8 := Nextval('Mod_Ctr_seq');
   upd timestamptz := CURRENT_TIMESTAMP;
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

  UPDATE Main_Record 
  SET Updated = Upd
  WHERE UUID = NEW.UUID;  

  IF NOT ( COALESCE(OLD.First_Name,'') = COALESCE(NEW.First_Name,'')  AND COALESCE(OLD.Last_Name,'') = COALESCE(NEW.Last_Name,'') ) THEN 
      NEW.Simple_Name := Simpleword(NEW.First_Name || ', '::text || NEW.Last_Name || ', '::text || NEW.Middle_Name);   
      NEW.Vec_Simple_Name := to_tsvector('Simple'::regConfig, unAccent(NEW.Simple_Name::text));
    NEW.Vec_Full_Simple_Name := setweight( to_tsvector('Simple'::regConfig, unAccent(NEW.Simple_Name::text)), 'A'::"char" );
  END IF;
  PERFORM pg_notify('mychan', payload::text);
  RETURN NEW;
END 
$function$
;

DROP TRIGGER IF EXISTS trgIdentity_Infoupd ON Identity_Info;
DROP TRIGGER IF EXISTS trgIdentity_InfoIns ON Identity_Info;

---------------------------------------------------------------------------
CREATE TRIGGER trgIdentity_InfoUpd
---------------------------------------------------------------------------
  BEFORE UPDATE ON public.Identity_Info
  FOR EACH ROW
  WHEN (OLD.* IS DISTINCT FROM NEW.*)
  EXECUTE PROCEDURE Identity_InfoUPDATE();

---------------------------------------------------------------------------
CREATE TRIGGER trgIdentity_InfoIns
---------------------------------------------------------------------------
  BEFORE INSERT ON Identity_Info
  FOR EACH ROW
  EXECUTE PROCEDURE Identity_InfoINSERT();

  DROP TRIGGER IF EXISTS trgIdentity_Info_del on Identity_Info;
---------------------------------------------------------------------------
CREATE TRIGGER trgIdentity_Info_del
---------------------------------------------------------------------------
    AFTER DELETE ON Identity_Info
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

-- View indexes:
CREATE INDEX IF NOT EXISTS idxIdentity_Infolexeme 
  ON public.Identity_Info
  USING gin (Vec_Full_Simple_Name);

CREATE OR REPLACE FUNCTION searchName (
    pName character varying, 
    psearchlimit integer) 
  RETURNS 
  TABLE (
    ID bigint,
    Last_Name character varying,
    First_Name character varying,
    Simple_Name character varying,
    rnkfull real
    ) 
  AS $$
  DECLARE 
      locor tsquery;
      locand tsquery;
      loc2 tsquery;
  BEGIN
    pName = SimpleWord(pName); 
    pName = pName || ' ';
    locor = plainto_tsquery_or(pName);
    locand = plainto_tsquery_and(pName);
    loc2 = lower(SUBSTR(pName,1, Position(' ' IN pName)));

    RETURN QUERY 
    SELECT  
      g.ID, g.Last_Name, g.First_Name, g.Simple_Name, ts_rank(Vec_Full_Simple_Name,loc2) rnkfull   
    FROM Identity_Info g   
    WHERE Vec_Full_Simple_Name @@ locor
    ORDER BY rnkfull desc
    limit psearchlimit;
  END; $$ Language 'plpgsql';
