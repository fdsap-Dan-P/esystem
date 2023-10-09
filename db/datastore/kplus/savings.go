package db

import (
	"context"
	"fmt"
)

const kPLUSCustSavingsListSQL = `-- name: kPLUSCustSavingsListSQL :one
SELECT * FROM
(SELECT 
  ids.id_number INAIIID, cus.cid,  a.acc, typ.code acctType, typ.account_type accDesc, a.open_date dopen, 
  stat.title statusDesc,  a.balance balance, stat.code status
FROM Account a
INNER JOIN customer cus  on cus.id = a.customer_id 
INNER JOIN account_type typ  on a.account_type_id = typ.id 
INNER JOIN reference stat on stat.code = a.status_code and lower(stat.ref_type) = 'accountstatus'
INNER JOIN product p on typ.product_id = p.id 
INNER JOIN IDs on ids.iiid  = cus.IIID
INNER JOIN Reference idType on idType.id = ids.Type_id and lower(idType.title) = 'inai-iiid'
WHERE lower(p.product_name) = 'savings' and abs(balance) > 0) d
`

func populateKPLUSCustSavingsList(q *QueriesKPlus, ctx context.Context, sql string) ([]Savings, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []Savings{}
	for rows.Next() {
		var i Savings
		err := rows.Scan(
			&i.INAIIID,
			&i.Cid,
			&i.Acc,
			&i.AcctType,
			&i.AccDesc,
			&i.Dopen,
			&i.StatusDesc,
			&i.Balance,
			&i.Status,
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
	return items, nil
}

func (q *QueriesKPlus) SavingsList(ctx context.Context, arg SavingsListParams) ([]Savings, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE lower(trim(INAIIID)) = '%v' LIMIT %d OFFSET %d",
			kPLUSCustSavingsListSQL, arg.INAIIID, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE lower(trim(INAIIID)) = '%v' ", kPLUSCustSavingsListSQL, arg.INAIIID)
	}
	return populateKPLUSCustSavingsList(q, ctx, sql)
}

type SavingsListParams struct {
	INAIIID int64 `json:"iNAIIID"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

// CustSavingsList(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSCustSavingsList, error)
// GetTransactionHistory(ctx context.Context, in *KPLUSGetTransactionHistoryRequest, opts ...grpc.CallOption) (*KPLUSGetTransactionHistoryResponse, error)
// GenerateColShtperCID(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSGenerateColShtperCIDResponse, error)
// K2CCallBackRef(ctx context.Context, in *KPLUSCallBackRefRequest, opts ...grpc.CallOption) (*KPLUSResponse, error)
// GetReferences(ctx context.Context, in *KPLUSGetReferencesRequest, opts ...grpc.CallOption) (*KPLUSGetReferencesResponse, error)
// MultiplePayment(ctx context.Context, in *KPLUSMultiplePaymentRequest, opts ...grpc.CallOption) (*KPLUSResponse, error)
// SearchLoanList(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSSearchLoanListResponse, error)
// LoanInfo(ctx context.Context, in *KPLUSAccRequest, opts ...grpc.CallOption) (*KPLUSLoanInfoResponse, error)
// GetSavingForSuperApp(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSGetSavingResponse, error)
// FundTransferRequest(ctx context.Context, in *KPLUSFundTransferRequest, opts ...grpc.CallOption) (*KPLUSResponse, error)
