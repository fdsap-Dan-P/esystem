DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'social_provider_type') THEN
        CREATE TYPE Social_Provider_Type AS ENUM
        (
           'facebook','twitter', 'google'
        );
    END IF;
END$$;

----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Social_Media_Credential (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  User_ID bigint NOT NULL,
  Provider_Key VarChar(128) NOT NULL,
  Provider_Type Social_Provider_Type NOT NULL,
  Other_Info jsonb NULL,

  CONSTRAINT Social_Media_Credential_pkey PRIMARY KEY (UUID),
  CONSTRAINT fkSocial_Media_Credential_user FOREIGN KEY (User_ID) REFERENCES Users(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxSocial_Media_Credential_UUID ON public.Social_Media_Credential(UUID);

DROP TRIGGER IF EXISTS trgSocial_Media_CredentialIns on Social_Media_Credential;
---------------------------------------------------------------------------
CREATE TRIGGER trgSocial_Media_CredentialIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Social_Media_Credential
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
  
DROP TRIGGER IF EXISTS trgSocial_Media_Credentialupd on Social_Media_Credential;
---------------------------------------------------------------------------
CREATE TRIGGER trgSocial_Media_Credentialupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Social_Media_Credential
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgSocial_Media_Credential_del on Social_Media_Credential;
---------------------------------------------------------------------------
CREATE TRIGGER trgSocial_Media_Credential_del
---------------------------------------------------------------------------
    AFTER DELETE ON Social_Media_Credential
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwSocial_Media_Credential
----------------------------------------------------------------------------------------
  AS SELECT 
    mr.UUID,
    User_ID, u.Login_Name, Provider_Key, Provider_Type, 
    ii.Last_Name, ii.First_Name, ii.Middle_Name,
    mr.Mod_Ctr,
    uat.Other_Info,
    mr.Created,
    mr.Updated 
   FROM Social_Media_Credential uat
   INNER JOIN Main_Record mr ON mr.UUID = uat.UUID
   INNER JOIN Users u ON u.ID = uat.User_ID
   INNER JOIN Identity_Info ii on ii.ID = u.IIID
;
