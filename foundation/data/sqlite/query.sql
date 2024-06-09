-- name: GetUsers :many
select * from users
order by created_at asc, role asc;

-- name: GetUser :one
select * from users
where id = ? limit 1;

-- name: CreateUser :one
insert into users (email, first_name, last_name, password, role)
values (?, ?, ?, ?, ?)
RETURNING *;

-- name: GetProjects :many
select sqlc.embed(project), sqlc.embed(users) from project
join users on project.user_id = users.id
order by project.created_at asc, project.name asc;

-- name: GetProject :one
select * from project
where id = ? limit 1;

-- name: CreateProject :one
insert into project (name, description, user_id)
values (?, ?, ?)
RETURNING *;