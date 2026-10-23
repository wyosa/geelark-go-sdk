package app

// VersionInfo is an available application version, including its ID and code.
type VersionInfo = ListAvailableResponseItemsItemAppVersionInfoListItem

// AvailableApp is an application and its installable versions.
type AvailableApp = ListAvailableResponseItemsItem

// InstalledApp is an application reported by the installed-applications API.
type InstalledApp = ListInstalledResponseItemsItem
