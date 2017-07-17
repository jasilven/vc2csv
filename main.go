package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	keys := parseArgs()
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		if strings.HasPrefix(s.Text(), "BEGIN:VCARD") {
			m := make(map[string]string)
			for s.Scan(); !(strings.HasPrefix(s.Text(), "END:VCARD")); s.Scan() {
				k, v := parseLine(keys[:], s.Text())
				m[k] = strings.TrimPrefix(m[k]+","+v, ",")
			}
			if csv := toCsv(m, keys[:]); csv != "" {
				fmt.Println(csv)
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatalln("scanning standard input: " + err.Error())
	}
}

func parseArgs() []string {
	fieldsPtr := flag.String("fields", "FN EMAIL TEL", "vcard fields to be extracted")
	flag.Parse()
	return strings.Split(*fieldsPtr, " ")
}

func parseLine(prefs []string, l string) (string, string) {
	for _, p := range prefs {
		if strings.HasPrefix(l, p+";") || strings.HasPrefix(l, p+":") {
			s := strings.Split(l, ":")
			return p, strings.Replace(s[1], ";", ",", -1)
		}
	}
	return "", ""
}

func toCsv(m map[string]string, keys []string) string {
	retval := ""
	for _, v := range keys {
		retval = retval + ";" + strings.TrimSpace(m[v])
	}
	return strings.TrimPrefix(retval, ";")
}
