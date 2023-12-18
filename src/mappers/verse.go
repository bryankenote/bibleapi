package mappers

import (
	biblev1 "github.com/bryankenote/bibleapi/src/codegen/pb/bible/v1"
	"github.com/bryankenote/bibleapi/src/codegen/sqlc"

	"connectrpc.com/connect"
)

func FromGetChapterRequest(req *connect.Request[biblev1.GetChapterRequest]) *sqlc.GetChapterParams {
	return &sqlc.GetChapterParams{
		Translation: req.Msg.Translation,
		Book:        req.Msg.Book,
		Chapter:     int64(req.Msg.Chapter),
	}
}

func ToVerseDtos(verses []sqlc.Verse) []*biblev1.Verse {
	var versesDto []*biblev1.Verse
	for _, verse := range verses {
		versesDto = append(versesDto, &biblev1.Verse{
			Translation: verse.Translation,
			Book:        verse.Book,
			Chapter:     int32(verse.Chapter),
			Verse:       int32(verse.Verse),
			Content:     verse.Content,
		})
	}
	return versesDto
}
