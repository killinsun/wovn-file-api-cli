package domain

import (
	"testing"

	"github.com/killinsun/wovn-file-api-cli/mock"
)

func TestNewURLListStore(t *testing.T) {
	path := "./url_test.txt"
	want := &UrlListStore{
		UrlStore: []UrlStore{
			{Value: "https://example.com"},
			{Value: "https://example.com/path"},
			{Value: "https://example.com/path/to"},
			{Value: "https://example.com/path/to/directory"},
		},
	}

	spyRepo := SpyFsRepository{}
	got := NewUrlListStore(&spyRepo, path)

	for i := range want.UrlStore {
		if want.UrlStore[i].Value != got.UrlStore[i].Value {
			t.Errorf(
				"cannot get URLListStore correctly. want: %q, but got: %q",
				want.UrlStore[i].Value,
				got.UrlStore[i].Value,
			)
		}
	}
}

func TestGetAllTranslatedHTML(t *testing.T) {
	path := "./url_test.txt"

	spyRepo := SpyFsRepository{}
	mockRepo := mockWovnAPIRepository{}

	urlListStore := NewUrlListStore(&spyRepo, path)
	err := urlListStore.GetAllTranslatedHTML(&spyRepo, &mockRepo, mock.Config)

	if err != nil {
		t.Errorf("cannot call GetAllTranslatedHTML correctly. error: %v", err)
	}
}

func BenchmarkGetAllTranslatedHTML(b *testing.B) {
	path := "./url_test.txt"

	spyRepo := SpyFsRepository{}
	mockRepo := mockWovnAPIRepository{}

	urlListStore := NewUrlListStore(&spyRepo, path)
	urlListStore.GetAllTranslatedHTML(&spyRepo, &mockRepo, mock.Config)

}

func TestSendReportByURLList(t *testing.T) {
	path := "./url_test.txt"

	spyRepo := SpyFsRepository{}
	mockRepo := mockWovnAPIRepository{}

	urlListStore := NewUrlListStore(&spyRepo, path)
	err := urlListStore.SendReportByURLList(&mockRepo, mock.Config)

	if err != nil {
		t.Errorf("cannot call SendReportByURLList correctly. error: %v", err)
	}
}

func BenchmarkSendReportByURLList(b *testing.B) {
	path := "./url_test.txt"

	spyRepo := SpyFsRepository{}
	mockRepo := mockWovnAPIRepository{}

	urlListStore := NewUrlListStore(&spyRepo, path)
	urlListStore.SendReportByURLList(&mockRepo, mock.Config)

}
