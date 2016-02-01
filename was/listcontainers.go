package was

import (
	"encoding/xml"
	"net/http"

	"github.com/martini-contrib/render"
)

/*
ListContainers endpoint
https://msdn.microsoft.com/en-us/library/azure/dd179352.aspx

The List Containers operation returns a list of the containers under the specified account.
*/
func ListContainers(req *http.Request, r render.Render) {

	resp := ListContainersResponse{}

	r.XML(200, resp)
}

/*
ListContainersRequest ...
*/
type ListContainersRequest struct {

	//   Optional. Filters the results to return only containers whose name begins with the specified prefix.
	prefix string
	//  Optional. A string value that identifies the portion of the list to be returned with the next list operation. The operation returns a marker value within the response body if the list returned was not complete. The marker value may then be used in a subsequent call to request the next set of list items.
	//The marker value is opaque to the client.
	marker string
	// Optional. Specifies the maximum number of containers to return. If the request does not specify maxresults, or specifies a value greater than 5,000, the server will return up to 5,000 items. If the parameter is set to a value less than or equal to zero, the server will return status code 400 (Bad Request).
	maxresults uint
	// (include=metadata) Optional. Include this parameter to specify that the container's metadata be returned as part of the response body.
	include string
	//Optional. The timeout parameter is expressed in seconds.
	timeout uint
}

/*
ListContainersRequestHeaders ...
*/
type ListContainersRequestHeaders struct {

	//	Required. Specifies the authentication scheme, account name, and signature. For more information, see Authentication for the Azure Storage Services.
	authorization string `header:"Authorization"`
	// Required. Specifies the Coordinated Universal Time (UTC) for the request. For more information, see Authentication for the Azure Storage Services.
	date    string `header:"Date"`
	xmsdate string `header:"x-ms-date"`
	// Required for all authenticated requests. Specifies the version of the operation to use for this request. For more information, see Versioning for the Azure Storage Services.
	xmsversion string `header:"x-ms-version"`
	// Optional. Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when storage analytics logging is enabled.
	xmsclientrequestid string `header:"x-ms-client-request-id"`
}

/*
   <?xml version="1.0" encoding="utf-8"?>
   <EnumerationResults ServiceEndpoint="https://myaccount.blob.core.windows.net">
     <Prefix>string-value</Prefix>
     <Marker>string-value</Marker>
     <MaxResults>int-value</MaxResults>
     <Containers>
       <Container>
         <Name>container-name</Name>
         <Properties>
           <Last-Modified>date/time-value</Last-Modified>
           <Etag>etag</Etag>
           <LeaseStatus>locked | unlocked</LeaseStatus>
           <LeaseState>available | leased | expired | breaking | broken</LeaseState>
           <LeaseDuration>infinite | fixed</LeaseDuration>
         </Properties>
         <Metadata>
           <metadata-name>value</metadata-name>
         </Metadata>
       </Container>
     </Containers>
     <NextMarker>marker-value</NextMarker>
   </EnumerationResults>
*/

/*
ListContainersResponse ...
*/
type ListContainersResponse struct {
	XMLName         xml.Name                          `xml:"EnumerationResults"`
	ServiceEndpoint string                            `xml:"ServiceEndpoint,attr"`
	Prefix          string                            `xml:"Prefix"`
	Marker          string                            `xml:"Marker"`
	MaxResults      uint                              `xml:"MaxResults"`
	NextMarker      string                            `xml:"NextMarker"`
	Containers      []ListContainersResponseContainer `xml:"Containers"`
}

/*
ListContainersResponseContainer ...
*/
type ListContainersResponseContainer struct {
	XMLName    xml.Name                                    `xml:"Container"`
	Name       string                                      `xml:"name"`
	Properties []ListContainersResponseContainerProperties `xml:"Properties"`
	Metadata   []ListContainersResponseContainerMetadata   `xml:"Metadata"`
}

/*
ListContainersResponseContainerProperties ...
*/
type ListContainersResponseContainerProperties struct {
	LastModified  string `xml:"Last-Modiefied"`
	Etag          string `xml:"Etag"`
	LeaseStatus   string `xml:"LeaseStatus"`
	LeaseState    string `xml:"LeaseState"`
	LeaseDuration string `xml:"LeaseDuration"`
}

/*
ListContainersResponseContainerMetadata ...
*/
type ListContainersResponseContainerMetadata struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}
