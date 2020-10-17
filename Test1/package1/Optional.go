package package1

import (
	"bufio"
	"flag"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//zad 1,2
func StringLength(s string) int {
	str := flag.String("napis", "", "use to declare string from console")
	flag.Parse()
	if *str != "" {
		return len(strings.Split(*str, ""))
	} else {
		return len(strings.Split(s, ""))
	}

}

//zad 3
func CheckForRepeats(path string) (map[string]int, error) {
	m := make(map[string]int)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for v := range strings.Split(scanner.Text(), " ") {
			m[strings.Split(scanner.Text(), " ")[v]]++
		}

	}
	return m, nil
}

//zad 4
func SaveMapToFile(filename string, m map[string]int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for v := range m {
		var str string = v + " - " + strconv.Itoa(m[(v)]) + "\n"
		io.WriteString(file, str)
	}
	return nil
}

//zad 5
func fileToString(filepath string) (string, error) {
	str := ""
	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str += scanner.Text() + "\n"
	}
	return str, nil
}
func servermsg(w http.ResponseWriter, r *http.Request) {
	message, _ := fileToString("lol2.txt")
	w.Write([]byte(message))
}
func CreateServer() {
	http.HandleFunc("/", servermsg)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

//zad 6
func ReadHttp(url string) string {
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	return "%s\n" + string(html)
}
