package converter

var mm2enChar = map[string]string{"က": "Ka", "ခ": "Kh", "ဂ": "Ga", "ဃ": "Gh", "င": "Ng", "စ": "Ca", "ဆ": "Ch", "ဇ": "Ja", "ဈ": "Jh", "ည": "Ny", "ဎ": "Dd", "ဏ": "Nn", "တ": "Ta", "ထ": "Th", "ဒ": "Da", "ဓ": "Dh", "န": "Na", "ပ": "Pa", "ဖ": "Ph", "ဗ": "Ba", "ဘ": "Bh", "မ": "Ma", "ယ": "Ya", "ရ": "Ra", "လ": "La", "ဝ": "Wa", "သ": "Sa", "ဟ": "Ha", "ဠ": "Ll", "ဥ": "U", "ဧ": "E"}
var en2mmChar = map[string]string{"Ka": "က", "Kh": "ခ", "Ga": "ဂ", "Gh": "ဃ", "Ng": "င", "Ca": "စ", "Ch": "ဆ", "Ja": "ဇ", "Jh": "ဈ", "Ny": "ည", "Dd": "ဎ", "Nn": "ဏ", "Ta": "တ", "Th": "ထ", "Da": "ဒ", "Dh": "ဓ", "Na": "န", "Pa": "ပ", "Ph": "ဖ", "Ba": "ဗ", "Bh": "ဘ", "Ma": "မ", "Ya": "ယ", "Ra": "ရ", "La": "လ", "Wa": "ဝ", "Sa": "သ", "Ha": "ဟ", "Ll": "ဠ", "U": "ဥ", "E": "ဧ"}
var nrcmm2en = map[string]string{`(\(ဧည့်\))`: `(AC)`, `(\(ပြု\))`: "(NC)", `(\(နိုင်\))`: "(C)", `(\(စ\))`: "(V)", "သီ": "N", `(\(သ\))`: "(M)"}
var nrcen2MM = map[string]string{`(\(AC\))`: "(ဧည့်)", `(\(NC\))`: "(ပြု)", `(\(C)\)`: "(နိုင်)", `(\(V\))`: "(စ)", "N": "သီ", `(\(M\))`: "(သ)"}

var mmNRCformatRegex = `([၀-၉]{1,2})/([ကခဂဃငစဆဇဈညဎဏတထဒဓနပဖဗဘမယရလဝသဟဠဥဧ]{3})(\([နိုင်,ဧည့်,ပြု,စ,သ,သီ]{1,6}\))([၀-၉]{6})$`
var enNRCformatRegex = `([1-9]{1,2})/([Ka|Kh|Ga|Gh|Ng|Ca|Ch|Ja|Jh|Ny|Dd|Nn|Ta|Th|Da|Dh|Na|Pa|Ph|Ba|Bh|Ma|Ya|Ra|La|Wa|Sa|Ha|Ll|u|E]{3,10})(\([C|AC|NC|V|M|N]{1,2}\))([0-9]{6})$`
