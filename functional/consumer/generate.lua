package.path = "../../types.lua;" .. package.path

local name = "consumer"

local function generate_consumer_named(type, name)
    local name = name:gsub("^%l", string.upper)
    
    io.write(string.format("\n// %s represents a function that accepts a %s.\n", name, type))
    io.write(string.format("type %s func(%s)\n", name, type))
end

local function generate_consumer(type)
    generate_consumer_named(type, type)
end

local function generate_consumers_named(type, name)
    local name = name:gsub("^%l", string.upper) .. "s"
    
    io.write(string.format("\n// %s represents a function that accepts a []%s.\n", name, type))
    io.write(string.format("type %s func([]%s)\n", name, type))
end

local function generate_consumers(type)
    generate_consumers_named(type, type)
end

file = io.open(name .. ".go", "w+")
io.output(file)

io.write("package " .. name .. "\n")

for _, v in ipairs(require("types")) do
    generate_consumer(v)
end

generate_consumer_named("interface{}", "Interface")

for _, v in ipairs(require("types")) do
    generate_consumers(v)
end

generate_consumers_named("interface{}", "Interface")

file:close()
io.output(io.stdout)
