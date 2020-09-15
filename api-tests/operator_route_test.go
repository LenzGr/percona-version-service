package api_tests

import (
	"testing"
	"time"

	"github.com/Percona-Lab/percona-version-service/client/version_service"
	"github.com/stretchr/testify/assert"
)

func TestOperatorRouteShouldReturnRigthOperatorVersion(t *testing.T) {
	cli := cli()

	cases := []struct {
		product string
		version string
	}{
		{"pxc-operator", "1.4.0"},
		{"pxc-operator", "1.5.0"},
		{"psmdb-operator", "1.5.0"},
	}

	for _, c := range cases {
		params := &version_service.VersionServiceOperatorParams{
			OperatorVersion: c.version,
			Product:         c.product,
		}
		params.WithTimeout(2 * time.Second)

		resp, err := cli.VersionService.VersionServiceOperator(params)
		assert.NoError(t, err)

		v := getVersion(resp.Payload.Versions[0].Matrix.Operator)
		assert.Equal(t, c.version, v)
		assert.Equal(t, c.version, resp.Payload.Versions[0].Operator)
		assert.Equal(t, c.product, resp.Payload.Versions[0].Product)
	}
}

func TestOperatorRoutePxcShouldReturnNotEmptyResponses(t *testing.T) {
	cli := cli()

	cases := []struct {
		product string
		version string
	}{
		{"pxc-operator", "1.4.0"},
		{"pxc-operator", "1.5.0"},
	}

	for _, c := range cases {
		params := &version_service.VersionServiceOperatorParams{
			OperatorVersion: c.version,
			Product:         c.product,
		}
		params.WithTimeout(2 * time.Second)

		resp, err := cli.VersionService.VersionServiceOperator(params)
		assert.NoError(t, err)

		assert.Len(t, resp.Payload.Versions, 1)
		assert.Len(t, resp.Payload.Versions[0].Matrix.Operator, 1)
		assert.Greater(t, len(resp.Payload.Versions[0].Matrix.Pxc), 0)
		assert.Greater(t, len(resp.Payload.Versions[0].Matrix.Pmm), 0)
		assert.Greater(t, len(resp.Payload.Versions[0].Matrix.Proxysql), 0)
		assert.Greater(t, len(resp.Payload.Versions[0].Matrix.Backup), 0)

		if c.version >= "1.5.0" {
			assert.Greater(t, len(resp.Payload.Versions[0].Matrix.Haproxy), 0)
		} else {
			assert.Equal(t, 0, len(resp.Payload.Versions[0].Matrix.Haproxy))
		}
	}
}

func TestOperatorRoutePsmdbShouldReturnNotEmptyResponses(t *testing.T) {
	cli := cli()

	cases := []struct {
		product string
		version string
	}{
		{"psmdb-operator", "1.5.0"},
	}

	for _, c := range cases {
		params := &version_service.VersionServiceOperatorParams{
			OperatorVersion: c.version,
			Product:         c.product,
		}
		params.WithTimeout(2 * time.Second)

		resp, err := cli.VersionService.VersionServiceOperator(params)
		assert.NoError(t, err)

		assert.Len(t, resp.Payload.Versions, 1)
		assert.Len(t, resp.Payload.Versions[0].Matrix.Operator, 1)
		assert.Greater(t, len(resp.Payload.Versions[0].Matrix.Mongod), 0)
		assert.Greater(t, len(resp.Payload.Versions[0].Matrix.Pmm), 0)
		assert.Greater(t, len(resp.Payload.Versions[0].Matrix.Backup), 0)
	}
}
