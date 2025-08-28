-- This scripts needs to do 2 things:
-- Update the expenses where the owner doesn't have shares or payments to have
-- some of his
-- Normalize the shares

local json = require("json")

local function printTable(t)
  for key, value in pairs(t) do
    if type(value) == "table" then
      printTable(value)
    else
      print(key, ":", value)
    end
  end
end

local function isOwnerInShares(exp)
  for _, share in pairs(exp.Shares) do
    if share.User.UserID == exp.Owner.UserID then
      return true
    end
  end

  return false
end

local function isOwnerInPayments(exp)
  for _, paym in pairs(exp.Payments) do
    if paym.User.UserID == exp.Owner.UserID then
      return true
    end
  end

  return false
end

---@diagnostic disable-next-line: undefined-global
local expenses, err = GetAllExpenses()
if not expenses then
  print("Error:", err)
end

for i = 1, #expenses do
  -- print "------------------------------------"
  local eJson = expenses[i].expense
  local e = json.decode(eJson)

  if not isOwnerInShares(e) then
    -- Insert a share in the owner user into shares as 0.0
    goto continue
  end

  if not isOwnerInPayments(e) then
    -- Insert a share in the owner user into shares as 0.0
    goto continue
  end

  normalizedExpense = NormalizeShare(eJson)
  if eJson == normalizedExpense then
    print "This one needs to be updated"
  end

  ::continue::
end
