package main

var systems = []system{{
	Name:        "Datasektionen.se",
	Description: "Hitta all information och nyheter om sektionen",
	URL:         "https://datasektionen.se",
	Icon:        "https://datasektionen.se/icons/favicon-96x96.png",
	Color:       "#ff9800",
}, {
	Name:        "Cashflow",
	Description: "Få tillbaka pengar för dina utlägg",
	URL:         "https://cashflow.datasektionen.se",
	Icon:        "https://cashflow.datasektionen.se/static/images/logos/favicon-96x96.png",
	Color:       "#216C2A",
}, {
	Name:        "METAmost",
	Description: "Tyvärr inte Matrix eller Zulip",
	URL:         "https://mattermost.datasektionen.se",
	Icon:        "https://mattermost.datasektionen.se/static/images/favicon/favicon-default-96x96.png",
	Color:       "#1b1d21",
}, {
	Name:        "Bokningssystemet",
	Description: "Boka META, Mötesrummet, bilarna, eller annat",
	URL:         "https://bokning.datasektionen.se",
	Icon:        "https://bokning.datasektionen.se/logos/favicon-96x96.png",
	Color:       "#03a9f4",
}, {
	Name:        "Valsystemet",
	Description: "Se till att någon tar ansvar",
	URL:         "https://val.datasektionen.se",
	Icon:        "https://val.datasektionen.se/images/logos/favicon-96x96.png",
	Color:       "#07d",
	Sensitive:   true,
}, {
	Name:        "Styrdokument",
	Description: "Stadgar, reglemente, PM, och policydokument",
	URL:         "https://styrdokument.datasektionen.se",
	Icon:        "https://styrdokument.datasektionen.se/static/favicon-96x96.png",
	Color:       "#ff9800",
}, {
	Name:        "Budgetsystemet",
	Description: "Se vad sektionen gör av sina pengar",
	URL:         "https://budget.datasektionen.se",
	Icon:        "https://budget.datasektionen.se/static/images/icons/favicon-96x96.png",
	Color:       "#3d3846",
}, {
	Name:        "Dfunkt",
	Description: "Vem gör vad egentligen?",
	URL:         "https://dfunkt.datasektionen.se",
	Icon:        "https://dfunkt.datasektionen.se/images/logos/favicon-96x96.png",
	Color:       "#9c27b0",
}, {
	Name:        "dJulkalendern",
	Description: "Vem är bäst på att data™?",
	URL:         "https://djulkalendern.se",
	Icon:        "https://djulkalendern.se/images/djuldanke.svg",
	Color:       "#B90e0a",
}, {
	Name:        "Aaallt",
	Description: "Det här \"systemet\"",
	URL:         "https://aaallt.datasektionen.se",
	Icon:        "https://aaallt.datasektionen.se/favicon-96x96.png",
	Color:       "#ffca28",
}, {
	Name:        "Betting",
	Description: "När slutar SM?",
	URL:         "https://betting.datasektionen.se",
	Icon:        "https://betting.datasektionen.se/images/icons/favicon-96x96.png",
	Color:       "#78909c",
}, {
	Name:        "Damm",
	Description: "Datasektionens historia",
	URL:         "https://damm.datasektionen.se",
	Icon:        "https://damm.datasektionen.se/images/icons/favicon-96x96.png",
	Color:       "#000000",
}, {
	Name:        "dbuggen",
	Description: "Når ut till hundratals genom allt förutom TV",
	URL:         "https://dbu.gg",
	Icon:        "https://dbu.gg/favicon/favicon-96x96.png",
	Color:       "#444",
	IconStyle:   "border-radius: 50%",
}, {
	Name:        "ddagen.se",
	Description: "Allt om D-dagen",
	URL:         "https://ddagen.se",
	Icon:        "https://ddagen.se/img/favicon.ico",
	Color:       "#ec5f99",
}, {
	Name:        "/dev/audio",
	Description: "Starkt är vackert",
	URL:         "https://audio.datasektionen.se",
	Icon:        "https://audio.datasektionen.se/favicon-32x32.png",
	Color:       "#ff7043",
}, {
	Name:        "dJubileet",
	Description: "Datasektionens 45-årsjubileum",
	URL:         "https://djubileet.se",
	Icon:        "https://djubileet.se/favicon.ico",
	Color:       "#4E0093",
}, {
	Name:        "Duckumentation",
	Description: "Dokumentation av API:er",
	URL:         "https://duckumentation.datasektionen.se",
	Icon:        "https://duckumentation.datasektionen.se/images/icons/favicon-96x96.png",
	Color:       "#4caf50",
}, {
	Name:        "DUrn",
	Description: "Röstningssystem vid urnval",
	URL:         "https://durn.datasektionen.se",
	Icon:        "https://durn.datasektionen.se/public/favicon.png",
	Color:       "#0033cc",
}, {
	Name:        "GitHub",
	Description: "Datasektionens GitHub",
	URL:         "https://github.com/datasektionen",
	Icon:        "https://github.com/favicon.ico",
	Color:       "#000",
}, {
	Name:        "Metaspexet",
	Description: "Datas och Medias spex!",
	URL:         "https://metaspexet.se",
	Icon:        "https://metaspexet.se/images/favicon.svg",
	Color:       "#800000",
	IconStyle:   "background-color: #6f1d1b; border: 2px solid #6f1d1b; border-radius: 50%",
}, {
	Name:        "META-TV",
	Description: "Nå ut till hundratals genom META-TV",
	URL:         "https://tv.datasektionen.se",
	Icon:        "https://tv.datasektionen.se/favicon.png",
	Color:       "#009688",
}, {
	Name:        "Nyhetssystemet",
	Description: "Skriv och posta nyheter till hemsidan",
	URL:         "https://calypso.datasektionen.se",
	Icon:        "https://calypso.datasektionen.se/favicon-96x96.png",
	Color:       "#ec407a",
}, {
	Name:        "Pico",
	Description: "Länkförkortare",
	URL:         "https://dsekt.se",
	Icon:        "https://dsekt.se/favicon-96x96.png",
	Color:       "#03a9f4",
}, {
	Name:        "Pls",
	Description: "Lär dig ta ett nej",
	URL:         "https://pls.datasektionen.se",
	Icon:        "https://pls.datasektionen.se/images/logos/favicon-96x96.png",
	Color:       "#9c27b0",
}, {
	Name:        "Sektionswikin",
	Description: "Här hittar du massvis ackumulerad kunskap.",
	URL:         "https://wiki.datasektionen.se",
	Icon:        "https://wiki.datasektionen.se/resources/assets/skold_svart_thumb.png",
	Color:       "#ff9800",
}, {
	Name:        "Status",
	Description: "Ta reda på om något av våra system är nere",
	URL:         "https://status.datasektionen.se",
	Icon:        "https://userfiles.uptimerobot.com/img/623458-1649958820-icon.png",
	Color:       "#32cd32",
}, {
	Name:        "STÖn",
	Description: "Suck",
	URL:         "https://ston.datasektionen.se",
	Icon:        "https://ston.datasektionen.se/images/logos/favicon-96x96.png",
	Color:       "#ec5f99",
	Sensitive:   true,
}, {
	Name:        "Yoggi",
	Description: "Ladda upp statiska filer i sektionssyfte",
	URL:         "https://static.datasektionen.se",
	Icon:        "https://static.datasektionen.se/favicon-96x96.png",
}, {
	Name:        "METAcraft",
	Description: "Datas och Medias Minecraftserver",
	URL:         "https://metacraft.nu",
	Icon:        "https://metacraft.nu/logo-square.png",
}}
