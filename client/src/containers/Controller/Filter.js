export const UserFilter = {
  usergroup_filter: {
    transitem: ["t.cruser_id","in",[{select:["id"], from:"employee", 
      where:["usergroup","=","?"]}]],
    transmovement: ["t.cruser_id","in",[{select:["id"], from:"employee", 
      where:["usergroup","=","?"]}]],
    transpayment: ["t.cruser_id","in",[{select:["id"], from:"employee", 
      where:["usergroup","=","?"]}]]
  },
  employee_filter: {
    transitem: ["t.cruser_id","=","?"],
    transmovement: ["t.cruser_id","=","?"],
    transpayment: ["t.cruser_id","=","?"]
  }
}