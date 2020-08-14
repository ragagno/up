package.path = "../../types.lua;" .. package.path

local name = "reverse"

local function generate_reverse(type)
    local name = type:gsub("^%l", string.upper)

    io.write(string.format("\n// %s returns slice in reverse order.\n", name))
    io.write(string.format("func %s(slice []%s) []%s {\n", name, type, type))
    io.write(string.format("    var reverse = make([]%s, len(slice))\n", type))
    io.write(string.format("\n"))
    io.write(string.format("    for index, value := range slice {\n"))
    io.write(string.format("        reverse[len(reverse) - 1 - index] = value\n"))
    io.write(string.format("    }\n"))
    io.write(string.format("\n"))
    io.write(string.format("    return reverse\n"))
    io.write(string.format("}\n"))
end

file = io.open(name .. ".go", "w+")
io.output(file)

io.write("package " .. name .. "\n")

for _, v in ipairs(require("types")) do
    generate_reverse(v)
end

file:close()
io.output(io.stdout)
