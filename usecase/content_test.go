package usecase_test

import (
	"fmt"
	"peanut/domain"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Content", func() {
	var (
		content   domain.Content
		filePath  string
		playTimeS string
		playTime  time.Time
	)

	BeforeEach(func() {
		filePath = "/public/uploads/0d041b1e-2332-41fd-9b26-be75ad62ca67.png"
		playTimeS = "2022-01-02 15:20:20"
		playTime, _ = time.Parse("2006-01-02 15:04:05", playTimeS)
		content = domain.Content{
			Title:       "Title of Content",
			Description: "Description of Content",
			PlayTime:    playTime,
			Resolution:  1080,
			AspectRatio: "16:9",
			FilePath:    filePath,
		}
	})

	Describe("API Create Content", func() {
		Context("with new content", func() {
			It("should be success", func() {
				// prepare
				filePath = "/public/uploads/0d041b1e-2332-41fd-9b26-be75ad62ca67.png"
				req := domain.CreateContentReq{
					Title:       "Title of Content",
					Description: "Description of Content",
					PlayTime:    playTimeS,
					Resolution:  1080,
					AspectRatio: "16:9",
				}
				contentRepo.EXPECT().CreateContent(content).Return(&content, nil)
				// do
				resp, err := contentUc.CreateContent(ctx, req, filePath)
				// check
				Expect(err).To(BeNil())
				Expect(resp).NotTo(BeNil())
			})
		})

		Context("with database error content", func() {
			It("should be failed", func() {
				// prepare
				filePath = "/public/uploads/0d041b1e-2332-41fd-9b26-be75ad62ca67.png"
				req := domain.CreateContentReq{
					Title:       "Title of Content",
					Description: "Description of Content",
					PlayTime:    playTimeS,
					Resolution:  1080,
					AspectRatio: "16:9",
				}
				contentRepo.EXPECT().CreateContent(content).Return(nil, fmt.Errorf("database error"))
				// do
				resp, err := contentUc.CreateContent(ctx, req, filePath)
				// check
				Expect(err).NotTo(BeNil())
				Expect(resp).To(BeZero())
			})
		})
	})
})
