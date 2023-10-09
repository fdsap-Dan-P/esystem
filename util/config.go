package util

import (
	"bytes"
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	DBHost              string        `mapstructure:"DB_HOST"`
	DBPort              string        `mapstructure:"DB_PORT"`
	DBUser              string        `mapstructure:"DB_USER"`
	DBPass              string        `mapstructure:"DB_PASS"`
	DBName              string        `mapstructure:"DB_NAME"`
	DockerSQLImgID      string        `mapstructure:"DOCKER_SQL_IMGID"`
	ESystemDump         string        `mapstructure:"ESYSTEMDUMP"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	DBeSystemLocal      string        `mapstructure:"DB_eSystemLocal"`
	DBDataLake          string        `mapstructure:"DB_DataLake"`
	HomeFolder          string        `mapstructure:"HOME_FOLDER"`
	DB2T24CB            string        `mapstructure:"DB2_T24_CB"`
	DB2T24RBI           string        `mapstructure:"DB2_T24_RBI"`
	DB2T24SME           string        `mapstructure:"DB2_T24_SME"`
	DB2DWHCB            string        `mapstructure:"DB2_DWH_CB"`
	DB2DWHRBI           string        `mapstructure:"DB2_DWH_RBI"`
	DB2DWHSME           string        `mapstructure:"DB2_DWH_SME"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string, key []byte) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	encryptedData := []byte(viper.GetString("encrypted_config"))
	// Decrypt the encrypted configuration data
	decryptedData, err := DecryptData(encryptedData, key)
	if err != nil {
		return
	}

	// Unmarshal the decrypted data into the Config struct
	err = viper.ReadConfig(bytes.NewBuffer(decryptedData))
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
