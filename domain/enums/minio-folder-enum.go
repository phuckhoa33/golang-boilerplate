package enums

// MinioFolderEnum is a type that represents the folder name in Minio
type MinioFolderEnum string

const (
	// MinioFolderEnumUser is a folder name for user
	MinioFolderEnumUser MinioFolderEnum = "user"
	// MinioFolderEnumProduct is a folder name for product
	MinioFolderEnumProduct MinioFolderEnum = "product"
	// MinioFolderEnumCategory is a folder name for category
	MinioFolderEnumCategory MinioFolderEnum = "category"
	// MinioFolderEnumBrand is a folder name for brand
	MinioFolderEnumBrand MinioFolderEnum = "brand"
	// MinioFolderEnumBanner is a folder name for banner
	MinioFolderEnumBanner MinioFolderEnum = "banner"
	// MinioFolderEnumArticle is a folder name for article
	MinioFolderEnumArticle MinioFolderEnum = "article"
	// MinioFolderEnumOther is a folder name for other
	MinioFolderEnumOther MinioFolderEnum = "other"
)
