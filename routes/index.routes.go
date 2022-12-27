//
// index.routes.go
// Copyright (C) 2022 tinix <tinix@archlinux>
//
// Distributed under terms of the MIT license.
//

package routes

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Tinix!"))
}
