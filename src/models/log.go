package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	"../api/proto/pb_src"
)

type Log struct {
	Id             int64
	UserId         int64
	Severity       int32
	LogType        string
	Section        string
	Description    string
	HappenedAt     time.Time
	AdditionalData string
}

func QueryGetLogs(limit int32, page int32, query []*pb_src.QueryItem) ([]Log, int64, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic triggered: %d", r)
		}
	}()

	if limit == 0 {
		limit = 20
	}

	if page == 0 {
		page = 1
	}

	var queryString strings.Builder

	queryString.WriteString("id IS NOT NULL")

	if len(query) > 0 {
		for i := range query {
			switch query[i].Param {
			case "user_id":
				queryString.WriteString(" AND user_id=" + query[i].Value)
			case "severity":
				queryString.WriteString(" AND severity=" + query[i].Value)
			case "log_type":
				queryString.WriteString(" AND log_type ILIKE '%" + query[i].Value + "%'")
			case "section":
				queryString.WriteString(" AND section ILIKE '%" + query[i].Value + "%'")
			case "description":
				queryString.WriteString(" AND description ILIKE '%" + query[i].Value + "%'")
			case "before":
				queryString.WriteString(" AND happened_at < to_timestamp(" + query[i].Value + ")")
			case "after":
				queryString.WriteString(" AND happened_at > to_timestamp(" + query[i].Value + ")")
			}
		}
	}

	offset := limit * (page - 1)

	rows, err := db.Query(fmt.Sprintf("SELECT user_id, severity, log_type, section, description, additional_data, happened_at FROM public.logs WHERE %s OFFSET $1 LIMIT $2;", queryString.String()), offset, limit)

	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var count int64

	row := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM public.logs WHERE %s;", queryString.String()))

	if err := row.Scan(&count); err != nil {
		log.Print(err)
	}

	var logs []Log

	for rows.Next() {
		var userId int64
		var severity int32
		var logType string
		var section string
		var description string
		var additionalData string
		var happenedAt time.Time

		err := rows.Scan(&userId, &severity, &logType, &section, &description, &additionalData, &happenedAt)
		if err != nil {
			log.Print(err)
		}

		logs = append(logs, Log{
			UserId:         userId,
			Severity:       severity,
			LogType:        logType,
			Section:        section,
			Description:    description,
			AdditionalData: additionalData,
			HappenedAt:     happenedAt,
		})
	}

	if err := rows.Err(); err != nil {
		log.Print(err)
		return nil, 0, fmt.Errorf("query failed: %s", err)
	}

	return logs, count, nil
}

func QueryCreateLog(logEntry Log) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic triggered: %d", r)
		}
	}()

	prep, err := db.Prepare("INSERT INTO public.logs(user_id, severity, log_type, section, description, additional_data, happened_at) VALUES($1, $2, $3, $4, $5, $6, $7)")

	if err != nil {
		log.Print(err)
		return "", err
	}

	result, err := prep.Exec(logEntry.UserId, logEntry.Severity, logEntry.LogType, logEntry.Section, logEntry.Description, logEntry.AdditionalData, logEntry.HappenedAt)

	if err != nil {
		log.Print(err)
		return "", err
	}

	if _, err := result.RowsAffected(); err != nil {
		log.Print(err)
		return "", err
	}

	return "Log saved successfully", nil
}
