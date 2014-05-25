package main

func createFakeData() {

	jasonStore := Store{Name: "jason", DefaultTheme: "simple"}
	jasonStore.Create()

	hostMapping := HostMapping{Host: "jason.mystore.com:3000", StoreId: jasonStore.Id}
	hostMapping.create()

	//themes
	simpleTheme := Theme{Name: "simple", IsDefault: true, StoreId: jasonStore.Id}
	simpleTheme.create()

	redTheme := Theme{Name: "red", IsDefault: false, StoreId: jasonStore.Id}
	redTheme.create()

	//templates
	basicTemplate := Template{Name: "basic", StoreId: jasonStore.Id, ThemeId: simpleTheme.Id}
	basicTemplate.Content =
		`<html>
		    <head>
		        <title>{{.Page.Title}}</title>
		        <link rel="stylesheet" href="/public/css/default.css"/>
		    </head>
		    <body>
				{{.Page.Content}}
		    </body>
		</html>`
	basicTemplate.create()

	homeTemplate := Template{Name: "home", StoreId: jasonStore.Id, ThemeId: simpleTheme.Id}
	homeTemplate.Content =
		`<html>
			<head>
			    <title>{{.Page.Title}}</title>
			</head>
			<body>		
				{{.Collections}}
				{{.Page.Content}}
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

	product_list_Page := Page{StoreId: jasonStore.Id, TemplateName: basicTemplate.Name, Name: "product_list"}
	product_list_Page.Content = "product list page."
	product_list_Page.create()

	//create collections
	menCollection := Collection{StoreId: jasonStore.Id, DisplayName: "DisplayName_Men", ResourceId: "Men", Tags: "shirt_tee, long_tee, polo, jeans, underwear, 領帶"}
	if err := menCollection.create(); err != nil {
		println(err.Error())
	}

	//create products

}
