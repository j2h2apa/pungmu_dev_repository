package main

import (
	"fmt"

	"gopkg.in/rana/ora.v4"
	"gopkg.in/rana/ora.v4/lg"
)

func main() {
	// use an optional log package for ora logging
	cfg := ora.Cfg()
	cfg.Log.Logger = lg.Log
	ora.SetCfg(cfg)

	env, err := ora.OpenEnv()
	defer env.Close()
	if err != nil {
		panic(err)
	}

	srvCfg := ora.SrvCfg{Dblink: "LTDB_57"}
	srv, err := env.OpenSrv(srvCfg)
	defer srv.Close()
	if err != nil {
		panic(err)
	}

	sesCfg := ora.SesCfg{
		Username: "ltuser",
		Password: "lt2005",
	}
	ses, err := srv.OpenSes(sesCfg)
	defer ses.Close()
	if err != nil {
		panic(err)
	}

	// test for server ping
	err = ses.Ping()
	if err == nil {
		fmt.Println("Ping successful")
	}

	// get server version
	version, err := srv.Version()
	if version != "" && err == nil {
		fmt.Printf("Received version from server %v", version)
	}

	// create table
	tableName := "t1"
	query := fmt.Sprintf(
		"CREATE TABLE %v "+
			"(C1 NUMBER(19,0), "+
			" C2 VARCHAR2(48 CHAR))",
		tableName,
	)
	var rowsAffected uint64
	// fmt.Println(query)
	// stmtTbl, err = ses.Prep(query)
	// defer stmtTbl.Close()
	// if err != nil {
	// 	panic(err)
	// }
	// rowsAffected, err = stmtTbl.Exe()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("crate table rowsAffected : %d\n", rowsAffected)

	// begin 1st transation
	tx, err := ses.StartTx()
	if err != nil {
		panic(err)
	}

	// insert record
	var id uint64
	var arg1 string = "Go is expressive, concise, clean, and efficient."
	query = fmt.Sprintf("insert into %v (c1, c2) values (J2H2S2APA_PERSON_ID.NEXTVAL, :c2) returning c1 into :c1", tableName)
	stmtIns, err := ses.Prep(query)
	defer stmtIns.Close()
	rowsAffected, err = stmtIns.Exe(arg1, &id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("insert result : rowsAffected %d - id %d\n", rowsAffected, id)

	// insert nullable Slice string
	var args1 []ora.String = make([]ora.String, 4)
	args1[0] = ora.String{Value: "Its concurrency mechanisms make it easy to"}
	args1[1] = ora.String{IsNull: true}
	args1[2] = ora.String{Value: "It's a fast, statically typed, compiled"}
	args1[3] = ora.String{Value: "One of Go's key design goals is code"}
	query = fmt.Sprintf("insert into %v (c1, c2) values (J2H2S2APA_PERSON_ID.NEXTVAL, :c2)", tableName)
	stmtSliceIns, err := ses.Prep(query)
	defer stmtSliceIns.Close()
	if err != nil {
		panic(err)
	}
	rowsAffected, err = stmtSliceIns.Exe(args1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Slice insert result : rowsAffected %d\n", rowsAffected)

	// fetch records
	query = fmt.Sprintf("select c1, c2 from %v", tableName)
	stmtQry, err := ses.Prep(query)
	defer stmtQry.Close()
	if err != nil {
		panic(err)
	}
	rset, err := stmtQry.Qry()
	if err != nil {
		panic(err)
	}
	fmt.Println(rset.ColumnIndex())
	for rset.Next() {
		fmt.Println(rset.Row[0], rset.Row[1], rset.Columns[0], rset.Columns[1])
	}
	if err := rset.Err(); err != nil {
		panic(err)
	}

	// commit 1st transation
	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	//////////////////////////////////////////////////////////////////////////
	// begin 2nd transation
	tx2, err := ses.StartTx()
	if err != nil {
		panic(err)
	}

	// insert null String
	query = fmt.Sprintf("insert into %v (c1, c2) values (J2H2S2APA_PERSON_ID.NEXTVAL, :c2)", tableName)
	nullableStr := ora.String{IsNull: true}
	stmtTrans, err := ses.Prep(query)
	defer stmtTrans.Close()
	if err != nil {
		panic(err)
	}
	rowsAffected, err = stmtTrans.Exe(nullableStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("insert nullable string result : rowsAffected %d\n", rowsAffected)
	// rollback 2nd transaction
	err = tx2.Rollback()
	if err != nil {
		panic(err)
	}

	// fetch and specify return type
	query = fmt.Sprintf("select count(c1) from %v where c2 is null", tableName)
	stmtCount, err := ses.Prep(query, ora.U8)
	defer stmtCount.Close()
	if err != nil {
		panic(err)
	}
	rset, err = stmtCount.Qry()
	if err != nil {
		panic(err)
	}
	row := rset.NextRow()
	if row != nil {
		fmt.Printf("fetch and specify return type %d\n", row[0])
	}
	if err := rset.Err(); err != nil {
		panic(err)
	}

	// create stored procedure with sys_refcursor
	// query = fmt.Sprintf(
	// 	"create or replace procedure j2h2s2apa_proc1(p1 out sys_refcursor) as begin "+
	// 		"open p1 for select c1, c2 from %v where c1 > 2 order by c1;"+
	// 		"end j2h2s2apa_proc1;",
	// 	tableName)
	// stmtProcCreate, err := ses.Prep(query)
	// defer stmtProcCreate.Close()
	// if err != nil {
	// 	panic(err)
	// }
	// rowsAffected, err = stmtProcCreate.Exe()
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("=======================================================")
	// call stored procedure
	query = "CALL J2H2S2APA_PROC1(:1)"
	stmtProcCall, err := ses.Prep(query)
	defer stmtProcCall.Close()
	if err != nil {
		panic(err)
	}
	var procRset *ora.Rset = &ora.Rset{}
	rowsAffected, err = stmtProcCall.Exe(procRset)
	if err != nil {
		panic(err)
	}
	if procRset.IsOpen() {
		for procRset.Next() {
			fmt.Println(procRset.Row[0], procRset.Row[1])
		}
		if err := procRset.Err(); err != nil {
			panic(err)
		}
		fmt.Printf("procRset len : %d", procRset.Len())
	}
}
