package backend

import (
	"log"
    "net"
	"path/filepath"
    "io/ioutil"
	"strconv"
    "bufio"
    "strings"
    "../tools/"
)


func NNTP_POST_ReadAndSave(conn net.Conn, groupname string) {

    num_message ,_ := strconv.AtoI( GetNumFilesByGroup(groupname string) )
    num_message++

    id_message := tools.RandSeq(32)

    header_file := filepath.Join(messages_folder,"h-" + groupname + "-" + strconv.Itoa(num_message) + "-" + id_message)
    body_file := filepath.Join(messages_folder,"b-" + groupname + "-" + strconv.Itoa(num_message) + "-" + id_message)

    answer_ok := "340 Ok, recommended ID <" + id_message + ">\n"
    conn.Write([]byte(answer_ok))

    var body []string
    var headers []string

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        line := scanner.Text()
        if line == "\n\n" {
            break
        }else{
        headers = append(headers,line,'\n')
        }
    }

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        line := scanner.Text()
        if line == ".\n" {
            break
        }else{
        body = append(body,line,'\n')
        }
    }



    err := ioutil.WriteFile(header_file, headers, 0644)
    if err != nil { log.Printf("[WTF] cannot write headers on %s", header_file) } else { log.Printf("[FYI] headers saved in %s", header_file) }

    err := ioutil.WriteFile(header_file, body, 0644)
    if err != nil { log.Printf("[WTF] cannot write body on %s", body_file) } else { log.Printf("[FYI] body saved in %s", body_file) }

    conn.Write([]byte("240 article posted ok"))


}
