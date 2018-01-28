package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/marove2000/hack-and-pay/contract"
	"github.com/marove2000/hack-and-pay/errors"

	"github.com/go-validator/validator"
)

func (h *Handler) publicUserIndex(ctx context.Context, r *http.Request, pathParams map[string]string) (interface{}, error) {
	logger := pkgLogger.ForFunc(ctx, "publicUserIndex")
	logger.Debug("enter handler")

	// get all user data
	users, err := h.repo.GetUsersWithBalance(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (h *Handler) signUp(ctx context.Context, r *http.Request, pathParams map[string]string) (interface{}, error) {
	logger := pkgLogger.ForFunc(ctx, "signUp")
	logger.Debug("enter handler")

	user := &contract.AddUserRequestBody{}

	// get body data
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(user)
	if err != nil {
		logger.WithError(err).Error("failed to parse body")
		return nil, errors.BadRequest(err.Error())
	}
	defer r.Body.Close()

	// validate data
	if err = validator.Validate(user); err != nil {
		logger.WithError(err).Warn("bad request")
		return nil, errors.BadRequest(err.Error())
	}

	// check if ldap is active
	var id int
	switch {
	case h.ldap.IsActive():
		// login with ldap
		err = h.ldap.Login(ctx, user.Name, user.Password)
		if err != nil {
			return nil, err
		}

		// login correct, create user in DB
		id, err = h.repo.AddLDAPUser(ctx, user)
		if err != nil {
			return nil, err
		}

	default:
		// ldap not active create user account
		id, err = h.repo.AddLocalUser(ctx, user)
		if err != nil {
			return nil, err
		}
	}
	return &contract.AddUserResponseBody{UserID: id}, err
}

func (h *Handler) login(ctx context.Context, r *http.Request, pathParams map[string]string) (interface{}, error) {
	//TODO Finish function
	logger := pkgLogger.ForFunc(ctx, "login")
	logger.Debug("enter handler")

	user := &contract.LoginRequestBody{}

	// get body data
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(user)
	if err != nil {
		logger.WithError(err).Error("failed to parse body")
		return nil, errors.BadRequest(err.Error())
	}
	defer r.Body.Close()

	if err = validator.Validate(user); err != nil {
		logger.WithError(err).Warn("bad request")
		return nil, errors.BadRequest(err.Error())
	}

	u, err := h.repo.GetPublicUserDataByUserName(ctx, user.Name)
	if err != nil {
		return nil, err
	}

	if h.ldap.IsActive() {
		err = h.ldap.Login(ctx, user.Name, user.Password)
		if err != nil {
			return nil, err
		}
	} else {
		// TODO local login
	}

	// TODO build jwt

	return u, nil

	// check if user is in Database and get public user data
	//dbUser, err = v1.GetPublicUserDataByUserName(user.UserName)
	//if err != nil {
	//	return nil, err
	//} else if dbUser.UserID == 0 {
	//	// user does not exist, check ldap first
	//	err = v1.GetLDAPAuthentication(user)
	//	if err != nil {
	//		// user does not exist in ldap too or credentials are wrong
	//		log.Println(err)
	//		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//		w.WriteHeader(http.StatusUnauthorized)
	//		return
	//	} else {
	//		// user is existing in ldap
	//		// TODO get LDAP Email Address from LDAP-field mail
	//		v1.AddUser(user, "ldap")
	//		user.UserJWT, err = v1.JWTGet(user)
	//		if err != nil {
	//			log.Println(err)
	//			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//			w.WriteHeader(http.StatusInternalServerError)
	//			return
	//		} else {
	//			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//			w.WriteHeader(http.StatusOK)
	//			if err := json.NewEncoder(w).Encode(user); err != nil {
	//				log.Println(err)
	//				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//				w.WriteHeader(http.StatusInternalServerError)
	//				return
	//			}
	//		}
	//	}
	//
	//	// if ldap also does not exist
	//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return
	//}
	//
	//if conf.LDAPActive == false {
	//	// LDAP not active, check password in database
	//
	//	err = v1.CheckPassword(dbUser, []byte(user.UserPassword))
	//	if err != nil {
	//		log.Println(err)
	//		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//		w.WriteHeader(http.StatusUnauthorized)
	//		return
	//	} else {
	//		dbUser.UserJWT, err = v1.JWTGet(dbUser)
	//		if err != nil {
	//			log.Println(err)
	//			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//			w.WriteHeader(http.StatusInternalServerError)
	//			return
	//		}
	//	}
	//} else {
	//	// LDAP active, check with LDAP
	//	err = v1.GetLDAPAuthentication(user)
	//	if err != nil {
	//		// user does not exist in ldap too or credentials are wrong
	//		log.Println(err)
	//		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//		w.WriteHeader(http.StatusUnauthorized)
	//		return
	//	} else {
	//		dbUser.UserJWT, err = v1.JWTGet(dbUser)
	//		if err != nil {
	//			log.Println(err)
	//			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//			w.WriteHeader(http.StatusInternalServerError)
	//			return
	//		}
	//	}
	//}
	//
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	//if err := json.NewEncoder(w).Encode(dbUser); err != nil {
	//	log.Println(err)
	//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}

}
