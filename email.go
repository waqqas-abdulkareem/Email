package email

import(
	"fmt"
	"bytes"
	"strings"
	_"encoding/base64"
	"regexp"
)

type Part struct{
	Headers map[string]string
	Payload []byte
}

type Email struct{
	Headers map[string]string
	Body string
	Part []Part
}

func Parse(email string) (*Email, error){
	Headers := parseHeaders(email)

	return &Email{Headers: Headers},nil
}

func parseHeaders(email string) map[string]string{
	headers := make(map[string]string)

	headersSectionBuffer := bytes.NewBufferString("")
	for _,line := range strings.Split(email,"\n"){
		if len(line) == 0 {
			break
		}
		headersSectionBuffer.WriteString(line)
	}

	headersSection := headersSectionBuffer.String()
	headerNameRegex := regexp.MustCompile("([A-Z][A-Za-z0-9\\-_]*):")
	headerNameRanges := headerNameRegex.FindAllStringIndex(headersSection,-1)
	if  len(headerNameRanges) == 0{
		return headers;//empty map
	}
	for i,headerNameRange := range headerNameRanges{
		
		headerNameStart := headerNameRange[0]
		headerNameEnd := headerNameRange[1]-1//-1 to ignore : inlcuded in regex
		headerValueStart := headerNameRange[1]+1
		headerValueEnd := headerValueStart

		if i + 1 >= len(headerNameRanges) {
			//read until end of header section
			headerValueEnd = len(headersSection)
		}else{
			//read until start of next header
			nextHeaderRange := headerNameRanges[i + 1]
			headerValueEnd = nextHeaderRange[0]
		}

		headerName := headersSection[headerNameStart:headerNameEnd]
		headerValue := headersSection[headerValueStart:headerValueEnd]
		
		headers[headerName] = headerValue
	}

	return headers
}


func (e Email) String() string{
	buffer := bytes.NewBufferString("")
	for name,value := range e.Headers {
		buffer.WriteString(fmt.Sprintf("%s:%s\n",name,value))
	}
	return buffer.String()
}