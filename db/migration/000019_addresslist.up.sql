CREATE TABLE IF NOT EXISTS public.Address_List (
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  IIID bigint NOT NULL,
  Series int2 NOT NULL,
  Detail varchar(150) NULL,
  URL varchar(200) NULL,
  Type_ID bigint NOT NULL,
  Geography_ID bigint NULL,
  Other_Info jsonb NULL,

  CONSTRAINT Address_List_pkey PRIMARY KEY (IIID, Series),
  CONSTRAINT fk_Address_Geography FOREIGN KEY (Geography_ID) REFERENCES Geography(ID),
  CONSTRAINT fk_AddressIdentity_Info FOREIGN KEY (IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fk_Address_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID)
);

CREATE UNIQUE INDEX IF NOT EXISTS idxAddress_List_UUID ON public.Address_List(UUID);
-- drop table Address_List cascade
DROP TRIGGER IF EXISTS trgAddress_List_Ins on Address_List;
---------------------------------------------------------------------------
CREATE TRIGGER trgAddress_List_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Address_List
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAddress_List_upd on Address_List;
---------------------------------------------------------------------------
CREATE TRIGGER trgAddress_List_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Address_List
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAddress_List_del on Address_List;
---------------------------------------------------------------------------
CREATE TRIGGER trgAddress_List_del
---------------------------------------------------------------------------
    AFTER DELETE ON Address_List
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


