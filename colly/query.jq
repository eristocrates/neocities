map(select(.hasModernList == true) | {fullName, listLink, teamListLink, testString, teams})
#map(select(.hasModernList == true) | {fullName, listLink, testString, teams: (.teams | map(.teamName))})
#map(select(.hasModernList == true) | {fullName, listLink, testString, testStringArray })
#map(select(.shortName == "bofnt") | {fullName, listLink, testString, testStringArray })