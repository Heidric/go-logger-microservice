package test

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"

	"./pb_src"

	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func TestMain(m *testing.M) {
	setup()

	time.Sleep(5 * time.Second)

	code := m.Run()

	cleanup()
	os.Exit(code)
}

func setup() {
	if os.Getenv("ENV") != "test" {
		err := os.Setenv("ENV", "test")

		if err != nil {
			log.Fatalln("ENV global variable is not set to 'test' and app has failed to set it itself")
		}
	}

	if os.Getenv("MAIN_PORT") == "" {
		err := os.Setenv("MAIN_PORT", "7777")

		if err != nil {
			log.Fatalln("MAIN_PORT global variable is not set and app has failed to set it itself")
		}
	}

	err := os.Chdir("../src/")

	err = exec.Command("/bin/sh", "build.sh").Run()

	if err != nil {
		log.Fatalln("Build failed:", err)
	}

	err = os.Chdir("../devops/")

	err = exec.Command("/bin/sh", "run.sh").Run()

	if err != nil {
		log.Fatalln("Start failed:", err)
	}

	err = os.Chdir("../test/")

	if err != nil {
		log.Fatalln("Failed to return to test folder:", err)
	}
}

func cleanup() {
	err := exec.Command("/bin/sh", "cleanup.sh").Run()

	if err != nil {
		log.Fatalln("Cleanup failed:", err)
	}
}

func TestCreateLogEntry(t *testing.T) {
	log.Println("This should successfully create a log entry")

	conn, client := connect()

	defer conn.Close()

	timeNow := time.Now()

	happenedAt, _ := ptypes.TimestampProto(timeNow)

	_, err := client.CreateEntry(context.Background(), &pb_src.LogEntry{
		UserId:         1,
		Severity:       1,
		Section:        "Main",
		LogType:        "New",
		Description:    "Qwerty",
		AdditionalData: "{ info: 'additional data' }",
		HappenedAt:     happenedAt,
	})

	if err != nil {
		t.Errorf("Error when calling CreateEntry: %s", err)
	}
}

func TestGetLogEntriesWithoutQuery(t *testing.T) {
	log.Println("This should successfully get a list of log entries")

	conn, client := connect()

	defer conn.Close()

	response, err := client.GetEntries(context.Background(), &pb_src.EntriesRequest{
		Limit: 20,
		Page:  1,
	})

	if err != nil {
		t.Errorf("Error when calling CreateEntry: %s", err)
		return
	}

	if response == nil {
		t.Error("Expected response to not be nil")
		return
	}

	if len(response.Entries) < 1 || response.Count < 1 {
		t.Error("Expected response to contain data")
	}
}

func TestGetLogEntriesWithQuery(t *testing.T) {
	log.Println("This also should successfully get a list of log entries")

	conn, client := connect()

	defer conn.Close()

	userIdQuery := pb_src.QueryItem{
		Param: "UserId",
		Value: "1",
	}

	query := []*pb_src.QueryItem{
		&userIdQuery,
	}

	response, err := client.GetEntries(context.Background(), &pb_src.EntriesRequest{
		Limit: 20,
		Page:  1,
		Query: query,
	})

	if err != nil {
		t.Errorf("Error when calling CreateEntry: %s", err)
		return
	}

	if response == nil {
		t.Error("Expected response to not be nil")
		return
	}

	if len(response.Entries) < 1 || response.Count < 1 {
		t.Error("Expected response to contain data")
	}
}

func TestDeleteLogEntry(t *testing.T) {
	log.Println("This should successfully delete an entry") // TODO: implement this
}

func connect() (*grpc.ClientConn, pb_src.LoggerServiceClient) {
	conn, err := grpc.Dial(fmt.Sprintf(":%s", os.Getenv("MAIN_PORT")), grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %s", err)
	}

	client := pb_src.NewLoggerServiceClient(conn)

	return conn, client
}
