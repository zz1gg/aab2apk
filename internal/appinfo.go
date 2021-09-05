package internal

import (
	"aab2apk/pkg"

	"github.com/fatih/color"

	"github.com/shogo82148/androidbinary/apk"

	"github.com/jedib0t/go-pretty/v6/table"

	"os"
)

type APPInfomation struct {
	appname      string
	description  string
	packageName  string
	mainActivity string
	versioncode  int32
	maxSDK       int32
	minSDK       int32
	versionname  string
	debuggable   bool
	allowBackup  bool
}

func APPInfo(outputname, target string) {

	//aab size
	aabsize, _ := os.Stat(target)
	apksize, _ := os.Stat(outputname)
	appinfo := AndroidBinInfo(outputname)
	//log.Println("Extract APP Information: ")
	color.Blue("APP Information:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Value"})
	t.AppendRows([]table.Row{
		{1, "AAB Name", target},
	})

	t.AppendRow([]interface{}{2, "APP Name", appinfo.appname})
	t.AppendRow([]interface{}{3, "Package Name", appinfo.packageName})
	t.AppendRow([]interface{}{4, "AAB Size(byte)", aabsize.Size()})
	t.AppendRow([]interface{}{5, "APP VersionName", appinfo.versionname})
	t.AppendRow([]interface{}{6, "AAB MD5", pkg.GetMD5Hash(target)})
	t.AppendRow([]interface{}{7, "MainActivity", appinfo.mainActivity})
	t.AppendRow([]interface{}{8, "Debuggable", appinfo.debuggable})
	t.AppendRow([]interface{}{9, "AllowBackup", appinfo.allowBackup})
	t.AppendRow([]interface{}{10, "VersionCode", appinfo.versioncode})
	t.AppendRow([]interface{}{11, "minSDK", appinfo.minSDK})
	t.AppendRow([]interface{}{12, "APK MD5", pkg.GetMD5Hash(outputname)})
	t.AppendRow([]interface{}{13, "APK Size(byte)", apksize.Size()})


	t.SetStyle(table.StyleColoredBlackOnCyanWhite)  //output style, it can be commented out
	t.Render()

}

func AndroidBinInfo(outputname string) APPInfomation {

	var appinfomation APPInfomation

	pkg, _ := apk.OpenFile(outputname)

	defer pkg.Close()
	pkg.Manifest().App.Description.MustString()
	appinfomation.appname = pkg.Manifest().App.Name.MustString()
	appinfomation.versionname = pkg.Manifest().VersionName.MustString()
	appinfomation.packageName = pkg.PackageName()
	appinfomation.mainActivity, _ = pkg.MainActivity()
	appinfomation.versioncode, _ = pkg.Manifest().VersionCode.Int32()
	appinfomation.maxSDK, _ = pkg.Manifest().SDK.Max.Int32()
	appinfomation.minSDK, _ = pkg.Manifest().SDK.Min.Int32()
	appinfomation.allowBackup = pkg.Manifest().App.AllowBackup.MustBool()
	appinfomation.debuggable = pkg.Manifest().App.Debuggable.MustBool()
	return appinfomation
}
