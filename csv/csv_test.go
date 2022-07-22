package csv

import (
	"context"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	NewCsvReader("./template").Use("robot_team.csv", "robot_name.csv").Loading(ctx, 1, false)
	time.Sleep(6 * time.Second)
}
