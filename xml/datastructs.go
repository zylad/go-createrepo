// Package main provides ...
package xml

//TODO, some defaults should be used here.
type OpenChecksumType struct {
	Type	string	`xml:"type,attr,omitempty"`
	Value	string	`xml:",chardata"`
}

type ChecksumType struct {
	Type	string	`xml:"type,attr,omitempty"`
	Value	string	`xml:",chardata"`
}

type LocationType struct {
	Href	string	`xml:"href,attr,omitempty"`
	Value	string	`xml:",chardata"`
}

type DataType struct {
	Type			string					`xml:"type,attr,omitempty"`
	Location		OpenChecksumType		`xml:"location"`
	Checksum		ChecksumType			`xml:"checksum"`
	Timestamp		string					`xml:"timestamp"`
	OpenChecksum	OpenChecksumType		`xml:"open-checksum"`
	Size			int						`xml:"size"`
	OpenSize		int						`xml:"open-size"`
}

type Header struct {
	XMLName		xml.Name	`xml:"http://linux.duke.edu/metadata/repo repomd"`
	Revision	string		`xml:"revision"`
	Data		[]DataType	`xml:"data"`
}
