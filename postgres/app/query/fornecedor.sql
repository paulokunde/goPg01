
-- name: ListFornecedores :many
select * from fornecedor
order by nome;

-- name: GetFornecedorById :one
select * from fornecedor
where id = $1 limit 1;

-- name: GetFornecedorByName :one
select * from fornecedor
where nome = $1 limit 1;

-- name: FindFornecedorByName :many
select * from fornecedor
where nome like $1;

-- name: CreateFornecedor :one
insert into fornecedor(
    nome)
values(
    $1
)returning *;

-- name: UpdateFornecedor :one
update fornecedor
set nome = $2
where id = $1 returning *; 

-- name: DeleteFornecedor :exec
delete from fornecedor
where id = $1;
