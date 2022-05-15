package azure

import "encoding/xml"

type EnumerationResults struct {
	XMLName         xml.Name    `xml:"EnumerationResults"`
	ServiceEndpoint string      `xml:"ServiceEndpoint"`
	Containers      []Container `xml:"Containers>Container"` //Select inner container name because of Azure XML structure
}

type Container struct {
	XMLName    xml.Name   `xml:"Container"`
	Name       string     `xml:"Name"`
	Properties Properties `xml:"Properties"`
}

type Properties struct {
	LastModified                          string `xml:"Last-Modified"`
	Etag                                  string `xml:"Etag"`
	LeaseStatus                           string `xml:"LeaseStatus"`
	LeaseState                            string `xml:"LeaseState"`
	PublicAccess                          string `xml:"PublicAccess"`
	DefaultEncryptionScope                string `xml:"DefaultEncryptionScope"`
	DenyEncryptionScopeOverride           string `xml:"DenyEncryptionScopeOverride"`
	HasImmutabilityPolicy                 bool   `xml:"HasImmutabilityPolicy"`
	HasLegalHold                          bool   `xml:"HasLegalHold"`
	ImmutableStorageWithVersioningEnabled bool   `xml:"ImmutableStorageWithVersioningEnabled"`
}
