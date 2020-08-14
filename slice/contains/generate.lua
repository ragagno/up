package.path = "../../types.lua;" .. package.path

local name = "contains"

local function generate_contains(type)
    local name = type:gsub("^%l", string.upper)

    io.write(string.format("\n// %s returns true if slice contains value.\n", name))
    io.write(string.format("func %s(slice []%s, value %s) bool {\n", name, type, type))
    io.write(string.format("    for _, e := range slice {\n"))
    io.write(string.format("        if value == e {\n"))
    io.write(string.format("            return true\n"))
    io.write(string.format("        }\n"))
    io.write(string.format("    }\n"))
    io.write(string.format("\n"))
    io.write(string.format("    return false\n"))
    io.write(string.format("}\n"))
end

file = io.open(name .. ".go", "w+")
io.output(file)

io.write("package " .. name .. "\n")

for _, v in ipairs(require("types")) do
    generate_contains(v)
end

file:close()
io.output(io.stdout)
