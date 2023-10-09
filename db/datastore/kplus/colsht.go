package db

import (
	"context"
	"fmt"
)

const kPLUSColShtSQL = `-- name: kPLUSColShtSQL :one
SELECT 
  id_number, BrCode, AppType, sht.Code, sht.Status, StatusDesc, sht.Acc, sht.Iiid, CustomerId, CentralOfficeId, CID, 
  UM, ClientName, CenterCode, CenterName, ManCode, Unit, AreaCode, Area, StaffName, AcctType, 
  AcctDesc, DisbDate, DateStart, Maturity, Principal, Interest, Gives, BalPrin IbalPrin, BalInt IbalInt, BalPrin, 
  BalInt, Amort, DuePrin, DueInt, LoanBal, SaveBal, WaivedInt, UnPaidCtr, WritenOff, Classification, 
  ClassDesc, WriteOff, 0 Pay, 0 Withdraw, 0 "Type", OrgName, OrgAddress, MeetingDate, MeetingDay, SharesOfStock, 
  DateEstablished, sht.Uuid
FROM ColSht sht
INNER JOIN IDs on ids.iiid  = sht.IIID
INNER JOIN Reference idType on idType.id = ids.Type_id and lower(idType.title) = 'inai-iiid'
WHERE lower(trim(IDs.id_number)) = $1::VarChar(30)
`

func populateKPLUScolSht(q *QueriesKPlus, ctx context.Context, sql string, arg ColShtParams) ([]ColShtperAcc, error) {
	rows, err := q.db.QueryContext(ctx, sql, arg.INAIIID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []ColShtperAcc{}
	for rows.Next() {
		var i ColShtperAcc
		err := rows.Scan(
			&i.INAIIID,
			&i.BrCode,
			&i.AppType,
			&i.Code,
			&i.Status,
			&i.StatusDesc,
			&i.Acc,
			&i.Iiid,
			&i.CustomerId,
			&i.CentralOfficeId,
			&i.CID,
			&i.UM,
			&i.ClientName,
			&i.CenterCode,
			&i.CenterName,
			&i.ManCode,
			&i.Unit,
			&i.AreaCode,
			&i.Area,
			&i.StaffName,
			&i.AcctType,
			&i.AcctDesc,
			&i.DisbDate,
			&i.DateStart,
			&i.Maturity,
			&i.Principal,
			&i.Interest,
			&i.Gives,
			&i.IbalPrin,
			&i.IbalInt,
			&i.BalPrin,
			&i.BalInt,
			&i.Amort,
			&i.DuePrin,
			&i.DueInt,
			&i.LoanBal,
			&i.SaveBal,
			&i.WaivedInt,
			&i.UnPaidCtr,
			&i.WritenOff,
			&i.Classification,
			&i.ClassDesc,
			&i.WriteOff,
			&i.Pay,
			&i.Withdraw,
			&i.Type,
			&i.OrgName,
			&i.OrgAddress,
			&i.MeetingDate,
			&i.MeetingDay,
			&i.SharesOfStock,
			&i.DateEstablished,
			&i.Uuid,
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

func (q *QueriesKPlus) ColSht(ctx context.Context, arg ColShtParams) ([]ColShtperAcc, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d",
			kPLUSColShtSQL, arg.Limit, arg.Offset)
	} else {
		sql = kPLUSColShtSQL
	}
	return populateKPLUScolSht(q, ctx, sql, arg)
}

type ColShtParams struct {
	INAIIID int64 `json:"iNAIIID"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

// GenerateColShtperCID(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSGenerateColShtperCIDResponse, error)
// K2CCallBackRef(ctx context.Context, in *KPLUSCallBackRefRequest, opts ...grpc.CallOption) (*KPLUSResponse, error)
// GetReferences(ctx context.Context, in *KPLUSGetReferencesRequest, opts ...grpc.CallOption) (*KPLUSGetReferencesResponse, error)
// MultiplePayment(ctx context.Context, in *KPLUSMultiplePaymentRequest, opts ...grpc.CallOption) (*KPLUSResponse, error)
// SearchLoanList(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSSearchLoanListResponse, error)
// LoanInfo(ctx context.Context, in *KPLUSAccRequest, opts ...grpc.CallOption) (*KPLUSLoanInfoResponse, error)
// GetSavingForSuperApp(ctx context.Context, in *KPLUSCustomerRequest, opts ...grpc.CallOption) (*KPLUSGetSavingResponse, error)
// FundTransferRequest(ctx context.Context, in *KPLUSFundTransferRequest, opts ...grpc.CallOption) (*KPLUSResponse, error)
