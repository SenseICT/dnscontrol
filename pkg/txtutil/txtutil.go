package txtutil

import "github.com/StackExchange/dnscontrol/v4/models"

// SplitSingleLongTxt finds TXT records with a single long string and splits it
// into 255-octet chunks. This is used by providers that, when a user specifies
// one long TXT string, split it into smaller strings behind the scenes.
// This should be called from GetZoneRecordsCorrections().
// Deprecated: SplitSingleLongTxt() is deprecated. No longer needed.
func SplitSingleLongTxt(records []*models.RecordConfig) {
	// for _, rc := range records {
	// 	if rc.HasFormatIdenticalToTXT() {
	// 		s := rc.TxtStrings[0]
	// 		if len(rc.TxtStrings) == 1 && len(s) > 255 {
	// 			rc.SetTargetTXTs(splitChunks(s, 255))
	// 		}
	// 	}
	// }
}

// ToChunks returns the string as chunks of 255-octet strings (the last string being the remainder).
func ToChunks(s string) []string {
	return splitChunks(s, 255)
}

func splitChunks(buf string, lim int) []string {
	var chunk string
	chunks := make([]string, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}
