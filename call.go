package kitex_demo

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/notedemo"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/notedemo/noteservice"
	"github.com/cloudwego/kitex/client"
)

var (
	noteClient noteservice.Client
	err        error
)

func init() {

	// 生成 resolver 初始化的代码
	noteClient, err = noteservice.NewClient(
		"demo note",
		client.WithHostPorts("127.0.0.1:8888"), // resolver
	)
	if err != nil {
		panic(err)
	}
}

// CreateNote create note info
func CreateNote(ctx context.Context, req *notedemo.CreateNoteRequest) (*notedemo.CreateNoteResponse, error) {
	resp, err := noteClient.CreateNote(ctx, req)
	if err != nil {
		hlog.CtxErrorf(ctx, "err = %s", err.Error())
		return nil, err
	}
	return resp, nil
}

// QueryNotes query list of note info
func QueryNotes(ctx context.Context, req *notedemo.QueryNoteRequest) (*notedemo.QueryNoteResponse, error) {
	resp, err := noteClient.QueryNote(ctx, req)
	if err != nil {
		hlog.CtxErrorf(ctx, "err = %s", err.Error())
		return nil, err
	}
	return resp, nil
}
