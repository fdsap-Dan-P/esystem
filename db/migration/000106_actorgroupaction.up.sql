
---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Actor_Group_Action (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Actor_Group_ID bigint NOT NULL,
  Actor_Group Varchar(200) NOT NULL,
  Ticket_Type_Action_ID bigint NOT NULL,
  Ticket_Type Varchar(200) NOT NULL,
  Action_Desc Varchar(200) NOT NULL,
  Other_Info jsonb NULL,
 
  CONSTRAINT Actor_Group_Action_pkey PRIMARY KEY (UUID),
  CONSTRAINT fk_Actor_Group_Action_Type FOREIGN KEY (Actor_Group_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Actor_Group_Action_Item FOREIGN KEY (Ticket_Type_Action_ID) REFERENCES Ticket_Type_Action(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxActor_Group_Action_Unq ON public.Actor_Group_Action(Actor_Group_ID, Ticket_Type_Action_ID);

DROP TRIGGER IF EXISTS trgActor_Group_Action_Ins on Actor_Group_Action;
---------------------------------------------------------------------------
CREATE TRIGGER trgActor_Group_Action_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Actor_Group_Action
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgActor_Group_Action_upd on Actor_Group_Action;
---------------------------------------------------------------------------
CREATE TRIGGER trgActor_Group_Action_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Actor_Group_Action
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgActor_Group_Action_del on Actor_Group_Action;
---------------------------------------------------------------------------
CREATE TRIGGER trgActor_Group_Action_del
---------------------------------------------------------------------------
    AFTER DELETE ON Actor_Group_Action
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
