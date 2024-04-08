//go:build !windows
// +build !windows

package mssql

import (
	"context"
	"fmt"
	"net"

	"github.com/BruceCatYu/go-mssqldb/msdsn"
)

func dialConnectionUsingNamedPipe(_ context.Context, _ *Connector, p msdsn.Config) (conn net.Conn, err error) {
	return nil, fmt.Errorf(
		"Named pipe protocol (\"%s\") is implemented only for Windows OS",
		p.Protocol)
}
