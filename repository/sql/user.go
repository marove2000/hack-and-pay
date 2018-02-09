package sql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/marove2000/hack-and-pay/contract"
	"github.com/marove2000/hack-and-pay/errors"
	"github.com/marove2000/hack-and-pay/repository/sql/models"

	sqlerror "github.com/pkg/errors"
	"github.com/vattle/sqlboiler/queries/qm"
	"gopkg.in/nullbio/null.v6"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/marove2000/hack-and-pay/config"
	"github.com/go-sql-driver/mysql"
)

func (m *Mysql) AddLocalUser(ctx context.Context, body *contract.AddUserRequestBody) (userID int, err error) {
	logger := pkgLogger.ForFunc(ctx, "AddLocalUser")
	logger.Debug("enter repository")

	tx, err := m.db.Beginx()
	defer func() {
		errr := tx.Rollback()
		if errr != nil && errr != sql.ErrTxDone {
			logger.WithError(errr).Error("failed to roll back tx")
			err = errors.InternalServerError("db error", errr)
		}
	}()

	usr := models.User{
		Name:  body.Name,
		Email: null.StringFrom(body.Email),
	}

	// add user
	err = usr.Insert(tx, models.UserColumns.Name, models.UserColumns.Email)
	if err != nil {
		sqlerr, ok := sqlerror.Cause(err).(*mysql.MySQLError)
		if !ok {
			logger.WithError(err).Error("failed to insert transaction")
			return 0, errors.InternalServerError("db error", err)
		}

		switch sqlerr.Number {
		case 1062:
			logger.WithField("username", body.Name).Warn("duplicate entry for username")
			return 0, errors.BadRequest("bad request")
		default:
			logger.WithError(err).Error("failed to insert transaction")
			return 0, errors.InternalServerError("db error", err)
		}

	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.WithError(err).Error("failed to create password hash")
		return 0, errors.InternalServerError("password hash error", err)
	}

	auth := models.UserAuth{
		Method: string(contract.AuthTypePasswd),
		UserID: usr.UserID,
		Value:  null.Bytes{Bytes: hashedPassword, Valid: true},
	}

	err = auth.Insert(tx, models.UserAuthColumns.UserID, models.UserAuthColumns.Method, models.UserAuthColumns.Value)
	if err != nil {
		logger.WithError(err).Error("failed to insert user auth")
		return 0, errors.InternalServerError("db error", err)
	}

	err = tx.Commit()
	if err != nil {
		logger.WithError(err).Error("failed to commit")
		return 0, errors.InternalServerError("db error", err)
	}

	return usr.UserID, err

}

func (m *Mysql) AddLDAPUser(ctx context.Context, body *contract.AddUserRequestBody) (userID int, err error) {
	logger := pkgLogger.ForFunc(ctx, "AddLDAPUser")
	logger.Debug("enter repository")

	tx, err := m.db.Beginx()
	defer func() {
		errr := tx.Rollback()
		if errr != nil && errr != sql.ErrTxDone {
			logger.WithError(errr).Error("failed to roll back tx")
			err = errors.InternalServerError("db error", errr)
		}
	}()

	usr := models.User{
		Name:  body.Name,
		Email: null.StringFrom(body.Email),
	}

	// add user
	err = usr.Insert(tx, models.UserColumns.Name, models.UserColumns.Email)
	if err != nil {
		logger.WithError(err).Error("failed to insert user")
		return 0, errors.InternalServerError("db error", err)
	}

	auth := models.UserAuth{
		Method: string(contract.AuthTypeLDAP),
		UserID: usr.UserID,
	}

	err = auth.Insert(tx, models.UserAuthColumns.UserID, models.UserAuthColumns.Method)
	if err != nil {
		logger.WithError(err).Error("failed to insert user auth")
		return 0, errors.InternalServerError("db error", err)
	}

	err = tx.Commit()
	if err != nil {
		logger.WithError(err).Error("failed to commit")
		return 0, errors.InternalServerError("db error", err)
	}

	return usr.UserID, err
}

func (m *Mysql) GetPublicUserDataByUserID(ctx context.Context, userID int) (*contract.User, error) {
	logger := pkgLogger.ForFunc(ctx, "GetPublicUserDataByUserID")
	logger.Debug("enter repository")

	user, err := models.Users(m.db, qm.Where("user_id=?", userID)).One()
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Warn("failed to find user")
			return nil, errors.NotFound("user not found")
		}

		logger.WithError(err).Error("failed to fetch users from db")
		return nil, errors.InternalServerError("db error", err)
	}

	return &contract.User{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email.String,
	}, nil
}

func (m *Mysql) CheckIsAdminJWT(ctx context.Context, JWT string, userID int) (error) {
	logger := pkgLogger.ForFunc(ctx, "CheckJWT")
	logger.Debug("enter repository")

	// read config
	conf := config.ReadConfig()

	token, err := jwt.Parse(JWT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			return nil, err
		}
		return []byte(conf.JWT.Secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if claims["userIsAdmin"].(bool) == true {
			// user is admin
			// TODO: Check Timeout
			return nil
		} else {
			err := fmt.Errorf("user is no admin")
			return err
		}

	} else {
		return err
	}
}

func (m *Mysql) GetPublicUserDataByUserName(ctx context.Context, name string) (*contract.User, error) {
	logger := pkgLogger.ForFunc(ctx, "GetPublicUserDataByUserName")
	logger.Debug("enter repository")

	user, err := models.Users(m.db, qm.Where("name=?", name)).One()
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Warn("failed to find user")
			return nil, errors.NotFound("user not found")
		}

		logger.WithError(err).Error("failed to fetch users from db")
		return nil, errors.InternalServerError("db error", err)
	}

	return &contract.User{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email.String,
	}, nil
}

func (m *Mysql) Login(ctx context.Context, name, pass string) error {
	logger := pkgLogger.ForFunc(ctx, "Login")
	logger.Debug("enter Login")

	// get user id by name
	user, err := models.Users(m.db, qm.Where("name=?", name)).One()
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Warn("failed to find user")
			return errors.NotFound("user not found")
		}

		logger.WithError(err).Error("failed to fetch users from db")
		return errors.InternalServerError("db error", err)
	}

	// get user_auth data by id
	auth, err := models.UserAuths(m.db, qm.Where("user_id=?", user.UserID)).One()
	if err != nil {
		logger.WithError(err).Error("failed to fetch users auth from db")
		return errors.InternalServerError("db error", err)
	}

	// check auth value
	err = bcrypt.CompareHashAndPassword([]byte(auth.Value.Bytes), []byte(pass))
	if err != nil {
		logger.WithError(err).Error("wrong password")
		return errors.Unauthorized()
	}

	return nil
}
