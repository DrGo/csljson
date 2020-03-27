package csljson

type CslData struct {
	Items []*Csljson `json:"items"`
}

type Csljson struct {
	Abstract                 string  `json:"abstract,omitempty"`
	Accessed                 Date    `json:"accessed"`
	Annote                   string  `json:"annote,omitempty"`
	Archive                  string  `json:"archive,omitempty"`
	ArchiveLocation          string  `json:"archive_location,omitempty"`
	ArchivePlace             string  `json:"archive-place,omitempty"`
	Author                   []Name  `json:"author"`
	Authority                string  `json:"authority,omitempty"`
	CallNumber               string  `json:"call-number,omitempty"`
	ChapterNumber            string  `json:"chapter-number,omitempty"`
	CitationLabel            string  `json:"citation-label,omitempty"`
	CitationNumber           string  `json:"citation-number,omitempty"`
	CollectionEditor         []Name  `json:"collection-editor"`
	CollectionNumber         string  `json:"collection-number,omitempty"`
	CollectionTitle          string  `json:"collection-title,omitempty"`
	Composer                 []Name  `json:"composer"`
	Container                Date    `json:"container"`
	ContainerAuthor          []Name  `json:"container-author"`
	ContainerTitle           string  `json:"container-title,omitempty"`
	ContainerTitleShort      string  `json:"container-title-short,omitempty"`
	Dimensions               string  `json:"dimensions,omitempty"`
	Director                 []Name  `json:"director"`
	DOI                      string  `json:"DOI"`
	Edition                  Variant `json:"edition,omitempty"`
	Editor                   []Name  `json:"editor"`
	EditorialDirector        []Name  `json:"editorial-director"`
	Event                    string  `json:"event,omitempty"`
	EventDate                Date    `json:"event-date"`
	EventPlace               string  `json:"event-place,omitempty"`
	FirstReferenceNoteNumber string  `json:"first-reference-note-number,omitempty"`
	Genre                    string  `json:"genre,omitempty"`
	ID                       Variant `json:"id"`
	Illustrator              []Name  `json:"illustrator"`
	Interviewer              []Name  `json:"interviewer"`
	ISBN                     string  `json:"ISBN,omitempty"`
	ISSN                     string  `json:"ISSN,omitempty"`
	Issue                    Variant `json:"issue"`
	Issued                   Date    `json:"issued"`
	Jurisdiction             string  `json:"jurisdiction,omitempty"`
	Keyword                  string  `json:"keyword,omitempty"`
	Locator                  string  `json:"locator,omitempty"`
	Medium                   string  `json:"medium,omitempty"`
	Note                     string  `json:"note,omitempty"`
	Number                   Variant `json:"number,omitempty"`
	NumberOfPages            Variant `json:"number-of-pages,omitempty"`
	NumberOfVolumes          Variant `json:"number-of-volumes,omitempty"`
	OriginalAuthor           []Name  `json:"original-author"`
	OriginalDate             Date    `json:"original-date"`
	OriginalPublisher        string  `json:"original-publisher,omitempty"`
	OriginalPublisherPlace   string  `json:"original-publisher-place,omitempty"`
	OriginalTitle            string  `json:"original-title,omitempty"`
	Page                     string  `json:"page,omitempty"`
	PageFirst                string  `json:"page-first,omitempty"`
	PMCID                    string  `json:"PMCID,omitempty"`
	PMID                     string  `json:"PMID,omitempty"`
	Publisher                string  `json:"publisher,omitempty"`
	PublisherPlace           string  `json:"publisher-place,omitempty"`
	Recipient                []Name  `json:"recipient"`
	References               string  `json:"references,omitempty"`
	ReviewedAuthor           []Name  `json:"reviewed-author"`
	ReviewedTitle            string  `json:"reviewed-title,omitempty"`
	Scale                    string  `json:"scale,omitempty"`
	Section                  string  `json:"section,omitempty"`
	Source                   string  `json:"source,omitempty"`
	Status                   string  `json:"status,omitempty"`
	Submitted                Date    `json:"submitted"`
	Title                    string  `json:"title,omitempty"`
	TitleShort               string  `json:"title-short,omitempty"`
	Translator               []Name  `json:"translator"`
	Type                     string  `json:"type"`
	URL                      string  `json:"URL,omitempty"`
	Version                  string  `json:"version,omitempty"`
	Volume                   Variant `json:"volume,omitempty"`
	YearSuffix               string  `json:"year-suffix,omitempty"`
}

type Name struct {
	CommaSuffix         Variant `json:"comma-suffix,omitempty"`
	DroppingParticle    string  `json:"dropping-particle,omitempty"`
	Family              string  `json:"family,omitempty"`
	Given               string  `json:"given,omitempty"`
	Literal             string  `json:"literal,omitempty"`
	NonDroppingParticle string  `json:"non-dropping-particle,omitempty"`
	ParseNames          Variant `json:"parse-names,omitempty"`
	StaticOrdering      Variant `json:"static-ordering,omitempty"`
	Suffix              string  `json:"suffix,omitempty"`
}

type Date struct {
	Circa     Variant     `json:"circa,omitempty"`
	DateParts []DateParts `json:"date-parts,omitempty"`
	Literal   string      `json:"literal,omitempty"`
	Raw       string      `json:"raw,omitempty"`
	Season    Variant     `json:"season,omitempty"`
}

type DateParts []Variant

// TODO: add constants

// "type": {
// 	"type": "string",
// 	"required": true,
// 	"enum" : [
// 			"article",
// 			"article-journal",
// 			"article-magazine",
// 			"article-newspaper",
// 			"bill",
// 			"book",
// 			"broadcast",
// 			"chapter",
// 			"dataset",
// 			"entry",
// 			"entry-dictionary",
// 			"entry-encyclopedia",
// 			"figure",
// 			"graphic",
// 			"interview",
// 			"legal_case",
// 			"legislation",
// 			"manuscript",
// 			"map",
// 			"motion_picture",
// 			"musical_score",
// 			"pamphlet",
// 			"paper-conference",
// 			"patent",
// 			"personal_communication",
// 			"post",
// 			"post-weblog",
// 			"report",
// 			"review",
// 			"review-book",
// 			"song",
// 			"speech",
// 			"thesis",
// 			"treaty",
// 			"webpage"
// 	]
// }
