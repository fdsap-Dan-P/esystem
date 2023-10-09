package api

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	ds "simplebank/db/datastore"
	account "simplebank/db/datastore/account"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type transferRequest struct {
	FromAccountId int64           `json:"fromAccountId" binding:"required,min=1"`
	ToAccountId   int64           `json:"toAccountId" binding:"required,min=1"`
	Amount        decimal.Decimal `json:"amount" binding:"required"`
	Currency      string          `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fromAccount, valid := server.validAccount(ctx, req.FromAccountId, req.Currency)
	if !valid {
		return
	}

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// Correct to validate user name
	if fromAccount.CustomerId != fromAccount.CustomerId { //authPayload.Username {
		err := errors.New("from account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	_, valid = server.validAccount(ctx, req.ToAccountId, req.Currency)
	if !valid {
		return
	}

	arg := ds.TransferTxParams{
		FromAccountId: req.FromAccountId,
		ToAccountId:   req.ToAccountId,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) (account.AccountInfo, bool) {
	log.Println("API: validAccounts..1--:", accountID, ctx)
	account, err := server.store.GetAccount(ctx, accountID)
	log.Println("API: validAccounts..2--: ", ctx)
	if err != nil {
		log.Println("API: validAccounts..3--: ", ctx)
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account, false
		}

		log.Println("API: validAccounts..4--: ", ctx)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}

	log.Println("API: validAccounts..5--: ", ctx)
	if account.Currency != currency {
		log.Println("API: validAccounts..6--: ", account)
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.Id, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return account, false
	}
	log.Println("API: validAccounts..7--: ", account.Balance)

	return account, true
}
