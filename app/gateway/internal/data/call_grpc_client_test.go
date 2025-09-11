package data

import (
	"context"
	"kratos-blog/api/v1/blog"
	"kratos-blog/api/v1/comment"
	"kratos-blog/api/v1/user"
	"testing"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/hashicorp/consul/api"
	grpcmd "google.golang.org/grpc/metadata"
)

func TestUserClient(t *testing.T) {
	consulCli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	r := consul.New(consulCli)
	client := NewUserServiceClient(r)
	t.Log(client.GetUser(context.Background(), &user.GetUserRequest{
		Name: "ljn",
	}))
}

func TestBlogClient(t *testing.T) {
	consulCli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	r := consul.New(consulCli)
	client := NewBlogServiceClient(r)
	t.Log(client.ListBlog(context.Background(), &blog.ListBlogRequest{}))
}

func TestCommentClient(t *testing.T) {
	consulCli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	r := consul.New(consulCli)
	client := NewCommentServiceClient(r)
	t.Log(client.ExtractParentComments(context.Background(), &comment.ExtractParentCommentsRequest{
		Id: "876",
	}))

}

func TestGetRole(t *testing.T) {
	consulCli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	r := consul.New(consulCli)
	commentClient := NewCommentServiceClient(r)
	ctx := context.Background()
	var md grpcmd.MD
	md.Set("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InRlc3QwMDEiLCJpc3MiOiJBdXRoX1NlcnZlciIsInN1YiI6IkF1dGgiLCJleHAiOjE3MTE2OTI0MDEsIm5iZiI6MTcxMTA4NzYwMiwiaWF0IjoxNzExMDg3NjAxLCJqdGkiOiJYbGZhaTh5YTZJIn0.w1OAMgR7VyNnspCZiWcLn40fuit49JhBuxPFjoekQ8o")
	body, _ := commentClient.AddComment(ctx, &comment.CreateCommentRequest{
		ArticleId:   "876",
		Comment:     "hello",
		CommentAddr: "58.216.213.53",
	})
	t.Log(body)
}
