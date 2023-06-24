package excel

import (
	"log"

	"github.com/xuri/excelize/v2"
)

func Read(filePath string) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	log.Println("#sheets", f.SheetCount)

}
