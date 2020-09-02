package constants

// BaseURL is the base for all URLs in JSON responses
// const BaseURL = "http://localhost:8001"

// NotFoundJSON is the JSON sent when there are no results (e.g. because of pagination overflow)
const NotFoundJSON = "{\"detail\": \"Not found\"}"

// PageSize is the number of items displayed in a JSON result
const PageSize = 10

// MaxUint is the maximum unsigned integer
const MaxUint = ^uint64(0)

// MinUint is the minimum unsigned integer, that is: 0
const MinUint = 0

// MaxInt is the maximum unsigned integer
const MaxInt = int(MaxUint >> 1)

// MinInt is the minimum signed integer
const MinInt = -MaxInt - 1
