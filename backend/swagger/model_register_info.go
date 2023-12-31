/*
 * PUN street Universal Access - OpenAPI 3.0
 *
 * pua
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RegisterInfo struct {
	UserName string `json:"user_name"`

	UserEmail string `json:"user_email"`

	Password string `json:"password"`

	Phone string `json:"phone"`

	Address string `json:"address"`

	Birthday string `json:"birthday"`

	StoreRegisterInfo *StoreRegisterInfo `json:"StoreRegisterInfo"`
}
