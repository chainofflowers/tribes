package backend

import (
	"log"
    "net"
	"path/filepath"
	"strconv"
    "bufio"
    "../tools/"
)


func NNTP_POST_ReadAndSave(conn net.Conn, groupname string) {

    num_message,_  := strconv.Atoi( GetNumFilesByGroup(groupname) )
    num_message++



    id_message := tools.RandSeq(32)

    answer_ok := "340 Ok, recommended ID <" + id_message + ">\n"
    conn.Write([]byte(answer_ok))
    log.Printf("[FYI] %s", answer_ok)

    var body []string
    var headers []string
    var is_header bool = true


    scanner_h := bufio.NewScanner(conn)
    for  {
        ok := scanner_h.Scan()
        if !ok {break}
        line := scanner_h.Text()
        if line == "" {
                        log.Printf("[FYI] header line was ->%s<-", line)
                        is_header = false
                        continue
                        }

        if is_header {
                    log.Printf("[FYI] header line is:  ->%s<-", line)
                    headers = append(headers, line)
        }else{
            log.Printf("[FYI] body line is:  ->%s<-", line)
            body = append(body,line)

        }
        if line == "." {break}
    }

    headers = append(headers, "Message-ID: <" + id_message + ">")

    header_file := filepath.Join(messages_folder,"h-" + groupname + "-" + strconv.Itoa(num_message) + "-" + id_message)
    body_file := filepath.Join(messages_folder,"b-" + groupname + "-" + strconv.Itoa(num_message) + "-" + id_message)

    err := tools.WriteMessages(headers, header_file)
    if err != nil { log.Printf("[WTF] cannot write headers on %s", header_file) } else { log.Printf("[FYI] headers saved in %s", header_file) }

    err = tools.WriteMessages(body, body_file)
    if err != nil { log.Printf("[WTF] cannot write body on %s", body_file) } else { log.Printf("[FYI] body saved in %s", body_file) }

    conn.Write([]byte("240 article posted ok\n"))


}
