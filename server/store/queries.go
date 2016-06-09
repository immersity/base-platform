package store

var (
	sqlInsertAccount            = string(MustAsset("sql/insert_account.sql"))
	sqlSelectAccountCredentials = string(MustAsset("sql/select_account_credentials.sql"))
)
