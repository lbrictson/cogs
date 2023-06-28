package web

import "embed"

// Assets represents the embedded files.
// You can add more files here by just extending this line, they will all be in the go executable
//
//go:embed templates/* static/css/* static/js/* static/img/*
var Assets embed.FS
