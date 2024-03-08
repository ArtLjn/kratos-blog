package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/hashicorp/consul/api"
	"kratos-blog/api/v1/blog"
	"kratos-blog/api/v1/comment"
	"kratos-blog/api/v1/user"
	"testing"
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
