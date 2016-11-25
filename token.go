package anthon

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS
	// Literals
	IDENT // main
	// Misc characters
	ASTERISK   // *
	COMMA      // ,
	DosPunt    // :
	ParDer     // (
	ParIsq     // )
	PuntCom    // ;
	LlaveDer   // {
	LlaveIsq   // }
	SaltoDeL   //n
	Asignaicon // =
	AgruppDER  // [
	AgruppISQ  // ]

	// Keywords
	SELECT
	FROM
	LOLIF
	LOLSWITCH
	LOLEND
	LOLINT // Tipodedato
	LOLFUNC
	LOLMAIN
	LOLELSE
	LOLFOR
	LOLELSEIF
	VAR // variables
	LOLBREAK
	LOLWHILE
	LOLDO
	LOLTRY
	LOLCATCH
	LOLFINALLY
	LOLPOINT
	LOLNEW
	LOLFIXED
	LOLFOREACH
	LOLUNCHECKED
	LOLIMPORT
	
)
