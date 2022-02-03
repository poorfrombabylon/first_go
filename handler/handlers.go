package handler

import (
	"fmt"
	"net/http"
	"test_task/repository/inmemory"
	"test_task/repository/postgres"
)

var db_in = inmemory.NewInMemory()
var db_out = postgres.NewPostgresDB()

var longUrlFormTmpl = []byte(`
<html>
	<body>
		<form action = "/" method = "POST">
			New LongUrl: <input type="text" name="longurl">
			<input type="submit" value="LOGIN">
		</form>
	</body>
	<body>
		<form action = "/" method = "GET">
			Find ShortUrl: <input type="text" name="longurl">
			<input type="submit" value="LOGIN">
		</form>
	</body>
</html>
`)

var storage string

func Stor_type(st_type string) {
	fmt.Println(st_type)
	storage = st_type
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		func() {
			w.Write(longUrlFormTmpl)
			return
		}()
		longUrl := r.FormValue("longurl")
		//fmt.Fprintln(w, "long: ", longUrl)

		if longUrl == "" {
			fmt.Fprintln(w, "incorrect input")
		}

		if storage == "in" {
			shortUrl, err := db_in.GetShortUrl(longUrl)
			fmt.Println("in memory")
			if err != nil {
				fmt.Fprintln(w, err)
			}
			w.Write([]byte(fmt.Sprintf("short: %v for %v", shortUrl, longUrl)))
		}
		if storage == "out" {
			shortUrl, err := db_out.GetUrl(longUrl)

			if err != nil {
				fmt.Fprintln(w, err)
			}

			w.Write([]byte(fmt.Sprintf("short: %v for %v", shortUrl, longUrl)))
		} else {
			fmt.Println("Storage type was not choosen")
		}

	}
	if r.Method == http.MethodPost {
		func() {
			w.Write(longUrlFormTmpl)
			return
		}()
		longUrl := r.FormValue("longurl")
		fmt.Fprintln(w, "long: ", longUrl, r.Method)

		if storage == "in" {
			fmt.Println("in memory")
			shortUrl, err := db_in.PostUrl(longUrl)
			if err != nil {
				fmt.Fprintln(w, err)
			}
			w.Write([]byte(fmt.Sprintf("short: %v for %v", shortUrl, longUrl)))
			//fmt.Println(db_in)
			return
		}
		if storage == "out" {
			shortUrl, err := db_out.PostUrl(longUrl)
			if err != nil {
				fmt.Fprintln(w, err)
			}
			w.Write([]byte(fmt.Sprintf("short: %v for %v", shortUrl, longUrl)))
			//fmt.Println(db_in)
			return
		}

	}
}
