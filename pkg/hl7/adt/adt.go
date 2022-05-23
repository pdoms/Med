package adt

type ADT_CONFIG struct {
	MSH MSH_CONFIG
	EVN EVN
	PID PID
	PD1 PD1
	RO1 RO1
	NK1 NK1
	PV1 PV1
	PV2 PV2
	RO2 RO2
	DB1 DB1
	OBX OBX
	Al1 Al1
	DG1 DG1
	DRG DRG
	PR1 PR1
	RO3 RO3
	GT1 GT1
	IN1 IN1
	IN2 IN2
	IN3 IN3
	RO4 RO4
	ACC ACC
	UB1 UB1
	UB2 UB2
	PDA PDA
}

type MSH_CONFIG struct {
	FieldSeparator                      string        `json:"FieldSeparator"`
	EndodingCharacters                  string        `json:"EncodingCharacters"`
	SendingApplication                  string        `json:"SendingApplication"`
	SendingFacility                     string        `json:"SendingFacility"`
	ReceivingApplication                string        `json:"ReceivingApplication"`
	ReceivingFacility                   string        `json:"ReceivingFacility"`
	DateTimeOfMessage                   string        `json:"DateTimeOfMessage"`
	Security                            string        `json:"Security"`
	MessageType                         string        `json:"MessageType"`
	MessageContrilId                    string        `json:"MessageControlId"`
	ProcessingId                        string        `json:"ProcessingId"`
	VersionId                           string        `json:"VersionId"`
	SequenceNumber                      string        `json:"SequenceNumber"`
	ContinuationNumber                  string        `json:"ContinuationNumber"`
	AcceptAcknowledgeType               string        `json:"AcceptAcknowledgeType"`
	ApplicationAcknowledgmentType       string        `json:"ApplicationAcknowledgmentType"`
	CountryCode                         string        `json:"CountryCode"`
	CharacterSet                        string        `json:"CharacterSet"`
	PrincipalLanguageOfMessage          string        `json:"PrincipalLanguageOfMessage"`
	AlternateCharacterSetHandlingScheme string        `json:"AlternateCharacterSetHandlingScheme"`
	MessgeProfileIdentifier             string        `json:"MessageProfileIdentifier"`
	SubFields                           MSH_SubFields `json:"SubFields"`
}

type MSH_SubFields struct {
	MessageType                MSH_SUB_MessageType                `json:"MessageType"`
	ProcessngId                MSH_SUB_ProcessingId               `json:"ProcessingId"`
	VersionIdentifier          MSH_SUB_VersionIdentifier          `json:"VersionIdentifier"`
	PrincipalLanguageOfMessage MSH_SUB_PrincipalLangaugeOfMessage `json:"PrincipalLanguageOfMessage"`
	MessageProfileIdentifier   MSH_SUB_MessageProfileIdentifier   `json:"MessageProfileIdentifier"`
}

type MSH_SUB_MessageType struct {
	MessageCode      string `json:"MessageCode"`
	TriggerEvent     string `json:"TriggerEvent"`
	MessageStructure string `json:"MessageStructure"`
}

type MSH_SUB_ProcessingId struct {
	ProcessingId   string `json:"ProcessingId"`
	ProcessingMode string `json:"ProcessingMode"`
}
type MSH_SUB_VersionIdentifier struct {
	VersonId                      string `json:"VersionId"`
	InternationalizationCode      string `json:"InternationalizationCode"`
	InternationalizationVersionId string `json:"InternationalizationVersionId"`
}
type MSH_SUB_PrincipalLangaugeOfMessage struct {
	Identifier                string `json:"Identifier"`
	Text                      string `json:"Text"`
	NameOfCodingSystem        string `json:"NameOfCodingSystem"`
	AlternateIdentifier       string `json:"AlternateIdentifier"`
	NameOfAlternateIdentifier string `json:"NameOfAlternateIdentifier"`
}
type MSH_SUB_MessageProfileIdentifier struct {
	EntitiyIdentifier string `json:"EntityIdentifier"`
	NamesSpaceId      string `json:"NameSpaceId"`
	UniversalId       string `json:"UniversalId"`
	UniversalIdType   string `json:"UniversalIdType"`
}

type EVN struct{}
type PID struct{}
type PD1 struct{}
type RO1 struct{}
type NK1 struct{}
type PV1 struct{}
type PV2 struct{}
type RO2 struct{}
type DB1 struct{}
type OBX struct{}
type Al1 struct{}
type DG1 struct{}
type DRG struct{}
type PR1 struct{}
type RO3 struct{}
type GT1 struct{}
type IN1 struct{}
type IN2 struct{}
type IN3 struct{}
type RO4 struct{}
type ACC struct{}
type UB1 struct{}
type UB2 struct{}
type PDA struct{}

//actual adt structs
