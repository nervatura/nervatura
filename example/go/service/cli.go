package service

import (
	"errors"
	"os"
	"os/exec"
	"slices"
	"strings"

	cu "github.com/nervatura/component/pkg/util"
)

type CliClient struct {
	Config cu.SM
}

func (cli *CliClient) connect(arg ...string) (result any, err error) {
	if cli.Config["NT_SERVICE_PATH"] == "docker" {
		arg = append([]string{"exec", "-i", "nervatura", "/nervatura"}, arg...)
	}
	cmd := exec.Command(cli.Config["NT_SERVICE_PATH"], arg...)
	cmd.Env = os.Environ()
	stdout, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	results := strings.Split(string(stdout), "\n")
	infoStr := results[len(results)-2]
	err = cu.ConvertFromByte([]byte(infoStr), &result)
	if err != nil {
		return nil, errors.New("command line error")
	}
	if values, valid := result.(cu.IM); valid {
		if value, found := values["code"]; found {
			if !slices.Contains([]int64{200, 201, 204}, cu.ToInteger(value, 0)) {
				return nil, errors.New(cu.ToString(values["message"], "Unknown error"))
			}
		}
	}
	return result, nil
}

func (cli *CliClient) Database(options cu.IM) (result any, err error) {
	var data []byte
	if data, err = cu.ConvertToByte(options); err != nil {
		return nil, err
	}
	return cli.connect("-c", "database", "-o", string(data))
}

func (cli *CliClient) ResetPassword(options cu.IM) (result any, err error) {
	var data []byte
	if data, err = cu.ConvertToByte(options); err != nil {
		return nil, err
	}
	return cli.connect("-c", "reset", "-o", string(data))
}

func (cli *CliClient) Create(model string, options, data cu.IM) (result any, err error) {
	var optData, dataBytes []byte
	if optData, err = cu.ConvertToByte(options); err != nil {
		return nil, err
	}
	if dataBytes, err = cu.ConvertToByte(data); err != nil {
		return nil, err
	}
	return cli.connect("-c", "create", "-m", model, "-o", string(optData), "-d", string(dataBytes))
}

func (cli *CliClient) Update(model string, options, data cu.IM) (result any, err error) {
	var optData, dataBytes []byte
	if optData, err = cu.ConvertToByte(options); err != nil {
		return nil, err
	}
	if dataBytes, err = cu.ConvertToByte(data); err != nil {
		return nil, err
	}
	return cli.connect("-c", "update", "-m", model, "-o", string(optData), "-d", string(dataBytes))
}

func (cli *CliClient) Delete(model string, options cu.IM) (result any, err error) {
	var optData []byte
	if optData, err = cu.ConvertToByte(options); err != nil {
		return nil, err
	}
	return cli.connect("-c", "delete", "-m", model, "-o", string(optData))
}

func (cli *CliClient) Get(model string, options cu.IM) (result any, err error) {
	var optData []byte
	if optData, err = cu.ConvertToByte(options); err != nil {
		return nil, err
	}
	return cli.connect("-c", "get", "-m", model, "-o", string(optData))
}

func (cli *CliClient) Query(model string, options cu.IM) (result any, err error) {
	var optData []byte
	if optData, err = cu.ConvertToByte(options); err != nil {
		return nil, err
	}
	return cli.connect("-c", "query", "-m", model, "-o", string(optData))
}

func (cli *CliClient) View(options cu.IM) (result any, err error) {
	var optData []byte
	if optData, err = cu.ConvertToByte(options); err != nil {
		return nil, err
	}
	return cli.connect("-c", "view", "-o", string(optData))
}

func (cli *CliClient) Function(options cu.IM) (result any, err error) {
	var optData []byte
	if optData, err = cu.ConvertToByte(options); err != nil {
		return nil, err
	}
	return cli.connect("-c", "function", "-o", string(optData))
}
