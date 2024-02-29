package models

type OriginalMessage struct {
	From    string `json:"from"`
	Sent    string `json:"sent"`
	To      string `json:"to"`
	Cc      string `json:"cc"`
	Bcc     string `json:"bcc"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

type Email struct {
	MessageID               string             `json:"message_id"`
	Date                    string             `json:"date"`
	From                    string             `json:"from"`
	To                      string             `json:"to"`
	Subject                 string             `json:"subject"`
	MimeVersion             string             `json:"mime_version"`
	ContentType             string             `json:"content_type"`
	ContentTransferEncoding string             `json:"content_transfer_encoding"`
	XFrom                   string             `json:"x_from"`
	XTo                     string             `json:"x_to"`
	Xcc                     string             `json:"x_cc"`
	Xbcc                    string             `json:"x_bcc"`
	XFolder                 string             `json:"x_folder"`
	XOrigin                 string             `json:"x_origin"`
	XFileName               string             `json:"x_file_name"`
	Content                 string             `json:"content"`
	OriginalMessage         *[]OriginalMessage `json:"original_message"`
}

var HeadersList = []string{
	"Message-ID",
	"Date",
	"From",
	"To",
	"Subject",
	"Mime-Version",
	"Content-Type",
	"Content-Transfer-Encoding",
	"X-From",
	"X-To",
	"X-cc",
	"X-bcc",
	"X-Folder",
	"X-Origin",
	"X-FileName",
}

var OriginalMessageSeparator = " -----Original Message-----"
