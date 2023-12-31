/*
 * PUN street Universal Access - OpenAPI 3.0
 *
 * pua
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type StoreRegisterInfo struct {
	Name string `json:"name"`

	Description string `json:"description"`

	Address string `json:"address"`

	ShippingFee int64 `json:"shipping_fee"`

	Picture string `json:"picture"`
}
