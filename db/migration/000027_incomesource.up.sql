----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Income_Source (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  IIID bigint NOT NULL,
  Series int2 NOT NULL,
  Source VarChar(200) NOT NULL,
  Type_ID bigint NOT NULL,
  Min_Income numeric(16,6) NOT NULL,
  Max_Income numeric(16,6) NOT NULL,
  Remarks Text,
  Other_Info jsonb NULL,
  
  CONSTRAINT Income_Source_pkey PRIMARY KEY (IIID, Series),
  CONSTRAINT fkIncome_Source_Identity FOREIGN KEY (IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fkIncome_Source_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID)
);

CREATE UNIQUE INDEX IF NOT EXISTS idxIncome_Source_UUID ON public.Income_Source(UUID);

DROP TRIGGER IF EXISTS trgIncome_SourceIns on Income_Source;
---------------------------------------------------------------------------
CREATE TRIGGER trgIncome_SourceIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Income_Source
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgIncome_Sourceupd on Income_Source;
---------------------------------------------------------------------------
CREATE TRIGGER trgIncome_Sourceupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Income_Source
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgIncome_Source_del on Income_Source;
---------------------------------------------------------------------------
CREATE TRIGGER trgIncome_Source_del
---------------------------------------------------------------------------
    AFTER DELETE ON Income_Source
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

/*
INSERT into Income_Source(
      IIID, Series, Source, Type_ID,
      Min_Income, Max_Income, Remarks) 
  
  SELECT 
      i.ID IIID, Series, Source, typ.ID Type_ID,
      Min_Income, Max_Income, Remarks
  FROM (Values
      ('100',1,'Private Company','Salary',8000,10000,'Primary Income')
      )   
    a(
      Alternate_ID, Series, Source, SourceType,
      Min_Income, Max_Income, Remarks
      )   
  LEFT JOIN Identity_Info i on i.Alternate_ID = a.Alternate_ID 
  LEFT JOIN vwReference typ on lower(typ.Ref_Type) = lower('SourceofIncome') and lower(typ.Title) = lower(a.SourceType)

  ON CONFLICT(IIID, Series) DO UPDATE SET
    Source = excluded.Source, 
    Type_ID = excluded.Type_ID, 
    Min_Income = excluded.Min_Income,
    Max_Income = excluded.Max_Income,
    Remarks = excluded.Remarks
    ;
*/