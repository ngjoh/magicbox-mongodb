package model

import (
	//"log"

	"fmt"
	"testing"

	"github.com/koksmat-com/koksmat/io"
	"github.com/stretchr/testify/assert"
)

func TestSyncHubSiteNavigation(t *testing.T) {

	t.Log("Syncing Site Navigation")
	err := SyncSitesNavigation()
	if err != nil {
		t.Error(err)
	}
	t.Log("Done")
}

func TestSyncHubSitePages(t *testing.T) {
	t.Log("Syncing Site Pages")

	err := SyncHubSitePages("b80f09f2-c5e5-4f69-9944-33e8fe18a96c")
	if err != nil {
		t.Error(err)
	}

	t.Log("Done")
}

func TestSyncHubSitePagesProducts(t *testing.T) {
	t.Log("Syncing Site Pages")

	err := SyncHubSitePages("2dbf44b0-6ac7-4d46-a14f-397cbe8ab728")
	if err != nil {
		t.Error(err)
	}

	t.Log("Done")
}

func TestSyncHubSites(t *testing.T) {
	t.Log("Syncing Hub Sites")

	err := SyncHubSites()
	if err != nil {
		t.Error(err)
	}

	t.Log("Done")
}

func TestSharePointHubSitemap(t *testing.T) {
	err := SharePointHubSitemap("https://christianiabpos.sharepoint.com/sites/nexiintra-home")
	if err != nil {
		t.Error(err)
	}
}

func TestSharePointSitemapIntranet(t *testing.T) {
	url := "https://christianiabpos.sharepoint.com/sites/nexiintra-home"
	siteMap, err := SharePointSitemap(url)

	if err != nil {
		t.Error(err)
	}
	//io.WriteFile("intranet.sitemap.json", siteMap)
	assert.Greater(t, len(*&siteMap.Menu), 0)
	tag := fmt.Sprintf("SITEMAP|%s", url)
	err = SetBlobJSON(tag, siteMap)
	if err != nil {
		t.Error(err)
	}
	x, err := GetBlob(tag)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, x)

}

func TestSharePointSitemapProducts(t *testing.T) {
	siteMap, err := SharePointSitemap("https://christianiabpos.sharepoint.com/sites/IssuerProducts")
	if err != nil {
		t.Error(err)
	}
	io.WriteFile("IssuerProducts.sitemap.json", siteMap)
	assert.Greater(t, len(*&siteMap.Menu), 0)
}
