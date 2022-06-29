module iHome

go 1.18

require (
	github.com/beego/beego/v2 v2.0.4
	github.com/sharabao13/fdfs_client v0.0.0-20220624082934-564f07351e3f
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5
)

require github.com/Elemlee/weilaihui_goconfig v0.0.0-20210112012228-a2993d5d7875 // indirect

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/smartystreets/goconvey v1.6.4
)

replace github.com/gomodule/redigo v2.0.0+incompatible => github.com/gomodule/redigo v1.8.4
