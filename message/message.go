package message

import (
	"fmt"
	"go-spg/utils"
)

const (
	ICON_INFORMATION = 0
	ICON_WARNING     = (1 << 0)
	ICON_ERROR       = (1 << 1)
	ICON_MASK = (ICON_INFORMATION | ICON_WARNING | ICON_ERROR)
	BTN_OK		= 0x00000400
	BTN_YES     = 0x00004000// QMessageBox::Yes
	BTN_NO      = 0x00010000// QMessageBox::No
	BTN_ABORT   = 0x00040000// QMessageBox::Abort
	BTN_RETRY   = 0x00080000// QMessageBox::Retry
	BTN_IGNORE  = 0x00100000// QMessageBox::Ignore
	BTN_CLOSE   = 0x00200000// QMessageBox::Close
	BTN_CANCEL  = 0x00400000// QMessageBox::Cancel
	BTN_DISCARD = 0x00800000// QMessageBox::Discard
	BTN_HELP    = 0x01000000// QMessageBox::Help
	BTN_APPLY   = 0x02000000// QMessageBox::Apply
	BTN_RESET   = 0x04000000// QMessageBox::Reset
	/**
	 * Mask of all available buttons in CClientUIInterface::MessageBoxFlags
	 * This needs to be updated, when buttons are changed there!
	 */
	BTN_MASK = (BTN_OK | BTN_YES | BTN_NO | BTN_ABORT | BTN_RETRY | BTN_IGNORE |
	BTN_CLOSE | BTN_CANCEL | BTN_DISCARD | BTN_HELP | BTN_APPLY | BTN_RESET)

	/* Force blocking, modal message box dialog (not just OS notification) */
	MODAL               = 0x10000000

	MSG_SECURE       = 0x40000000

	/* Predefined combinations for certain default usage cases */
	MSG_INFORMATION = ICON_INFORMATION
	MSG_WARNING = (ICON_WARNING | BTN_OK | MODAL)
	MSG_ERROR = (ICON_ERROR | BTN_OK | MODAL)
)

type MessageBox struct {
	message string
	caption string
	style   int
}

type MessageQuestion struct {
	message string
	caption string
	style   int
}

type MessageInit struct {
	Message string
}

var MessageBoxChan chan MessageBox

var MessageInitChan chan MessageInit

func ThreadSafeMessageBox(ch chan MessageBox) {
	for c:=range ch{
		message:=c.message
		caption:=c.caption
		style:=c.style
		fSecure := style & MSG_SECURE

		var strCaption string
		switch style {
		case MSG_ERROR:
			strCaption = "Error"
		case MSG_WARNING:
			strCaption = "Warning"
		case MSG_INFORMATION:
			strCaption = "Information"
		default:
			strCaption = caption
		}

		if fSecure == 0 {
			fmt.Println(strCaption, message)
		}
	}
}

func ThreadSafeMessageInit(ch chan MessageInit) {
	for c := range ch {
		utils.PrintInfo(c.Message)
	}
}

func SetupMessageThread() {
	utils.PrintInfo("SetupMessageThread ...")

	MessageBoxChan = make(chan MessageBox)
	go ThreadSafeMessageBox(MessageBoxChan)

	MessageInitChan = make(chan MessageInit)
	go ThreadSafeMessageInit(MessageInitChan)
}

func SendMessageBox(message string, caption string, style int){
	go func() {
		MessageBoxChan <- MessageBox{message, caption, style}
	}()
}

func SendMessageQuestion(message string, caption string, style int){
	SendMessageBox(message, caption, style)
}

func SendMessageInit(message string){
	go func() {
		MessageInitChan <- MessageInit{message}
	}()
}
