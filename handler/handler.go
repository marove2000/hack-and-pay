// Hack and Pay.
//
// Pay your shit while hacking shit
//
// Schemes: http
// Version: 0.1
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// SecurityDefinitions:
//   oauth2:
//     type: oauth2
//     authorizationUrl: /oauth2/auth
//     tokenUrl: /oauth2/token
//     in: header
//     flow: password
//     scopes:
//       bla: foo
//
// swagger:meta
package handler

import (
	"github.com/marove2000/hack-and-pay/log"
	"github.com/marove2000/hack-and-pay/repository/ldap"
	"github.com/marove2000/hack-and-pay/repository/sql"
	"github.com/marove2000/hack-and-pay/router"
)

var pkgLogger = log.New("handler")

type Handler struct {
	repo       *sql.Mysql
	ldap       *ldap.LDAP
	authorizer authorizer
}

func New(repo *sql.Mysql, ldap *ldap.LDAP, authorizer authorizer) *Handler {
	return &Handler{
		repo:       repo,
		ldap:       ldap,
		authorizer: authorizer,
	}
}

func (h *Handler) Routes() []*router.Route {
	return []*router.Route{
		// swagger:route GET /v1/user users GetUserIndex
		//
		// Lists all users with their public data.
		//
		// This will show all available pets by default.
		// You can get the pets that are out of stock
		//
		// Consumes:
		// - application/json
		//
		// Produces:
		// - application/json
		//
		// Schemes: http
		//
		// Security:
		//   oauth2: admin, user
		//
		// Responses:
		//   200: UserSlice
		//   400: error
		{
			"GetUserIndex",
			"GET",
			"/v1/user",
			wrap(h.publicUserIndex),
		},
		// swagger:route POST /v1/user/login users userLogin
		//
		// Used for login.
		//
		// Yadda yadda yadda
		//
		// Consumes:
		// - application/json
		//
		// Produces:
		// - application/json
		//
		// Schemes: http
		//
		// Parameters:
		// Body      LoginRequestBody     in:body required "comment woop"
		//
		// Security:
		// oauth2: admin, user
		//
		// Responses:
		//   200: User
		//   400: error
		//	 500: error
		{
			"Login",
			"POST",
			"/v1/login",
			wrap(h.login),
		},
		// swagger:route POST /v1/user users addUser
		//
		// Adds a user.
		//
		// Yadda yadda yadda
		//
		// Consumes:
		// - application/json
		//
		// Produces:
		// - application/json
		//
		// Schemes: http
		//
		// Parameters:
		// Body      AddUserRequestBody     in:body required "comment woop"
		//
		// Security:
		// oauth2: admin, user
		//
		// Responses:
		//   200: UserSlice
		//   400: error
		//	 500: error
		{
			"AddUser",
			"POST",
			"/v1/user",
			wrap(h.signUp),
		},
		{
			"GetUserDetail",
			"GET",
			"/v1/user/:id",
			wrap(h.getUserDetail),
		},
		{
			"GetProductIndex",
			"GET",
			"/v1/product",
			wrap(h.productIndex),
		},
		{
			"GetProductDetail",
			"GET",
			"/v1/product/:sku",
			wrap(h.productDetail),
		},
		{
			"AddTransaction",
			"POST",
			"/v1/user/:id/transaction",
			wrap(h.authorizer.Authorize(h.addTransaction, authTypeAll)),
		},
	}
}

//
//var RoutesIndex = []*router.Route{
//	{
//		"Index",
//		"GET",
//		"/v1",
//		h, //TODO Edit Landing-Page
//	},
//	{
//		"publicUserIndex",
//		"GET",
//		"/v1/user",
//		v1.publicUserIndex,
//	},
//	{
//		"AddUser",
//		"POST",
//		"/v1/user",
//		v1.AddUser,
//	},
//	{
//		"GetPublicUserDetail",
//		"GET",
//		"/v1/user/{id}",
//		v1.GetPublicUserDetail,
//	},
//	{
//		"GetAuthentication",
//		"POST",
//		"/v1/authentication",
//		v1.GetAuthentication,
//	},
//	{
//		"ChangeBalance",
//		"POST",
//		"/v1/user/{id}/transaction",
//		v1.ChangeBalance,
//	},
//	{
//		"PutProduct",
//		"PUT",
//		"/v1/put/product",
//		v1.PutProduct,
//	},
//	{
//		"GetProductIndex",
//		"GET",
//		"/v1/get/product",
//		v1.GetProductIndex,
//	},
//}
