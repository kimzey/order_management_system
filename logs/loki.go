package logger

//import (
//	"context"
//	"github.com/grafana/loki-client-go/loki"
//	"github.com/grafana/loki-client-go/loki/util/urlutil"
//	"github.com/sirupsen/logrus"
//	"net/url"
//	"time"
//)
//
//var (
//	LokiClient *loki.Client
//)
//
//func InitLokiClient() {
//	lokiURL, err := url.Parse("http://localhost:3100/loki/api/v1/push")
//	if err != nil {
//		logrus.Fatalf("Failed to parse Loki URL: %v", err)
//	}
//
//	lokiURLValue := urlutil.NewURLValue(lokiURL) // Convert url.URL to urlutil.URLValue
//
//	cfg := loki.Config{
//		URL:       &lokiURLValue,
//		BatchWait: 1 * time.Second,
//		BatchSize: 10000,
//	}
//
//	LokiClient, err = loki.New(cfg)
//	if err != nil {
//		logrus.Fatalf("Unable to create Loki client: %v", err)
//	}
//}
//
//func LogToLoki(level logrus.Level, message string, fields logrus.Fields) {
//	entry := logrus.NewEntry(logrus.StandardLogger()).WithFields(fields).WithTime(time.Now())
//	entry.Level = level
//	entry.Message = message
//
//	ctx := context.Background()
//	err := LokiClient.Handle(ctx, loki.Entry{
//		Labels:    loki.LabelSet{"job": "order_management", "level": level.String()},
//		Timestamp: entry.Time,
//		Line:      entry.String(),
//	})
//	if err != nil {
//		logrus.Errorf("Unable to send log to Loki: %v", err)
//	}
//}
