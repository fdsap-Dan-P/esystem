
with trn as
 (SELECT account_id, 
    sum(case when trn_prin > 0 THEN  trn_prin ELSE 0 END) trn_prin_Credit, 
    sum(case when trn_prin < 0 THEN  -trn_prin ELSE 0 END) trn_prin_Debit, 
    sum(case when trn_int  > 0 THEN   trn_int ELSE 0 END) trn_int_Credit, 
    sum(case when trn_int  < 0 THEN  -trn_int ELSE 0 END) trn_int_Debit

  FROM Account_Tran 
  WHERE trn_type_code not in (3400, 3100)
  GROUP BY account_id)

  UPDATE Account SET 
  Debit  = trn_prin_Debit,
  Credit = trn_prin_Credit
FROM trn  
WHERE Account.id = trn.Account_Id ;



with trn as
 (SELECT account_id, 
    sum(case when trn_prin > 0 THEN  trn_prin ELSE 0 END) trn_prin_Credit, 
    sum(case when trn_prin < 0 THEN  -trn_prin ELSE 0 END) trn_prin_Debit, 
    sum(case when trn_int  > 0 THEN   trn_int ELSE 0 END) trn_int_Credit, 
    sum(case when trn_int  < 0 THEN  -trn_int ELSE 0 END) trn_int_Debit
  FROM Account_Tran 
  WHERE trn_type_code not in (3400, 3100)
  GROUP BY account_id)

  UPDATE Account_Interest SET 
  Debit  = trn_int_Debit,
  Credit = trn_int_Credit
FROM trn  
WHERE Account_Interest.Account_Id = trn.Account_Id 

