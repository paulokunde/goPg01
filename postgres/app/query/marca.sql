
-- name: ListMarcas :many
select * from marca
order by nome;

-- name: GetMarcaById :one
select * from marca
where id = $1 limit 1;

-- name: GetMarcaByName :one
select * from marca
where nome = $1 limit 1;

-- name: FindMarcaByName :many
select * from marca
where nome like $1
limit $2 offset $3;

-- name: CreateMarca :one
insert into marca(
    nome)
values(
    $1
)returning *;

-- name: UpdateMarca :one
update marca 
set nome = $2
where id = $1 returning *; 

-- name: DeleteMarca :exec
delete from marca
where id = $1;