package service

import (
	"slices"
	"time"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
)

type QueueService struct {
	cls *ClientService
}

func NewQueueService(cls *ClientService) *QueueService {
	return &QueueService{
		cls: cls,
	}
}

func (s *QueueService) Data(evt ct.ResponseEvent, params cu.IM) (data cu.IM, err error) {
	client := evt.Trigger.(*ct.Client)
	data = cu.IM{
		"filters": cu.IM{
			"report_type": "",
			"start_date":  time.Now().AddDate(0, 0, -30).Format(time.DateOnly),
			"end_date":    time.Now().Format(time.DateOnly),
			"ref_code":    "",
			"export_mode": "PDF",
			"orientation": "p",
			"paper_size":  "a4",
		},
		"print_queue":  []cu.IM{},
		"session_id":   client.Ticket.SessionID,
		"editor_icon":  ct.IconPrint,
		"editor_title": client.Msg("queue_title"),
	}

	return data, nil
}

func (s *QueueService) Response(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	switch evt.Name {
	case ct.FormEventOK:
		return s.formNext(evt)

	case ct.ClientEventSideMenu:
		return s.sideMenu(evt)

	default:
		return s.editorField(evt)
	}
}

func (s *QueueService) search(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)
	filters := cu.ToIM(stateData["filters"], cu.IM{})
	user := cu.ToIM(client.Ticket.User, cu.IM{})
	query := md.Query{
		Fields: []string{"q.*", "r.report_name",
			"a.user_name", "'" + client.Msg("queue_delete") + "' as delete_lbl", "'office_queue' as editor"},
		From:    `config_print_queue q inner join config_report r on q.report_code = r.code inner join auth a on q.auth_code = a.code`,
		OrderBy: []string{"q.id"},
	}
	if cu.ToString(filters["report_type"], "") != "" {
		query.Filters = append(query.Filters, md.Filter{Field: "q.ref_type", Comp: "==", Value: filters["report_type"]})
	}
	if cu.ToString(filters["start_date"], "") != "" {
		query.Filters = append(query.Filters, md.Filter{Field: "date(q.time_stamp)", Comp: ">=", Value: filters["start_date"]})
	}
	if cu.ToString(filters["end_date"], "") != "" {
		query.Filters = append(query.Filters, md.Filter{Field: "date(q.time_stamp)", Comp: "<=", Value: filters["end_date"]})
	}
	if cu.ToString(filters["ref_code"], "") != "" {
		query.Filters = append(query.Filters, md.Filter{Field: "q.ref_code", Comp: "like", Value: "%" + cu.ToString(filters["ref_code"], "") + "%"})
	}
	if cu.ToString(user["user_group"], "") != md.UserGroupAdmin.String() {
		query.Filters = append(query.Filters, md.Filter{Field: "q.auth_code", Comp: "==", Value: user["code"]})
	}
	var rows []cu.IM = []cu.IM{}
	if rows, err = ds.StoreDataQuery(query, false); err != nil {
		return evt, err
	}
	stateData["print_queue"] = rows
	stateData["view"] = "items"
	client.SetEditor("office_queue", cu.ToString(stateData["view"], ""), stateData)
	return evt, err
}

func (s *QueueService) sideMenu(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)

	menuMap := map[string]func() (re ct.ResponseEvent, err error){
		"queue_search": func() (re ct.ResponseEvent, err error) {
			return s.search(evt)
		},
		"editor_cancel": func() (re ct.ResponseEvent, err error) {
			client.ResetEditor()
			return evt, err
		},
	}

	if fn, ok := menuMap[cu.ToString(evt.Value, "")]; ok {
		return fn()
	}

	return evt, err
}

func (s *QueueService) formNext(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	ds := s.cls.getDataStore(client.Ticket.Database)

	frmValues := cu.ToIM(evt.Value, cu.IM{})
	frmData := cu.ToIM(frmValues["data"], cu.IM{})

	nextMap := map[string]func() (re ct.ResponseEvent, err error){
		"queue_delete": func() (re ct.ResponseEvent, err error) {
			printQueue := cu.ToIMA(stateData["print_queue"], []cu.IM{})
			if idx := slices.IndexFunc(printQueue, func(c cu.IM) bool {
				return cu.ToInteger(c["id"], 0) == cu.ToInteger(frmData["id"], 0)
			}); idx > int(-1) {
				if _, err = ds.StoreDataUpdate(md.Update{Model: "config",
					IDKey: cu.ToInteger(printQueue[idx]["id"], 0)}); err == nil {
					printQueue = append(printQueue[:idx], printQueue[idx+1:]...)
					stateData["print_queue"] = printQueue
				}
			}
			return evt, err
		},
	}

	if fn, ok := nextMap[cu.ToString(frmData["next"], "")]; ok {
		return fn()
	}
	return evt, err
}

func (s *QueueService) editorField(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, _, stateData := client.GetStateData()
	filters := cu.ToIM(stateData["filters"], cu.IM{})

	resultUpdate := func() (re ct.ResponseEvent, err error) {
		stateData["filters"] = filters
		client.SetEditor("office_queue", cu.ToString(stateData["view"], ""), stateData)
		return evt, err
	}

	values := cu.ToIM(evt.Value, cu.IM{})
	fieldName := cu.ToString(values["name"], "")
	value := cu.ToString(values["value"], "")
	valueData := cu.ToIM(values["value"], cu.IM{})

	fieldMap := map[string]func() (re ct.ResponseEvent, err error){
		ct.TableEventEditCell: func() (re ct.ResponseEvent, err error) {
			row := cu.ToIM(valueData["row"], cu.IM{})
			modal := cu.IM{
				"warning_label":   client.Msg("inputbox_delete"),
				"warning_message": "",
				"next":            "queue_delete",
				"id":              cu.ToInteger(row["id"], 0),
			}
			client.SetForm("warning", modal, 0, true)
			return evt, nil
		},

		"report_type": func() (re ct.ResponseEvent, err error) {
			filters[fieldName] = value
			return resultUpdate()
		},

		"start_date": func() (re ct.ResponseEvent, err error) {
			filters[fieldName] = value
			return resultUpdate()
		},

		"end_date": func() (re ct.ResponseEvent, err error) {
			filters[fieldName] = value
			return resultUpdate()
		},

		"ref_code": func() (re ct.ResponseEvent, err error) {
			filters[fieldName] = value
			return resultUpdate()
		},
	}

	if fn, ok := fieldMap[fieldName]; ok {
		return fn()
	}
	return evt, nil
}
