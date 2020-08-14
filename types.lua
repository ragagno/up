local types = {}

setmetatable(types, { __shl = function (t, v) t[#t + 1] = v return t end })

_ = types << "bool"
_ = types << "int"       << "int8"    << "int16"  << "int32"  << "int64"
_ = types << "uint"      << "uint8"   << "uint16" << "uint32" << "uint64" << "uintptr"
_ = types << "float32"   << "float64"
_ = types << "complex64" << "complex128"
_ = types << "byte"      << "rune"
_ = types << "string"
_ = types << "error"

setmetatable(types, nil)

return types
