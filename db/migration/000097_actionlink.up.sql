---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Action_Link (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Event_Name Varchar(200) NOT NULL,
  Type_ID bigint NOT NULL,
  End_Point_Call text NULL,
  Server_ID bigint NULL,
  Other_Info jsonb NULL,
 
  CONSTRAINT Action_Link_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Action_Link_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Action_Link_Server FOREIGN KEY (Server_ID) REFERENCES Server(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAction_Link_UUID ON public.Action_Link(UUID);

DROP TRIGGER IF EXISTS trgAction_Link_Ins on Action_Link;
---------------------------------------------------------------------------
CREATE TRIGGER trgAction_Link_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Action_Link
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAction_Link_upd on Action_Link;
---------------------------------------------------------------------------
CREATE TRIGGER trgAction_Link_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Action_Link
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAction_Link_del on Action_Link;
---------------------------------------------------------------------------
CREATE TRIGGER trgAction_Link_del
---------------------------------------------------------------------------
    AFTER DELETE ON Action_Link
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
  
