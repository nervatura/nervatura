package nervatura

// TimeLayout DateTime format
const TimeLayout = "2006-01-02 15:04:05"

// IM is a map[string]interface{} type short alias
type IM = map[string]interface{}

// SM is a map[string]string type short alias
type SM = map[string]string

// IL is a []interface{} type short alias
type IL = []interface{}

// IList is a []interface{} string const
const IList = "[]interface{}"

// SL is a []string type short alias
type SL = []string

// DataDriver a general database interface
type DataDriver interface {
	Properties() struct {
		SQL, Transaction bool
	} //DataDriver features
	Connection() struct {
		Alias     string
		Connected bool
		Engine    string
	} //Returns the database connection
	CreateConnection(string, string) error                                                  //Create a new database connection
	CreateDatabase(logData []SM) ([]SM, error)                                              //Create a Nervatura Database
	CheckHashtable(hashtable string) error                                                  //Check/create a password ref. table
	UpdateHashtable(hashtable, refname, value string) error                                 //Set a password
	Query(queries []Query, transaction interface{}) ([]IM, error)                           //Query is a basic nosql friendly queries the database
	QuerySQL(sqlString string, params []interface{}, transaction interface{}) ([]IM, error) //Executes a SQL query
	QueryKey(options IM, transaction interface{}) ([]IM, error)                             //Complex data queries
	Update(options Update) (int64, error)                                                   //Update is a basic nosql friendly update/insert/delete and returns the update/insert id
	BeginTransaction() (interface{}, error)                                                 //Begins a transaction and returns an it
	CommitTransaction(trans interface{}) error                                              //Commit a transaction
	RollbackTransaction(trans interface{}) error                                            //Rollback a transaction
	CloseConnection() error
}

// Filter query filter type
type Filter struct {
	Or    bool   // and (False) or (True)
	Field string //Fieldname and alias
	Comp  string //==,!=,<,<=,>,>=,in,is
	Value interface{}
}

// Query data desc. type
type Query struct {
	Fields  []string //Returns fields
	From    string   //Table or doc. name and alias
	Filters []Filter
	Filter  string //filter string (eg. "id=1 and field='value'")
	OrderBy []string
}

// Update data desc. type
type Update struct {
	Values IM
	Model  string
	IDKey  int64       //Update, delete or insert exec
	Trans  interface{} //Transaction
}

// User - Nervatura user properties
type User struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	Empnumber  string `json:"empnumber"`
	Usergroup  int64  `json:"usergroup"`
	Scope      string `json:"scope"`
	Department string `json:"department,omitempty"`
}
