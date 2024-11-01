// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/2xSvW4TTRR9Fet+X2Wtdp2AKLZLA4qwBUKiSlxMdq/tiTw/zMwaW9FKWUdBpkKKEBEl",
	"naEwNVGUlxks0uUV0J1dmxDR7M7cn3PvOWdOIFNCK4nSWUhPwGYjFCwcn6F7bdF0uXWv0GolLVKYOxQh",
	"r43SaBzHcOM5fQfKCOYgBS7dk8cQgZtprK84RANlBJKJgNNkrDNcDqEsIzD4puAGc0gPCK4p7W9B1NEx",
	"Zo4wmgAzhs2gpF4uByqgcjem1F6v29p7uQ8RTNBYriSk0Ik78Q61K42SaQ4pPIo78S5EoJkbBRoJ0zyZ",
	"7CSFRRMCQ3T0y9FmhmtXI5Euvlr5s29+/sOfLfz8Yv3h0/rm0leXvrrx1Wd/Oj+UEEYZRl37OaT3JQXi",
	"W6sa5ux2OvTLlHQow0im9ZhnoTk5tjR3Yw+d/jc4gBT+S/74lzTmJf9yLqj0N4sXz4PsthCCmdmG1pZT",
	"tao5BR6ODS35QiXQD2AWzSSIdPBQnkb8Vl0BERRmDCmMnNM2TZLpdDqNhZI4Gyjzlpk8zpQgp5jh7Ghc",
	"64Fywo2SohEjxwErxvSyyLgIUBaC9qEbrdPfrvhwmXa7dmvpq+/tdqt1KPd6XSK3+Prr4/L29PzuetF7",
	"un537qvVz6uru+v3fn5x+2V538rNW6wlgrJf/g4AAP//XdkeMDgDAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
