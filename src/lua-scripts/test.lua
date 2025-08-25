---@diagnostic disable-next-line: undefined-global
local expenses, err = GetAllExpenses()

if not expenses then
  print("Error:", err)
else
  for _, e in ipairs(expenses) do
    print(e.id, e.description, e.amount)
  end
end

---@diagnostic disable-next-line: undefined-global
local a = AddDecimal("0.1", "0.25")
print(a)
