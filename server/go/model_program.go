/*
 * radio program API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Program struct {

	Id string `json:"id"`

	Title string `json:"title"`

	Detail string `json:"detail"`

	Url string `json:"url,omitempty"`
}
