package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strings"
)

type cfgmail struct {
	Username string
	Password string
	Smtphost string
	Mailto   string
}

type cfg struct {
	Name, Text string
}

func main() {

	subject := flag.String("s", "email by linfei", "mail subject")
	cfgfile := flag.String("c", "conf.json", "config file")
	bodyfile := flag.String("f", "index.html", "mail body file")
	flag.Parse()
	//从json文件中读取发送邮件服务器配置信息
	cfgjson := getConf(*cfgfile)

	var cfg cfgmail
	dec := json.NewDecoder(strings.NewReader(cfgjson))
	for {

		if err := dec.Decode(&cfg); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		//fmt.Printf("%s\n%s\n%s\n", cfg.Username, cfg.Password, cfg.Smtphost)

	}

	username := cfg.Username
	password := cfg.Password
	host := cfg.Smtphost
	mailto := cfg.Mailto
	to := strings.Split(mailto, ";")

	fmt.Printf("subject : %s\n", *subject)
	fmt.Printf("config file : %s\n", *cfgfile)
	fmt.Printf("mailbody file : %s\n", *bodyfile)
	body := readfile(*bodyfile)
	fmt.Println(body)
	//fmt.Printf("============")
	//fmt.Println(username)
	//subject := "Test send email by golang"

	/*body := `
	  <html>
	  <body>
	  <h3>
	  "Test send email by golang，来个测试试一下"
	  </h3>
	  </body>
	  </html>
	  `
	*/

	fmt.Printf("%s\n%s\n%s\n", username, password, host)
	//for _, value := range to {
	//	fmt.Printf("%s\n", value)
	//}
	//fmt.Printf("%s\n%s\n%s\n", cfg.Username, cfg.Password, cfg.Smtphost)

	for _, mailtoaddr := range to {
		err := SendMail(username, password, host, mailtoaddr, *subject, body, "html")
		if err != nil {
			fmt.Println("send mail error!")
			fmt.Println(err)
		} else {
			fmt.Println("send mail success!")
		}
	}

}

func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func getConf(filename string) string {
	//filename := "conf.json"
	file, err := os.Open(filename)

	defer file.Close()
	if err != nil {
		fmt.Println("read conf file error")
		log.Fatal(err)
	}

	buf := make([]byte, 512)
	var str1 string
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		//os.Stdout.Write(buf[:n])

		str := string(buf[:n])

		str1 = str1 + str
	}
	return str1
}

func readfile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	//fmt.Println(string(fd))
	return string(fd)
}
