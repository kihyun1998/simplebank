package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	// db 연결 테스트
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect : ", err)
	}

	// db 접속 객체를 testQueries에 저장
	testQueries = New(conn)

	// 테스트 시작
	os.Exit(m.Run())
}