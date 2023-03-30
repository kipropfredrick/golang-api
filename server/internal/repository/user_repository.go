package repository

import (
	"context"
	"database/sql"
	"server/internal/models"

)

//inner most layer interacting with database
//create interface
type DTBX interface{
	
	ExecContext(ctx context.Context,query string,args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context,query string)(*sql.Stmt,error)
	QueryContext(context.Context,string, ...interface{})(*sql.Rows,error)
	QueryRowContext(context.Context,string, ...interface{}) *sql.Row
}
//create struct type of repository
type repository struct {
   db DTBX
}
//create func and inject it with interface
func NewUserRepository(db DTBX) models.Repository {
	// return a pointer to the repository
    return &repository{
		db: db,
	}
}
//function to create user
func (repo *repository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	var lastinsertid int
	query := "INSERT INTO userS (username,email,password) VALUES ($1,$2,$3) returning id"
	err:= repo.db.QueryRowContext(ctx, query,user.Username,user.Email,user.Password).Scan(&lastinsertid)
	if err != nil {
		return &models.User{},err
	}
	//typecast id to interface
	user.ID=int64(lastinsertid)
	return user, nil
}
//get user by email
func (user *repository) GetByEmail(ctx context.Context,email string)(*models.User,error){
   //empty struct
   u := models.User{}
   query := "SELECT id, email, username, password FROM users WHERE email = $1"
   err := user.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password)
   if err != nil {
	   return &models.User{}, nil
   }
   return &u,nil
}