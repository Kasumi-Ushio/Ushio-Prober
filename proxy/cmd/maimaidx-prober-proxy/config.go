package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Kasumi-Ushio/Ushio-Prober/proxy/lib"
)

type workingMode int

const (
	workingModeUpdate workingMode = 0
	workingModeExport workingMode = 1 // only for debug or other
)

type config struct {
	UserName          string   `json:"username"`
	Password          string   `json:"password"`
	Mode              string   `json:"mode,omitempty"`
	MaiDiffs          []string `json:"mai_diffs,omitempty"`
	Verbose           bool     `json:"verbose" default:"false"`
	Addr              string   `json:"addr" default:":8033"`
	NoEditGlobalProxy bool     `json:"no_edit_global_proxy" default:"false"`
	NetworkTimeout    int      `json:"timeout" default:"30"`
	Slice             bool     `json:"slice" default:"false"`
	ProberApiUrl	  string   `json:"prober_api_url" default:"https://www.diving-fish.com/api/maimaidxprober"`
	// intermediate value
	MaiIntDiffs []int
}

/* func (c *config) FlagOverride(set *flag.FlagSet) (err error) {
	set.Visit(func(f *flag.Flag) {
		if f.Name == "v" {
			c.Verbose = f.Value.(flag.Getter).Get().(bool)
		} else if f.Name == "addr" {
			c.Addr = f.Value.(flag.Getter).Get().(string)
		} else if f.Name == "no-edit-global-proxy" {
			c.NoEditGlobalProxy = f.Value.(flag.Getter).Get().(bool)
		} else if f.Name == "timeout" {
			c.NetworkTimeout = f.Value.(flag.Getter).Get().(int)
		} else if f.Name == "mai-diffs" {
			maiDiffs := strings.Split(f.Value.String(), ",")
			if len(maiDiffs) == 1 && maiDiffs[0] == "" {
				maiDiffs = c.MaiDiffs
			} else {
				c.MaiDiffs = maiDiffs
			}
		} else if f.Name == "slice" {
			c.Slice = f.Value.(flag.Getter).Get().(bool)
		}
	})
	return
} */

func (c *config) FlagOverride(set *flag.FlagSet) (err error) {
	set.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "v":
			var verbose bool
			if err := set.Lookup("v").Value.Set(f.Value.String()); err != nil {
				err = err
			} else {
				c.Verbose = verbose
			}
		case "addr":
			c.Addr = f.Value.String()
		case "no-edit-global-proxy":
			var noEditGlobalProxy bool
			if err := set.Lookup("no-edit-global-proxy").Value.Set(f.Value.String()); err != nil {
				err = err
			} else {
				c.NoEditGlobalProxy = noEditGlobalProxy
			}
		case "timeout":
			var timeout int
			if err := set.Lookup("timeout").Value.Set(f.Value.String()); err != nil {
				err = err
			} else {
				c.NetworkTimeout = timeout
			}
		case "mai-diffs":
			maiDiffs := strings.Split(f.Value.String(), ",")
			if len(maiDiffs) == 1 && maiDiffs[0] == "" {
				maiDiffs = c.MaiDiffs
			} else {
				c.MaiDiffs = maiDiffs
			}
		case "slice":
			var slice bool
			if err := set.Lookup("slice").Value.Set(f.Value.String()); err != nil {
				err = err
			} else {
				c.Slice = slice
			}
		case "prober-api-url"
			c.ProberApiUrl = f.Value.String()
	})
	return
}

func (c *config) getWorkingMode() workingMode {
	if c.Mode == "export" {
		return workingModeExport
	}
	return workingModeUpdate
}

func getMaiDiffs(MaiDiffs []string) (diffs []int, err error) {
	maiDiffMap := map[string]int{
		"0":         0,
		"bas":       0,
		"basic":     0,
		"1":         1,
		"adv":       1,
		"advanced":  1,
		"2":         2,
		"exp":       2,
		"expert":    2,
		"3":         3,
		"mas":       3,
		"master":    3,
		"4":         4,
		"rem":       4,
		"remaster":  4,
		"re:master": 4,
	}
	if len(MaiDiffs) == 0 {
		for i := 0; i <= 4; i++ {
			diffs = append(diffs, i) // 添加元素
		}
	} else {
		diffList := []string{"Basic", "Advanced", "Expert", "Master", "Re:MASTER"}
		diffStr := ""
		for _, diff := range MaiDiffs {
			if intDiff, exist := maiDiffMap[strings.ToLower(diff)]; exist {
				diffs = append(diffs, intDiff)
				diffStr += diffList[intDiff] + " "
			} else {
				Log(LogLevelWarning, "未找到 %s 难度等级，已跳过……", diff)
				Log(LogLevelWarning, "Can not found the diff level: %s, skiped...", diff)
			}
		}
		if len(diffs) == 0 {
			Log(LogLevelWarning, "未为舞萌 DX 指定任何难度等级，导入将不会产生任何效果")
			Log(LogLevelWarning, "You did not specific any diff level for maimai DX. Import data will make no effects.")
		} else {
			Log(LogLevelInfo, "您已修改舞萌的难度等级为：%s", diffStr)
			Log(LogLevelInfo, "You have modified the maimai DX diff level to: %s", diffStr)
		}
	}
	return
}

func initConfig(path string) (config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		// First run
		lib.GenerateCert()
		os.WriteFile(path, []byte("{\"username\": \"\", \"password\": \"\"}"), 0644)
		return config{}, fmt.Errorf("初次使用请填写 %s 文件，并依据教程完成根证书的安装。/ Please write the %s file, and import the cert file by following the guide.", path)
	}

	obj := config{
		Addr:              ":8033",
		NetworkTimeout:    30,
		Slice:             false,
		Verbose:           false,
		NoEditGlobalProxy: false,
	}

	err = json.Unmarshal(b, &obj)
	if err != nil {
		return config{}, fmt.Errorf("配置文件格式有误，无法解析：%w。请检查 %s 文件的内容 / Configuration file format was incorrect, we can not recognize: %w, please check the context in %s. ", err, path)
	}

	return obj, nil
}
