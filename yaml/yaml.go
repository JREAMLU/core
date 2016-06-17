package yaml

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var Yconf Conf

type Conf struct {
	Platform struct {
		//command
		Command_first  string `yaml:"command_first"`
		Command_second string `yaml:"command_second"`
		Command_third  string `yaml:"command_third"`
		Command_kill   string `yaml:"command_kill"`
		Command_pgrep  string `yaml:"command_pgrep"`

		//kafka
		Kafka_topic                        string `yaml:"kafka_topic"`
		Kafka_addrs                        string `yaml:"kafka_addrs"`
		Kafka_chroot                       string `yaml:"kafka_chroot"`
		Kafka_zkaddrs                      string `yaml:"kafka_zkaddrs"`
		Kafka_group_name                   string `yaml:"kafka_group_name"`
		Offsets_commit_interval            int    `yaml:"offsets_commit_interval"`
		Offsets_processing_timeout_seconds int    `yaml:"offsets_processing_timeout_seconds"`

		//database master
		Db_alias_m       string `yaml:"db_alias_m"`
		Db_driver_m      string `yaml:"db_driver_m"`
		Db_username_m    string `yaml:"db_username_m"`
		Db_password_m    string `yaml:"db_password_m"`
		Db_server_m      string `yaml:"db_server_m"`
		Db_port_m        string `yaml:"db_port_m"`
		Db_name_m        string `yaml:"db_name_m"`
		Db_charset_m     string `yaml:"db_charset_m"`
		Db_maxidle_m     int    `yaml:"db_maxidle_m"`
		Db_debug         bool   `yaml:"db_debug"`
		Db_singulartable bool   `yaml:"db_singulartable"`
		Db_logmode       bool   `yaml:"db_logmode"`
		Db_setting       string `yaml:"db_setting"`

		//process status
		Processing string `yaml:"processing"`
		Stop       string `yaml:"stop"`

		//定时检查进程时间
		Check_time int `yaml:"check_time"`

		//log path
		Log_path string `yaml:"log_path"`

		//mongodb
		Mgo_db_name string `yaml:"mgo_db_name"`
		Mongodb     string `yaml:"mongodb"`

		//url
		Base_url string `yaml:"base_url"`

		//alarm
		Alarm_mobile string `yaml:"alarm_mobile"`

		//rpc_port
		Rpc_port string `yaml:"rpc_port"`
	}
}

func LoadConf(path string) {
	data := read(path)
	err := yaml.Unmarshal([]byte(data), &Yconf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func read(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}
