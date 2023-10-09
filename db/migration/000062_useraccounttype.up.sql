----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.User_Account_Type (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  User_ID bigint NOT NULL,
  Account_Type_ID bigint NOT NULL,
  Allow boolean NOT NULL,
  Other_Info jsonb NULL,

  CONSTRAINT User_Account_Type_pkey PRIMARY KEY (User_ID, Account_Type_ID),  
  CONSTRAINT fkUserAcc_Type FOREIGN KEY (Account_Type_ID) REFERENCES Account_Type(ID),
  CONSTRAINT fkUserAcc_User FOREIGN KEY (User_ID) REFERENCES Users(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxUser_Account_Type_UUID ON public.User_Account_Type(UUID);

DROP TRIGGER IF EXISTS trgUser_Account_TypeIns on User_Account_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgUser_Account_TypeIns
---------------------------------------------------------------------------
    BEFORE INSERT ON User_Account_Type
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
  
DROP TRIGGER IF EXISTS trgUser_Account_Typeupd on User_Account_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgUser_Account_Typeupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON User_Account_Type
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgUser_Account_Type_del on User_Account_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgUser_Account_Type_del
---------------------------------------------------------------------------
    AFTER DELETE ON User_Account_Type
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

    
----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwUsersAcccountType
----------------------------------------------------------------------------------------
  AS SELECT 
    mr.UUID, u.ID User_ID, u.Login_Name,

    y.ID Account_Type_ID, y.UUID Account_TypeUUID, y.Account_Type,
        
    uat.Allow,
   
    mr.Mod_Ctr,
    uat.Other_Info,
    mr.Created,
    mr.Updated 
   FROM User_Account_Type uat
   INNER JOIN Main_Record mr on mr.UUID = uat.UUID
   LEFT JOIN Account_Type y ON y.ID = uat.Account_Type_ID
   LEFT JOIN Users u ON u.ID = uat.User_ID
    ;

  INSERT INTO User_Account_Type
    (User_ID, Account_Type_ID, Allow)
  SELECT 
     u.ID User_ID, y.ID Account_Type_ID, a.Allow
  FROM
   (Values
    ('erick1421@gmail.com', 'Sikap 1', true))
     a(Login_Name, Account_Type, Allow)  
  INNER JOIN Account_Type y on y.Account_Type = a.Account_Type
  INNER JOIN Users    u on u.Login_Name = a.Login_Name
  
  ON CONFLICT(User_ID, Account_Type_ID) DO UPDATE SET
    Allow = EXCLUDED.Allow,
    Other_Info = EXCLUDED.Other_Info
  ;
