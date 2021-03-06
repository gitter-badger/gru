package resource

import (
	"testing"

	"github.com/coreos/go-systemd/util"
)

func TestService(t *testing.T) {
	if !util.IsRunningSystemd() {
		return
	}

	L := newLuaState()
	defer L.Close()

	const code = `
	svc = service.new("nginx")
	`

	if err := L.DoString(code); err != nil {
		t.Fatal(err)
	}

	svc := luaResource(L, "svc").(*Service)
	errorIfNotEqual(t, "service", svc.Type)
	errorIfNotEqual(t, "nginx", svc.Name)
	errorIfNotEqual(t, StateRunning, svc.State)
	errorIfNotEqual(t, []string{}, svc.After)
	errorIfNotEqual(t, []string{}, svc.Before)
	errorIfNotEqual(t, true, svc.Enable)
}
