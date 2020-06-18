package main
import(
	"rura/downLoad/message"
	"strings"
	"testing"
)
func Test_Messages (t *testing.T){
	s:=message.MakeStartRegion(123)
	reg:=message.ReadStartRegion(s)
	if reg!=123 {
		t.Errorf("Start 123 <> %d",reg)
	}
	s=message.MakeEndRegion(123)
	reg=message.ReadEndRegion(s)
	if reg!=123 {
		t.Errorf("End 123 <> %d",reg)
	}
	s=message.MakeStartSendingTable("table")
	is,table:=message.ReadStartSendingTable(s)

	if !is || strings.Compare("table",table)!=0{
		t.Error("startSendingTable ",table)
	}
	s=message.MakeEndSendingTable("table")
	is,table=message.ReadEndSendingTable(s)

	if !is || strings.Compare("table",table)!=0{
		t.Error("endSendingTable ",table)
	}
}
