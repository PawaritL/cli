package configure

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/databricks/bricks/cmd/root"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

type Configs struct {
	Host    string `ini:"host"`
	Token   string `ini:"token,omitempty"`
	Profile string `ini:"-"`
}

var tokenMode bool

func (cfg *Configs) loadNonInteractive(cmd *cobra.Command) error {
	host, err := cmd.Flags().GetString("host")
	if err != nil || host == "" {
		return fmt.Errorf("use --host to specify host in non interactive mode: %w", err)
	}
	cfg.Host = host

	if !tokenMode {
		return nil
	}

	n, err := fmt.Scanf("%s\n", &cfg.Token)
	if err != nil {
		return err
	}
	if n != 1 {
		return fmt.Errorf("exactly 1 argument required")
	}
	return nil
}

var configureCmd = &cobra.Command{
	Use:    "configure",
	Short:  "Configure authentication",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		profile, err := cmd.Flags().GetString("profile")
		if err != nil {
			return fmt.Errorf("read --profile flag: %w", err)
		}

		path := os.Getenv("DATABRICKS_CONFIG_FILE")
		if path == "" {
			path, err = os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("homedir: %w", err)
			}
		}
		if filepath.Base(path) == ".databrickscfg" {
			path = filepath.Dir(path)
		}
		err = os.MkdirAll(path, os.ModeDir|os.ModePerm)
		if err != nil {
			return fmt.Errorf("create config dir: %w", err)
		}
		cfgPath := filepath.Join(path, ".databrickscfg")
		_, err = os.Stat(cfgPath)
		if errors.Is(err, os.ErrNotExist) {
			file, err := os.Create(cfgPath)
			if err != nil {
				return fmt.Errorf("create config file: %w", err)
			}
			file.Close()
		} else if err != nil {
			return fmt.Errorf("open config file: %w", err)
		}

		ini_cfg, err := ini.Load(cfgPath)
		if err != nil {
			return fmt.Errorf("load config file: %w", err)
		}
		cfg := &Configs{"", "", profile}
		err = ini_cfg.Section(profile).MapTo(cfg)
		if err != nil {
			return fmt.Errorf("unmarshal loaded config: %w", err)
		}

		err = cfg.loadNonInteractive(cmd)
		if err != nil {
			return fmt.Errorf("reading configs: %w", err)
		}

		err = ini_cfg.Section(profile).ReflectFrom(cfg)
		if err != nil {
			return fmt.Errorf("marshall config: %w", err)
		}

		var buffer bytes.Buffer
		if ini_cfg.Section("DEFAULT").Body() != "" {
			//This configuration makes the ini library write the DEFAULT header explicitly.
			//DEFAULT section might be empty
			ini.DefaultHeader = true
		}
		_, err = ini_cfg.WriteTo(&buffer)
		if err != nil {
			return fmt.Errorf("write config to buffer: %w", err)
		}
		err = os.WriteFile(cfgPath, buffer.Bytes(), os.ModePerm)
		if err != nil {
			return fmt.Errorf("write congfig to file: %w", err)
		}

		return nil
	},
}

func init() {
	root.RootCmd.AddCommand(configureCmd)
	configureCmd.Flags().BoolVarP(&tokenMode, "token", "t", false, "Configure using Databricks Personal Access Token")
	configureCmd.Flags().String("host", "", "Host to connect to.")
	configureCmd.Flags().String("profile", "DEFAULT", "CLI connection profile to use.")
}
