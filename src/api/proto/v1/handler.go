package v1

import (
	"log"

	"github.com/golang/protobuf/ptypes"

	"../../../models"
	"../pb_src"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) CreateEntry(ctx context.Context, in *pb_src.LogEntry) (*pb_src.LogCreationResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic triggered: %d", r)
		}
	}()

	happenedAt, _ := ptypes.Timestamp(in.HappenedAt)

	entry := models.Log{
		UserId:         in.UserId,
		Severity:       in.Severity,
		LogType:        in.LogType,
		Section:        in.Section,
		Description:    in.Description,
		AdditionalData: in.AdditionalData,
		HappenedAt:     happenedAt,
	}

	result, err := models.QueryCreateLog(entry)

	if err != nil {
		return nil, err
	}

	return &pb_src.LogCreationResponse{Result: result, Error: ""}, nil
}

func (s *Server) GetEntries(ctx context.Context, in *pb_src.EntriesRequest) (*pb_src.EntriesResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic triggered: %d", r)
		}
	}()

	logs, count, err := models.QueryGetLogs(in.Limit, in.Page, in.Query)
	if err != nil {
		return nil, err
	}

	data := make([]*pb_src.LogEntry, len(logs))

	for i := range data {
		happenedAt, _ := ptypes.TimestampProto(logs[i].HappenedAt)

		data[i] = &pb_src.LogEntry{
			UserId:         logs[i].UserId,
			Severity:       logs[i].Severity,
			LogType:        logs[i].LogType,
			Section:        logs[i].Section,
			Description:    logs[i].Description,
			AdditionalData: logs[i].AdditionalData,
			HappenedAt:     happenedAt,
		}
	}

	return &pb_src.EntriesResponse{Entries: data, Count: count}, nil
}
