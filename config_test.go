package gop

import (
	"os"
	"testing"
	"time"

	"github.com/cocoonlife/testify/assert"
)

func TestHandleFeatureOverrides(t *testing.T) {
	a := assert.New(t)

	app := InitCmd("cocoon", "gop_config_test")
	f, err := os.Open("testdata/ini_overrides.conf")
	a.NoError(err)
	defer f.Close()

	app.Cfg.TransientOverride("gop", "log_level", "info")
	app.Cfg.TransientOverride("gop", "listen_addr", "localhost:1732")
	app.Cfg.TransientOverride("brain", "alert_cooldown", "10m")
	app.Cfg.TransientOverride("brain", "alert_cooldown_maximum", "15m")

	err = app.Cfg.HandleFeatureOverrides(f, 2*time.Second)
	a.NoError(err)

	v, ok := app.Cfg.Get("gop", "log_level", "")
	a.True(ok, "log_level does exist")
	a.Equal(v, "debug", "correct override")

	v, ok = app.Cfg.Get("gop", "listen_addr", "")
	a.True(ok, "listen_addr does exist")
	a.Equal(v, "192.168.1.1:8080", "correct listen_addr")

	v, ok = app.Cfg.Get("brain", "alert_cooldown", "")
	a.True(ok, "alert_cooldown does exist")
	a.Equal(v, "0m", "correct alert_cooldown")

	v, ok = app.Cfg.Get("brain", "alert_cooldown_maximum", "")
	a.True(ok, "alert_cooldown_maximum does exist")
	a.Equal(v, "5m", "correct alert_cooldown_maximum")

	time.Sleep(5 * time.Second)

	v, ok = app.Cfg.Get("gop", "log_level", "")
	a.True(ok, "log_level does exist")
	a.Equal(v, "info", "correct override")

	v, ok = app.Cfg.Get("gop", "listen_addr", "")
	a.True(ok, "listen_addr does exist")
	a.Equal(v, "localhost:1732", "correct listen_addr")

	v, ok = app.Cfg.Get("brain", "alert_cooldown", "")
	a.True(ok, "alert_cooldown does exist")
	a.Equal(v, "10m", "correct alert_cooldown")

	v, ok = app.Cfg.Get("brain", "alert_cooldown_maximum", "")
	a.True(ok, "alert_cooldown_maximum does exist")
	a.Equal(v, "15m", "correct alert_cooldown_maximum")

}

func TestHandleFeatureOverridesBroken(t *testing.T) {
	a := assert.New(t)

	app := InitCmd("cocoon", "gop_config_test")
	f, err := os.Open("testdata/ini_overrides_broken.conf")
	a.NoError(err)
	defer f.Close()

	err = app.Cfg.HandleFeatureOverrides(f, 1*time.Hour)
	a.Error(err)
}
