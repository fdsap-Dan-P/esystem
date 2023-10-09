package db

import (
	"context"
	"errors"
	"log"
	"simplebank/util"
	"time"

	dsTrn "simplebank/db/datastore/transaction"

	"github.com/shopspring/decimal"
)

const kPLUSMultiplePaymentSQL = `-- name: kPLUSMultiplePaymentSQL :one
SELECT trn_serial 
FROM trn_head th
INNER JOIN Users u on u.id = th.user_id 
WHERE reference  = $1 and u.login_name = 'konek2CARD'
`

type AppType int64

const (
	AppTypeSavings AppType = 0
	AppTypeMBA     AppType = 1
	AppTypeLoan    AppType = 3
	AppTypeSource  AppType = 704
	AppTypeFee     AppType = 706
)

type Payment struct {
	AppType   int64           `json:"appType"`
	Acc       string          `json:"acc"`
	Pay       decimal.Decimal `json:"pay"`
	Withdraw  decimal.Decimal `json:"withdraw"`
	PaidPrin  decimal.Decimal `json:"paidPrin"`
	PaidInt   decimal.Decimal `json:"paidInt"`
	WaivedInt decimal.Decimal `json:"waivedInt"`
}

type MultiplePaymentRequest struct {
	RemitterINAIIID string    `json:"iNAIIID"`
	CustomerId      int64     `json:"customerId"`
	IIID            int64     `json:"iIID"`
	Cid             int64     `json:"cid"`
	PrNumber        string    `json:"prNumber"`
	SourceId        int64     `json:"sourceId"`
	OrNumber        int64     `json:"orNumber"`
	Username        string    `json:"username"`
	Trndate         string    `json:"trndate"`
	TotalCollection float64   `json:"totalCollection"`
	Particulars     string    `json:"particulars"`
	Payment         []Payment `json:"payment"`
}

func (q *QueriesKPlus) MultiplePayment(
	ctx context.Context, req MultiplePaymentRequest) (KPLUSResponse, error) {
	var i KPLUSResponse
	var err error
	// amort := make(map[string]map[int16]dsAcc.ScheduleInfo)
	// ln := make(map[string]dsAcc.AccountStat)
	balPrin := util.SetDecimal("0")
	balInt := util.SetDecimal("0")
	duePrin := util.SetDecimal("0")
	dueInt := util.SetDecimal("0")
	payPrin := util.SetDecimal("0")
	payInt := util.SetDecimal("0")
	newBal := util.SetDecimal("0")
	pay := util.SetDecimal("0")
	// waivedInt := util.SetDecimal("0")
	var saturday time.Time

	custs, custErr := q.GetCustomersInfo(context.Background(), CustomersInfoParam{
		INAIIID: req.RemitterINAIIID,
	})
	if custErr != nil {
		return i, custErr
	}
	if len(custs) <= 0 {
		return i, errors.New("remitter ID not Found.")
	}

	tic, ticErr := q.NewDailyTicket(context.Background(),
		dsTrn.NewDailyTicketRequest{
			CentralOffice: "CI",
			TicketType:    "API",
			TicketDate:    util.SetDate(req.Trndate),
			Postedby:      req.Username,
			Status:        "10",
			Remarks:       util.SetNullString(req.Particulars),
		})

	if ticErr != nil {
		return i, nil
	}

	typ, typErr := q.GetReferenceInfobyTitle(context.Background(), "TrnHeadType", 0, "Multiple Payment")

	if typErr != nil {
		return i, nil
	}

	trnHead := dsTrn.TrnHeadRequest{
		Uuid:            util.UUID(),
		TicketId:        tic.Id,
		TrnDate:         util.SetDate(req.Trndate),
		TrnSerial:       util.UUID().String(),
		TypeId:          typ.Id,
		OfficeId:        custs[0].OfficeId,
		UserId:          tic.PostedbyId,
		TransactingIiid: util.SetNullInt64(custs[0].Iiid),
		Orno:            util.NullString(),
		Isfinal:         util.SetNullBool(true),
		Ismanual:        util.SetNullBool(false),
		AlternateTrn:    util.NullString(),
		Reference:       req.PrNumber,
		Particular:      util.SetNullString(typ.Title),
		// OtherInfo:    sql.NullString(sql.NullString{String: string(info), Valid: true}),
	}

	log.Printf("trnHead:%s ", util.Struct2Json(trnHead))

	// lnStat := make(map[string]dsAcc.AccountStat)
	var paidN int16
	accs := []string{}

	for _, d := range req.Payment {
		// log.Printf("i:%v d:%v", i, d)
		accs = append(accs, d.Acc)
		// log.Printf("amort:%v err:%v", amort, err)
	}

	ln, err := q.GetAccountStat(ctx, accs)
	if err != nil {
		return KPLUSResponse{
			RetCode:   0,
			Message:   err.Error(),
			Reference: "",
		}, err
	}
	amort, err := q.GetSchedulebyAcc(ctx, accs)
	if err != nil {
		return KPLUSResponse{
			RetCode:   0,
			Message:   err.Error(),
			Reference: "",
		}, err
	}
	// log.Printf("i:%v acc:%v ln:%v err:%v", i, ln[d.Acc].Acc, ln[d.Acc].BalPrin, err)

	for i, d := range req.Payment {
		balPrin = ln[d.Acc].BalPrin
		balInt = ln[d.Acc].BalInt
		newBal = balPrin.Add(balInt)
		newBal = newBal.Sub(d.Pay)
		saturday = ln[d.Acc].MeetingDate
		saturday = saturday.Add(time.Duration(saturday.Weekday()) * 24 * time.Hour)
		pay = d.Pay
		// waivedInt = decimal.Zero
		log.Printf("Start compute Application i:%v. acc:%v, balprin:%v, balInt:%v, nrebal:%v",
			i, ln[d.Acc].Acc, balPrin, balInt, newBal)
		for k := int16(1); k <= int16(len(amort[d.Acc])); k++ {
			s := amort[d.Acc][k]
			// log.Printf("k: %v EndPrin:%v EndInt:%v newBal:%v ",
			// 	k, s.EndPrin, s.EndInt, newBal)
			if s.EndPrin.Add(s.EndInt).LessThanOrEqual(newBal) {
				paidN = k
				log.Printf("Found dNum: %v", paidN)
				break
			}
			// log.Printf("%v ---| Pass Loop", k)
		}
		s := amort[d.Acc][paidN]
		duePrin = balPrin.Sub(s.EndPrin)
		dueInt = balInt.Sub(s.EndInt)
		// Apply Interest Payment First
		if dueInt.LessThanOrEqual(pay) {
			payInt = dueInt
			pay = pay.Sub(dueInt)
		} else {
			payInt = pay
			pay = decimal.Zero
		}

		// Apply Remaining to Principal
		if duePrin.LessThanOrEqual(pay) {
			payPrin = duePrin
			pay = pay.Sub(duePrin)
		} else {
			payPrin = pay
			pay = decimal.Zero
		}
		log.Printf("Acc:%v Last->balPrin:%v balInt:%v payPrin:%v payInt:%v newBal:%v pay:%v err:%v:",
			d.Acc, balPrin, balInt, payPrin, payInt, newBal, pay, err)
	}

	// dtTrnHead, trnHeadErr := q.CreateTrnHead(ctx, trnHead)

	log.Printf("trnHead: %v ", trnHead)
	// var serial string
	// row := q.db.QueryRowContext(ctx, kPLUSMultiplePaymentSQL, prNo)
	// err := row.Scan(
	// 	&serial,
	// )
	// if err == nil {
	// 	i.RetCode = 1001
	// 	i.Message = "Transaction Exist!"
	// 	i.Reference = serial
	// } else {
	// 	i.RetCode = 0
	// 	i.Message = "Transaction not Exist!"
	// 	i.Reference = ""
	// 	err = nil
	// }
	return i, nil

}
