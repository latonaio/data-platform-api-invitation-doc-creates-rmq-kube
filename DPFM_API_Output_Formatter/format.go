package dpfm_api_output_formatter

import (
	"data-platform-api-invitation-doc-creates-rmq-kube/DPFM_API_Caller/requests"
)

func ConvertToHeaderDoc(headerDoc *HeaderDoc) *requests.HeaderDoc {
	pm := &requests.HeaderDoc{}

	pm = &requests.HeaderDoc{
		Invitation:	              headerDoc.Invitation,
		DocType:                  headerDoc.DocType,
		DocVersionID:             headerDoc.DocVersionID,
		DocID:                    headerDoc.DocID,
		FileExtension:            headerDoc.FileExtension,
		FileName:                 headerDoc.FileName,
		FilePath:                 headerDoc.FilePath,
		DocIssuerBusinessPartner: headerDoc.DocIssuerBusinessPartner,
	}

	return pm
}
