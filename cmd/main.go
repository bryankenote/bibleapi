package main

import (
	"context"
	"flag"

	"log"
	"net/http"

	biblev1 "github.com/bryankenote/bibleapi/codegen/pb/bible/v1"
	"github.com/bryankenote/bibleapi/codegen/pb/bible/v1/biblev1connect"
	"github.com/bryankenote/bibleapi/mappers"

	"github.com/bryankenote/bibleapi/db"

	"connectrpc.com/connect"
	"github.com/joho/godotenv"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type BibleServer struct{}

func (s *BibleServer) GetChapter(
	ctx context.Context,
	req *connect.Request[biblev1.GetChapterRequest],
) (*connect.Response[biblev1.GetChapterResponse], error) {
	log.Println("Request headers: ", req.Header())

	verses, err := db.Instance.GetChapter(ctx, *mappers.FromGetChapterRequest(req))
	if err != nil {
		log.Fatalf("Unable to load verse: %s", err.Error())
	}

	res := connect.NewResponse(&biblev1.GetChapterResponse{
		Verses: mappers.ToVerseDtos(verses),
	})
	res.Header().Set("Bible-Version", "v1")
	return res, nil
}

func init() {
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env file")
	}

	db.ConnectToDB()
}

func main() {
	bibleServer := &BibleServer{}
	mux := http.NewServeMux()
	path, handler := biblev1connect.NewBibleServiceHandler(bibleServer)
	mux.Handle(path, handler)
	http.ListenAndServe("localhost:8000", h2c.NewHandler(mux, &http2.Server{}))
}
