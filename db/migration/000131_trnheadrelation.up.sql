---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Trn_Head_Relation (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Trn_Head_ID bigint NOT NULL,
  Related_ID bigint NOT NULL,
  Type_ID bigint NOT NULL,
  Remarks varchar(100) NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT Trn_Head_Relation_pkey   PRIMARY KEY (Trn_Head_ID, Related_ID),
  CONSTRAINT fkTrn_Head_Relation_Head FOREIGN KEY (Trn_Head_ID) REFERENCES Trn_Head(ID),
  CONSTRAINT fkTrn_Head_Relation_Rel  FOREIGN KEY (Related_ID) REFERENCES Trn_Head(ID),
  CONSTRAINT fkTrn_Head_Relation_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxTrn_Head_Relation_UUID ON public.Trn_Head_Relation(UUID);

DROP TRIGGER IF EXISTS trgTrn_Head_RelationIns on Trn_Head_Relation;
---------------------------------------------------------------------------
CREATE TRIGGER trgTrn_Head_RelationIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Trn_Head_Relation
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgTrn_Head_Relationupd on Trn_Head_Relation;
---------------------------------------------------------------------------
CREATE TRIGGER trgTrn_Head_Relationupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Trn_Head_Relation
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgTrn_Head_Relation_del on Trn_Head_Relation;
---------------------------------------------------------------------------
CREATE TRIGGER trgTrn_Head_Relation_del
---------------------------------------------------------------------------
    AFTER DELETE ON Trn_Head_Relation
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

  
  INSERT INTO Trn_Head_Relation(
    Trn_Head_ID, Related_ID, Type_ID, Remarks)
  SELECT 
    h.ID Trn_Head_ID, r.ID Related_ID, typ.ID ActionType, a.Remarks
    
   FROM (Values
      ('26dfab18-f80b-46cf-9c54-be79d4fc5d23'::UUID, '2af90d74-3bee-48c5-8935-443edafb8f5a'::UUID,'Automated', 'Remarks')
      )   
    a(Trn_HeadUUID, RelTrn_HeadUUID, ActionType, Remarks)  

  INNER JOIN Trn_Head h on h.UUID = a.Trn_HeadUUID 
  INNER JOIN Trn_Head r on r.UUID = a.RelTrn_HeadUUID 
  INNER JOIN vwReference typ  on lower(typ.Title) = lower(a.ActionType) and lower(typ.Ref_Type) = 'trnheadaction'

  ON CONFLICT(Trn_Head_ID, Related_ID)
  DO UPDATE SET
    Type_ID = excluded.Type_ID,
    Remarks = excluded.Remarks
  ;   



