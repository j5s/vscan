package poc

import (
	"fmt"
	"github.com/veo/vscan/brute"
	"github.com/veo/vscan/poc/ThinkPHP"
	"github.com/veo/vscan/poc/fastjson"
	"github.com/veo/vscan/poc/jboss"
	"github.com/veo/vscan/poc/jenkins"
	"github.com/veo/vscan/poc/phpunit"
	"github.com/veo/vscan/poc/seeyon"
	"github.com/veo/vscan/poc/shiro"
	"github.com/veo/vscan/poc/tomcat"
	"github.com/veo/vscan/poc/weblogic"
	"net/url"
)

func POCcheck(wappalyzertechnologies []string, URL string, finalURL string) []string {
	var HOST string
	var technologies []string
	if host, err := url.Parse(URL); err == nil {
		HOST = host.Host
	}
	for tech := range wappalyzertechnologies {
		switch wappalyzertechnologies[tech] {
		case "Shiro":
			key := shiro.CVE_2016_4437(finalURL)
			if key != "" {
				technologies = append(technologies, fmt.Sprintf("exp-Shiro|key:%s", key))
			}
		case "Apache Tomcat":
			username, password := brute.Tomcat_brute(URL)
			if username != "" {
				technologies = append(technologies, fmt.Sprintf("brute-Tomcat|%s:%s", username, password))
			}
			if tomcat.CVE_2020_1938(HOST) {
				technologies = append(technologies, "exp-Tomcat|CVE_2020_1938")
			}
			if tomcat.CVE_2017_12615(URL) {
				technologies = append(technologies, "exp-Tomcat|CVE_2017_12615")
			}
		case "Basic":
			username, password := brute.Basic_brute(URL)
			if username != "" {
				technologies = append(technologies, fmt.Sprintf("brute-basic|%s:%s", username, password))
			}
		case "WebLogic":
			username, password := brute.Weblogic_brute(URL)
			if username != "" {
				if username == "login_page" {
					technologies = append(technologies, "WebLogic_login_page")
				} else {
					technologies = append(technologies, fmt.Sprintf("brute-WebLogic|%s:%s", username, password))
				}
			}
			if weblogic.CVE_2014_4210(URL) {
				technologies = append(technologies, "exp-WebLogic|CVE_2014_4210")
			}
			if weblogic.CVE_2017_3506(URL) {
				technologies = append(technologies, "exp-WebLogic|CVE_2017_3506")
			}
			if weblogic.CVE_2017_10271(URL) {
				technologies = append(technologies, "exp-WebLogic|CVE_2017_10271")
			}
			if weblogic.CVE_2018_2894(URL) {
				technologies = append(technologies, "exp-WebLogic|CVE_2018_2894")
			}
			if weblogic.CVE_2019_2725(URL) {
				technologies = append(technologies, "exp-WebLogic|CVE_2019_2725")
			}
			if weblogic.CVE_2019_2729(URL) {
				technologies = append(technologies, "exp-WebLogic|CVE_2019_2729")
			}
			if weblogic.CVE_2020_2883(URL) {
				technologies = append(technologies, "exp-WebLogic|CVE_2020_2883")
			}
			if weblogic.CVE_2020_14882(URL) {
				technologies = append(technologies, "exp-WebLogic|CVE_2020_14882")
			}
			if weblogic.CVE_2020_14883(URL) {
				technologies = append(technologies, "exp-WebLogic|CVE_2020_14883")
			}
			if weblogic.CVE_2021_2109(URL) {
				technologies = append(technologies, "exp-WebLogic|CVE_2021_2109")
			}
		case "JBoss Application Server":
			if jboss.CVE_2017_12149(URL) {
				technologies = append(technologies, "exp-jboss|CVE_2017_12149")
			}
			username, password := brute.Jboss_brute(URL)
			if username != "" {
				technologies = append(technologies, fmt.Sprintf("brute-jboss|%s:%s", username, password))
			}
		case "JSON":
			fastjsonversion := fastjson.Check(URL)
			if fastjsonversion != "" {
				technologies = append(technologies, fmt.Sprintf("fastjson|%s", fastjsonversion))
			}
		case "Jenkins":
			if jenkins.Unauthorized(URL) {
				technologies = append(technologies, "exp-jenkins|Unauthorized script")
			}
			if jenkins.CVE_2018_1000110(URL) {
				technologies = append(technologies, "exp-jenkins|CVE_2018_1000110")
			}
			if jenkins.CVE_2018_1000861(URL) {
				technologies = append(technologies, "exp-jenkins|CVE_2018_1000861")
			}
			if jenkins.CVE_2019_10003000(URL) {
				technologies = append(technologies, "exp-jenkins|CVE_2019_10003000")
			}
		case "ThinkPHP":
			if ThinkPHP.RCE(URL) {
				technologies = append(technologies, "exp-ThinkPHP")
			}
		case "phpunit":
			if phpunit.CVE_2017_9841(URL) {
				technologies = append(technologies, "exp-phpunit|CVE_2017_9841")
			}
		case "致远OA":
			if seeyon.SeeyonFastjson(URL) {
				technologies = append(technologies, "exp-seeyon|SeeyonFastjson")
			}
			if seeyon.SessionUpload(URL) {
				technologies = append(technologies, "exp-seeyon|SessionUpload")
			}
			if seeyon.CNVD_2019_19299(URL) {
				technologies = append(technologies, "exp-seeyon|CNVD_2019_19299")
			}
			if seeyon.CNVD_2020_62422(URL) {
				technologies = append(technologies, "exp-seeyon|CNVD_2020_62422")
			}
			if seeyon.CNVD_2021_01627(URL) {
				technologies = append(technologies, "exp-seeyon|CNVD_2021_01627")
			}
			if seeyon.CreateMysql(URL) {
				technologies = append(technologies, "exp-seeyon|CreateMysql")
			}
			if seeyon.DownExcelBeanServlet(URL) {
				technologies = append(technologies, "exp-seeyon|DownExcelBeanServlet")
			}
			if seeyon.GetSessionList(URL) {
				technologies = append(technologies, "exp-seeyon|GetSessionList")
			}
			if seeyon.InitDataAssess(URL) {
				technologies = append(technologies, "exp-seeyon|InitDataAssess")
			}
			if seeyon.ManagementStatus(URL) {
				technologies = append(technologies, "exp-seeyon|ManagementStatus")
			}
			if seeyon.BackdoorScan(URL) {
				technologies = append(technologies, "exp-seeyon|Backdoor")
			}
		case "LoginPage":
			username, password, loginurl := brute.Admin_brute(URL)
			if loginurl != "" {
				technologies = append(technologies, fmt.Sprintf("brute-admin|%s:%s", username, password))
			}
		}
	}
	return technologies
}
