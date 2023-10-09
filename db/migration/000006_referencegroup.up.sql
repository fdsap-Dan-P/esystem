---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.reference_group (
---------------------------------------------------------------------------
  Group_Id bigint NOT NULL,
  Ref_Id bigint NOT NULL,
  Title varchar(200) NOT NULL,
  Active Boolean NOT NULL,
  CONSTRAINT referencegroup_pkey PRIMARY KEY (Group_Id, Ref_Id),
  CONSTRAINT fk_ReferenceGroup_GroupId FOREIGN KEY (Group_Id) REFERENCES Reference(ID),
  CONSTRAINT fk_ReferenceGroup_Reference FOREIGN KEY (Ref_Id) REFERENCES Reference(ID)
);

DROP TRIGGER IF EXISTS trgreference_group_Ins on reference_group;
---------------------------------------------------------------------------
CREATE TRIGGER trgreference_group_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON reference_group
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();

DROP TRIGGER IF EXISTS trgreference_group_upd on reference_group;
---------------------------------------------------------------------------
CREATE TRIGGER trgreference_group_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON reference_group
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgreference_group_del on reference_group;
---------------------------------------------------------------------------
CREATE TRIGGER trgreference_group_del
---------------------------------------------------------------------------
    AFTER DELETE ON reference_group
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

 