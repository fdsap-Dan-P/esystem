----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.User_Office (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  User_ID bigint NOT NULL,
  Office_ID bigint NOT NULL,
  Allow boolean NOT NULL,
  Other_Info jsonb NULL,

  CONSTRAINT User_Office_pkey PRIMARY KEY (User_ID, Office_ID ),
  CONSTRAINT fkUser_Office FOREIGN KEY (Office_ID ) REFERENCES Office(ID),
  CONSTRAINT fkUserUser FOREIGN KEY (User_ID) REFERENCES Users(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxUser_Office_UUID ON public.User_Office(UUID);

DROP TRIGGER IF EXISTS trgUser_OfficeIns on User_Office;
---------------------------------------------------------------------------
CREATE TRIGGER trgUser_OfficeIns
---------------------------------------------------------------------------
    BEFORE INSERT ON User_Office
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
  
DROP TRIGGER IF EXISTS trgUser_Officeupd on User_Office;
---------------------------------------------------------------------------
CREATE TRIGGER trgUser_Officeupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON User_Office
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgUser_Office_del on User_Office;
---------------------------------------------------------------------------
CREATE TRIGGER trgUser_Office_del
---------------------------------------------------------------------------
    AFTER DELETE ON User_Office
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

    
----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwUser_Office
----------------------------------------------------------------------------------------
  AS SELECT 
    mr.UUID,
    u.ID User_ID, u.Login_Name,

    y.ID Office_ID , y.UUID OfficeUUID, y.Office_Name,
        
    uat.Allow,
   
    mr.Mod_Ctr,
    uat.Other_Info,
    mr.Created,
    mr.Updated 
   FROM User_Office uat
   INNER JOIN Main_Record mr on mr.UUID = uat.UUID
   LEFT JOIN Office y ON y.ID = uat.Office_ID 
   LEFT JOIN Users u ON u.ID = uat.User_ID
    ;

  INSERT INTO User_Office
    (User_ID, Office_ID , Allow)
  SELECT 
     u.ID User_ID, y.ID Office_ID , a.Allow
  FROM
   (Values
    ('erick1421@gmail.com', 'United Church of Christ in the Philippines', true))
     a(Login_Name, Office_Name, Allow)  
  INNER JOIN Office y on y.Office_Name = a.Office_Name
  INNER JOIN Users    u on u.Login_Name = a.Login_Name
  
  ON CONFLICT(User_ID, Office_ID ) DO UPDATE SET
    Allow = EXCLUDED.Allow
  ;
