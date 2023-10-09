CREATE OR REPLACE FUNCTION public.loaddata(tablename text, filepathname text, brcode character varying)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
BEGIN    
   EXECUTE format('DELETE FROM %s WHERE BrCode = %L ;', tableName, brCode);  
   EXECUTE format('COPY %s FROM %L DELIMITER %L CSV ;', tableName, filepathname, '|');  
END
$function$
;

CREATE OR REPLACE FUNCTION public.loaddata(tablename text, filepathname text)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
BEGIN    
   EXECUTE format('COPY %s FROM %L DELIMITER %L CSV HEADER;', tableName, filepathname, '|');  
END
$function$
;


CREATE or replace FUNCTION public.temp_runner()
  RETURNS bigint AS
$func$
DECLARE
   lastAcc bigint;
BEGIN
   SELECT RIGHT(param_value,12) into lastAcc from m_param  WHERE param_name='ACC_SERIAL';
   EXECUTE format('CREATE SEQUENCE IF NOT EXISTS seq_acc START %s', lastAcc+1);

   SELECT param_value into lastAcc from m_param  WHERE param_name='EPN_SERIAL';
   EXECUTE format('CREATE SEQUENCE IF NOT EXISTS seq_epn START %s', lastAcc+1);
   RETURN lastAcc;
END
$func$ LANGUAGE plpgsql;


select f_checkdigit(param_value || to_char(nextval('seq_acc'), 'fm000000000000') ) ""accountNumber"" from m_param where param_name = 'ACC_PREFIX';