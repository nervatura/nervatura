insert into ui_audit (usergroup, nervatype, subtype, inputfilter, supervisor)
(select usergroup.id as usergroup, nervatype.id as nervatype, null as subtype, inputfilter.id as inputfilter, 0 as supervisor
from (select id from groups where groupname='usergroup' 
  and groupvalue in('demo','guest','user')) usergroup, 
  (select id from groups where groupname='nervatype' 
    and groupvalue in('audit','customer','employee','event','price','product','project','tool','setting')) nervatype,
  (select id from groups where groupname='inputfilter' and groupvalue='disabled') inputfilter
union select usergroup.id as usergroup, nervatype.id as nervatype, subtype.id as subtype, inputfilter.id as inputfilter, 0 as supervisor
from (select id from groups where groupname='usergroup' 
  and groupvalue in('demo','guest','user')) usergroup, 
  (select id from groups where groupname='nervatype' and groupvalue='trans') nervatype,
  (select id from groups where groupname='transtype' 
    and groupvalue in('bank','cash','delivery','formula','inventory','invoice','receipt','offer','order','production','rent','waybill','worksheet')) subtype,
  (select id from groups where groupname='inputfilter' and groupvalue='disabled') inputfilter
union select usergroup.id as usergroup, nervatype.id as nervatype, subtype.id as subtype, inputfilter.id as inputfilter, 0 as supervisor
from (select id from groups where groupname='usergroup' 
  and groupvalue in('demo','guest','user')) usergroup, 
  (select id from groups where groupname='nervatype' and groupvalue='report') nervatype,
  (select id from ui_report) subtype,
  (select id from groups where groupname='inputfilter' and groupvalue='disabled') inputfilter
union select usergroup.id as usergroup, nervatype.id as nervatype, subtype.id as subtype, inputfilter.id as inputfilter, 0 as supervisor
from (select id from groups where groupname='usergroup' 
  and groupvalue in('demo','guest','user')) usergroup, 
  (select id from groups where groupname='nervatype' and groupvalue='menu') nervatype,
  (select id from ui_menu) subtype,
  (select id from groups where groupname='inputfilter' and groupvalue='disabled') inputfilter
order by usergroup, nervatype, subtype)