package backend

import (
	"fmt"
	"log"
	"net"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"tribes/tools"
)

func HeaderCompliance(header []string) (safe_header []string) {
	var hdr map[string]string
	hdr = make(map[string]string)
	var elem []string
	var s_hdr []string
	var lastseen int = 0

	for i := range header {
		if strings.Count(header[i], ": ") == 0 {
			header[lastseen] += " " + header[i]
			header[i] = "Averno: yes"
		} else {
			lastseen = i
		}

	}

	for i := range header {
		elem = strings.SplitN(header[i], ": ", 2)
		hdr[elem[0]] = strings.Trim(elem[1], " ")
		log.Printf("[HEADER-RFC] header %s: saved as %s", elem[0], hdr[elem[0]])
	}

	delete(hdr, "Averno")

	if _, ok := hdr["Subject"]; ok == false {
		hdr["Subject"] = "Too Lazy to add a Subject"
	}
	if _, ok := hdr["From"]; ok == false {
		hdr["From"] = "I am not Anonymous, I am just lazy"
	}

	if _, ok := hdr["Newsgroups"]; ok == false {
		hdr["Newsgroups"] = "garbage"
	}

	// only one format for date. 'cause I can.

	const layout = "Mon Jan 2 2006 15:04:05 MST"
	orario := time.Now()
	hdr["Date"] = orario.Format(layout)

	for key, value := range hdr {
		s_hdr = append(s_hdr, key+": "+value)
	}

	return s_hdr

}

func GenerateXOVERLineFromHeader(header []string, groupname string, message_id string, msgnum_str string) []string {
	var hdr map[string]string
	hdr = make(map[string]string)
	var xover_line []string
	var elem []string

	for i := range header {

		elem = strings.SplitN(header[i], ":", 2)
		hdr[elem[0]] = strings.Trim(elem[1], " ")
		log.Printf("[XOVER-SAVE] header %s: saved as %s", elem[0], hdr[elem[0]])
	}

	// some stupid newsreaders want XOVER but they can't understand a message with no references
	if _, ok := hdr["References"]; ok == false {
		hdr["References"] = "00001"
	}

	xover_line = append(xover_line, fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\tXref: %s", msgnum_str, hdr["Subject"], hdr["From"], hdr["Date"], hdr["Message-ID"], hdr["References"], hdr["Bytes"], hdr["Lines"], hdr["Xref"]))

	return xover_line

}

// Adds the xover file during the POST phase
func SaveXOVERLineForPost(header []string, groupname string, message_id string, msgnum_str string) {

	xover_line := GenerateXOVERLineFromHeader(header, groupname, message_id, msgnum_str)

	xover_file := filepath.Join(messages_folder, "x-"+groupname+"-"+msgnum_str+"-"+message_id)

	err := tools.WriteMessages(xover_line, xover_file)
	if err != nil {
		log.Printf("[WTF] cannot write xover on %s", xover_file)
	} else {
		log.Printf("[FYI] xover saved in %s", xover_file)
	}

}

// answers for the stupid LIST OVERVIEW.FMT command

func NNTP_LIST_OVERVIEW_FMT(conn net.Conn) {

	conn.Write([]byte("215 Order of fields in overview database\r\n"))
	list_overview := "Subject:\r\nFrom:\r\nDate:\r\nMessage-ID:\r\nReferences:\r\nBytes:\r\nLines:\r\nXref:full\r\n.\r\n"
	conn.Write([]byte(list_overview))
}

// frontend for the XOVER and OVER commands.

func NNTP_XOVER_Frontend(groupname string, conn net.Conn, pattern string) {

	//  only checks WTF the client is asking for

	if matches, _ := regexp.MatchString("(?i)^X?OVER[ ]*$", pattern); matches == true {
		fake_id := GetFirstNumByGroup(groupname)
		log.Printf("[XOVER-FE] Matched with no arguments for  ", groupname)
		NNTP_XOVER_ReturnByInterval(groupname, conn, fake_id, fake_id)
	}

	if matches, _ := regexp.MatchString("(?i)^X?OVER[ ]*[0-9]+$", pattern); matches == true {
		ids := strings.Split(pattern, " ")
		log.Printf("[XOVER-FE] Matched with one argument (%s) for %s", ids[1], groupname)
		NNTP_XOVER_ReturnByInterval(groupname, conn, ids[1], ids[1])
	}

	if matches, _ := regexp.MatchString("(?i)^X?OVER[ ]*[0-9]+-$", pattern); matches == true {
		ids := strings.Split(pattern, " ")
		low := strings.TrimRight(ids[1], "-")
		high := GetLastNumByGroup(groupname)
		log.Printf("[XOVER-FE] Open interval from (%s) to (%s) for %s", low, high, groupname)
		NNTP_XOVER_ReturnByInterval(groupname, conn, low, high)
	}

	if matches, _ := regexp.MatchString("(?i)^X?OVER[ ]*([0-9]+)-([0-9]+)$", pattern); matches == true {

		ids := strings.Split(pattern, " ")
		interval := strings.Split(ids[1], "-")

		m_high_int, _ := strconv.Atoi(interval[1])
		m_low_int, _ := strconv.Atoi(interval[0])

		if m_high_int < m_low_int {

			swap := m_high_int
			m_high_int = m_low_int
			m_low_int = swap
		}

		high, _ := strconv.Atoi(GetLastNumByGroup(groupname))
		low, _ := strconv.Atoi(GetFirstNumByGroup(groupname))

		log.Printf("[XOVER-FE] Step 1 interval from (%d) to (%d) for %s", m_low_int, m_high_int, groupname)
		log.Printf("[XOVER-FE] Step 1 limit from (%d) to (%d) for %s", low, high, groupname)

		if m_high_int > high {
			m_high_int = high
		}
		if m_low_int < low {
			m_low_int = low
		}

		log.Printf("[XOVER-FE] Step 2 interval from (%d) to (%d) for %s", m_low_int, m_high_int, groupname)
		log.Printf("[XOVER-FE] Step 2 limit from (%d) to (%d) for %s", low, high, groupname)

		interval[0] = fmt.Sprintf("%05d", m_low_int)
		interval[1] = fmt.Sprintf("%05d", m_high_int)

		log.Printf("[XOVER-FE] Closed interval from (%s) to (%s) for %s", interval[0], interval[1], groupname)
		NNTP_XOVER_ReturnByInterval(groupname, conn, interval[0], interval[1])
	}

}

func NNTP_XOVER_ReturnByInterval(groupname string, conn net.Conn, m_low string, m_high string) (err error) {

	log.Printf("[XOVER-BE] Request for %s was: from %s to %s   ", groupname, m_low, m_high)

	m_high_int, _ := strconv.Atoi(m_high)
	m_low_int, _ := strconv.Atoi(m_low)

	// some stupid newsreader are using offset 0.... no words.

	if m_high_int == 0 {
		m_high_int = 1
	}
	if m_low_int == 0 {
		m_low_int = 1
	}

	m_low = fmt.Sprintf("%05d", m_low_int)
	m_high = fmt.Sprintf("%05d", m_high_int)

	log.Printf("[XOVER-BE] Start processing %s from %s to %s   ", groupname, m_low, m_high)

	// retrieve all the xover info files for the group

	if files, err := filepath.Glob(messages_folder + "/x-" + groupname + "-*"); err != nil {
		log.Printf("[SOB] Cannot search  in %s  ", groupname)
		conn.Write([]byte("224 No articles in " + m_low + "-" + m_high + "\r\n"))
		return err
	} else {
		if files == nil {
			log.Printf("[SOB] %s is empty ", groupname)
			conn.Write([]byte("224 No articles in " + groupname + "\r\n"))
			return nil
		}

		// it the group is populated, then sort it

		if files != nil {
			sort.Strings(files)
		}

		// only one number was submitted or min==max for any reason

		if m_low == m_high {
			conn.Write([]byte("224 Overview information for " + m_low + " follows\r\n"))
			log.Printf("[XOVER-BE] Equal executing for %s,%s  ", m_low, m_high)
			if pos := tools.StringPosInSlice(m_low, files); pos >= 0 {
				log.Printf("[XOVER-BE] Retrieved file %d for name %s  ", pos, files[pos])
				Transmit_Article(conn, files[pos])
			} else {
				log.Printf("[XOVER-BE] Whatever: %d for name %s  ", pos, files)
			}
			conn.Write([]byte(".\r\n"))
			return nil
		}

		//

		if m_low != m_high {

			log.Printf("[XOVER-BE] executing for %s -> %s  ", m_low, m_high)
			msg := "224 Overview information for " + m_low + "-" + m_high + " follows\r\n"
			conn.Write([]byte(msg))
			log.Printf("[XOVER-BE] starting cycle %s -> %s  ", m_low_int, m_high_int)
			for i := m_low_int; i <= m_high_int; i++ {
				filetoken := "-" + fmt.Sprintf("%05d", i) + "-"
				pos := tools.StringPosInSlice(filetoken, files)
				Transmit_Article(conn, files[pos])
				log.Printf("[END] %d xover printed ", i)
			}
			conn.Write([]byte(".\r\n"))
			return nil
		}

	}
	return nil
}
