package backend

import (
	"../tools/"

	"fmt"
	"log"
	"net"
	"path/filepath"

	"strings"
)

func SaveXOVERLineForPost(header []string, groupname string, message_id string, msgnum_str string) {
	var hdr map[string]string
	hdr = make(map[string]string)
	var xover_line []string
	var elem []string

	for i := range header {

		elem = strings.SplitN(header[i], ":",2)
		hdr[elem[0]] = strings.Trim(elem[1], " ")
        log.Printf("[FYI] header %s: saved as %s", elem[0], hdr[elem[0]])
	}
    log.Printf("[XOVER] preparing for write" )

	xover_line = append(xover_line, fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\tXref: %s", msgnum_str, hdr["Subject"], hdr["From"],hdr["Date"], hdr["Message-ID"], hdr["References"], hdr["Bytes"], hdr["Lines"], hdr["Xref"]) )




    log.Printf("[XOVER] string composed" )

	xover_file := filepath.Join(messages_folder, "x-"+groupname+"-"+msgnum_str+"-"+message_id)

	err := tools.WriteMessages(xover_line, xover_file)
	if err != nil {
		log.Printf("[WTF] cannot write xover on %s", xover_file)
	} else {
		log.Printf("[FYI] xover saved in %s", xover_file)
	}

}

func NNTP_LIST_OVERVIEW_FMT(conn net.Conn){

        conn.Write([]byte("215 Order of fields in overview database\r\n"))
        list_overview := "Subject:\r\nFrom:\r\nDate:\r\nMessage-ID:\r\nReferences:\r\nBytes:\r\nLines:\r\nXref:full\r\n.\r\n"
        conn.Write([]byte(list_overview))
}
