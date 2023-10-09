----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.User_Object (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  User_ID bigint NOT NULL,
  Object_ID bigint NOT NULL,
  Allow boolean NOT NULL,
  Other_Info jsonb NULL,

  CONSTRAINT User_Object_pkey PRIMARY KEY (User_ID, Object_ID),  
  CONSTRAINT fkUser_Object FOREIGN KEY (Object_ID) REFERENCES Reference(ID),
  CONSTRAINT fkUserUser FOREIGN KEY (User_ID) REFERENCES Users(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxUser_Object_UUID ON public.User_Object(UUID);

DROP TRIGGER IF EXISTS trgUser_ObjectIns on User_Object;
---------------------------------------------------------------------------
CREATE TRIGGER trgUser_ObjectIns
---------------------------------------------------------------------------
    BEFORE INSERT ON User_Object
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
  
DROP TRIGGER IF EXISTS trgUser_Objectupd on User_Object;
---------------------------------------------------------------------------
CREATE TRIGGER trgUser_Objectupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON User_Object
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgUser_Object_del on User_Object;
---------------------------------------------------------------------------
CREATE TRIGGER trgUser_Object_del
---------------------------------------------------------------------------
    AFTER DELETE ON User_Object
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

    
----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwUser_Object
----------------------------------------------------------------------------------------
  AS SELECT 
    mr.UUID,
    u.ID User_ID, u.Login_Name,

    y.ID Object_ID, y.UUID ObjectUUID, y.Title ObjectName,
        
    uat.Allow,
   
    mr.Mod_Ctr,
    uat.Other_Info,
    mr.Created,
    mr.Updated 
   FROM User_Object uat
   INNER JOIN Main_Record mr on mr.UUID = uat.UUID
   LEFT JOIN Reference y ON y.ID = uat.Object_ID
   LEFT JOIN Users u ON u.ID = uat.User_ID
    ;

  INSERT INTO User_Object
    (User_ID, Object_ID, Allow)
  SELECT 
     u.ID User_ID, y.ID Object_ID, a.Allow
  FROM
   (Values
    ('erick1421@gmail.com', 'File', true))
     a(Login_Name, ObjectName, Allow)  
  INNER JOIN vwReference y on y.Title = a.ObjectName and y.Ref_Type= 'Access_Object'
  INNER JOIN Users    u on u.Login_Name = a.Login_Name
  
  ON CONFLICT(User_ID, Object_ID) DO UPDATE SET
    Allow = EXCLUDED.Allow,
    Other_Info = EXCLUDED.Other_Info
  ;
