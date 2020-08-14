package.path = "../../types.lua;" .. package.path

local name = "predicate"

local function generate_predicate_named(type, name)
    local name = name:gsub("^%l", string.upper)
    
    io.write(string.format("\n// %s represents a function that accepts a %s and returns a bool if the predicate is validated.\n", name, type))
    io.write(string.format("type %s func(%s) bool\n", name, type))
end

local function generate_predicate(type)
    generate_predicate_named(type, type)
end

local function generate_predicates_named(type, name)
    local name = name:gsub("^%l", string.upper) .. "s"
    
    io.write(string.format("\n// %s represents a function that accepts a []%s and returns a bool if the predicate is validated.\n", name, type))
    io.write(string.format("type %s func([]%s) bool\n", name, type))
end

local function generate_predicates(type)
    generate_predicates_named(type, type)
end

file = io.open(name .. ".go", "w+")
io.output(file)

io.write("package " .. name .. "\n")

for _, v in ipairs(require("types")) do
    generate_predicate(v)
end

generate_predicate_named("interface{}", "Interface")

for _, v in ipairs(require("types")) do
    generate_predicates(v)
end

generate_predicates_named("interface{}", "Interface")

file:close()
io.output(io.stdout)
