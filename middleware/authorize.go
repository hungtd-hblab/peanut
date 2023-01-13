package middleware

import (
	"peanut/pkg/apierrors"
	"peanut/pkg/jwt"
	"peanut/pkg/response"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func Authorize(obj string, act string, e *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, err := jwt.ExtractUserID(ctx)
		if err != nil {
			err = apierrors.NewErrorf(apierrors.Unauthorized, "Unauthorized")
			response.Error(ctx, err)
			ctx.Abort()
			return
		}

		user := strconv.Itoa(userID)
		if ok, _ := e.Enforce(user, obj, act); ok {
			ctx.Next()
			return
		}

		err = apierrors.NewErrorf(apierrors.Unauthorized, "Do not have permission")
		response.Error(ctx, err)
		ctx.Abort()
	}
}
