// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package database

import (
	"database/sql"
	"time"
)

type InfraAnalise struct {
	Atende       sql.NullInt32  `json:"atende"`
	Obs          sql.NullString `json:"obs"`
	CreateTime   sql.NullTime   `json:"create_time"`
	Username     sql.NullString `json:"username"`
	ID           int32          `json:"id"`
	ItensID      int32          `json:"itens_id"`
	FornecedorID int32          `json:"fornecedor_id"`
}

type InfraComponente struct {
	Username   sql.NullString `json:"username"`
	CreateTime sql.NullTime   `json:"create_time"`
	ID         int32          `json:"id"`
	Nome       sql.NullString `json:"nome"`
	ModeloID   int32          `json:"modelo_id"`
}

type InfraEquipamento struct {
	Username   sql.NullString `json:"username"`
	CreateTime sql.NullTime   `json:"create_time"`
	ID         int32          `json:"id"`
	Nome       sql.NullString `json:"nome"`
	TipoID     int32          `json:"tipo_id"`
	ModeloID   int32          `json:"modelo_id"`
}

type InfraFornecedor struct {
	Username   sql.NullString `json:"username"`
	CreateTime sql.NullTime   `json:"create_time"`
	ID         int32          `json:"id"`
	Nome       sql.NullString `json:"nome"`
}

type InfraIten struct {
	Data          time.Time `json:"data"`
	ID            int32     `json:"id"`
	PregaoNumero  int32     `json:"pregao_numero"`
	PregaoAno     int32     `json:"pregao_ano"`
	EquipamentoID int32     `json:"equipamento_id"`
	SolicitanteID int32     `json:"solicitante_id"`
}

type InfraMarca struct {
	Username   string    `json:"username"`
	CreateTime time.Time `json:"create_time"`
	ID         int32     `json:"id"`
	Nome       string    `json:"nome"`
}

type InfraModelo struct {
	Username   sql.NullString `json:"username"`
	CreateTime sql.NullTime   `json:"create_time"`
	ID         int32          `json:"id"`
	Nome       sql.NullString `json:"nome"`
	MarcaID    int32          `json:"marca_id"`
}

type InfraPeca struct {
	EquipamentoID int32          `json:"equipamento_id"`
	ComponenteID  int32          `json:"componente_id"`
	CreateTime    sql.NullTime   `json:"create_time"`
	Username      sql.NullString `json:"username"`
}

type InfraPregao struct {
	Username   sql.NullString `json:"username"`
	CreateTime sql.NullTime   `json:"create_time"`
	UpdateTime sql.NullTime   `json:"update_time"`
	Numero     int32          `json:"numero"`
	Ano        int32          `json:"ano"`
}

type InfraSolicitante struct {
	ID         int32          `json:"id"`
	Nome       string         `json:"nome"`
	CreateTime time.Time      `json:"create_time"`
	UpdateTime sql.NullTime   `json:"update_time"`
	Usuario    sql.NullString `json:"usuario"`
}

type InfraTipo struct {
	Username   sql.NullString `json:"username"`
	CreateTime sql.NullTime   `json:"create_time"`
	ID         int32          `json:"id"`
	Nome       sql.NullString `json:"nome"`
}