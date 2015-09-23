// Package main provides ...
package xml

type OpenChecksumType struct {
	Type	string	`xml:"type,attr,omitempty"`
	Value	string	`xml:",chardata"`
}

type DataType struct {
	Type			string		`xml:"type,attr,omitempty"`
	Location		string		`xml:"location"`
	Checksum		string		`xml:"checksum"`
	Timestamp		string		`xml:"timestamp"`
	OpenChecksum	OpenChecksumType	`xml:"open-checksum"`
}

type Header struct {
	XMLName		xml.Name	`xml:"http://linux.duke.edu/metadata/repo repomd"`
	Revision	string		`xml:"revision"`
	Data		[]DataType	`xml:"data"`
}
