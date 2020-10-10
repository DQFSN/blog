package microimp

import (
	mpb "blog/api/micro"
	db "blog/server/db"
	"blog/server/model"
	"fmt"
	"golang.org/x/net/context"
	"log"
)

type BlogServer struct {
}

func (bs BlogServer) PublishBlog(ctx context.Context, in *mpb.PublishRequest, out *mpb.PublishReply) error {

	if len(in.Title) > 0 && len(in.Author) > 0 {
		mysqlDB := db.DB()
		err := mysqlDB.Create(&model.Blog{Title: in.Title, Content: in.Content, Author: in.Author}).Error
		if err != nil {
			out = &mpb.PublishReply{
				Status: fmt.Sprintf("publish blog : %s", err),
			}
			return  err
		}

		out =  &mpb.PublishReply{
			Status: fmt.Sprintf("publish ok : %s", in.Title),
		}
		return nil
	}

	log.Fatal("publish failed, title or author can not be empty")
	out = &mpb.PublishReply{
		Status: fmt.Sprintln("publish failed, title or author can not be empty"),
	}
	return nil
}

func (bs BlogServer) GetBlogs(ctx context.Context, in *mpb.BlogsRequest, out *mpb.BlogsReply)  error {
	mysqlDB := db.DB()

	var blogs []*mpb.Blog
	fmt.Printf("author-->%s\n", in.Author)
	if len(in.Author) > 0 {
		mysqlDB.Where(mpb.Blog{Author: in.Author}).Find(&blogs)
	} else {
		mysqlDB.Where(mpb.Blog{}).Find(&blogs)
	}

	out = &mpb.BlogsReply{Blogs: blogs}
	return  nil

}

func (bs BlogServer) ModifyBlog(ctx context.Context, in *mpb.ModifyBlogRequest, out *mpb.ModifyBlogReply) (err error) {
	mysqlDB := db.DB()

	err = mysqlDB.Where(mpb.ModifyBlogRequest{Id: in.Id}).Error

	if err != nil {
		out.Status = fmt.Sprintf("update blog err ：%v", err)
		return nil
	}
	out.Status = "update blog Ok"
	return nil
}
