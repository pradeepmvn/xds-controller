package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// LoadConfig initializes and parses yaml fileinto struct
func LoadConfig() *ControllerConfig {
	cc := &ControllerConfig{}
	yamlFile, err := ioutil.ReadFile("/config/config.yaml")
	if err != nil {
		log.Fatal("Failed to open Configuration file at config/config.yaml")
	}
	err = yaml.Unmarshal(yamlFile, cc)
	if err != nil {
		log.Panic("Unable to Load Config", err)
	}

	cj, _ := json.Marshal(cc)
	log.Printf("Resolved Config: %s", cj)
	return cc
}

// ControllerConfig cretes a list of services that are configured
type ControllerConfig struct {
	NodeId               string     `yaml:"node_id"`
	LogDebug             bool       `yaml:"log_debug"`
	MaxConcurrentStreams int        `yaml:"max_concurrent_streams"`
	ListenerPort         int        `yaml:"listener_port"`
	PrometheusPort       int        `yaml:"prometheus_port"`
	Clusters             []*Cluster `yaml:"clusters"`
}

// Cluster a specific clluster config
type Cluster struct {
	Name                 string `yaml:"name"`
	ResolverType         string `yaml:"resolver_type"`
	RefreshIntervalInSec int    `yaml:"refresh_interval_in_sec"`
	Address              string `yaml:"address"`
	NameSpace            string `yaml:"name_space"`
	Port                 uint32 `yaml:"port"`
	LbPolicy             string `yaml:"lb_policy"`
	MaxRequests          uint32 `yaml:"max_requests"`
	Retry                Retry  `yaml:"retry"`
}

// Retry Configuration specific to a cluster
type Retry struct {
	Enabled                 bool   `yaml:"enabled"`
	RetryOn                 string `yaml:"retry_on"`
	NumRetries              uint32 `yaml:"num_retries"`
	BackoffIntervalInSec    uint32 `yaml:"backoff_interval_in_sec"`
	BackoffMaxIntervalInSec uint32 `yaml:"backoff_max_interval_in_sec"`
}
