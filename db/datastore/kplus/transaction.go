package db

import (
	"context"
	"fmt"
	"log"
	"simplebank/util"
	"sort"
	"time"
)

const kPLUStransactionSQL = `-- name: kPLUStransactionSQL :one
WITH typ as
  (SELECT trntype
   FROM 
    (VALUES
      (1),(2),(3),(5),(13),(210),(214),(227),(231),(233),
      (234),(236),(238),(506),(560),(3001),(3097),(3098),(3099),(3899),(3201),(3202) )
    typ(trntype)),
acc as
  (SELECT
     acc.ID Account_Id, acc.acc, acc.principal + acc.debit - acc.credit bal_Prin, 
     COALESCE(ai.interest,0)-COALESCE(ai.credit,0)+COALESCE(ai.debit,0) Bal_Int,
	 Normal_Balance
   FROM account acc  
   INNER JOIN Account_Type typ on acc.Account_Type_Id = typ.Id
   LEFT JOIN account_interest ai on acc.id = ai.account_id 
   WHERE acc = $1
 ),
bal as
( SELECT acc.account_id, Sum(trn_prin) Paid_Prin,sum(trn_int) Paid_Int
  FROM account_tran at2
  INNER JOIN trn_head th on th.id = at2.trn_head_id 
  INNER JOIN typ on typ.trntype = at2.trn_type_code, acc
  WHERE at2.account_id = acc.account_id and at2.trn_type_code = typ.trntype  
     and th.trn_date > $3
  GROUP BY acc.account_id
) 
SELECT 
  acc.Account_id, acc.acc,  to_char(th.trn_date, 'mm-dd-yyyy') trndate, 
  t.value_date ValueDate, th.id trn_Head_id, t.series trn,
  t.alternate_key, t.trn_prin prin, t.trn_int intr,  
  t.trn_prin + t.trn_int trnamount, acc.bal_Prin, acc.Bal_Int, acc.bal_Prin + acc.Bal_Int Balance, 
  COALESCE(bal.Paid_Prin,0) Paid_Prin, COALESCE(bal.Paid_Int,0) Paid_Int,
  th.particular particulars, COALESCE(r.Title,'') trnType, u.login_name username,
  typ.TrnType is NOT NULL isFinancial, acc.Normal_Balance
FROM
 Account_Tran t
INNER JOIN trn_head th on t.trn_head_id = th.id
LEFT JOIN typ on typ.trntype = t.trn_type_code 
LEFT JOIN Reference r on r.Code = t.trn_type_code  and lower(r.Ref_Type) = 'trntype'
INNER JOIN users u on th.user_id = u.id
LEFT JOIN bal on bal.account_id = t.account_id, acc
WHERE t.account_id = acc.Account_id and th.trn_date Between $2 and $3
ORDER BY trndate desc, trn desc
`

func populateKPLUStransaction(q *QueriesKPlus, ctx context.Context, sql string, arg TransactionParams) ([]TransactionHistory, error) {
	rows, err := q.db.QueryContext(ctx, sql, arg.Acc, arg.DateFrom, arg.DateTo)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []TransactionHistory{}
	for rows.Next() {
		var i TransactionHistory
		err := rows.Scan(
			&i.AccountId,
			&i.Acc,
			&i.TrnDate,
			&i.ValueDate,
			&i.TrnHeadId,
			&i.Trn,
			&i.AlternateKey,
			&i.Prin,
			&i.Intr,
			&i.TrnAmount,
			&i.BalPrin,
			&i.BalInt,
			&i.Balance,
			&i.PaidPrin,
			&i.PaidInt,
			&i.Particulars,
			&i.TrnType,
			&i.Username,
			&i.IsFinancial,
			&i.NormalBalance,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	sort.Sort(ByTrnDateTrn(items))

	balPrin := float64(0)
	balInt := float64(0)
	norBal := true

	for i, d := range items {
		if i == 0 {
			norBal = items[0].NormalBalance
			if norBal {
				balPrin = util.RoundTo(items[0].BalPrin+items[0].PaidPrin, 2)
				balInt = util.RoundTo(items[0].BalInt+items[0].PaidInt, 2)
			} else {
				balPrin = util.RoundTo(-items[0].BalPrin-items[0].PaidPrin, 2)
				balInt = util.RoundTo(-items[0].BalInt-items[0].PaidInt, 2)
			}
		}
		items[i].BalPrin = util.RoundTo(balPrin, 2)
		items[i].BalInt = util.RoundTo(balInt, 2)
		items[i].Balance = util.RoundTo(balPrin+balInt, 2)
		log.Printf("items[i].Balance: %v", items[i].Balance)

		if norBal && items[i].Prin != 0 {
			items[i].Prin = -items[i].Prin
		}
		if norBal && items[i].Intr != 0 {
			items[i].Intr = -items[i].Intr
		}

		balPrin = balPrin + d.Prin
		balInt = balInt + d.Intr
	}
	return items, nil
}

func (q *QueriesKPlus) Transaction(ctx context.Context, arg TransactionParams) ([]TransactionHistory, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d",
			kPLUStransactionSQL, arg.Limit, arg.Offset)
	} else {
		sql = kPLUStransactionSQL
	}
	return populateKPLUStransaction(q, ctx, sql, arg)
}

type TransactionParams struct {
	Acc      string    `json:"acc"`
	DateFrom time.Time `json:"date_from"`
	DateTo   time.Time `json:"date_to"`
	Limit    int32     `json:"limit"`
	Offset   int32     `json:"offset"`
}

// GetTransactionHistory(ctx context.Context, in *KPLUSGetTransactionHistoryRequest, opts ...grpc.CallOption) (*KPLUSGetTransactionHistoryResponse, error)
// GenerateColShtperCID(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSGenerateColShtperCIDResponse, error)
// K2CCallBackRef(ctx context.Context, in *KPLUSCallBackRefRequest, opts ...grpc.CallOption) (*KPLUSResponse, error)
// GetReferences(ctx context.Context, in *KPLUSGetReferencesRequest, opts ...grpc.CallOption) (*KPLUSGetReferencesResponse, error)
// MultiplePayment(ctx context.Context, in *KPLUSMultiplePaymentRequest, opts ...grpc.CallOption) (*KPLUSResponse, error)
// SearchLoanList(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSSearchLoanListResponse, error)
// LoanInfo(ctx context.Context, in *KPLUSAccRequest, opts ...grpc.CallOption) (*KPLUSLoanInfoResponse, error)
// GetSavingForSuperApp(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSGetSavingResponse, error)
// FundTransferRequest(ctx context.Context, in *KPLUSFundTransferRequest, opts ...grpc.CallOption) (*KPLUSResponse, error)
