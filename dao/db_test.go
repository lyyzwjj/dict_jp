package dao

import (
	"github.com/lyyzwjj/dict_jp/model"
	"testing"
)

func TestReMigrate(t *testing.T) {
	TestMigrateDrop(t)
	TestMigrateInit(t)
}

func TestMigrateInit(t *testing.T) {
	InitMySQLDB()
	if err := Repo.AutoMigrate(&model.Vocabulary{}, &model.Word{}, &model.WordMeaning{}); err != nil {
		return
	}
}

func TestMigrateDrop(t *testing.T) {
	InitMySQLDB()
	Repo.Exec("DROP TABLE word_meanings;")
	Repo.Exec("DROP TABLE words;")
	Repo.Exec("DROP TABLE vocabularies;")
}

func TestTruncate(t *testing.T) {
	InitMySQLDB()
	Repo.Exec("SET FOREIGN_KEY_CHECKS=0;")
	Repo.Exec("TRUNCATE TABLE word_meanings;")
	Repo.Exec("TRUNCATE TABLE words;")
	Repo.Exec("TRUNCATE TABLE vocabularies;")
	Repo.Exec("SET FOREIGN_KEY_CHECKS=1;")
	Repo.Exec("SELECT @@FOREIGN_KEY_CHECKS;")
}
