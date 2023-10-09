
\COPY dat_Area FROM 'NCR4/stg_Area.csv' DELIMITER '|' CSV;
\COPY dat_Unit FROM 'NCR4/stg_Managers.csv' DELIMITER '|' CSV;
\COPY dat_Center FROM 'NCR4/stg_Center.csv' DELIMITER '|' CSV;
\COPY dat_Customer FROM 'NCR4/stg_Customer.csv' DELIMITER '|' CSV;
\COPY dat_lnMaster FROM 'NCR4/stg_lnMaster.csv' DELIMITER '|' CSV;
\COPY dat_saMaster FROM 'NCR4/stg_saMaster.csv' DELIMITER '|' CSV;
\COPY dat_trnMaster FROM 'NCR4/stg_trnMaster.csv' DELIMITER '|' CSV;
\COPY dat_trnMaster FROM 'NCR4/stg_satrnMaster.csv' DELIMITER '|' CSV;
\COPY dat_LoanInst FROM 'NCR4/stg_LoanInst.csv' DELIMITER '|' CSV;
\COPY dat_Mutual_Fund FROM 'NCR4/stg_Mutual_Fund.csv' DELIMITER '|' CSV;


              
                     

delete from dat_Area where brcode in ('N1');
delete from dat_Unit where brcode in ('N1');
delete from dat_Center where brcode in ('N1');
delete from dat_Customer where brcode in ('N1');
delete from dat_lnMaster where brcode in ('N1');
delete from dat_saMaster where brcode in ('N1');
delete from dat_trnMaster where brcode in ('N1');
delete from dat_LoanInst where brcode in ('N1');
delete from dat_lnChrgData where brcode in ('N1');
delete from dat_Mutual_Fund where brcode in ('N1');

'P7','W5','04','C1','L7', 'O1', '2F'

delete from dat_Area ;
delete from dat_Unit ;
delete from dat_Center ;
delete from dat_Customer ;
delete from dat_lnMaster ;
delete from dat_saMaster ;
delete from dat_trnMaster ;
delete from dat_LoanInst ;
delete from dat_lnChrgData ;
delete from dat_Mutual_Fund ;

