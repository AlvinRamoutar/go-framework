package data

type ENPack struct{}

var EN_DATA = map[string]string{
	"ENLIBCOMM000": "Unknown error has occured",
	"ENLIBCOMM007": "Module must be started synchronously, skipping asynchronous start",
	"ENLIBCOMM008": "Module must be started asynchronously, skipping synchronous start",
	"ENLIBCOMM009": "Module failed to start",
	"ENLIBCOMM010": "Module started successfully",
	"ENLIBCOMM011": "Module status",
	"ENLIBCOMM012": "Reloading module with live config",
	"ENLIBCONF010": "Failed to load config file from given path",
	"ENLIBCONF011": "Failed to unmarshal config file from given path; it was found, but the format was not as expected",
	"ENLIBHTTP009": "Internal routine failure serving HTTP",
	"ENLIBHTTP010": "Could not add this route as one with the given name already exists",
	"ENLIBHTTP011": "Could not add this route as its of path type regex, and the provided regex is invalid",
	"ENLIBHTTP015": "Could not delete this route as one with the given name doesn't exist",
	"ENLIBHTTP050": "HTTP request inbound",
}

func (pack ENPack) New() map[string]string {
	return EN_DATA
}
