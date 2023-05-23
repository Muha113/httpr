package httpr

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGroup(t *testing.T) {
	var (
		mw1HitCtn          int8
		mw2HitCtn          int8
		mw3HitCtn          int8
		mw4HitCtn          int8
		fake1HandlerHitCtn int8
		fake2HandlerHitCtn int8
		fake3HandlerHitCtn int8

		triggerOrder     []string
		wantTriggerOrder []string
	)

	fake1Handler := func(w http.ResponseWriter, _ *http.Request, _ Params) {
		triggerOrder = append(triggerOrder, "fake1Handler")
		fake1HandlerHitCtn++
		w.WriteHeader(http.StatusOK)
	}
	fake2Handler := func(w http.ResponseWriter, _ *http.Request, _ Params) {
		triggerOrder = append(triggerOrder, "fake2Handler")
		fake2HandlerHitCtn++
		w.WriteHeader(http.StatusOK)
	}
	fake3Handler := func(w http.ResponseWriter, _ *http.Request, _ Params) {
		triggerOrder = append(triggerOrder, "fake3Handler")
		fake3HandlerHitCtn++
		w.WriteHeader(http.StatusOK)
	}
	mw1 := func(next Handle) Handle {
		return func(w http.ResponseWriter, r *http.Request, p Params) {
			triggerOrder = append(triggerOrder, "mw1")
			mw1HitCtn++
			next(w, r, p)
		}
	}
	mw2 := func(next Handle) Handle {
		return func(w http.ResponseWriter, r *http.Request, p Params) {
			triggerOrder = append(triggerOrder, "mw2")
			mw2HitCtn++
			next(w, r, p)
		}
	}
	mw3 := func(next Handle) Handle {
		return func(w http.ResponseWriter, r *http.Request, p Params) {
			triggerOrder = append(triggerOrder, "mw3")
			mw3HitCtn++
			next(w, r, p)
		}
	}
	mw4 := func(next Handle) Handle {
		return func(w http.ResponseWriter, r *http.Request, p Params) {
			triggerOrder = append(triggerOrder, "mw4")
			mw4HitCtn++
			next(w, r, p)
		}
	}

	router := New()

	baseGroup := router.Group("base", mw1)
	baseGroup.GET("get", fake1Handler)

	noMWsGroup := baseGroup.Group("no_mws")
	noMWsGroup.POST("", fake2Handler)

	twoMWsGroup := baseGroup.Group("2mws", mw2, mw3)
	twoMWsGroup.DELETE("smth/:id/delete", fake3Handler, mw4)
	twoMWsGroup.POST("smth", fake3Handler)

	triggerOrder = []string{}
	wantTriggerOrder = []string{"mw1", "fake2Handler"}
	r, _ := http.NewRequest(http.MethodPost, "/base/no_mws", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if !(w.Code == http.StatusOK && mw1HitCtn == 1 && fake2HandlerHitCtn == 1 && equalSlices(triggerOrder, wantTriggerOrder)) {
		t.Errorf("Group routing failed with router group.")
		t.FailNow()
	}

	triggerOrder = []string{}
	wantTriggerOrder = []string{"mw1", "fake1Handler"}
	r, _ = http.NewRequest(http.MethodGet, "/base/get", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if !(w.Code == http.StatusOK && mw1HitCtn == 2 && fake1HandlerHitCtn == 1 && equalSlices(triggerOrder, wantTriggerOrder)) {
		t.Errorf("Group routing failed with router group.")
		t.FailNow()
	}

	triggerOrder = []string{}
	wantTriggerOrder = []string{"mw1", "mw2", "mw3", "mw4", "fake3Handler"}
	r, _ = http.NewRequest(http.MethodDelete, "/base/2mws/smth/:id/delete", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if !(w.Code == http.StatusOK && mw1HitCtn == 3 && mw2HitCtn == 1 && mw3HitCtn == 1 && mw4HitCtn == 1 && fake3HandlerHitCtn == 1 && equalSlices(triggerOrder, wantTriggerOrder)) {
		t.Errorf("Group routing failed with router group.")
		t.FailNow()
	}

	triggerOrder = []string{}
	wantTriggerOrder = []string{"mw1", "mw2", "mw3", "fake3Handler"}
	r, _ = http.NewRequest(http.MethodPost, "/base/2mws/smth", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if !(w.Code == http.StatusOK && mw1HitCtn == 4 && mw2HitCtn == 2 && mw3HitCtn == 2 && mw4HitCtn == 1 && fake3HandlerHitCtn == 2 && equalSlices(triggerOrder, wantTriggerOrder)) {
		t.Errorf("Group routing failed with router group.")
		t.FailNow()
	}
}

func equalSlices(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
