package middleware

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"peanut/pkg/apierrors"
	"peanut/pkg/array"
	"peanut/pkg/jwt"
	"peanut/pkg/response"

	"github.com/gin-gonic/gin"
)

var roles Roles

type Roles struct {
	Resources map[string]Permission `json:"resources"`
}

type Permission struct {
	Get  []string `json:"get"`
	Post []string `json:"post"`
	Del  []string `json:"del"`
}

func init() {
	loadRoles()
}

func loadRoles() {
	jsonFile, err := os.Open("./config/roles.json")
	if err != nil {
		log.Printf("Load role config file err: %v", err.Error())
	}
	defer jsonFile.Close()

	byteRole, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteRole, &roles)
}

func Authorize(act, resr string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtRoles, err := jwt.ExtractRoles(ctx)
		if err != nil || jwtRoles == nil {
			err = apierrors.NewErrorf(apierrors.Unauthorized, "Unauthorized")
			response.Error(ctx, err)
			ctx.Abort()
			return
		}

		for _, r := range jwtRoles {
			if HasPermission(r, act, resr) {
				ctx.Next()
				return
			}
		}

		err = apierrors.NewErrorf(apierrors.Unauthorized, "Do not have permission")
		response.Error(ctx, err)
		ctx.Abort()
	}
}

func HasPermission(sub, act, obj string) bool {
	resr := roles.Resources[obj]

	var perm []string
	switch act {
	case "get":
		perm = resr.Get
	case "post":
		perm = resr.Post
	case "del":
		perm = resr.Del
	}

	return array.IsIn(sub, perm)
}
