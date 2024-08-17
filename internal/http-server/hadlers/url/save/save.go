package save

import (
	"golang.org/x/exp/slog"
	"net/http"
)

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
