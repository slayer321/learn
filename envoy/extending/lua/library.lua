LIBRARY = {}

function LIBRARY.RandomString()
  local result = ""
  for i = 1, 24 do
    result = result .. string.char(math.random(97, 122))
  end
  return result
end

return LIBRARY