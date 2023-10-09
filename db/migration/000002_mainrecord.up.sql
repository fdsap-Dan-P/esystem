CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
---------------------------------------------------------------------------
CREATE Sequence IF NOT EXISTS public.Mod_Ctr_seq
---------------------------------------------------------------------------
  INCREMENT BY 1
  MINValue 1
  MAXValue 9223372036854775807
  START 1
  CACHE 1
  NO CYCLE;

----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Main_Record (
----------------------------------------------------------------------------------------
  UUID     uuid NOT NULL,
  Mod_Ctr   int8 NOT NULL,
  Tablename VarChar(255),
  Created  timestamptz NULL,
  Updated  timestamptz NULL,
  CONSTRAINT Main_Record_pkey PRIMARY KEY (UUID)
);

----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Modified (
----------------------------------------------------------------------------------------
  Mod_Ctr int8 NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  
  Updated timestamptz NULL,
  Other_Info jsonb NULL,
  CONSTRAINT Modified_pkey PRIMARY KEY (Mod_Ctr),
  CONSTRAINT fkModifiedID FOREIGN KEY (UUID) REFERENCES Main_Record(UUID)
);

CREATE INDEX IF NOT EXISTS idxModified_UUID ON public.Modified(UUID);

-- INSERT Trigger


---------------------------------------------------------------------------
CREATE or REPLACE FUNCTION trgGenericInsert() returns trigger AS 
---------------------------------------------------------------------------
$$ 
DECLARE 
   ctr int8 := Nextval('Mod_Ctr_seq');
   upd timestamptz := CURRENT_TIMESTAMP;
   payload JSON;
   tbl VarChar(255);
   existtbl VarChar(255);
BEGIN
   
   tbl = TG_TABLE_NAME::regclass::text;

   payload = json_build_object( 
      'action', TG_OP,
      'table',  tbl,
      'data',  json_build_object( 
          'UUID', NEW.UUID,
          'Mod_Ctr', ctr
      ) 
   );
 
--   SELECT TableName into existtbl FROM Main_Record mr WHERE mr.UUID = NEW.UUID;
--   existtbl = coalesce(existtbl, tbl);
 
--   IF existtbl <> tbl THEN
 --    RAISE 'Duplicate UUID: %', UUID USING ERRCODE = 'unique_violation';
 --  ELSE
     INSERT INTO Main_Record(UUID, TableName, Mod_Ctr, Created)
     SELECT NEW.UUID, tbl, ctr, upd
     ON CONFLICT(UUID) WHERE tbl = TableName DO NOTHING;
 --  END IF;
 
   INSERT INTO Modified(Mod_Ctr, UUID, Updated)
   SELECT ctr, NEW.UUID, upd;

   PERFORM pg_notify('mychan', payload::text);   
   RETURN NEW;
   
END $$ Language plpgsql;



-- UPDATE Trigger
---------------------------------------------------------------------------
CREATE or REPLACE FUNCTION trgGenericUpdate() returns trigger AS 
---------------------------------------------------------------------------
$$ 
DECLARE 
   ctr int8 := Nextval('Mod_Ctr_seq');
   Upd timestamptz := CURRENT_TIMESTAMP;
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
  ON CONFLICT(UUID) DO 
  UPDATE SET Updated = Upd;    


  INSERT INTO Modified(Mod_Ctr, UUID, Updated)
  ValueS(ctr, NEW.UUID, upd);

--   UPDATE Main_Record 
--   SET Updated = Upd
--   WHERE UUID = NEW.UUID;

  PERFORM pg_notify('mychan', payload::text);
  RETURN NEW;
END $$ Language plpgsql;

---------------------------------------------------------------------------
--- Specs Trigger
---------------------------------------------------------------------------
CREATE or REPLACE FUNCTION trgGenericSpecsInsert() returns trigger AS 
---------------------------------------------------------------------------
$$ 
DECLARE 
   ctr int8 := Nextval('Mod_Ctr_seq');
   upd timestamptz := CURRENT_TIMESTAMP;
   payload JSON;
   tbl VarChar(255);
   existtbl VarChar(255);
   SpecsCode VarChar(50);
BEGIN
   
   tbl = TG_TABLE_NAME::regclass::text;

   payload = json_build_object( 
      'action', TG_OP,
      'table',  tbl,
      'data',  json_build_object( 
          'UUID', NEW.UUID,
          'Mod_Ctr', ctr
      ) 
   );
 
--   SELECT TableName into existtbl FROM Main_Record mr WHERE mr.UUID = NEW.UUID;
--   existtbl = coalesce(existtbl, tbl);
 
--   IF existtbl <> tbl THEN
 --    RAISE 'Duplicate UUID: %', UUID USING ERRCODE = 'unique_violation';
 --  ELSE
     INSERT INTO Main_Record(UUID, TableName, Mod_Ctr, Created)
     SELECT NEW.UUID, tbl, ctr, upd
     ON CONFLICT(UUID) WHERE tbl = TableName DO NOTHING;
 --  END IF;
 
   INSERT INTO Modified(Mod_Ctr, UUID, Updated)
   SELECT ctr, NEW.UUID, upd;

   SpecsCode := Short_Name FROM Reference WHERE id = NEW.Specs_Id;
   NEW.Specs_Code := SpecsCode;
   PERFORM pg_notify('mychan', payload::text);   

   RETURN NEW;
END $$ Language plpgsql;

-- UPDATE Specs Trigger
---------------------------------------------------------------------------
CREATE or REPLACE FUNCTION trgGenericSpecsUpdate() returns trigger AS 
---------------------------------------------------------------------------
$$ 
DECLARE 
   ctr int8 := Nextval('Mod_Ctr_seq');
   Upd timestamptz := CURRENT_TIMESTAMP;
   payload JSON;
   SpecsCode VarChar(50);
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
  ON CONFLICT(UUID) DO 
  UPDATE SET Updated = Upd;    


  INSERT INTO Modified(Mod_Ctr, UUID, Updated)
  ValueS(ctr, NEW.UUID, upd);

--   UPDATE Main_Record 
--   SET Updated = Upd
--   WHERE UUID = NEW.UUID;

  SpecsCode := Short_Name FROM Reference WHERE id = NEW.Specs_Id;
  NEW.Specs_Code := SpecsCode;
  PERFORM pg_notify('mychan', payload::text);

  RETURN NEW;
END $$ Language plpgsql;

---------------------------------------------------------------------------
--- Item Trigger
---------------------------------------------------------------------------
CREATE or REPLACE FUNCTION trgGenericItemInsert() returns trigger AS 
---------------------------------------------------------------------------
$$ 
DECLARE 
   ctr int8 := Nextval('Mod_Ctr_seq');
   upd timestamptz := CURRENT_TIMESTAMP;
   payload JSON;
   tbl VarChar(255);
   existtbl VarChar(255);
   ItemCode VarChar(50);
BEGIN
   
   tbl = TG_TABLE_NAME::regclass::text;

   payload = json_build_object( 
      'action', TG_OP,
      'table',  tbl,
      'data',  json_build_object( 
          'UUID', NEW.UUID,
          'Mod_Ctr', ctr
      ) 
   );
 
--   SELECT TableName into existtbl FROM Main_Record mr WHERE mr.UUID = NEW.UUID;
--   existtbl = coalesce(existtbl, tbl);
 
--   IF existtbl <> tbl THEN
 --    RAISE 'Duplicate UUID: %', UUID USING ERRCODE = 'unique_violation';
 --  ELSE
     INSERT INTO Main_Record(UUID, TableName, Mod_Ctr, Created)
     SELECT NEW.UUID, tbl, ctr, upd
     ON CONFLICT(UUID) WHERE tbl = TableName DO NOTHING;
 --  END IF;
 
   INSERT INTO Modified(Mod_Ctr, UUID, Updated)
   SELECT ctr, NEW.UUID, upd;

  ItemCode := Short_Name FROM Reference WHERE id = NEW.Item_Id;
  NEW.Item_Code := ItemCode;
  PERFORM pg_notify('mychan', payload::text);

  RETURN NEW;
END $$ Language plpgsql;

-- UPDATE Item Trigger
---------------------------------------------------------------------------
CREATE or REPLACE FUNCTION trgGenericItemUpdate() returns trigger AS 
---------------------------------------------------------------------------
$$ 
DECLARE 
   ctr int8 := Nextval('Mod_Ctr_seq');
   Upd timestamptz := CURRENT_TIMESTAMP;
   payload JSON;
   ItemCode VarChar(50);
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
  ON CONFLICT(UUID) DO 
  UPDATE SET Updated = Upd;    


  INSERT INTO Modified(Mod_Ctr, UUID, Updated)
  ValueS(ctr, NEW.UUID, upd);

--   UPDATE Main_Record 
--   SET Updated = Upd
--   WHERE UUID = NEW.UUID;

  ItemCode := Short_Name FROM Reference WHERE id = NEW.Item_Id;
  NEW.Item_Code := ItemCode;
  PERFORM pg_notify('mychan', payload::text);

  RETURN NEW;
END $$ Language plpgsql;


---------------------------------------------------------------------------
--- config Item Trigger
---------------------------------------------------------------------------
CREATE or REPLACE FUNCTION trgGenericConfigInsert() returns trigger AS 
---------------------------------------------------------------------------
$$ 
DECLARE 
   ctr int8 := Nextval('Mod_Ctr_seq');
   upd timestamptz := CURRENT_TIMESTAMP;
   payload JSON;
   tbl VarChar(255);
   existtbl VarChar(255);
   ItemCode VarChar(50);
   configCode VarChar(50);
BEGIN
   
   tbl = TG_TABLE_NAME::regclass::text;

   payload = json_build_object( 
      'action', TG_OP,
      'table',  tbl,
      'data',  json_build_object( 
          'UUID', NEW.UUID,
          'Mod_Ctr', ctr
      ) 
   );
 
--   SELECT TableName into existtbl FROM Main_Record mr WHERE mr.UUID = NEW.UUID;
--   existtbl = coalesce(existtbl, tbl);
 
--   IF existtbl <> tbl THEN
 --    RAISE 'Duplicate UUID: %', UUID USING ERRCODE = 'unique_violation';
 --  ELSE
     INSERT INTO Main_Record(UUID, TableName, Mod_Ctr, Created)
     SELECT NEW.UUID, tbl, ctr, upd
     ON CONFLICT(UUID) WHERE tbl = TableName DO NOTHING;
 --  END IF;
 
   INSERT INTO Modified(Mod_Ctr, UUID, Updated)
   SELECT ctr, NEW.UUID, upd;

  configCode := Short_Name FROM Reference WHERE id = NEW.config_Id;
  NEW.config_Code := configCode;

  PERFORM pg_notify('mychan', payload::text);

  RETURN NEW;
END $$ Language plpgsql;

-- UPDATE config Item Trigger
---------------------------------------------------------------------------
CREATE or REPLACE FUNCTION trgGenericConfigUpdate() returns trigger AS 
---------------------------------------------------------------------------
$$ 
DECLARE 
   ctr int8 := Nextval('Mod_Ctr_seq');
   Upd timestamptz := CURRENT_TIMESTAMP;
   payload JSON;
   ItemCode VarChar(50);
   configCode VarChar(50);
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
  ON CONFLICT(UUID) DO 
  UPDATE SET Updated = Upd;    


  INSERT INTO Modified(Mod_Ctr, UUID, Updated)
  ValueS(ctr, NEW.UUID, upd);

--   UPDATE Main_Record 
--   SET Updated = Upd
--   WHERE UUID = NEW.UUID;

  configCode := Short_Name FROM Reference WHERE id = NEW.config_Id;
  NEW.config_Code := configCode;

  PERFORM pg_notify('mychan', payload::text);

  RETURN NEW;
END $$ Language plpgsql;


---------------------------------------------------------------------------
CREATE or REPLACE FUNCTION trgGenericDelete() returns trigger AS 
---------------------------------------------------------------------------
$$ 
DECLARE 
   ctr int8 := Nextval('Mod_Ctr_seq');
   Upd timestamptz := CURRENT_TIMESTAMP;
   payload JSON;
BEGIN

  payload = json_build_object( 
       'action', TG_OP, 
       'data', json_build_object('UUID', OLD.UUID ) 
     );

  INSERT INTO Main_Record(UUID, TableName, Mod_Ctr, Created)
  SELECT OLD.UUID, TG_TABLE_NAME::regclass::text, ctr, upd
  ON CONFLICT(UUID) DO 
  UPDATE SET Updated = Upd;    
  
  INSERT INTO Modified(Mod_Ctr, UUID, Updated, Other_Info)
  ValueS(ctr, OLD.UUID, upd, row_to_json( NEW ));
     
  PERFORM pg_notify('mychan', payload::text);
  RETURN NULL;
END $$ Language plpgsql;

--------------------------------------------------------------------------------------------------
CREATE EXTENSION IF NOT EXISTS unAccent;
--------------------------------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION simpleword
--------------------------------------------------------------------------------------------------
  (a character varying)
  
RETURNS character varying AS $BODY$
DECLARE
  i Int = 0;
  s character varying = ' ';
  t character varying(10);
  t2 character varying(2);
  t3 character varying(3);
  p character varying(1) = '';
BEGIN  
  IF (a is null) THEN
     RETURN '';    
  END IF;
  a = unaccent(upper(a)) || '  ';
  a = Replace(a,'maria','ma');
  a = Replace(a,'barangay','brgy');
  a = Replace(a,'belle','bel'); 

  a = replace(a, chr(31),'');
  a = replace(a, chr(30),'');
  a = replace(a, chr(9),'');
  a = replace(a, '%','');
  LOOP 
    i = i + 1;
    EXIT WHEN i > LENGTH(a);
    t = SubString(a,i,1); 
    t2 = SubString(a,i,2); 


    IF '·•….,-—–\/;”“' like '%t%' THEN t = ' '; END IF; 
      
    IF 'ABCDEFGHIJKLMNOPQRSTUVWXYZÑ1234567890¤ ' NOT LIKE '%' || t || '%' THEN 
       t = '';
    END IF;
  
    IF t2 in ('EA','EE') THEN 
      t = 'I';  
      i = i + 1;
    ELSIF t2 = 'OO' THEN 
      t = 'U'; 
      i = i + 1;
    ELSIF t2 = 'CK' THEN 
      t = 'K';
      i = i + 1;
    ELSIF p = t THEN 
      t = ''; 
    ELSIF t = 'H' and 'AEIOU' NOT LIKE '%' || p || '%' THEN 
      t = ''; 
    ELSIF (t = 'I' or t = 'Y') and not ('AEOU' like '%' || p || '%') THEN
      t = 'E'; 
    ELSIF t = 'U' and not ('AEOU' like '%' || p || '%') THEN
      t = 'O';     
    ELSIF t = 'V' THEN t = 'B';
    ELSIF t = 'Z' THEN t = 'S';
    END IF;
  
    s = s || t;
    p = t;
  END LOOP;

  i = 0;
  p = ' ';
  a = s;
  s = '';
  LOOP
    EXIT WHEN i > LENGTH(a)-2;
    i = i + 1;
    t = SubString(a,i,1);    
    t2 = SubString(a,i+1,1); 
    t3 = SubString(a,i+2,1); 
        
    IF t = 'C' THEN
       t = CASE WHEN SubString(a,i+1,1) in ('E','I') THEN 'S' 
                       WHEN SubString(a,i+1,1) in ('A','O','U') THEN 'K' ELSE t END;
    ELSIF t || t2 || t3 = 'BLE' THEN 
       t = 'BEL'; -- BLE to BEL 
       i = i + 2; 
   -- ELSIF 'BCDFGHJKLNPQRSTVWXYZ' like '%' || t ||'%' and 'AEIOU' like '%' || t2 || '%' and t3 = 'R' THEN 
   --    t = t || 'R' || t2;  -- [B..Z]ER to [B..Z]RE
   --    i = i + 2; 
    END IF;

    s = s || t;
  END LOOP;
  RETURN trim(lower(s));    
END;
$BODY$ LANGUAGE plpgsql volatile;
 
--------------------------------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION plainto_tsquery_or (query text) 
--------------------------------------------------------------------------------------------------
RETURNS tsquery AS $BODY$
BEGIN  
  --query = Simpleword(query);
  RETURN replace(cast(plainto_tsquery('Simple', query) as text),'&','|');   
END;
$BODY$ Language plpgsql volatile;

--------------------------------------------------------------------------------------------------
CREATE OR REPLACE FUNCTION plainto_tsquery_and (query text) 
--------------------------------------------------------------------------------------------------
RETURNS tsquery AS $BODY$
BEGIN  
  --query = Simpleword(query);
  RETURN plainto_tsquery('Simple', query);   
END;
$BODY$ Language plpgsql volatile;

CREATE OR REPLACE FUNCTION FullNameTFMLS(
  title character varying,
  lname character varying,
  fname character varying,
  mname character varying,
  suffix character varying
) RETURNS character varying AS
$BODY$

BEGIN
  RETURN 
    CASE WHEN TRIM(COALESCE(title,'')) = '' THEN '' ELSE title || ' ' END || 
    CASE WHEN TRIM(COALESCE(fname,'')) = '' THEN '' ELSE fname END || 
    CASE WHEN TRIM(COALESCE(mname,'')) = '' THEN '' ELSE ' ' || LEFT(mname,1) || '.' END || 
    CASE WHEN TRIM(COALESCE(lname,'')) = '' THEN '' ELSE ' ' || lname END || 
    CASE WHEN TRIM(COALESCE(suffix,'')) = '' THEN '' ELSE ' ' || suffix END
  ;
END;
$BODY$
LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION FullName(
  lname character varying,
  fname character varying,
  mname character varying,
  suffix character varying
) RETURNS character varying AS
$BODY$

BEGIN
  RETURN 
    CASE WHEN TRIM(lname,'') = '' THEN '' ELSE ' ' || lname END || 
    CASE WHEN TRIM(COALESCE(fname,'')) = '' THEN '' ELSE ', ' || fname END || 
    CASE WHEN TRIM(COALESCE(suffix,'')) = '' THEN '' ELSE ' ' || suffix END ||
    CASE WHEN TRIM(COALESCE(mname,'')) = '' THEN '' ELSE ' ' || LEFT(mname,1) || '.' END
  ;
END;
$BODY$
LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION FullName(
  lname character varying,
  fname character varying,
  mname character varying
) RETURNS character varying AS
$BODY$

BEGIN
  RETURN 
    CASE WHEN TRIM(lname,'') = '' THEN '' ELSE ' ' || lname END || 
    CASE WHEN TRIM(COALESCE(fname,'')) = '' THEN '' ELSE ', ' || fname END || 
    CASE WHEN TRIM(COALESCE(mname,'')) = '' THEN '' ELSE ' ' || LEFT(mname,1) || '.' END
  ;
END;
$BODY$
LANGUAGE plpgsql;
