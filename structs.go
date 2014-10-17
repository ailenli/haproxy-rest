package main

import(
	"sync"
)

type Metric struct {

	Name string `json:"name"`
	Value int `json:"value"`
	Timestamp int64 `json:"timestamp"`

}

// Struct to hold the output from the /stats endpoint
type StatsGroup struct {
	Pxname string `json:"pxname"`
	Svname string `json:"svname"`
	Qcur string `json:"qcur"`
	Qmax string `json:"qmax"`
	Scur string `json:"scur"`
	Smax string `json:"smax"`
	Slim string `json:"slim"`
	Stot string `json:"stot"`
	Bin string `json:"bin"`
	Bout string `json:"bout"`
	Dreq string `json:"dreq"`
	Dresp string `json:"dresp"`
	Ereq string `json:"ereq"`
	Econ string `json:"econ"`
	Eresp string `json:"eresp"`
	Wretr string `json:"wretr"`
	Wredis string `json:"wredis"`
	Status string `json:"status"`
	Weight string `json:"weight"`
	Act string `json:"act"`
	Bck string `json:"bck"`
	Chkfail string `json:"chkfail"`
	Chkdown string `json:"chkdown"`
	Lastchg string `json:"lastchg"`
	Downtime string `json:"downtime"`
	Qlimit string `json:"qlimit"`
	Pid string `json:"pid"`
	Iid string `json:"iid"`
	Sid string `json:"sid"`
	Throttle string `json:"throttle"`
	Lbtot string `json:"lbtot"`
	Tracked string `json:"tracked"`
	_Type string `json:"type"`
	Rate string `json:"rate"`
	Rate_lim string `json:"rate_lim"`
	Rate_max string `json:"rate_max"`
	Check_status string `json:"check_status"`
	Check_code string `json:"check_code"`
	Check_duration string `json:"check_duration"`
	Hrsp_1xx string `json:"hrsp_1xx"`
	Hrsp_2xx string `json:"hrsp_2xx"`
	Hrsp_3xx string `json:"hrsp_3xx"`
	Hrsp_4xx string `json:"hrsp_4xx"`
	Hrsp_5xx string `json:"hrsp_5xx"`
	Hrsp_other string `json:"hrsp_other"`
	Hanafail string `json:"hanafail"`
	Req_rate string `json:"req_rate"`
	Req_rate_max string `json:"req_rate_max"`
	Req_tot string `json:"req_tot"`
	Cli_abrt string `json:"cli_abrt"`
	Srv_abrt string `json:"srv_abrt"`
	Comp_in string `json:"comp_in"`
	Comp_out string `json:"comp_out"`
	Comp_byp string `json:"comp_byp"`
	Comp_rsp string `json:"comp_rsp"`
	Lastsess string `json:"lastsess"`
	Last_chk string `json:"last_chk"`
	Last_agt string `json:"last_agt"`
	Qtime string `json:"qtime"`
	Ctime string `json:"ctime"`
	Rtime string `json:"rtime"`
	Ttime string `json:"ttime"`
}

// struct to hold the output from the /info endpoint
type Info struct {
	Name string `json:"Name"`
	Version string `json:"Version"`
	Release_date string `json:"Release_date"`
	Nbproc string `json:"Nbproc"`
	Process_num string `json:"Process_num"`
	Pid string `json:"Pid"`
	Uptime string `json:"Uptime"`
	Uptime_sec string `json:"Uptime_sec"`
	Memmax_MB string `json:"Memmax_MB"`
	Ulimitn string `json:"Ulimit-n"`
	Maxsock string `json:"Maxsock"`
	Maxconn string `json:"Maxconn"`
	Hard_maxconn string `json:"Hard_maxconn"`
	CurrConns string `json:"CurrConns"`
	CumConns string `json:"CumConns"`
	CumReq string `json:"CumReq"`
	MaxSslConns string `json:"MaxSslConns"`
	CurrSslConns string `json:"CurrSslConns"`
	CumSslConns string `json:"CumSslConns"`
	Maxpipes string `json:"Maxpipes"`
	PipesUsed string `json:"PipesUsed"`
	PipesFree string `json:"PipesFree"`
	ConnRate string `json:"ConnRate"`
	ConnRateLimit string `json:"ConnRateLimit"`
	MaxConnRate string `json:"MaxConnRate"`
	SessRate string `json:"SessRate"`
	SessRateLimit string `json:"SessRateLimit"`
	MaxSessRate string `json:"MaxSessRate"`
	SslRate string `json:"SslRate"`
	SslRateLimit string `json:"SslRateLimit"`
	MaxSslRate string `json:"MaxSslRate"`
	SslFrontendKeyRate string `json:"SslFrontendKeyRate"`
	SslFrontendMaxKeyRate string `json:"SslFrontendMaxKeyRate"`
	SslFrontendSessionReuse_pct string `json:"SslFrontendSessionReuse_pct"`
	SslBackendKeyRate string `json:"SslBackendKeyRate"`
	SslBackendMaxKeyRate string `json:"SslBackendMaxKeyRate"`
	SslCacheLookups string `json:"SslCacheLookups"`
	SslCacheMisses string `json:"SslCacheMisses"`
	CompressBpsIn string `json:"CompressBpsIn"`
	CompressBpsOut string `json:"CompressBpsOut"`
	CompressBpsRateLim string `json:"CompressBpsRateLim"`
	ZlibMemUsage string `json:"ZlibMemUsage"`
	MaxZlibMemUsage string `json:"MaxZlibMemUsage"`
	Tasks string `json:"Tasks"`
	Run_queue string `json:"Run_queue"`
	Idle_pct string `json:"Idle_pct"`
	node string `json:"node"`
	description string `json:"description"`
}

// Main configuration object for load balancers. This contains all variables and is passed to
// the templating engine.
type Config struct {
	Frontends [] 	*Frontend			`json:frontends"`
	Backends [] 	*Backend 			`json:"backends"`
	Services[]		*Service			`json:"services"`
	PidFile  		string              	`json:"-"`
	Mutex    		*sync.RWMutex       	`json:"-"`
	templateFile    *string       			`json:"-"`
	configFile    	*string       			`json:"-"`
}

func (f *Config) SetPid(pid string) {
	f.PidFile = pid
}

func (f *Config) GetPid()(string) {
	return f.PidFile
}

type Foo struct {
	name string
}

func (f *Foo) SetName(name string) {
	f.name = name
}

type Service struct {

	Name           	string                  `json:"name"`
	BindPort		int						`json:"bindPort"`
	EndPoint		string					`json:"endPoint"`
	Mode			string					`json:"mode"`
}

// Defines a single haproxy "backend".
type Backend struct {
	Name           string                    `json:"name"`
	Mode		   string					 `json:"mode"`
	BackendServers [] *BackendServer 		 `json:"servers"`
	Options        ProxyOptions              `json:"options"`
}

// Defines a single haproxy "frontend".
type Frontend struct {
	Name           string                    `json:"name"`
	Mode		   string					 `json:"mode"`
	BindPort	   int			 			 `json:"bindPort"`
	BindIp		   string			 	 	 `json:"bindIp"`
	Options        ProxyOptions              `json:"options"`
	UseBackend	   string		 			 `json:"useBackend"`
}



type ProxyOptions struct {
	AbortOnClose    bool `json:"abortOnClose"`
	AllBackups      bool `json:"allBackups"`
	CheckCache      bool `json:"checkCache"`
	ForwardFor      bool `json:"forwardFor"`
	HttpClose       bool `json:"httpClose"`
	HttpCheck       bool `json:"httpCheck"`
	LdapCheck       bool `json:"ldapCheck"`
	MysqlCheck      bool `json:"mysqlCheck"`
	PgsqlCheck      bool `json:"pgsqlCheck"`
	RedisCheck      bool `json:"redisCheck"`
	SmtpCheck       bool `json:"smtpCheck"`
	SslHelloCheck   bool `json:"sslHelloCheck"`
	TcpKeepAlive    bool `json:"tcpKeepAlive"`
	TcpLog          bool `json:"tcpLog"`
	TcpSmartAccept  bool `json:"tcpSmartAccept"`
	TcpSmartConnect bool `json:"tcpSmartConnect"`
	Transparent     bool `json:"transparent"`
}

// Defines a server which exists in a backend.
type BackendServer struct {
	Name          string `json:"name"`
	Host		  string `json:"host"`
	Port          int 	 `json:"port"`
	Weight        int    `json:"weight"`
	MaxConn       int    `json:"maxconn"`
	Check         bool   `json:"check"`
	CheckInterval int    `json:"checkInterval"`
}
