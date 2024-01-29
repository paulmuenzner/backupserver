package email

import (
	"fmt"

	"github.com/paulmuenzner/golang-backupserver/config"
	convert "github.com/paulmuenzner/golang-backupserver/utils/convert"
	envHandler "github.com/paulmuenzner/golang-backupserver/utils/env"
)

// Retrieve configuration data (eg. email provider, smtp port) from .env file for production settings only
// Base parameter for dependency injection of email client (production)
func EmailProductionConfig() (emailClientConfig *EmailClientConfigData, err error) {
	// Retrieve .env values by keys provided in config file

	// HOST
	host, err := envHandler.GetEnvValue(config.EmailProviderHostEnv, "")
	if err != nil {
		return nil, fmt.Errorf("Cannot retrieve .env value for host of email provider in 'EmailProductionConfig()'. Env key: %s. No default value has been employed. Error: %v", config.EmailProviderHostEnv, err)
	}

	// USERNAME
	smtpUsername, err := envHandler.GetEnvValue(config.EmailProviderUserNameEnv, "")
	if err != nil {
		return nil, fmt.Errorf("Cannot retrieve .env value for smtp user name in 'EmailProductionConfig()'. Env key: %s. No default value has been employed. Error: %v", config.EmailProviderUserNameEnv, err)
	}

	// PASSWORD
	smtpPassword, err := envHandler.GetEnvValue(config.EmailProviderPasswordEnv, "")
	if err != nil {
		return nil, fmt.Errorf("Cannot retrieve .env value for smtp password in 'EmailProductionConfig()'. Env key: %s. No default value has been employed. Error: %v", config.EmailProviderPasswordEnv, err)
	}

	// SMTP PORT
	smtpPortString, err := envHandler.GetEnvValue(config.EmailProviderSmtpPortEnv, "")
	if err != nil {
		return nil, fmt.Errorf("Cannot retrieve .env value for smtp port in 'EmailProductionConfig()'. Env key: %s. No default value has been employed. Error: %v", config.EmailProviderSmtpPortEnv, err)
	}
	smtPortInt, err := convert.ConvertStringToInt(smtpPortString)
	if err != nil {
		return nil, fmt.Errorf("Cannot convert smtpPortString to int in 'EmailProductionConfig()'. smtpPortString: %s. No default value has been employed. Error: %v", smtpPortString, err)
	}

	// Configure EmailClientConfigData structure
	emailClientConfig = &EmailClientConfigData{Host: host,
		SmtpUsername: smtpUsername,
		SmtpPassword: smtpPassword,
		SmtpPort:     smtPortInt}

	return emailClientConfig, nil
}
