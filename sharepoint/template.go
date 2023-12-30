package sharepoint

import (
	"encoding/xml"
	"log"
)

func FilterTemplate(templateName string) error {

	template := &Provisioning2{}
	byteValue, err := Assets.ReadFile(templateName)
	if err != nil {
		return err
	}
	xml.Unmarshal(byteValue, &template)
	log.Println("Loaded template")
	log.Println(len(template.Templates.ProvisioningTemplate.ClientSidePages.ClientSidePage), "Main Pages")
	countOfTranslations := 0
	for ix, page := range template.Templates.ProvisioningTemplate.ClientSidePages.ClientSidePage {

		log.Println("      ", ix, page.PageName)
		for _, translatedPage := range page.Translations.ClientSidePage {
			log.Println(">", translatedPage.LCID, ix, page.PageName)
			countOfTranslations++
		}

	}
	log.Println(countOfTranslations, "Translations")
	return nil
}

func LoadTemplate(templateName string) error {
	_, err := Assets.ReadFile(templateName)
	if err != nil {
		return err
	}

	return nil
}
