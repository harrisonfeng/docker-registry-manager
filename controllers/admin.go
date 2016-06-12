package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/stefannaglee/docker-registry-manager/utilities"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {

	c.Data["releaseVersion"] = utils.ReleaseVersion

	c.TplName = "admin.tpl"
}

func (c *AdminController) GetLogs() {
	logs := utils.ParseLogFile()

	c.Data["json"] = &logs
	c.ServeJSON()
}

func (c *AdminController) ClearLogs() {

	err := utils.ClearLogFile()
	if err == nil {
		c.CustomAbort(200, "Success")
	}
	c.CustomAbort(404, "Failed to clear log: "+err.Error())

}

func (c *AdminController) ArchiveLogs() {

	err := utils.ArchiveLogFile()
	if err == nil {
		c.CustomAbort(200, "Success")
	}
	c.CustomAbort(404, "Failed to clear log: "+err.Error())

}

func (c *AdminController) ToggleDebug() {

	utils.ToggleDebug()
	c.CustomAbort(200, "Success")
}

func (c *AdminController) GetLogLevel() {

	currentLevel := utils.Log.Level
	type level struct {
		LogLevel string
	}
	l := level{
		LogLevel: currentLevel.String(),
	}
	c.Data["json"] = &l
	c.ServeJSON()

}

func (c *AdminController) GetReleaseVersion() {

	currentRelease := utils.ReleaseVersion
	type ReleaseVersion struct {
		ReleaseVersion string
	}
	r := ReleaseVersion{
		ReleaseVersion: currentRelease,
	}
	c.Data["json"] = &r
	c.ServeJSON()

}

// GetLiveStatistics returns stats on request information tracked by beego
func (c *AdminController) GetLiveStatistics() {

	r := toolbox.StatisticsMap
	rs := r.GetMapData()

	// Convert beego times to seconds for sorted
	for _, req := range rs {
		var err error
		req["min_s"], err = utils.StatToSeconds(req["min_time"].(string))
		if err != nil {
			utils.Log.Error(err)
		}
		req["max_s"], err = utils.StatToSeconds(req["max_time"].(string))
		if err != nil {
			utils.Log.Error(err)
		}
		req["avg_s"], err = utils.StatToSeconds(req["avg_time"].(string))
		if err != nil {
			utils.Log.Error(err)
		}
		req["total_s"], err = utils.StatToSeconds(req["total_time"].(string))
		if err != nil {
			utils.Log.Error(err)
		}
	}

	fmt.Println(rs)

	c.Data["json"] = &rs
	c.ServeJSON()

}
