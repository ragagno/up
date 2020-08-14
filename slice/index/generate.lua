package.path = "../../types.lua;" .. package.path

local name = "index"

local function generate_index(type)
    local name = type:gsub("^%l", string.upper)

    io.write(string.format("\n// %s returns the index of value in slice, or -1 if it doesn't exist.\n", name))
    io.write(string.format("func %s(slice []%s, value %s) int {\n", name, type, type))
    io.write(string.format("    for i, e := range slice {\n"))
    io.write(string.format("        if value == e {\n"))
    io.write(string.format("            return i\n"))
    io.write(string.format("        }\n"))
    io.write(string.format("    }\n"))
    io.write(string.format("\n"))
    io.write(string.format("    return -1\n"))
    io.write(string.format("}\n"))
end

file = io.open(name .. ".go", "w+")
io.output(file)

io.write("package " .. name .. "\n")

for _, v in ipairs(require("types")) do
    generate_index(v)
end

file:close()
io.output(io.stdout)
