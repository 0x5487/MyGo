package main

func createFakeData() {

	jasonStore := Store{Name: "jason", DefaultTheme: "simple"}
	jasonStore.Create()

	hostTable := HostTable{Host: "jason.mystore.com:3000", StoreId: jasonStore.Id}
	hostTable.Create()

	simpleTheme := Theme{Name: "simple", IsDefault: true, StoreId: jasonStore.Id}
	simpleTheme.Create()

	basicTemplate := Template{Name: "basic", StoreId: jasonStore.Id, ThemeId: simpleTheme.Id}
	basicTemplate.create()

	homeTemplate := Template{Name: "home", StoreId: jasonStore.Id, ThemeId: simpleTheme.Id}
	homeTemplate.create()

	redTheme := Theme{Name: "red", IsDefault: false, StoreId: jasonStore.Id}
	redTheme.Create()

	basicTemplate = Template{Name: "basic", StoreId: jasonStore.Id, ThemeId: redTheme.Id}
	basicTemplate.create()

	homeTemplate = Template{Name: "home", StoreId: jasonStore.Id, ThemeId: redTheme.Id}
	homeTemplate.create()

}
