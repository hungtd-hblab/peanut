package usecase_test

import (
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"peanut/domain"
)

var _ = Describe("User", func() {
	var (
		createdU domain.User
		u        domain.SignupReq
	)
	BeforeEach(func() {
		u = domain.SignupReq{
			Username: "hung",
			Email:    "hung@example.com",
		}

		createdU = domain.User{}
	})

	Describe("API Sign up", func() {
		Context("with existed username", func() {
			It("should be error", func() {
				// TODO: fill in your test in this case
			})
		})
		Context("with new user", func() {
			It("should be success", func() {
				// prepare
				userRepo.EXPECT().GetUserByUsername(u.Username).Return(nil, nil)
				userRepo.EXPECT().CreateUser(gomock.All()).Return(&createdU, nil)
				// do
				resp, err := userUc.SignUp(ctx, u)
				// check
				Expect(err).To(BeNil())
				Expect(resp).NotTo(BeNil())
			})
		})
		Context("with database error response", func() {
			It("should be err", func() {
				// prepare
				userRepo.EXPECT().GetUserByUsername(u.Username).Return(nil, nil)
				userRepo.EXPECT().CreateUser(gomock.Any()).
					Return(nil, fmt.Errorf("database error"))
				// do
				resp, err := userUc.SignUp(ctx, u)
				// check
				Expect(err).NotTo(BeNil())
				Expect(resp).To(BeZero())
			})
		})
	})

	Describe("API Get user", func() {
		Context("with existed id", func() {
			It("should be return user", func() {
				// Do something
				Expect(u).NotTo(Equal(createdU))
			})
		})
	})
})
