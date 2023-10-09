---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Server (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Code varchar(30) NOT NULL,
  Connectivity int2 NOT NULL, -- 0:local, 1:network, 2:service
  Net_Address varchar(100) NOT NULL,
  Certificate text,
  HomePath varchar(1000) NOT NULL DEFAULT '',
  Description varchar(1000) NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT Server_pkey PRIMARY KEY (ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxServer_UUID ON public.Server(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxServer_Code ON public.Server(lower(Code));

DROP TRIGGER IF EXISTS trgServer_Ins on Server;
---------------------------------------------------------------------------
CREATE TRIGGER trgServer_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Server
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgServer_upd on Server;
---------------------------------------------------------------------------
CREATE TRIGGER trgServer_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Server
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgServer_del on Server;
---------------------------------------------------------------------------
CREATE TRIGGER trgServer_del
---------------------------------------------------------------------------
    AFTER DELETE ON Server
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

INSERT INTO Server(
    UUID, Code, Connectivity, Net_Address, Certificate, Description)
  SELECT 
    a.UUID, a.Code, 
    CASE a.Connectivity WHEN 'local' THEN 0 WHEN 'network' THEN 1 WHEN 'service' THEN 2 ELSE 0 END Connectivity, 
    a.Net_Address, a.Certificate, a.Description
    
   FROM (Values
      ('bb295886-65ba-435d-a937-18fb9bd3204b'::UUID,'Local','Local','localhost',NULL,'Local Server')
      )   
    a( UUID, Code, Connectivity, Net_Address, Certificate, Description)  

  ON CONFLICT(lower(Code))
  DO UPDATE SET
    Connectivity = excluded.Connectivity,
    Net_Address = excluded.Net_Address,
    Certificate = excluded.Certificate,
    Description = excluded.Description
  ;  
