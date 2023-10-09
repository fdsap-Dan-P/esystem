---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Actor_Group_Member (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Actor_Group_ID bigint NOT NULL,
  Actor_Group Varchar(200) NOT NULL,
  User_ID bigint NOT NULL,
  Other_Info jsonb NULL,
 
  CONSTRAINT Actor_Group_Member_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Actor_Group_Member_Group FOREIGN KEY (Actor_Group_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Actor_Group_Member_User FOREIGN KEY (User_ID) REFERENCES Users(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxActor_Group_Member_UUID ON public.Actor_Group_Member(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxActor_Group_Member_Unq ON public.Actor_Group_Member(Actor_Group_ID, User_ID);

DROP TRIGGER IF EXISTS trgActor_Group_Member_Ins on Actor_Group_Member;
---------------------------------------------------------------------------
CREATE TRIGGER trgActor_Group_Member_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Actor_Group_Member
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgActor_Group_Member_upd on Actor_Group_Member;
---------------------------------------------------------------------------
CREATE TRIGGER trgActor_Group_Member_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Actor_Group_Member
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgActor_Group_Member_del on Actor_Group_Member;
---------------------------------------------------------------------------
CREATE TRIGGER trgActor_Group_Member_del
---------------------------------------------------------------------------
    AFTER DELETE ON Actor_Group_Member
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
