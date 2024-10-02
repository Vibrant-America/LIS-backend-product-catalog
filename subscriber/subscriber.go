package subscriber

import (
	"context"
	"time"

	// "go-micro.dev/v4/util/log"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
)

type User struct {
	UserId     int    `json:"user_id"`
	Email      string `json:"email"`
	PracticeId int    `json:"practice_id"`
	ClinicId   int    `json:"clinic_id"`
	PatientId  int    `json:"patient_id,omitempty"`
	AccessList string `json:"access_list"`
	Iat        int    `json:"iat"`
	Exp        int    `json:"exp"`
}
type AuditLog struct {
	User        User        `json:"user"`
	FeildName   string      `json:"feildName,omitempty"`
	ServiceName string      `json:"serviceName,omitempty"`
	Args        interface{} `json:"args,omitempty"`
	Result      interface{} `json:"result"`
	Date        time.Time   `json:"date"`
}

func SubEv(ctx context.Context, event *AuditLog) error {
	// md, _ := metadata.FromContext(ctx)
	log.Infof("[pubsub] Received event %+v with metadata %+v\n", event)
	// do something with event
	return nil
}

type Sub struct{}

// Method can be of any name
func (s *Sub) Process(ctx context.Context, event *AuditLog) error {
	md, _ := metadata.FromContext(ctx)
	log.Infof("[pubsub.1] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}
