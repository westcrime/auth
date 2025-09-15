package query

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/westcrime/auth/internal/model"
)

func CreateUser(ctx context.Context, pool *pgxpool.Pool, createUser *model.CreateUser) (error, int64) {
	builderInsert := sq.Insert("users").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "email", "password_hash", "role", "created_at").
		Values(createUser.Name, createUser.Email, createUser.Password, createUser.Role, time.Now()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err, -1
	}

	var id int64
	err = pool.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return err, -1
	}

	return nil, id
}

func UpdateUser(ctx context.Context, pool *pgxpool.Pool, updateUser *model.UpdateUser) error {
	builderInsert := sq.Update("users").
		PlaceholderFormat(sq.Dollar).
		Set("email", updateUser.Info.Email).
		Set("name", updateUser.Info.Name).
		Where(sq.Eq{"id": updateUser.Id})

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	res, err := pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	fmt.Printf("Rows affected: %d\n", res.RowsAffected())
	fmt.Printf("Command: %s\n", res.String())
	return nil
}

func DeleteUser(ctx context.Context, pool *pgxpool.Pool, user_id int64) error {
	builderInsert := sq.Delete("users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": user_id})

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	res, err := pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	fmt.Printf("Rows affected: %d\n", res.RowsAffected())
	fmt.Printf("Command: %s\n", res.String())
	return nil
}

// func GetUsers(ctx context.Context, pool *pgxpool.Pool) (error, []model.User) {
// 	builderInsert := sq.Select("id", "email", "name", "role", "created_at", "updated_at").
// 		From("users").
// 		PlaceholderFormat(sq.Dollar).
// 		OrderBy("id ASC").
// 		Limit(100)

// 	query, args, err := builderInsert.ToSql()
// 	if err != nil {
// 		return err, nil
// 	}

// 	rows, err := pool.Query(ctx, query, args...)
// 	if err != nil {
// 		return err, nil
// 	}

// 	var id, role int64
// 	var email, name string
// 	var createdAt time.Time
// 	var updatedAt sql.NullTime

// 	users := make([]model.User, 2)

// 	for rows.Next() {
// 		err = rows.Scan(&id, &email, &name, &role, &createdAt, &updatedAt)
// 		if err != nil {
// 			return err, nil
// 		}
// 		users = append(users, model.User{
// 			Id:        id,
// 			Name:      name,
// 			Email:     email,
// 			Role:      model.Role(role),
// 			CreatedAt: createdAt,
// 			UpdatedAt: updatedAt.Time,
// 		})
// 	}
// 	return nil, users
// }

func GetUser(ctx context.Context, pool *pgxpool.Pool, user_id int64) (error, model.User) {
	builderInsert := sq.Select("id", "email", "name", "role", "created_at", "updated_at").
		From("users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": user_id}).
		Limit(1)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err, model.User{}
	}

	rows, err := pool.Query(ctx, query, args...)
	if err != nil {
		return err, model.User{}
	}

	var id, role int64
	var email, name string
	var createdAt time.Time
	var updatedAt sql.NullTime

	user := model.User{}

	for rows.Next() {
		err = rows.Scan(&id, &email, &name, &role, &createdAt, &updatedAt)
		if err != nil {
			return err, model.User{}
		}
		user = model.User{
			Id:        id,
			Name:      name,
			Email:     email,
			Role:      model.Role(role),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt.Time,
		}
	}
	return nil, user
}
