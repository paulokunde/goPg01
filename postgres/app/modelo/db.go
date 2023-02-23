// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package modelo

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createFornecedorStmt, err = db.PrepareContext(ctx, createFornecedor); err != nil {
		return nil, fmt.Errorf("error preparing query CreateFornecedor: %w", err)
	}
	if q.createMarcaStmt, err = db.PrepareContext(ctx, createMarca); err != nil {
		return nil, fmt.Errorf("error preparing query CreateMarca: %w", err)
	}
	if q.createModeloStmt, err = db.PrepareContext(ctx, createModelo); err != nil {
		return nil, fmt.Errorf("error preparing query CreateModelo: %w", err)
	}
	if q.deleteFornecedorStmt, err = db.PrepareContext(ctx, deleteFornecedor); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteFornecedor: %w", err)
	}
	if q.deleteMarcaStmt, err = db.PrepareContext(ctx, deleteMarca); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteMarca: %w", err)
	}
	if q.deleteModeloStmt, err = db.PrepareContext(ctx, deleteModelo); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteModelo: %w", err)
	}
	if q.findFornecedorByNameStmt, err = db.PrepareContext(ctx, findFornecedorByName); err != nil {
		return nil, fmt.Errorf("error preparing query FindFornecedorByName: %w", err)
	}
	if q.findMarcaByNameStmt, err = db.PrepareContext(ctx, findMarcaByName); err != nil {
		return nil, fmt.Errorf("error preparing query FindMarcaByName: %w", err)
	}
	if q.findMarcasByNameStmt, err = db.PrepareContext(ctx, findMarcasByName); err != nil {
		return nil, fmt.Errorf("error preparing query FindMarcasByName: %w", err)
	}
	if q.findModelosByMarcaStmt, err = db.PrepareContext(ctx, findModelosByMarca); err != nil {
		return nil, fmt.Errorf("error preparing query FindModelosByMarca: %w", err)
	}
	if q.findModelosByNameStmt, err = db.PrepareContext(ctx, findModelosByName); err != nil {
		return nil, fmt.Errorf("error preparing query FindModelosByName: %w", err)
	}
	if q.getFornecedorByIdStmt, err = db.PrepareContext(ctx, getFornecedorById); err != nil {
		return nil, fmt.Errorf("error preparing query GetFornecedorById: %w", err)
	}
	if q.getFornecedorByNameStmt, err = db.PrepareContext(ctx, getFornecedorByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetFornecedorByName: %w", err)
	}
	if q.getMarcaByIdStmt, err = db.PrepareContext(ctx, getMarcaById); err != nil {
		return nil, fmt.Errorf("error preparing query GetMarcaById: %w", err)
	}
	if q.getMarcaByNameStmt, err = db.PrepareContext(ctx, getMarcaByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetMarcaByName: %w", err)
	}
	if q.getModeloByIdStmt, err = db.PrepareContext(ctx, getModeloById); err != nil {
		return nil, fmt.Errorf("error preparing query GetModeloById: %w", err)
	}
	if q.getModeloByNameStmt, err = db.PrepareContext(ctx, getModeloByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetModeloByName: %w", err)
	}
	if q.listFornecedoresStmt, err = db.PrepareContext(ctx, listFornecedores); err != nil {
		return nil, fmt.Errorf("error preparing query ListFornecedores: %w", err)
	}
	if q.listMarcasStmt, err = db.PrepareContext(ctx, listMarcas); err != nil {
		return nil, fmt.Errorf("error preparing query ListMarcas: %w", err)
	}
	if q.listModelosStmt, err = db.PrepareContext(ctx, listModelos); err != nil {
		return nil, fmt.Errorf("error preparing query ListModelos: %w", err)
	}
	if q.updateFornecedorStmt, err = db.PrepareContext(ctx, updateFornecedor); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateFornecedor: %w", err)
	}
	if q.updateMarcaStmt, err = db.PrepareContext(ctx, updateMarca); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMarca: %w", err)
	}
	if q.updateModeloStmt, err = db.PrepareContext(ctx, updateModelo); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateModelo: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createFornecedorStmt != nil {
		if cerr := q.createFornecedorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createFornecedorStmt: %w", cerr)
		}
	}
	if q.createMarcaStmt != nil {
		if cerr := q.createMarcaStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createMarcaStmt: %w", cerr)
		}
	}
	if q.createModeloStmt != nil {
		if cerr := q.createModeloStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createModeloStmt: %w", cerr)
		}
	}
	if q.deleteFornecedorStmt != nil {
		if cerr := q.deleteFornecedorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteFornecedorStmt: %w", cerr)
		}
	}
	if q.deleteMarcaStmt != nil {
		if cerr := q.deleteMarcaStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteMarcaStmt: %w", cerr)
		}
	}
	if q.deleteModeloStmt != nil {
		if cerr := q.deleteModeloStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteModeloStmt: %w", cerr)
		}
	}
	if q.findFornecedorByNameStmt != nil {
		if cerr := q.findFornecedorByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findFornecedorByNameStmt: %w", cerr)
		}
	}
	if q.findMarcaByNameStmt != nil {
		if cerr := q.findMarcaByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findMarcaByNameStmt: %w", cerr)
		}
	}
	if q.findMarcasByNameStmt != nil {
		if cerr := q.findMarcasByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findMarcasByNameStmt: %w", cerr)
		}
	}
	if q.findModelosByMarcaStmt != nil {
		if cerr := q.findModelosByMarcaStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findModelosByMarcaStmt: %w", cerr)
		}
	}
	if q.findModelosByNameStmt != nil {
		if cerr := q.findModelosByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findModelosByNameStmt: %w", cerr)
		}
	}
	if q.getFornecedorByIdStmt != nil {
		if cerr := q.getFornecedorByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFornecedorByIdStmt: %w", cerr)
		}
	}
	if q.getFornecedorByNameStmt != nil {
		if cerr := q.getFornecedorByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFornecedorByNameStmt: %w", cerr)
		}
	}
	if q.getMarcaByIdStmt != nil {
		if cerr := q.getMarcaByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMarcaByIdStmt: %w", cerr)
		}
	}
	if q.getMarcaByNameStmt != nil {
		if cerr := q.getMarcaByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMarcaByNameStmt: %w", cerr)
		}
	}
	if q.getModeloByIdStmt != nil {
		if cerr := q.getModeloByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getModeloByIdStmt: %w", cerr)
		}
	}
	if q.getModeloByNameStmt != nil {
		if cerr := q.getModeloByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getModeloByNameStmt: %w", cerr)
		}
	}
	if q.listFornecedoresStmt != nil {
		if cerr := q.listFornecedoresStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listFornecedoresStmt: %w", cerr)
		}
	}
	if q.listMarcasStmt != nil {
		if cerr := q.listMarcasStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listMarcasStmt: %w", cerr)
		}
	}
	if q.listModelosStmt != nil {
		if cerr := q.listModelosStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listModelosStmt: %w", cerr)
		}
	}
	if q.updateFornecedorStmt != nil {
		if cerr := q.updateFornecedorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateFornecedorStmt: %w", cerr)
		}
	}
	if q.updateMarcaStmt != nil {
		if cerr := q.updateMarcaStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMarcaStmt: %w", cerr)
		}
	}
	if q.updateModeloStmt != nil {
		if cerr := q.updateModeloStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateModeloStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                       DBTX
	tx                       *sql.Tx
	createFornecedorStmt     *sql.Stmt
	createMarcaStmt          *sql.Stmt
	createModeloStmt         *sql.Stmt
	deleteFornecedorStmt     *sql.Stmt
	deleteMarcaStmt          *sql.Stmt
	deleteModeloStmt         *sql.Stmt
	findFornecedorByNameStmt *sql.Stmt
	findMarcaByNameStmt      *sql.Stmt
	findMarcasByNameStmt     *sql.Stmt
	findModelosByMarcaStmt   *sql.Stmt
	findModelosByNameStmt    *sql.Stmt
	getFornecedorByIdStmt    *sql.Stmt
	getFornecedorByNameStmt  *sql.Stmt
	getMarcaByIdStmt         *sql.Stmt
	getMarcaByNameStmt       *sql.Stmt
	getModeloByIdStmt        *sql.Stmt
	getModeloByNameStmt      *sql.Stmt
	listFornecedoresStmt     *sql.Stmt
	listMarcasStmt           *sql.Stmt
	listModelosStmt          *sql.Stmt
	updateFornecedorStmt     *sql.Stmt
	updateMarcaStmt          *sql.Stmt
	updateModeloStmt         *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                       tx,
		tx:                       tx,
		createFornecedorStmt:     q.createFornecedorStmt,
		createMarcaStmt:          q.createMarcaStmt,
		createModeloStmt:         q.createModeloStmt,
		deleteFornecedorStmt:     q.deleteFornecedorStmt,
		deleteMarcaStmt:          q.deleteMarcaStmt,
		deleteModeloStmt:         q.deleteModeloStmt,
		findFornecedorByNameStmt: q.findFornecedorByNameStmt,
		findMarcaByNameStmt:      q.findMarcaByNameStmt,
		findMarcasByNameStmt:     q.findMarcasByNameStmt,
		findModelosByMarcaStmt:   q.findModelosByMarcaStmt,
		findModelosByNameStmt:    q.findModelosByNameStmt,
		getFornecedorByIdStmt:    q.getFornecedorByIdStmt,
		getFornecedorByNameStmt:  q.getFornecedorByNameStmt,
		getMarcaByIdStmt:         q.getMarcaByIdStmt,
		getMarcaByNameStmt:       q.getMarcaByNameStmt,
		getModeloByIdStmt:        q.getModeloByIdStmt,
		getModeloByNameStmt:      q.getModeloByNameStmt,
		listFornecedoresStmt:     q.listFornecedoresStmt,
		listMarcasStmt:           q.listMarcasStmt,
		listModelosStmt:          q.listModelosStmt,
		updateFornecedorStmt:     q.updateFornecedorStmt,
		updateMarcaStmt:          q.updateMarcaStmt,
		updateModeloStmt:         q.updateModeloStmt,
	}
}