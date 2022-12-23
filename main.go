package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

var (
	previousCookie = "7DEy26LL0H1YI~zeG~3qaBRSJE3j0desD-kfTbGks28~FrsN1XOhy6D_REDACTEDNnv_S-EfF-"
	siteKey        = "288922D4BE1987530B4E5D4A17952C"
	step           = 0
)

func main() {
	// just an example
	if previousCookie == "" {
		previousCookie = "null"
	}
	t := time.Now()
	rand.Seed(t.UnixNano())
	example, _ := os.ReadFile("payload_example.json")
	finger := &fingerprint{}
	json.Unmarshal(example, &finger)
	// Datadome is shit so you can math/random this but it can't be higher than 499
	finger.TimeSinceLastEvent = rand.Float64() + float64(rand.Intn(450))
	_, offset := t.Zone()
	finger.TimezoneOffset = offset
	uaLower := strings.ToLower(finger.Ua)
	if strings.Contains(uaLower, "Chrome") || strings.Contains(uaLower, "Opera") || strings.Contains(uaLower, "Other") {
		finger.EvalLength = 33
	} else if strings.Contains(uaLower, "Internet Explorer") || strings.Contains(uaLower, "Other") {
		finger.EvalLength = 39
	} else if strings.Contains(uaLower, "Safari") || strings.Contains(uaLower, "Firefox") || strings.Contains(uaLower, "Other") {
		finger.EvalLength = 37
	}
	finger.TimeSeconds = (t.UnixMilli() / 1000) - 1000
	data := &url.Values{}
	if step == 1 {
		// make sure it's not null, null would result in 0
		if len(previousCookie) != 4 {
			finger.CookieLength = 1
		}
		e := &events{PageTimestamp: t.UnixMilli(), Width: finger.ArsW, Height: finger.ArsH}
		e.PopulateEvents()
		finger.EsSigmdn = e.Sigmdn
		finger.EsMumdn = e.Mumdn
		finger.EsDistmdn = e.Dismdn
		finger.EsAngsmdn = e.Angsmdn
		finger.EsAngemdn = e.Angemdn
		// these values I've only seen on footlocker, not on reddit.
		finger.MpCx = int(e.LastPosition.X)
		finger.MpCy = int(e.LastPosition.Y)
		finger.MpTr = true
		finger.MpMx = &e.LastPosition.MovementX
		finger.MpMy = &e.LastPosition.MovementY
		finger.MpSx = e.LastPosition.ScreenX
		finger.MpSy = e.LastPosition.ScreenY
		finger.MmMd = e.WeirdMDValue
		f, _ := json.Marshal(finger)
		fmt.Println(string(f))
		data.Set("jsData", string(f))
		f, _ = json.Marshal(e)
		data.Set("eventCounters", string(f))
		data.Set("jsType", "le")
		data.Set("cid", previousCookie)
		data.Set("ddk", siteKey)
	} else {
		finger.MpMx = nil
		finger.MpMy = nil
		f, _ := json.Marshal(finger)
		fmt.Println(string(f))
		data.Set("jsData", string(f))
		data.Set("eventCounters", "[]")
		data.Set("jsType", "ch")
		data.Set("cid", previousCookie)
		data.Set("ddk", siteKey)
	}
	data.Set("Referer", "https://www.footlocker.com/")
	data.Set("request", "/")
	data.Set("responsePage", "origin")
	data.Set("ddv", "4.6.0")
	fmt.Println(data.Encode())
}

func (e *events) PopulateEvents() {
	// If you want GOOD mouse movements, don't do this, this is completely random is really shit and is just here to show the implementation in golang
	// *though I do think that this would work fine for datadome but mainly cause it's shit*
	mouseMovements := []*mouseMove{}
	eventsNum := rand.Intn(400) + 2800
	e.MouseMove = eventsNum
	sigmas := []float64{}
	mumdn := []float64{}
	dists := []float64{}
	startAngles := []float64{}
	endAngles := []float64{}
	for i := 0; i < eventsNum; i++ {
		// get a random x and random y, not gonna make it based off previous x and y cause lazy lol
		mm := &mouseMove{}
		mm.X = float64(rand.Intn(e.Width))
		mm.Y = float64(rand.Intn(e.Height))
		mm.MovementX = rand.Intn(2) - 1 // possibility to get 1, 0, or -1
		mm.MovementY = rand.Intn(2) - 1 // possibility to get 1, 0, or -1
		mm.ScreenX = rand.Intn(e.Width)
		mm.ScreenY = rand.Intn(e.Height)
		// probably a bit high to be realistic but idc
		mm.Timestamp = rand.Float64() + float64(rand.Intn(450))
		mouseMovements = append(mouseMovements, mm)
		if i%1000 == 0 || i == eventsNum-1 {
			if i == eventsNum-1 {
				e.LastPosition = mm
			}
			if e.WeirdMD == nil || e.WeirdMD.Timestamp == 0.0 {
				e.WeirdMD = mm
			} else {
				// if a cache of previous movement exists, we do the math
				a := math.Sqrt((mm.X-e.WeirdMD.X)*(mm.X-e.WeirdMD.X) + (mm.Y-e.WeirdMD.Y)*(mm.Y-e.WeirdMD.Y))
				if int(a) > e.WeirdMDValue {
					e.WeirdMDValue = int(a)
				}
			}
			sig, mu, dist, startAngle, endAngle := getValues(mouseMovements[i/1000:])
			// sigmas is broken, don't wanna fix it, feel free to do it yourself
			// the issue is really only that float64 doesn't give us enough decimal places of percision, big.float would probably work
			sigmas = append(sigmas, sig)
			mumdn = append(mumdn, mu)
			dists = append(dists, dist)
			startAngles = append(startAngles, startAngle)
			endAngles = append(endAngles, endAngle)
		}
	}
	fmt.Println(len(mouseMovements))
	e.Sigmdn = selectFromArray(sigmas)
	e.Mumdn = selectFromArray(mumdn)
	e.Dismdn = selectFromArray(dists)
	e.Angsmdn = selectFromArray(startAngles)
	e.Angemdn = selectFromArray(endAngles)
	e.Click = rand.Intn(30)
	e.Scroll = rand.Intn(10)
	e.Keydown = rand.Intn(20)
	e.Keyup = e.Keydown
}

func getValues(movements []*mouseMove) (float64, float64, float64, float64, float64) {
	x := 0.0
	q := len(movements)
	y := 0.0
	for _, m := range movements {
		l := math.Log(m.Timestamp)
		x += l
		y += l * l
	}
	B := movements[0]
	C := movements[q-1]
	mouseX := float64(B.X)
	mouseY := float64(B.Y)
	previousMouseX := float64(C.X)
	previousMouseY := float64(C.Y)
	differenceX := previousMouseX - mouseX
	differenceY := previousMouseY - mouseY
	D := 3
	if q < 4 {
		D = q - 1
	}
	E := movements[D]
	F := movements[q-D-1]
	return math.Sqrt((float64(q)*y-x*x)/float64(q)*(float64(q)-1)) / 1000.0, x / float64(q), math.Sqrt(differenceX*differenceX + differenceY*differenceY), calculateAngle(B.X, B.Y, E.X, E.Y), calculateAngle(C.X, C.Y, F.X, F.Y)
}

func calculateAngle(x1, y1, x2, y2 float64) float64 {
	v := x2 - x1
	w := y2 - y1
	x := math.Acos(v / math.Sqrt(v*v+(w*w)))
	if w < 0 {
		return -x
	}
	return x
}

func selectFromArray(arr []float64) float64 {
	sort.Float64s(arr)
	u := (len(arr) - 1) * 50 / 100
	v := 1
	if arr[v] != 0.0 {
		w := u - v
		return arr[v] + float64(w)*(arr[v+1]-arr[v])
	}
	return arr[v]
}
