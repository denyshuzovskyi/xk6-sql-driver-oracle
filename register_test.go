package oracle

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/grafana/xk6-sql/sqltest"
	"github.com/stretchr/testify/require"
	"net"
	"runtime"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

//go:embed testdata/script.js
var script string

func TestIntegration(t *testing.T) { //nolint:paralleltest
	if testing.Short() {
		t.Skip()
	}

	if runtime.GOOS != "linux" {
		t.Skip("Works only on Linux (Testcontainers)")
	}

	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "gvenzl/oracle-free:23",
		ExposedPorts: []string{"1521/tcp"},
		Env: map[string]string{
			"ORACLE_PASSWORD":   "oracle",
			"APP_USER":          "oracle",
			"APP_USER_PASSWORD": "oracle",
		},
		WaitingFor: wait.ForLog("DATABASE IS READY TO USE!").WithStartupTimeout(time.Minute * 5),
	}

	ctrReq := testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	}

	ctr, err := testcontainers.GenericContainer(ctx, ctrReq)
	require.NoError(t, err)
	defer func() { require.NoError(t, ctr.Terminate(ctx)) }()

	host, err := ctr.Host(ctx)
	require.NoError(t, err)

	mappedPort, err := ctr.MappedPort(ctx, "1521/tcp")
	require.NoError(t, err)

	dsn := fmt.Sprintf("oracle://oracle:oracle@%s/FREEPDB1", net.JoinHostPort(host, mappedPort.Port()))

	sqltest.RunScript(t, "oracle", dsn, script)
}
