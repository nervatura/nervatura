package nervatura

var testData = struct {
	version, hashTable, tokenKey, apiKey,
	validPass, validHash,
	adminToken, userToken, customerToken string
	getNervaStore func(config IM, user *User, customer IM) *NervaStore
}{
	version:   "test",
	hashTable: "ref17890714",
	tokenKey:  "TEST_TOKEN_KEY",
	apiKey:    "TEST_API_KEY",
	validPass: "TestPassword",
	validHash: "$argon2id$v=19$m=65536,t=1,p=2$h7cS5ggipbI1O3m4i+ItAA$yhlMIvM74BvuCUTyfHb+ABx2DkwmlX1S7xyAGVnH7S4",
	getNervaStore: func(config IM, user *User, customer IM) (nstore *NervaStore) {
		nstore = New(&testDriver{Config: config}, config)
		nstore.User = user
		nstore.Customer = customer
		return nstore
	},
}

type testDriver struct {
	Config IM
}

func (ds *testDriver) Properties() struct {
	SQL, Transaction bool
} {
	return struct{ SQL, Transaction bool }{SQL: true, Transaction: true}
}
func (ds *testDriver) Connection() struct {
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
	return struct {
		Alias     string
		Connected bool
		Engine    string
	}{
		Alias:     "test",
		Connected: true,
		Engine:    "test",
	}
}
func (ds *testDriver) CreateConnection(alias, connStr string) error {
	if value, found := ds.Config["CreateConnection"].(func() error); found {
		return value()
	}
	return nil
}
func (ds *testDriver) CloseConnection() error {
	return nil
}
func (ds *testDriver) CreateDatabase(logData []SM) ([]SM, error) {
	if value, found := ds.Config["CreateDatabase"].(func() ([]SM, error)); found {
		return value()
	}
	return logData, nil
}
func (ds *testDriver) CheckHashtable(hashtable string) error {
	if value, found := ds.Config["CheckHashtable"].(func() error); found {
		return value()
	}
	return nil
}
func (ds *testDriver) UpdateHashtable(hashtable, refname, value string) error {
	if value, found := ds.Config["UpdateHashtable"].(func() error); found {
		return value()
	}
	return nil
}
func (ds *testDriver) Query(queries []Query, transaction interface{}) ([]IM, error) {
	if value, found := ds.Config["Query"].(func([]Query) ([]IM, error)); found {
		return value(queries)
	}
	return []IM{}, nil
}
func (ds *testDriver) QuerySQL(sqlString string, params []interface{}, transaction interface{}) ([]IM, error) {
	if value, found := ds.Config["QuerySQL"].(func(sqlString string) ([]IM, error)); found {
		return value(sqlString)
	}
	return []IM{}, nil
}
func (ds *testDriver) QueryKey(options IM, transaction interface{}) ([]IM, error) {
	if value, found := ds.Config["QueryKey"].(func(IM) ([]IM, error)); found {
		return value(options)
	}
	return []IM{}, nil
}
func (ds *testDriver) Update(options Update) (int64, error) {
	if value, found := ds.Config["Update"].(func(Update) (int64, error)); found {
		return value(options)
	}
	return 0, nil
}
func (ds *testDriver) BeginTransaction() (interface{}, error) {
	if value, found := ds.Config["BeginTransaction"].(func() (interface{}, error)); found {
		return value()
	}
	return IM{}, nil
}
func (ds *testDriver) CommitTransaction(trans interface{}) error {
	if value, found := ds.Config["CommitTransaction"].(func() error); found {
		return value()
	}
	return nil
}
func (ds *testDriver) RollbackTransaction(trans interface{}) error {
	if value, found := ds.Config["RollbackTransaction"].(func() error); found {
		return value()
	}
	return nil
}
