package disk

import (
	"github.com/ricochet2200/go-disk-usage/du"
	"log"
	"os"
)

func GetDisk() *du.DiskUsage {
	pwd, err := os.Executable()
	if err != nil {
		log.Println(err)
		return nil
	}
	usage := du.NewDiskUsage(pwd)
	return usage
}
