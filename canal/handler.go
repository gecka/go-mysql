package canal

import (
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
)

type EventHandler interface {
	OnRotate(rotateEvent *replication.RotateEvent, pos mysql.Position, rawData []byte) error
	// OnTableChanged is called when the table is created, altered, renamed or dropped.
	// You need to clear the associated data like cache with the table.
	// It will be called before OnDDL.
	OnTableChanged(schema string, table string) error
	OnDDL(queryEvent *replication.QueryEvent, pos mysql.Position, rawData []byte) error
	OnRow(e *replication.RowsEvent, ne *RowsEvent, pos mysql.Position, rawData []byte) error
	OnXID(e *replication.XIDEvent, pos mysql.Position, rawData []byte) error
	OnMariaGTID(e *replication.MariadbGTIDEvent, gtid mysql.GTIDSet, pos mysql.Position, rawData []byte) error
	OnGTID(e *replication.GTIDEvent, gtid mysql.GTIDSet, pos mysql.Position, rawData []byte) error
	// OnPosSynced Use your own way to sync position. When force is true, sync position immediately.
	OnPosSynced(pos mysql.Position, set mysql.GTIDSet, force bool) error
	OnDefault(e *replication.Event, pos mysql.Position, rawData []byte) error
	String() string
}

type DummyEventHandler struct {
}

func (h *DummyEventHandler) OnRotate(rotateEvent *replication.RotateEvent, pos mysql.Position, rawData []byte) error {
	return nil
}
func (h *DummyEventHandler) OnTableChanged(schema string, table string) error { return nil }
func (h *DummyEventHandler) OnDDL(queryEvent *replication.QueryEvent, pos mysql.Position, rawData []byte) error {
	return nil
}
func (h *DummyEventHandler) OnRow(e *replication.RowsEvent, ne *RowsEvent, pos mysql.Position, rawData []byte) error {
	return nil
}
func (h *DummyEventHandler) OnXID(e *replication.XIDEvent, pos mysql.Position, rawData []byte) error {
	return nil
}
func (h *DummyEventHandler) OnGTID(e *replication.GTIDEvent, gtid mysql.GTIDSet, pos mysql.Position, rawData []byte) error {
	return nil
}
func (h *DummyEventHandler) OnMariaGTID(e *replication.MariadbGTIDEvent, gtid mysql.GTIDSet, pos mysql.Position, rawData []byte) error {
	return nil
}
func (h *DummyEventHandler) OnPosSynced(mysql.Position, mysql.GTIDSet, bool) error { return nil }
func (h *DummyEventHandler) OnDefault(e replication.Event, pos mysql.Position, rawData []byte) error {
	return nil
}

func (h *DummyEventHandler) String() string { return "DummyEventHandler" }

// `SetEventHandler` registers the sync handler, you must register your
// own handler before starting Canal.
func (c *Canal) SetEventHandler(h EventHandler) {
	c.eventHandler = h
}
