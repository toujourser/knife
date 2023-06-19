package beego_log

import "testing"

func TestBeegoLogger(t *testing.T) {
	logger, err := NewLogger(7, "test.log")
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("test")
	logger.WithField("key", "value").Info("test")
}
