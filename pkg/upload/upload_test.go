package upload

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
	_ = &UploadRepo{
		UploadPath: "E:\\HongDou-Go-Blog\\kratos-blog",
		Domain:     "http://localhost:8000",
		MaxSize:    "10 MB",
	}
	file, err := os.Open("./text.txt")
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "file", file)
	_, ok := http.RequestFromServerContext(ctx)
	if !ok {
		panic(errors.New("ctx false"))
	}
}

func TestContext(t *testing.T) {

}
