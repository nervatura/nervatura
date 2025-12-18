package service

import (
	"slices"
	"strings"
	"time"

	ct "github.com/nervatura/component/pkg/component"
	cu "github.com/nervatura/component/pkg/util"
	md "github.com/nervatura/nervatura/v6/pkg/model"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

type SearchService struct {
	cls *ClientService
}

func NewSearchService(cls *ClientService) *SearchService {
	return &SearchService{
		cls: cls,
	}
}

func (s *SearchService) Data(evt ct.ResponseEvent, params cu.IM) (data cu.IM, err error) {
	client := evt.Trigger.(*ct.Client)
	ds := s.cls.getDataStore(client.Ticket.Database)

	data = cu.IM{"result": []cu.IM{}}
	view := cu.ToString(params["view"], "")
	query := md.Query{}
	if pq, found := params["query"].(md.Query); found {
		query = pq
	}
	filters := []ct.BrowserFilter{}
	if pf, found := params["filters"].([]ct.BrowserFilter); found {
		filters = pf
	}

	queryFilters := []string{}
	for _, filter := range filters {
		queryFilters = s.cls.UI.SearchConfig.Filter(view, filter, queryFilters)
	}

	if len(queryFilters) > 0 {
		queryFilter := strings.Join(queryFilters, " ")
		queryFilter, _ = strings.CutPrefix(queryFilter, "or ")
		queryFilter, _ = strings.CutPrefix(queryFilter, "and ")
		queryFilter = "(" + queryFilter + ")"
		if len(query.Filter) > 0 {
			queryFilter = " and " + queryFilter
		}
		query.Filter += queryFilter
	}

	data["result"], err = ds.StoreDataQuery(query, false)
	return data, err
}

func (s *SearchService) sideMenu(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, stateKey, stateData := client.GetStateData()
	groups := s.cls.UI.SearchConfig.SideGroups(client.Labels())
	switch strings.Split(cu.ToString(evt.Value, ""), "_")[0] {

	case "group":
		viewName := stateKey
		simple := cu.ToBoolean(stateData["simple"], false)
		if cu.ToString(stateData["side_group"], "") == cu.ToString(evt.Value, "") {
			stateData["side_group"] = ""
		} else {
			stateData["side_group"] = evt.Value
			if idx := slices.IndexFunc(groups, func(g md.SideGroup) bool {
				return g.Name == evt.Value
			}); idx > int(-1) {
				viewName = groups[idx].Views[0]
				simple = s.cls.UI.SearchConfig.View(viewName, client.Labels()).Simple
				stateData["rows"] = []cu.IM{}
				stateData["filter_value"] = ""
			}
		}
		client.SetSearch(viewName, stateData, simple)

	default:
		stateData["rows"] = []cu.IM{}
		stateData["filter_value"] = ""
		client.SetSearch(cu.ToString(evt.Value, ""), stateData, s.cls.UI.SearchConfig.View(cu.ToString(evt.Value, ""), cu.SM{}).Simple)

	}
	return evt, err
}

func (s *SearchService) Response(evt ct.ResponseEvent) (re ct.ResponseEvent, err error) {
	client := evt.Trigger.(*ct.Client)
	_, stateKey, stateData := client.GetStateData()

	switch evt.Name {
	case ct.ClientEventSideMenu:
		return s.sideMenu(evt)

	case ct.BrowserEventBookmark:
		modal := cu.IM{
			"title":         client.Msg("bookmark_new"),
			"icon":          ct.IconStar,
			"label":         client.Msg("bookmark_enter"),
			"placeholder":   "",
			"field_name":    "value",
			"default_value": "",
			"required":      false,
			"next":          "bookmark_add",
		}
		client.SetForm("input_string", modal, 0, true)

	case ct.FormEventOK:
		sConf := s.cls.UI.SearchConfig.View(stateKey, client.CustomFunctions.Labels(client.Lang))
		visibleColumns := client.GetSearchVisibleColumns(ut.ToBoolMap(sConf.VisibleColumns, map[string]bool{}))
		viewData := cu.ToIM(stateData[stateKey], cu.IM{})
		frmValues := cu.ToIM(evt.Value, cu.IM{})
		frmValue := cu.ToIM(frmValues["value"], cu.IM{})
		label := cu.ToString(frmValue["value"], "")
		bookmark := md.Bookmark{
			BookmarkType: md.BookmarkTypeBrowser,
			Label:        label,
			Key:          stateKey,
			Code:         sConf.Title,
			Filters:      viewData["filters"],
			Columns:      visibleColumns,
			TimeStamp:    time.Now().Format(time.RFC3339),
		}
		return s.cls.addBookmark(evt, bookmark), nil

	}
	return evt, err
}
