package setup

//Set переменная для хранения текущих настроек
var SetServer *Server
var SetClient *Client

//SetupServer общая структура для настройки системы сервера
type Server struct {
	LogPath string   `toml:"log"`
	Home    DataBase `toml:"home"`
	Regions []Region `toml:"regions"` //Подключенные регионы
}

//SetupClient общая структура для настройки системы клиента
type Client struct {
	LogPath     string   `toml:"log"`
	Home        DataBase `toml:"home"`
	ID          int      `toml:"id"`
	Tables      []Table  `toml:"tables"`
	PathCrosses string   `toml:"crosses"`
}

//Region описание одного региона
type Region struct {
	Name        string `toml:"name"`
	IDIn        int    `toml:"idIn"`
	IDOut       int    `toml:"idOut"`
	Description string `toml:"description"`
}

//Table описание передаваемых таблиц
type Table struct {
	Name   string  `toml:"name"`
	Fields []Field `toml:"fields"`
}

//Field описание колонок таблицы для передачи
type Field struct {
	Name string `toml:"name"`
	Type string `toml:"type"`
}

//DataBase настройки базы данных postresql
type DataBase struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
}
