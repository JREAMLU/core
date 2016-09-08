package useragent

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var ConfigCache = NewCache()

var (
	app            []byte
	browser        []byte
	engine         []byte
	library        []byte
	pim            []byte
	player         []byte
	reader         []byte
	camera         []byte
	carBrowser     []byte
	console        []byte
	media          []byte
	mobile         []byte
	tv             []byte
	oss            []byte
	robot          []byte
	vendorfragment []byte
)

type Cache struct {
	Robots       []Robot
	OSs          []OS
	Vendors      []Vendor
	Apps         []Client
	Browsers     []Browser
	Engines      []Engine
	Libraries    []Client
	Pims         []Client
	MediaPlayers []Client
	Readers      []Reader
	Cameras      []Device
	CarBrowsers  []Device
	Consoles     []Device
	Medias       []Device
	Mobiles      []Device
	TVs          []Device
}

func NewCache() *Cache {
	app, _ = loadFile("./conf/client/app.json")
	browser, _ = loadFile("./conf/client/browser.json")
	engine, _ = loadFile("./conf/client/engine.json")
	library, _ = loadFile("./conf/client/library.json")
	pim, _ = loadFile("./conf/client/pim.json")
	player, _ = loadFile("./conf/client/player.json")
	reader, _ = loadFile("./conf/client/reader.json")

	camera, _ = loadFile("./conf/device/camera.json")
	carBrowser, _ = loadFile("./conf/device/car_browser.json")
	console, _ = loadFile("./conf/device/console.json")
	media, _ = loadFile("./conf/device/media.json")
	mobile, _ = loadFile("./conf/device/mobile.json")
	tv, _ = loadFile("./conf/device/tv.json")

	oss, _ = loadFile("./conf/system/os.json")
	robot, _ = loadFile("./conf/system/robot.json")
	vendorfragment, _ = loadFile("./conf/system/vendorfragment.json")

	cache := &Cache{}
	cache.loadRobots()
	cache.loadOSs()
	cache.loadVendor()
	cache.loadApps()
	cache.loadBrowser()
	cache.loadEngines()
	cache.loadLibraries()
	cache.loadPims()
	cache.loadMeidaPlayers()
	cache.loadReaders()
	cache.loadCameras()
	cache.loadCarBrowser()
	cache.loadConsole()
	cache.loadMedia()
	cache.loadMobile()
	cache.loadTV()
	return cache
}

func loadFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (this *Cache) loadRobots() {
	err := json.Unmarshal(robot, &this.Robots)
	checkErr(err)
}

func (this *Cache) loadOSs() {
	err := json.Unmarshal(oss, &this.OSs)
	checkErr(err)
}

func (this *Cache) loadVendor() {
	temp := make(map[string][]string)
	err := json.Unmarshal(vendorfragment, &temp)
	checkErr(err)
	for k, v := range temp {
		ven := Vendor{}
		ven.Producer = k
		for _, value := range v {
			ven.Regex = append(ven.Regex, value)
		}
		this.Vendors = append(this.Vendors, ven)
	}
}

func (this *Cache) loadApps() {
	err := json.Unmarshal(app, &this.Apps)
	checkErr(err)
}

func (this *Cache) loadBrowser() {
	err := json.Unmarshal(browser, &this.Browsers)
	checkErr(err)
}

func (this *Cache) loadEngines() {
	err := json.Unmarshal(engine, &this.Engines)
	checkErr(err)
}

func (this *Cache) loadLibraries() {
	err := json.Unmarshal(library, &this.Libraries)
	checkErr(err)
}

func (this *Cache) loadPims() {
	err := json.Unmarshal(pim, &this.Pims)
	checkErr(err)
}

func (this *Cache) loadMeidaPlayers() {
	err := json.Unmarshal(player, &this.MediaPlayers)
	checkErr(err)
}

func (this *Cache) loadReaders() {
	err := json.Unmarshal(reader, &this.Readers)
	checkErr(err)
}

func (this *Cache) loadCameras() {
	err := json.Unmarshal(camera, &this.Cameras)
	checkErr(err)
}

func (this *Cache) loadCarBrowser() {
	err := json.Unmarshal(carBrowser, &this.CarBrowsers)
	checkErr(err)
}

func (this *Cache) loadConsole() {
	err := json.Unmarshal(console, &this.Consoles)
	checkErr(err)
}

func (this *Cache) loadMedia() {
	err := json.Unmarshal(media, &this.Medias)
	checkErr(err)
}

func (this *Cache) loadMobile() {
	err := json.Unmarshal(mobile, &this.Mobiles)
	checkErr(err)
}

func (this *Cache) loadTV() {
	err := json.Unmarshal(tv, &this.TVs)
	checkErr(err)
}
