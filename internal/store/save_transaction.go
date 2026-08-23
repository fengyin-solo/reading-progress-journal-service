package store

import "readinglog/internal/model"

type SaveTransaction struct {
	attempt   int
	committed bool
}

func BeginSave(attempt int) *SaveTransaction {
	return &SaveTransaction{attempt: attempt}
}

func (tx *SaveTransaction) WriteEntry() error {
	if tx.attempt == 1 {
		return model.ErrJournalWrite
	}
	return nil
}

func (tx *SaveTransaction) Commit() error {
	tx.committed = true
	return nil
}

func (tx *SaveTransaction) Rollback() error {
	if tx.attempt == 1 {
		return model.ErrRollbackCleanup
	}
	return nil
}

func (tx *SaveTransaction) Finish() (err error) {
	defer func() {
		if err != nil {
			err = tx.Rollback()
		}
	}()
	if err = tx.WriteEntry(); err != nil {
		return err
	}
	return tx.Commit()
}

func (tx *SaveTransaction) Committed() bool { return tx.committed }
