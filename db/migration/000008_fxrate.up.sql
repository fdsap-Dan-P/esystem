----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.fxRate (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Buy_Rate numeric(16,10) NOT NULL DEFAULT 0,
  Cutof_Date Date NOT NULL,
  Sell_Rate numeric(16,10) NOT NULL DEFAULT 0,
  Base_Currency varchar(3) NOT NULL,
  Currency varchar(3) NOT NULL,

  Other_Info jsonb NULL,

  CONSTRAINT fxRate_pkey PRIMARY KEY (Base_Currency, Currency, Cutof_Date)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxfxRate_UUID ON public.fxRate(UUID);

DROP TRIGGER IF EXISTS trgFXRate_Ins on FXRate;
---------------------------------------------------------------------------
CREATE TRIGGER trgFXRate_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON FXRate
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();

DROP TRIGGER IF EXISTS trgFXRate_upd on FXRate;
---------------------------------------------------------------------------
CREATE TRIGGER trgFXRate_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON FXRate
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
  EXECUTE PROCEDURE trgGenericUpdate();

  DROP TRIGGER IF EXISTS trgFXRate_del on FXRate;
---------------------------------------------------------------------------
CREATE TRIGGER trgFXRate_del
---------------------------------------------------------------------------
    AFTER DELETE ON FXRate
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwfxRate
----------------------------------------------------------------------------------------
AS  SELECT 
    mr.UUID,
    r.Base_Currency,
    r.Currency,
    r.Cutof_Date,
    r.Buy_Rate,
    r.Sell_Rate,
  
    mr.Mod_Ctr,
    r.Other_Info,
    mr.Created,
    mr.Updated 
   FROM fxRate r
     JOIN Main_Record  mr on mr.UUID = r.UUID;
   
----------------------------------------------------------------------------------------
   INSERT INTO 
     FXRate(Base_Currency,   Currency,   Cutof_Date,   Buy_Rate,   Sell_Rate)   
   SELECT 'PHP' Base_Currency,   'USD' Currency,   Date '2019-01-15' Cutof_Date,   
     50 Buy_Rate,   52 Sell_Rate
   ON CONFLICT(Base_Currency,   Currency,   Cutof_Date)
   DO UPDATE SET Buy_Rate = EXCLUDED.Buy_Rate, Sell_Rate = EXCLUDED.Sell_Rate;

