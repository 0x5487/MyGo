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
	basicTemplate := Template{Name: "basic", StoreId: jasonStore.Id}
	basicTemplate.create()

	homeTemplate := Template{Name: "home", StoreId: jasonStore.Id}
	homeTemplate.create()

	//pages
	homePage := Page{StoreId: jasonStore.Id, TemplateId: homeTemplate.Id, Name: "home"}
	homePage.create()

}
