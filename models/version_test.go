package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDrvAppVersion_CreatedAtUnix(t *testing.T) {
	ts := int64(1000)
	v := DrvAppVersion{CreatedAt: time.Unix(ts, 0)}
	assert.Equal(t, ts, v.CreatedAtUnix())
}

func TestDrvAppVersion_ToProto(t *testing.T) {
	ts := int64(1000)
	v := DrvAppVersion{Version: "1.1", CreatedAt: time.Unix(ts, 0)}
	res := v.ToProto()
	assert.Equal(t, v.Version, res.Version)
	assert.Equal(t, v.CreatedAtUnix(), res.CreatedAt)
}
