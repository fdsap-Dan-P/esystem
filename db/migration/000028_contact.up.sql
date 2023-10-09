----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Contact (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  IIID bigint NOT NULL,
  Series int2 NOT NULL,
  Contact varchar(50) NOT NULL,
  Type_ID bigint NOT NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT Contact_pkey PRIMARY KEY (IIID, Series),
  CONSTRAINT fk_Contact_Identity FOREIGN KEY (IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fk_Contact_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxContact_UUID ON public.Contact(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxContact_Type ON public.Contact(IIID, Type_ID, lower(trim(Contact)));

CREATE INDEX IF NOT EXISTS idxContact_Contact ON public.Contact USING btree (Contact);

DROP TRIGGER IF EXISTS trgContact_Ins on Contact;
---------------------------------------------------------------------------
CREATE TRIGGER trgContact_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Contact
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgContact_upd on Contact;
---------------------------------------------------------------------------
CREATE TRIGGER trgContact_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Contact
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgContact_del on Contact;
---------------------------------------------------------------------------
CREATE TRIGGER trgContact_del
---------------------------------------------------------------------------
    AFTER DELETE ON Contact
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwContact
----------------------------------------------------------------------------------------
AS SELECT
    mr.UUID,
    ii.ID IIID, ii.Alternate_ID,
    ii.Last_Name, ii.First_Name, ii.Middle_Name, ii.BirthDay, ii.Sex,
    
    c.Series, c.Contact,

    typ.ID Type_ID, typ.UUID TypeUUID, typ.Title ContactType,
    
    mr.Mod_Ctr,
    c.Other_Info,
    mr.Created,
    mr.Updated 
   FROM Contact c
   JOIN Main_Record mr ON mr.UUID = c.UUID
   JOIN Identity_Info ii ON c.IIID = ii.ID
   LEFT JOIN Reference typ ON typ.ID = c.Type_ID; 
 
 /*
 INSERT into Contact(IIID, Series, Type_ID, Contact, Other_Info) 
  SELECT 
    ii.ID, a.Series,
    typ.ID Type_ID, a.Contact, cast(a.Other_Info as jsonb) Other_Info
   FROM (Values
      ('100',1,'Cellphone','+63 998 851 3220',NULL),
      ('100',2,'eMail','erick1421@gmail.com',NULL)
         )   
    a(Alternate_ID, Series, Contact_Type, Contact, Other_Info)  
      
  LEFT JOIN Identity_Info ii   on ii.Alternate_ID = a.Alternate_ID
  LEFT JOIN Contact c          on c.IIID = ii.ID and c.Series = a.Series
  LEFT JOIN vwReference typ    on lower(typ.Title) = lower(a.Contact_Type) 
     and typ.Ref_Type = 'ContactType'  
  ON CONFLICT(IIID, Series) DO UPDATE SET  
    Type_ID = excluded.Type_ID,   
    Contact = excluded.Contact
  ;  
*/