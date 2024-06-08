-- name: GetUsers :many
select * from users
order by (created_at, role) asc;

-- name: GetUser :one
select * from users
where id = ? limit 1;

-- name: CreateUser :exec
insert into users (email, first_name, last_name, password, role)
values (?, ?, ?, ?, ?);

-- name: GetProjects :many
select * from project
order by (created_at, name) asc;

-- name: GetProject :one
select * from project
where id = ? limit 1;

-- name: CreateProject :exec
insert into project (name, description, user_id)
values (?, ?, ?);