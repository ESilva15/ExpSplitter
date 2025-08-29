-- This scripts needs to do 2 things:
-- Update the expenses where the owner doesn't have shares or payments to have
-- some of his because of how the normalization function works
-- Normalize the shares and update the expense

local json = require("json")

local info = debug.getinfo(1, 'S')
print("Running ", info.source)

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

local function insertOwnerShare(exp)
  local share = {
    User = exp.Owner,
    Share = 0,
  }

  local ok, err = InsertShare(json.encode(share), exp.ExpID)
  return ok, err
end

local function insertOwnerPayment(exp)
  local payment = {
    User = exp.Owner,
    PayedAmount = 0,
  }

  local ok, err = InsertPayment(json.encode(payment), exp.ExpID)
  return ok, err
end

local function validateSharesAndPayments(exp)
  if not isOwnerInShares(exp) then
    local ok, err = insertOwnerShare(exp)
    if not ok then
      print("Error inserting owner share: ", err)
      return false
    end
    print("Inserted share in expense: ", exp.ExpID)
  end

  if not isOwnerInPayments(exp) then
    local ok, err = insertOwnerPayment(exp)
    if not ok then
      print("Error inserting owner payment: ", err)
      return false
    end
    print("Inserted payment in expense: ", exp.ExpID)
  end

  return true
end

local function requiresUpdate(exp, normalized)
  return exp ~= normalized
end

---@diagnostic disable-next-line: undefined-global
local ok, data = GetAllExpenses()
if not ok or data == nil then
  print("Error:", data)
  return
end

for i = 1, #data do
  local eJson = data[i].expense
  local e = json.decode(eJson)

  local res = validateSharesAndPayments(e)
  if not res then
    goto continue
  end

  -- if we inserted shares and payments we need to update the expense here
  local updatedEJson
  ok, updatedEJson = GetExpense(e.ExpID)
  if not ok then
    goto continue
  end

  local normalizedExpense
  ok, normalizedExpense = NormalizeShare(updatedEJson)
  if not ok then
    print("Failed to normalize share: ", normalizedExpense)
    goto continue
  end

  local newExp = json.decode(updatedEJson)
  local normalized = json.decode(normalizedExpense)
  if requiresUpdate(updatedEJson, normalizedExpense) then
    local err
    ok, err = UpdateExpense(normalizedExpense)
    if not ok then
      print("Failed to update expense ", normalized.ExpID, ": ", err)
      goto continue
    end
    print("Updated expense ", normalized.ExpID)
  end

  ::continue::
end
