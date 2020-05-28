package transaction

import (
	"context"
	"database/sql"
	"errors"
	"gihub.com/team3_qgame/database"
	"log"

	"github.com/google/uuid"
)

var (
	transactionContextKey = "transactionContextKey"

	// The transaction has been already rollbacked
	ErrTxWasRollbacked = errors.New("the transaction has been rollbacked")
	// The transaction has not been started
	ErrNotInTransaction = errors.New("not in a transaction, but the method requires it (Begin() has not been called")
)

type (
	ITransactionContext interface {
		Begin() (uuid.UUID, error)
		Commit(uuid.UUID) error
		Rollback() error
		Provider() *sql.DB
	}

	transactionContext struct {
		dbManager       *database.DBConnection
		tx              *sql.Tx
		transactionUUID *uuid.UUID

		rollbacked bool
	}
)

func (t *transactionContext) GetTransactionContext(ctx context.Context) (*transactionContext, context.Context) {
	return getTransactionContext(ctx, t.dbManager)
}

func getTransactionContext(ctx context.Context, dbConn *database.DBConnection) (*transactionContext, context.Context) {
	transactionContext, found := ctx.Value(transactionContextKey).(*transactionContext)
	if !found {
		transactionContext := newTransactionContext(dbConn)
		newContext := context.WithValue(ctx, transactionContextKey, transactionContext)
		return transactionContext, newContext
	}
	return transactionContext, ctx
}

// Begin transaction
func (t *transactionContext) Begin() (id uuid.UUID, err error) {
	if t.wasRollbacked() {
		err = ErrTxWasRollbacked
		return
	}

	id, err = uuid.NewRandom()
	if err != nil {
		return
	}

	if !t.inTransaction() {
		t.transactionUUID = &id
		t.tx, err = t.dbManager.GetConnection().Begin()

		if err != nil {
			log.Printf("cant begin transaction (%v)\n", id)
			return
		}

		log.Printf("new transaction: %v", t.transactionUUID)
	} else {
		log.Printf("use old transaction: %v", t.transactionUUID)
	}

	return
}

func (t *transactionContext) Provider() *sql.DB {
	if t.wasRollbacked() {
		log.Println("transaction has been rollback!")
		return nil
	}

	if t.inTransaction() {
		return t.dbManager.
	}

	return t.providerWithoutTransaction()
}

func (t *transactionContext) Commit(id uuid.UUID) error {
	if t.wasRollbacked() {
		return ErrTxWasRollbacked
	}

	if !t.inTransaction() {
		return ErrNotInTransaction
	}

	// the current committer is a not transaction holder
	if *t.transactionUUID != id {
		return nil
	}

	defer t.dispose()

	if err := t.tx.Commit(); err != nil {
		log.Printf("cant commit transaction: %v; err: %s\n", t.transactionUUID, err)
		return err
	}

	return nil
}

func (t *transactionContext) Rollback() error {
	if t.wasRollbacked() {
		return ErrTxWasRollbacked
	}
	if !t.inTransaction() {
		log.Println("not in transaction; transaction is nil")
		return nil
	}

	defer t.disposeAfterRollback()

	if err := t.tx.Rollback(); err != nil {
		log.Printf("cant rollback (%v): %s", t.transactionUUID, err)
		return err
	}

	return nil
}

func (t *transactionContext) inTransaction() bool {
	return t.tx != nil && t.transactionUUID != nil
}

func (t *transactionContext) dispose() {
	log.Printf("disposing transaction (%v)\n", t.transactionUUID)

	t.tx = nil
	t.transactionUUID = nil
}

func (t *transactionContext) disposeAfterRollback() {
	t.rollbacked = true
	t.dispose()
}

func (t *transactionContext) wasRollbacked() bool {
	return t.rollbacked
}

func (t *transactionContext) providerWithoutTransaction() *sql.DB {
	return t.dbManager.GetConnection()
}

func newTransactionContext(dbHolder *database.DBConnection) *transactionContext {
	return &transactionContext{dbManager: dbHolder}
}
