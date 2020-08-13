package.path = "../types.lua;" .. package.path

local name = "producer"

local function generate_producer_named(type, name)
    local name = name:gsub("^%l", string.upper)
    
    io.write(string.format("\n// The type `%s` represents a function that returns a `%s`.\n", name, type))
    io.write(string.format("type %s func() %s\n", name, type))
end

local function generate_producer(type)
    generate_producer_named(type, type)
end

local function generate_producers_named(type, name)
    local name = name:gsub("^%l", string.upper) .. "s"
    
    io.write(string.format("\n// The type `%s` represents a function that returns a `[]%s`.\n", name, type))
    io.write(string.format("type %s func() []%s\n", name, type))
end

local function generate_producers(type)
    generate_producers_named(type, type)
end

file = io.open(name .. ".go", "w+")
io.output(file)

io.write("package " .. name .. "\n")

for _, v in ipairs(require("types")) do
    generate_producer(v)
end

generate_producer_named("interface{}", "Interface")

for _, v in ipairs(require("types")) do
    generate_producers(v)
end

generate_producers_named("interface{}", "Interface")

file:close()
io.output(io.stdout)
