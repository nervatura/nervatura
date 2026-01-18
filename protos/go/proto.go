package proto

import (
	"encoding/json"
	"strings"

	cu "github.com/nervatura/component/pkg/util"
	ut "github.com/nervatura/nervatura/v6/pkg/service/utils"
)

func (c *JsonBytes) MarshalJSON() ([]byte, error) {
	return c.Data, nil
}

func (c *JsonBytes) UnmarshalJSON(data []byte) error {
	c.Data = data
	return nil
}

func (mt *JsonString) MarshalJSON() ([]byte, error) {
	return json.Marshal(ut.SMToJS(mt.Data))
}

func (mt *JsonString) UnmarshalJSON(data []byte) error {
	var err error
	var mapData cu.IM
	if err = cu.ConvertFromByte(data, &mapData); err == nil {
		mt.Data = cu.IMToSM(mapData)
	}
	return err
}

func (u *UserGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

func (u *UserGroup) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := UserGroup_value[value]; found {
		*u = UserGroup(UserGroup_value[value])
	}
	return nil
}

func (c *CustomerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *CustomerType) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := CustomerType_value[value]; found {
		*c = CustomerType(CustomerType_value[value])
	}
	return nil
}

func (l *LinkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}

func (l *LinkType) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := LinkType_value[value]; found {
		*l = LinkType(LinkType_value[value])
	}
	return nil
}

func (l *LogType) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}

func (l *LogType) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := LogType_value[value]; found {
		*l = LogType(LogType_value[value])
	}
	return nil
}

func (m *MovementType) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

func (m *MovementType) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := MovementType_value[value]; found {
		*m = MovementType(MovementType_value[value])
	}
	return nil
}

func (p *PlaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *PlaceType) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := PlaceType_value[value]; found {
		*p = PlaceType(PlaceType_value[value])
	}
	return nil
}

func (p *PriceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *PriceType) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := PriceType_value[value]; found {
		*p = PriceType(PriceType_value[value])
	}
	return nil
}

func (b *BarcodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *BarcodeType) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := BarcodeType_value[value]; found {
		*b = BarcodeType(BarcodeType_value[value])
	}
	return nil
}

func (p *ProductType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *ProductType) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := ProductType_value[value]; found {
		*p = ProductType(ProductType_value[value])
	}
	return nil
}

func (r *RateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *RateType) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := RateType_value[value]; found {
		*r = RateType(RateType_value[value])
	}
	return nil
}

func (p *PaidType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *PaidType) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := PaidType_value[value]; found {
		*p = PaidType(PaidType_value[value])
	}
	return nil
}

func (t *TransStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *TransStatus) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := TransStatus_value[value]; found {
		*t = TransStatus(TransStatus_value[value])
	}
	return nil
}

func (t *TransState) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *TransState) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := TransState_value[value]; found {
		*t = TransState(TransState_value[value])
	}
	return nil
}

func (d *Direction) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Direction) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := Direction_value[value]; found {
		*d = Direction(Direction_value[value])
	}
	return nil
}

func (t *TransType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *TransType) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := TransType_value[value]; found {
		*t = TransType(TransType_value[value])
	}
	return nil
}

func (c *ConfigType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *ConfigType) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := ConfigType_value[value]; found {
		*c = ConfigType(ConfigType_value[value])
	}
	return nil
}

func (f *FieldType) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.String())
}

func (f *FieldType) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := FieldType_value[value]; found {
		*f = FieldType(FieldType_value[value])
	}
	return nil
}

func (m *MapFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

func (m *MapFilter) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := MapFilter_value[value]; found {
		*m = MapFilter(MapFilter_value[value])
	}
	return nil
}

func (m *ShortcutMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

func (m *ShortcutMethod) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := ShortcutMethod_value[value]; found {
		*m = ShortcutMethod(ShortcutMethod_value[value])
	}
	return nil
}

func (f *ShortcutField) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.String())
}

func (f *ShortcutField) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := ShortcutField_value[value]; found {
		*f = ShortcutField(ShortcutField_value[value])
	}
	return nil
}

func (f *FileType) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.String())
}

func (f *FileType) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if _, found := FileType_value[value]; found {
		*f = FileType(FileType_value[value])
	}
	return nil
}

func (c *Config) UnmarshalJSON(data []byte) error {
	var cfData map[string]interface{}
	var err error
	if err = cu.ConvertFromByte(data, &cfData); err == nil {
		c.Id = cu.ToInteger(cfData["id"], 0)
		c.Code = cu.ToString(cfData["code"], "")
		c.ConfigType = ConfigType(ConfigType_value[cu.ToString(cfData["config_type"], "")])
		c.TimeStamp = cu.ToString(cfData["time_stamp"], "")
		switch c.ConfigType {
		case ConfigType_CONFIG_MAP:
			var mapConfig ConfigMap
			if err = ut.ConvertToType(cfData["data"], &mapConfig); err == nil {
				c.Data = &Config_Map{Map: &mapConfig}
			}
		case ConfigType_CONFIG_SHORTCUT:
			var shortcutConfig ConfigShortcut
			if err = ut.ConvertToType(cfData["data"], &shortcutConfig); err == nil {
				c.Data = &Config_Shortcut{Shortcut: &shortcutConfig}
			}
		case ConfigType_CONFIG_MESSAGE:
			var messageConfig ConfigMessage
			if err = ut.ConvertToType(cfData["data"], &messageConfig); err == nil {
				c.Data = &Config_Message{Message: &messageConfig}
			}
		case ConfigType_CONFIG_PATTERN:
			var patternConfig ConfigPattern
			if err = ut.ConvertToType(cfData["data"], &patternConfig); err == nil {
				c.Data = &Config_Pattern{Pattern: &patternConfig}
			}
		case ConfigType_CONFIG_REPORT:
			var reportConfig ConfigReport
			if err = ut.ConvertToType(cfData["data"], &reportConfig); err == nil {
				c.Data = &Config_Report{Report: &reportConfig}
			}
		case ConfigType_CONFIG_PRINT_QUEUE:
			var printQueueConfig ConfigPrintQueue
			if err = ut.ConvertToType(cfData["data"], &printQueueConfig); err == nil {
				c.Data = &Config_PrintQueue{PrintQueue: &printQueueConfig}
			}
		case ConfigType_CONFIG_DATA:
			mapData := cu.IMToSM(cu.ToIM(cfData["data"], cu.IM{}))
			c.Data = &Config_ConfigData{ConfigData: &JsonString{Data: mapData}}
		}
		return err
	}
	return json.Unmarshal(data, c)
}

func (cd *Config_ConfigData) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &cd.ConfigData.Data)
}

func (cd *Config_ConfigData) MarshalJSON() ([]byte, error) {
	return json.Marshal(cd.ConfigData.Data)
}

func (m *Config_Map) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &m.Map)
}

func (m *Config_Map) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Map)
}

func (s *Config_Shortcut) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.Shortcut)
}

func (s *Config_Shortcut) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Shortcut)
}

func (m *Config_Message) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &m.Message)
}

func (m *Config_Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Message)
}

func (p *Config_Pattern) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &p.Pattern)
}

func (p *Config_Pattern) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Pattern)
}

func (r *Config_Report) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Report)
}

func (r *Config_Report) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Report)
}

func (p *Config_PrintQueue) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &p.PrintQueue)
}

func (p *Config_PrintQueue) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.PrintQueue)
}
