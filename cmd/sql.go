package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"tour/internal/sql2struct"
)

var (
	username  string
	password  string
	host      string
	charset   string
	dbType    string
	dbName    string
	tableName string
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换和处理",
	Long:  "sql转换和处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql 转换",
	Long:  "sql 转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %s", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %s", err)
		}
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %s", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	// go run . sql struct -u=root -p=? --db=blog_service --table=blog_tag
	sql2structCmd.Flags().StringVarP(&username, "username", "u", "root", "数据库账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "p", "", "数据库密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "数据库HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "数据库编码")
	sql2structCmd.Flags().StringVarP(&dbType, "dbtype", "", "mysql", "表名称")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "blog_service", "数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "blog_tag", "表名称")
}
