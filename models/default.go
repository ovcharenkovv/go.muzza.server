package models

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
