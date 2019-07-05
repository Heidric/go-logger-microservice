package logger_client

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"

	"./pb_src"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func example() {
	timeNow := time.Now()

	happenedAt, _ := ptypes.TimestampProto(timeNow)

	result, err := CreateLogEntry(&pb_src.LogEntry{
		UserId:         1,
		Severity:       1,
		Section:        "Main",
		LogType:        "New",
		Description:    "Qwerty",
		AdditionalData: "{ info: 'additional data' }",
		HappenedAt:     happenedAt,
	})

	if err != nil {
		log.Println(err)
	}

	if result != "" {
		log.Println(result)
	}

	userIdQuery := pb_src.QueryItem{
		Param: "UserId",
		Value: "1",
	}

	query := []*pb_src.QueryItem{
		&userIdQuery,
	}

	entries, count, err := GetLogEntriesList(20, 1, query)

	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Entries: %+v, count: %d", entries, count)
	}
}

func CreateLogEntry(entry *pb_src.LogEntry) (string, error) {
	conn, client := connect()

	defer conn.Close()

	response, err := client.CreateEntry(context.Background(), entry)

	if err != nil {
		log.Printf("Failed to create entry with the following error: %s", err)
		return "", err
	}

	return response.Result, nil
}

func GetLogEntriesList(limit int32, page int32, query []*pb_src.QueryItem) (entries []*pb_src.LogEntry, count int64, err error) {
	conn, client := connect()

	defer conn.Close()

	response, err := client.GetEntries(context.Background(), &pb_src.EntriesRequest{
		Limit: limit,
		Page:  page,
		Query: query,
	})

	if err != nil {
		log.Printf("Failed to get entries: %s", err)
		return nil, 0, err
	}

	return response.Entries, response.Count, nil
}

func connect() (*grpc.ClientConn, pb_src.LoggerServiceClient) {
	conn, err := grpc.Dial(fmt.Sprintf(":%s", os.Getenv("MAIN_PORT")), grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %s", err)
	}

	client := pb_src.NewLoggerServiceClient(conn)

	return conn, client
}
