-- name: ListMarcas :many
select * from infra.marca
order by nome;

-- name: GetMarcaById :one
select * from infra.marca
where id = $1 limit 1;

-- name: GetMarcaByName :one
select * from infra.marca
where nome = $1 limit 1;

-- name: FindMarcaByName :many
select * from infra.marca
where nome like $1;

-- name: CreateMarca :execlastid
insert into infra.marca(
    username,nome)
values(
    $1, $2
);

-- name: UpdateMarca :exec
update infra.marca 
set nome = $2, usuario = $3
where id = $1; 

-- name: DeleteMarca :exec
delete from infra.marca
where id = $1;