/* Nervatura test database driver
 */
package driver

import (
	"io"
	"net"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

type TestDriver struct {
	Config map[string]interface{}
}

func (ds *TestDriver) Properties() struct {
	SQL, Transaction bool
} {
	return struct{ SQL, Transaction bool }{SQL: true, Transaction: true}
}

func (ds *TestDriver) Connection() struct {
	Alias     string
	Connected bool
	Engine    string
} {
	if value, found := ds.Config["Connection"].(func() struct {
		Alias     string
		Connected bool
		Engine    string
	}); found {
		return value()
	}
	alias := cu.ToString(ds.Config["alias"], "test")
	engine := cu.ToString(ds.Config["engine"], "test")
	return struct {
		Alias     string
		Connected bool
		Engine    string
	}{
		Alias:     alias,
		Connected: true,
		Engine:    engine,
	}
}

func (ds *TestDriver) CreateConnection(alias, connStr string) error {
	if value, found := ds.Config["CreateConnection"].(func() error); found {
		return value()
	}
	return nil
}

func (ds *TestDriver) CloseConnection() error {
	return nil
}

func (ds *TestDriver) Query(queries []md.Query, transaction interface{}) ([]map[string]interface{}, error) {
	if value, found := ds.Config["Query"].(func([]md.Query) ([]map[string]interface{}, error)); found {
		return value(queries)
	}
	return []map[string]interface{}{}, nil
}

func (ds *TestDriver) QuerySQL(sqlString string, params []interface{}, transaction interface{}) ([]map[string]interface{}, error) {
	if value, found := ds.Config["QuerySQL"].(func(sqlString string) ([]map[string]interface{}, error)); found {
		return value(sqlString)
	}
	return []map[string]interface{}{}, nil
}

func (ds *TestDriver) Update(options md.Update) (int64, error) {
	if value, found := ds.Config["Update"].(func(md.Update) (int64, error)); found {
		return value(options)
	}
	return 0, nil
}

func (ds *TestDriver) UpdateSQL(sqlString string, transaction interface{}) error {
	if value, found := ds.Config["UpdateSQL"].(func(sqlString string, transaction interface{}) error); found {
		return value(sqlString, transaction)
	}
	return nil
}

func (ds *TestDriver) BeginTransaction() (interface{}, error) {
	if value, found := ds.Config["BeginTransaction"].(func() (interface{}, error)); found {
		return value()
	}
	return map[string]interface{}{}, nil
}

func (ds *TestDriver) CommitTransaction(trans interface{}) error {
	if value, found := ds.Config["CommitTransaction"].(func() error); found {
		return value()
	}
	return nil
}

func (ds *TestDriver) RollbackTransaction(trans interface{}) error {
	if value, found := ds.Config["RollbackTransaction"].(func() error); found {
		return value()
	}
	return nil
}

// TestConn implements net.Conn interface for testing
type TestConn struct {
	net.Conn
	readData  []byte
	writeData []byte
}

func NewTestConn() *TestConn {
	return &TestConn{
		readData:  []byte{},
		writeData: []byte{},
	}
}

func (m *TestConn) Read(b []byte) (n int, err error) {
	if len(m.readData) == 0 {
		return 0, io.EOF
	}
	n = copy(b, m.readData)
	m.readData = m.readData[n:]
	return n, nil
}

func (m *TestConn) Write(b []byte) (n int, err error) {
	m.writeData = append(m.writeData, b...)
	return len(b), nil
}

func (m *TestConn) Close() error                       { return nil }
func (m *TestConn) LocalAddr() net.Addr                { return &net.TCPAddr{} }
func (m *TestConn) RemoteAddr() net.Addr               { return &net.TCPAddr{} }
func (m *TestConn) SetDeadline(t time.Time) error      { return nil }
func (m *TestConn) SetReadDeadline(t time.Time) error  { return nil }
func (m *TestConn) SetWriteDeadline(t time.Time) error { return nil }
