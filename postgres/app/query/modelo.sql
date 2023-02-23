
-- name: ListModelos :many
select * from modelo
order by nome;

-- name: GetModeloById :one
select * from modelo
where id = $1 limit 1;

-- name: GetModeloByName :one
select * from modelo
where nome = $1 limit 1;

-- name: FindModelosByName :many
select * from modelo
where nome like $1
limit $2 offset $3;

-- name: FindModelosByMarca :many
select * from modelo
where marca_id = $1
limit $2 offset $3;

-- name: CreateModelo :one
insert into modelo(
    nome,marca_id)
values(
    $1,$2
)returning *;

-- name: UpdateModelo :one
update modelo 
set nome = $2,marca_id = $3
where id = $1 returning *; 

-- name: DeleteModelo :exec
delete from modelo
where id = $1;

-- name: FindMarcasByName :many
select * from marca
where nome like $1;