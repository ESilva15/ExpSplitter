local cjson = require("cjson")

---@diagnostic disable-next-line: undefined-global
local expenses, err = GetAllExpenses()
if not expenses then
  print("Error:", err)
end

for i = 1, #expenses do
  local eJson = expenses[i].expense
  local e = cjson.decode(eJson)

  print(e)
end
