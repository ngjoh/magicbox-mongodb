package restapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/koksmat-com/koksmat/model"
	"github.com/koksmat-com/koksmat/powershell"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

const sharePointTag = "SharePoint"

func getBlobb(path string) (*string, error) {
	return nil, nil
}

func getBlob() usecase.Interactor {
	type BlobRequest struct {
		Tag string `json:"tag" path:"tag" example:"SITEMAP%7Chttps%3A%2F%2Fchristianiabpos.sharepoint.com%2Fsites%2Fnexiintra-home"  binding:"required"`
	}
	type BlobResponse struct {
		Content map[string]interface{} `json:"content,inline"`
		Cache   string                 `header:"Cache-Control" json:"-"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, input BlobRequest, output *BlobResponse) error {

		o, err := model.GetBlob(input.Tag)
		if err != nil {
			return err
		}
		br := &BlobResponse{
			Content: *o,
			Cache:   "public, max-age=60",
		}
		*output = *br
		return err

	})

	u.SetTitle("Reading blob")
	u.SetDescription(`


Returns a piece of unstructured data 

`)
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(authenticationTag)
	return u
}

func copyLibrary() usecase.Interactor {
	type CopyLibraryRequest struct {
		FromUrl    string `json:"fromurl"  example:"https://christianiabpos.sharepoint.com/sites/nexi"  binding:"required"`
		ToUrl      string `json:"tourl" example:"https://christianiabpos.sharepoint.com/sites/nexiintra-home"  binding:"required"`
		FromLibray string `json:"fromlibrary"  example:"Shared Documents"  binding:"required"`
		ToLibrary  string `json:"tolibrary" example:"Copy of Shared Documents"  binding:"required"`
	}
	type CopyLibraryResponse struct {
	}

	u := usecase.NewInteractor(func(ctx context.Context, input CopyLibraryRequest, output *CopyLibraryResponse) error {

		_, err := powershell.CopyLibrary(input.FromUrl, input.ToUrl, input.FromLibray, input.ToLibrary)

		return err

	})

	u.SetTitle("Copy a library ")
	u.SetDescription(`
	Copy a library from one site to another site, can also copy internally in the same site

	Future: Copy a library from one site to another site, cross tenancy

`)
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(sharePointTag)
	return u
}

func copyPage() usecase.Interactor {
	type CopyPageRequest struct {
		FromPageUrl string `json:"frompageurl"  example:"https://christianiabpos.sharepoint.com/sites/nexi/SitePages/Home.aspx?mode=Edit"  binding:"required"`
		ToSiteUrl   string `json:"tositeurl"  example:"https://christianiabpos.sharepoint.com/sites/nexiintra-home/SitePages/Home.aspx"  binding:"required"`
	}
	type CopyPageResponse struct {
		NewPageUrl string `json:"newpageurl"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, input CopyPageRequest, output *CopyPageResponse) error {
		split := strings.Split(input.FromPageUrl, "/SitePages/")
		fromUrl := fmt.Sprintf("%s", split[0])

		pageNameRaw1 := fmt.Sprintf("%s", split[1])
		pageName := fmt.Sprintf("%s", strings.Split(pageNameRaw1, "?")[0])
		destPageName := pageName
		if len(strings.Split(pageName, "/")) > 1 {
			destPageName = strings.Split(pageName, "/")[1]
		}

		result, err := powershell.CopyPage(fromUrl, input.ToSiteUrl, pageName, destPageName)

		if (err == nil) && (result != nil) {
			output.NewPageUrl = result.NewPageURL
		}
		return err

	})

	u.SetTitle("Copy a page ")
	u.SetDescription(`
	

`)
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(sharePointTag)
	return u
}

func renameLibrary() usecase.Interactor {
	type CopyLibraryRequest struct {
		SiteUrl        string `json:"siteurl"  example:"https://christianiabpos.sharepoint.com/sites/nexiintra"  binding:"required"`
		OldLibraryName string `json:"oldlibraryname" example:"Import1"  binding:"required"`
		NewLibraryName string `json:"newlibraryname"  example:"Regulatory Documents"  binding:"required"`
		NewLibraryURL  string `json:"newurl" example:"regulatory_documents"  binding:"required"`
	}
	type CopyLibraryResponse struct {
	}

	u := usecase.NewInteractor(func(ctx context.Context, input CopyLibraryRequest, output *CopyLibraryResponse) error {

		_, err := powershell.RenameLibrary(input.SiteUrl, input.OldLibraryName, input.NewLibraryName, input.NewLibraryURL)

		return err

	})

	u.SetTitle("Rename a Library or List ")
	u.SetDescription(`
	Rename a library title and URL

`)
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(sharePointTag)
	return u
}
