package main

func createFakeData() {

	jasonStore := Store{Name: "jason", DefaultTheme: "simple"}
	jasonStore.Create()

	hostTable := HostTable{Host: "jason.mystore.com:3000", StoreId: jasonStore.Id}
	hostTable.Create()

	//themes
	simpleTheme := Theme{Name: "simple", IsDefault: true, StoreId: jasonStore.Id}
	simpleTheme.Create()

	redTheme := Theme{Name: "red", IsDefault: false, StoreId: jasonStore.Id}
	redTheme.Create()

	//templates
	basicTemplate := Template{Name: "basic", StoreId: jasonStore.Id, ThemeId: simpleTheme.Id}
	basicTemplate.Content =
		`<html>
		    <head>
		        <title>{{.Title}}</title>
		        <link rel="stylesheet" href="/public/css/default.css"/>
		    </head>
		    <body>
				{{.Content}}
		    </body>
		</html>`
	basicTemplate.create()

	homeTemplate := Template{Name: "home", StoreId: jasonStore.Id, ThemeId: simpleTheme.Id}
	homeTemplate.Content =
		`<html>
			<head>
			    <title>{{.Title}}</title>
			</head>
			<body>
			    {{.Content}}
			</body>
		</html>`
	homeTemplate.create()

	//pages
	homePage := Page{StoreId: jasonStore.Id, TemplateName: homeTemplate.Name, Name: "home"}
	homePage.Content = "Hello Jason"
	homePage.create()

	helloPage := Page{StoreId: jasonStore.Id, TemplateName: basicTemplate.Name, Name: "hello"}
	helloPage.Content = "Welcome to my hello page."
	helloPage.create()

}
