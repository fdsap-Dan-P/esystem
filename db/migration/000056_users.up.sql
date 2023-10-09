
----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Users (
----------------------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  IIID bigint NOT NULL,
  Login_Name varchar(100) NOT NULL,
  Display_Name varchar(100) NULL,
  Access_Role_ID bigint NOT NULL,
  Status_Code integer NOT NULL,
  Date_Given timestamptz NULL,
  Date_Expired timestamptz NULL,
  Date_Locked timestamptz NULL,
  Password_Changed_At timestamptz NULL,
  Hashed_Password bytea NOT NULL,
  Attempt int2 NOT NULL,
  Isloggedin  bool NULL,
  Thumbnail bytea NULL,
  Other_Info jsonb NULL,
 
  CONSTRAINT UserLogin_pkey PRIMARY KEY (ID),
  CONSTRAINT fkUserLoginIdentity_Info FOREIGN KEY (IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fkUserLoginrole FOREIGN KEY (Access_Role_ID) REFERENCES Access_Role(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxUsers_UUID ON public.Users(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxUserLogin_Name ON public.Users(lower(Login_Name));

DROP TRIGGER IF EXISTS trgUsersIns on Users;
---------------------------------------------------------------------------
CREATE TRIGGER trgUsersIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Users
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgUsersupd on Users;
---------------------------------------------------------------------------
CREATE TRIGGER trgUsersupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Users
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgUsers_del on Users;
---------------------------------------------------------------------------
CREATE TRIGGER trgUsers_del
---------------------------------------------------------------------------
    AFTER DELETE ON Users
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwUsers
----------------------------------------------------------------------------------------
AS SELECT 
    usr.ID, mr.UUID,
    p.IIID IIID,
    p.Title,    
    p.Last_Name,
    p.First_Name,
    p.Middle_Name,
    p.Mother_Maiden_Name,
    p.Birthday,
    p.Sex,
    
    p.gender_id, p.Gender,    
    
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

    usr.Login_Name,
    
    usr.Status_Code, sta.ID Status_ID, sta.UUID StatusUUID, sta.Title Statusdesc,    
    ar.ID Access_Role_ID, ar.Access_Name,
    
    usr.Date_Given,
    usr.Date_Expired,
    usr.Date_Locked,
    usr.Password_Changed_At,
    usr.Hashed_Password,
    usr.Attempt,
    usr.isloggedin,
    
    mr.Mod_Ctr,
    usr.Other_Info,
    mr.Created,
    mr.Updated 
    
  FROM Users usr
  INNER JOIN Main_Record mr on mr.UUID = usr.UUID  
  LEFT JOIN vwperson p ON p.IIID = usr.IIID
  LEFT JOIN Access_Role ar ON ar.ID = usr.Access_Role_ID
  LEFT JOIN Reference sta ON sta.Code = usr.Status_Code and lower(sta.Ref_Type) = 'userstatus'
  ;

----------------------------------------------------------------------------------------
  INSERT into Users(
    IIID, Login_Name, Access_Role_ID, Status_Code,
    Date_Given, Date_Expired, Date_Locked, Password_Changed_At, Hashed_Password, Attempt, isLoggedin,
    Other_Info)
  SELECT 
    ii.ID IIID, a.Login_Name, ar.ID Access_Role_ID, stat.Code Status_Code, 
    a.Date_Given, a.Date_Expired, a.Date_Locked, a.Date_Given, decode('01', 'hex'), a.Attempt, a.isLoggedin,
    NULL Other_Info
  FROM (Values
      ('100','erick1421@gmail.com','Bookkeeper','Active','01/01/2020'::date,'01/01/2030'::date,'01/01/2020'::date,'Password',0,FALSE),
      ('101','olive.mercado0609@gmail.com','Bookkeeper','Active','01/01/2020'::date,'01/01/2030'::date,'01/01/2020'::date,'Password',0,FALSE)
    )   
  a(Alternate_ID, Login_Name, Access_Role, Status,
    Date_Given, Date_Expired, Date_Locked, Hashed_Password, Attempt, isLoggedin)  
  
  LEFT JOIN Identity_Info ii on ii.Alternate_ID = a.Alternate_ID
  LEFT JOIN Access_Role ar on lower(ar.Access_Name) = lower(a.Access_Role)
  LEFT JOIN Reference stat  on lower(stat.Title) = lower(a.Status) and lower(stat.Ref_Type) = 'userstatus'    
  ON CONFLICT(lower(Login_Name))   
  DO UPDATE SET
    IIID            = EXCLUDED.IIID, 
    Access_Role_ID  = EXCLUDED.Access_Role_ID, 
    Status_Code     = EXCLUDED.Status_Code, 
    Date_Given      = EXCLUDED.Date_Given,
    Date_Expired    = EXCLUDED.Date_Expired,   
    Date_Locked     = EXCLUDED.Date_Locked, 
    Password_Changed_At = EXCLUDED.Password_Changed_At,
    Hashed_Password = EXCLUDED.Hashed_Password,
    Attempt         = EXCLUDED.Attempt, 
    isloggedin      = EXCLUDED.isloggedin
   ;

