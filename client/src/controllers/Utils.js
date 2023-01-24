export const getSql = (engine, _sql) => {
  let prmCount = 0;
  const engineType = (sql_) => {
    let sql = sql_
    switch (engine) {
      case "sqlite":
      case "sqlite3":
        sql = sql.replace(/{CCS}/g, "");
        sql = sql.replace(/{SEP}/g, "||");
        sql = sql.replace(/{CCE}/g, "");
        sql = sql.replace(/{CAS_TEXT}/g, "cast(");
        sql = sql.replace(/{CAE_TEXT}/g, " as text)");
        sql = sql.replace(/{CAS_INT}/g, "cast("); // cast as integer - start
        sql = sql.replace(/{CAE_INT}/g, " as integer)"); // cast as integer - end
        sql = sql.replace(/{CAS_FLOAT}/g, "cast(");
        sql = sql.replace(/{CAE_FLOAT}/g, " as double)"); // " as real)");
        sql = sql.replace(/{CAS_DATE}/g, "");
        sql = sql.replace(/{CASF_DATE}/g, "");
        sql = sql.replace(/{CAE_DATE}/g, "");
        sql = sql.replace(/{CAEF_DATE}/g, "");
        sql = sql.replace(/{FMSF_NUMBER}/g, "");
        sql = sql.replace(/{FMSF_DATE}/g, "");
        sql = sql.replace(/{FMEF_CONVERT}/g, "");
        sql = sql.replace(/{FMS_FLOAT}/g, "");
        sql = sql.replace(/{FME_FLOAT}/g, "");
        sql = sql.replace(/{FMS_INT}/g, "");
        sql = sql.replace(/{FME_INT}/g, "");
        sql = sql.replace(/{FMS_DATE}/g, "substr("); // format to iso date - start
        sql = sql.replace(/{FME_DATE}/g, ",1,10)"); // format to iso date - end
        sql = sql.replace(/{FMS_DATETIME}/g, "substr("); // format to iso datetime - start
        sql = sql.replace(/{FME_DATETIME}/g, ",1,19)"); // format to iso datetime - end
        sql = sql.replace(/{FMS_TIME}/g, "substr(time(");
        sql = sql.replace(/{FME_TIME}/g, "),0,6)");
        sql = sql.replace(/{JOKER}/g, "'%'");
        sql = sql.replace(/{CUR_DATE}/g, "date('now')");
        break;
      case "mysql":
        sql = sql.replace(/{CCS}/g, "concat(");
        sql = sql.replace(/{SEP}/g, ",");
        sql = sql.replace(/{CCE}/g, ")");
        sql = sql.replace(/{CAS_TEXT}/g, "cast(");
        sql = sql.replace(/{CAE_TEXT}/g, " as char)");
        sql = sql.replace(/{CAS_INT}/g, "cast("); // cast as integer - start
        sql = sql.replace(/{CAE_INT}/g, " as signed)"); // cast as integer - end
        sql = sql.replace(/{CAS_FLOAT}/g, "cast(");
        sql = sql.replace(/{CAE_FLOAT}/g, " as decimal)"); // " as decimal)");
        sql = sql.replace(/{CAS_DATE}/g, "cast(");
        sql = sql.replace(/{CASF_DATE}/g, "cast(");
        sql = sql.replace(/{CAE_DATE}/g, " as date)");
        sql = sql.replace(/{CAEF_DATE}/g, " as date)");
        sql = sql.replace(/{FMSF_NUMBER}/g, "");
        sql = sql.replace(/{FMSF_DATE}/g, "");
        sql = sql.replace(/{FMEF_CONVERT}/g, "");
        sql = sql.replace(/{FMS_FLOAT}/g, "replace(format(cast(");
        sql = sql.replace(/{FME_FLOAT}/g, " as decimal(10,2)),2),'.00','')");
        sql = sql.replace(/{FMS_INT}/g, "format(cast(");
        sql = sql.replace(/{FME_INT}/g, " as signed), 0)");
        sql = sql.replace(/{FMS_DATE}/g, "date_format(");
        sql = sql.replace(/{FME_DATE}/g, ", '%Y-%m-%d')");
        sql = sql.replace(/{FMS_DATETIME}/g, "date_format(");
        sql = sql.replace(/{FME_DATETIME}/g, ", '%Y-%m-%dT%H:%i:%s')");
        sql = sql.replace(/{FMS_TIME}/g, "cast(cast(");
        sql = sql.replace(/{FME_TIME}/g, " as time) as char)");
        sql = sql.replace(/{JOKER}/g, "'%'");
        sql = sql.replace(/{CUR_DATE}/g, "current_date");
        break;
      case "postgres":
        sql = sql.replace(/{CCS}/g, "");
        sql = sql.replace(/{SEP}/g, "||");
        sql = sql.replace(/{CCE}/g, "");
        sql = sql.replace(/{CAS_TEXT}/g, "cast(");
        sql = sql.replace(/{CAE_TEXT}/g, " as text)");
        sql = sql.replace(/{CAS_INT}/g, "cast("); // cast as integer - start
        sql = sql.replace(/{CAE_INT}/g, " as integer)"); // cast as integer - end
        sql = sql.replace(/{CAS_FLOAT}/g, "cast(");
        sql = sql.replace(/{CAE_FLOAT}/g, " as float)");
        sql = sql.replace(/{CAS_DATE}/g, "cast(");
        sql = sql.replace(/{CASF_DATE}/g, "cast(");
        sql = sql.replace(/{CAE_DATE}/g, " as date)");
        sql = sql.replace(/{CAEF_DATE}/g, " as date)");
        sql = sql.replace(
          /{FMSF_NUMBER}/g,
          "case when rf_number.fieldname is null then 0 else "
        );
        sql = sql.replace(
          /{FMSF_DATE}/g,
          "case when rf_date.fieldname is null then current_date else "
        );
        sql = sql.replace(/{FMEF_CONVERT}/g, " end ");
        sql = sql.replace(/{FMS_FLOAT}/g, "replace(to_char(cast(");
        sql = sql.replace(
          /{FME_FLOAT}/g,
          " as float),'999,999,990.00'),'.00','')"
        );
        sql = sql.replace(/{FMS_INT}/g, "to_char(cast(");
        sql = sql.replace(/{FME_INT}/g, " as integer), '999,999,999')");
        sql = sql.replace(/{FMS_DATE}/g, "to_char(");
        sql = sql.replace(/{FME_DATE}/g, ", 'YYYY-MM-DD')");
        sql = sql.replace(/{FMS_DATETIME}/g, "to_char(");
        sql = sql.replace(/{FME_DATETIME}/g, ", 'YYYY-MM-DD\"T\"HH24:MI:SS')");
        sql = sql.replace(/{FMS_TIME}/g, "substr(cast(cast(");
        sql = sql.replace(/{FME_TIME}/g, " as time) as text), 0, 6)");
        sql = sql.replace(/{JOKER}/g, "chr(37)");
        sql = sql.replace(/{CUR_DATE}/g, "current_date");
        break;
      case "mssql":
        sql = sql.replace(/{CCS}/g, "");
        sql = sql.replace(/{SEP}/g, "+");
        sql = sql.replace(/{CCE}/g, "");
        sql = sql.replace(/{CAS_TEXT}/g, "cast(");
        sql = sql.replace(/{CAE_TEXT}/g, " as nvarchar)");
        sql = sql.replace(/{CAS_INT}/g, "cast(");
        sql = sql.replace(/{CAE_INT}/g, " as int)");
        sql = sql.replace(/{CAS_FLOAT}/g, "cast(");
        sql = sql.replace(/{CAE_FLOAT}/g, " as real)");
        sql = sql.replace(/{CAS_DATE}/g, "cast(");
        sql = sql.replace(/{CASF_DATE}/g, "");
        sql = sql.replace(/{CAE_DATE}/g, " as date)");
        sql = sql.replace(/{CAEF_DATE}/g, "");
        sql = sql.replace(/{FMSF_NUMBER}/g, "");
        sql = sql.replace(/{FMSF_DATE}/g, "");
        sql = sql.replace(/{FMEF_CONVERT}/g, "");
        sql = sql.replace(/{FMS_FLOAT}/g, "replace(convert(varchar,cast("); // mssql 2012+ format(???,'N2')
        sql = sql.replace(/{FME_FLOAT}/g, " as money),1),'.00','')");
        sql = sql.replace(/{FMS_INT}/g, "replace(convert(varchar,cast(");
        sql = sql.replace(/{FME_INT}/g, " as money),1), '.00','')");
        sql = sql.replace(/{FMS_DATE}/g, "convert(varchar(10),");
        sql = sql.replace(/{FME_DATE}/g, ", 120)");
        sql = sql.replace(/{FMS_DATETIME}/g, "convert(varchar(19),");
        sql = sql.replace(/{FME_DATETIME}/g, ", 120)");
        sql = sql.replace(/{FMS_TIME}/g, "SUBSTRING(cast(cast(");
        sql = sql.replace(/{FME_TIME}/g, " as time) as nvarchar),0,6)");
        sql = sql.replace(/{JOKER}/g, "'%'");
        sql = sql.replace(/{CUR_DATE}/g, "cast(GETDATE() as DATE)");
        break;
      default:
        break;
    }
    return sql;
  };

  const sqlDecode = (data_, key) => {
    let data = data_
    let sql = "";
    if (Array.isArray(data)) {
      let sep = " ";
      let startBr = "";
      let endBr = "";
      if (data.length > 0) {
        if (["select","select_distinct","union_select","order_by","group_by"].includes(key) ||
          data[0].length === 0
        ) {
          sep = ", ";
        }
      }
      data.forEach((element_) => {
        let element = element_
        if (typeof element === "undefined" || element === null) {
          element = "null";
        }
        if (element.length === 0) {
          if (key !== "set") {
            startBr = "(";
            endBr = ")";
          }
        } else if (
          data.length === 2 &&
          (element === "and" || element === "or")
        ) {
          sql += `${element  } (`;
          endBr = ")";
        } else if (key && data.length === 1 && typeof data[0] === "object") {
          sql += ` (${  sqlDecode(element, key)  })`;
        } else {
          sql += sep + sqlDecode(element, key);
        }
      });
      if (sep === ", ") {
        sql = sql.substr(2);
      }
      if (key && data.includes("on")) {
        sql = key.replace("_", " ") + sql;
      }
      return startBr + sql.toString().trim() + endBr;
    }
    if (typeof data === "object") {
      Object.keys(data).forEach(_key => {
        if (_key === "inner_join" || _key === "left_join") {
          sql += ` ${  sqlDecode(data[_key], _key)}`;
        } else {
          sql +=
            ` ${  _key.replace("_", " ")  } ${  sqlDecode(data[_key], _key)}`;
        }
      })
      return sql;
    }
    if (data.includes("?") && key !== "select") {
      prmCount += 1;
      if(engine === "postgres"){
        data = data.replace("?",`$${  prmCount}`);
      }
    }
    return data;
  };

  if (typeof _sql === "string") {
    return {sql: engineType(_sql), prmCount};
  } 
    return {sql: engineType(sqlDecode(_sql)), prmCount};
  
};

export const request = async (url, options) => {
  const parseJSON = (response) => {
    if (response.status === 401)
      return { code: 401, message: "Unauthorized" }
    if (response.status === 400)
      return response.json()
    if (response.status === 204 || response.status === 205) {
      return null;
    }
    switch (response.headers.get('content-type').split(";")[0]) {
      case "application/pdf":
      case "application/xml":
      case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
        return response.blob()

      case "application/json":
        return response.json()
      
      case "text/plain":
      case "text/csv":
        return response.text()
    
      default:
        return response
    }
  }

  const checkStatus = (response) => {
    if ((response.status >= 200 && response.status < 300) || (response.status === 400) || (response.status === 401)) {
      return response;
    }

    const error = new Error(response.statusText);
    error.response = response;
    throw error;
  }

  const fetchResponse = await fetch(url, options);
  const response = checkStatus(fetchResponse);
  return parseJSON(response);
}

export const saveToDisk = (fileUrl, fileName) => {
  const element = document.createElement("a")
  element.href = fileUrl
  element.download = fileName || fileUrl
  document.body.appendChild(element) // Required for this to work in FireFox
  element.click()
}